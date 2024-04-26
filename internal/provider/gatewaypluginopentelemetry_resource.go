// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
	"math/big"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginOpentelemetryResource{}
var _ resource.ResourceWithImportState = &GatewayPluginOpentelemetryResource{}

func NewGatewayPluginOpentelemetryResource() resource.Resource {
	return &GatewayPluginOpentelemetryResource{}
}

// GatewayPluginOpentelemetryResource defines the resource implementation.
type GatewayPluginOpentelemetryResource struct {
	client *sdk.Konnect
}

// GatewayPluginOpentelemetryResourceModel describes the resource data model.
type GatewayPluginOpentelemetryResourceModel struct {
	Config         tfTypes.CreateOpentelemetryPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                    `tfsdk:"consumer"`
	ControlPlaneID types.String                            `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                             `tfsdk:"created_at"`
	Enabled        types.Bool                              `tfsdk:"enabled"`
	ID             types.String                            `tfsdk:"id"`
	Protocols      []types.String                          `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                    `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                    `tfsdk:"service"`
	Tags           []types.String                          `tfsdk:"tags"`
}

func (r *GatewayPluginOpentelemetryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_opentelemetry"
}

func (r *GatewayPluginOpentelemetryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginOpentelemetry Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"batch_flush_delay": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The delay, in seconds, between two consecutive batches.`,
					},
					"batch_span_count": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of spans to be sent in a single batch.`,
					},
					"connect_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(1000),
						Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2. Default: 1000`,
					},
					"endpoint": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"header_type": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("preserve"),
						Description: `must be one of ["preserve", "ignore", "b3", "b3-single", "w3c", "jaeger", "ot", "aws", "gcp", "datadog"]; Default: "preserve"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"preserve",
								"ignore",
								"b3",
								"b3-single",
								"w3c",
								"jaeger",
								"ot",
								"aws",
								"gcp",
								"datadog",
							),
						},
					},
					"headers": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `The custom headers to be added in the HTTP request sent to the OTLP server. This setting is useful for adding the authentication headers (token) for the APM backend.`,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
					},
					"http_response_header_for_traceid": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"queue": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"initial_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Default:     numberdefault.StaticBigFloat(big.NewFloat(0.01)),
								Description: `Time in seconds before the initial retry is made for a failing batch. Default: 0.01`,
							},
							"max_batch_size": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(1),
								Description: `Maximum number of entries that can be processed at a time. Default: 1`,
							},
							"max_bytes": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum number of bytes that can be waiting on a queue, requires string content.`,
							},
							"max_coalescing_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Default:     numberdefault.StaticBigFloat(big.NewFloat(1)),
								Description: `Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler. Default: 1`,
							},
							"max_entries": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(10000),
								Description: `Maximum number of entries that can be waiting on the queue. Default: 10000`,
							},
							"max_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Default:     numberdefault.StaticBigFloat(big.NewFloat(60)),
								Description: `Maximum time in seconds between retries, caps exponential backoff. Default: 60`,
							},
							"max_retry_time": schema.NumberAttribute{
								Computed:    true,
								Optional:    true,
								Default:     numberdefault.StaticBigFloat(big.NewFloat(60)),
								Description: `Time in seconds before the queue gives up calling a failed handler for a batch. Default: 60`,
							},
						},
					},
					"read_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(5000),
						Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2. Default: 5000`,
					},
					"resource_attributes": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
					},
					"sampling_rate": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Tracing sampling rate for configuring the probability-based sampler. When set, this value supersedes the global ` + "`" + `tracing_sampling_rate` + "`" + ` setting from kong.conf.`,
					},
					"send_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(5000),
						Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2. Default: 5000`,
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
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Description: `Whether the plugin is applied. Default: true`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the Plugin to lookup`,
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
		},
	}
}

func (r *GatewayPluginOpentelemetryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginOpentelemetryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginOpentelemetryResourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	createOpentelemetryPlugin := data.ToSharedCreateOpentelemetryPlugin()
	request := operations.CreateOpentelemetryPluginRequest{
		ControlPlaneID:            controlPlaneID,
		CreateOpentelemetryPlugin: createOpentelemetryPlugin,
	}
	res, err := r.client.Plugins.CreateOpentelemetryPlugin(ctx, request)
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
	if res.OpentelemetryPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOpentelemetryPlugin(res.OpentelemetryPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginOpentelemetryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginOpentelemetryResourceModel
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

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.GetOpentelemetryPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetOpentelemetryPlugin(ctx, request)
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
	if res.OpentelemetryPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOpentelemetryPlugin(res.OpentelemetryPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginOpentelemetryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginOpentelemetryResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	createOpentelemetryPlugin := data.ToSharedCreateOpentelemetryPlugin()
	request := operations.UpdateOpentelemetryPluginRequest{
		PluginID:                  pluginID,
		ControlPlaneID:            controlPlaneID,
		CreateOpentelemetryPlugin: createOpentelemetryPlugin,
	}
	res, err := r.client.Plugins.UpdateOpentelemetryPlugin(ctx, request)
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
	if res.OpentelemetryPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOpentelemetryPlugin(res.OpentelemetryPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginOpentelemetryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginOpentelemetryResourceModel
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

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.DeleteOpentelemetryPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteOpentelemetryPlugin(ctx, request)
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

func (r *GatewayPluginOpentelemetryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_plugin_opentelemetry. Reason: composite imports strings not supported.")
}
