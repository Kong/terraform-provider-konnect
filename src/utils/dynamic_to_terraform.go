// This file is based on the following file and is distributed under the same MPL 2.0 license.
// https://github.com/Tobotimus/terraform-provider-toml/blob/67f006a3c4a75d1c66bbcd8b7f973827d8c865fd/internal/provider/toml_decode.go
// This file has been edited for the purpose of this project.

package utils

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ConvertToTerraformType(dynamicValue any) (attr.Type, attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	switch value := dynamicValue.(type) {
	case string:
		return types.StringType, types.StringValue(value), diags
	case int:
		return types.Int64Type, types.Int64Value(int64(value)), diags
	case int64:
		return types.Int64Type, types.Int64Value(value), diags
	case float32:
		return types.Float64Type, types.Float64Value(float64(value)), diags
	case float64:
		return types.Float64Type, types.Float64Value(value), diags
	case bool:
		return types.BoolType, types.BoolValue(value), diags
	case time.Time:
		return types.StringType, types.StringValue(value.Format(time.RFC3339)), diags
	case nil:
		return types.DynamicType, types.DynamicNull(), diags
	case []any:
		elementTypes := make([]attr.Type, len(value))
		elementValues := make([]attr.Value, len(value))
		for i, dynamicElementValue := range value {
			elementType, elementValue, elementDiags := ConvertToTerraformType(dynamicElementValue)
			elementTypes[i] = elementType
			elementValues[i] = elementValue
			diags.Append(elementDiags...)
		}
		if diags.HasError() {
			return nil, nil, diags
		}
		result, tupleDiags := types.TupleValue(elementTypes, elementValues)
		diags.Append(tupleDiags...)
		return types.TupleType{ElemTypes: elementTypes}, result, diags
	case map[string]any:
		attributeTypes := make(map[string]attr.Type, len(value))
		attributeValues := make(map[string]attr.Value, len(value))
		for attributeName, dynamicAttributeValue := range value {
			attributeType, attributeValue, attributeDiags := ConvertToTerraformType(dynamicAttributeValue)
			attributeTypes[attributeName] = attributeType
			attributeValues[attributeName] = attributeValue
			diags.Append(attributeDiags...)
		}
		if diags.HasError() {
			return nil, nil, diags
		}
		result, objectDiags := types.ObjectValue(attributeTypes, attributeValues)
		diags.Append(objectDiags...)
		return types.ObjectType{AttrTypes: attributeTypes}, result, diags
	default:
		diags.AddError(
			"Invalid type to convert to Terraform type",
			fmt.Sprintf("Unable to convert value %v (type %T) to Terraform type", value, value),
		)
		return nil, nil, diags
	}
}
