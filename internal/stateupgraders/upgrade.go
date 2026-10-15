// Package stateupgraders backfills Terraform state for resources that gained
// a new attribute (e.g. `workspace`) after they were already in use.
package stateupgraders

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var schemaTypes = map[string]tftypes.Type{}

func RegisterSchemaType(resourceTypeName string, t tftypes.Type) {
	schemaTypes[resourceTypeName] = t
}

// Decodes prior state as raw JSON, applies transform to backfill it, then
// re-decodes against the current schema - so fields the transform doesn't
// touch survive schema drift, unlike a hand-written field-by-field copy.
func upgradeToCurrentSchema(resourceTypeName string, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse, transform func(map[string]any)) {
	schemaType, ok := schemaTypes[resourceTypeName]
	if !ok {
		resp.Diagnostics.AddError("State upgrade failed", fmt.Sprintf("no schema type registered for %q", resourceTypeName))
		return
	}

	var raw map[string]any
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		resp.Diagnostics.AddError("State upgrade failed", "could not parse prior state JSON: "+err.Error())
		return
	}

	if transform != nil {
		transform(raw)
	}

	upgraded, err := json.Marshal(raw)
	if err != nil {
		resp.Diagnostics.AddError("State upgrade failed", "could not re-marshal upgraded state: "+err.Error())
		return
	}

	value, err := (tfprotov6.RawState{JSON: upgraded}).UnmarshalWithOpts(schemaType, tfprotov6.UnmarshalOpts{
		ValueFromJSONOpts: tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true},
	})
	if err != nil {
		resp.Diagnostics.AddError("State upgrade failed", "could not decode upgraded state against current schema: "+err.Error())
		return
	}

	dynamicValue, err := tfprotov6.NewDynamicValue(schemaType, value)
	if err != nil {
		resp.Diagnostics.AddError("State upgrade failed", "could not encode upgraded state: "+err.Error())
		return
	}

	resp.DynamicValue = &dynamicValue
}
