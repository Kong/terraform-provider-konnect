// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateGraphqlProxyCacheAdvancedPluginConfig struct {
	BypassOnErr types.Bool                                   `tfsdk:"bypass_on_err"`
	CacheTTL    types.Int64                                  `tfsdk:"cache_ttl"`
	Memory      *CreateGraphqlProxyCacheAdvancedPluginMemory `tfsdk:"memory"`
	Redis       *CreateGraphqlProxyCacheAdvancedPluginRedis  `tfsdk:"redis"`
	Strategy    types.String                                 `tfsdk:"strategy"`
	VaryHeaders []types.String                               `tfsdk:"vary_headers"`
}
