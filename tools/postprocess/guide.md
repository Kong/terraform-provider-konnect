# Post-Processing Layer — Code Guide

A short orientation to the code added for the dynamic-type post-processor. For
the *why* and the full design, read `/postprocess.md`. This doc covers the
*what* and *where*.

## The problem in one paragraph

Speakeasy regenerates the provider on every `make speakeasy` and cannot emit
`types.Dynamic` / `schema.DynamicAttribute`. Some Kong fields (canonically
`redis.port`) are polymorphic — a number *or* a referenceable vault string — and
must be modeled as a Terraform **dynamic** attribute. We flag such a field in the
OpenAPI so Speakeasy stamps it with a **marker custom type**
(`referenceabletypes.ReferenceableStringType` / `ReferenceableString`); this tool
runs right after generation, detects the marker, and rewrites the generated code
to use a real dynamic attribute. It is deterministic and idempotent, so CI can
assert a clean `git diff`.

## Map of what was added

```
internal/provider/dynamic/        # runtime helpers (hand-written, never regenerated)
  convert.go                       # ToAny / ToFramework
  convert_test.go

tools/postprocess/                 # the build-time tool (run via `go run ./tools/postprocess`)
  main.go                          # CLI: build registry -> engine.Run
  engine/                          # workflow-agnostic machinery
    engine.go                      # Bundle, Pass, Registry, Run, RenderFile, build gate
    loader.go                      # go/packages -> decorated dst trees + types info
    target.go                      # Target (one unit of work)
  passes/dynamictype/              # the first (and only) workflow
    pass.go                        # Pass impl + per-edit payload types + dispatch
    detect.go                      # find marker + correlate the 5 edit sites
    apply.go                       # the 5 edits (pure dst -> dst)
    helpers.go                     # type-resolution + dst-construction utilities
    pass_test.go                   # golden + detection-count tests
    unit_test.go                   # translateDefault + structural-helper tests
    testdata/fixture/              # compilable miniature input + expected.go.golden

Makefile                           # `speakeasy` runs the tool; `postprocess`, `postprocess-check`
```

## The two layers, and why they're separate

**Runtime layer — `internal/provider/dynamic`.** Plain, unit-tested Go that the
generated code *calls at runtime*. `ToAny` turns a `types.Dynamic` from Terraform
into the value sent to the API; `ToFramework` does the reverse. The tool only
ever rewires call sites to these two symbols, so the hard conversion logic lives
here and is tested independently of any AST machinery. This is never regenerated.

**Build-time layer — `tools/postprocess`.** Runs once after generation. It never
ships to users; it only *rewrites generated source*.

## Core abstractions (engine)

The engine is deliberately decoupled from any workflow. Three things to know:

- **`Bundle`** — the loaded, mutable world: decorated `dst` trees for the
  packages we may rewrite, plus their `go/types` info, plus a `touched` set so
  only modified files are re-rendered. Built by `engine.Load`.

- **`Pass`** — the extensibility seam:
  ```go
  type Pass interface {
      Name() string
      Detect(b *Bundle) ([]Target, error)  // find the units of work
      Apply(b *Bundle, t Target) error      // mutate dst for one unit (pure AST->AST)
  }
  ```
  Adding a future workflow = implement this and `Registry.Register` it. The
  loader, runner, formatter and build gate are untouched.

- **`Target`** — one unit of work. `Name`/`Pkg`/`File` are engine-visible; `Data`
  is an opaque pass-specific payload the pass type-asserts in `Apply`.

**`engine.Run`** is the whole algorithm: `Load` → for each pass `Detect` then
`Apply` each target → `RenderFile` every touched file (restore dst → goimports) →
write → `go build ./...` as a hard gate. A non-zero build leaves the tree for
triage.

### Two libraries doing the heavy lifting

- **`go/packages`** loads the provider with full **type information**. We resolve
  a field's *type* (not its spelling) to decide if it's the marker — robust to
  aliases, renamed imports, and formatting.
- **`dave/dst`** (Decorated Syntax Tree) does the rewriting. Unlike
  `go/ast`+`printer`, it preserves comments and formatting on round-trip, so
  unrelated code stays byte-for-byte identical. The `decorator` ties the two
  together: `engine.DstFor(pkg, astNode)` maps a typed `go/ast` node to the
  `dst` node we mutate.

> Detection reads the **AST + types**; edits mutate the **dst**. `DstFor` is the
> bridge. (This is also why in-memory idempotency isn't unit-tested: detection
> reads the unchanged AST. Idempotency is a *file-level* property — re-load the
> rewritten file and the marker is gone — verified by re-running the tool.)

## The DynamicType pass

**Marker** (resolved via `go/types`, never string-matched):
`referenceabletypes.ReferenceableStringType` on a schema attribute's `CustomType`,
and `referenceabletypes.ReferenceableString` as a model field's type.

**`detect.go`** finds five kinds of edit site and emits one `Target` each. The
marker anchors three of them directly; the other two are *correlated* by resolved
identity (no file names, no line numbers):

| Site | How it's found |
| --- | --- |
| model field (b) | struct field whose type resolves to the marker value |
| schema attr (a) | `KeyValueExpr` whose value is a composite lit with a marker `CustomType` |
| refresh (c) | the `…ReferenceableStringType{}.ValueFromString(…)` 3-statement block |
| build (d) | `X := new(string)` + `if …ValueString()` block whose model field matches a marked `(struct, field)` |
| SDK field/getter (e) | derived from the refresh source expr `resp.…Port` — its named type + field name locate the declaration in the shared package |

The `(struct, field)` and SDK correlation use `fieldRef{pkgPath, type, field}`
keys (see `helpers.go`), matched by resolved name rather than pointer identity so
it works whether deps are loaded from source or export data.

**`apply.go`** holds the five pure dst edits:

| Edit | Transformation |
| --- | --- |
| `applySchema` | `schema.StringAttribute` → `schema.DynamicAttribute`; drop `CustomType` + `Validators`; `translateDefault` rewrites a kind-typed default to a `dynamicdefault.StaticValue(...)` |
| `applyModel` | field type → `types.Dynamic` |
| `applyRefresh` | replace the block with `<model> = dynamic.ToFramework(<src>)` |
| `applyBuild` | replace the block with `<var> := dynamic.ToAny(<model>)` |
| `applySDKField` / `applySDKGetter` | SDK field + getter result → `any` (so the `any` from `ToAny` is assignable) |

`translateDefault` maps `StaticInt64`/`StaticFloat64` → number, `StaticString` →
string, `StaticBool` → bool, and **errors** on anything unmapped so the runner
fails loudly rather than emitting wrong code.

### How new symbols and imports are handled

Edits build `dst.Ident`s carrying a package **`Path`** (e.g.
`ident("ToFramework", dynamicPkg)`). On render, `RenderFile` restores with import
management (so a path-bearing ident expands to a qualified `dynamic.ToFramework`
and the import is added), then runs **goimports** as the final authority —
adding the helper/framework packages and dropping the now-unused
`referenceabletypes` / kind-typed default+validator imports.

## Tests

- `pass_test.go` — loads `testdata/fixture/` (a compilable miniature of all five
  locations) with real type info, runs detect+apply, and golden-compares the
  rendered output to `expected.go.golden`. Run `go test ./... -update` to refresh
  the golden. Also asserts the detection finds exactly one site of each kind.
- `unit_test.go` — `translateDefault` across all default kinds (incl. the error
  path) and the structural helpers (`unwrapConv`, `buildModelExpr`, tag parsing…).
- `internal/provider/dynamic/convert_test.go` — round-trips number/string/bool.

## Running it

```bash
go run ./tools/postprocess          # detect + transform + `go build ./...` gate
make postprocess                    # same, via Make
make postprocess-check              # runs it, then asserts a clean git diff (CI)
```

It runs automatically inside `make speakeasy`, after `speakeasy run` and before
`go generate`, so generated docs reflect the dynamic attribute.

## Adding the next dynamic field

Mark it in the OpenAPI overlay so Speakeasy stamps the
`ReferenceableStringType` marker on it. That's it — the pass picks it up by the
marker; no code change here.

## Adding a different workflow

Create `passes/<workflow>/`, implement the `Pass` interface, register it in
`main.go`, add a `testdata/<workflow>/` golden case. The engine is untouched.
