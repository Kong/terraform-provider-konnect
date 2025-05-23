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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ServerlessCloudGatewayDataSource{}
var _ datasource.DataSourceWithConfigure = &ServerlessCloudGatewayDataSource{}

func NewServerlessCloudGatewayDataSource() datasource.DataSource {
	return &ServerlessCloudGatewayDataSource{}
}

// ServerlessCloudGatewayDataSource is the data source implementation.
type ServerlessCloudGatewayDataSource struct {
	client *sdk.Konnect
}

// ServerlessCloudGatewayDataSourceModel describes the data model.
type ServerlessCloudGatewayDataSourceModel struct {
	ControlPlane    tfTypes.ServerlessControlPlane `tfsdk:"control_plane"`
	CreatedAt       types.String                   `tfsdk:"created_at"`
	GatewayEndpoint types.String                   `tfsdk:"gateway_endpoint"`
	Labels          map[string]types.String        `tfsdk:"labels"`
	UpdatedAt       types.String                   `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *ServerlessCloudGatewayDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_serverless_cloud_gateway"
}

// Schema defines the schema for the data source.
func (r *ServerlessCloudGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ServerlessCloudGateway DataSource",

		Attributes: map[string]schema.Attribute{
			"control_plane": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `ID of the serverless cloud gateway CP.`,
					},
					"prefix": schema.StringAttribute{
						Computed:    true,
						Description: `The prefix of the serverless cloud gateway CP.`,
					},
					"region": schema.StringAttribute{
						Computed:    true,
						Description: `The control plane region.`,
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"gateway_endpoint": schema.StringAttribute{
				Computed:    true,
				Description: `Endpoint for the serverless cloud gateway.`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `Labels to facilitate tagged search on serverless cloud gateways. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'.`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ServerlessCloudGatewayDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ServerlessCloudGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ServerlessCloudGatewayDataSourceModel
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

	request, requestDiags := data.ToOperationsGetServerlessCloudGatewayRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.ServerlessCloudGateways.GetServerlessCloudGateway(ctx, *request)
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
	if !(res.ServerlessCloudGateway != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedServerlessCloudGateway(ctx, res.ServerlessCloudGateway)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
