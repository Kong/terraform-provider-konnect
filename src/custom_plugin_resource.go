package custom

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfReflect "github.com/kong/terraform-provider-konnect/internal/provider/reflect"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CustomPluginResource{}
var _ resource.ResourceWithImportState = &CustomPluginResource{}

func NewCustomPluginResource() resource.Resource {
	return &CustomPluginResource{}
}

type CustomPluginResourceModel struct {
	ID             types.String  `tfsdk:"id"`
	Name           types.String  `tfsdk:"name"`
	Config         types.Dynamic `tfsdk:"config"`
	ControlPlaneID types.String  `tfsdk:"control_plane_id"`
}

func (r *CustomPluginResourceModel) ToSharedPluginInput() (shared.PluginInput, error) {
	config := r.Config.UnderlyingValue()

	// Use .String() to convert the dynamic value to a string,
	// then unmarshal the JSON string into a map[string]any
	var configJson map[string]any
	err := json.Unmarshal([]byte(config.String()), &configJson)
	if err != nil {
		return shared.PluginInput{}, err
	}

	return shared.PluginInput{
		Name:   sdk.String(r.Name.ValueString()),
		Config: configJson,
	}, nil
}

type CustomPluginResource struct {
	client *sdk.Konnect
}

func (r *CustomPluginResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_custom_plugin"
}

func (r *CustomPluginResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Gateway Custom Plugin Resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Plugin ID`,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{},
				Required:      true,
				Description:   `Plugin Name`,
			},
			"config": schema.DynamicAttribute{
				PlanModifiers: []planmodifier.Dynamic{},
				Required:      true,
				Description:   `Configuration`,
			},
			"control_plane_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{},
				Required:      true,
				Description:   `Control Plane ID`,
			},
		},
	}
}

func (r *CustomPluginResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CustomPluginResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plugin *CustomPluginResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &plugin, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := plugin.ControlPlaneID.ValueString()
	pluginData, err := plugin.ToSharedPluginInput()
	if err != nil {
		resp.Diagnostics.AddError("failed to convert PluginInput", err.Error())
		return
	}

	request := operations.CreatePluginRequest{
		ControlPlaneID: controlPlaneID,
		Plugin:         pluginData,
	}
	res, err := r.client.Plugins.CreatePlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.Plugin == nil {
		resp.Diagnostics.AddError("unexpected response from API.", "No response body")
		return
	}
	plugin.RefreshFromResponse(ctx, r.client, resp.Diagnostics, res.Plugin)
	refreshPlan(ctx, plan, &plugin, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plugin)...)
}

func merge(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, target interface{}) {
	var plan types.Object
	var state types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	val, err := state.ToTerraformValue(ctx)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Object Conversion Error", "An unexpected error was encountered trying to convert object. This is always an error in the provider. Please report the following to the provider developer:\n\n"+err.Error()))
		return
	}
	resp.Diagnostics.Append(tfReflect.Into(ctx, types.ObjectType{AttrTypes: state.AttributeTypes(ctx)}, val, target, tfReflect.Options{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	}, path.Empty())...)
	if resp.Diagnostics.HasError() {
		return
	}

	refreshPlan(ctx, plan, target, resp.Diagnostics)
}

func refreshPlan(ctx context.Context, plan types.Object, target interface{}, diagnostics diag.Diagnostics) {
	obj := types.ObjectType{AttrTypes: plan.AttributeTypes(ctx)}
	val, err := plan.ToTerraformValue(ctx)
	if err != nil {
		diagnostics.Append(diag.NewErrorDiagnostic("Object Conversion Error", "An unexpected error was encountered trying to convert object. This is always an error in the provider. Please report the following to the provider developer:\n\n"+err.Error()))
		return
	}
	diagnostics.Append(tfReflect.Into(ctx, obj, val, target, tfReflect.Options{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
		SourceType:              tfReflect.SourceTypePlan,
	}, path.Empty())...)
}

func (r *CustomPluginResourceModel) RefreshFromResponse_WORKING(ctx context.Context, diags diag.Diagnostics, resp *shared.Plugin) {
	r.ID = types.StringPointerValue(resp.ID)
	r.Name = types.StringPointerValue(resp.Name)

	c, d := types.MapValueFrom(ctx, types.DynamicType, resp.Config)
	if d != nil {
		diags.Append(d...)
	}

	r.Config = types.DynamicValue(c)
}

type PluginSchema struct {
	Fields []map[string]PluginField `json:"fields"`
}
type PluginField struct {
	Type    string                   `json:"type"`
	Default interface{}              `json:"default"`
	Fields  []map[string]PluginField `json:"fields"`
}

func pruneConfigValues(ctx context.Context, config map[string]interface{}, configSchema []map[string]PluginField) {
	for _, fieldConfig := range configSchema {
		for fieldKey, field := range fieldConfig {
			configValue, ok := config[fieldKey]
			if ok {
				// If field type is record, then recurse
				if field.Type == "record" {
					childConfigValue := configValue.(map[string]interface{})
					pruneConfigValues(ctx, childConfigValue, field.Fields)
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

func extractPluginConfigSchema(fields []map[string]any, diags diag.Diagnostics) (PluginSchema, diag.Diagnostics) {
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

func (r *CustomPluginResourceModel) RefreshFromResponse(ctx context.Context, client *sdk.Konnect, diags diag.Diagnostics, resp *shared.Plugin) {
	r.ID = types.StringPointerValue(resp.ID)
	r.Name = types.StringPointerValue(resp.Name)

	// Remove
	schema, err := client.Plugins.FetchPluginSchema(ctx, operations.FetchPluginSchemaRequest{
		PluginName:     r.Name.ValueString(),
		ControlPlaneID: r.ControlPlaneID.ValueString(),
	})

	if err != nil {
		diags.AddError("failed to fetch plugin schema", err.Error())
		return
	}

	config, diag := extractPluginConfigSchema(schema.Object.Fields, diags)
	if diag.HasError() {
		diags.Append(diag...)
		return
	}

	// Find the schema
	pruneConfigValues(ctx, resp.Config, config.Fields)

	_, c, diag := convertToTerraformType(resp.Config)
	if diag.HasError() {
		diags.Append(diag...)
	}

	r.Config = types.DynamicValue(c)
}

func (r *CustomPluginResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CustomPluginResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.GetPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       data.ID.ValueString(),
	}
	res, err := r.client.Plugins.GetPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", "")
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), "")
		return
	}
	if res.Plugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", "")
		return
	}

	// Remove any fields that are not in the plan
	// TODO: This is a temporary workaround until the framework provides a better way to handle this
	// TODO: Add plugin schema endpoint to the generated SDK

	data.RefreshFromResponse(ctx, r.client, resp.Diagnostics, res.Plugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomPluginResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CustomPluginResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Load data from state
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	// Load data from plan
	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)

	controlPlaneID := data.ControlPlaneID.ValueString()
	plugin, err := data.ToSharedPluginInput()
	if err != nil {
		resp.Diagnostics.AddError("failed to convert PluginInput", err.Error())
		return
	}

	request := operations.UpsertPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       data.ID.ValueString(),
		Plugin:         plugin,
	}
	res, err := r.client.Plugins.UpsertPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError("unexpected response from API.", fmt.Sprintf("Got an unexpected response code %v", res.StatusCode))
		return
	}
	if res.Plugin == nil {
		resp.Diagnostics.AddError("unexpected response from API.", "No response body")
		return
	}
	data.RefreshFromResponse(ctx, r.client, resp.Diagnostics, res.Plugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	data.ID = types.StringPointerValue(res.Plugin.ID)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CustomPluginResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CustomPluginResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.DeletePluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       data.ID.ValueString(),
	}
	res, err := r.client.Plugins.DeletePlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", "")
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), "")
		return
	}
}

func (r *CustomPluginResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_custom_plugin.")
}

func convertToTerraformType(dynamicValue any) (attr.Type, attr.Value, diag.Diagnostics) {
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
			elementType, elementValue, elementDiags := convertToTerraformType(dynamicElementValue)
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
			attributeType, attributeValue, attributeDiags := convertToTerraformType(dynamicAttributeValue)
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
