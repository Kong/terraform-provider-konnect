// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OpenTelemetry struct {
	Endpoint        types.String `tfsdk:"endpoint"`
	RefreshInterval types.String `tfsdk:"refresh_interval"`
}