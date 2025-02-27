// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshTraceItemBackends struct {
	Datadog       *Datadog                    `tfsdk:"datadog"`
	OpenTelemetry *MeshTraceItemOpenTelemetry `tfsdk:"open_telemetry"`
	Type          types.String                `tfsdk:"type"`
	Zipkin        *Zipkin                     `tfsdk:"zipkin"`
}
