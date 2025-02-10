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
var _ datasource.DataSource = &CloudGatewayProviderAccountListDataSource{}
var _ datasource.DataSourceWithConfigure = &CloudGatewayProviderAccountListDataSource{}

func NewCloudGatewayProviderAccountListDataSource() datasource.DataSource {
	return &CloudGatewayProviderAccountListDataSource{}
}

// CloudGatewayProviderAccountListDataSource is the data source implementation.
type CloudGatewayProviderAccountListDataSource struct {
	client *sdk.Konnect
}

// CloudGatewayProviderAccountListDataSourceModel describes the data model.
type CloudGatewayProviderAccountListDataSourceModel struct {
	Data       []tfTypes.ProviderAccount `tfsdk:"data"`
	Meta       tfTypes.PaginatedMeta     `tfsdk:"meta"`
	PageNumber types.Int64               `tfsdk:"page_number"`
	PageSize   types.Int64               `tfsdk:"page_size"`
}

// Metadata returns the data source type name.
func (r *CloudGatewayProviderAccountListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_provider_account_list"
}

// Schema defines the schema for the data source.
func (r *CloudGatewayProviderAccountListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayProviderAccountList DataSource",

		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: `An RFC-3339 timestamp representation of provider account creation date.`,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"provider": schema.StringAttribute{
							Computed:    true,
							Description: `Name of cloud provider.`,
						},
						"provider_account_id": schema.StringAttribute{
							Computed:    true,
							Description: `ID of the cloud provider account.`,
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: `An RFC-3339 timestamp representation of provider account update date.`,
						},
					},
				},
			},
			"meta": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"page": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"number": schema.NumberAttribute{
								Computed: true,
							},
							"size": schema.NumberAttribute{
								Computed: true,
							},
							"total": schema.NumberAttribute{
								Computed: true,
							},
						},
						Description: `Contains pagination query parameters and the total number of objects returned.`,
					},
				},
				Description: `returns the pagination information`,
			},
			"page_number": schema.Int64Attribute{
				Optional:    true,
				Description: `Determines which page of the entities to retrieve.`,
			},
			"page_size": schema.Int64Attribute{
				Optional:    true,
				Description: `The maximum number of items to include per page. The last page of a collection may include fewer items.`,
			},
		},
	}
}

func (r *CloudGatewayProviderAccountListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *CloudGatewayProviderAccountListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CloudGatewayProviderAccountListDataSourceModel
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

	pageSize := new(int64)
	if !data.PageSize.IsUnknown() && !data.PageSize.IsNull() {
		*pageSize = data.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	pageNumber := new(int64)
	if !data.PageNumber.IsUnknown() && !data.PageNumber.IsNull() {
		*pageNumber = data.PageNumber.ValueInt64()
	} else {
		pageNumber = nil
	}
	request := operations.ListProviderAccountsRequest{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
	res, err := r.client.CloudGateways.ListProviderAccounts(ctx, request)
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
	if !(res.ListProviderAccountsResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedListProviderAccountsResponse(res.ListProviderAccountsResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
