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
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PortalAuthDataSource{}
var _ datasource.DataSourceWithConfigure = &PortalAuthDataSource{}

func NewPortalAuthDataSource() datasource.DataSource {
	return &PortalAuthDataSource{}
}

// PortalAuthDataSource is the data source implementation.
type PortalAuthDataSource struct {
	client *sdk.Konnect
}

// PortalAuthDataSourceModel describes the data model.
type PortalAuthDataSourceModel struct {
	BasicAuthEnabled       types.Bool                `tfsdk:"basic_auth_enabled"`
	IdpMappingEnabled      types.Bool                `tfsdk:"idp_mapping_enabled"`
	KonnectMappingEnabled  types.Bool                `tfsdk:"konnect_mapping_enabled"`
	OidcAuthEnabled        types.Bool                `tfsdk:"oidc_auth_enabled"`
	OidcConfig             *tfTypes.PortalOIDCConfig `tfsdk:"oidc_config"`
	OidcTeamMappingEnabled types.Bool                `tfsdk:"oidc_team_mapping_enabled"`
	PortalID               types.String              `tfsdk:"portal_id"`
	SamlAuthEnabled        types.Bool                `tfsdk:"saml_auth_enabled"`
}

// Metadata returns the data source type name.
func (r *PortalAuthDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_auth"
}

// Schema defines the schema for the data source.
func (r *PortalAuthDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PortalAuth DataSource",

		Attributes: map[string]schema.Attribute{
			"basic_auth_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `The portal has basic auth enabled or disabled.`,
			},
			"idp_mapping_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `IdP groups determine the Portal Teams a developer has. This will soon replace oidc_team_mapping_enabled.`,
			},
			"konnect_mapping_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `A Konnect Identity Admin assigns teams to a developer.`,
			},
			"oidc_auth_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `The portal has OIDC enabled or disabled.`,
			},
			"oidc_config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"claim_mappings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"email": schema.StringAttribute{
								Computed: true,
							},
							"groups": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `Mappings from a portal developer atribute to an Identity Provider claim.`,
					},
					"client_id": schema.StringAttribute{
						Computed: true,
					},
					"issuer": schema.StringAttribute{
						Computed: true,
					},
					"scopes": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
				Description: `Configuration properties for an OpenID Connect Identity Provider.`,
			},
			"oidc_team_mapping_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `IdP groups determine the Portal Teams a developer has.`,
			},
			"portal_id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the portal.`,
			},
			"saml_auth_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `The portal has SAML enabled or disabled.`,
			},
		},
	}
}

func (r *PortalAuthDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PortalAuthDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PortalAuthDataSourceModel
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

	var portalID string
	portalID = data.PortalID.ValueString()

	request := operations.GetPortalAuthenticationSettingsRequest{
		PortalID: portalID,
	}
	res, err := r.client.PortalAuthSettings.GetPortalAuthenticationSettings(ctx, request)
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
	if !(res.PortalAuthenticationSettingsResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPortalAuthenticationSettingsResponse(ctx, res.PortalAuthenticationSettingsResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
