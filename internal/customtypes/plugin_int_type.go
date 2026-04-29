package plugindynamic

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

//
// TYPE
//

type MyIntType struct {
	basetypes.Int64Type
}

func (t MyIntType) Equal(o attr.Type) bool {
	other, ok := o.(MyIntType)
	return ok && t.Int64Type.Equal(other.Int64Type)
}

func (t MyIntType) String() string {
	return "MyIntType"
}

func (t MyIntType) ValueType(ctx context.Context) attr.Value {
	return MyIntValue{}
}

// ValueFromInt64 returns an Int64Valuable type given an Int64Value.
func (t MyIntType) ValueFromInt64(ctx context.Context, in basetypes.Int64Value) (basetypes.Int64Valuable, diag.Diagnostics) {
	return MyIntValue{
		Int64Value: in,
	}, nil
}

func (t MyIntType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.Int64Valuable, diag.Diagnostics) {
	valuestr := in.ValueString()
	// convert string to int64
	var intVal int64
	_, err := fmt.Sscanf(valuestr, "%d", &intVal)
	if err != nil {
		return MyIntValue{
			Int64Value: basetypes.NewInt64Null(),
		}, diag.Diagnostics{diag.NewErrorDiagnostic("Invalid String for MyIntType", fmt.Sprintf("The provided string value '%s' could not be parsed as an integer: %s", valuestr, err.Error()))}
	}

	return MyIntValue{
		Int64Value: basetypes.NewInt64PointerValue(&intVal),
	}, nil
}
