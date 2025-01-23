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
var _ datasource.DataSource = &GatewayPluginRateLimitingAdvancedDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginRateLimitingAdvancedDataSource{}

func NewGatewayPluginRateLimitingAdvancedDataSource() datasource.DataSource {
	return &GatewayPluginRateLimitingAdvancedDataSource{}
}

// GatewayPluginRateLimitingAdvancedDataSource is the data source implementation.
type GatewayPluginRateLimitingAdvancedDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginRateLimitingAdvancedDataSourceModel describes the data model.
type GatewayPluginRateLimitingAdvancedDataSourceModel struct {
	Config         tfTypes.RateLimitingAdvancedPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer       `tfsdk:"consumer" tfPlanOnly:"true"`
	ConsumerGroup  *tfTypes.ACLWithoutParentsConsumer       `tfsdk:"consumer_group" tfPlanOnly:"true"`
	ControlPlaneID types.String                             `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                              `tfsdk:"created_at"`
	Enabled        types.Bool                               `tfsdk:"enabled"`
	ID             types.String                             `tfsdk:"id"`
	InstanceName   types.String                             `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering               `tfsdk:"ordering"`
	Protocols      []types.String                           `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer       `tfsdk:"route" tfPlanOnly:"true"`
	Service        *tfTypes.ACLWithoutParentsConsumer       `tfsdk:"service" tfPlanOnly:"true"`
	Tags           []types.String                           `tfsdk:"tags"`
	UpdatedAt      types.Int64                              `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginRateLimitingAdvancedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_rate_limiting_advanced"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginRateLimitingAdvancedDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginRateLimitingAdvanced DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"consumer_groups": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of consumer groups allowed to override the rate limiting settings for the given Route or Service. Required if ` + "`" + `enforce_consumer_groups` + "`" + ` is set to ` + "`" + `true` + "`" + `.`,
					},
					"dictionary_name": schema.StringAttribute{
						Computed:    true,
						Description: `The shared dictionary where counters are stored. When the plugin is configured to synchronize counter data externally (that is ` + "`" + `config.strategy` + "`" + ` is ` + "`" + `cluster` + "`" + ` or ` + "`" + `redis` + "`" + ` and ` + "`" + `config.sync_rate` + "`" + ` isn't ` + "`" + `-1` + "`" + `), this dictionary serves as a buffer to populate counters in the data store on each synchronization cycle.`,
					},
					"disable_penalty": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to ` + "`" + `true` + "`" + `, this doesn't count denied requests (status = ` + "`" + `429` + "`" + `). If set to ` + "`" + `false` + "`" + `, all requests, including denied ones, are counted. This parameter only affects the ` + "`" + `sliding` + "`" + ` window_type.`,
					},
					"enforce_consumer_groups": schema.BoolAttribute{
						Computed:    true,
						Description: `Determines if consumer groups are allowed to override the rate limiting settings for the given Route or Service. Flipping ` + "`" + `enforce_consumer_groups` + "`" + ` from ` + "`" + `true` + "`" + ` to ` + "`" + `false` + "`" + ` disables the group override, but does not clear the list of consumer groups. You can then flip ` + "`" + `enforce_consumer_groups` + "`" + ` to ` + "`" + `true` + "`" + ` to re-enforce the groups.`,
					},
					"error_code": schema.NumberAttribute{
						Computed:    true,
						Description: `Set a custom error code to return when the rate limit is exceeded.`,
					},
					"error_message": schema.StringAttribute{
						Computed:    true,
						Description: `Set a custom error message to return when the rate limit is exceeded.`,
					},
					"header_name": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing an HTTP header name.`,
					},
					"hide_client_headers": schema.BoolAttribute{
						Computed:    true,
						Description: `Optionally hide informative response headers that would otherwise provide information about the current status of limits and counters.`,
					},
					"identifier": schema.StringAttribute{
						Computed:    true,
						Description: `The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be ` + "`" + `ip` + "`" + `, ` + "`" + `credential` + "`" + `, ` + "`" + `consumer` + "`" + `, ` + "`" + `service` + "`" + `, ` + "`" + `header` + "`" + `, ` + "`" + `path` + "`" + ` or ` + "`" + `consumer-group` + "`" + `.`,
					},
					"limit": schema.ListAttribute{
						Computed:    true,
						ElementType: types.NumberType,
						Description: `One or more requests-per-window limits to apply. There must be a matching number of window limits and sizes specified.`,
					},
					"namespace": schema.StringAttribute{
						Computed:    true,
						Description: `The rate limiting library namespace to use for this plugin instance. Counter data and sync configuration is isolated in each namespace. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. ` + "`" + `strategy` + "`" + `, ` + "`" + `redis` + "`" + `, ` + "`" + `sync_rate` + "`" + `, ` + "`" + `window_size` + "`" + `, ` + "`" + `dictionary_name` + "`" + `, need to be the same.`,
					},
					"path": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).`,
					},
					"redis": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"cluster_max_redirections": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum retry attempts for redirection.`,
							},
							"cluster_nodes": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Computed:    true,
											Description: `A string representing a host name, such as example.com.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Description: `An integer representing a port number between 0 and 65535, inclusive.`,
										},
									},
								},
								Description: `Cluster addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.`,
							},
							"connect_timeout": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"connection_is_proxied": schema.BoolAttribute{
								Computed:    true,
								Description: `If the connection to Redis is proxied (e.g. Envoy), set it ` + "`" + `true` + "`" + `. Set the ` + "`" + `host` + "`" + ` and ` + "`" + `port` + "`" + ` to point to the proxy address.`,
							},
							"database": schema.Int64Attribute{
								Computed:    true,
								Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy`,
							},
							"host": schema.StringAttribute{
								Computed:    true,
								Description: `A string representing a host name, such as example.com.`,
							},
							"keepalive_backlog": schema.Int64Attribute{
								Computed:    true,
								Description: `Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return ` + "`" + `nil` + "`" + `. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than ` + "`" + `keepalive_pool_size` + "`" + `. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than ` + "`" + `keepalive_pool_size` + "`" + `.`,
							},
							"keepalive_pool_size": schema.Int64Attribute{
								Computed:    true,
								Description: `The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither ` + "`" + `keepalive_pool_size` + "`" + ` nor ` + "`" + `keepalive_backlog` + "`" + ` is specified, no pool is created. If ` + "`" + `keepalive_pool_size` + "`" + ` isn't specified but ` + "`" + `keepalive_backlog` + "`" + ` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.`,
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
							},
							"port": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a port number between 0 and 65535, inclusive.`,
							},
							"read_timeout": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"send_timeout": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"sentinel_master": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.`,
							},
							"sentinel_nodes": schema.ListNestedAttribute{
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
								Description: `Sentinel node addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.`,
							},
							"sentinel_password": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.`,
							},
							"sentinel_role": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel role to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Sentinel.`,
							},
							"sentinel_username": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.`,
							},
							"server_name": schema.StringAttribute{
								Computed:    true,
								Description: `A string representing an SNI (server name indication) value for TLS.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Description: `If set to true, uses SSL to connect to Redis.`,
							},
							"ssl_verify": schema.BoolAttribute{
								Computed:    true,
								Description: `If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` in ` + "`" + `kong.conf` + "`" + ` to specify the CA (or server) certificate used by your Redis server. You may also need to configure ` + "`" + `lua_ssl_verify_depth` + "`" + ` accordingly.`,
							},
							"username": schema.StringAttribute{
								Computed:    true,
								Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
							},
						},
					},
					"retry_after_jitter_max": schema.NumberAttribute{
						Computed:    true,
						Description: `The upper bound of a jitter (random delay) in seconds to be added to the ` + "`" + `Retry-After` + "`" + ` header of denied requests (status = ` + "`" + `429` + "`" + `) in order to prevent all the clients from coming back at the same time. The lower bound of the jitter is ` + "`" + `0` + "`" + `; in this case, the ` + "`" + `Retry-After` + "`" + ` header is equal to the ` + "`" + `RateLimit-Reset` + "`" + ` header.`,
					},
					"strategy": schema.StringAttribute{
						Computed:    true,
						Description: `The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: ` + "`" + `local` + "`" + ` and ` + "`" + `cluster` + "`" + `.`,
					},
					"sync_rate": schema.NumberAttribute{
						Computed:    true,
						Description: `How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 will sync the counters in the specified number of seconds. The minimum allowed interval is 0.02 seconds (20ms).`,
					},
					"window_size": schema.ListAttribute{
						Computed:    true,
						ElementType: types.NumberType,
						Description: `One or more window sizes to apply a limit to (defined in seconds). There must be a matching number of window limits and sizes specified.`,
					},
					"window_type": schema.StringAttribute{
						Computed:    true,
						Description: `Sets the time window type to either ` + "`" + `sliding` + "`" + ` (default) or ` + "`" + `fixed` + "`" + `. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters.`,
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

func (r *GatewayPluginRateLimitingAdvancedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginRateLimitingAdvancedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginRateLimitingAdvancedDataSourceModel
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

	request := operations.GetRatelimitingadvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetRatelimitingadvancedPlugin(ctx, request)
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
	if !(res.RateLimitingAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRateLimitingAdvancedPlugin(res.RateLimitingAdvancedPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
