// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConfigurationDataPlaneGroupAutoscaleStatic struct {
	InstanceType       types.String `tfsdk:"instance_type"`
	Kind               types.String `tfsdk:"kind"`
	RequestedInstances types.Int64  `tfsdk:"requested_instances"`
}
