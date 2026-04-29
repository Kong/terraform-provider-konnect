package plugindynamic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

//
// TYPE
//

type StringOrIntType struct {
	basetypes.DynamicType
}

func (t StringOrIntType) Equal(o attr.Type) bool {
	other, ok := o.(StringOrIntType)
	return ok && t.DynamicType.Equal(other.DynamicType)
}

func (t StringOrIntType) String() string {
	return "StringOrIntType"
}

func (t StringOrIntType) ValueType(ctx context.Context) attr.Value {
	return StringOrIntValue{}
}
