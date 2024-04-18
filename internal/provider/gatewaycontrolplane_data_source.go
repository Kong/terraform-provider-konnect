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
var _ datasource.DataSource = &GatewayControlPlaneDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayControlPlaneDataSource{}

func NewGatewayControlPlaneDataSource() datasource.DataSource {
	return &GatewayControlPlaneDataSource{}
}

// GatewayControlPlaneDataSource is the data source implementation.
type GatewayControlPlaneDataSource struct {
	client *sdk.Konnect
}

// GatewayControlPlaneDataSourceModel describes the data model.
type GatewayControlPlaneDataSourceModel struct {
	Config      *tfTypes.Config         `tfsdk:"config"`
	Description types.String            `tfsdk:"description"`
	ID          types.String            `tfsdk:"id"`
	Labels      map[string]types.String `tfsdk:"labels"`
	Name        types.String            `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (r *GatewayControlPlaneDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_control_plane"
}

// Schema defines the schema for the data source.
func (r *GatewayControlPlaneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayControlPlane DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `Control Plane Endpoint.`,
					},
					"telemetry_endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `Telemetry Endpoint.`,
					},
				},
				Description: `CP configuration object for related access endpoints.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `The description of the control plane in Konnect.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The control plane ID`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
					`` + "\n" +
					`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".` + "\n" +
					``,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the control plane.`,
			},
		},
	}
}

func (r *GatewayControlPlaneDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayControlPlaneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayControlPlaneDataSourceModel
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

	id := data.ID.ValueString()
	request := operations.GetControlPlaneRequest{
		ID: id,
	}
	res, err := r.client.ControlPlanes.GetControlPlane(ctx, request)
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
	if res.ControlPlane == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedControlPlane(res.ControlPlane)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
