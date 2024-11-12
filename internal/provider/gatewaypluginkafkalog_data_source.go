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
var _ datasource.DataSource = &GatewayPluginKafkaLogDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginKafkaLogDataSource{}

func NewGatewayPluginKafkaLogDataSource() datasource.DataSource {
	return &GatewayPluginKafkaLogDataSource{}
}

// GatewayPluginKafkaLogDataSource is the data source implementation.
type GatewayPluginKafkaLogDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginKafkaLogDataSourceModel describes the data model.
type GatewayPluginKafkaLogDataSourceModel struct {
	Config         tfTypes.KafkaLogPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer         `tfsdk:"consumer" tfPlanOnly:"true"`
	ConsumerGroup  *tfTypes.ACLConsumer         `tfsdk:"consumer_group" tfPlanOnly:"true"`
	ControlPlaneID types.String                 `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                  `tfsdk:"created_at"`
	Enabled        types.Bool                   `tfsdk:"enabled"`
	ID             types.String                 `tfsdk:"id"`
	InstanceName   types.String                 `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering   `tfsdk:"ordering"`
	Protocols      []types.String               `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer         `tfsdk:"route" tfPlanOnly:"true"`
	Service        *tfTypes.ACLConsumer         `tfsdk:"service" tfPlanOnly:"true"`
	Tags           []types.String               `tfsdk:"tags"`
	UpdatedAt      types.Int64                  `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginKafkaLogDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_kafka_log"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginKafkaLogDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginKafkaLog DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"authentication": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"mechanism": schema.StringAttribute{
								Computed:    true,
								Description: `The SASL authentication mechanism.  Supported options: ` + "`" + `PLAIN` + "`" + ` or ` + "`" + `SCRAM-SHA-256` + "`" + `.`,
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Description: `Password for SASL authentication.`,
							},
							"strategy": schema.StringAttribute{
								Computed:    true,
								Description: `The authentication strategy for the plugin, the only option for the value is ` + "`" + `sasl` + "`" + `.`,
							},
							"tokenauth": schema.BoolAttribute{
								Computed:    true,
								Description: `Enable this to indicate ` + "`" + `DelegationToken` + "`" + ` authentication`,
							},
							"user": schema.StringAttribute{
								Computed:    true,
								Description: `Username for SASL authentication.`,
							},
						},
					},
					"bootstrap_servers": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Computed:    true,
									Description: `A string representing a host name, such as example.com.`,
								},
								"port": schema.Int64Attribute{
									Computed:    true,
									Description: `An integer representing a port number between 0 and 65535, inclusive.`,
								},
							},
						},
						Description: `Set of bootstrap brokers in a ` + "`" + `{host: host, port: port}` + "`" + ` list format.`,
					},
					"cluster_name": schema.StringAttribute{
						Computed:    true,
						Description: `An identifier for the Kafka cluster. By default, this field generates a random string. You can also set your own custom cluster identifier.  If more than one Kafka plugin is configured without a ` + "`" + `cluster_name` + "`" + ` (that is, if the default autogenerated value is removed), these plugins will use the same producer, and by extension, the same cluster. Logs will be sent to the leader of the cluster.`,
					},
					"custom_fields_by_lua": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Lua code as a key-value map`,
					},
					"keepalive": schema.Int64Attribute{
						Computed: true,
					},
					"keepalive_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"producer_async": schema.BoolAttribute{
						Computed:    true,
						Description: `Flag to enable asynchronous mode.`,
					},
					"producer_async_buffering_limits_messages_in_memory": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of messages that can be buffered in memory in asynchronous mode.`,
					},
					"producer_async_flush_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum time interval in milliseconds between buffer flushes in asynchronous mode.`,
					},
					"producer_request_acks": schema.Int64Attribute{
						Computed:    true,
						Description: `The number of acknowledgments the producer requires the leader to have received before considering a request complete. Allowed values: 0 for no acknowledgments; 1 for only the leader; and -1 for the full ISR (In-Sync Replica set).`,
					},
					"producer_request_limits_bytes_per_request": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum size of a Produce request in bytes.`,
					},
					"producer_request_limits_messages_per_request": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of messages to include into a single Produce request.`,
					},
					"producer_request_retries_backoff_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Backoff interval between retry attempts in milliseconds.`,
					},
					"producer_request_retries_max_attempts": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of retry attempts per single Produce request.`,
					},
					"producer_request_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Time to wait for a Produce response in milliseconds`,
					},
					"security": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"certificate_id": schema.StringAttribute{
								Computed:    true,
								Description: `UUID of certificate entity for mTLS authentication.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Description: `Enables TLS.`,
							},
						},
					},
					"timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Socket timeout in milliseconds.`,
					},
					"topic": schema.StringAttribute{
						Computed:    true,
						Description: `The Kafka topic to publish to.`,
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
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
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
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
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

func (r *GatewayPluginKafkaLogDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginKafkaLogDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginKafkaLogDataSourceModel
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

	request := operations.GetKafkalogPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetKafkalogPlugin(ctx, request)
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
	if !(res.KafkaLogPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedKafkaLogPlugin(res.KafkaLogPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
