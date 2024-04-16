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
var _ datasource.DataSource = &GatewayConsumerDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayConsumerDataSource{}

func NewGatewayConsumerDataSource() datasource.DataSource {
	return &GatewayConsumerDataSource{}
}

// GatewayConsumerDataSource is the data source implementation.
type GatewayConsumerDataSource struct {
	client *sdk.Konnect
}

// GatewayConsumerDataSourceModel describes the data model.
type GatewayConsumerDataSourceModel struct {
	ControlPlaneID types.String   `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64    `tfsdk:"created_at"`
	CustomID       types.String   `tfsdk:"custom_id"`
	ID             types.String   `tfsdk:"id"`
	Tags           []types.String `tfsdk:"tags"`
	Username       types.String   `tfsdk:"username"`
}

// Metadata returns the data source type name.
func (r *GatewayConsumerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_consumer"
}

// Schema defines the schema for the data source.
func (r *GatewayConsumerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayConsumer DataSource",

		Attributes: map[string]schema.Attribute{
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"custom_id": schema.StringAttribute{
				Computed:    true,
				Description: `Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or ` + "`" + `username` + "`" + ` with the request.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the Consumer to lookup`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Consumer for grouping and filtering.`,
			},
			"username": schema.StringAttribute{
				Computed:    true,
				Description: `The unique username of the Consumer. You must send either this field or ` + "`" + `custom_id` + "`" + ` with the request.`,
			},
		},
	}
}

func (r *GatewayConsumerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayConsumerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayConsumerDataSourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	consumerID := data.ID.ValueString()
	request := operations.GetConsumerRequest{
		ControlPlaneID: controlPlaneID,
		ConsumerID:     consumerID,
	}
	res, err := r.client.Consumers.GetConsumer(ctx, request)
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
	if res.Consumer == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConsumer(res.Consumer)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
