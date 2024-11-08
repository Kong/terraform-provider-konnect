// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginStatsdAdvancedResource{}
var _ resource.ResourceWithImportState = &GatewayPluginStatsdAdvancedResource{}

func NewGatewayPluginStatsdAdvancedResource() resource.Resource {
	return &GatewayPluginStatsdAdvancedResource{}
}

// GatewayPluginStatsdAdvancedResource defines the resource implementation.
type GatewayPluginStatsdAdvancedResource struct {
	client *sdk.Konnect
}

// GatewayPluginStatsdAdvancedResourceModel describes the resource data model.
type GatewayPluginStatsdAdvancedResourceModel struct {
	Config         *tfTypes.CreateStatsdAdvancedPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                      `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer                      `tfsdk:"consumer_group"`
	ControlPlaneID types.String                              `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                               `tfsdk:"created_at"`
	Enabled        types.Bool                                `tfsdk:"enabled"`
	ID             types.String                              `tfsdk:"id"`
	InstanceName   types.String                              `tfsdk:"instance_name"`
	Ordering       *tfTypes.CreateACLPluginOrdering          `tfsdk:"ordering"`
	Protocols      []types.String                            `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                      `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                      `tfsdk:"service"`
	Tags           []types.String                            `tfsdk:"tags"`
	UpdatedAt      types.Int64                               `tfsdk:"updated_at"`
}

func (r *GatewayPluginStatsdAdvancedResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_statsd_advanced"
}

func (r *GatewayPluginStatsdAdvancedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginStatsdAdvanced Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"allow_status_codes": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of status code ranges that are allowed to be logged in metrics.`,
					},
					"consumer_identifier_default": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The default consumer identifier for metrics. This will take effect when a metric's consumer identifier is omitted. Allowed values are ` + "`" + `custom_id` + "`" + `, ` + "`" + `consumer_id` + "`" + `, ` + "`" + `username` + "`" + `. must be one of ["consumer_id", "custom_id", "username"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"consumer_id",
								"custom_id",
								"username",
							),
						},
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"hostname_in_prefix": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Include the ` + "`" + `hostname` + "`" + ` in the ` + "`" + `prefix` + "`" + ` for each metric name.`,
					},
					"metrics": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"consumer_identifier": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `must be one of ["consumer_id", "custom_id", "username"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"consumer_id",
											"custom_id",
											"username",
										),
									},
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Not Null; must be one of ["kong_latency", "latency", "request_count", "request_per_user", "request_size", "response_size", "status_count", "status_count_per_user", "unique_users", "upstream_latency", "status_count_per_workspace", "status_count_per_user_per_route", "shdict_usage", "cache_datastore_hits_total", "cache_datastore_misses_total"]`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
										stringvalidator.OneOf(
											"kong_latency",
											"latency",
											"request_count",
											"request_per_user",
											"request_size",
											"response_size",
											"status_count",
											"status_count_per_user",
											"unique_users",
											"upstream_latency",
											"status_count_per_workspace",
											"status_count_per_user_per_route",
											"shdict_usage",
											"cache_datastore_hits_total",
											"cache_datastore_misses_total",
										),
									},
								},
								"sample_rate": schema.NumberAttribute{
									Computed: true,
									Optional: true,
								},
								"service_identifier": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `must be one of ["service_id", "service_name", "service_host", "service_name_or_host"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"service_id",
											"service_name",
											"service_host",
											"service_name_or_host",
										),
									},
								},
								"stat_type": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Not Null; must be one of ["counter", "gauge", "histogram", "meter", "set", "timer"]`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
										stringvalidator.OneOf(
											"counter",
											"gauge",
											"histogram",
											"meter",
											"set",
											"timer",
										),
									},
								},
								"workspace_identifier": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `must be one of ["workspace_id", "workspace_name"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"workspace_id",
											"workspace_name",
										),
									},
								},
							},
						},
						Description: `List of Metrics to be logged.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
						Validators: []validator.Int64{
							int64validator.AtMost(65535),
						},
					},
					"prefix": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `String to prefix to each metric's name.`,
					},
					"queue": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"concurrency_limit": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The number of of queue delivery timers. -1 indicates unlimited. must be one of ["-1", "1"]`,
								Validators: []validator.Int64{
									int64validator.OneOf(-1, 1),
								},
							},
							"initial_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Time in seconds before the initial retry is made for a failing batch.`,
							},
							"max_batch_size": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of entries that can be processed at a time.`,
								Validators: []validator.Int64{
									int64validator.Between(1, 1000000),
								},
							},
							"max_bytes": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of bytes that can be waiting on a queue, requires string content.`,
							},
							"max_coalescing_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.`,
							},
							"max_entries": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of entries that can be waiting on the queue.`,
								Validators: []validator.Int64{
									int64validator.Between(1, 1000000),
								},
							},
							"max_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum time in seconds between retries, caps exponential backoff.`,
							},
							"max_retry_time": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Time in seconds before the queue gives up calling a failed handler for a batch.`,
							},
						},
					},
					"service_identifier_default": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The default service identifier for metrics. This will take effect when a metric's service identifier is omitted. Allowed values are ` + "`" + `service_name_or_host` + "`" + `, ` + "`" + `service_id` + "`" + `, ` + "`" + `service_name` + "`" + `, ` + "`" + `service_host` + "`" + `. must be one of ["service_id", "service_name", "service_host", "service_name_or_host"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"service_id",
								"service_name",
								"service_host",
								"service_name_or_host",
							),
						},
					},
					"udp_packet_size": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Combine UDP packet up to the size configured. If zero (0), don't combine the UDP packet. Must be a number between 0 and 65507 (inclusive).`,
					},
					"use_tcp": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Use TCP instead of UDP.`,
					},
					"workspace_identifier_default": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The default workspace identifier for metrics. This will take effect when a metric's workspace identifier is omitted. Allowed values are ` + "`" + `workspace_id` + "`" + `, ` + "`" + `workspace_name` + "`" + `. must be one of ["workspace_id", "workspace_name"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"workspace_id",
								"workspace_name",
							),
						},
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
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
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
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

func (r *GatewayPluginStatsdAdvancedResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginStatsdAdvancedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginStatsdAdvancedResourceModel
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

	createStatsdAdvancedPlugin := data.ToSharedCreateStatsdAdvancedPlugin()
	request := operations.CreateStatsdadvancedPluginRequest{
		ControlPlaneID:             controlPlaneID,
		CreateStatsdAdvancedPlugin: createStatsdAdvancedPlugin,
	}
	res, err := r.client.Plugins.CreateStatsdadvancedPlugin(ctx, request)
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
	if !(res.StatsdAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedStatsdAdvancedPlugin(res.StatsdAdvancedPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginStatsdAdvancedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginStatsdAdvancedResourceModel
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

	request := operations.GetStatsdadvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetStatsdadvancedPlugin(ctx, request)
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
	if !(res.StatsdAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedStatsdAdvancedPlugin(res.StatsdAdvancedPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginStatsdAdvancedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginStatsdAdvancedResourceModel
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

	createStatsdAdvancedPlugin := data.ToSharedCreateStatsdAdvancedPlugin()
	request := operations.UpdateStatsdadvancedPluginRequest{
		PluginID:                   pluginID,
		ControlPlaneID:             controlPlaneID,
		CreateStatsdAdvancedPlugin: createStatsdAdvancedPlugin,
	}
	res, err := r.client.Plugins.UpdateStatsdadvancedPlugin(ctx, request)
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
	if !(res.StatsdAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedStatsdAdvancedPlugin(res.StatsdAdvancedPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginStatsdAdvancedResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginStatsdAdvancedResourceModel
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

	request := operations.DeleteStatsdadvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteStatsdadvancedPlugin(ctx, request)
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

func (r *GatewayPluginStatsdAdvancedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
