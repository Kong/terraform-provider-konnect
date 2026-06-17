// Package fixture is a compilable miniature of the four generated locations the
// dynamictype pass rewrites (model, schema, refresh, build) plus the SDK field
// they feed. It is loaded with full type information by the pass's golden test;
// the expected post-processed form lives in expected.go.golden.
package fixture

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v3/internal/referenceabletypes"
)

// Redis is the resource model (mirrors internal/provider/types).
type Redis struct {
	Host types.String                           `tfsdk:"host"`
	Port referenceabletypes.ReferenceableString `tfsdk:"port"`
}

// Config nests the redis model.
type Config struct {
	Redis *Redis
}

// Model is the top-level resource model.
type Model struct {
	Config *Config
}

// SDKRedis is the SDK shared model (mirrors internal/sdk/models/shared).
type SDKRedis struct {
	// A port number between 0 and 65535, inclusive, or a referenceable value.
	Port *string `json:"port,omitempty"`
}

// GetPort is the SDK getter for the port field.
func (r *SDKRedis) GetPort() *string {
	if r == nil {
		return nil
	}
	return r.Port
}

// SDKConfig nests the SDK redis model.
type SDKConfig struct {
	Redis *SDKRedis
}

// SDKResp is the SDK response wrapper.
type SDKResp struct {
	Config *SDKConfig
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
	}
}

// Refresh mirrors the generated RefreshFromShared flow (API -> model).
func (r *Model) Refresh(ctx context.Context, resp *SDKResp) {
	var diags diag.Diagnostics
	portValuable, portDiags := referenceabletypes.ReferenceableStringType{}.ValueFromString(ctx, types.StringPointerValue(resp.Config.Redis.Port))
	diags.Append(portDiags...)
	r.Config.Redis.Port = portValuable.(referenceabletypes.ReferenceableString)
	_ = diags
}

// Build mirrors the generated ToShared flow (model -> API).
func (r *Model) Build() *SDKRedis {
	port1 := new(string)
	if !r.Config.Redis.Port.IsUnknown() && !r.Config.Redis.Port.IsNull() {
		*port1 = r.Config.Redis.Port.ValueString()
	} else {
		port1 = nil
	}
	return &SDKRedis{
		Port: port1,
	}
}
