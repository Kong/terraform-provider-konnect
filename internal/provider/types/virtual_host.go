// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type VirtualHost struct {
	JSONPatches []JSONPatches                                                    `tfsdk:"json_patches"`
	Match       MeshProxyPatchItemSpecDefaultAppendModificationsVirtualHostMatch `tfsdk:"match"`
	Operation   types.String                                                     `tfsdk:"operation"`
	Value       types.String                                                     `tfsdk:"value"`
}