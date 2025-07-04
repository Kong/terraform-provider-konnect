// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PortalProductVersionResource{}
var _ resource.ResourceWithImportState = &PortalProductVersionResource{}

func NewPortalProductVersionResource() resource.Resource {
	return &PortalProductVersionResource{}
}

// PortalProductVersionResource defines the resource implementation.
type PortalProductVersionResource struct {
	client *sdk.Konnect
}

// PortalProductVersionResourceModel describes the resource data model.
type PortalProductVersionResourceModel struct {
	ApplicationRegistrationEnabled types.Bool             `tfsdk:"application_registration_enabled"`
	AuthStrategies                 []tfTypes.AuthStrategy `tfsdk:"auth_strategies"`
	AuthStrategyIds                []types.String         `tfsdk:"auth_strategy_ids"`
	AutoApproveRegistration        types.Bool             `tfsdk:"auto_approve_registration"`
	CreatedAt                      types.String           `tfsdk:"created_at"`
	Deprecated                     types.Bool             `tfsdk:"deprecated"`
	ID                             types.String           `tfsdk:"id"`
	NotifyDevelopers               types.Bool             `tfsdk:"notify_developers"`
	PortalID                       types.String           `tfsdk:"portal_id"`
	ProductVersionID               types.String           `tfsdk:"product_version_id"`
	PublishStatus                  types.String           `tfsdk:"publish_status"`
	UpdatedAt                      types.String           `tfsdk:"updated_at"`
}

func (r *PortalProductVersionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_product_version"
}

func (r *PortalProductVersionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PortalProductVersion Resource",
		Attributes: map[string]schema.Attribute{
			"application_registration_enabled": schema.BoolAttribute{
				Required:    true,
				Description: `Whether the application registration on this portal for the api product version is enabled`,
			},
			"auth_strategies": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_credentials": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"auth_methods": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"available_scopes": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Possible developer selectable scopes for an application. Only present when using DCR Provider that supports it.`,
								},
								"credential_type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["client_credentials", "self_managed_client_credentials"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"client_credentials",
											"self_managed_client_credentials",
										),
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `The Application Auth Strategy ID.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
							Description: `Client Credential Auth strategy that the application uses.`,
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.Expressions{
									path.MatchRelative().AtParent().AtName("key_auth"),
								}...),
							},
						},
						"key_auth": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"credential_type": schema.StringAttribute{
									Computed:    true,
									Description: `must be "key_auth"`,
									Validators: []validator.String{
										stringvalidator.OneOf("key_auth"),
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `The Application Auth Strategy ID.`,
								},
								"key_names": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
							Description: `KeyAuth Auth strategy that the application uses.`,
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.Expressions{
									path.MatchRelative().AtParent().AtName("client_credentials"),
								}...),
							},
						},
					},
				},
				Description: `A list of authentication strategies`,
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
			},
			"auth_strategy_ids": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: `A list of authentication strategy IDs`,
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
			},
			"auto_approve_registration": schema.BoolAttribute{
				Required:    true,
				Description: `Whether the application registration auto approval on this portal for the api product version is enabled`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"deprecated": schema.BoolAttribute{
				Required:    true,
				Description: `Whether the api product version on the portal is deprecated`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Contains a unique identifier used for this resource.`,
			},
			"notify_developers": schema.BoolAttribute{
				Optional:    true,
				Description: `Whether to notify developers who are affected by this change`,
			},
			"portal_id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the portal.`,
			},
			"product_version_id": schema.StringAttribute{
				Required:    true,
				Description: `API product version identifier`,
			},
			"publish_status": schema.StringAttribute{
				Required:    true,
				Description: `Publication status of the API product version on the portal. must be one of ["published", "unpublished"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"published",
						"unpublished",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `An ISO-8601 timestamp representation of entity update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *PortalProductVersionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PortalProductVersionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PortalProductVersionResourceModel
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

	request, requestDiags := data.ToOperationsReplacePortalProductVersionRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.PortalProductVersions.ReplacePortalProductVersion(ctx, *request)
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
	if !(res.PortalProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPortalProductVersion(ctx, res.PortalProductVersion)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetPortalProductVersionRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.PortalProductVersions.GetPortalProductVersion(ctx, *request1)
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
	if !(res1.PortalProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPortalProductVersion(ctx, res1.PortalProductVersion)...)

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

func (r *PortalProductVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PortalProductVersionResourceModel
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

	request, requestDiags := data.ToOperationsGetPortalProductVersionRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.PortalProductVersions.GetPortalProductVersion(ctx, *request)
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
	if !(res.PortalProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPortalProductVersion(ctx, res.PortalProductVersion)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalProductVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PortalProductVersionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsReplacePortalProductVersionRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.PortalProductVersions.ReplacePortalProductVersion(ctx, *request)
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
	if !(res.PortalProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPortalProductVersion(ctx, res.PortalProductVersion)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetPortalProductVersionRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.PortalProductVersions.GetPortalProductVersion(ctx, *request1)
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
	if !(res1.PortalProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedPortalProductVersion(ctx, res1.PortalProductVersion)...)

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

func (r *PortalProductVersionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PortalProductVersionResourceModel
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

	request, requestDiags := data.ToOperationsDeletePortalProductVersionRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.PortalProductVersions.DeletePortalProductVersion(ctx, *request)
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

func (r *PortalProductVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		PortalID         string `json:"portal_id"`
		ProductVersionID string `json:"product_version_id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The import ID is not valid. It is expected to be a JSON object string with the format: '{ "portal_id": "",  "product_version_id": "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"}': `+err.Error())
		return
	}

	if len(data.PortalID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field portal_id is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("portal_id"), data.PortalID)...)
	if len(data.ProductVersionID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field product_version_id is required but was not found in the json encoded ID. It's expected to be a value alike '"5f9fd312-a987-4628-b4c5-bb4f4fddd5f7"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("product_version_id"), data.ProductVersionID)...)

}
