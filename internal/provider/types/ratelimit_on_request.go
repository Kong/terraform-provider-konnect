// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type RatelimitOnRequest struct {
	Kind   types.String `tfsdk:"kind"`
	Limits []Limits     `tfsdk:"limits"`
	Name   types.String `tfsdk:"name"`
}
