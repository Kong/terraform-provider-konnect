package dynamictype

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/kong/terraform-provider-konnect/v3/tools/postprocess/engine"
)

// Detect resolves every marker-flagged field and the four generated locations
// it touches (model, schema, refresh, build) plus the SDK field that feeds it.
// All resolution uses go/types, never string matching.
func (*Pass) Detect(b *engine.Bundle) ([]engine.Target, error) {
	var targets []engine.Target

	typesD := b.Pkg(engine.TypesPkg)
	providerD := b.Pkg(engine.ProviderPkg)
	sharedD := b.Pkg(engine.SharedPkg)
	if typesD == nil || providerD == nil || sharedD == nil {
		return nil, fmt.Errorf("required packages not loaded")
	}

	// 1. Model fields carrying the marker value type; record their (struct,field)
	//    so the build edit can correlate its primitive-builder block.
	marked := map[fieldRef]bool{}
	detectModel(typesD, &targets, marked)

	// 2. Provider schema + refresh + build. The refresh sites reveal which SDK
	//    fields feed a dynamic value.
	sdkRefs := map[fieldRef]bool{}
	detectSchema(providerD, &targets)
	if err := detectRefresh(providerD, &targets, sdkRefs); err != nil {
		return nil, err
	}
	detectBuild(providerD, &targets, marked)

	// 3. SDK struct fields + getters feeding a dynamic value.
	detectSDK(sharedD, &targets, sdkRefs)

	return targets, nil
}

// forEachFile iterates a package's syntax with its decorated file in hand.
func forEachFile(p *decorator.Package, fn func(af *ast.File, file *dst.File)) {
	for _, af := range p.Package.Syntax {
		dstFile, _ := engine.DstFor(p, af).(*dst.File)
		if dstFile == nil {
			continue
		}
		fn(af, dstFile)
	}
}

func detectModel(p *decorator.Package, targets *[]engine.Target, marked map[fieldRef]bool) {
	forEachFile(p, func(af *ast.File, file *dst.File) {
		for _, decl := range af.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, f := range st.Fields.List {
					if !isMarkerNamed(p.TypesInfo.TypeOf(f.Type), markerValue) {
						continue
					}
					df, ok := engine.DstFor(p, f).(*dst.Field)
					if !ok {
						continue
					}
					for _, nm := range f.Names {
						marked[fieldRef{pkgPath: p.PkgPath, typ: ts.Name.Name, field: nm.Name}] = true
					}
					name := fmt.Sprintf("model:%s.%s", ts.Name.Name, tfsdkTag(f))
					*targets = append(*targets, target(name, p, file, kindModel, &modelSite{field: df}))
				}
			}
		}
	})
}

func detectSchema(p *decorator.Package, targets *[]engine.Target) {
	forEachFile(p, func(af *ast.File, file *dst.File) {
		ast.Inspect(af, func(n ast.Node) bool {
			kv, ok := n.(*ast.KeyValueExpr)
			if !ok {
				return true
			}
			name, ok := stringLit(kv.Key)
			if !ok {
				return true
			}
			lit, ok := kv.Value.(*ast.CompositeLit)
			if !ok || !hasMarkerCustomType(p, lit) {
				return true
			}
			if dl, ok := engine.DstFor(p, lit).(*dst.CompositeLit); ok {
				*targets = append(*targets, target("schema:"+name, p, file, kindSchema, &schemaSite{lit: dl}))
			}
			return true
		})
	})
}

// hasMarkerCustomType reports whether a composite literal has a
// `CustomType: referenceabletypes.ReferenceableStringType{}` element.
func hasMarkerCustomType(p *decorator.Package, lit *ast.CompositeLit) bool {
	for _, el := range lit.Elts {
		kv, ok := el.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := kv.Key.(*ast.Ident)
		if !ok || key.Name != "CustomType" {
			continue
		}
		if isMarkerNamed(p.TypesInfo.TypeOf(kv.Value), markerType) {
			return true
		}
	}
	return false
}

func detectRefresh(p *decorator.Package, targets *[]engine.Target, sdkRefs map[fieldRef]bool) error {
	var detectErr error
	forEachFile(p, func(af *ast.File, file *dst.File) {
		ast.Inspect(af, func(n ast.Node) bool {
			blk, ok := n.(*ast.BlockStmt)
			if !ok {
				return true
			}
			for i, stmt := range blk.List {
				as, ok := stmt.(*ast.AssignStmt)
				if !ok || len(as.Rhs) != 1 {
					continue
				}
				c, ok := as.Rhs[0].(*ast.CallExpr)
				if !ok {
					continue
				}
				sel, ok := c.Fun.(*ast.SelectorExpr)
				if !ok || sel.Sel.Name != "ValueFromString" {
					continue
				}
				if !isMarkerNamed(p.TypesInfo.TypeOf(sel.X), markerType) {
					continue
				}
				if len(c.Args) < 2 {
					continue
				}
				src := unwrapConv(c.Args[1])

				lhs, count := refreshAssignment(p, blk.List, i)
				if lhs == nil {
					detectErr = fmt.Errorf("refresh: could not find model assignment after ValueFromString")
					return false
				}

				dstBlk, ok := engine.DstFor(p, blk).(*dst.BlockStmt)
				if !ok {
					continue
				}
				site := &refreshSite{
					block: dstBlk,
					lhs:   dstExpr(p, lhs),
					src:   dstExpr(p, src),
					start: i,
					count: count,
				}
				captureLeading(dstBlk, i, &site.leadingBefore, &site.leading)

				// The SDK field feeding the dynamic value (e.g. resp.Config.Redis.Port).
				if sse, ok := src.(*ast.SelectorExpr); ok {
					if ref, ok := fieldRefOf(p.TypesInfo, sse.X, sse.Sel.Name); ok {
						sdkRefs[ref] = true
					}
				}

				name := "refresh"
				if sse, ok := src.(*ast.SelectorExpr); ok {
					name = "refresh:" + sse.Sel.Name
				}
				*targets = append(*targets, target(name, p, file, kindRefresh, site))
			}
			return true
		})
	})
	return detectErr
}

// refreshAssignment finds the `<model> = valuable.(Referenceable...)` statement
// following the ValueFromString assignment at index start, returning its LHS and
// the number of statements (inclusive) to replace.
func refreshAssignment(p *decorator.Package, list []ast.Stmt, start int) (ast.Expr, int) {
	for j := start + 1; j < len(list) && j <= start+3; j++ {
		as, ok := list[j].(*ast.AssignStmt)
		if !ok || len(as.Rhs) != 1 || len(as.Lhs) != 1 {
			continue
		}
		ta, ok := as.Rhs[0].(*ast.TypeAssertExpr)
		if !ok || !isMarkerNamed(p.TypesInfo.TypeOf(ta.Type), markerValue) {
			continue
		}
		return as.Lhs[0], j - start + 1
	}
	return nil, 0
}

func detectBuild(p *decorator.Package, targets *[]engine.Target, marked map[fieldRef]bool) {
	forEachFile(p, func(af *ast.File, file *dst.File) {
		ast.Inspect(af, func(n ast.Node) bool {
			blk, ok := n.(*ast.BlockStmt)
			if !ok {
				return true
			}
			for i := 0; i+1 < len(blk.List); i++ {
				as, ok := blk.List[i].(*ast.AssignStmt)
				if !ok || as.Tok != token.DEFINE || len(as.Lhs) != 1 || len(as.Rhs) != 1 {
					continue
				}
				varIdent, ok := as.Lhs[0].(*ast.Ident)
				if !ok || !isNewCall(as.Rhs[0]) {
					continue
				}
				ifs, ok := blk.List[i+1].(*ast.IfStmt)
				if !ok {
					continue
				}
				model := buildModelExpr(ifs, varIdent.Name)
				if model == nil {
					continue
				}
				msel, ok := model.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				ref, ok := fieldRefOf(p.TypesInfo, msel.X, msel.Sel.Name)
				if !ok || !marked[ref] {
					continue
				}

				dstBlk, ok := engine.DstFor(p, blk).(*dst.BlockStmt)
				if !ok {
					continue
				}
				site := &buildSite{
					block:   dstBlk,
					varName: varIdent.Name,
					model:   dstExpr(p, model),
					start:   i,
					count:   2,
				}
				captureLeading(dstBlk, i, &site.leadingBefore, &site.leading)
				*targets = append(*targets, target("build:"+msel.Sel.Name, p, file, kindBuild, site))
			}
			return true
		})
	})
}

// isNewCall reports whether e is `new(<type>)`.
func isNewCall(e ast.Expr) bool {
	c, ok := e.(*ast.CallExpr)
	if !ok || len(c.Args) != 1 {
		return false
	}
	id, ok := c.Fun.(*ast.Ident)
	return ok && id.Name == "new"
}

// buildModelExpr finds `*<varName> = <model>.ValueString()` inside an if body
// and returns the <model> receiver expression.
func buildModelExpr(ifs *ast.IfStmt, varName string) ast.Expr {
	var found ast.Expr
	ast.Inspect(ifs, func(n ast.Node) bool {
		as, ok := n.(*ast.AssignStmt)
		if !ok || len(as.Lhs) != 1 || len(as.Rhs) != 1 {
			return true
		}
		star, ok := as.Lhs[0].(*ast.StarExpr)
		if !ok {
			return true
		}
		id, ok := star.X.(*ast.Ident)
		if !ok || id.Name != varName {
			return true
		}
		c, ok := as.Rhs[0].(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := c.Fun.(*ast.SelectorExpr)
		if !ok || sel.Sel.Name != "ValueString" {
			return true
		}
		found = sel.X
		return false
	})
	return found
}

func detectSDK(p *decorator.Package, targets *[]engine.Target, sdkRefs map[fieldRef]bool) {
	if len(sdkRefs) == 0 {
		return
	}
	forEachFile(p, func(af *ast.File, file *dst.File) {
		for _, decl := range af.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				if d.Tok != token.TYPE {
					continue
				}
				for _, spec := range d.Specs {
					ts, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					st, ok := ts.Type.(*ast.StructType)
					if !ok {
						continue
					}
					for _, f := range st.Fields.List {
						for _, nm := range f.Names {
							ref := fieldRef{pkgPath: p.PkgPath, typ: ts.Name.Name, field: nm.Name}
							if !sdkRefs[ref] {
								continue
							}
							if df, ok := engine.DstFor(p, f).(*dst.Field); ok {
								*targets = append(*targets, target("sdkfield:"+ts.Name.Name+"."+nm.Name, p, file, kindSDKField, &sdkFieldSite{field: df}))
							}
						}
					}
				}
			case *ast.FuncDecl:
				if d.Recv == nil || len(d.Recv.List) != 1 {
					continue
				}
				recv := recvTypeName(d.Recv.List[0].Type)
				for ref := range sdkRefs {
					if ref.typ == recv && d.Name.Name == "Get"+ref.field {
						if fd, ok := engine.DstFor(p, d).(*dst.FuncDecl); ok {
							*targets = append(*targets, target("sdkgetter:"+d.Name.Name, p, file, kindSDKGetter, &sdkGetterSite{fn: fd}))
						}
					}
				}
			}
		}
	})
}

// recvTypeName returns the receiver type name from `(r *T)` or `(r T)`.
func recvTypeName(e ast.Expr) string {
	if star, ok := e.(*ast.StarExpr); ok {
		e = star.X
	}
	if id, ok := e.(*ast.Ident); ok {
		return id.Name
	}
	return ""
}

// captureLeading records the Before space and Start decorations of the first
// statement being replaced so the replacement preserves spacing/comments.
func captureLeading(blk *dst.BlockStmt, idx int, before *dst.SpaceType, start *dst.Decorations) {
	if idx < 0 || idx >= len(blk.List) {
		return
	}
	d := blk.List[idx].Decorations()
	*before = d.Before
	*start = d.Start
}
