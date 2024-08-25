// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type APIProductVersionPortal struct {
	ApplicationRegistrationEnabled types.Bool                      `tfsdk:"application_registration_enabled"`
	AuthStrategies                 []APIProductVersionAuthStrategy `tfsdk:"auth_strategies"`
	AutoApproveRegistration        types.Bool                      `tfsdk:"auto_approve_registration"`
	Deprecated                     types.Bool                      `tfsdk:"deprecated"`
	PortalID                       types.String                    `tfsdk:"portal_id"`
	PortalName                     types.String                    `tfsdk:"portal_name"`
	PortalProductVersionID         types.String                    `tfsdk:"portal_product_version_id"`
	PublishStatus                  types.String                    `tfsdk:"publish_status"`
}