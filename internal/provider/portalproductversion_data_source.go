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
var _ datasource.DataSource = &PortalProductVersionDataSource{}
var _ datasource.DataSourceWithConfigure = &PortalProductVersionDataSource{}

func NewPortalProductVersionDataSource() datasource.DataSource {
	return &PortalProductVersionDataSource{}
}

// PortalProductVersionDataSource is the data source implementation.
type PortalProductVersionDataSource struct {
	client *sdk.Konnect
}

// PortalProductVersionDataSourceModel describes the data model.
type PortalProductVersionDataSourceModel struct {
	ApplicationRegistrationEnabled types.Bool             `tfsdk:"application_registration_enabled"`
	AuthStrategies                 []tfTypes.AuthStrategy `tfsdk:"auth_strategies"`
	AutoApproveRegistration        types.Bool             `tfsdk:"auto_approve_registration"`
	CreatedAt                      types.String           `tfsdk:"created_at"`
	Deprecated                     types.Bool             `tfsdk:"deprecated"`
	ID                             types.String           `tfsdk:"id"`
	PortalID                       types.String           `tfsdk:"portal_id"`
	ProductVersionID               types.String           `tfsdk:"product_version_id"`
	PublishStatus                  types.String           `tfsdk:"publish_status"`
	UpdatedAt                      types.String           `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PortalProductVersionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_product_version"
}

// Schema defines the schema for the data source.
func (r *PortalProductVersionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PortalProductVersion DataSource",

		Attributes: map[string]schema.Attribute{
			"application_registration_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the application registration on this portal for the api product version is enabled`,
			},
			"auth_strategies": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_credentials": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"auth_methods": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"credential_type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["client_credentials", "self_managed_client_credentials"]`,
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `The Application Auth Strategy ID.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
							Description: `Client Credential Auth strategy that the application uses.`,
						},
						"key_auth": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"credential_type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["key_auth"]`,
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `The Application Auth Strategy ID.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
							Description: `KeyAuth Auth strategy that the application uses.`,
						},
					},
				},
				Description: `A list of authentication strategies`,
			},
			"auto_approve_registration": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the application registration auto approval on this portal for the api product version is enabled`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
			},
			"deprecated": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the api product version on the portal is deprecated`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Contains a unique identifier used by the API for this resource.`,
			},
			"portal_id": schema.StringAttribute{
				Required:    true,
				Description: `portal identifier`,
			},
			"product_version_id": schema.StringAttribute{
				Required:    true,
				Description: `API product version identifier`,
			},
			"publish_status": schema.StringAttribute{
				Computed:    true,
				Description: `Publication status of the API product version on the portal. must be one of ["published", "unpublished"]`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
			},
		},
	}
}

func (r *PortalProductVersionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PortalProductVersionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PortalProductVersionDataSourceModel
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

	productVersionID := data.ProductVersionID.ValueString()
	portalID := data.PortalID.ValueString()
	request := operations.GetPortalProductVersionRequest{
		ProductVersionID: productVersionID,
		PortalID:         portalID,
	}
	res, err := r.client.PortalProductVersions.GetPortalProductVersion(ctx, request)
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
	if !(res.PortalProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPortalProductVersion(res.PortalProductVersion)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
