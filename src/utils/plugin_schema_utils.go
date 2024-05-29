package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type PluginSchema struct {
	Fields []map[string]PluginField `json:"fields"`
}
type PluginField struct {
	Type    string                   `json:"type"`
	Default interface{}              `json:"default"`
	Fields  []map[string]PluginField `json:"fields"`
}

func PruneConfigValues(ctx context.Context, config map[string]interface{}, configSchema []map[string]PluginField) {
	for _, fieldConfig := range configSchema {
		for fieldKey, field := range fieldConfig {
			configValue, ok := config[fieldKey]
			if ok {
				// If field type is record, then recurse
				if field.Type == "record" {
					childConfigValue := configValue.(map[string]interface{})
					PruneConfigValues(ctx, childConfigValue, field.Fields)
					// If all fields of child config are defaults, then remove entire child config from parent
					if len(childConfigValue) == 0 {
						delete(config, fieldKey)
					}
				} else if areConfigValuesEqual(ctx, fieldKey, configValue, field.Default) {
					delete(config, fieldKey)
				}
			}
		}
	}
}

func areConfigValuesEqual(ctx context.Context, key string, configValue interface{}, schemaDefault interface{}) bool {
	configValueJSON, _ := json.Marshal(configValue)
	configValueString := string(configValueJSON[:])
	schemaFieldDefaultJSON, _ := json.Marshal(schemaDefault)
	schemaFieldString := string(schemaFieldDefaultJSON[:])
	return configValueString == schemaFieldString
}

func ExtractPluginConfigSchema(fields []map[string]any, diags diag.Diagnostics) (PluginSchema, diag.Diagnostics) {
	// Each field is a map with a single key
	var configSchema any
	for _, value := range fields {
		valueMap, ok := value["config"]
		if ok {
			configSchema, ok = valueMap.(PluginSchema)
			if !ok {
				diags.AddError("failed to extract plugin config schema", fmt.Sprintf("expected PluginSchema, got %T", valueMap))
				return PluginSchema{}, diags
			}
			break
		}
	}

	return configSchema.(PluginSchema), diags
}
