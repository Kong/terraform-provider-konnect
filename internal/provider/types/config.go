// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Config struct {
	ControlPlaneEndpoint types.String `tfsdk:"control_plane_endpoint"`
	TelemetryEndpoint    types.String `tfsdk:"telemetry_endpoint"`
}
