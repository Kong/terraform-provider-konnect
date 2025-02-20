// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/stringvalidators"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DcrProviderResource{}
var _ resource.ResourceWithImportState = &DcrProviderResource{}

func NewDcrProviderResource() resource.Resource {
	return &DcrProviderResource{}
}

// DcrProviderResource defines the resource implementation.
type DcrProviderResource struct {
	client *sdk.Konnect
}

// DcrProviderResourceModel describes the resource data model.
type DcrProviderResourceModel struct {
	Active      types.Bool       `tfsdk:"active"`
	Auth0       *tfTypes.Auth0   `tfsdk:"auth0" tfPlanOnly:"true"`
	AzureAd     *tfTypes.AzureAd `tfsdk:"azure_ad" tfPlanOnly:"true"`
	Curity      *tfTypes.AzureAd `tfsdk:"curity" tfPlanOnly:"true"`
	DisplayName types.String     `tfsdk:"display_name"`
	HTTP        *tfTypes.HTTP    `tfsdk:"http" tfPlanOnly:"true"`
	ID          types.String     `tfsdk:"id"`
	Issuer      types.String     `tfsdk:"issuer"`
	Name        types.String     `tfsdk:"name"`
	Okta        *tfTypes.Okta    `tfsdk:"okta" tfPlanOnly:"true"`
}

func (r *DcrProviderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dcr_provider"
}

func (r *DcrProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DcrProvider Resource",
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: `At least one active auth strategy is using this DCR provider.`,
			},
			"auth0": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Computed:    true,
						Description: `At least one active auth strategy is using this DCR provider.`,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dcr_config": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"initial_client_audience": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This is the audience value used for the initial client.` + "\n" +
									`If using a custom domain on Auth0, this must be set as to the Auth0 Management API audience value.` + "\n" +
									`If left blank, the issuer will be used instead.`,
								Validators: []validator.String{
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
							"initial_client_id": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This ID should be copied from your identity provider's settings after you create a client` + "\n" +
									`and assign it as the management client for DCR for this developer portal` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
							"initial_client_secret": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This secret should be copied from your identity provider's settings after you create a client` + "\n" +
									`and assign it as the management client for DCR for this developer portal` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
							"use_developer_managed_scopes": schema.BoolAttribute{
								Computed: true,
								Optional: true,
							},
						},
						Description: `Payload to create an Auth0 DCR provider. Not Null`,
						Validators: []validator.Object{
							speakeasy_objectvalidators.NotNull(),
						},
					},
					"display_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.`,
						Validators: []validator.String{
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `Contains a unique identifier used by the API for this resource.`,
					},
					"issuer": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthAtMost(256),
						},
					},
					"labels": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
							`` + "\n" +
							`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI. Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"provider_type": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Not Null; must be "auth0"; Requires replacement if changed.`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf("auth0"),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `Payload to update an Auth0 DCR provider.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("azure_ad"),
						path.MatchRelative().AtParent().AtName("curity"),
						path.MatchRelative().AtParent().AtName("http"),
						path.MatchRelative().AtParent().AtName("okta"),
					}...),
				},
			},
			"azure_ad": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Computed:    true,
						Description: `At least one active auth strategy is using this DCR provider.`,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dcr_config": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"initial_client_id": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This ID should be copied from your identity provider's settings after you create a client` + "\n" +
									`and assign it as the management client for DCR for this developer portal` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
							"initial_client_secret": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This secret should be copied from your identity provider's settings after you create a client` + "\n" +
									`and assign it as the management client for DCR for this developer portal` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
						},
						Description: `Payload to create an Azure AD DCR provider. Not Null`,
						Validators: []validator.Object{
							speakeasy_objectvalidators.NotNull(),
						},
					},
					"display_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.`,
						Validators: []validator.String{
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `Contains a unique identifier used by the API for this resource.`,
					},
					"issuer": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthAtMost(256),
						},
					},
					"labels": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
							`` + "\n" +
							`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI. Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"provider_type": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Not Null; must be "azureAd"; Requires replacement if changed.`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf("azureAd"),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `Payload to update an Azure AD DCR provider.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("auth0"),
						path.MatchRelative().AtParent().AtName("curity"),
						path.MatchRelative().AtParent().AtName("http"),
						path.MatchRelative().AtParent().AtName("okta"),
					}...),
				},
			},
			"curity": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Computed:    true,
						Description: `At least one active auth strategy is using this DCR provider.`,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dcr_config": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"initial_client_id": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This ID should be copied from your identity provider's settings after you create a client` + "\n" +
									`and assign it as the management client for DCR for this developer portal` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
							"initial_client_secret": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This secret should be copied from your identity provider's settings after you create a client` + "\n" +
									`and assign it as the management client for DCR for this developer portal` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
						},
						Description: `Payload to create a Curity DCR provider. Not Null`,
						Validators: []validator.Object{
							speakeasy_objectvalidators.NotNull(),
						},
					},
					"display_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.`,
						Validators: []validator.String{
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `Contains a unique identifier used by the API for this resource.`,
					},
					"issuer": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthAtMost(256),
						},
					},
					"labels": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
							`` + "\n" +
							`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI. Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"provider_type": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Not Null; must be "curity"; Requires replacement if changed.`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf("curity"),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `Payload to update a Curity DCR provider.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("auth0"),
						path.MatchRelative().AtParent().AtName("azure_ad"),
						path.MatchRelative().AtParent().AtName("http"),
						path.MatchRelative().AtParent().AtName("okta"),
					}...),
				},
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Description: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthBetween(1, 256),
				},
			},
			"http": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Computed:    true,
						Description: `At least one active auth strategy is using this DCR provider.`,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dcr_config": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"api_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This is the API Key that will be sent with each HTTP request to the custom DCR server. It can be` + "\n" +
									`verified on the server to ensure that incoming requests are coming from Konnect.` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthBetween(12, 256),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_]+$`), "must match pattern "+regexp.MustCompile(`^[a-zA-Z0-9_]+$`).String()),
								},
							},
							"dcr_base_url": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `The base URL of the DCR server. This is the URL that will be used to make the HTTP requests from Konnect to the DCR provider.` + "\n" +
									`This URL must be accessible from the Konnect service.` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
							"disable_event_hooks": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `This flag disables all the event-hooks on the application flow for the DCR provider.`,
							},
							"disable_refresh_secret": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `This flag disable the refresh-secret endpoint on the application flow for the DCR provider.`,
							},
						},
						Description: `Payload to create an HTTP DCR provider. Not Null`,
						Validators: []validator.Object{
							speakeasy_objectvalidators.NotNull(),
						},
					},
					"display_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.`,
						Validators: []validator.String{
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `Contains a unique identifier used by the API for this resource.`,
					},
					"issuer": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthAtMost(256),
						},
					},
					"labels": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
							`` + "\n" +
							`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI. Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"provider_type": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Not Null; must be "http"; Requires replacement if changed.`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf("http"),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `Payload to update an HTTP DCR provider.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("auth0"),
						path.MatchRelative().AtParent().AtName("azure_ad"),
						path.MatchRelative().AtParent().AtName("curity"),
						path.MatchRelative().AtParent().AtName("okta"),
					}...),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Contains a unique identifier used by the API for this resource.`,
			},
			"issuer": schema.StringAttribute{
				Computed:    true,
				Description: `The issuer of the DCR provider.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(256),
				},
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthBetween(1, 256),
				},
			},
			"okta": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Computed:    true,
						Description: `At least one active auth strategy is using this DCR provider.`,
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity creation date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"dcr_config": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"dcr_token": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `This secret should be copied from your identity provider's settings after you create a client` + "\n" +
									`and assign it as the management client for DCR for this developer portal` + "\n" +
									`Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.UTF8LengthAtMost(256),
								},
							},
						},
						Description: `Payload to create an Okta DCR provider. Not Null`,
						Validators: []validator.Object{
							speakeasy_objectvalidators.NotNull(),
						},
					},
					"display_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.`,
						Validators: []validator.String{
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Description: `Contains a unique identifier used by the API for this resource.`,
					},
					"issuer": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthAtMost(256),
						},
					},
					"labels": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
							`` + "\n" +
							`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The name of the DCR provider. This is used to identify the DCR provider in the Konnect UI. Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.UTF8LengthBetween(1, 256),
						},
					},
					"provider_type": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `Not Null; must be "okta"; Requires replacement if changed.`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf("okta"),
						},
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: `An ISO-8601 timestamp representation of entity update date.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
				},
				Description: `Payload to update an Okta DCR provider.`,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.Expressions{
						path.MatchRelative().AtParent().AtName("auth0"),
						path.MatchRelative().AtParent().AtName("azure_ad"),
						path.MatchRelative().AtParent().AtName("curity"),
						path.MatchRelative().AtParent().AtName("http"),
					}...),
				},
			},
		},
	}
}

func (r *DcrProviderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DcrProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DcrProviderResourceModel
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

	request := *data.ToSharedCreateDcrProviderRequest()
	res, err := r.client.DCRProviders.CreateDcrProvider(ctx, request)
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
	if !(res.CreateDcrProviderResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCreateDcrProviderResponse(res.CreateDcrProviderResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var dcrProviderID string
	dcrProviderID = data.ID.ValueString()

	request1 := operations.GetDcrProviderRequest{
		DcrProviderID: dcrProviderID,
	}
	res1, err := r.client.DCRProviders.GetDcrProvider(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.DcrProviderResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDcrProviderResponse(res1.DcrProviderResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DcrProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DcrProviderResourceModel
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

	var dcrProviderID string
	dcrProviderID = data.ID.ValueString()

	request := operations.GetDcrProviderRequest{
		DcrProviderID: dcrProviderID,
	}
	res, err := r.client.DCRProviders.GetDcrProvider(ctx, request)
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
	if !(res.DcrProviderResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDcrProviderResponse(res.DcrProviderResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DcrProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DcrProviderResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var dcrProviderID string
	dcrProviderID = data.ID.ValueString()

	updateDcrProviderRequest := *data.ToSharedUpdateDcrProviderRequest()
	request := operations.UpdateDcrProviderRequest{
		DcrProviderID:            dcrProviderID,
		UpdateDcrProviderRequest: updateDcrProviderRequest,
	}
	res, err := r.client.DCRProviders.UpdateDcrProvider(ctx, request)
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
	if !(res.DcrProviderResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDcrProviderResponse(res.DcrProviderResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var dcrProviderId1 string
	dcrProviderId1 = data.ID.ValueString()

	request1 := operations.GetDcrProviderRequest{
		DcrProviderID: dcrProviderId1,
	}
	res1, err := r.client.DCRProviders.GetDcrProvider(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.DcrProviderResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDcrProviderResponse(res1.DcrProviderResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DcrProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DcrProviderResourceModel
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

	var dcrProviderID string
	dcrProviderID = data.ID.ValueString()

	request := operations.DeleteDcrProviderRequest{
		DcrProviderID: dcrProviderID,
	}
	res, err := r.client.DCRProviders.DeleteDcrProvider(ctx, request)
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

func (r *DcrProviderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
