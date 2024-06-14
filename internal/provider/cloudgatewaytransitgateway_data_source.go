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
var _ datasource.DataSource = &CloudGatewayTransitGatewayDataSource{}
var _ datasource.DataSourceWithConfigure = &CloudGatewayTransitGatewayDataSource{}

func NewCloudGatewayTransitGatewayDataSource() datasource.DataSource {
	return &CloudGatewayTransitGatewayDataSource{}
}

// CloudGatewayTransitGatewayDataSource is the data source implementation.
type CloudGatewayTransitGatewayDataSource struct {
	client *sdk.Konnect
}

// CloudGatewayTransitGatewayDataSourceModel describes the data model.
type CloudGatewayTransitGatewayDataSourceModel struct {
	CidrBlocks                     []types.String                         `tfsdk:"cidr_blocks"`
	CreatedAt                      types.String                           `tfsdk:"created_at"`
	DNSConfig                      []tfTypes.TransitGatewayDNSConfig      `tfsdk:"dns_config"`
	EntityVersion                  types.Int64                            `tfsdk:"entity_version"`
	ID                             types.String                           `tfsdk:"id"`
	Name                           types.String                           `tfsdk:"name"`
	NetworkID                      types.String                           `tfsdk:"network_id"`
	State                          types.String                           `tfsdk:"state"`
	TransitGatewayAttachmentConfig tfTypes.TransitGatewayAttachmentConfig `tfsdk:"transit_gateway_attachment_config"`
	UpdatedAt                      types.String                           `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *CloudGatewayTransitGatewayDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_transit_gateway"
}

// Schema defines the schema for the data source.
func (r *CloudGatewayTransitGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayTransitGateway DataSource",

		Attributes: map[string]schema.Attribute{
			"cidr_blocks": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				MarkdownDescription: `CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning` + "\n" +
					`network.` + "\n" +
					``,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of transit gateway creation date.`,
			},
			"dns_config": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"domain_proxy_list": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							MarkdownDescription: `Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,` + "\n" +
								`for a transit gateway.` + "\n" +
								``,
						},
						"remote_dns_server_ip_addresses": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway.`,
						},
					},
				},
				MarkdownDescription: `List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway` + "\n" +
					`attachment.` + "\n" +
					``,
			},
			"entity_version": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: `Monotonically-increasing version count of the transit gateway, to indicate the order of updates to the` + "\n" +
					`transit gateway.` + "\n" +
					``,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the transit gateway to operate on.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Human-readable name of the transit gateway.`,
			},
			"network_id": schema.StringAttribute{
				Required:    true,
				Description: `The network to operate on.`,
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: `State of the transit gateway. must be one of ["created", "initializing", "ready", "terminating", "terminated"]`,
			},
			"transit_gateway_attachment_config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"aws_transit_gateway_attachment_config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["aws-transit-gateway-attachment"]`,
							},
							"ram_share_arn": schema.StringAttribute{
								Computed:    true,
								Description: `Resource Share ARN to verify request to create transit gateway attachment.`,
							},
							"transit_gateway_id": schema.StringAttribute{
								Computed:    true,
								Description: `AWS Transit Gateway ID to create attachment to.`,
							},
						},
					},
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of transit gateway update date.`,
			},
		},
	}
}

func (r *CloudGatewayTransitGatewayDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *CloudGatewayTransitGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CloudGatewayTransitGatewayDataSourceModel
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

	networkID := data.NetworkID.ValueString()
	transitGatewayID := data.ID.ValueString()
	request := operations.GetTransitGatewayRequest{
		NetworkID:        networkID,
		TransitGatewayID: transitGatewayID,
	}
	res, err := r.client.TransitGateways.GetTransitGateway(ctx, request)
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
	if !(res.TransitGateway != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTransitGateway(res.TransitGateway)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
