// Package dynamic holds the hand-written runtime conversion helpers used by the
// generated provider for fields modeled as a Terraform dynamic attribute.
//
// Speakeasy does not support types.Dynamic / schema.DynamicAttribute, so some
// genuinely polymorphic Kong fields (the canonical case is redis.port, which is
// normally an integer but may also be a referenceable vault string such as
// "{vault://...}") are rewired to a dynamic attribute by the post-processing
// layer (tools/postprocess). That layer only ever rewrites call sites to the two
// symbols exported here; the conversion logic lives in normal, unit-tested Go so
// it is never regenerated.
package dynamic

import (
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ToAny converts a dynamic value coming from Terraform into the plain Go value
// sent to the SDK/API. It supports a numeric value as well as any other concrete
// value (e.g. a string vault reference). A null or unknown value yields nil so
// the field is omitted.
func ToAny(d types.Dynamic) any {
	if d.IsNull() || d.IsUnknown() {
		return nil
	}

	switch v := d.UnderlyingValue().(type) {
	case types.Number:
		if v.IsNull() || v.IsUnknown() {
			return nil
		}
		bf := v.ValueBigFloat()
		if bf == nil {
			return nil
		}
		if i, acc := bf.Int64(); acc == big.Exact {
			return i
		}
		f, _ := bf.Float64()
		return f
	case types.Int64:
		return v.ValueInt64()
	case types.Float64:
		return v.ValueFloat64()
	case types.String:
		return v.ValueString()
	case types.Bool:
		return v.ValueBool()
	default:
		// Fall back to the raw underlying value; this keeps other concrete
		// kinds flowing through unchanged.
		if v.IsNull() || v.IsUnknown() {
			return nil
		}
		return v
	}
}

// ToFramework converts a value returned by the SDK/API into a Terraform dynamic
// value. Numeric values are normalized to a Number so they round-trip
// consistently with numbers supplied in HCL config; strings and bools are
// preserved as-is. A nil value (including a typed nil pointer) yields a null
// dynamic value.
func ToFramework(v any) types.Dynamic {
	switch val := v.(type) {
	case nil:
		return types.DynamicNull()
	case int64:
		return types.DynamicValue(types.NumberValue(new(big.Float).SetInt64(val)))
	case int:
		return types.DynamicValue(types.NumberValue(new(big.Float).SetInt64(int64(val))))
	case float64:
		// JSON numbers decode as float64; represent them as Number to match
		// the concrete type Terraform produces for numeric HCL values.
		return types.DynamicValue(types.NumberValue(big.NewFloat(val)))
	case float32:
		return types.DynamicValue(types.NumberValue(big.NewFloat(float64(val))))
	case *int64:
		if val == nil {
			return types.DynamicNull()
		}
		return types.DynamicValue(types.NumberValue(new(big.Float).SetInt64(*val)))
	case *int:
		if val == nil {
			return types.DynamicNull()
		}
		return types.DynamicValue(types.NumberValue(new(big.Float).SetInt64(int64(*val))))
	case *float64:
		if val == nil {
			return types.DynamicNull()
		}
		return types.DynamicValue(types.NumberValue(big.NewFloat(*val)))
	case *string:
		if val == nil {
			return types.DynamicNull()
		}
		return types.DynamicValue(types.StringValue(*val))
	case string:
		return types.DynamicValue(types.StringValue(val))
	case *bool:
		if val == nil {
			return types.DynamicNull()
		}
		return types.DynamicValue(types.BoolValue(*val))
	case bool:
		return types.DynamicValue(types.BoolValue(val))
	default:
		return types.DynamicNull()
	}
}
