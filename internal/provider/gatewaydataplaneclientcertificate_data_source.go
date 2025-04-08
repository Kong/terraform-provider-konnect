// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayDataPlaneClientCertificateDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayDataPlaneClientCertificateDataSource{}

func NewGatewayDataPlaneClientCertificateDataSource() datasource.DataSource {
	return &GatewayDataPlaneClientCertificateDataSource{}
}

// GatewayDataPlaneClientCertificateDataSource is the data source implementation.
type GatewayDataPlaneClientCertificateDataSource struct {
	client *sdk.Konnect
}

// GatewayDataPlaneClientCertificateDataSourceModel describes the data model.
type GatewayDataPlaneClientCertificateDataSourceModel struct {
	Cert           types.String `tfsdk:"cert"`
	ControlPlaneID types.String `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64  `tfsdk:"created_at"`
	ID             types.String `tfsdk:"id"`
	UpdatedAt      types.Int64  `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayDataPlaneClientCertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_data_plane_client_certificate"
}

// Schema defines the schema for the data source.
func (r *GatewayDataPlaneClientCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayDataPlaneClientCertificate DataSource",

		Attributes: map[string]schema.Attribute{
			"cert": schema.StringAttribute{
				Computed:    true,
				Description: `JSON escaped string of the certificate.`,
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Date certificate was created.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Unique ID of the certificate entity.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Date certificate was last updated.`,
			},
		},
	}
}

func (r *GatewayDataPlaneClientCertificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayDataPlaneClientCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayDataPlaneClientCertificateDataSourceModel
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

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	var certificateID string
	certificateID = data.ID.ValueString()

	request := operations.GetDataplaneCertificateRequest{
		ControlPlaneID: controlPlaneID,
		CertificateID:  certificateID,
	}
	res, err := r.client.DPCertificates.GetDataplaneCertificate(ctx, request)
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
	if !(res.DataPlaneClientCertificateResponse != nil && res.DataPlaneClientCertificateResponse.Item != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedDataPlaneClientCertificate(ctx, res.DataPlaneClientCertificateResponse.Item)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
