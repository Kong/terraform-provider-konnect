// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudGatewayTransitGatewayResource{}
var _ resource.ResourceWithImportState = &CloudGatewayTransitGatewayResource{}

func NewCloudGatewayTransitGatewayResource() resource.Resource {
	return &CloudGatewayTransitGatewayResource{}
}

// CloudGatewayTransitGatewayResource defines the resource implementation.
type CloudGatewayTransitGatewayResource struct {
	client *sdk.Konnect
}

// CloudGatewayTransitGatewayResourceModel describes the resource data model.
type CloudGatewayTransitGatewayResourceModel struct {
	AWSTransitGateway            *tfTypes.AWSTransitGateway            `queryParam:"inline" tfsdk:"aws_transit_gateway" tfPlanOnly:"true"`
	AwsTransitGatewayResponse    *tfTypes.AwsTransitGatewayResponse    `queryParam:"inline" tfsdk:"aws_transit_gateway_response" tfPlanOnly:"true"`
	AWSVpcPeeringGateway         *tfTypes.AWSVpcPeeringGateway         `queryParam:"inline" tfsdk:"aws_vpc_peering_gateway" tfPlanOnly:"true"`
	AwsVpcPeeringGatewayResponse *tfTypes.AwsVpcPeeringGatewayResponse `queryParam:"inline" tfsdk:"aws_vpc_peering_gateway_response" tfPlanOnly:"true"`
	AzureTransitGateway          *tfTypes.AzureTransitGateway          `queryParam:"inline" tfsdk:"azure_transit_gateway" tfPlanOnly:"true"`
	AzureTransitGatewayResponse  *tfTypes.AzureTransitGatewayResponse  `queryParam:"inline" tfsdk:"azure_transit_gateway_response" tfPlanOnly:"true"`
	EntityVersion                types.Int64                           `tfsdk:"entity_version"`
	ID                           types.String                          `tfsdk:"id"`
	Name                         types.String                          `tfsdk:"name"`
	NetworkID                    types.String                          `tfsdk:"network_id"`
}

func (r *CloudGatewayTransitGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_transit_gateway"
}

func (r *CloudGatewayTransitGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayTransitGateway Resource",
		Attributes: map[string]schema.Attribute{
			"aws_transit_gateway": schema.SingleNestedAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Attributes: map[string]schema.Attribute{
					"cidr_blocks": schema.ListAttribute{
						Required: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						ElementType: types.StringType,
						MarkdownDescription: `CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning` + "\n" +
							`network.` + "\n" +
							`Requires replacement if changed.`,
					},
					"dns_config": schema.ListNestedAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
							},
							Attributes: map[string]schema.Attribute{
								"domain_proxy_list": schema.ListAttribute{
									Required: true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.RequiresReplaceIfConfigured(),
									},
									ElementType: types.StringType,
									MarkdownDescription: `Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,` + "\n" +
										`for a transit gateway.` + "\n" +
										`Requires replacement if changed.`,
								},
								"remote_dns_server_ip_addresses": schema.ListAttribute{
									Required: true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.RequiresReplaceIfConfigured(),
									},
									ElementType: types.StringType,
									Description: `Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway. Requires replacement if changed.`,
								},
							},
						},
						MarkdownDescription: `List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway` + "\n" +
							`attachment.` + "\n" +
							`Requires replacement if changed.`,
					},
					"name": schema.StringAttribute{
						Required: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Description: `Human-readable name of the transit gateway. Requires replacement if changed.`,
					},
					"transit_gateway_attachment_config": schema.SingleNestedAttribute{
						Required: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
						},
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `must be "aws-transit-gateway-attachment"; Requires replacement if changed.`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"aws-transit-gateway-attachment",
									),
								},
							},
							"ram_share_arn": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `Resource Share ARN to verify request to create transit gateway attachment. Requires replacement if changed.`,
							},
							"transit_gateway_id": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `AWS Transit Gateway ID to create attachment to. Requires replacement if changed.`,
							},
						},
						Description: `Requires replacement if changed.`,
					},
				},
				Description: `Requires replacement if changed.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("aws_transit_gateway_response"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway_response"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway_response"),
					}...),
				},
			},
			"aws_transit_gateway_response": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"cidr_blocks": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						MarkdownDescription: `CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning` + "\n" +
							`network.`,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An RFC-3339 timestamp representation of transit gateway creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dns_config": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"domain_proxy_list": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									MarkdownDescription: `Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,` + "\n" +
										`for a transit gateway.`,
								},
								"remote_dns_server_ip_addresses": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway.`,
								},
							},
						},
						MarkdownDescription: `List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway` + "\n" +
							`attachment.`,
					},
					"entity_version": schema.Int64Attribute{
						Computed: true,
						MarkdownDescription: `Monotonically-increasing version count of the transit gateway, to indicate the order of updates to the` + "\n" +
							`transit gateway.`,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `Human-readable name of the transit gateway.`,
					},
					"state": schema.StringAttribute{
						Computed: true,
						MarkdownDescription: `The current state of the Transit Gateway. Possible values:` + "\n" +
							`- ` + "`" + `created` + "`" + ` - The attachment has been created but is not attached to transit gateway.` + "\n" +
							`- ` + "`" + `initializing` + "`" + ` - The attachment is in the process of being initialized and is setting up necessary resources.` + "\n" +
							`- ` + "`" + `pending-acceptance` + "`" + ` The attachment request is awaiting acceptance in customer VPC.` + "\n" +
							`- ` + "`" + `ready` + "`" + ` - The transit gateway attachment is fully operational and can route traffic as configured.` + "\n" +
							`- ` + "`" + `terminating` + "`" + ` - The attachment is in the process of being deleted and is no longer accepting new traffic.` + "\n" +
							`- ` + "`" + `terminated` + "`" + ` - The attachment has been fully deleted and is no longer available.` + "\n" +
							`must be one of ["created", "initializing", "pending-acceptance", "ready", "terminating", "terminated"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"created",
								"initializing",
								"pending-acceptance",
								"ready",
								"terminating",
								"terminated",
							),
						},
					},
					"transit_gateway_attachment_config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Computed:    true,
								Description: `must be "aws-transit-gateway-attachment"`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"aws-transit-gateway-attachment",
									),
								},
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
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An RFC-3339 timestamp representation of transit gateway update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("aws_transit_gateway"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway_response"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway_response"),
					}...),
				},
			},
			"aws_vpc_peering_gateway": schema.SingleNestedAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Attributes: map[string]schema.Attribute{
					"cidr_blocks": schema.ListAttribute{
						Required: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						ElementType: types.StringType,
						MarkdownDescription: `CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning` + "\n" +
							`network.` + "\n" +
							`Requires replacement if changed.`,
					},
					"dns_config": schema.ListNestedAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
							},
							Attributes: map[string]schema.Attribute{
								"domain_proxy_list": schema.ListAttribute{
									Required: true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.RequiresReplaceIfConfigured(),
									},
									ElementType: types.StringType,
									MarkdownDescription: `Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,` + "\n" +
										`for a transit gateway.` + "\n" +
										`Requires replacement if changed.`,
								},
								"remote_dns_server_ip_addresses": schema.ListAttribute{
									Required: true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.RequiresReplaceIfConfigured(),
									},
									ElementType: types.StringType,
									Description: `Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway. Requires replacement if changed.`,
								},
							},
						},
						MarkdownDescription: `List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway` + "\n" +
							`attachment.` + "\n" +
							`Requires replacement if changed.`,
					},
					"name": schema.StringAttribute{
						Required: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Description: `Human-readable name of the transit gateway. Requires replacement if changed.`,
					},
					"transit_gateway_attachment_config": schema.SingleNestedAttribute{
						Required: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
						},
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `must be "aws-vpc-peering-attachment"; Requires replacement if changed.`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"aws-vpc-peering-attachment",
									),
								},
							},
							"peer_account_id": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `Requires replacement if changed.`,
							},
							"peer_vpc_id": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `Requires replacement if changed.`,
							},
							"peer_vpc_region": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `Requires replacement if changed.`,
							},
						},
						Description: `Requires replacement if changed.`,
					},
				},
				Description: `Requires replacement if changed.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("aws_transit_gateway"),
						path.MatchRelative().AtParent().AtName("aws_transit_gateway_response"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway_response"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway_response"),
					}...),
				},
			},
			"aws_vpc_peering_gateway_response": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"cidr_blocks": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						MarkdownDescription: `CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning` + "\n" +
							`network.`,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An RFC-3339 timestamp representation of transit gateway creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dns_config": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"domain_proxy_list": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									MarkdownDescription: `Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,` + "\n" +
										`for a transit gateway.`,
								},
								"remote_dns_server_ip_addresses": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway.`,
								},
							},
						},
						MarkdownDescription: `List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway` + "\n" +
							`attachment.`,
					},
					"entity_version": schema.Int64Attribute{
						Computed: true,
						MarkdownDescription: `Monotonically-increasing version count of the transit gateway, to indicate the order of updates to the` + "\n" +
							`transit gateway.`,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `Human-readable name of the transit gateway.`,
					},
					"state": schema.StringAttribute{
						Computed: true,
						MarkdownDescription: `The current state of the Transit Gateway. Possible values:` + "\n" +
							`- ` + "`" + `created` + "`" + ` - The attachment has been created but is not attached to transit gateway.` + "\n" +
							`- ` + "`" + `initializing` + "`" + ` - The attachment is in the process of being initialized and is setting up necessary resources.` + "\n" +
							`- ` + "`" + `pending-acceptance` + "`" + ` The attachment request is awaiting acceptance in customer VPC.` + "\n" +
							`- ` + "`" + `ready` + "`" + ` - The transit gateway attachment is fully operational and can route traffic as configured.` + "\n" +
							`- ` + "`" + `terminating` + "`" + ` - The attachment is in the process of being deleted and is no longer accepting new traffic.` + "\n" +
							`- ` + "`" + `terminated` + "`" + ` - The attachment has been fully deleted and is no longer available.` + "\n" +
							`must be one of ["created", "initializing", "pending-acceptance", "ready", "terminating", "terminated"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"created",
								"initializing",
								"pending-acceptance",
								"ready",
								"terminating",
								"terminated",
							),
						},
					},
					"transit_gateway_attachment_config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Computed:    true,
								Description: `must be "aws-vpc-peering-attachment"`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"aws-vpc-peering-attachment",
									),
								},
							},
							"peer_account_id": schema.StringAttribute{
								Computed: true,
							},
							"peer_vpc_id": schema.StringAttribute{
								Computed: true,
							},
							"peer_vpc_region": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An RFC-3339 timestamp representation of transit gateway update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("aws_transit_gateway"),
						path.MatchRelative().AtParent().AtName("aws_transit_gateway_response"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway_response"),
					}...),
				},
			},
			"azure_transit_gateway": schema.SingleNestedAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Attributes: map[string]schema.Attribute{
					"dns_config": schema.ListNestedAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
							},
							Attributes: map[string]schema.Attribute{
								"domain_proxy_list": schema.ListAttribute{
									Required: true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.RequiresReplaceIfConfigured(),
									},
									ElementType: types.StringType,
									MarkdownDescription: `Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,` + "\n" +
										`for a transit gateway.` + "\n" +
										`Requires replacement if changed.`,
								},
								"remote_dns_server_ip_addresses": schema.ListAttribute{
									Required: true,
									PlanModifiers: []planmodifier.List{
										listplanmodifier.RequiresReplaceIfConfigured(),
									},
									ElementType: types.StringType,
									Description: `Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway. Requires replacement if changed.`,
								},
							},
						},
						MarkdownDescription: `List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway` + "\n" +
							`attachment.` + "\n" +
							`Requires replacement if changed.`,
					},
					"name": schema.StringAttribute{
						Required: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Description: `Human-readable name of the transit gateway. Requires replacement if changed.`,
					},
					"transit_gateway_attachment_config": schema.SingleNestedAttribute{
						Required: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
						},
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `must be "azure-vnet-peering-attachment"; Requires replacement if changed.`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"azure-vnet-peering-attachment",
									),
								},
							},
							"resource_group_name": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `Resource Group Name for the Azure VNET Peering attachment. Requires replacement if changed.`,
							},
							"subscription_id": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `Subscription ID for the Azure VNET Peering attachment. Requires replacement if changed.`,
							},
							"tenant_id": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `Tenant ID for the Azure VNET Peering attachment. Requires replacement if changed.`,
							},
							"vnet_name": schema.StringAttribute{
								Required: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
								},
								Description: `VNET Name for the Azure VNET Peering attachment. Requires replacement if changed.`,
							},
						},
						Description: `Requires replacement if changed.`,
					},
				},
				Description: `Requires replacement if changed.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("aws_transit_gateway"),
						path.MatchRelative().AtParent().AtName("aws_transit_gateway_response"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway_response"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway_response"),
					}...),
				},
			},
			"azure_transit_gateway_response": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An RFC-3339 timestamp representation of transit gateway creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dns_config": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"domain_proxy_list": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									MarkdownDescription: `Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,` + "\n" +
										`for a transit gateway.`,
								},
								"remote_dns_server_ip_addresses": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway.`,
								},
							},
						},
						MarkdownDescription: `List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway` + "\n" +
							`attachment.`,
					},
					"entity_version": schema.Int64Attribute{
						Computed: true,
						MarkdownDescription: `Monotonically-increasing version count of the transit gateway, to indicate the order of updates to the` + "\n" +
							`transit gateway.`,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `Human-readable name of the transit gateway.`,
					},
					"state": schema.StringAttribute{
						Computed: true,
						MarkdownDescription: `The current state of the Transit Gateway. Possible values:` + "\n" +
							`- ` + "`" + `created` + "`" + ` - The attachment has been created but is not attached to transit gateway.` + "\n" +
							`- ` + "`" + `initializing` + "`" + ` - The attachment is in the process of being initialized and is setting up necessary resources.` + "\n" +
							`- ` + "`" + `pending-acceptance` + "`" + ` The attachment request is awaiting acceptance in customer VPC.` + "\n" +
							`- ` + "`" + `ready` + "`" + ` - The transit gateway attachment is fully operational and can route traffic as configured.` + "\n" +
							`- ` + "`" + `terminating` + "`" + ` - The attachment is in the process of being deleted and is no longer accepting new traffic.` + "\n" +
							`- ` + "`" + `terminated` + "`" + ` - The attachment has been fully deleted and is no longer available.` + "\n" +
							`must be one of ["created", "initializing", "pending-acceptance", "ready", "terminating", "terminated"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"created",
								"initializing",
								"pending-acceptance",
								"ready",
								"terminating",
								"terminated",
							),
						},
					},
					"transit_gateway_attachment_config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Computed:    true,
								Description: `must be "azure-vnet-peering-attachment"`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"azure-vnet-peering-attachment",
									),
								},
							},
							"resource_group_name": schema.StringAttribute{
								Computed:    true,
								Description: `Resource Group Name for the Azure VNET Peering attachment.`,
							},
							"subscription_id": schema.StringAttribute{
								Computed:    true,
								Description: `Subscription ID for the Azure VNET Peering attachment.`,
							},
							"tenant_id": schema.StringAttribute{
								Computed:    true,
								Description: `Tenant ID for the Azure VNET Peering attachment.`,
							},
							"vnet_name": schema.StringAttribute{
								Computed:    true,
								Description: `VNET Name for the Azure VNET Peering attachment.`,
							},
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An RFC-3339 timestamp representation of transit gateway update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("aws_transit_gateway"),
						path.MatchRelative().AtParent().AtName("aws_transit_gateway_response"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway"),
						path.MatchRelative().AtParent().AtName("aws_vpc_peering_gateway_response"),
						path.MatchRelative().AtParent().AtName("azure_transit_gateway"),
					}...),
				},
			},
			"entity_version": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: `Monotonically-increasing version count of the transit gateway, to indicate the order of updates to the` + "\n" +
					`transit gateway.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Human-readable name of the transit gateway.`,
			},
			"network_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `The network to operate on. Requires replacement if changed.`,
			},
		},
	}
}

func (r *CloudGatewayTransitGatewayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CloudGatewayTransitGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CloudGatewayTransitGatewayResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var networkID string
	networkID = data.NetworkID.ValueString()

	createTransitGatewayRequest := *data.ToSharedCreateTransitGatewayRequest()
	request := operations.CreateTransitGatewayRequest{
		NetworkID:                   networkID,
		CreateTransitGatewayRequest: createTransitGatewayRequest,
	}
	res, err := r.client.CloudGateways.CreateTransitGateway(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.TransitGatewayResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedTransitGatewayResponse(ctx, res.TransitGatewayResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayTransitGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CloudGatewayTransitGatewayResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	networkID = data.NetworkID.ValueString()

	var transitGatewayID string
	transitGatewayID = data.ID.ValueString()

	request := operations.GetTransitGatewayRequest{
		NetworkID:        networkID,
		TransitGatewayID: transitGatewayID,
	}
	res, err := r.client.CloudGateways.GetTransitGateway(ctx, request)
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
	if !(res.TransitGatewayResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedTransitGatewayResponse(ctx, res.TransitGatewayResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayTransitGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CloudGatewayTransitGatewayResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayTransitGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CloudGatewayTransitGatewayResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	networkID = data.NetworkID.ValueString()

	var transitGatewayID string
	transitGatewayID = data.ID.ValueString()

	request := operations.DeleteTransitGatewayRequest{
		NetworkID:        networkID,
		TransitGatewayID: transitGatewayID,
	}
	res, err := r.client.CloudGateways.DeleteTransitGateway(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *CloudGatewayTransitGatewayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		NetworkID string `json:"network_id"`
		ID        string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The import ID is not valid. It is expected to be a JSON object string with the format: '{ "network_id": "36ae63d3-efd1-4bec-b246-62aa5d3f5695",  "id": "0850820b-d153-4a2a-b9be-7d2204779139"}': `+err.Error())
		return
	}

	if len(data.NetworkID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field network_id is required but was not found in the json encoded ID. It's expected to be a value alike '"36ae63d3-efd1-4bec-b246-62aa5d3f5695"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), data.NetworkID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"0850820b-d153-4a2a-b9be-7d2204779139"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
