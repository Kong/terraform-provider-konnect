// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StandardWebhooksPluginConfig struct {
	SecretV1        types.String `tfsdk:"secret_v1"`
	ToleranceSecond types.Int64  `tfsdk:"tolerance_second"`
}
