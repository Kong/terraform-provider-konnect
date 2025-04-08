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
var _ datasource.DataSource = &GatewayPluginRequestTransformerAdvancedDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginRequestTransformerAdvancedDataSource{}

func NewGatewayPluginRequestTransformerAdvancedDataSource() datasource.DataSource {
	return &GatewayPluginRequestTransformerAdvancedDataSource{}
}

// GatewayPluginRequestTransformerAdvancedDataSource is the data source implementation.
type GatewayPluginRequestTransformerAdvancedDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginRequestTransformerAdvancedDataSourceModel describes the data model.
type GatewayPluginRequestTransformerAdvancedDataSourceModel struct {
	Config         *tfTypes.RequestTransformerAdvancedPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer              `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLWithoutParentsConsumer              `tfsdk:"consumer_group"`
	ControlPlaneID types.String                                    `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                                     `tfsdk:"created_at"`
	Enabled        types.Bool                                      `tfsdk:"enabled"`
	ID             types.String                                    `tfsdk:"id"`
	InstanceName   types.String                                    `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering                      `tfsdk:"ordering"`
	Protocols      []types.String                                  `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer              `tfsdk:"route"`
	Service        *tfTypes.ACLWithoutParentsConsumer              `tfsdk:"service"`
	Tags           []types.String                                  `tfsdk:"tags"`
	UpdatedAt      types.Int64                                     `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginRequestTransformerAdvancedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_request_transformer_advanced"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginRequestTransformerAdvancedDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginRequestTransformerAdvanced DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"add": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"body": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"headers": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"json_types": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"querystring": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"allow": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"body": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"append": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"body": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"headers": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"json_types": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"querystring": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"dots_in_keys": schema.BoolAttribute{
						Computed:    true,
						Description: `Specify whether dots (for example, ` + "`" + `customers.info.phone` + "`" + `) should be treated as part of a property name or used to descend into nested JSON objects.  See [Arrays and nested objects](#arrays-and-nested-objects).`,
					},
					"http_method": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing an HTTP method, such as GET, POST, PUT, or DELETE. The string must contain only uppercase letters.`,
					},
					"remove": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"body": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"headers": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"querystring": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"rename": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"body": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"headers": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"querystring": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"replace": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"body": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"headers": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"json_types": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"querystring": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"uri": schema.StringAttribute{
								Computed: true,
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

func (r *GatewayPluginRequestTransformerAdvancedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginRequestTransformerAdvancedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginRequestTransformerAdvancedDataSourceModel
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

	request := operations.GetRequesttransformeradvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetRequesttransformeradvancedPlugin(ctx, request)
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
	if !(res.RequestTransformerAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedRequestTransformerAdvancedPlugin(ctx, res.RequestTransformerAdvancedPlugin)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
