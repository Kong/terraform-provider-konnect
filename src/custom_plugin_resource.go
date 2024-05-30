package custom

import (
	"context"
	"fmt"

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
	"github.com/kong/terraform-provider-konnect/src/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CustomPluginResource{}
var _ resource.ResourceWithImportState = &CustomPluginResource{}

func NewCustomPluginResource() resource.Resource {
	return &CustomPluginResource{}
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

	pluginData, err := plugin.ToSharedPluginInput()
	if err != nil {
		resp.Diagnostics.AddError("failed to convert PluginInput", err.Error())
		return
	}

	res, err := r.client.Plugins.CreatePlugin(ctx, operations.CreatePluginRequest{
		ControlPlaneID: plugin.ControlPlaneID.ValueString(),
		Plugin:         pluginData,
	})

	checkPluginResponse(res, err, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	plugin.RefreshFromResponse(ctx, r.client, resp.Diagnostics, res.Plugin)
	refreshPlan(ctx, plan, &plugin, resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plugin)...)
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

	res, err := r.client.Plugins.GetPlugin(ctx, operations.GetPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       data.ID.ValueString(),
	})

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
	if res.Plugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", "")
		return
	}

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

	// Merge state + plan
	utils.Merge(ctx, req, resp, &data)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	plugin, err := data.ToSharedPluginInput()
	if err != nil {
		resp.Diagnostics.AddError("failed to convert PluginInput", err.Error())
		return
	}

	res, err := r.client.Plugins.UpsertPlugin(ctx, operations.UpsertPluginRequest{
		ControlPlaneID: data.ControlPlaneID.ValueString(),
		PluginID:       data.ID.ValueString(),
		Plugin:         plugin,
	})

	checkPluginResponse(res, err, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	data.RefreshFromResponse(ctx, r.client, resp.Diagnostics, res.Plugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

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

	res, err := r.client.Plugins.DeletePlugin(ctx, operations.DeletePluginRequest{
		ControlPlaneID: data.ControlPlaneID.ValueString(),
		PluginID:       data.ID.ValueString(),
	})

	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}

	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
}

func (r *CustomPluginResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_custom_plugin.")
}

func (r *CustomPluginResourceModel) RefreshFromResponse(ctx context.Context, client *sdk.Konnect, diags diag.Diagnostics, resp *shared.Plugin) {
	r.ID = types.StringPointerValue(resp.ID)
	r.Name = types.StringPointerValue(resp.Name)

	// Remove default values from the config
	// Fetch plugin schema + default values
	schema, err := client.Plugins.FetchPluginSchema(ctx, operations.FetchPluginSchemaRequest{
		PluginName:     r.Name.ValueString(),
		ControlPlaneID: r.ControlPlaneID.ValueString(),
	})

	if err != nil {
		diags.AddError("failed to fetch plugin schema", err.Error())
		return
	}

	// Extract the plugin schema from the response
	config, diag := utils.ExtractPluginConfigSchema(schema.Object.Fields, diags)
	if diag.HasError() {
		diags.Append(diag...)
		return
	}

	// Remove default values
	utils.PruneConfigValues(ctx, resp.Config, config.Fields)

	// Convert to Terraform types
	_, c, diag := utils.ConvertToTerraformType(resp.Config)
	if diag.HasError() {
		diags.Append(diag...)
	}

	r.Config = types.DynamicValue(c)
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

func checkPluginResponse(res any, err error, diagnostics *diag.Diagnostics) {
	if err != nil {
		diagnostics.AddError("failure to invoke API", err.Error())
		return
	}

	if res == nil {
		diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}

	plugin := getPluginFromResponse(res)

	if plugin == nil {
		diagnostics.AddError("could not cast API response to plugin", fmt.Sprintf("%T", res))
		return
	}
}

func getPluginFromResponse(res any) *shared.Plugin {
	switch t := res.(type) {
	case *operations.GetPluginResponse:
		return t.Plugin
	case *operations.UpsertPluginResponse:
		return t.Plugin
	case *operations.CreatePluginResponse:
		return t.Plugin
	default:
		return nil
	}
}
