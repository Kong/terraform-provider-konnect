package referenceabletypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = (*ReferenceableStringType)(nil)

// ReferenceableStringType is an attribute type that represents a string that is base64 encoded, but only in configuration and state, not in the response. It has
// custom semantic equality defined in the Value type, which decodes the string and compares it with the response after create / update.
type ReferenceableStringType struct {
	basetypes.StringType
}

// String returns a human readable string of the type name.
func (t ReferenceableStringType) String() string {
	return "customtypes.encodedstring.Base64InputType"
}

// ValueType returns the Value type.
func (t ReferenceableStringType) ValueType(ctx context.Context) attr.Value {
	return ReferenceableString{}
}

// Equal returns true if the given type is equivalent.
func (t ReferenceableStringType) Equal(o attr.Type) bool {
	other, ok := o.(ReferenceableStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t ReferenceableStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return ReferenceableString{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t ReferenceableStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}
