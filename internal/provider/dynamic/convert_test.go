package dynamic

import (
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestToAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   types.Dynamic
		want any
	}{
		{"null", types.DynamicNull(), nil},
		{"unknown", types.DynamicUnknown(), nil},
		{
			name: "integer number",
			in:   types.DynamicValue(types.NumberValue(big.NewFloat(6379))),
			want: int64(6379),
		},
		{
			name: "fractional number",
			in:   types.DynamicValue(types.NumberValue(big.NewFloat(1.5))),
			want: 1.5,
		},
		{
			name: "string vault reference",
			in:   types.DynamicValue(types.StringValue("{vault://my-vault/redis-port}")),
			want: "{vault://my-vault/redis-port}",
		},
		{
			name: "bool",
			in:   types.DynamicValue(types.BoolValue(true)),
			want: true,
		},
		{
			name: "int64 underlying",
			in:   types.DynamicValue(types.Int64Value(7000)),
			want: int64(7000),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, ToAny(tt.in))
		})
	}
}

func TestToFramework(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   any
		want types.Dynamic
	}{
		{"nil", nil, types.DynamicNull()},
		{"typed nil string pointer", (*string)(nil), types.DynamicNull()},
		{"typed nil int64 pointer", (*int64)(nil), types.DynamicNull()},
		{
			name: "int64",
			in:   int64(6379),
			want: types.DynamicValue(types.NumberValue(new(big.Float).SetInt64(6379))),
		},
		{
			name: "float64 whole",
			in:   float64(6379),
			want: types.DynamicValue(types.NumberValue(big.NewFloat(6379))),
		},
		{
			name: "string",
			in:   "{vault://my-vault/redis-port}",
			want: types.DynamicValue(types.StringValue("{vault://my-vault/redis-port}")),
		},
		{
			name: "string pointer",
			in:   ptr("6379"),
			want: types.DynamicValue(types.StringValue("6379")),
		},
		{
			name: "bool",
			in:   true,
			want: types.DynamicValue(types.BoolValue(true)),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.True(t, tt.want.Equal(ToFramework(tt.in)), "got %v want %v", ToFramework(tt.in), tt.want)
		})
	}
}

// TestRoundTrip ensures a numeric value supplied to the SDK survives the
// API -> framework -> API round-trip with a stable concrete type.
func TestRoundTrip(t *testing.T) {
	t.Parallel()

	// number round-trip
	fw := ToFramework(int64(6379))
	assert.Equal(t, int64(6379), ToAny(fw))

	// string round-trip
	fw = ToFramework("{vault://v/p}")
	assert.Equal(t, "{vault://v/p}", ToAny(fw))
}

func ptr[T any](v T) *T { return &v }
