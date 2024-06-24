// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &APIProductSpecificationResource{}
var _ resource.ResourceWithImportState = &APIProductSpecificationResource{}

func NewAPIProductSpecificationResource() resource.Resource {
	return &APIProductSpecificationResource{}
}

// APIProductSpecificationResource defines the resource implementation.
type APIProductSpecificationResource struct {
	client *sdk.Konnect
}

// APIProductSpecificationResourceModel describes the resource data model.
type APIProductSpecificationResourceModel struct {
	APIProductID        types.String `tfsdk:"api_product_id"`
	APIProductVersionID types.String `tfsdk:"api_product_version_id"`
	Content             types.String `tfsdk:"content"`
	CreatedAt           types.String `tfsdk:"created_at"`
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
}

func (r *APIProductSpecificationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_product_specification"
}

func (r *APIProductSpecificationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "APIProductSpecification Resource",
		Attributes: map[string]schema.Attribute{
			"api_product_id": schema.StringAttribute{
				Required:    true,
				Description: `The API product identifier`,
			},
			"api_product_version_id": schema.StringAttribute{
				Required:    true,
				Description: `The API product version identifier`,
			},
			"content": schema.StringAttribute{
				Required:    true,
				Description: `The base64 encoded contents of the API product version specification`,
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
				Description: `The API product version specification identifier`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The name of the API product version specification`,
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

func (r *APIProductSpecificationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *APIProductSpecificationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *APIProductSpecificationResourceModel
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
	apiProductVersionID := data.APIProductVersionID.ValueString()
	createAPIProductVersionSpecDTO := *data.ToSharedCreateAPIProductVersionSpecDTO()
	request := operations.CreateAPIProductVersionSpecRequest{
		APIProductID:                   apiProductID,
		APIProductVersionID:            apiProductVersionID,
		CreateAPIProductVersionSpecDTO: createAPIProductVersionSpecDTO,
	}
	res, err := r.client.APIProductVersionSpecification.CreateAPIProductVersionSpec(ctx, request)
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
	if !(res.APIProductVersionSpec != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersionSpec(res.APIProductVersionSpec)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductSpecificationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *APIProductSpecificationResourceModel
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
	apiProductVersionID := data.APIProductVersionID.ValueString()
	specificationID := data.ID.ValueString()
	request := operations.GetAPIProductVersionSpecRequest{
		APIProductID:        apiProductID,
		APIProductVersionID: apiProductVersionID,
		SpecificationID:     specificationID,
	}
	res, err := r.client.APIProductVersionSpecification.GetAPIProductVersionSpec(ctx, request)
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
	if !(res.APIProductVersionSpec != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersionSpec(res.APIProductVersionSpec)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductSpecificationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *APIProductSpecificationResourceModel
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
	apiProductVersionID := data.APIProductVersionID.ValueString()
	specificationID := data.ID.ValueString()
	updateAPIProductVersionSpecDTO := *data.ToSharedUpdateAPIProductVersionSpecDTO()
	request := operations.UpdateAPIProductVersionSpecRequest{
		APIProductID:                   apiProductID,
		APIProductVersionID:            apiProductVersionID,
		SpecificationID:                specificationID,
		UpdateAPIProductVersionSpecDTO: updateAPIProductVersionSpecDTO,
	}
	res, err := r.client.APIProductVersionSpecification.UpdateAPIProductVersionSpec(ctx, request)
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
	if !(res.APIProductVersionSpec != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAPIProductVersionSpec(res.APIProductVersionSpec)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIProductSpecificationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *APIProductSpecificationResourceModel
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
	apiProductVersionID := data.APIProductVersionID.ValueString()
	specificationID := data.ID.ValueString()
	request := operations.DeleteAPIProductVersionSpecRequest{
		APIProductID:        apiProductID,
		APIProductVersionID: apiProductVersionID,
		SpecificationID:     specificationID,
	}
	res, err := r.client.APIProductVersionSpecification.DeleteAPIProductVersionSpec(ctx, request)
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

func (r *APIProductSpecificationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		APIProductID        string `json:"api_product_id"`
		APIProductVersionID string `json:"api_product_version_id"`
		ID                  string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "api_product_id": "d32d905a-ed33-46a3-a093-d8f536af9a8a",  "api_product_version_id": "9f5061ce-78f6-4452-9108-ad7c02821fd5",  "id": "742ff9f1-fb89-4aeb-a599-f0e278c7aeaa"}': `+err.Error())
		return
	}

	if len(data.APIProductID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field api_product_id is required but was not found in the json encoded ID. It's expected to be a value alike '"d32d905a-ed33-46a3-a093-d8f536af9a8a"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("api_product_id"), data.APIProductID)...)
	if len(data.APIProductVersionID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field api_product_version_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9f5061ce-78f6-4452-9108-ad7c02821fd5"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("api_product_version_id"), data.APIProductVersionID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"742ff9f1-fb89-4aeb-a599-f0e278c7aeaa"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
