// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Propagation struct {
	Clear         []types.String `tfsdk:"clear"`
	DefaultFormat types.String   `tfsdk:"default_format"`
	Extract       []types.String `tfsdk:"extract"`
	Inject        []types.String `tfsdk:"inject"`
}