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

	// 2. Provider schema + refresh + build. Both the refresh sites (SDK source
	//    expr) and the build sites (model field that flows into the SDK struct)
	//    reveal which SDK fields feed a dynamic value.
	sdkRefs := map[fieldRef]bool{}
	buildNames := map[[2]string]bool{}
	detectSchema(providerD, &targets)
	if err := detectRefresh(providerD, &targets, sdkRefs); err != nil {
		return nil, err
	}
	detectBuild(providerD, &targets, marked, buildNames)

	// 3. SDK struct fields + getters feeding a dynamic value. A field qualifies
	//    if a refresh names it directly, or a build site flows a dynamic value
	//    into the identically named SDK struct field (Speakeasy mirrors the
	//    model and SDK struct/field names).
	detectSDK(sharedD, &targets, sdkRefs, buildNames)

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

// detectBuild locates the primitive-builder block for each marked model field.
// Speakeasy emits one of two shapes depending on whether the SDK field is a
// pointer or a value:
//
//	// pointer (e.g. redis.port -> *string)
//	port1 := new(string)
//	if !M.IsUnknown() && !M.IsNull() { *port1 = M.ValueString() } else { port1 = nil }
//
//	// value (e.g. acl rule resource_names -> string)
//	var resourceNames string
//	resourceNames = M.ValueString()
//
// Both are replaced by `<var> := dynamic.ToAny(M)`. Each matched site records
// its (struct, field) name in buildNames so the SDK field can be retyped even
// when there is no refresh site (nested slices refresh via a prior-state copy).
func detectBuild(p *decorator.Package, targets *[]engine.Target, marked map[fieldRef]bool, buildNames map[[2]string]bool) {
	forEachFile(p, func(af *ast.File, file *dst.File) {
		ast.Inspect(af, func(n ast.Node) bool {
			blk, ok := n.(*ast.BlockStmt)
			if !ok {
				return true
			}
			for i := 0; i < len(blk.List); i++ {
				varName, model, start, count := matchBuilder(p, blk.List, i, marked)
				if model == nil {
					continue
				}
				msel := model.(*ast.SelectorExpr)
				ref, _ := fieldRefOf(p.TypesInfo, msel.X, msel.Sel.Name)
				buildNames[[2]string{ref.typ, ref.field}] = true

				dstBlk, ok := engine.DstFor(p, blk).(*dst.BlockStmt)
				if !ok {
					continue
				}
				site := &buildSite{
					block:   dstBlk,
					varName: varName,
					model:   dstExpr(p, model),
					start:   start,
					count:   count,
				}
				captureLeading(dstBlk, start, &site.leadingBefore, &site.leading)
				*targets = append(*targets, target("build:"+msel.Sel.Name, p, file, kindBuild, site))
			}
			return true
		})
	})
}

// matchBuilder tries to match a builder shape anchored at blk.List[i]. On a
// match it returns the result variable name, the marked model expression, and
// the [start,count) span of statements to replace; otherwise model is nil.
func matchBuilder(p *decorator.Package, list []ast.Stmt, i int, marked map[fieldRef]bool) (varName string, model ast.Expr, start, count int) {
	// Pointer shape: `X := new(T)` followed by an if writing `*X = M.ValueString()`.
	if as, ok := list[i].(*ast.AssignStmt); ok && as.Tok == token.DEFINE &&
		len(as.Lhs) == 1 && len(as.Rhs) == 1 && isNewCall(as.Rhs[0]) {
		if id, ok := as.Lhs[0].(*ast.Ident); ok && i+1 < len(list) {
			if ifs, ok := list[i+1].(*ast.IfStmt); ok {
				if m := buildModelExpr(ifs, id.Name); m != nil && markedSelector(p, m, marked) {
					return id.Name, m, i, 2
				}
			}
		}
	}

	// Value shape: an assignment `X = M.ValueString()` (or `X := …`), optionally
	// preceded by a `var X T` declaration.
	if as, ok := list[i].(*ast.AssignStmt); ok && len(as.Lhs) == 1 && len(as.Rhs) == 1 {
		id, ok := as.Lhs[0].(*ast.Ident)
		if !ok {
			return "", nil, 0, 0
		}
		m := valueStringReceiver(as.Rhs[0])
		if m == nil || !markedSelector(p, m, marked) {
			return "", nil, 0, 0
		}
		if i > 0 && declStmtDeclares(list[i-1], id.Name) {
			return id.Name, m, i - 1, 2
		}
		return id.Name, m, i, 1
	}

	return "", nil, 0, 0
}

// markedSelector reports whether e is a selector on a marked model field.
func markedSelector(p *decorator.Package, e ast.Expr, marked map[fieldRef]bool) bool {
	sel, ok := e.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	ref, ok := fieldRefOf(p.TypesInfo, sel.X, sel.Sel.Name)
	return ok && marked[ref]
}

// valueStringReceiver returns M if e is `M.ValueString()`, else nil.
func valueStringReceiver(e ast.Expr) ast.Expr {
	c, ok := e.(*ast.CallExpr)
	if !ok || len(c.Args) != 0 {
		return nil
	}
	sel, ok := c.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "ValueString" {
		return nil
	}
	return sel.X
}

// declStmtDeclares reports whether stmt is `var <name> T`.
func declStmtDeclares(stmt ast.Stmt, name string) bool {
	ds, ok := stmt.(*ast.DeclStmt)
	if !ok {
		return false
	}
	gd, ok := ds.Decl.(*ast.GenDecl)
	if !ok || gd.Tok != token.VAR {
		return false
	}
	for _, spec := range gd.Specs {
		vs, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		for _, n := range vs.Names {
			if n.Name == name {
				return true
			}
		}
	}
	return false
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

func detectSDK(p *decorator.Package, targets *[]engine.Target, sdkRefs map[fieldRef]bool, buildNames map[[2]string]bool) {
	if len(sdkRefs) == 0 && len(buildNames) == 0 {
		return
	}
	// wants reports whether the SDK field (struct, name) feeds a dynamic value.
	wants := func(structName, name string) bool {
		return sdkRefs[fieldRef{pkgPath: p.PkgPath, typ: structName, field: name}] ||
			buildNames[[2]string{structName, name}]
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
							if !wants(ts.Name.Name, nm.Name) {
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
				const getterPrefix = "Get"
				if recv == "" || len(d.Name.Name) <= len(getterPrefix) || d.Name.Name[:len(getterPrefix)] != getterPrefix {
					continue
				}
				if wants(recv, d.Name.Name[len(getterPrefix):]) {
					if fd, ok := engine.DstFor(p, d).(*dst.FuncDecl); ok {
						*targets = append(*targets, target("sdkgetter:"+d.Name.Name, p, file, kindSDKGetter, &sdkGetterSite{fn: fd}))
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
