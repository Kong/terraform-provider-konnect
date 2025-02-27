// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshAccessLogItemBackends struct {
	File          *File                                   `tfsdk:"file"`
	OpenTelemetry *MeshAccessLogItemSpecFromOpenTelemetry `tfsdk:"open_telemetry"`
	TCP           *MeshAccessLogItemSpecFromTCP           `tfsdk:"tcp"`
	Type          types.String                            `tfsdk:"type"`
}
