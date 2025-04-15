// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayPluginStatsdDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginStatsdDataSource{}

func NewGatewayPluginStatsdDataSource() datasource.DataSource {
	return &GatewayPluginStatsdDataSource{}
}

// GatewayPluginStatsdDataSource is the data source implementation.
type GatewayPluginStatsdDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginStatsdDataSourceModel describes the data model.
type GatewayPluginStatsdDataSourceModel struct {
	Config         *tfTypes.StatsdPluginConfig        `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer"`
	ControlPlaneID types.String                       `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                        `tfsdk:"created_at"`
	Enabled        types.Bool                         `tfsdk:"enabled"`
	ID             types.String                       `tfsdk:"id"`
	InstanceName   types.String                       `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering         `tfsdk:"ordering"`
	Protocols      []types.String                     `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer `tfsdk:"route"`
	Service        *tfTypes.ACLWithoutParentsConsumer `tfsdk:"service"`
	Tags           []types.String                     `tfsdk:"tags"`
	UpdatedAt      types.Int64                        `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginStatsdDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_statsd"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginStatsdDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginStatsd DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allow_status_codes": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of status code ranges that are allowed to be logged in metrics.`,
					},
					"consumer_identifier_default": schema.StringAttribute{
						Computed: true,
					},
					"flush_timeout": schema.Float64Attribute{
						Computed: true,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `The IP address or hostname of StatsD server to send data to.`,
					},
					"hostname_in_prefix": schema.BoolAttribute{
						Computed: true,
					},
					"metrics": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"consumer_identifier": schema.StringAttribute{
									Computed:    true,
									Description: `Authenticated user detail.`,
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `StatsD metric’s name.`,
								},
								"sample_rate": schema.Float64Attribute{
									Computed:    true,
									Description: `Sampling rate`,
								},
								"service_identifier": schema.StringAttribute{
									Computed:    true,
									Description: `Service detail.`,
								},
								"stat_type": schema.StringAttribute{
									Computed:    true,
									Description: `Determines what sort of event a metric represents.`,
								},
								"workspace_identifier": schema.StringAttribute{
									Computed:    true,
									Description: `Workspace detail.`,
								},
							},
						},
						Description: `List of metrics to be logged.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Description: `The port of StatsD server to send data to.`,
					},
					"prefix": schema.StringAttribute{
						Computed:    true,
						Description: `String to prefix to each metric's name.`,
					},
					"queue": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"concurrency_limit": schema.Int64Attribute{
								Computed:    true,
								Description: `The number of of queue delivery timers. -1 indicates unlimited.`,
							},
							"initial_retry_delay": schema.Float64Attribute{
								Computed:    true,
								Description: `Time in seconds before the initial retry is made for a failing batch.`,
							},
							"max_batch_size": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum number of entries that can be processed at a time.`,
							},
							"max_bytes": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum number of bytes that can be waiting on a queue, requires string content.`,
							},
							"max_coalescing_delay": schema.Float64Attribute{
								Computed:    true,
								Description: `Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.`,
							},
							"max_entries": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum number of entries that can be waiting on the queue.`,
							},
							"max_retry_delay": schema.Float64Attribute{
								Computed:    true,
								Description: `Maximum time in seconds between retries, caps exponential backoff.`,
							},
							"max_retry_time": schema.Float64Attribute{
								Computed:    true,
								Description: `Time in seconds before the queue gives up calling a failed handler for a batch.`,
							},
						},
					},
					"queue_size": schema.Int64Attribute{
						Computed: true,
					},
					"retry_count": schema.Int64Attribute{
						Computed: true,
					},
					"service_identifier_default": schema.StringAttribute{
						Computed: true,
					},
					"tag_style": schema.StringAttribute{
						Computed: true,
					},
					"udp_packet_size": schema.Float64Attribute{
						Computed: true,
					},
					"use_tcp": schema.BoolAttribute{
						Computed: true,
					},
					"workspace_identifier_default": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
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
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A set of strings representing protocols.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
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

func (r *GatewayPluginStatsdDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayPluginStatsdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginStatsdDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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

	request := operations.GetStatsdPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetStatsdPlugin(ctx, request)
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
	if !(res.StatsdPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedStatsdPlugin(ctx, res.StatsdPlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
