// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BasicAuthPluginConfig struct {
	Anonymous       types.String `tfsdk:"anonymous"`
	HideCredentials types.Bool   `tfsdk:"hide_credentials"`
	Realm           types.String `tfsdk:"realm"`
}
