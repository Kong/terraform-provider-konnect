// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &APIProductVersionResource{}
var _ resource.ResourceWithImportState = &APIProductVersionResource{}

func NewAPIProductVersionResource() resource.Resource {
	return &APIProductVersionResource{}
}

// APIProductVersionResource defines the resource implementation.
type APIProductVersionResource struct {
	client *sdk.Konnect
}

// APIProductVersionResourceModel describes the resource data model.
type APIProductVersionResourceModel struct {
	APIProductID   types.String                      `tfsdk:"api_product_id"`
	CreatedAt      types.String                      `tfsdk:"created_at"`
	Deprecated     types.Bool                        `tfsdk:"deprecated"`
	GatewayService *tfTypes.GatewayServicePayload    `tfsdk:"gateway_service"`
	ID             types.String                      `tfsdk:"id"`
	Name           types.String                      `tfsdk:"name"`
	Portals        []tfTypes.APIProductVersionPortal `tfsdk:"portals"`
	UpdatedAt      types.String                      `tfsdk:"updated_at"`
}

func (r *APIProductVersionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_product_version"
}

func (r *APIProductVersionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "APIProductVersion Resource",
		Attributes: map[string]schema.Attribute{
			"api_product_id": schema.StringAttribute{
				Required:    true,
				Description: `The API Product ID`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"deprecated": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Indicates if the version of the API product is deprecated. Applies deprecation or removes deprecation from all related portal product versions. This field is deprecated: Use [PortalProductVersion.deprecated](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.`,
			},
			"gateway_service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"control_plane_id": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The identifier of the control plane that the gateway service resides in. Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
					"id": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The identifier of a gateway service associated with the version of the API product. Not Null`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
						},
					},
					"runtime_group_id": schema.StringAttribute{
						Computed:    true,
						Description: `This field is deprecated, please use ` + "`" + `control_plane_id` + "`" + ` instead. The identifier of the control plane that the gateway service resides in`,
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The API product version identifier`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The version name of the API product version.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"portals": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"application_registration_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"auth_strategies": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"auto_approve_registration": schema.BoolAttribute{
							Computed: true,
						},
						"deprecated": schema.BoolAttribute{
							Computed: true,
						},
						"portal_id": schema.StringAttribute{
							Computed: true,
						},
						"portal_name": schema.StringAttribute{
							Computed: true,
						},
						"portal_product_version_id": schema.StringAttribute{
							Computed: true,
						},
						"publish_status": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["published", "unpublished"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"published",
									"unpublished",
								),
							},
						},
					},
				},
				Description: `The list of portals which this API product version is configured for`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *APIProductVersionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *APIProductVersionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *APIProductVersionResourceModel
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

	apiProductID := data.APIProductID.ValueString()
	createAPIProductVersionDTO := *data.ToSharedCreateAPIProductVersionDTO()
	request := operations.CreateAPIProductVersionRequest{
		APIProductID:               apiProductID,
		CreateAPIProductVersionDTO: createAPIProductVersionDTO,
	}
	res, err := r.client.APIProductVersions.CreateAPIProductVersion(ctx, request)
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
	if !(res.APIProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersion(res.APIProductVersion)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *APIProductVersionResourceModel
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

	apiProductID := data.APIProductID.ValueString()
	id := data.ID.ValueString()
	request := operations.GetAPIProductVersionRequest{
		APIProductID: apiProductID,
		ID:           id,
	}
	res, err := r.client.APIProductVersions.GetAPIProductVersion(ctx, request)
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
	if !(res.APIProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersion(res.APIProductVersion)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *APIProductVersionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	apiProductID := data.APIProductID.ValueString()
	id := data.ID.ValueString()
	updateAPIProductVersionDTO := *data.ToSharedUpdateAPIProductVersionDTO()
	request := operations.UpdateAPIProductVersionRequest{
		APIProductID:               apiProductID,
		ID:                         id,
		UpdateAPIProductVersionDTO: updateAPIProductVersionDTO,
	}
	res, err := r.client.APIProductVersions.UpdateAPIProductVersion(ctx, request)
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
	if !(res.APIProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersion(res.APIProductVersion)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	apiProductId1 := data.APIProductID.ValueString()
	id1 := data.ID.ValueString()
	request1 := operations.GetAPIProductVersionRequest{
		APIProductID: apiProductId1,
		ID:           id1,
	}
	res1, err := r.client.APIProductVersions.GetAPIProductVersion(ctx, request1)
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
	if !(res1.APIProductVersion != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersion(res1.APIProductVersion)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductVersionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *APIProductVersionResourceModel
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

	apiProductID := data.APIProductID.ValueString()
	id := data.ID.ValueString()
	request := operations.DeleteAPIProductVersionRequest{
		APIProductID: apiProductID,
		ID:           id,
	}
	res, err := r.client.APIProductVersions.DeleteAPIProductVersion(ctx, request)
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

func (r *APIProductVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		APIProductID string `json:"api_product_id"`
		ID           string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "api_product_id": "d32d905a-ed33-46a3-a093-d8f536af9a8a",  "id": "9f5061ce-78f6-4452-9108-ad7c02821fd5"}': `+err.Error())
		return
	}

	if len(data.APIProductID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field api_product_id is required but was not found in the json encoded ID. It's expected to be a value alike '"d32d905a-ed33-46a3-a093-d8f536af9a8a"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("api_product_id"), data.APIProductID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"9f5061ce-78f6-4452-9108-ad7c02821fd5"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
