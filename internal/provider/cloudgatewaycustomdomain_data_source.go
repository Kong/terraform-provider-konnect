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
var _ datasource.DataSource = &CloudGatewayCustomDomainDataSource{}
var _ datasource.DataSourceWithConfigure = &CloudGatewayCustomDomainDataSource{}

func NewCloudGatewayCustomDomainDataSource() datasource.DataSource {
	return &CloudGatewayCustomDomainDataSource{}
}

// CloudGatewayCustomDomainDataSource is the data source implementation.
type CloudGatewayCustomDomainDataSource struct {
	client *sdk.Konnect
}

// CloudGatewayCustomDomainDataSourceModel describes the data model.
type CloudGatewayCustomDomainDataSourceModel struct {
	CertificateID   types.String                      `tfsdk:"certificate_id"`
	ControlPlaneGeo types.String                      `tfsdk:"control_plane_geo"`
	ControlPlaneID  types.String                      `tfsdk:"control_plane_id"`
	CreatedAt       types.String                      `tfsdk:"created_at"`
	Domain          types.String                      `tfsdk:"domain"`
	EntityVersion   types.Int64                       `tfsdk:"entity_version"`
	ID              types.String                      `tfsdk:"id"`
	SniID           types.String                      `tfsdk:"sni_id"`
	State           types.String                      `tfsdk:"state"`
	StateMetadata   tfTypes.CustomDomainStateMetadata `tfsdk:"state_metadata"`
	UpdatedAt       types.String                      `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *CloudGatewayCustomDomainDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_custom_domain"
}

// Schema defines the schema for the data source.
func (r *CloudGatewayCustomDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayCustomDomain DataSource",

		Attributes: map[string]schema.Attribute{
			"certificate_id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `Certificate ID for the certificate representing this domain and stored on data-planes for this` + "\n" +
					`control-plane. Can be retrieved via the control-planes API for this custom domain's control-plane.` + "\n" +
					``,
			},
			"control_plane_geo": schema.StringAttribute{
				Computed:    true,
				Description: `Set of control-plane geos supported for deploying cloud-gateways configurations. must be one of ["us", "eu", "au"]`,
			},
			"control_plane_id": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of custom domain creation date.`,
			},
			"domain": schema.StringAttribute{
				Computed:    true,
				Description: `Domain name of the custom domain.`,
			},
			"entity_version": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: `Monotonically-increasing version count of the custom domain, to indicate the order of updates to the custom` + "\n" +
					`domain.` + "\n" +
					``,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the custom domain to operate on.`,
			},
			"sni_id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `Server Name Indication ID for this domain and stored on data-planes for this control-plane. Can be retrieved` + "\n" +
					`via the control-planes API for this custom domain's control-plane.` + "\n" +
					``,
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: `State of the custom domain. must be one of ["created", "initializing", "ready", "terminating", "terminated", "error"]`,
			},
			"state_metadata": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"reason": schema.StringAttribute{
						Computed: true,
						MarkdownDescription: `Reason why the custom domain may be in an erroneous state, reported from backing infrastructure.` + "\n" +
							``,
					},
					"reported_status": schema.StringAttribute{
						Computed:    true,
						Description: `Reported status of the custom domain from backing infrastructure.`,
					},
				},
				MarkdownDescription: `Metadata describing the backing state of the custom domain and why it may be in an erroneous state.` + "\n" +
					``,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of custom domain update date.`,
			},
		},
	}
}

func (r *CloudGatewayCustomDomainDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *CloudGatewayCustomDomainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CloudGatewayCustomDomainDataSourceModel
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

	customDomainID := data.ID.ValueString()
	request := operations.GetCustomDomainRequest{
		CustomDomainID: customDomainID,
	}
	res, err := r.client.CustomDomains.GetCustomDomain(ctx, request)
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
	if res.CustomDomain == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCustomDomain(res.CustomDomain)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
