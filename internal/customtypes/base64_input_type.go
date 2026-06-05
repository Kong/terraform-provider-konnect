package customtypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = (*ReferenceableIntegerType)(nil)

// ReferenceableIntegerType is an attribute type representing a string value with
// custom semantic equality defined in the value type.
type ReferenceableIntegerType struct {
	basetypes.StringType
}

// String returns a human readable string of the type name.
func (t ReferenceableIntegerType) String() string {
	return "customtypes.referenceableinteger.ReferenceableIntegerType"
}

// ValueType returns the Value type.
func (t ReferenceableIntegerType) ValueType(ctx context.Context) attr.Value {
	return ReferenceableInteger{}
}

// Equal returns true if the given type is equivalent.
func (t ReferenceableIntegerType) Equal(o attr.Type) bool {
	other, ok := o.(ReferenceableIntegerType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t ReferenceableIntegerType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return ReferenceableInteger{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t ReferenceableIntegerType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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
