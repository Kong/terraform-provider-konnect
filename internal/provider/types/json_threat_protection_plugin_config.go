// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type JSONThreatProtectionPluginConfig struct {
	AllowDuplicateObjectEntryName types.Bool   `tfsdk:"allow_duplicate_object_entry_name"`
	EnforcementMode               types.String `tfsdk:"enforcement_mode"`
	ErrorMessage                  types.String `tfsdk:"error_message"`
	ErrorStatusCode               types.Int64  `tfsdk:"error_status_code"`
	MaxArrayElementCount          types.Int64  `tfsdk:"max_array_element_count"`
	MaxBodySize                   types.Int64  `tfsdk:"max_body_size"`
	MaxContainerDepth             types.Int64  `tfsdk:"max_container_depth"`
	MaxObjectEntryCount           types.Int64  `tfsdk:"max_object_entry_count"`
	MaxObjectEntryNameLength      types.Int64  `tfsdk:"max_object_entry_name_length"`
	MaxStringValueLength          types.Int64  `tfsdk:"max_string_value_length"`
}
