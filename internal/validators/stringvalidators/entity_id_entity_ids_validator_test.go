package stringvalidators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// unknownIDs stands for an entity_ids list that is not known at plan time, for
// example one that references an attribute of a resource that does not exist
// yet.
type unknownIDs struct{}

// roleSchema mirrors the subset of the role resources the validator reads:
// entity_id is the validated attribute, entity_ids is its sibling.
func roleSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"entity_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"entity_ids": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

// encode converts an entity_id fixture to its tftypes value: nil is null, a
// string is known.
func encode(entityID any) tftypes.Value {
	switch v := entityID.(type) {
	case nil:
		return tftypes.NewValue(tftypes.String, nil)
	case string:
		return tftypes.NewValue(tftypes.String, v)
	default:
		panic("unsupported entity_id fixture type")
	}
}

// encodeList converts an entity_ids fixture to its tftypes value: nil is a null
// list, unknownIDs an unknown one, a []string a known list in the given order.
func encodeList(entityIDs any) tftypes.Value {
	listType := tftypes.List{ElementType: tftypes.String}

	switch v := entityIDs.(type) {
	case nil:
		return tftypes.NewValue(listType, nil)
	case unknownIDs:
		return tftypes.NewValue(listType, tftypes.UnknownValue)
	case []string:
		elems := make([]tftypes.Value, 0, len(v))
		for _, s := range v {
			elems = append(elems, tftypes.NewValue(tftypes.String, s))
		}
		return tftypes.NewValue(listType, elems)
	default:
		panic("unsupported entity_ids fixture type")
	}
}

// request builds the validation request the framework sends at plan time. The
// config value of entity_id and the value in the plan config always agree.
func request(t *testing.T, entityID any, entityIDs any) validator.StringRequest {
	t.Helper()

	s := roleSchema()
	objType := s.Type().TerraformType(context.Background()).(tftypes.Object)

	configValue := types.StringNull()
	if v, ok := entityID.(string); ok {
		configValue = types.StringValue(v)
	}

	return validator.StringRequest{
		Path: path.Root("entity_id"),
		Config: tfsdk.Config{
			Schema: s,
			Raw: tftypes.NewValue(objType, map[string]tftypes.Value{
				"entity_id":  encode(entityID),
				"entity_ids": encodeList(entityIDs),
			}),
		},
		ConfigValue: configValue,
	}
}

func TestEntityIDEntityIdsValidator(t *testing.T) {
	cp := "5f2d8a1e-0f6b-4b3e-9a4d-1c2b3a4d5e6f"
	ws := "7c9e4f2a-3b1d-4c5e-8f6a-9b0c1d2e3f4a"

	tests := []struct {
		name       string
		entityID   any
		entityIDs  any
		wantErrors int
	}{
		{
			name:       "entity_id only",
			entityID:   cp,
			entityIDs:  nil,
			wantErrors: 0,
		},
		{
			name:       "entity_ids only",
			entityID:   nil,
			entityIDs:  []string{cp},
			wantErrors: 0,
		},
		{
			name:       "entity_ids with a workspace parent and child",
			entityID:   nil,
			entityIDs:  []string{cp, ws},
			wantErrors: 0,
		},
		{
			name:       "both set",
			entityID:   cp,
			entityIDs:  []string{cp},
			wantErrors: 1,
		},
		{
			name:       "both set with two ids",
			entityID:   cp,
			entityIDs:  []string{cp, ws},
			wantErrors: 1,
		},
		{
			// The validator checks the conflict only. The API rejects a body
			// with neither field (400-entity-id-or-entity-ids-required).
			name:       "neither set",
			entityID:   nil,
			entityIDs:  nil,
			wantErrors: 0,
		},
		{
			name:       "empty entity_ids",
			entityID:   cp,
			entityIDs:  []string{},
			wantErrors: 0,
		},
		{
			name:       "entity_ids unknown",
			entityID:   cp,
			entityIDs:  unknownIDs{},
			wantErrors: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &validator.StringResponse{}
			EntityIDEntityIdsValidator().ValidateString(context.Background(), request(t, tt.entityID, tt.entityIDs), resp)

			if got := resp.Diagnostics.ErrorsCount(); got != tt.wantErrors {
				t.Fatalf("got %d errors, want %d: %v", got, tt.wantErrors, resp.Diagnostics)
			}
			if tt.wantErrors == 0 {
				return
			}

			diag := resp.Diagnostics.Errors()[0]
			if diag.Summary() != "Conflicting Configuration" {
				t.Errorf("got summary %q, want %q", diag.Summary(), "Conflicting Configuration")
			}
			if diag.Detail() != "The 'entity_id' and 'entity_ids' fields are mutually exclusive. Only one of them can be set." {
				t.Errorf("unexpected detail: %q", diag.Detail())
			}
		})
	}
}

// TestEntityIDEntityIdsValidatorUnknownEntityID covers a plan where entity_id
// itself is not known yet. The validator must stay quiet and let a later plan
// do the check.
func TestEntityIDEntityIdsValidatorUnknownEntityID(t *testing.T) {
	s := roleSchema()
	objType := s.Type().TerraformType(context.Background()).(tftypes.Object)

	req := validator.StringRequest{
		Path: path.Root("entity_id"),
		Config: tfsdk.Config{
			Schema: s,
			Raw: tftypes.NewValue(objType, map[string]tftypes.Value{
				"entity_id":  tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"entity_ids": encodeList([]string{"5f2d8a1e-0f6b-4b3e-9a4d-1c2b3a4d5e6f"}),
			}),
		},
		ConfigValue: types.StringUnknown(),
	}

	resp := &validator.StringResponse{}
	EntityIDEntityIdsValidator().ValidateString(context.Background(), req, resp)

	if got := resp.Diagnostics.ErrorsCount(); got != 0 {
		t.Fatalf("got %d errors, want 0: %v", got, resp.Diagnostics)
	}
}
