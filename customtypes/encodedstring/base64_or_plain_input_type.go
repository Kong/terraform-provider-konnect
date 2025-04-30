// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encodedstring

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable = (*Base64OrPlainInputType)(nil)
)

// Base64OrPlainInputType is an attribute type that represents a string that is optionally base64 encoded, but only in configuration and state, not in the response. It has
// custom semantic equality defined in the Value type, which does a double comparison with the value in response of Read/Create/Update - with and without decoding.
type Base64OrPlainInputType struct {
	basetypes.StringType
}

// String returns a human readable string of the type name.
func (t Base64OrPlainInputType) String() string {
	return "customtypes.encodedstring.Base64OrPlainInputType"
}

// ValueType returns the Value type.
func (t Base64OrPlainInputType) ValueType(ctx context.Context) attr.Value {
	return Base64OrPlainInput{}
}

// Equal returns true if the given type is equivalent.
func (t Base64OrPlainInputType) Equal(o attr.Type) bool {
	other, ok := o.(Base64OrPlainInputType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t Base64OrPlainInputType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return Base64OrPlainInput{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t Base64OrPlainInputType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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
