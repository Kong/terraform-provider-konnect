// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshPassthroughItemDefault struct {
	AppendMatch     []AppendMatch `tfsdk:"append_match"`
	PassthroughMode types.String  `tfsdk:"passthrough_mode"`
}