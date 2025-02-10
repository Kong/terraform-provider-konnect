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
var _ datasource.DataSource = &GatewayPluginLdapAuthDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginLdapAuthDataSource{}

func NewGatewayPluginLdapAuthDataSource() datasource.DataSource {
	return &GatewayPluginLdapAuthDataSource{}
}

// GatewayPluginLdapAuthDataSource is the data source implementation.
type GatewayPluginLdapAuthDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginLdapAuthDataSourceModel describes the data model.
type GatewayPluginLdapAuthDataSourceModel struct {
	Config         tfTypes.LdapAuthPluginConfig       `tfsdk:"config"`
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
func (r *GatewayPluginLdapAuthDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_ldap_auth"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginLdapAuthDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginLdapAuth DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure ` + "`" + `4xx` + "`" + `.`,
					},
					"attribute": schema.StringAttribute{
						Computed:    true,
						Description: `Attribute to be used to search the user; e.g. cn`,
					},
					"base_dn": schema.StringAttribute{
						Computed:    true,
						Description: `Base DN as the starting point for the search; e.g., dc=example,dc=com`,
					},
					"cache_ttl": schema.NumberAttribute{
						Computed:    true,
						Description: `Cache expiry time in seconds.`,
					},
					"header_type": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string to use as part of the Authorization header`,
					},
					"hide_credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value telling the plugin to hide the credential to the upstream server. It will be removed by Kong before proxying the request.`,
					},
					"keepalive": schema.NumberAttribute{
						Computed:    true,
						Description: `An optional value in milliseconds that defines how long an idle connection to LDAP server will live before being closed.`,
					},
					"ldap_host": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"ldap_port": schema.Int64Attribute{
						Computed:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
					},
					"ldaps": schema.BoolAttribute{
						Computed:    true,
						Description: `Set to ` + "`" + `true` + "`" + ` to connect using the LDAPS protocol (LDAP over TLS).  When ` + "`" + `ldaps` + "`" + ` is configured, you must use port 636. If the ` + "`" + `ldap` + "`" + ` setting is enabled, ensure the ` + "`" + `start_tls` + "`" + ` setting is disabled.`,
					},
					"realm": schema.StringAttribute{
						Computed:    true,
						Description: `When authentication fails the plugin sends ` + "`" + `WWW-Authenticate` + "`" + ` header with ` + "`" + `realm` + "`" + ` attribute value.`,
					},
					"start_tls": schema.BoolAttribute{
						Computed:    true,
						Description: `Set it to ` + "`" + `true` + "`" + ` to issue StartTLS (Transport Layer Security) extended operation over ` + "`" + `ldap` + "`" + ` connection. If the ` + "`" + `start_tls` + "`" + ` setting is enabled, ensure the ` + "`" + `ldaps` + "`" + ` setting is disabled.`,
					},
					"timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `An optional timeout in milliseconds when waiting for connection with LDAP server.`,
					},
					"verify_ldap_host": schema.BoolAttribute{
						Computed:    true,
						Description: `Set to ` + "`" + `true` + "`" + ` to authenticate LDAP server. The server certificate will be verified according to the CA certificates specified by the ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` directive.`,
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
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support tcp and tls.`,
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

func (r *GatewayPluginLdapAuthDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginLdapAuthDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginLdapAuthDataSourceModel
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

	request := operations.GetLdapauthPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetLdapauthPlugin(ctx, request)
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
	if !(res.LdapAuthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLdapAuthPlugin(res.LdapAuthPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
