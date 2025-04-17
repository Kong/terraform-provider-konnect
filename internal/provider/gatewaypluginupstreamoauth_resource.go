// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
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
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginUpstreamOauthResource{}
var _ resource.ResourceWithImportState = &GatewayPluginUpstreamOauthResource{}

func NewGatewayPluginUpstreamOauthResource() resource.Resource {
	return &GatewayPluginUpstreamOauthResource{}
}

// GatewayPluginUpstreamOauthResource defines the resource implementation.
type GatewayPluginUpstreamOauthResource struct {
	client *sdk.Konnect
}

// GatewayPluginUpstreamOauthResourceModel describes the resource data model.
type GatewayPluginUpstreamOauthResourceModel struct {
	Config         *tfTypes.UpstreamOauthPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer_group"`
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

func (r *GatewayPluginUpstreamOauthResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_upstream_oauth"
}

func (r *GatewayPluginUpstreamOauthResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginUpstreamOauth Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"behavior": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"idp_error_response_body_template": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The template to use to create the body of the response to return to the consumer if Kong fails to obtain a token from the IdP.`,
							},
							"idp_error_response_content_type": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The Content-Type of the response to return to the consumer if Kong fails to obtain a token from the IdP.`,
							},
							"idp_error_response_message": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The message to embed in the body of the response to return to the consumer if Kong fails to obtain a token from the IdP.`,
							},
							"idp_error_response_status_code": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The response code to return to the consumer if Kong fails to obtain a token from the IdP.`,
								Validators: []validator.Int64{
									int64validator.Between(500, 599),
								},
							},
							"purge_token_on_upstream_status_codes": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.Int64Type,
								Description: `An array of status codes which will force an access token to be purged when returned by the upstream. An empty array will disable this functionality.`,
							},
							"upstream_access_token_header_name": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The name of the header used to send the access token (obtained from the IdP) to the upstream service.`,
							},
						},
					},
					"cache": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"default_ttl": schema.Float64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The lifetime of a token without an explicit ` + "`" + `expires_in` + "`" + ` value.`,
							},
							"eagerly_expire": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The number of seconds to eagerly expire a cached token. By default, a cached token expires 5 seconds before its lifetime as defined in ` + "`" + `expires_in` + "`" + `.`,
							},
							"memory": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"dictionary_name": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `The shared dictionary used by the plugin to cache tokens if ` + "`" + `config.cache.strategy` + "`" + ` is set to ` + "`" + `memory` + "`" + `.`,
									},
								},
							},
							"redis": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"cluster_max_redirections": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `Maximum retry attempts for redirection.`,
									},
									"cluster_nodes": schema.ListNestedAttribute{
										Computed: true,
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Validators: []validator.Object{
												speakeasy_objectvalidators.NotNull(),
											},
											Attributes: map[string]schema.Attribute{
												"ip": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Description: `A string representing a host name, such as example.com.`,
												},
												"port": schema.Int64Attribute{
													Computed:    true,
													Optional:    true,
													Description: `An integer representing a port number between 0 and 65535, inclusive.`,
													Validators: []validator.Int64{
														int64validator.AtMost(65535),
													},
												},
											},
										},
										Description: `Cluster addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.`,
									},
									"connect_timeout": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
										Validators: []validator.Int64{
											int64validator.AtMost(2147483646),
										},
									},
									"connection_is_proxied": schema.BoolAttribute{
										Computed:    true,
										Optional:    true,
										Description: `If the connection to Redis is proxied (e.g. Envoy), set it ` + "`" + `true` + "`" + `. Set the ` + "`" + `host` + "`" + ` and ` + "`" + `port` + "`" + ` to point to the proxy address.`,
									},
									"database": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy`,
									},
									"host": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `A string representing a host name, such as example.com.`,
									},
									"keepalive_backlog": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return ` + "`" + `nil` + "`" + `. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than ` + "`" + `keepalive_pool_size` + "`" + `. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than ` + "`" + `keepalive_pool_size` + "`" + `.`,
										Validators: []validator.Int64{
											int64validator.AtMost(2147483646),
										},
									},
									"keepalive_pool_size": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither ` + "`" + `keepalive_pool_size` + "`" + ` nor ` + "`" + `keepalive_backlog` + "`" + ` is specified, no pool is created. If ` + "`" + `keepalive_pool_size` + "`" + ` isn't specified but ` + "`" + `keepalive_backlog` + "`" + ` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.`,
										Validators: []validator.Int64{
											int64validator.Between(1, 2147483646),
										},
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
									},
									"port": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `An integer representing a port number between 0 and 65535, inclusive.`,
										Validators: []validator.Int64{
											int64validator.AtMost(65535),
										},
									},
									"read_timeout": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
										Validators: []validator.Int64{
											int64validator.AtMost(2147483646),
										},
									},
									"send_timeout": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
										Validators: []validator.Int64{
											int64validator.AtMost(2147483646),
										},
									},
									"sentinel_master": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.`,
									},
									"sentinel_nodes": schema.ListNestedAttribute{
										Computed: true,
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Validators: []validator.Object{
												speakeasy_objectvalidators.NotNull(),
											},
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Description: `A string representing a host name, such as example.com.`,
												},
												"port": schema.Int64Attribute{
													Computed:    true,
													Optional:    true,
													Description: `An integer representing a port number between 0 and 65535, inclusive.`,
													Validators: []validator.Int64{
														int64validator.AtMost(65535),
													},
												},
											},
										},
										Description: `Sentinel node addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.`,
									},
									"sentinel_password": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.`,
									},
									"sentinel_role": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Sentinel role to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Sentinel. must be one of ["any", "master", "slave"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"any",
												"master",
												"slave",
											),
										},
									},
									"sentinel_username": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.`,
									},
									"server_name": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `A string representing an SNI (server name indication) value for TLS.`,
									},
									"ssl": schema.BoolAttribute{
										Computed:    true,
										Optional:    true,
										Description: `If set to true, uses SSL to connect to Redis.`,
									},
									"ssl_verify": schema.BoolAttribute{
										Computed:    true,
										Optional:    true,
										Description: `If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` in ` + "`" + `kong.conf` + "`" + ` to specify the CA (or server) certificate used by your Redis server. You may also need to configure ` + "`" + `lua_ssl_verify_depth` + "`" + ` accordingly.`,
									},
									"username": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
									},
								},
							},
							"strategy": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The method Kong should use to cache tokens issued by the IdP. must be one of ["memory", "redis"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"memory",
										"redis",
									),
								},
							},
						},
					},
					"client": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"auth_method": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The authentication method used in client requests to the IdP. Supported values are: ` + "`" + `client_secret_basic` + "`" + ` to send ` + "`" + `client_id` + "`" + ` and ` + "`" + `client_secret` + "`" + ` in the ` + "`" + `Authorization: Basic` + "`" + ` header, ` + "`" + `client_secret_post` + "`" + ` to send ` + "`" + `client_id` + "`" + ` and ` + "`" + `client_secret` + "`" + ` as part of the request body, or ` + "`" + `client_secret_jwt` + "`" + ` to send a JWT signed with the ` + "`" + `client_secret` + "`" + ` using the client assertion as part of the body. must be one of ["client_secret_basic", "client_secret_jwt", "client_secret_post", "none"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"client_secret_basic",
										"client_secret_jwt",
										"client_secret_post",
										"none",
									),
								},
							},
							"client_secret_jwt_alg": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The algorithm to use with JWT when using ` + "`" + `client_secret_jwt` + "`" + ` authentication. must be one of ["HS256", "HS512"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"HS256",
										"HS512",
									),
								},
							},
							"http_proxy": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The proxy to use when making HTTP requests to the IdP.`,
							},
							"http_proxy_authorization": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The ` + "`" + `Proxy-Authorization` + "`" + ` header value to be used with ` + "`" + `http_proxy` + "`" + `.`,
							},
							"http_version": schema.Float64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The HTTP version used for requests made by this plugin. Supported values: ` + "`" + `1.1` + "`" + ` for HTTP 1.1 and ` + "`" + `1.0` + "`" + ` for HTTP 1.0.`,
							},
							"https_proxy": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The proxy to use when making HTTPS requests to the IdP.`,
							},
							"https_proxy_authorization": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The ` + "`" + `Proxy-Authorization` + "`" + ` header value to be used with ` + "`" + `https_proxy` + "`" + `.`,
							},
							"keep_alive": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Whether to use keepalive connections to the IdP.`,
							},
							"no_proxy": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A comma-separated list of hosts that should not be proxied.`,
							},
							"ssl_verify": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Whether to verify the certificate presented by the IdP when using HTTPS.`,
							},
							"timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Network I/O timeout for requests to the IdP in milliseconds.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
						},
					},
					"oauth": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"audience": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `List of audiences passed to the IdP when obtaining a new token.`,
							},
							"client_id": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The client ID for the application registration in the IdP.`,
							},
							"client_secret": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The client secret for the application registration in the IdP.`,
							},
							"grant_type": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The OAuth grant type to be used. must be one of ["client_credentials", "password"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"client_credentials",
										"password",
									),
								},
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The password to use if ` + "`" + `config.oauth.grant_type` + "`" + ` is set to ` + "`" + `password` + "`" + `.`,
							},
							"scopes": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								Default:     listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
								ElementType: types.StringType,
								Description: `List of scopes to request from the IdP when obtaining a new token.`,
							},
							"token_endpoint": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The token endpoint URI.`,
							},
							"token_headers": schema.MapAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `Extra headers to be passed in the token endpoint request.`,
								Validators: []validator.Map{
									mapvalidator.ValueStringsAre(validators.IsValidJSON()),
								},
							},
							"token_post_args": schema.MapAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `Extra post arguments to be passed in the token endpoint request.`,
								Validators: []validator.Map{
									mapvalidator.ValueStringsAre(validators.IsValidJSON()),
								},
							},
							"username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `The username to use if ` + "`" + `config.oauth.grant_type` + "`" + ` is set to ` + "`" + `password` + "`" + `.`,
							},
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
			"consumer_group": schema.SingleNestedAttribute{
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
				Description: `If set, the plugin will activate only for requests where the specified consumer group has been authenticated. (Note that some plugins can not be restricted to consumers groups this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer Groups`,
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
				Optional:    true,
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
				Description: `A set of strings representing HTTP protocols.`,
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
				Optional:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayPluginUpstreamOauthResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginUpstreamOauthResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginUpstreamOauthResourceModel
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

	upstreamOauthPlugin := *data.ToSharedUpstreamOauthPlugin()
	request := operations.CreateUpstreamoauthPluginRequest{
		ControlPlaneID:      controlPlaneID,
		UpstreamOauthPlugin: upstreamOauthPlugin,
	}
	res, err := r.client.Plugins.CreateUpstreamoauthPlugin(ctx, request)
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
	if !(res.UpstreamOauthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedUpstreamOauthPlugin(ctx, res.UpstreamOauthPlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginUpstreamOauthResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginUpstreamOauthResourceModel
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

	request := operations.GetUpstreamoauthPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetUpstreamoauthPlugin(ctx, request)
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
	if !(res.UpstreamOauthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedUpstreamOauthPlugin(ctx, res.UpstreamOauthPlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginUpstreamOauthResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginUpstreamOauthResourceModel
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

	upstreamOauthPlugin := *data.ToSharedUpstreamOauthPlugin()
	request := operations.UpdateUpstreamoauthPluginRequest{
		PluginID:            pluginID,
		ControlPlaneID:      controlPlaneID,
		UpstreamOauthPlugin: upstreamOauthPlugin,
	}
	res, err := r.client.Plugins.UpdateUpstreamoauthPlugin(ctx, request)
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
	if !(res.UpstreamOauthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedUpstreamOauthPlugin(ctx, res.UpstreamOauthPlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginUpstreamOauthResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginUpstreamOauthResourceModel
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

	request := operations.DeleteUpstreamoauthPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteUpstreamoauthPlugin(ctx, request)
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

func (r *GatewayPluginUpstreamOauthResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The import ID is not valid. It is expected to be a JSON object string with the format: '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "id": "3473c251-5b6c-4f45-b1ff-7ede735a366d"}': `+err.Error())
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
