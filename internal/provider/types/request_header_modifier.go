// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type RequestHeaderModifier struct {
	Add    []MeshGlobalRateLimitItemSpecAdd `tfsdk:"add"`
	Remove []types.String                   `tfsdk:"remove"`
	Set    []MeshGlobalRateLimitItemSpecAdd `tfsdk:"set"`
}
