// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshAccessLogItemSpecTCP struct {
	Address types.String `tfsdk:"address"`
	Format  *Format      `tfsdk:"format"`
}