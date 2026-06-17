package dynamictype

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"testing"

	"github.com/dave/dst"
	"github.com/stretchr/testify/require"
)

// numericLit builds an integer basic-literal dst expression.
func numericLit(v string) *dst.BasicLit { return &dst.BasicLit{Kind: token.INT, Value: v} }

func newFset() *token.FileSet { return token.NewFileSet() }

func mustParseExpr(t *testing.T, src string) ast.Expr {
	t.Helper()
	e, err := parser.ParseExpr(src)
	require.NoError(t, err)
	return e
}

func exprString(e ast.Expr) string {
	var buf bytes.Buffer
	_ = printer.Fprint(&buf, token.NewFileSet(), e)
	return buf.String()
}

func TestTranslateDefault(t *testing.T) {
	tests := []struct {
		name    string
		fn      string
		arg     dst.Expr
		want    string // expected innermost constructor name
		wantErr bool
	}{
		{"int64", "StaticInt64", numericLit("6379"), "NumberValue", false},
		{"float64", "StaticFloat64", &dst.BasicLit{Kind: token.FLOAT, Value: "1.5"}, "NumberValue", false},
		{"string", "StaticString", &dst.BasicLit{Kind: token.STRING, Value: `"6379"`}, "StringValue", false},
		{"bool", "StaticBool", dst.NewIdent("true"), "BoolValue", false},
		{"unknown", "StaticDuration", numericLit("1"), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := call(ident(tt.fn, "irrelevant"), tt.arg)
			got, err := translateDefault(in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// Outer call must be dynamicdefault.StaticValue(types.DynamicValue(<inner>(...))).
			outer := got.(*dst.CallExpr)
			require.Equal(t, "StaticValue", outer.Fun.(*dst.Ident).Name)
			require.Equal(t, dynamicDefaultPkg, outer.Fun.(*dst.Ident).Path)

			dynamicVal := outer.Args[0].(*dst.CallExpr)
			require.Equal(t, "DynamicValue", dynamicVal.Fun.(*dst.Ident).Name)

			inner := dynamicVal.Args[0].(*dst.CallExpr)
			require.Equal(t, tt.want, inner.Fun.(*dst.Ident).Name)
			require.Equal(t, typesPkg, inner.Fun.(*dst.Ident).Path)

			if tt.want == "NumberValue" {
				bigCall := inner.Args[0].(*dst.CallExpr)
				require.Equal(t, "NewFloat", bigCall.Fun.(*dst.Ident).Name)
				require.Equal(t, bigPkg, bigCall.Fun.(*dst.Ident).Path)
			}
		})
	}
}

func TestStructTagGet(t *testing.T) {
	tag := structTag(`tfsdk:"port" json:"port,omitempty"`)
	require.Equal(t, "port", tag.Get("tfsdk"))
	require.Equal(t, "port,omitempty", tag.Get("json"))
	require.Equal(t, "", tag.Get("missing"))
}

func TestUnwrapConv(t *testing.T) {
	// types.StringPointerValue(resp.Config.Redis.Port) -> resp.Config.Redis.Port
	e := mustParseExpr(t, "types.StringPointerValue(resp.Config.Redis.Port)")
	got := unwrapConv(e)
	require.Equal(t, "resp.Config.Redis.Port", exprString(got))

	// A non-conversion expression is returned unchanged.
	plain := mustParseExpr(t, "resp.Config.Redis.Port")
	require.Equal(t, plain, unwrapConv(plain))
}

func TestIsNewCall(t *testing.T) {
	require.True(t, isNewCall(mustParseExpr(t, "new(string)")))
	require.False(t, isNewCall(mustParseExpr(t, "make([]int, 0)")))
	require.False(t, isNewCall(mustParseExpr(t, "foo()")))
}

func TestRecvTypeName(t *testing.T) {
	require.Equal(t, "Redis", recvTypeName(mustParseExpr(t, "(*Redis)").(*ast.ParenExpr).X))
	require.Equal(t, "Redis", recvTypeName(mustParseExpr(t, "Redis")))
}

func TestBuildModelExpr(t *testing.T) {
	src := `package p
func f() {
	port1 := new(string)
	if !r.Config.Redis.Port.IsUnknown() && !r.Config.Redis.Port.IsNull() {
		*port1 = r.Config.Redis.Port.ValueString()
	} else {
		port1 = nil
	}
}`
	fset := newFset()
	file, err := parser.ParseFile(fset, "p.go", src, 0)
	require.NoError(t, err)

	var ifs *ast.IfStmt
	ast.Inspect(file, func(n ast.Node) bool {
		if s, ok := n.(*ast.IfStmt); ok {
			ifs = s
			return false
		}
		return true
	})
	require.NotNil(t, ifs)

	model := buildModelExpr(ifs, "port1")
	require.NotNil(t, model)
	require.Equal(t, "r.Config.Redis.Port", exprString(model))

	// Wrong variable name yields no match.
	require.Nil(t, buildModelExpr(ifs, "other"))
}

func TestStringLit(t *testing.T) {
	s, ok := stringLit(mustParseExpr(t, `"port"`))
	require.True(t, ok)
	require.Equal(t, "port", s)

	_, ok = stringLit(mustParseExpr(t, "42"))
	require.False(t, ok)
}
