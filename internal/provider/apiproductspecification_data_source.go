// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &APIProductSpecificationDataSource{}
var _ datasource.DataSourceWithConfigure = &APIProductSpecificationDataSource{}

func NewAPIProductSpecificationDataSource() datasource.DataSource {
	return &APIProductSpecificationDataSource{}
}

// APIProductSpecificationDataSource is the data source implementation.
type APIProductSpecificationDataSource struct {
	client *sdk.Konnect
}

// APIProductSpecificationDataSourceModel describes the data model.
type APIProductSpecificationDataSourceModel struct {
	APIProductID        types.String `tfsdk:"api_product_id"`
	APIProductVersionID types.String `tfsdk:"api_product_version_id"`
	Content             types.String `tfsdk:"content"`
	CreatedAt           types.String `tfsdk:"created_at"`
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *APIProductSpecificationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_product_specification"
}

// Schema defines the schema for the data source.
func (r *APIProductSpecificationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "APIProductSpecification DataSource",

		Attributes: map[string]schema.Attribute{
			"api_product_id": schema.StringAttribute{
				Required:    true,
				Description: `The API product identifier`,
			},
			"api_product_version_id": schema.StringAttribute{
				Required:    true,
				Description: `The API product version identifier`,
			},
			"content": schema.StringAttribute{
				Computed:    true,
				Description: `The contents of the API product version specification`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The API product version specification identifier`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the API product version specification`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
			},
		},
	}
}

func (r *APIProductSpecificationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *APIProductSpecificationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *APIProductSpecificationDataSourceModel
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
	apiProductVersionID := data.APIProductVersionID.ValueString()
	specificationID := data.ID.ValueString()
	request := operations.GetAPIProductVersionSpecRequest{
		APIProductID:        apiProductID,
		APIProductVersionID: apiProductVersionID,
		SpecificationID:     specificationID,
	}
	res, err := r.client.APIProductVersionSpecification.GetAPIProductVersionSpec(ctx, request)
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
	if !(res.APIProductVersionSpec != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersionSpec(res.APIProductVersionSpec)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
