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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayHMACAuthDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayHMACAuthDataSource{}

func NewGatewayHMACAuthDataSource() datasource.DataSource {
	return &GatewayHMACAuthDataSource{}
}

// GatewayHMACAuthDataSource is the data source implementation.
type GatewayHMACAuthDataSource struct {
	client *sdk.Konnect
}

// GatewayHMACAuthDataSourceModel describes the data model.
type GatewayHMACAuthDataSourceModel struct {
	ConsumerID     types.String   `tfsdk:"consumer_id"`
	ControlPlaneID types.String   `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64    `tfsdk:"created_at"`
	ID             types.String   `tfsdk:"id"`
	Secret         types.String   `tfsdk:"secret"`
	Tags           []types.String `tfsdk:"tags"`
	Username       types.String   `tfsdk:"username"`
}

// Metadata returns the data source type name.
func (r *GatewayHMACAuthDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_hmac_auth"
}

// Schema defines the schema for the data source.
func (r *GatewayHMACAuthDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayHMACAuth DataSource",

		Attributes: map[string]schema.Attribute{
			"consumer_id": schema.StringAttribute{
				Required:    true,
				Description: `Consumer ID for nested entities`,
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"secret": schema.StringAttribute{
				Computed: true,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"username": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *GatewayHMACAuthDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayHMACAuthDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayHMACAuthDataSourceModel
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

	request, requestDiags := data.ToOperationsGetHmacAuthWithConsumerRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.HMACAuthCredentials.GetHmacAuthWithConsumer(ctx, *request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.HMACAuth != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedHMACAuth(ctx, res.HMACAuth)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
