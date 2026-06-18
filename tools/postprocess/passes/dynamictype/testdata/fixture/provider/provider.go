// Package provider is the provider-side half of the dynamictype golden fixture
// (mirrors internal/provider + internal/provider/types): the model structs, the
// schema, the refresh flow and the build flow. The expected post-processed form
// lives in expected.go.golden; the SDK half it feeds lives in ../shared.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v3/internal/referenceabletypes"
	"github.com/kong/terraform-provider-konnect/v3/tools/postprocess/passes/dynamictype/testdata/fixture/shared"
)

// Redis pairs with shared.Redis through a refresh site (different struct names).
type Redis struct {
	Host types.String                           `tfsdk:"host"`
	Port referenceabletypes.ReferenceableString `tfsdk:"port"`
}

// ACLRule pairs with shared.ACLRule by name through a build site (no refresh).
type ACLRule struct {
	ResourceNames referenceabletypes.ReferenceableString `tfsdk:"resource_names"`
}

// Model is the top-level resource model.
type Model struct {
	Redis *Redis
	Rules []ACLRule
}

// Schema mirrors the generated resource schema attribute map.
func Schema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"port": schema.StringAttribute{
			CustomType:  referenceabletypes.ReferenceableStringType{},
			Computed:    true,
			Optional:    true,
			Default:     stringdefault.StaticString("6379"),
			Description: `An integer representing a port number between 0 and 65535, inclusive.`,
		},
		"resource_names": schema.StringAttribute{
			CustomType:  referenceabletypes.ReferenceableStringType{},
			Required:    true,
			Description: `If any of these entries match, the resource name matches for this rule.`,
		},
	}
}

// Refresh mirrors RefreshFromShared: the pointer/referenceable shape.
func (m *Model) Refresh(ctx context.Context, redis *shared.Redis) {
	var diags diag.Diagnostics
	portValuable, portDiags := referenceabletypes.ReferenceableStringType{}.ValueFromString(ctx, types.StringPointerValue(redis.Port))
	diags.Append(portDiags...)
	m.Redis.Port = portValuable.(referenceabletypes.ReferenceableString)
	_ = diags
}

// Build mirrors ToShared: the pointer builder shape (Redis.Port) and the value
// builder shape nested in a slice loop (Rules[i].ResourceNames).
func (m *Model) Build() (*shared.Redis, []shared.ACLRule) {
	port1 := new(string)
	if !m.Redis.Port.IsUnknown() && !m.Redis.Port.IsNull() {
		*port1 = m.Redis.Port.ValueString()
	} else {
		port1 = nil
	}
	redis := &shared.Redis{
		Port: port1,
	}

	rules := make([]shared.ACLRule, 0, len(m.Rules))
	for rulesIndex := range m.Rules {
		var resourceNames string
		resourceNames = m.Rules[rulesIndex].ResourceNames.ValueString()
		rules = append(rules, shared.ACLRule{
			ResourceNames: resourceNames,
		})
	}
	return redis, rules
}
