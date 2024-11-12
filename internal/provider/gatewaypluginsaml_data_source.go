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
var _ datasource.DataSource = &GatewayPluginSamlDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginSamlDataSource{}

func NewGatewayPluginSamlDataSource() datasource.DataSource {
	return &GatewayPluginSamlDataSource{}
}

// GatewayPluginSamlDataSource is the data source implementation.
type GatewayPluginSamlDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginSamlDataSourceModel describes the data model.
type GatewayPluginSamlDataSourceModel struct {
	Config         tfTypes.SamlPluginConfig   `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer       `tfsdk:"consumer" tfPlanOnly:"true"`
	ConsumerGroup  *tfTypes.ACLConsumer       `tfsdk:"consumer_group" tfPlanOnly:"true"`
	ControlPlaneID types.String               `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                `tfsdk:"created_at"`
	Enabled        types.Bool                 `tfsdk:"enabled"`
	ID             types.String               `tfsdk:"id"`
	InstanceName   types.String               `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering `tfsdk:"ordering"`
	Protocols      []types.String             `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer       `tfsdk:"route" tfPlanOnly:"true"`
	Service        *tfTypes.ACLConsumer       `tfsdk:"service" tfPlanOnly:"true"`
	Tags           []types.String             `tfsdk:"tags"`
	UpdatedAt      types.Int64                `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginSamlDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_saml"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginSamlDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginSaml DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer. If not set, a Kong Consumer must exist for the SAML IdP user credentials, mapping the username format to the Kong Consumer username.`,
					},
					"assertion_consumer_path": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).`,
					},
					"idp_certificate": schema.StringAttribute{
						Computed:    true,
						Description: `The public certificate provided by the IdP. This is used to validate responses from the IdP.  Only include the contents of the certificate. Do not include the header (` + "`" + `BEGIN CERTIFICATE` + "`" + `) and footer (` + "`" + `END CERTIFICATE` + "`" + `) lines.`,
					},
					"idp_sso_url": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"issuer": schema.StringAttribute{
						Computed:    true,
						Description: `The unique identifier of the IdP application. Formatted as a URL containing information about the IdP so the SP can validate that the SAML assertions it receives are issued from the correct IdP.`,
					},
					"nameid_format": schema.StringAttribute{
						Computed:    true,
						Description: `The requested ` + "`" + `NameId` + "`" + ` format. Options available are: - ` + "`" + `Unspecified` + "`" + ` - ` + "`" + `EmailAddress` + "`" + ` - ` + "`" + `Persistent` + "`" + ` - ` + "`" + `Transient` + "`" + ``,
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
							"prefix": schema.StringAttribute{
								Computed:    true,
								Description: `The Redis session key prefix.`,
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
							"socket": schema.StringAttribute{
								Computed:    true,
								Description: `The Redis unix socket path.`,
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
					"request_digest_algorithm": schema.StringAttribute{
						Computed:    true,
						Description: `The digest algorithm for Authn requests: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA1` + "`" + ``,
					},
					"request_signature_algorithm": schema.StringAttribute{
						Computed:    true,
						Description: `The signature algorithm for signing Authn requests. Options available are: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA384` + "`" + ` - ` + "`" + `SHA512` + "`" + ``,
					},
					"request_signing_certificate": schema.StringAttribute{
						Computed:    true,
						Description: `The certificate for signing requests.`,
					},
					"request_signing_key": schema.StringAttribute{
						Computed:    true,
						Description: `The private key for signing requests.  If this parameter is set, requests sent to the IdP are signed.  The ` + "`" + `request_signing_certificate` + "`" + ` parameter must be set as well.`,
					},
					"response_digest_algorithm": schema.StringAttribute{
						Computed:    true,
						Description: `The algorithm for verifying digest in SAML responses: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA1` + "`" + ``,
					},
					"response_encryption_key": schema.StringAttribute{
						Computed:    true,
						Description: `The private encryption key required to decrypt encrypted assertions.`,
					},
					"response_signature_algorithm": schema.StringAttribute{
						Computed:    true,
						Description: `The algorithm for validating signatures in SAML responses. Options available are: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA384` + "`" + ` - ` + "`" + `SHA512` + "`" + ``,
					},
					"session_absolute_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid.`,
					},
					"session_audience": schema.StringAttribute{
						Computed:    true,
						Description: `The session audience, for example "my-application"`,
					},
					"session_cookie_domain": schema.StringAttribute{
						Computed:    true,
						Description: `The session cookie domain flag.`,
					},
					"session_cookie_http_only": schema.BoolAttribute{
						Computed:    true,
						Description: `Forbids JavaScript from accessing the cookie, for example, through the ` + "`" + `Document.cookie` + "`" + ` property.`,
					},
					"session_cookie_name": schema.StringAttribute{
						Computed:    true,
						Description: `The session cookie name.`,
					},
					"session_cookie_path": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).`,
					},
					"session_cookie_same_site": schema.StringAttribute{
						Computed:    true,
						Description: `Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks.`,
					},
					"session_cookie_secure": schema.BoolAttribute{
						Computed:    true,
						Description: `The cookie is only sent to the server when a request is made with the https:scheme (except on localhost), and therefore is more resistant to man-in-the-middle attacks.`,
					},
					"session_enforce_same_subject": schema.BoolAttribute{
						Computed:    true,
						Description: `When set to ` + "`" + `true` + "`" + `, audiences are forced to share the same subject.`,
					},
					"session_hash_storage_key": schema.BoolAttribute{
						Computed:    true,
						Description: `When set to ` + "`" + `true` + "`" + `, the storage key (session ID) is hashed for extra security. Hashing the storage key means it is impossible to decrypt data from the storage without a cookie.`,
					},
					"session_hash_subject": schema.BoolAttribute{
						Computed:    true,
						Description: `When set to ` + "`" + `true` + "`" + `, the value of subject is hashed before being stored. Only applies when ` + "`" + `session_store_metadata` + "`" + ` is enabled.`,
					},
					"session_idling_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `The session cookie idle time in seconds.`,
					},
					"session_memcached_host": schema.StringAttribute{
						Computed:    true,
						Description: `The memcached host.`,
					},
					"session_memcached_port": schema.Int64Attribute{
						Computed:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
					},
					"session_memcached_prefix": schema.StringAttribute{
						Computed:    true,
						Description: `The memcached session key prefix.`,
					},
					"session_memcached_socket": schema.StringAttribute{
						Computed:    true,
						Description: `The memcached unix socket path.`,
					},
					"session_remember": schema.BoolAttribute{
						Computed:    true,
						Description: `Enables or disables persistent sessions`,
					},
					"session_remember_absolute_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `Persistent session absolute timeout in seconds.`,
					},
					"session_remember_cookie_name": schema.StringAttribute{
						Computed:    true,
						Description: `Persistent session cookie name`,
					},
					"session_remember_rolling_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `Persistent session rolling timeout in seconds.`,
					},
					"session_request_headers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"session_response_headers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"session_rolling_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid.`,
					},
					"session_secret": schema.StringAttribute{
						Computed:    true,
						Description: `The session secret. This must be a random string of 32 characters from the base64 alphabet (letters, numbers, ` + "`" + `/` + "`" + `, ` + "`" + `_` + "`" + ` and ` + "`" + `+` + "`" + `). It is used as the secret key for encrypting session data as well as state information that is sent to the IdP in the authentication exchange.`,
					},
					"session_storage": schema.StringAttribute{
						Computed:    true,
						Description: `The session storage for session data: - ` + "`" + `cookie` + "`" + `: stores session data with the session cookie. The session cannot be invalidated or revoked without changing the session secret, but is stateless, and doesn't require a database. - ` + "`" + `memcached` + "`" + `: stores session data in memcached - ` + "`" + `redis` + "`" + `: stores session data in Redis`,
					},
					"session_store_metadata": schema.BoolAttribute{
						Computed:    true,
						Description: `Configures whether or not session metadata should be stored. This includes information about the active sessions for the ` + "`" + `specific_audience` + "`" + ` belonging to a specific subject.`,
					},
					"validate_assertion_signature": schema.BoolAttribute{
						Computed:    true,
						Description: `Enable signature validation for SAML responses.`,
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

func (r *GatewayPluginSamlDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginSamlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginSamlDataSourceModel
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

	request := operations.GetSamlPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetSamlPlugin(ctx, request)
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
	if !(res.SamlPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSamlPlugin(res.SamlPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
