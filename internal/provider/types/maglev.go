// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Maglev struct {
	HashPolicies []HashPolicies `tfsdk:"hash_policies"`
	TableSize    types.Int64    `tfsdk:"table_size"`
}
