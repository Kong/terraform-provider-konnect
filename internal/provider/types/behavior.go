// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Behavior struct {
	IdpErrorResponseBodyTemplate    types.String  `tfsdk:"idp_error_response_body_template"`
	IdpErrorResponseContentType     types.String  `tfsdk:"idp_error_response_content_type"`
	IdpErrorResponseMessage         types.String  `tfsdk:"idp_error_response_message"`
	IdpErrorResponseStatusCode      types.Int64   `tfsdk:"idp_error_response_status_code"`
	PurgeTokenOnUpstreamStatusCodes []types.Int64 `tfsdk:"purge_token_on_upstream_status_codes"`
	UpstreamAccessTokenHeaderName   types.String  `tfsdk:"upstream_access_token_header_name"`
}
