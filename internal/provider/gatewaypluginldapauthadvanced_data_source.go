// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayPluginLdapAuthAdvancedDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginLdapAuthAdvancedDataSource{}

func NewGatewayPluginLdapAuthAdvancedDataSource() datasource.DataSource {
	return &GatewayPluginLdapAuthAdvancedDataSource{}
}

// GatewayPluginLdapAuthAdvancedDataSource is the data source implementation.
type GatewayPluginLdapAuthAdvancedDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginLdapAuthAdvancedDataSourceModel describes the data model.
type GatewayPluginLdapAuthAdvancedDataSourceModel struct {
	Config         *tfTypes.CreateLdapAuthAdvancedPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                        `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer                        `tfsdk:"consumer_group"`
	ControlPlaneID types.String                                `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                                 `tfsdk:"created_at"`
	Enabled        types.Bool                                  `tfsdk:"enabled"`
	ID             types.String                                `tfsdk:"id"`
	InstanceName   types.String                                `tfsdk:"instance_name"`
	Ordering       *tfTypes.CreateACLPluginOrdering            `tfsdk:"ordering"`
	Protocols      []types.String                              `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                        `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                        `tfsdk:"service"`
	Tags           []types.String                              `tfsdk:"tags"`
	UpdatedAt      types.Int64                                 `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginLdapAuthAdvancedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_ldap_auth_advanced"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginLdapAuthAdvancedDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginLdapAuthAdvanced DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request will fail with an authentication failure ` + "`" + `4xx` + "`" + `. Note that this value must refer to the consumer ` + "`" + `id` + "`" + ` or ` + "`" + `username` + "`" + ` attribute, and **not** its ` + "`" + `custom_id` + "`" + `.`,
					},
					"attribute": schema.StringAttribute{
						Computed:    true,
						Description: `Attribute to be used to search the user; e.g., "cn".`,
					},
					"base_dn": schema.StringAttribute{
						Computed:    true,
						Description: `Base DN as the starting point for the search; e.g., 'dc=example,dc=com'.`,
					},
					"bind_dn": schema.StringAttribute{
						Computed:    true,
						Description: `The DN to bind to. Used to perform LDAP search of user. This ` + "`" + `bind_dn` + "`" + ` should have permissions to search for the user being authenticated.`,
					},
					"cache_ttl": schema.NumberAttribute{
						Computed:    true,
						Description: `Cache expiry time in seconds.`,
					},
					"consumer_by": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Whether to authenticate consumers based on ` + "`" + `username` + "`" + `, ` + "`" + `custom_id` + "`" + `, or both.`,
					},
					"consumer_optional": schema.BoolAttribute{
						Computed:    true,
						Description: `Whether consumer mapping is optional. If ` + "`" + `consumer_optional=true` + "`" + `, the plugin will not attempt to associate a consumer with the LDAP authenticated user.`,
					},
					"group_base_dn": schema.StringAttribute{
						Computed:    true,
						Description: `Sets a distinguished name (DN) for the entry where LDAP searches for groups begin. This field is case-insensitive.',dc=com'.`,
					},
					"group_member_attribute": schema.StringAttribute{
						Computed:    true,
						Description: `Sets the attribute holding the members of the LDAP group. This field is case-sensitive.`,
					},
					"group_name_attribute": schema.StringAttribute{
						Computed:    true,
						Description: `Sets the attribute holding the name of a group, typically called ` + "`" + `name` + "`" + ` (in Active Directory) or ` + "`" + `cn` + "`" + ` (in OpenLDAP). This field is case-insensitive.`,
					},
					"groups_required": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `The groups required to be present in the LDAP search result for successful authorization. This config parameter works in both **AND** / **OR** cases. - When ` + "`" + `["group1 group2"]` + "`" + ` are in the same array indices, both ` + "`" + `group1` + "`" + ` AND ` + "`" + `group2` + "`" + ` need to be present in the LDAP search result. - When ` + "`" + `["group1", "group2"]` + "`" + ` are in different array indices, either ` + "`" + `group1` + "`" + ` OR ` + "`" + `group2` + "`" + ` need to be present in the LDAP search result.`,
					},
					"header_type": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string to use as part of the Authorization header. By default, a valid Authorization header looks like this: ` + "`" + `Authorization: ldap base64(username:password)` + "`" + `. If ` + "`" + `header_type` + "`" + ` is set to "basic", then the Authorization header would be ` + "`" + `Authorization: basic base64(username:password)` + "`" + `. Note that ` + "`" + `header_type` + "`" + ` can take any string, not just ` + "`" + `'ldap'` + "`" + ` and ` + "`" + `'basic'` + "`" + `.`,
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
						Description: `Host on which the LDAP server is running.`,
					},
					"ldap_password": schema.StringAttribute{
						Computed:    true,
						Description: `The password to the LDAP server.`,
					},
					"ldap_port": schema.NumberAttribute{
						Computed:    true,
						Description: `TCP port where the LDAP server is listening. 389 is the default port for non-SSL LDAP and AD. 636 is the port required for SSL LDAP and AD. If ` + "`" + `ldaps` + "`" + ` is configured, you must use port 636.`,
					},
					"ldaps": schema.BoolAttribute{
						Computed:    true,
						Description: `Set it to ` + "`" + `true` + "`" + ` to use ` + "`" + `ldaps` + "`" + `, a secure protocol (that can be configured to TLS) to connect to the LDAP server. When ` + "`" + `ldaps` + "`" + ` is configured, you must use port 636. If the ` + "`" + `ldap` + "`" + ` setting is enabled, ensure the ` + "`" + `start_tls` + "`" + ` setting is disabled.`,
					},
					"log_search_results": schema.BoolAttribute{
						Computed:    true,
						Description: `Displays all the LDAP search results received from the LDAP server for debugging purposes. Not recommended to be enabled in a production environment.`,
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

func (r *GatewayPluginLdapAuthAdvancedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginLdapAuthAdvancedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginLdapAuthAdvancedDataSourceModel
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

	request := operations.GetLdapauthadvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetLdapauthadvancedPlugin(ctx, request)
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
	if !(res.LdapAuthAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedLdapAuthAdvancedPlugin(res.LdapAuthAdvancedPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
