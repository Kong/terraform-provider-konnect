// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshRateLimitItemHTTP struct {
	Disabled    types.Bool                                  `tfsdk:"disabled"`
	OnRateLimit *OnRateLimit                                `tfsdk:"on_rate_limit"`
	RequestRate *MeshGlobalRateLimitItemSpecFromRequestRate `tfsdk:"request_rate"`
}