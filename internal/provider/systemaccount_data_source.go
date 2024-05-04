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
var _ datasource.DataSource = &SystemAccountDataSource{}
var _ datasource.DataSourceWithConfigure = &SystemAccountDataSource{}

func NewSystemAccountDataSource() datasource.DataSource {
	return &SystemAccountDataSource{}
}

// SystemAccountDataSource is the data source implementation.
type SystemAccountDataSource struct {
	client *sdk.Konnect
}

// SystemAccountDataSourceModel describes the data model.
type SystemAccountDataSourceModel struct {
	CreatedAt      types.String `tfsdk:"created_at"`
	Description    types.String `tfsdk:"description"`
	ID             types.String `tfsdk:"id"`
	KonnectManaged types.Bool   `tfsdk:"konnect_managed"`
	Name           types.String `tfsdk:"name"`
	UpdatedAt      types.String `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *SystemAccountDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_account"
}

// Schema defines the schema for the data source.
func (r *SystemAccountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SystemAccount DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `Timestamp of when the system account was created.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `Description of the system account.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the system account.`,
			},
			"konnect_managed": schema.BoolAttribute{
				Computed:    true,
				Description: `The system account is managed by Konnect (true/false).`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Name of the system account.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `Timestamp of when the system account was last updated.`,
			},
		},
	}
}

func (r *SystemAccountDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SystemAccountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SystemAccountDataSourceModel
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

	accountID := data.ID.ValueString()
	request := operations.GetSystemAccountsIDRequest{
		AccountID: accountID,
	}
	res, err := r.client.SystemAccounts.GetSystemAccountsID(ctx, request)
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
	if res.SystemAccount == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSystemAccount(res.SystemAccount)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
