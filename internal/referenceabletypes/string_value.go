package referenceabletypes

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ basetypes.StringValuableWithSemanticEquals = (*ReferenceableString)(nil)

// ReferenceableString represents a valid base64 encoded string. Custom semantic equality
// logic is defined for ReferenceableString, where the encoded value is decoded, and then compared to the value received in response from Read / Create / Update.
type ReferenceableString struct {
	basetypes.StringValue
}

// Type returns an Base64InputType.
func (v ReferenceableString) Type(_ context.Context) attr.Type {
	return ReferenceableStringType{}
}

// Equal returns true if the given value is equivalent.
func (v ReferenceableString) Equal(o attr.Value) bool {
	other, ok := o.(ReferenceableString)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// NewBase64InputNull creates an Base64Input with a null value. Determine whether the value is null via IsNull method.
func NewBase64InputNull() ReferenceableString {
	return ReferenceableString{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewBase64InputUnknown creates an Base64Input with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewBase64InputUnknown() ReferenceableString {
	return ReferenceableString{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewBase64InputValue creates an Base64Input with a known value. Access the value via ValueString method.
func NewBase64InputValue(value string) ReferenceableString {
	return ReferenceableString{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewBase64InputPointerValue creates an Base64Input with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewBase64InputPointerValue(value *string) ReferenceableString {
	return ReferenceableString{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

// StringSemanticEquals decodes givenValuable, compares it with the receiver string value and returns whether they are inconsequentially equal.
// Semantic equality is checked during planning phase, and after receiving response in apply phase. In planning phase, givenValuable comes from the state, and v
// is the current value - from Read method response. In the apply phase, givenValuable is from the plan, and v is from response of Create / Update.
func (v ReferenceableString) StringSemanticEquals(ctx context.Context, givenValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	// The framework should always pass the correct value type, but always check
	_, ok := givenValuable.(ReferenceableString)

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

	decodedGivenStringBytes, errors := base64.StdEncoding.DecodeString(givenStringValue.ValueString())
	if errors != nil {
		// Not base64 encoded, return false so value from response is used in plan to compare with config.
		diags.AddError(
			"Custom String Decode Error",
			"An unexpected error was encountered trying to decode a custom string. This is always an error in the provider. Please report the following to the provider developer:\n\n",
		)
		return false, diags
	}

	return string(decodedGivenStringBytes) == v.ValueString(), diags
}
