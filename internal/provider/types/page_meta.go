// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type PageMeta struct {
	Number types.Number `tfsdk:"number"`
	Size   types.Number `tfsdk:"size"`
	Total  types.Number `tfsdk:"total"`
}