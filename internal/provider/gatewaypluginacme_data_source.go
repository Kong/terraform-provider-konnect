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
var _ datasource.DataSource = &GatewayPluginAcmeDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginAcmeDataSource{}

func NewGatewayPluginAcmeDataSource() datasource.DataSource {
	return &GatewayPluginAcmeDataSource{}
}

// GatewayPluginAcmeDataSource is the data source implementation.
type GatewayPluginAcmeDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginAcmeDataSourceModel describes the data model.
type GatewayPluginAcmeDataSourceModel struct {
	Config         *tfTypes.AcmePluginConfig  `tfsdk:"config"`
	ControlPlaneID types.String               `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                `tfsdk:"created_at"`
	Enabled        types.Bool                 `tfsdk:"enabled"`
	ID             types.String               `tfsdk:"id"`
	InstanceName   types.String               `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering `tfsdk:"ordering"`
	Protocols      []types.String             `tfsdk:"protocols"`
	Tags           []types.String             `tfsdk:"tags"`
	UpdatedAt      types.Int64                `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginAcmeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_acme"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginAcmeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginAcme DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"account_email": schema.StringAttribute{
						Computed:    true,
						Description: `The account identifier. Can be reused in a different plugin instance.`,
					},
					"account_key": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"key_id": schema.StringAttribute{
								Computed:    true,
								Description: `The Key ID.`,
							},
							"key_set": schema.StringAttribute{
								Computed:    true,
								Description: `The ID of the key set to associate the Key ID with.`,
							},
						},
						Description: `The private key associated with the account.`,
					},
					"allow_any_domain": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to ` + "`" + `true` + "`" + `, the plugin allows all domains and ignores any values in the ` + "`" + `domains` + "`" + ` list.`,
					},
					"api_uri": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"cert_type": schema.StringAttribute{
						Computed:    true,
						Description: `The certificate type to create. The possible values are ` + "`" + `rsa` + "`" + ` for RSA certificate or ` + "`" + `ecc` + "`" + ` for EC certificate.`,
					},
					"domains": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `An array of strings representing hosts. A valid host is a string containing one or more labels separated by periods, with at most one wildcard label ('*')`,
					},
					"eab_hmac_key": schema.StringAttribute{
						Computed:    true,
						Description: `External account binding (EAB) base64-encoded URL string of the HMAC key. You usually don't need to set this unless it is explicitly required by the CA.`,
					},
					"eab_kid": schema.StringAttribute{
						Computed:    true,
						Description: `External account binding (EAB) key id. You usually don't need to set this unless it is explicitly required by the CA.`,
					},
					"enable_ipv4_common_name": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that controls whether to include the IPv4 address in the common name field of generated certificates.`,
					},
					"fail_backoff_minutes": schema.NumberAttribute{
						Computed: true,
						MarkdownDescription: `Minutes to wait for each domain that fails to create a certificate. This applies to both a` + "\n" +
							`new certificate and a renewal certificate.`,
					},
					"preferred_chain": schema.StringAttribute{
						Computed:    true,
						Description: `A string value that specifies the preferred certificate chain to use when generating certificates.`,
					},
					"renew_threshold_days": schema.NumberAttribute{
						Computed:    true,
						Description: `Days remaining to renew the certificate before it expires.`,
					},
					"rsa_key_size": schema.Int64Attribute{
						Computed:    true,
						Description: `RSA private key size for the certificate. The possible values are 2048, 3072, or 4096.`,
					},
					"storage": schema.StringAttribute{
						Computed:    true,
						Description: `The backend storage type to use. In DB-less mode and Konnect, ` + "`" + `kong` + "`" + ` storage is unavailable. In hybrid mode and Konnect, ` + "`" + `shm` + "`" + ` storage is unavailable. ` + "`" + `shm` + "`" + ` storage does not persist during Kong restarts and does not work for Kong running on different machines, so consider using one of ` + "`" + `kong` + "`" + `, ` + "`" + `redis` + "`" + `, ` + "`" + `consul` + "`" + `, or ` + "`" + `vault` + "`" + ` in production.`,
					},
					"storage_config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"consul": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Computed:    true,
										Description: `A string representing a host name, such as example.com.`,
									},
									"https": schema.BoolAttribute{
										Computed:    true,
										Description: `Boolean representation of https.`,
									},
									"kv_path": schema.StringAttribute{
										Computed:    true,
										Description: `KV prefix path.`,
									},
									"port": schema.Int64Attribute{
										Computed:    true,
										Description: `An integer representing a port number between 0 and 65535, inclusive.`,
									},
									"timeout": schema.NumberAttribute{
										Computed:    true,
										Description: `Timeout in milliseconds.`,
									},
									"token": schema.StringAttribute{
										Computed:    true,
										Description: `Consul ACL token.`,
									},
								},
							},
							"kong": schema.MapAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"redis": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"database": schema.Int64Attribute{
										Computed:    true,
										Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy`,
									},
									"extra_options": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"namespace": schema.StringAttribute{
												Computed:    true,
												Description: `A namespace to prepend to all keys stored in Redis.`,
											},
											"scan_count": schema.NumberAttribute{
												Computed:    true,
												Description: `The number of keys to return in Redis SCAN calls.`,
											},
										},
										Description: `Custom ACME Redis options`,
									},
									"host": schema.StringAttribute{
										Computed:    true,
										Description: `A string representing a host name, such as example.com.`,
									},
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
									},
									"port": schema.Int64Attribute{
										Computed:    true,
										Description: `An integer representing a port number between 0 and 65535, inclusive.`,
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
									"timeout": schema.Int64Attribute{
										Computed:    true,
										Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
									},
									"username": schema.StringAttribute{
										Computed:    true,
										Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
									},
								},
							},
							"shm": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"shm_name": schema.StringAttribute{
										Computed:    true,
										Description: `Name of shared memory zone used for Kong API gateway storage`,
									},
								},
							},
							"vault": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_method": schema.StringAttribute{
										Computed:    true,
										Description: `Auth Method, default to token, can be 'token' or 'kubernetes'.`,
									},
									"auth_path": schema.StringAttribute{
										Computed:    true,
										Description: `Vault's authentication path to use.`,
									},
									"auth_role": schema.StringAttribute{
										Computed:    true,
										Description: `The role to try and assign.`,
									},
									"host": schema.StringAttribute{
										Computed:    true,
										Description: `A string representing a host name, such as example.com.`,
									},
									"https": schema.BoolAttribute{
										Computed:    true,
										Description: `Boolean representation of https.`,
									},
									"jwt_path": schema.StringAttribute{
										Computed:    true,
										Description: `The path to the JWT.`,
									},
									"kv_path": schema.StringAttribute{
										Computed:    true,
										Description: `KV prefix path.`,
									},
									"port": schema.Int64Attribute{
										Computed:    true,
										Description: `An integer representing a port number between 0 and 65535, inclusive.`,
									},
									"timeout": schema.NumberAttribute{
										Computed:    true,
										Description: `Timeout in milliseconds.`,
									},
									"tls_server_name": schema.StringAttribute{
										Computed:    true,
										Description: `SNI used in request, default to host if omitted.`,
									},
									"tls_verify": schema.BoolAttribute{
										Computed:    true,
										Description: `Turn on TLS verification.`,
									},
									"token": schema.StringAttribute{
										Computed:    true,
										Description: `Consul ACL token.`,
									},
								},
							},
						},
					},
					"tos_accepted": schema.BoolAttribute{
						Computed:    true,
						Description: `If you are using Let's Encrypt, you must set this to ` + "`" + `true` + "`" + ` to agree the terms of service.`,
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
				Description: `A set of strings representing HTTP protocols.`,
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

func (r *GatewayPluginAcmeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginAcmeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginAcmeDataSourceModel
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

	request := operations.GetAcmePluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetAcmePlugin(ctx, request)
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
	if !(res.AcmePlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAcmePlugin(res.AcmePlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
