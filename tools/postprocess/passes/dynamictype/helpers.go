package dynamictype

import (
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/kong/terraform-provider-konnect/v3/tools/postprocess/engine"
)

// fieldRef identifies a struct field by its declaring named type and field name.
type fieldRef struct {
	pkgPath string
	typ     string
	field   string
}

// namedOf strips pointers/aliases and returns the underlying *types.Named, or nil.
func namedOf(t types.Type) *types.Named {
	for {
		switch tt := t.(type) {
		case *types.Pointer:
			t = tt.Elem()
		case *types.Named:
			return tt
		case *types.Alias:
			t = types.Unalias(tt)
		default:
			return nil
		}
	}
}

// isMarkerNamed reports whether t resolves to referenceabletypes.<name>.
func isMarkerNamed(t types.Type, name string) bool {
	n := namedOf(t)
	if n == nil || n.Obj() == nil || n.Obj().Pkg() == nil {
		return false
	}
	return n.Obj().Pkg().Path() == engine.ReferencePkg && n.Obj().Name() == name
}

// fieldRefOf resolves the named type of expr and pairs it with sel to form a
// fieldRef. expr is the receiver of a field selector (e.g. resp.Config.Redis),
// sel is the field name (e.g. Port).
func fieldRefOf(info *types.Info, expr ast.Expr, sel string) (fieldRef, bool) {
	tv, ok := info.Types[expr]
	if !ok {
		return fieldRef{}, false
	}
	n := namedOf(tv.Type)
	if n == nil || n.Obj() == nil || n.Obj().Pkg() == nil {
		return fieldRef{}, false
	}
	return fieldRef{pkgPath: n.Obj().Pkg().Path(), typ: n.Obj().Name(), field: sel}, true
}

// unwrapConv unwraps a types.String[Pointer]Value(x) conversion to x.
func unwrapConv(e ast.Expr) ast.Expr {
	call, ok := e.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return e
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return e
	}
	if strings.HasPrefix(sel.Sel.Name, "String") && strings.HasSuffix(sel.Sel.Name, "Value") {
		return call.Args[0]
	}
	return e
}

// stringLit returns the unquoted value of a string-literal expression.
func stringLit(e ast.Expr) (string, bool) {
	bl, ok := e.(*ast.BasicLit)
	if !ok || bl.Kind.String() != "STRING" {
		return "", false
	}
	s, err := strconv.Unquote(bl.Value)
	if err != nil {
		return "", false
	}
	return s, true
}

// tfsdkTag extracts the tfsdk tag value from a struct field tag.
func tfsdkTag(f *ast.Field) string {
	if f.Tag == nil {
		return ""
	}
	raw, err := strconv.Unquote(f.Tag.Value)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(structTag(raw).Get("tfsdk")))
}

// structTag mirrors reflect.StructTag.Get without importing reflect for clarity.
type structTag string

func (t structTag) Get(key string) string {
	tag := string(t)
	for tag != "" {
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := tag[:i]
		tag = tag[i+1:]
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := tag[:i+1]
		tag = tag[i+1:]
		if name == key {
			value, err := strconv.Unquote(qvalue)
			if err != nil {
				return ""
			}
			return value
		}
	}
	return ""
}

// dstExpr returns a clone of the dst expression for an ast expression so it can
// be re-parented into a freshly built statement.
func dstExpr(p *decorator.Package, e ast.Expr) dst.Expr {
	n := engine.DstFor(p, e)
	if n == nil {
		return nil
	}
	expr, ok := dst.Clone(n).(dst.Expr)
	if !ok {
		return nil
	}
	return expr
}

// ident builds a (possibly package-qualified) dst identifier. An empty path
// yields a local identifier.
func ident(name, path string) *dst.Ident { return &dst.Ident{Name: name, Path: path} }

// call builds a call expression.
func call(fun dst.Expr, args ...dst.Expr) *dst.CallExpr {
	return &dst.CallExpr{Fun: fun, Args: args}
}
