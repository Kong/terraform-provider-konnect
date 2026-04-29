package plugindynamic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

//
// VALUE
//

type MyIntValue struct {
	basetypes.Int64Value
}

func (v MyIntValue) Equal(o attr.Value) bool {
	other, ok := o.(MyIntValue)
	return ok && v.Int64Value.Equal(other.Int64Value)
}

func (v MyIntValue) Type(ctx context.Context) attr.Type {
	return MyIntType{}
}

// Helper
func (v MyIntValue) ValueInt64() int64 {
	return v.Int64Value.ValueInt64()
}

//
// Constructors
//

func NewMyIntNull() MyIntValue {
	return MyIntValue{
		Int64Value: basetypes.NewInt64Null(),
	}
}

func NewMyIntUnknown() MyIntValue {
	return MyIntValue{
		Int64Value: basetypes.NewInt64Unknown(),
	}
}

func NewMyIntValue(v int64) MyIntValue {
	return MyIntValue{
		Int64Value: basetypes.NewInt64Value(v),
	}
}
