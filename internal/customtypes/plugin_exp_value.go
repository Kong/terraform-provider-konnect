package plugindynamic

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

//
// VALUE
//

type StringOrIntValue struct {
	basetypes.DynamicValue
}

func (v StringOrIntValue) Equal(o attr.Value) bool {
	other, ok := o.(StringOrIntValue)
	return ok && v.DynamicValue.Equal(other.DynamicValue)
}

func (v StringOrIntValue) Type(ctx context.Context) attr.Type {
	return StringOrIntType{}
}

// Helpers

func (v StringOrIntValue) IsString() bool {
	_, ok := v.UnderlyingValue().(basetypes.StringValue)
	return ok
}

func (v StringOrIntValue) IsInt() bool {
	num, ok := v.UnderlyingValue().(basetypes.NumberValue)
	if !ok {
		return false
	}

	f, _ := num.ValueBigFloat().Int(nil)
	asFloat := new(big.Float).SetInt(f)

	return num.ValueBigFloat().Cmp(asFloat) == 0
}

func (v StringOrIntValue) AsString() (string, error) {
	s, ok := v.UnderlyingValue().(basetypes.StringValue)
	if !ok {
		return "", fmt.Errorf("not a string")
	}
	return s.ValueString(), nil
}

func (v StringOrIntValue) AsInt64() (int64, error) {
	num, ok := v.UnderlyingValue().(basetypes.NumberValue)
	if !ok {
		return 0, fmt.Errorf("not a number")
	}

	i, acc := num.ValueBigFloat().Int64()
	if acc != big.Exact {
		return 0, fmt.Errorf("not an exact int64")
	}

	return i, nil
}

//
// Constructor
//

func NewStringOrIntValue(v attr.Value) StringOrIntValue {
	return StringOrIntValue{
		DynamicValue: basetypes.NewDynamicValue(v),
	}
}
