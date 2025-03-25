// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginSyslogResource{}
var _ resource.ResourceWithImportState = &GatewayPluginSyslogResource{}

func NewGatewayPluginSyslogResource() resource.Resource {
	return &GatewayPluginSyslogResource{}
}

// GatewayPluginSyslogResource defines the resource implementation.
type GatewayPluginSyslogResource struct {
	client *sdk.Konnect
}

// GatewayPluginSyslogResourceModel describes the resource data model.
type GatewayPluginSyslogResourceModel struct {
	Config         tfTypes.SyslogPluginConfig         `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer"`
	ControlPlaneID types.String                       `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                        `tfsdk:"created_at"`
	Enabled        types.Bool                         `tfsdk:"enabled"`
	ID             types.String                       `tfsdk:"id"`
	InstanceName   types.String                       `tfsdk:"instance_name"`
	Ordering       map[string]types.String            `tfsdk:"ordering"`
	Protocols      []types.String                     `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer `tfsdk:"route"`
	Service        *tfTypes.ACLWithoutParentsConsumer `tfsdk:"service"`
	Tags           []types.String                     `tfsdk:"tags"`
	UpdatedAt      types.Int64                        `tfsdk:"updated_at"`
}

func (r *GatewayPluginSyslogResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_syslog"
}

func (r *GatewayPluginSyslogResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginSyslog Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"client_errors_severity": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["alert", "crit", "debug", "emerg", "err", "info", "notice", "warning"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"alert",
								"crit",
								"debug",
								"emerg",
								"err",
								"info",
								"notice",
								"warning",
							),
						},
					},
					"custom_fields_by_lua": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `Lua code as a key-value map`,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
					},
					"facility": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The facility is used by the operating system to decide how to handle each log message. must be one of ["auth", "authpriv", "cron", "daemon", "ftp", "kern", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7", "lpr", "mail", "news", "syslog", "user", "uucp"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"auth",
								"authpriv",
								"cron",
								"daemon",
								"ftp",
								"kern",
								"local0",
								"local1",
								"local2",
								"local3",
								"local4",
								"local5",
								"local6",
								"local7",
								"lpr",
								"mail",
								"news",
								"syslog",
								"user",
								"uucp",
							),
						},
					},
					"log_level": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["alert", "crit", "debug", "emerg", "err", "info", "notice", "warning"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"alert",
								"crit",
								"debug",
								"emerg",
								"err",
								"info",
								"notice",
								"warning",
							),
						},
					},
					"server_errors_severity": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["alert", "crit", "debug", "emerg", "err", "info", "notice", "warning"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"alert",
								"crit",
								"debug",
								"emerg",
								"err",
								"info",
								"notice",
								"warning",
							),
						},
					},
					"successful_severity": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `must be one of ["alert", "crit", "debug", "emerg", "err", "info", "notice", "warning"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"alert",
								"crit",
								"debug",
								"emerg",
								"err",
								"info",
								"notice",
								"warning",
							),
						},
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"control_plane_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.MapAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A set of strings representing protocols.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayPluginSyslogResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginSyslogResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginSyslogResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	syslogPlugin := *data.ToSharedSyslogPluginInput()
	request := operations.CreateSyslogPluginRequest{
		ControlPlaneID: controlPlaneID,
		SyslogPlugin:   syslogPlugin,
	}
	res, err := r.client.Plugins.CreateSyslogPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.SyslogPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSyslogPlugin(res.SyslogPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginSyslogResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginSyslogResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetSyslogPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetSyslogPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
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
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.SyslogPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSyslogPlugin(res.SyslogPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginSyslogResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginSyslogResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	syslogPlugin := *data.ToSharedSyslogPluginInput()
	request := operations.UpdateSyslogPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
		SyslogPlugin:   syslogPlugin,
	}
	res, err := r.client.Plugins.UpdateSyslogPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.SyslogPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSyslogPlugin(res.SyslogPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginSyslogResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginSyslogResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.DeleteSyslogPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteSyslogPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GatewayPluginSyslogResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "plugin_id": "3473c251-5b6c-4f45-b1ff-7ede735a366d"}': `+err.Error())
		return
	}

	if len(data.ControlPlaneID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field control_plane_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9524ec7d-36d9-465d-a8c5-83a3c9390458"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("control_plane_id"), data.ControlPlaneID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"3473c251-5b6c-4f45-b1ff-7ede735a366d"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
