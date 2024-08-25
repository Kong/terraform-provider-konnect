// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AppAuthStrategyOpenIDConnectRequest struct {
	Active        types.Bool                                 `tfsdk:"active"`
	Configs       AppAuthStrategyOpenIDConnectRequestConfigs `tfsdk:"configs"`
	CreatedAt     types.String                               `tfsdk:"created_at"`
	DcrProvider   *DcrProvider                               `tfsdk:"dcr_provider"`
	DcrProviderID types.String                               `tfsdk:"dcr_provider_id"`
	DisplayName   types.String                               `tfsdk:"display_name"`
	ID            types.String                               `tfsdk:"id"`
	Labels        map[string]types.String                    `tfsdk:"labels"`
	Name          types.String                               `tfsdk:"name"`
	StrategyType  types.String                               `tfsdk:"strategy_type"`
	UpdatedAt     types.String                               `tfsdk:"updated_at"`
}
