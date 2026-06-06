package customtypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ basetypes.StringValuableWithSemanticEquals = (*ReferenceableInteger)(nil)

// ReferenceableInteger represents a string value with custom semantic equality
// logic.
type ReferenceableInteger struct {
	basetypes.StringValue
}

// Type returns a ReferenceableIntegerType.
func (v ReferenceableInteger) Type(_ context.Context) attr.Type {
	return ReferenceableIntegerType{}
}

// Equal returns true if the given value is equivalent.
func (v ReferenceableInteger) Equal(o attr.Value) bool {
	other, ok := o.(ReferenceableInteger)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// NewReferenceableIntegerNull creates a ReferenceableInteger with a null value.
func NewReferenceableIntegerNull() ReferenceableInteger {
	return ReferenceableInteger{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewReferenceableIntegerUnknown creates a ReferenceableInteger with an unknown value.
func NewReferenceableIntegerUnknown() ReferenceableInteger {
	return ReferenceableInteger{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewReferenceableIntegerValue creates a ReferenceableInteger with a known value.
func NewReferenceableIntegerValue(value string) ReferenceableInteger {
	return ReferenceableInteger{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewReferenceableIntegerPointerValue creates a ReferenceableInteger with a null value if nil or a known value.
func NewReferenceableIntegerPointerValue(value *string) ReferenceableInteger {
	return ReferenceableInteger{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

// StringSemanticEquals compares givenValuable with the receiver by raw string
// value and returns whether they are inconsequentially equal.
func (v ReferenceableInteger) StringSemanticEquals(ctx context.Context, givenValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	// The framework should always pass the correct value type, but always check
	_, ok := givenValuable.(ReferenceableInteger)

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

	return givenStringValue.ValueString() == v.ValueString(), diags
}
