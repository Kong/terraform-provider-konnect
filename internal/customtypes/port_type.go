package plugindynamic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// PortDynamicType is a string-based attribute type for a port number (0–65535).
// The value is stored as a decimal integer string (e.g. "6379").
type PortDynamicType struct {
	basetypes.StringType
}

var _ basetypes.StringTypable = (*PortDynamicType)(nil)

// String returns a human-readable string of the type name.
func (t PortDynamicType) String() string {
	return "plugindynamic.PortDynamicType"
}

// ValueType returns the Value type.
func (t PortDynamicType) ValueType(_ context.Context) attr.Value {
	return PortDynamicValue{}
}

// Equal returns true if the given type is equivalent.
func (t PortDynamicType) Equal(o attr.Type) bool {
	other, ok := o.(PortDynamicType)
	if !ok {
		return false
	}
	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a PortDynamicValue from a StringValue.
func (t PortDynamicType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return PortDynamicValue{StringValue: in}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.
func (t PortDynamicType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

// PortDynamicValue represents a port number (0–65535) stored as a decimal string.
type PortDynamicValue struct {
	basetypes.StringValue
}

var (
	_ basetypes.StringValuable    = (*PortDynamicValue)(nil)
	_ xattr.ValidateableAttribute = (*PortDynamicValue)(nil)
)

// Type returns a PortDynamicType.
func (v PortDynamicValue) Type(_ context.Context) attr.Type {
	return PortDynamicType{}
}

// Equal returns true if the given value represents the same port number.
func (v PortDynamicValue) Equal(o attr.Value) bool {
	other, ok := o.(PortDynamicValue)
	if !ok {
		return false
	}
	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true when both values represent the same port.
func (v PortDynamicValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newValue, ok := newValuable.(PortDynamicValue)
	if !ok {
		return false, diags
	}
	return v.ValueString() == newValue.ValueString(), diags
}

// ValidateAttribute validates that the stored value is an integer in 0–65535.
func (v PortDynamicValue) ValidateAttribute(_ context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}
	port, err := strconv.ParseInt(v.ValueString(), 10, 64)
	if err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid port value",
			fmt.Sprintf("Port must be a decimal integer, got: %q", v.ValueString()))
		return
	}
	if port < 0 || port > 65535 {
		resp.Diagnostics.AddAttributeError(req.Path, "Port out of range",
			fmt.Sprintf("Port must be between 0 and 65535 inclusive, got: %d", port))
	}
}

// NewPortDynamicValueNull creates a PortDynamicValue with a null value.
func NewPortDynamicValueNull() PortDynamicValue {
	return PortDynamicValue{StringValue: basetypes.NewStringNull()}
}

// NewPortDynamicValueUnknown creates a PortDynamicValue with an unknown value.
func NewPortDynamicValueUnknown() PortDynamicValue {
	return PortDynamicValue{StringValue: basetypes.NewStringUnknown()}
}

// NewPortDynamicValue creates a PortDynamicValue from a decimal integer string.
func NewPortDynamicValue(value string) (PortDynamicValue, diag.Diagnostics) {
	return PortDynamicValue{StringValue: basetypes.NewStringValue(value)}, nil
}

// NewPortDynamicValueMust creates a PortDynamicValue from a decimal integer string.
func NewPortDynamicValueMust(value string) PortDynamicValue {
	return PortDynamicValue{StringValue: basetypes.NewStringValue(value)}
}

// NewPortDynamicValueFromInt creates a PortDynamicValue from an integer.
func NewPortDynamicValueFromInt(port int64) PortDynamicValue {
	return PortDynamicValue{StringValue: basetypes.NewStringValue(strconv.FormatInt(port, 10))}
}

// NewPortDynamicPointerValue creates a null PortDynamicValue when value is nil,
// otherwise creates a known value from the integer pointer.
func NewPortDynamicPointerValue(value *int64) (PortDynamicValue, diag.Diagnostics) {
	if value == nil {
		return NewPortDynamicValueNull(), nil
	}
	return NewPortDynamicValue(strconv.FormatInt(*value, 10))
}
