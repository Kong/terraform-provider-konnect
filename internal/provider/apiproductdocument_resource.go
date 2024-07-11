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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &APIProductDocumentResource{}
var _ resource.ResourceWithImportState = &APIProductDocumentResource{}

func NewAPIProductDocumentResource() resource.Resource {
	return &APIProductDocumentResource{}
}

// APIProductDocumentResource defines the resource implementation.
type APIProductDocumentResource struct {
	client *sdk.Konnect
}

// APIProductDocumentResourceModel describes the resource data model.
type APIProductDocumentResourceModel struct {
	APIProductID     types.String      `tfsdk:"api_product_id"`
	Content          types.String      `tfsdk:"content"`
	CreatedAt        types.String      `tfsdk:"created_at"`
	ID               types.String      `tfsdk:"id"`
	Metadata         *tfTypes.Metadata `tfsdk:"metadata"`
	ParentDocumentID types.String      `tfsdk:"parent_document_id"`
	Slug             types.String      `tfsdk:"slug"`
	Status           types.String      `tfsdk:"status"`
	Title            types.String      `tfsdk:"title"`
	UpdatedAt        types.String      `tfsdk:"updated_at"`
}

func (r *APIProductDocumentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_product_document"
}

func (r *APIProductDocumentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "APIProductDocument Resource",
		Attributes: map[string]schema.Attribute{
			"api_product_id": schema.StringAttribute{
				Required:    true,
				Description: `The API product identifier`,
			},
			"content": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Can be markdown string content or base64 encoded string`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The document identifier related to the API product`,
			},
			"metadata": schema.SingleNestedAttribute{
				Computed:    true,
				Optional:    true,
				Attributes:  map[string]schema.Attribute{},
				Description: `metadata of the document`,
			},
			"parent_document_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `parent document id`,
			},
			"slug": schema.StringAttribute{
				Required:    true,
				Description: `document slug. must be unique accross documents belonging to an api product`,
			},
			"status": schema.StringAttribute{
				Required:    true,
				Description: `document publish status. must be one of ["published", "unpublished"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"published",
						"unpublished",
					),
				},
			},
			"title": schema.StringAttribute{
				Required:    true,
				Description: `document title`,
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

func (r *APIProductDocumentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *APIProductDocumentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *APIProductDocumentResourceModel
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
	createAPIProductDocumentDTO := *data.ToSharedCreateAPIProductDocumentDTO()
	request := operations.CreateAPIProductDocumentRequest{
		APIProductID:                apiProductID,
		CreateAPIProductDocumentDTO: createAPIProductDocumentDTO,
	}
	res, err := r.client.APIProductDocumentation.CreateAPIProductDocument(ctx, request)
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
	if !(res.APIProductDocument != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductDocument(res.APIProductDocument)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductDocumentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *APIProductDocumentResourceModel
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
	request := operations.GetAPIProductDocumentRequest{
		APIProductID: apiProductID,
		ID:           id,
	}
	res, err := r.client.APIProductDocumentation.GetAPIProductDocument(ctx, request)
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
	if !(res.APIProductDocument != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductDocument(res.APIProductDocument)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductDocumentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *APIProductDocumentResourceModel
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
	updateAPIProductDocumentDTO := *data.ToSharedUpdateAPIProductDocumentDTO()
	request := operations.UpdateAPIProductDocumentRequest{
		APIProductID:                apiProductID,
		ID:                          id,
		UpdateAPIProductDocumentDTO: updateAPIProductDocumentDTO,
	}
	res, err := r.client.APIProductDocumentation.UpdateAPIProductDocument(ctx, request)
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
	if !(res.APIProductDocument != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductDocument(res.APIProductDocument)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductDocumentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *APIProductDocumentResourceModel
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
	request := operations.DeleteAPIProductDocumentRequest{
		APIProductID: apiProductID,
		ID:           id,
	}
	res, err := r.client.APIProductDocumentation.DeleteAPIProductDocument(ctx, request)
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

func (r *APIProductDocumentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		APIProductID string `json:"api_product_id"`
		ID           string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "api_product_id": "d32d905a-ed33-46a3-a093-d8f536af9a8a",  "id": "de5c9818-be5c-42e6-b514-e3d4bc30ddeb"}': `+err.Error())
		return
	}

	if len(data.APIProductID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field api_product_id is required but was not found in the json encoded ID. It's expected to be a value alike '"d32d905a-ed33-46a3-a093-d8f536af9a8a"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("api_product_id"), data.APIProductID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"de5c9818-be5c-42e6-b514-e3d4bc30ddeb"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
