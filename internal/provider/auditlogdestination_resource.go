// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuditLogDestinationResource{}
var _ resource.ResourceWithImportState = &AuditLogDestinationResource{}

func NewAuditLogDestinationResource() resource.Resource {
	return &AuditLogDestinationResource{}
}

// AuditLogDestinationResource defines the resource implementation.
type AuditLogDestinationResource struct {
	client *sdk.Konnect
}

// AuditLogDestinationResourceModel describes the resource data model.
type AuditLogDestinationResourceModel struct {
	Authorization       types.String `tfsdk:"authorization"`
	CreatedAt           types.String `tfsdk:"created_at"`
	Endpoint            types.String `tfsdk:"endpoint"`
	ID                  types.String `tfsdk:"id"`
	LogFormat           types.String `tfsdk:"log_format"`
	Name                types.String `tfsdk:"name"`
	SkipSslVerification types.Bool   `tfsdk:"skip_ssl_verification"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
}

func (r *AuditLogDestinationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_audit_log_destination"
}

func (r *AuditLogDestinationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AuditLogDestination Resource",
		Attributes: map[string]schema.Attribute{
			"authorization": schema.StringAttribute{
				Optional:    true,
				Description: `The value to include in the ` + "`" + `Authorization` + "`" + ` header when sending audit logs to the webhook.`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `Timestamp when this webhook was created.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"endpoint": schema.StringAttribute{
				Required:    true,
				Description: `The endpoint that will receive audit log messages.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The unique ID of the audit log destination.`,
			},
			"log_format": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(`cef`),
				Description: `The output format of each log messages. Default: "cef"; must be one of ["cef", "json", "cps"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"cef",
						"json",
						"cps",
					),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The name of the audit log destination.`,
			},
			"skip_ssl_verification": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				MarkdownDescription: `Indicates if the SSL certificate verification of the host endpoint should be skipped when delivering payloads.` + "\n" +
					`We strongly recommend not setting this to 'true' as you are subject to man-in-the-middle and other attacks.` + "\n" +
					`This option should be considered only for self-signed SSL certificates used in a non-production environment.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `Timestamp when this webhook was last updated. Initial value is 0001-01-01T00:00:0Z.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *AuditLogDestinationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AuditLogDestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AuditLogDestinationResourceModel
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

	request, requestDiags := data.ToSharedCreateAuditLogDestination(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.AuditLogs.CreateAuditLogDestination(ctx, request)
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
	if !(res.AuditLogDestination != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAuditLogDestination(ctx, res.AuditLogDestination)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetAuditLogDestinationRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.AuditLogs.GetAuditLogDestination(ctx, *request1)
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
	if !(res1.AuditLogDestination != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAuditLogDestination(ctx, res1.AuditLogDestination)...)

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

func (r *AuditLogDestinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AuditLogDestinationResourceModel
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

	request, requestDiags := data.ToOperationsGetAuditLogDestinationRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.AuditLogs.GetAuditLogDestination(ctx, *request)
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
	if !(res.AuditLogDestination != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAuditLogDestination(ctx, res.AuditLogDestination)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditLogDestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AuditLogDestinationResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsUpdateAuditLogDestinationRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.AuditLogs.UpdateAuditLogDestination(ctx, *request)
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
	if !(res.AuditLogDestination != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAuditLogDestination(ctx, res.AuditLogDestination)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetAuditLogDestinationRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.AuditLogs.GetAuditLogDestination(ctx, *request1)
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
	if !(res1.AuditLogDestination != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedAuditLogDestination(ctx, res1.AuditLogDestination)...)

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

func (r *AuditLogDestinationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AuditLogDestinationResourceModel
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

	request, requestDiags := data.ToOperationsDeleteAuditLogDestinationRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.AuditLogs.DeleteAuditLogDestination(ctx, *request)
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

func (r *AuditLogDestinationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
