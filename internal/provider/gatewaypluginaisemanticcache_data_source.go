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
var _ datasource.DataSource = &GatewayPluginAiSemanticCacheDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginAiSemanticCacheDataSource{}

func NewGatewayPluginAiSemanticCacheDataSource() datasource.DataSource {
	return &GatewayPluginAiSemanticCacheDataSource{}
}

// GatewayPluginAiSemanticCacheDataSource is the data source implementation.
type GatewayPluginAiSemanticCacheDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginAiSemanticCacheDataSourceModel describes the data model.
type GatewayPluginAiSemanticCacheDataSourceModel struct {
	Config         *tfTypes.AiSemanticCachePluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer   `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLWithoutParentsConsumer   `tfsdk:"consumer_group"`
	ControlPlaneID types.String                         `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                          `tfsdk:"created_at"`
	Enabled        types.Bool                           `tfsdk:"enabled"`
	ID             types.String                         `tfsdk:"id"`
	InstanceName   types.String                         `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering           `tfsdk:"ordering"`
	Protocols      []types.String                       `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer   `tfsdk:"route"`
	Service        *tfTypes.ACLWithoutParentsConsumer   `tfsdk:"service"`
	Tags           []types.String                       `tfsdk:"tags"`
	UpdatedAt      types.Int64                          `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginAiSemanticCacheDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_ai_semantic_cache"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginAiSemanticCacheDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginAiSemanticCache DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"cache_control": schema.BoolAttribute{
						Computed:    true,
						Description: `When enabled, respect the Cache-Control behaviors defined in RFC7234.`,
					},
					"cache_ttl": schema.Int64Attribute{
						Computed:    true,
						Description: `TTL in seconds of cache entities. Must be a value greater than 0.`,
					},
					"embeddings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"allow_override": schema.BoolAttribute{
										Computed:    true,
										Description: `If enabled, the authorization header or parameter can be overridden in the request by the value configured in the plugin.`,
									},
									"aws_access_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_ACCESS_KEY_ID environment variable for this plugin instance.`,
									},
									"aws_secret_access_key": schema.StringAttribute{
										Computed:    true,
										Description: `Set this if you are using an AWS provider (Bedrock) and you are authenticating using static IAM User credentials. Setting this will override the AWS_SECRET_ACCESS_KEY environment variable for this plugin instance.`,
									},
									"azure_client_id": schema.StringAttribute{
										Computed:    true,
										Description: `If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client ID.`,
									},
									"azure_client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the client secret.`,
									},
									"azure_tenant_id": schema.StringAttribute{
										Computed:    true,
										Description: `If azure_use_managed_identity is set to true, and you need to use a different user-assigned identity for this LLM instance, set the tenant ID.`,
									},
									"azure_use_managed_identity": schema.BoolAttribute{
										Computed:    true,
										Description: `Set true to use the Azure Cloud Managed Identity (or user-assigned identity) to authenticate with Azure-provider models.`,
									},
									"gcp_service_account_json": schema.StringAttribute{
										Computed:    true,
										Description: `Set this field to the full JSON of the GCP service account to authenticate, if required. If null (and gcp_use_service_account is true), Kong will attempt to read from environment variable ` + "`" + `GCP_SERVICE_ACCOUNT` + "`" + `.`,
									},
									"gcp_use_service_account": schema.BoolAttribute{
										Computed:    true,
										Description: `Use service account auth for GCP-based providers and models.`,
									},
									"header_name": schema.StringAttribute{
										Computed:    true,
										Description: `If AI model requires authentication via Authorization or API key header, specify its name here.`,
									},
									"header_value": schema.StringAttribute{
										Computed:    true,
										Description: `Specify the full auth header value for 'header_name', for example 'Bearer key' or just 'key'.`,
									},
									"param_location": schema.StringAttribute{
										Computed:    true,
										Description: `Specify whether the 'param_name' and 'param_value' options go in a query string, or the POST form/JSON body.`,
									},
									"param_name": schema.StringAttribute{
										Computed:    true,
										Description: `If AI model requires authentication via query parameter, specify its name here.`,
									},
									"param_value": schema.StringAttribute{
										Computed:    true,
										Description: `Specify the full parameter value for 'param_name'.`,
									},
								},
							},
							"model": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed:    true,
										Description: `Model name to execute.`,
									},
									"options": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"upstream_url": schema.StringAttribute{
												Computed:    true,
												Description: `upstream url for the embeddings`,
											},
										},
										Description: `Key/value settings for the model`,
									},
									"provider": schema.StringAttribute{
										Computed:    true,
										Description: `AI provider format to use for embeddings API`,
									},
								},
							},
						},
					},
					"exact_caching": schema.BoolAttribute{
						Computed:    true,
						Description: `When enabled, a first check for exact query will be done. It will impact DB size`,
					},
					"ignore_assistant_prompts": schema.BoolAttribute{
						Computed:    true,
						Description: `Ignore and discard any assistant prompts when Vectorizing the request`,
					},
					"ignore_system_prompts": schema.BoolAttribute{
						Computed:    true,
						Description: `Ignore and discard any system prompts when Vectorizing the request`,
					},
					"ignore_tool_prompts": schema.BoolAttribute{
						Computed:    true,
						Description: `Ignore and discard any tool prompts when Vectorizing the request`,
					},
					"message_countback": schema.Float64Attribute{
						Computed:    true,
						Description: `Number of messages in the chat history to Vectorize/Cache`,
					},
					"stop_on_failure": schema.BoolAttribute{
						Computed:    true,
						Description: `Halt the LLM request process in case of a caching system failure`,
					},
					"vectordb": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"dimensions": schema.Int64Attribute{
								Computed:    true,
								Description: `the desired dimensionality for the vectors`,
							},
							"distance_metric": schema.StringAttribute{
								Computed:    true,
								Description: `the distance metric to use for vector searches`,
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
							"strategy": schema.StringAttribute{
								Computed:    true,
								Description: `which vector database driver to use`,
							},
							"threshold": schema.Float64Attribute{
								Computed:    true,
								Description: `the default similarity threshold for accepting semantic search results (float)`,
							},
						},
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
				Description: `If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups`,
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
				Description: `A set of strings representing HTTP protocols.`,
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

func (r *GatewayPluginAiSemanticCacheDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginAiSemanticCacheDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginAiSemanticCacheDataSourceModel
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

	request := operations.GetAisemanticcachePluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetAisemanticcachePlugin(ctx, request)
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
	if !(res.AiSemanticCachePlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAiSemanticCachePlugin(ctx, res.AiSemanticCachePlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
