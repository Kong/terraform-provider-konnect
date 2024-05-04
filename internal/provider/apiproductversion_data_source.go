// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

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
var _ datasource.DataSource = &APIProductVersionDataSource{}
var _ datasource.DataSourceWithConfigure = &APIProductVersionDataSource{}

func NewAPIProductVersionDataSource() datasource.DataSource {
	return &APIProductVersionDataSource{}
}

// APIProductVersionDataSource is the data source implementation.
type APIProductVersionDataSource struct {
	client *sdk.Konnect
}

// APIProductVersionDataSourceModel describes the data model.
type APIProductVersionDataSourceModel struct {
	APIProductID   types.String                      `tfsdk:"api_product_id"`
	CreatedAt      types.String                      `tfsdk:"created_at"`
	Deprecated     types.Bool                        `tfsdk:"deprecated"`
	GatewayService *tfTypes.GatewayService           `tfsdk:"gateway_service"`
	ID             types.String                      `tfsdk:"id"`
	Name           types.String                      `tfsdk:"name"`
	Portals        []tfTypes.APIProductVersionPortal `tfsdk:"portals"`
	UpdatedAt      types.String                      `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *APIProductVersionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_product_version"
}

// Schema defines the schema for the data source.
func (r *APIProductVersionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "APIProductVersion DataSource",

		Attributes: map[string]schema.Attribute{
			"api_product_id": schema.StringAttribute{
				Required:    true,
				Description: `The API product identifier`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
			},
			"deprecated": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether this API product version is deprecated in at least one portal. This field is deprecated: Use PortalProductVersion.deprecated instead`,
			},
			"gateway_service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"control_plane_id": schema.StringAttribute{
						Computed:    true,
						Description: `The identifier of the control plane that the gateway service resides in`,
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `The identifier of a gateway service associated with the version of the API product.`,
					},
					"runtime_group_id": schema.StringAttribute{
						Computed:    true,
						Description: `This field is deprecated, please use ` + "`" + `control_plane_id` + "`" + ` instead. The identifier of the control plane that the gateway service resides in`,
					},
				},
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The API product version identifier`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The version of the API product`,
			},
			"portals": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"application_registration_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"auth_strategies": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"deprecated": schema.BoolAttribute{
							Computed: true,
						},
						"portal_id": schema.StringAttribute{
							Computed: true,
						},
						"portal_name": schema.StringAttribute{
							Computed: true,
						},
						"portal_product_version_id": schema.StringAttribute{
							Computed: true,
						},
						"publish_status": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["published", "unpublished"]`,
						},
					},
				},
				Description: `The list of portals which this API product version is configured for`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
			},
		},
	}
}

func (r *APIProductVersionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *APIProductVersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *APIProductVersionDataSourceModel
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

	apiProductID := data.APIProductID.ValueString()
	id := data.ID.ValueString()
	request := operations.GetAPIProductVersionRequest{
		APIProductID: apiProductID,
		ID:           id,
	}
	res, err := r.client.APIProductVersions.GetAPIProductVersion(ctx, request)
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
	if res.APIProductVersion == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersion(res.APIProductVersion)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
