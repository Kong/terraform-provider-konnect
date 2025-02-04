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
var _ datasource.DataSource = &GatewayPluginStatsdAdvancedDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginStatsdAdvancedDataSource{}

func NewGatewayPluginStatsdAdvancedDataSource() datasource.DataSource {
	return &GatewayPluginStatsdAdvancedDataSource{}
}

// GatewayPluginStatsdAdvancedDataSource is the data source implementation.
type GatewayPluginStatsdAdvancedDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginStatsdAdvancedDataSourceModel describes the data model.
type GatewayPluginStatsdAdvancedDataSourceModel struct {
	Config         tfTypes.StatsdAdvancedPluginConfig `tfsdk:"config"`
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
func (r *GatewayPluginStatsdAdvancedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_statsd_advanced"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginStatsdAdvancedDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginStatsdAdvanced DataSource",

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
						Computed:    true,
						Description: `The default consumer identifier for metrics. This will take effect when a metric's consumer identifier is omitted. Allowed values are ` + "`" + `custom_id` + "`" + `, ` + "`" + `consumer_id` + "`" + `, ` + "`" + `username` + "`" + `.`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"hostname_in_prefix": schema.BoolAttribute{
						Computed:    true,
						Description: `Include the ` + "`" + `hostname` + "`" + ` in the ` + "`" + `prefix` + "`" + ` for each metric name.`,
					},
					"metrics": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"consumer_identifier": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"sample_rate": schema.NumberAttribute{
									Computed: true,
								},
								"service_identifier": schema.StringAttribute{
									Computed: true,
								},
								"stat_type": schema.StringAttribute{
									Computed: true,
								},
								"workspace_identifier": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Description: `List of Metrics to be logged.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
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
							"initial_retry_delay": schema.NumberAttribute{
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
							"max_coalescing_delay": schema.NumberAttribute{
								Computed:    true,
								Description: `Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.`,
							},
							"max_entries": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum number of entries that can be waiting on the queue.`,
							},
							"max_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Description: `Maximum time in seconds between retries, caps exponential backoff.`,
							},
							"max_retry_time": schema.NumberAttribute{
								Computed:    true,
								Description: `Time in seconds before the queue gives up calling a failed handler for a batch.`,
							},
						},
					},
					"service_identifier_default": schema.StringAttribute{
						Computed:    true,
						Description: `The default service identifier for metrics. This will take effect when a metric's service identifier is omitted. Allowed values are ` + "`" + `service_name_or_host` + "`" + `, ` + "`" + `service_id` + "`" + `, ` + "`" + `service_name` + "`" + `, ` + "`" + `service_host` + "`" + `.`,
					},
					"udp_packet_size": schema.NumberAttribute{
						Computed:    true,
						Description: `Combine UDP packet up to the size configured. If zero (0), don't combine the UDP packet. Must be a number between 0 and 65507 (inclusive).`,
					},
					"use_tcp": schema.BoolAttribute{
						Computed:    true,
						Description: `Use TCP instead of UDP.`,
					},
					"workspace_identifier_default": schema.StringAttribute{
						Computed:    true,
						Description: `The default workspace identifier for metrics. This will take effect when a metric's workspace identifier is omitted. Allowed values are ` + "`" + `workspace_id` + "`" + `, ` + "`" + `workspace_name` + "`" + `.`,
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

func (r *GatewayPluginStatsdAdvancedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginStatsdAdvancedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginStatsdAdvancedDataSourceModel
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
