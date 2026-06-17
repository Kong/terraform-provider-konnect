package engine

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"golang.org/x/tools/go/packages"
)

// Load loads the provider packages the post-processor operates on, with full
// type information, into decorated (dst) syntax trees.
//
// Only the packages that may be rewritten are requested as initial patterns, so
// their syntax is decorated; their import graph is type-checked from source
// (NeedDeps) so identifiers in one package resolve to the same named types as
// the package that declares them.
func Load(dir string) (*Bundle, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedTypesSizes |
			packages.NeedSyntax,
		Dir:  dir,
		Fset: token.NewFileSet(),
	}

	pkgs, err := decorator.Load(cfg, ProviderPkg, TypesPkg, SharedPkg)
	if err != nil {
		return nil, err
	}

	b := &Bundle{
		byPath:    map[string]*decorator.Package{},
		fileOwner: map[*dst.File]*decorator.Package{},
		touched:   map[*dst.File]bool{},
	}

	var loadErrs []string
	for _, p := range pkgs {
		for _, e := range p.Errors {
			loadErrs = append(loadErrs, e.Error())
		}
		b.Pkgs = append(b.Pkgs, p)
		b.byPath[p.PkgPath] = p
		for _, f := range p.Syntax {
			b.fileOwner[f] = p
		}
	}
	if len(loadErrs) > 0 {
		return nil, fmt.Errorf("packages loaded with errors:\n%s", join(loadErrs))
	}

	for _, path := range []string{ProviderPkg, TypesPkg, SharedPkg} {
		if b.byPath[path] == nil {
			return nil, fmt.Errorf("required package not loaded: %s", path)
		}
	}

	return b, nil
}

func join(ss []string) string {
	out := ""
	for i, s := range ss {
		if i > 0 {
			out += "\n"
		}
		out += "  " + s
	}
	return out
}
