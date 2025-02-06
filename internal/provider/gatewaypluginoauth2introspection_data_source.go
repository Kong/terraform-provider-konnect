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
var _ datasource.DataSource = &GatewayPluginOauth2IntrospectionDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginOauth2IntrospectionDataSource{}

func NewGatewayPluginOauth2IntrospectionDataSource() datasource.DataSource {
	return &GatewayPluginOauth2IntrospectionDataSource{}
}

// GatewayPluginOauth2IntrospectionDataSource is the data source implementation.
type GatewayPluginOauth2IntrospectionDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginOauth2IntrospectionDataSourceModel describes the data model.
type GatewayPluginOauth2IntrospectionDataSourceModel struct {
	Config         tfTypes.Oauth2IntrospectionPluginConfig `tfsdk:"config"`
	ControlPlaneID types.String                            `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                             `tfsdk:"created_at"`
	Enabled        types.Bool                              `tfsdk:"enabled"`
	ID             types.String                            `tfsdk:"id"`
	InstanceName   types.String                            `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering              `tfsdk:"ordering"`
	Protocols      []types.String                          `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer      `tfsdk:"route"`
	Service        *tfTypes.ACLWithoutParentsConsumer      `tfsdk:"service"`
	Tags           []types.String                          `tfsdk:"tags"`
	UpdatedAt      types.Int64                             `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginOauth2IntrospectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_oauth2_introspection"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginOauth2IntrospectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginOauth2Introspection DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure ` + "`" + `4xx` + "`" + `. Note that this value must refer to the consumer ` + "`" + `id` + "`" + ` or ` + "`" + `username` + "`" + ` attribute, and **not** its ` + "`" + `custom_id` + "`" + `.`,
					},
					"authorization_value": schema.StringAttribute{
						Computed:    true,
						Description: `The value to set as the ` + "`" + `Authorization` + "`" + ` header when querying the introspection endpoint. This depends on the OAuth 2.0 server, but usually is the ` + "`" + `client_id` + "`" + ` and ` + "`" + `client_secret` + "`" + ` as a Base64-encoded Basic Auth string (` + "`" + `Basic MG9hNWl...` + "`" + `).`,
					},
					"consumer_by": schema.StringAttribute{
						Computed:    true,
						Description: `A string indicating whether to associate OAuth2 ` + "`" + `username` + "`" + ` or ` + "`" + `client_id` + "`" + ` with the consumer's username. OAuth2 ` + "`" + `username` + "`" + ` is mapped to a consumer's ` + "`" + `username` + "`" + ` field, while an OAuth2 ` + "`" + `client_id` + "`" + ` maps to a consumer's ` + "`" + `custom_id` + "`" + `.`,
					},
					"custom_claims_forward": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of custom claims to be forwarded from the introspection response to the upstream request. Claims are forwarded in headers with prefix ` + "`" + `X-Credential-{claim-name}` + "`" + `.`,
					},
					"custom_introspection_headers": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `A list of custom headers to be added in the introspection request.`,
					},
					"hide_credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value telling the plugin to hide the credential to the upstream API server. It will be removed by Kong before proxying the request.`,
					},
					"introspect_request": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean indicating whether to forward information about the current downstream request to the introspect endpoint. If true, headers ` + "`" + `X-Request-Path` + "`" + ` and ` + "`" + `X-Request-Http-Method` + "`" + ` will be inserted into the introspect request.`,
					},
					"introspection_url": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"keepalive": schema.Int64Attribute{
						Computed:    true,
						Description: `An optional value in milliseconds that defines how long an idle connection lives before being closed.`,
					},
					"run_on_preflight": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that indicates whether the plugin should run (and try to authenticate) on ` + "`" + `OPTIONS` + "`" + ` preflight requests. If set to ` + "`" + `false` + "`" + `, then ` + "`" + `OPTIONS` + "`" + ` requests will always be allowed.`,
					},
					"timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `An optional timeout in milliseconds when sending data to the upstream server.`,
					},
					"token_type_hint": schema.StringAttribute{
						Computed:    true,
						Description: `The ` + "`" + `token_type_hint` + "`" + ` value to associate to introspection requests.`,
					},
					"ttl": schema.NumberAttribute{
						Computed:    true,
						Description: `The TTL in seconds for the introspection response. Set to 0 to disable the expiration.`,
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

func (r *GatewayPluginOauth2IntrospectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginOauth2IntrospectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginOauth2IntrospectionDataSourceModel
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

	request := operations.GetOauth2introspectionPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetOauth2introspectionPlugin(ctx, request)
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
	if !(res.Oauth2IntrospectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOauth2IntrospectionPlugin(res.Oauth2IntrospectionPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
