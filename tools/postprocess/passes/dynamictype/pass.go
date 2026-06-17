// Package dynamictype is the post-processing workflow that rewires a field
// flagged as a referenceable (polymorphic) value into a Terraform dynamic
// attribute.
//
// Speakeasy cannot emit types.Dynamic / schema.DynamicAttribute, so a field
// such as redis.port — normally an integer but optionally a referenceable vault
// string — is generated as a string carrying the marker custom type
// referenceabletypes.ReferenceableStringType. This pass detects that marker and
// rewrites the four generated locations (schema, model, refresh, build) plus the
// SDK field it feeds, so the attribute accepts either a number or a string.
package dynamictype

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/kong/terraform-provider-konnect/v3/tools/postprocess/engine"
)

// Marker type identity (resolved via go/types, never matched as a string).
const (
	markerValue = "ReferenceableString"     // model field type
	markerType  = "ReferenceableStringType" // schema CustomType
)

// Import paths of the symbols the pass introduces.
const (
	dynamicPkg        = engine.ModulePath + "/internal/provider/dynamic"
	typesPkg          = "github.com/hashicorp/terraform-plugin-framework/types"
	dynamicDefaultPkg = "github.com/hashicorp/terraform-plugin-framework/resource/schema/dynamicdefault"
	bigPkg            = "math/big"
)

// kind discriminates the payload carried by a Target.
type kind int

const (
	kindModel kind = iota
	kindSchema
	kindRefresh
	kindBuild
	kindSDKField
	kindSDKGetter
)

// modelSite retypes a model struct field to types.Dynamic.
type modelSite struct {
	field *dst.Field
}

// schemaSite rewrites a schema.StringAttribute composite literal to a
// schema.DynamicAttribute.
type schemaSite struct {
	lit *dst.CompositeLit
}

// refreshSite replaces the marker's ValueFromString block with a single
// dynamic.ToFramework assignment.
type refreshSite struct {
	block         *dst.BlockStmt
	lhs           dst.Expr // model field path, e.g. r.Config.Redis.Port
	src           dst.Expr // SDK source value, e.g. resp.Config.Redis.Port
	start         int      // index of the first statement to replace
	count         int      // number of statements to replace
	leadingBefore dst.SpaceType
	leading       dst.Decorations
}

// buildSite replaces the primitive-builder block with a single dynamic.ToAny
// assignment.
type buildSite struct {
	block         *dst.BlockStmt
	varName       string   // the builder variable, e.g. port1
	model         dst.Expr // model field path, e.g. r.Config.Redis.Port
	start         int
	count         int
	leadingBefore dst.SpaceType
	leading       dst.Decorations
}

// sdkFieldSite retypes the SDK struct field that feeds the dynamic value to any.
type sdkFieldSite struct {
	field *dst.Field
}

// sdkGetterSite retypes the SDK getter's return type to any.
type sdkGetterSite struct {
	fn *dst.FuncDecl
}

type payload struct {
	kind kind
	data any
}

// Pass implements engine.Pass.
type Pass struct{}

// New returns the dynamic-type pass.
func New() *Pass { return &Pass{} }

// Name identifies the pass.
func (*Pass) Name() string { return "dynamictype" }

// Apply dispatches to the per-kind edit and marks the file for rewriting.
func (pass *Pass) Apply(b *engine.Bundle, t engine.Target) error {
	b.Touch(t.File)
	return pass.apply(t)
}

// apply performs the edit without touching the bundle, so it can be exercised in
// isolation by tests.
func (*Pass) apply(t engine.Target) error {
	p, ok := t.Data.(payload)
	if !ok {
		return fmt.Errorf("unexpected target payload %T", t.Data)
	}
	switch p.kind {
	case kindModel:
		return applyModel(p.data.(*modelSite))
	case kindSchema:
		return applySchema(p.data.(*schemaSite))
	case kindRefresh:
		return applyRefresh(p.data.(*refreshSite))
	case kindBuild:
		return applyBuild(p.data.(*buildSite))
	case kindSDKField:
		return applySDKField(p.data.(*sdkFieldSite))
	case kindSDKGetter:
		return applySDKGetter(p.data.(*sdkGetterSite))
	default:
		return fmt.Errorf("unknown edit kind %d", p.kind)
	}
}

// target is a small helper to build an engine.Target for a located site.
func target(name string, pkg *decorator.Package, file *dst.File, k kind, data any) engine.Target {
	return engine.Target{Name: name, Pkg: pkg, File: file, Data: payload{kind: k, data: data}}
}
