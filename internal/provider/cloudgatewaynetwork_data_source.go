// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

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
var _ datasource.DataSource = &CloudGatewayNetworkDataSource{}
var _ datasource.DataSourceWithConfigure = &CloudGatewayNetworkDataSource{}

func NewCloudGatewayNetworkDataSource() datasource.DataSource {
	return &CloudGatewayNetworkDataSource{}
}

// CloudGatewayNetworkDataSource is the data source implementation.
type CloudGatewayNetworkDataSource struct {
	client *sdk.Konnect
}

// CloudGatewayNetworkDataSourceModel describes the data model.
type CloudGatewayNetworkDataSourceModel struct {
	AvailabilityZones             []types.String                  `tfsdk:"availability_zones"`
	CidrBlock                     types.String                    `tfsdk:"cidr_block"`
	CloudGatewayProviderAccountID types.String                    `tfsdk:"cloud_gateway_provider_account_id"`
	ConfigurationReferenceCount   types.Int64                     `tfsdk:"configuration_reference_count"`
	CreatedAt                     types.String                    `tfsdk:"created_at"`
	DdosProtection                types.Bool                      `tfsdk:"ddos_protection"`
	Default                       types.Bool                      `tfsdk:"default"`
	EntityVersion                 types.Int64                     `tfsdk:"entity_version"`
	Firewall                      *tfTypes.NetworkFirewallConfig  `tfsdk:"firewall"`
	ID                            types.String                    `tfsdk:"id"`
	Name                          types.String                    `tfsdk:"name"`
	ProviderMetadata              tfTypes.NetworkProviderMetadata `tfsdk:"provider_metadata"`
	Region                        types.String                    `tfsdk:"region"`
	State                         types.String                    `tfsdk:"state"`
	TransitGatewayCount           types.Int64                     `tfsdk:"transit_gateway_count"`
	UpdatedAt                     types.String                    `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *CloudGatewayNetworkDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_network"
}

// Schema defines the schema for the data source.
func (r *CloudGatewayNetworkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayNetwork DataSource",

		Attributes: map[string]schema.Attribute{
			"availability_zones": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `List of availability zones that the network is attached to.`,
			},
			"cidr_block": schema.StringAttribute{
				Computed:    true,
				Description: `CIDR block configuration for the network.`,
			},
			"cloud_gateway_provider_account_id": schema.StringAttribute{
				Computed: true,
			},
			"configuration_reference_count": schema.Int64Attribute{
				Computed:    true,
				Description: `The number of configurations that reference this network.`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of network creation date.`,
			},
			"ddos_protection": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether DDOS protection is enabled for the network.`,
			},
			"default": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: `Whether the network is a default network or not. Default networks are Networks that are created` + "\n" +
					`automatically by Konnect when an organization is linked to a provider account.` + "\n" +
					``,
			},
			"entity_version": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: `Monotonically-increasing version count of the network, to indicate the order of updates to the network.` + "\n" +
					``,
			},
			"firewall": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"allowed_cidr_blocks": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of allowed CIDR blocks to access a network.`,
					},
					"denied_cidr_blocks": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of denied CIDR blocks to access a network.`,
					},
				},
				Description: `Firewall configuration for a network.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Human-readable name of the network.`,
			},
			"provider_metadata": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"subnet_ids": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"vpc_id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Metadata describing attributes returned by cloud-provider for the network.`,
			},
			"region": schema.StringAttribute{
				Computed:    true,
				Description: `Region ID for cloud provider region.`,
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: `State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]`,
			},
			"transit_gateway_count": schema.Int64Attribute{
				Computed:    true,
				Description: `The number of transit gateways attached to this network.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of network update date.`,
			},
		},
	}
}

func (r *CloudGatewayNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *CloudGatewayNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CloudGatewayNetworkDataSourceModel
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

	var networkID string
	networkID = data.ID.ValueString()

	request := operations.GetNetworkRequest{
		NetworkID: networkID,
	}
	res, err := r.client.Networks.GetNetwork(ctx, request)
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
	if !(res.Network != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedNetwork(res.Network)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}