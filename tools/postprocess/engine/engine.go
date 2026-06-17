// Package engine is the shared machinery of the post-processing layer: it loads
// the generated provider packages into decorated syntax trees (dst) with full
// go/types information, runs an ordered pipeline of Passes over them, then
// formats, verifies and writes the result.
//
// The engine is deliberately decoupled from any individual workflow. Adding a
// new post-processing workflow means implementing the Pass interface and
// registering it; the loader, runner, formatter and verifier are untouched.
package engine

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"os/exec"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/guess"
	"golang.org/x/tools/imports"
)

// ModulePath is the Go module path of the provider.
const ModulePath = "github.com/kong/terraform-provider-konnect/v3"

// Package import paths the engine loads and may rewrite.
const (
	ProviderPkg  = ModulePath + "/internal/provider"
	TypesPkg     = ModulePath + "/internal/provider/types"
	SharedPkg    = ModulePath + "/internal/sdk/models/shared"
	ReferencePkg = ModulePath + "/internal/referenceabletypes"
)

// Bundle is the loaded, mutable view of the provider package(s): decorated dst
// trees plus the *types.Info needed to resolve identifiers to the marker type.
type Bundle struct {
	// Pkgs are the explicitly loaded packages, in load order.
	Pkgs []*decorator.Package
	// byPath indexes Pkgs by import path.
	byPath map[string]*decorator.Package
	// fileOwner maps every decorated file back to its package.
	fileOwner map[*dst.File]*decorator.Package
	// touched records files mutated by a pass; only these are rewritten.
	touched map[*dst.File]bool
}

// Pkg returns the loaded package for an import path, or nil.
func (b *Bundle) Pkg(path string) *decorator.Package { return b.byPath[path] }

// OwnerOf returns the package that owns a decorated file.
func (b *Bundle) OwnerOf(f *dst.File) *decorator.Package { return b.fileOwner[f] }

// Touch marks a file as modified so it will be re-rendered and written.
func (b *Bundle) Touch(f *dst.File) { b.touched[f] = true }

// DstFor maps a go/ast node from a package's syntax to its decorated dst node.
func DstFor(p *decorator.Package, n ast.Node) dst.Node {
	if n == nil {
		return nil
	}
	return p.Decorator.Dst.Nodes[n]
}

// Pass is the extensibility seam. A pass detects the units of work it will act
// on, then mutates the bundle's dst trees one target at a time. Apply is a pure
// AST→AST transform: it must not perform I/O.
type Pass interface {
	Name() string
	Detect(b *Bundle) ([]Target, error)
	Apply(b *Bundle, t Target) error
}

// Registry holds an ordered list of passes.
type Registry struct{ passes []Pass }

// Register appends a pass to the pipeline.
func (r *Registry) Register(p Pass) { r.passes = append(r.passes, p) }

// Passes returns the registered passes in order.
func (r *Registry) Passes() []Pass { return r.passes }

// Run loads the provider packages, applies every registered pass, then formats,
// writes and compile-verifies the result. A non-nil error leaves the working
// tree as-is for triage.
func Run(dir string, reg *Registry) error {
	b, err := Load(dir)
	if err != nil {
		return fmt.Errorf("load: %w", err)
	}

	for _, p := range reg.passes {
		targets, err := p.Detect(b)
		if err != nil {
			return fmt.Errorf("pass %q detect: %w", p.Name(), err)
		}
		for _, t := range targets {
			if err := p.Apply(b, t); err != nil {
				return fmt.Errorf("pass %q apply (%s): %w", p.Name(), t.Name, err)
			}
		}
		fmt.Printf("postprocess: pass %q transformed %d target(s)\n", p.Name(), len(targets))
	}

	written, err := b.render()
	if err != nil {
		return fmt.Errorf("render: %w", err)
	}
	for _, f := range written {
		fmt.Printf("postprocess: wrote %s\n", f)
	}

	if len(written) == 0 {
		fmt.Println("postprocess: no changes")
		return nil
	}

	return verify(dir)
}

// render restores every touched dst file to source, runs goimports, writes it
// back, and returns the list of written file paths.
func (b *Bundle) render() ([]string, error) {
	var written []string
	for f, ok := range b.touched {
		if !ok {
			continue
		}
		owner := b.fileOwner[f]
		path := owner.Decorator.Filenames[f]

		formatted, err := RenderFile(owner, f)
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(path, formatted, 0o644); err != nil {
			return nil, fmt.Errorf("write %s: %w", path, err)
		}
		written = append(written, path)
	}
	return written, nil
}

// RenderFile restores a (possibly mutated) decorated file to formatted source.
// It restores with import management so dst.Idents carrying a package Path
// expand back to qualified selectors and the import block is rebuilt, then runs
// goimports as the final authority on imports — adding the helper/framework
// packages introduced by edits and dropping the now-unused kind-typed ones.
func RenderFile(owner *decorator.Package, f *dst.File) ([]byte, error) {
	path := owner.Decorator.Filenames[f]
	r := decorator.NewRestorerWithImports(owner.PkgPath, guess.New())
	var buf bytes.Buffer
	if err := r.Fprint(&buf, f); err != nil {
		return nil, fmt.Errorf("restore %s: %w", path, err)
	}
	formatted, err := imports.Process(path, buf.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("goimports %s: %w", path, err)
	}
	return formatted, nil
}

// verify runs `go build ./...` as a hard correctness gate: post-processed output
// must always compile, so a break turns the runner (and CI) red. `go vet` is not
// used as a gate because the generated SDK does not pass it cleanly (it carries
// unexported json-tagged fields by design, predating this layer).
func verify(dir string) error {
	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go build ./... failed:\n%s", out)
	}
	fmt.Println("postprocess: go build ./... passed")
	return nil
}
