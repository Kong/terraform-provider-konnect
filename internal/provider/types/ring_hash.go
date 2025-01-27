// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type RingHash struct {
	HashFunction types.String   `tfsdk:"hash_function"`
	HashPolicies []HashPolicies `tfsdk:"hash_policies"`
	MaxRingSize  types.Int64    `tfsdk:"max_ring_size"`
	MinRingSize  types.Int64    `tfsdk:"min_ring_size"`
}
