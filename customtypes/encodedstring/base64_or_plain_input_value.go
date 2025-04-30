// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encodedstring

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable = (*Base64OrPlainInput)(nil)
)

// Base64OrPlainInput represents a valid IPv4 address string (dotted decimal, no leading zeroes). No semantic equality
// logic is defined for Base64OrPlainInput, so it will follow Terraform's data-consistency rules for strings, which must match byte-for-byte.
type Base64OrPlainInput struct {
	basetypes.StringValue
}

// Type returns an Base64OrPlainInputType.
func (v Base64OrPlainInput) Type(_ context.Context) attr.Type {
	return Base64OrPlainInputType{}
}

// Equal returns true if the given value is equivalent.
func (v Base64OrPlainInput) Equal(o attr.Value) bool {
	other, ok := o.(Base64OrPlainInput)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// NewBase64OrPlainInputNull creates an Base64OrPlainInput with a null value. Determine whether the value is null via IsNull method.
func NewBase64OrPlainInputNull() Base64OrPlainInput {
	return Base64OrPlainInput{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewBase64OrPlainInputUnknown creates an Base64OrPlainInput with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewBase64OrPlainInputUnknown() Base64OrPlainInput {
	return Base64OrPlainInput{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewBase64OrPlainInputValue creates an Base64OrPlainInput with a known value. Access the value via ValueString method.
func NewBase64OrPlainInputValue(value string) Base64OrPlainInput {
	return Base64OrPlainInput{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewBase64OrPlainInputPointerValue creates an Base64OrPlainInput with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewBase64OrPlainInputPointerValue(value *string) Base64OrPlainInput {
	return Base64OrPlainInput{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

// StringSemanticEquals decodes givenValuable, compares it with the receiver string value and returns whether they are inconsequentially equal.
// Semantic equality is checked during planning phase, and after receiving response in apply phase.
// In planning phase, givenValuable comes from the state, and v
// is the current value - from Read method response. In the apply phase, givenValuable is from the plan, and v is from response of Create / Update.
func (v Base64OrPlainInput) StringSemanticEquals(ctx context.Context, givenValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	// The framework should always pass the correct value type, but always check
	_, ok := givenValuable.(Base64OrPlainInput)

	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", givenValuable),
		)

		return false, diags
	}

	givenStringValue, err := givenValuable.ToStringValue(ctx)
	if err != nil {
		// Not a StringValue type.
		diags.AddError(
			"Custom String Conversion Error",
			"An unexpected error was encountered trying to convert a custom string. This is always an error in the provider. Please report the following to the provider developer:\n\n",
		)
		return false, diags
	}

	if v.ValueString() == givenStringValue.ValueString() {
		// given value is a plain string, and equal to v, so do nothing else.
		return true, diags
	}

	decodedGivenValueBytes, errors := base64.StdEncoding.DecodeString(givenStringValue.ValueString())
	if errors != nil {
		// Error while decoding, return false so value from response is used in plan to compare with config.
		diags.AddError(
			"Custom String Decode Error",
			"An unexpected error was encountered trying to decode a custom string. This is always an error in the provider. Please report the following to the provider developer:\n\n",
		)
		return false, diags
	}

	return string(decodedGivenValueBytes) == v.ValueString(), diags
}
