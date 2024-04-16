// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_int64planmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/int64planmodifier"
	speakeasy_listplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/listplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayTargetResource{}
var _ resource.ResourceWithImportState = &GatewayTargetResource{}

func NewGatewayTargetResource() resource.Resource {
	return &GatewayTargetResource{}
}

// GatewayTargetResource defines the resource implementation.
type GatewayTargetResource struct {
	client *sdk.Konnect
}

// GatewayTargetResourceModel describes the resource data model.
type GatewayTargetResourceModel struct {
	ControlPlaneID types.String         `tfsdk:"control_plane_id"`
	CreatedAt      types.Number         `tfsdk:"created_at"`
	ID             types.String         `tfsdk:"id"`
	Tags           []types.String       `tfsdk:"tags"`
	Target         types.String         `tfsdk:"target"`
	Upstream       *tfTypes.ACLConsumer `tfsdk:"upstream"`
	UpstreamID     types.String         `tfsdk:"upstream_id"`
	Weight         types.Int64          `tfsdk:"weight"`
}

func (r *GatewayTargetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_target"
}

func (r *GatewayTargetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayTarget Resource",
		Attributes: map[string]schema.Attribute{
			"control_plane_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed. `,
			},
			"created_at": schema.NumberAttribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the Target to lookup`,
			},
			"tags": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Target for grouping and filtering. Requires replacement if changed. `,
			},
			"target": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The target address (ip or hostname) and port. If the hostname resolves to an SRV record, the ` + "`" + `port` + "`" + ` value will be overridden by the value from the DNS record. Requires replacement if changed. `,
			},
			"upstream": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"upstream_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `ID or target of the Target to lookup. Requires replacement if changed. `,
			},
			"weight": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
					speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: `The weight this target gets within the upstream loadbalancer (` + "`" + `0` + "`" + `-` + "`" + `65535` + "`" + `). If the hostname resolves to an SRV record, the ` + "`" + `weight` + "`" + ` value will be overridden by the value from the DNS record. Requires replacement if changed. ; Default: 100`,
			},
		},
	}
}

func (r *GatewayTargetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayTargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayTargetResourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	upstreamIDForTarget := data.UpstreamID.ValueString()
	createTargetWithoutParents := *data.ToSharedCreateTargetWithoutParents()
	request := operations.CreateTargetWithUpstreamRequest{
		ControlPlaneID:             controlPlaneID,
		UpstreamIDForTarget:        upstreamIDForTarget,
		CreateTargetWithoutParents: createTargetWithoutParents,
	}
	res, err := r.client.Targets.CreateTargetWithUpstream(ctx, request)
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
	if res.Target == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTarget(res.Target)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayTargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayTargetResourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	upstreamIDForTarget := data.UpstreamID.ValueString()
	targetID := data.ID.ValueString()
	request := operations.GetTargetWithUpstreamRequest{
		ControlPlaneID:      controlPlaneID,
		UpstreamIDForTarget: upstreamIDForTarget,
		TargetID:            targetID,
	}
	res, err := r.client.Targets.GetTargetWithUpstream(ctx, request)
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
	if res.Target == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTarget(res.Target)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayTargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayTargetResourceModel
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

func (r *GatewayTargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayTargetResourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	upstreamIDForTarget := data.UpstreamID.ValueString()
	targetID := data.ID.ValueString()
	request := operations.DeleteTargetWithUpstreamRequest{
		ControlPlaneID:      controlPlaneID,
		UpstreamIDForTarget: upstreamIDForTarget,
		TargetID:            targetID,
	}
	res, err := r.client.Targets.DeleteTargetWithUpstream(ctx, request)
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

func (r *GatewayTargetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_target. Reason: composite imports strings not supported.")
}
