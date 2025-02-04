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
var _ datasource.DataSource = &GatewayPluginOasValidationDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginOasValidationDataSource{}

func NewGatewayPluginOasValidationDataSource() datasource.DataSource {
	return &GatewayPluginOasValidationDataSource{}
}

// GatewayPluginOasValidationDataSource is the data source implementation.
type GatewayPluginOasValidationDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginOasValidationDataSourceModel describes the data model.
type GatewayPluginOasValidationDataSourceModel struct {
	Config         tfTypes.OasValidationPluginConfig  `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer `tfsdk:"consumer"`
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
func (r *GatewayPluginOasValidationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_oas_validation"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginOasValidationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginOasValidation DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed_header_parameters": schema.StringAttribute{
						Computed:    true,
						Description: `List of header parameters in the request that will be ignored when performing HTTP header validation. These are additional headers added to an API request beyond those defined in the API specification.  For example, you might include the HTTP header ` + "`" + `User-Agent` + "`" + `, which lets servers and network peers identify the application, operating system, vendor, and/or version of the requesting user agent.`,
					},
					"api_spec": schema.StringAttribute{
						Computed:    true,
						Description: `The API specification defined using either Swagger or the OpenAPI. This can be either a JSON or YAML based file. If using a YAML file, the spec needs to be URI-Encoded to preserve the YAML format.`,
					},
					"api_spec_encoded": schema.BoolAttribute{
						Computed:    true,
						Description: `Indicates whether the api_spec is URI-Encoded.`,
					},
					"custom_base_path": schema.StringAttribute{
						Computed:    true,
						Description: `The base path to be used for path match evaluation. This value is ignored if ` + "`" + `include_base_path` + "`" + ` is set to ` + "`" + `false` + "`" + `.`,
					},
					"header_parameter_check": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, checks if HTTP header parameters in the request exist in the API specification.`,
					},
					"include_base_path": schema.BoolAttribute{
						Computed:    true,
						Description: `Indicates whether to include the base path when performing path match evaluation.`,
					},
					"notify_only_request_validation_failure": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, notifications via event hooks are enabled, but request based validation failures don't affect the request flow.`,
					},
					"notify_only_response_body_validation_failure": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, notifications via event hooks are enabled, but response validation failures don't affect the response flow.`,
					},
					"query_parameter_check": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, checks if query parameters in the request exist in the API specification.`,
					},
					"validate_request_body": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, validates the request body content against the API specification.`,
					},
					"validate_request_header_params": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, validates HTTP header parameters against the API specification.`,
					},
					"validate_request_query_params": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, validates query parameters against the API specification.`,
					},
					"validate_request_uri_params": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, validates URI parameters in the request against the API specification.`,
					},
					"validate_response_body": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, validates the response from the upstream services against the API specification. If validation fails, it results in an ` + "`" + `HTTP 406 Not Acceptable` + "`" + ` status code.`,
					},
					"verbose_response": schema.BoolAttribute{
						Computed:    true,
						Description: `If set to true, returns a detailed error message for invalid requests & responses. This is useful while testing.`,
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

func (r *GatewayPluginOasValidationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginOasValidationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginOasValidationDataSourceModel
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

	request := operations.GetOasvalidationPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetOasvalidationPlugin(ctx, request)
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
	if !(res.OasValidationPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOasValidationPlugin(res.OasValidationPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
