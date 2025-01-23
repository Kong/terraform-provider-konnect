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
var _ datasource.DataSource = &GatewayPluginOauth2DataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginOauth2DataSource{}

func NewGatewayPluginOauth2DataSource() datasource.DataSource {
	return &GatewayPluginOauth2DataSource{}
}

// GatewayPluginOauth2DataSource is the data source implementation.
type GatewayPluginOauth2DataSource struct {
	client *sdk.Konnect
}

// GatewayPluginOauth2DataSourceModel describes the data model.
type GatewayPluginOauth2DataSourceModel struct {
	Config         tfTypes.Oauth2PluginConfig         `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer" tfPlanOnly:"true"`
	ConsumerGroup  *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer_group" tfPlanOnly:"true"`
	ControlPlaneID types.String                       `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                        `tfsdk:"created_at"`
	Enabled        types.Bool                         `tfsdk:"enabled"`
	ID             types.String                       `tfsdk:"id"`
	InstanceName   types.String                       `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering         `tfsdk:"ordering"`
	Protocols      []types.String                     `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer `tfsdk:"route" tfPlanOnly:"true"`
	Service        *tfTypes.ACLWithoutParentsConsumer `tfsdk:"service" tfPlanOnly:"true"`
	Tags           []types.String                     `tfsdk:"tags"`
	UpdatedAt      types.Int64                        `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginOauth2DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_oauth2"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginOauth2DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginOauth2 DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"accept_http_if_already_terminated": schema.BoolAttribute{
						Computed:    true,
						Description: `Accepts HTTPs requests that have already been terminated by a proxy or load balancer.`,
					},
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.`,
					},
					"auth_header_name": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the header that is supposed to carry the access token.`,
					},
					"enable_authorization_code": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value to enable the three-legged Authorization Code flow (RFC 6742 Section 4.1).`,
					},
					"enable_client_credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value to enable the Client Credentials Grant flow (RFC 6742 Section 4.4).`,
					},
					"enable_implicit_grant": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value to enable the Implicit Grant flow which allows to provision a token as a result of the authorization process (RFC 6742 Section 4.2).`,
					},
					"enable_password_grant": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value to enable the Resource Owner Password Credentials Grant flow (RFC 6742 Section 4.3).`,
					},
					"global_credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value that allows using the same OAuth credentials generated by the plugin with any other service whose OAuth 2.0 plugin configuration also has ` + "`" + `config.global_credentials=true` + "`" + `.`,
					},
					"hide_credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value telling the plugin to show or hide the credential from the upstream service.`,
					},
					"mandatory_scope": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value telling the plugin to require at least one ` + "`" + `scope` + "`" + ` to be authorized by the end user.`,
					},
					"persistent_refresh_token": schema.BoolAttribute{
						Computed: true,
					},
					"pkce": schema.StringAttribute{
						Computed:    true,
						Description: `Specifies a mode of how the Proof Key for Code Exchange (PKCE) should be handled by the plugin.`,
					},
					"provision_key": schema.StringAttribute{
						Computed:    true,
						Description: `The unique key the plugin has generated when it has been added to the Service.`,
					},
					"realm": schema.StringAttribute{
						Computed:    true,
						Description: `When authentication fails the plugin sends ` + "`" + `WWW-Authenticate` + "`" + ` header with ` + "`" + `realm` + "`" + ` attribute value.`,
					},
					"refresh_token_ttl": schema.NumberAttribute{
						Computed:    true,
						Description: `Time-to-live value for data`,
					},
					"reuse_refresh_token": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value that indicates whether an OAuth refresh token is reused when refreshing an access token.`,
					},
					"scopes": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Describes an array of scope names that will be available to the end user. If ` + "`" + `mandatory_scope` + "`" + ` is set to ` + "`" + `true` + "`" + `, then ` + "`" + `scopes` + "`" + ` are required.`,
					},
					"token_expiration": schema.NumberAttribute{
						Computed:    true,
						Description: `An optional integer value telling the plugin how many seconds a token should last, after which the client will need to refresh the token. Set to ` + "`" + `0` + "`" + ` to disable the expiration.`,
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

func (r *GatewayPluginOauth2DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginOauth2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginOauth2DataSourceModel
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

	request := operations.GetOauth2PluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetOauth2Plugin(ctx, request)
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
	if !(res.Oauth2Plugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOauth2Plugin(res.Oauth2Plugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
