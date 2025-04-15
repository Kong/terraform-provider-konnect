// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Active struct {
	Concurrency            types.Int64             `tfsdk:"concurrency"`
	Headers                map[string]types.String `tfsdk:"headers"`
	Healthy                *Healthy                `tfsdk:"healthy"`
	HTTPPath               types.String            `tfsdk:"http_path"`
	HTTPSSni               types.String            `tfsdk:"https_sni"`
	HTTPSVerifyCertificate types.Bool              `tfsdk:"https_verify_certificate"`
	Timeout                types.Float64           `tfsdk:"timeout"`
	Type                   types.String            `tfsdk:"type"`
	Unhealthy              *Unhealthy              `tfsdk:"unhealthy"`
}
