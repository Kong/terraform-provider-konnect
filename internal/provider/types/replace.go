// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Replace struct {
	Body        []types.String `tfsdk:"body"`
	Headers     []types.String `tfsdk:"headers"`
	Querystring []types.String `tfsdk:"querystring"`
	URI         types.String   `tfsdk:"uri"`
}
