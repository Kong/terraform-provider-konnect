// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Filters struct {
	RequestHeaderModifier  *RequestHeaderModifier `tfsdk:"request_header_modifier"`
	RequestMirror          *RequestMirror         `tfsdk:"request_mirror"`
	RequestRedirect        *RequestRedirect       `tfsdk:"request_redirect"`
	ResponseHeaderModifier *RequestHeaderModifier `tfsdk:"response_header_modifier"`
	Type                   types.String           `tfsdk:"type"`
	URLRewrite             *URLRewrite            `tfsdk:"url_rewrite"`
}