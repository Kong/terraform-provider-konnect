// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/stringplanmodifier"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemAccountRoleResource{}
var _ resource.ResourceWithImportState = &SystemAccountRoleResource{}

func NewSystemAccountRoleResource() resource.Resource {
	return &SystemAccountRoleResource{}
}

// SystemAccountRoleResource defines the resource implementation.
type SystemAccountRoleResource struct {
	client *sdk.Konnect
}

// SystemAccountRoleResourceModel describes the resource data model.
type SystemAccountRoleResourceModel struct {
	AccountID      types.String `tfsdk:"account_id"`
	EntityID       types.String `tfsdk:"entity_id"`
	EntityRegion   types.String `tfsdk:"entity_region"`
	EntityTypeName types.String `tfsdk:"entity_type_name"`
	ID             types.String `tfsdk:"id"`
	RoleName       types.String `tfsdk:"role_name"`
}

func (r *SystemAccountRoleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_account_role"
}

func (r *SystemAccountRoleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SystemAccountRole Resource",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `ID of the system account. Requires replacement if changed.`,
			},
			"entity_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The ID of the entity. Requires replacement if changed.`,
			},
			"entity_region": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The region of the team. must be one of ["us", "eu", "au", "*"]; Requires replacement if changed.`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"us",
						"eu",
						"au",
						"*",
					),
				},
			},
			"entity_type_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The type of entity. must be one of ["API Products", "Application Auth Strategies", "Audit Logs", "Control Planes", "DCR Providers", "Identity", "Mesh Control Planes", "Networks", "Portals", "Service Hub"]; Requires replacement if changed.`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"API Products",
						"Application Auth Strategies",
						"Audit Logs",
						"Control Planes",
						"DCR Providers",
						"Identity",
						"Mesh Control Planes",
						"Networks",
						"Portals",
						"Service Hub",
					),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the role assignment.`,
			},
			"role_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The desired role. must be one of ["Admin", "Appearance Maintainer", "Application Registration", "Certificate Admin", "Cloud Gateway Cluster Admin", "Cloud Gateway Cluster Viewer", "Consumer Admin", "Creator", "Deployer", "Discovery Admin", "Discovery Viewer", "Gateway Service Admin", "Integration Admin", "Integration Viewer", "Key Admin", "Maintainer", "Network Admin", "Network Creator", "Network Viewer", "Plugin Admin", "Plugins Admin", "Product Publisher", "Publisher", "Route Admin", "SNI Admin", "Service Admin", "Service Creator", "Service Viewer", "Upstream Admin", "Vault Admin", "Viewer"]; Requires replacement if changed.`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Admin",
						"Appearance Maintainer",
						"Application Registration",
						"Certificate Admin",
						"Cloud Gateway Cluster Admin",
						"Cloud Gateway Cluster Viewer",
						"Consumer Admin",
						"Creator",
						"Deployer",
						"Discovery Admin",
						"Discovery Viewer",
						"Gateway Service Admin",
						"Integration Admin",
						"Integration Viewer",
						"Key Admin",
						"Maintainer",
						"Network Admin",
						"Network Creator",
						"Network Viewer",
						"Plugin Admin",
						"Plugins Admin",
						"Product Publisher",
						"Publisher",
						"Route Admin",
						"SNI Admin",
						"Service Admin",
						"Service Creator",
						"Service Viewer",
						"Upstream Admin",
						"Vault Admin",
						"Viewer",
					),
				},
			},
		},
	}
}

func (r *SystemAccountRoleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SystemAccountRoleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SystemAccountRoleResourceModel
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

	var accountID string
	accountID = data.AccountID.ValueString()

	assignRole := data.ToSharedAssignRole()
	request := operations.PostSystemAccountsAccountIDAssignedRolesRequest{
		AccountID:  accountID,
		AssignRole: assignRole,
	}
	res, err := r.client.SystemAccountsRoles.PostSystemAccountsAccountIDAssignedRoles(ctx, request)
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
	if !(res.AssignedRole != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAssignedRole(res.AssignedRole)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemAccountRoleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SystemAccountRoleResourceModel
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemAccountRoleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SystemAccountRoleResourceModel
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

func (r *SystemAccountRoleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SystemAccountRoleResourceModel
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

	var accountID string
	accountID = data.AccountID.ValueString()

	var roleID string
	roleID = data.ID.ValueString()

	request := operations.DeleteSystemAccountsAccountIDAssignedRolesRoleIDRequest{
		AccountID: accountID,
		RoleID:    roleID,
	}
	res, err := r.client.SystemAccountsRoles.DeleteSystemAccountsAccountIDAssignedRolesRoleID(ctx, request)
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

func (r *SystemAccountRoleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource system_account_role.")
}
