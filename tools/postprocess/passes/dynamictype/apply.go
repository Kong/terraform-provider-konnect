package dynamictype

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"
)

// applyModel retypes the model field to types.Dynamic. (edit b)
func applyModel(s *modelSite) error {
	s.field.Type = ident("Dynamic", typesPkg)
	return nil
}

// applySchema rewrites schema.StringAttribute{CustomType: marker, ...} into a
// schema.DynamicAttribute, translating any Default and dropping the kind-typed
// CustomType and Validators. (edit a)
func applySchema(s *schemaSite) error {
	typeIdent, ok := s.lit.Type.(*dst.Ident)
	if !ok {
		return fmt.Errorf("schema attribute has unexpected type expression %T", s.lit.Type)
	}
	typeIdent.Name = "DynamicAttribute"

	kept := make([]dst.Expr, 0, len(s.lit.Elts))
	for _, el := range s.lit.Elts {
		kv, ok := el.(*dst.KeyValueExpr)
		if !ok {
			kept = append(kept, el)
			continue
		}
		key, ok := kv.Key.(*dst.Ident)
		if !ok {
			kept = append(kept, el)
			continue
		}
		switch key.Name {
		case "CustomType", "Validators":
			// Drop: a dynamic attribute has no string custom type, and a
			// kind-typed validator list is invalid on it.
			continue
		case "Default":
			nv, err := translateDefault(kv.Value)
			if err != nil {
				return err
			}
			kv.Value = nv
			kept = append(kept, kv)
		default:
			kept = append(kept, kv)
		}
	}
	s.lit.Elts = kept
	return nil
}

// translateDefault converts a kind-typed static default into a dynamic default.
func translateDefault(v dst.Expr) (dst.Expr, error) {
	c, ok := v.(*dst.CallExpr)
	if !ok {
		return nil, fmt.Errorf("unsupported Default expression %T", v)
	}
	fn, ok := c.Fun.(*dst.Ident)
	if !ok || len(c.Args) != 1 {
		return nil, fmt.Errorf("unsupported Default call shape")
	}
	arg, ok := dst.Clone(c.Args[0]).(dst.Expr)
	if !ok {
		return nil, fmt.Errorf("unsupported Default argument %T", c.Args[0])
	}

	var inner dst.Expr
	switch fn.Name {
	case "StaticInt64", "StaticFloat64":
		inner = call(ident("NumberValue", typesPkg), call(ident("NewFloat", bigPkg), arg))
	case "StaticString":
		inner = call(ident("StringValue", typesPkg), arg)
	case "StaticBool":
		inner = call(ident("BoolValue", typesPkg), arg)
	default:
		return nil, fmt.Errorf("no dynamic translation for default %q", fn.Name)
	}
	return call(
		ident("StaticValue", dynamicDefaultPkg),
		call(ident("DynamicValue", typesPkg), inner),
	), nil
}

// applyRefresh replaces the marker's ValueFromString block with a single
// `<model> = dynamic.ToFramework(<src>)` assignment. (edit c)
func applyRefresh(s *refreshSite) error {
	stmt := &dst.AssignStmt{
		Lhs: []dst.Expr{s.lhs},
		Tok: token.ASSIGN,
		Rhs: []dst.Expr{call(ident("ToFramework", dynamicPkg), s.src)},
	}
	stmt.Decs.Before = s.leadingBefore
	stmt.Decs.Start = s.leading
	stmt.Decs.After = dst.NewLine
	return splice(s.block, s.start, s.count, stmt)
}

// applyBuild replaces the primitive-builder block with a single
// `<var> := dynamic.ToAny(<model>)` assignment. (edit d)
func applyBuild(s *buildSite) error {
	stmt := &dst.AssignStmt{
		Lhs: []dst.Expr{ident(s.varName, "")},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{call(ident("ToAny", dynamicPkg), s.model)},
	}
	stmt.Decs.Before = s.leadingBefore
	stmt.Decs.Start = s.leading
	stmt.Decs.After = dst.NewLine
	return splice(s.block, s.start, s.count, stmt)
}

// applySDKField retypes the SDK struct field that receives the dynamic value to
// any, and drops the now-meaningless default tag.
func applySDKField(s *sdkFieldSite) error {
	s.field.Type = ident("any", "")
	return nil
}

// applySDKGetter retypes the SDK getter's result to any.
func applySDKGetter(s *sdkGetterSite) error {
	if s.fn.Type == nil || s.fn.Type.Results == nil || len(s.fn.Type.Results.List) != 1 {
		return fmt.Errorf("getter %s has unexpected result list", s.fn.Name.Name)
	}
	s.fn.Type.Results.List[0].Type = ident("any", "")
	return nil
}

// splice replaces block.List[start:start+count] with stmt.
func splice(block *dst.BlockStmt, start, count int, stmt dst.Stmt) error {
	if start < 0 || start+count > len(block.List) {
		return fmt.Errorf("splice out of range: start=%d count=%d len=%d", start, count, len(block.List))
	}
	tail := append([]dst.Stmt{stmt}, block.List[start+count:]...)
	block.List = append(block.List[:start:start], tail...)
	return nil
}
