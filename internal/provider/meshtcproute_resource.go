// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	custom_listplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/listplanmodifier"
	speakeasy_listplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/listplanmodifier"
	custom_stringplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/stringplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
	speakeasy_listvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/listvalidators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshTCPRouteResource{}
var _ resource.ResourceWithImportState = &MeshTCPRouteResource{}

func NewMeshTCPRouteResource() resource.Resource {
	return &MeshTCPRouteResource{}
}

// MeshTCPRouteResource defines the resource implementation.
type MeshTCPRouteResource struct {
	client *sdk.Konnect
}

// MeshTCPRouteResourceModel describes the resource data model.
type MeshTCPRouteResourceModel struct {
	CpID             types.String                 `tfsdk:"cp_id"`
	CreationTime     types.String                 `tfsdk:"creation_time"`
	Labels           map[string]types.String      `tfsdk:"labels"`
	Mesh             types.String                 `tfsdk:"mesh"`
	ModificationTime types.String                 `tfsdk:"modification_time"`
	Name             types.String                 `tfsdk:"name"`
	Spec             tfTypes.MeshTCPRouteItemSpec `tfsdk:"spec"`
	Type             types.String                 `tfsdk:"type"`
	Warnings         []types.String               `tfsdk:"warnings"`
}

func (r *MeshTCPRouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_tcp_route"
}

func (r *MeshTCPRouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshTCPRoute Resource",
		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `Id of the Konnect resource`,
			},
			"creation_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Time at which the resource was created`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"labels": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: `The labels to help identity resources`,
			},
			"mesh": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `name of the mesh`,
			},
			"modification_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Time at which the resource was updated`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					custom_stringplanmodifier.RequiresReplaceModifier(),
				},
				Description: `name of the MeshTCPRoute`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"target_ref": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Required:    true,
								Description: `Kind of the referenced resource. must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Mesh",
										"MeshSubset",
										"MeshGateway",
										"MeshService",
										"MeshExternalService",
										"MeshMultiZoneService",
										"MeshServiceSubset",
										"MeshHTTPRoute",
										"Dataplane",
									),
								},
							},
							"labels": schema.MapAttribute{
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
									`Name and Namespace can be used.`,
							},
							"mesh": schema.StringAttribute{
								Optional:    true,
								Description: `Mesh is reserved for future use to identify cross mesh resources.`,
							},
							"name": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
									`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
							},
							"namespace": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
									`will be targeted.`,
							},
							"proxy_types": schema.ListAttribute{
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.List{
									custom_listplanmodifier.SupressZeroNullModifier(),
								},
								ElementType: types.StringType,
								MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
									`all data plane types are targeted by the policy.`,
							},
							"section_name": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
									`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
							},
							"tags": schema.MapAttribute{
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
									`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
							},
						},
						MarkdownDescription: `TargetRef is a reference to the resource the policy takes an effect on.` + "\n" +
							`The resource could be either a real store object or virtual resource` + "\n" +
							`defined in-place.`,
					},
					"to": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
						},
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"rules": schema.ListNestedAttribute{
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.List{
										custom_listplanmodifier.SupressZeroNullModifier(),
									},
									NestedObject: schema.NestedAttributeObject{
										Validators: []validator.Object{
											speakeasy_objectvalidators.NotNull(),
										},
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"backend_refs": schema.ListNestedAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.List{
															custom_listplanmodifier.SupressZeroNullModifier(),
														},
														NestedObject: schema.NestedAttributeObject{
															Validators: []validator.Object{
																speakeasy_objectvalidators.NotNull(),
															},
															Attributes: map[string]schema.Attribute{
																"kind": schema.StringAttribute{
																	Optional:    true,
																	Description: `Kind of the referenced resource. Not Null; must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]`,
																	Validators: []validator.String{
																		speakeasy_stringvalidators.NotNull(),
																		stringvalidator.OneOf(
																			"Mesh",
																			"MeshSubset",
																			"MeshGateway",
																			"MeshService",
																			"MeshExternalService",
																			"MeshMultiZoneService",
																			"MeshServiceSubset",
																			"MeshHTTPRoute",
																			"Dataplane",
																		),
																	},
																},
																"labels": schema.MapAttribute{
																	Optional:    true,
																	ElementType: types.StringType,
																	MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
																		`Name and Namespace can be used.`,
																},
																"mesh": schema.StringAttribute{
																	Optional:    true,
																	Description: `Mesh is reserved for future use to identify cross mesh resources.`,
																},
																"name": schema.StringAttribute{
																	Optional: true,
																	MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
																		`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
																},
																"namespace": schema.StringAttribute{
																	Optional: true,
																	MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
																		`will be targeted.`,
																},
																"port": schema.Int32Attribute{
																	Optional:    true,
																	Description: `Port is only supported when this ref refers to a real MeshService object`,
																},
																"proxy_types": schema.ListAttribute{
																	Computed: true,
																	Optional: true,
																	PlanModifiers: []planmodifier.List{
																		custom_listplanmodifier.SupressZeroNullModifier(),
																	},
																	ElementType: types.StringType,
																	MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
																		`all data plane types are targeted by the policy.`,
																},
																"section_name": schema.StringAttribute{
																	Optional: true,
																	MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
																		`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
																},
																"tags": schema.MapAttribute{
																	Optional:    true,
																	ElementType: types.StringType,
																	MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
																		`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
																},
																"weight": schema.Int64Attribute{
																	Computed:    true,
																	Optional:    true,
																	Default:     int64default.StaticInt64(1),
																	Description: `Default: 1`,
																},
															},
														},
													},
												},
												MarkdownDescription: `Default holds routing rules that can be merged with rules from other` + "\n" +
													`policies.` + "\n" +
													`Not Null`,
												Validators: []validator.Object{
													speakeasy_objectvalidators.NotNull(),
												},
											},
										},
									},
									MarkdownDescription: `Rules contains the routing rules applies to a combination of top-level` + "\n" +
										`targetRef and the targetRef in this entry.` + "\n" +
										`Not Null`,
									Validators: []validator.List{
										speakeasy_listvalidators.NotNull(),
										listvalidator.SizeAtMost(1),
									},
								},
								"target_ref": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Optional:    true,
											Description: `Kind of the referenced resource. Not Null; must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
												stringvalidator.OneOf(
													"Mesh",
													"MeshSubset",
													"MeshGateway",
													"MeshService",
													"MeshExternalService",
													"MeshMultiZoneService",
													"MeshServiceSubset",
													"MeshHTTPRoute",
													"Dataplane",
												),
											},
										},
										"labels": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
												`Name and Namespace can be used.`,
										},
										"mesh": schema.StringAttribute{
											Optional:    true,
											Description: `Mesh is reserved for future use to identify cross mesh resources.`,
										},
										"name": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
												`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
										},
										"namespace": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
												`will be targeted.`,
										},
										"proxy_types": schema.ListAttribute{
											Computed: true,
											Optional: true,
											PlanModifiers: []planmodifier.List{
												custom_listplanmodifier.SupressZeroNullModifier(),
											},
											ElementType: types.StringType,
											MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
												`all data plane types are targeted by the policy.`,
										},
										"section_name": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
												`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
										},
										"tags": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
												`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
										},
									},
									MarkdownDescription: `TargetRef is a reference to the resource that represents a group of` + "\n" +
										`destinations.` + "\n" +
										`Not Null`,
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
								},
							},
						},
						MarkdownDescription: `To list makes a match between the consumed services and corresponding` + "\n" +
							`configurations`,
						Validators: []validator.List{
							listvalidator.SizeAtLeast(1),
						},
					},
				},
				Description: `Spec is the specification of the Kuma MeshTCPRoute resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "MeshTCPRoute"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MeshTCPRoute",
					),
				},
			},
			"warnings": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					custom_listplanmodifier.SupressZeroNullModifier(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				ElementType: types.StringType,
				MarkdownDescription: `warnings is a list of warning messages to return to the requesting Kuma API clients.` + "\n" +
					`Warning messages describe a problem the client making the API request should correct or be aware of.`,
			},
		},
	}
}

func (r *MeshTCPRouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshTCPRouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshTCPRouteResourceModel
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

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	meshTCPRouteItem := *data.ToSharedMeshTCPRouteItemInput()
	request := operations.CreateMeshTCPRouteRequest{
		CpID:             cpID,
		Mesh:             mesh,
		Name:             name,
		MeshTCPRouteItem: meshTCPRouteItem,
	}
	res, err := r.client.MeshTCPRoute.CreateMeshTCPRoute(ctx, request)
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
	if !(res.MeshTCPRouteCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTCPRouteCreateOrUpdateSuccessResponse(res.MeshTCPRouteCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshTCPRouteRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshTCPRoute.GetMeshTCPRoute(ctx, request1)
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
	if !(res1.MeshTCPRouteItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTCPRouteItem(res1.MeshTCPRouteItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshTCPRouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshTCPRouteResourceModel
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

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	request := operations.GetMeshTCPRouteRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshTCPRoute.GetMeshTCPRoute(ctx, request)
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
	if !(res.MeshTCPRouteItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTCPRouteItem(res.MeshTCPRouteItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshTCPRouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshTCPRouteResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	meshTCPRouteItem := *data.ToSharedMeshTCPRouteItemInput()
	request := operations.UpdateMeshTCPRouteRequest{
		CpID:             cpID,
		Mesh:             mesh,
		Name:             name,
		MeshTCPRouteItem: meshTCPRouteItem,
	}
	res, err := r.client.MeshTCPRoute.UpdateMeshTCPRoute(ctx, request)
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
	if !(res.MeshTCPRouteCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTCPRouteCreateOrUpdateSuccessResponse(res.MeshTCPRouteCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshTCPRouteRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshTCPRoute.GetMeshTCPRoute(ctx, request1)
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
	if !(res1.MeshTCPRouteItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTCPRouteItem(res1.MeshTCPRouteItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshTCPRouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshTCPRouteResourceModel
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

	var cpID string
	cpID = data.CpID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	request := operations.DeleteMeshTCPRouteRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshTCPRoute.DeleteMeshTCPRoute(ctx, request)
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

}

func (r *MeshTCPRouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		CpID string `json:"cp_id"`
		Mesh string `json:"mesh"`
		Name string `json:"name"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "cp_id": "bf138ba2-c9b1-4229-b268-04d9d8a6410b",  "mesh": "",  "name": ""}': `+err.Error())
		return
	}

	if len(data.CpID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field cp_id is required but was not found in the json encoded ID. It's expected to be a value alike '"bf138ba2-c9b1-4229-b268-04d9d8a6410b"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cp_id"), data.CpID)...)
	if len(data.Mesh) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field mesh is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("mesh"), data.Mesh)...)
	if len(data.Name) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field name is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), data.Name)...)

}
