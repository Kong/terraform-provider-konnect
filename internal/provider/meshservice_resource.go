// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
	speakeasy_int64validators "github.com/kong/terraform-provider-konnect/v2/internal/validators/int64validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/stringvalidators"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshServiceResource{}
var _ resource.ResourceWithImportState = &MeshServiceResource{}

func NewMeshServiceResource() resource.Resource {
	return &MeshServiceResource{}
}

// MeshServiceResource defines the resource implementation.
type MeshServiceResource struct {
	client *sdk.Konnect
}

// MeshServiceResourceModel describes the resource data model.
type MeshServiceResourceModel struct {
	CpID             types.String                   `tfsdk:"cp_id"`
	CreationTime     types.String                   `tfsdk:"creation_time"`
	Labels           map[string]types.String        `tfsdk:"labels"`
	Mesh             types.String                   `tfsdk:"mesh"`
	ModificationTime types.String                   `tfsdk:"modification_time"`
	Name             types.String                   `tfsdk:"name"`
	Spec             tfTypes.MeshServiceItemSpec    `tfsdk:"spec"`
	Status           *tfTypes.MeshServiceItemStatus `tfsdk:"status"`
	Type             types.String                   `tfsdk:"type"`
	Warnings         []types.String                 `tfsdk:"warnings"`
}

func (r *MeshServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_service"
}

func (r *MeshServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshService Resource",
		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"creation_time": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Time at which the resource was created`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `The labels to help identity resources`,
			},
			"mesh": schema.StringAttribute{
				Required:    true,
				Description: `name of the mesh`,
			},
			"modification_time": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Time at which the resource was updated`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `name of the MeshService`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"identities": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Not Null; must be "ServiceTag"`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
										stringvalidator.OneOf(
											"ServiceTag",
										),
									},
								},
								"value": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
							},
						},
					},
					"ports": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"app_protocol": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Protocol identifies a protocol supported by a service.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"port": schema.Int64Attribute{
									Computed:    true,
									Optional:    true,
									Description: `Not Null`,
									Validators: []validator.Int64{
										speakeasy_int64validators.NotNull(),
									},
								},
								"target_port": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"integer": schema.Int64Attribute{
											Computed: true,
											Optional: true,
											Validators: []validator.Int64{
												int64validator.ConflictsWith(path.Expressions{
													path.MatchRelative().AtParent().AtName("str"),
												}...),
											},
										},
										"str": schema.StringAttribute{
											Computed: true,
											Optional: true,
											Validators: []validator.String{
												stringvalidator.ConflictsWith(path.Expressions{
													path.MatchRelative().AtParent().AtName("integer"),
												}...),
											},
										},
									},
								},
							},
						},
					},
					"selector": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"dataplane_ref": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
								},
							},
							"dataplane_tags": schema.MapAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"state": schema.StringAttribute{
						Computed: true,
						Optional: true,
						MarkdownDescription: `State of MeshService. Available if there is at least one healthy endpoint. Otherwise, Unavailable.` + "\n" +
							`It's used for cross zone communication to check if we should send traffic to it, when MeshService is aggregated into MeshMultiZoneService.` + "\n" +
							`must be one of ["Available", "Unavailable"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Available",
								"Unavailable",
							),
						},
					},
				},
				Description: `Spec is the specification of the Kuma MeshService resource.`,
			},
			"status": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"addresses": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"hostname_generator_ref": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"core_name": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Not Null`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
											},
										},
									},
								},
								"origin": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
					},
					"dataplane_proxies": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"connected": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Number of data plane proxies connected to the zone control plane`,
							},
							"healthy": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Number of data plane proxies with all healthy inbounds selected by this MeshService.`,
							},
							"total": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Total number of data plane proxies.`,
							},
						},
						Description: `Data plane proxies statistics selected by this MeshService.`,
					},
					"hostname_generators": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"conditions": schema.ListNestedAttribute{
									Computed: true,
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Validators: []validator.Object{
											speakeasy_objectvalidators.NotNull(),
										},
										Attributes: map[string]schema.Attribute{
											"message": schema.StringAttribute{
												Computed: true,
												Optional: true,
												MarkdownDescription: `message is a human readable message indicating details about the transition.` + "\n" +
													`This may be an empty string.` + "\n" +
													`Not Null`,
												Validators: []validator.String{
													speakeasy_stringvalidators.NotNull(),
													stringvalidator.UTF8LengthAtMost(32768),
												},
											},
											"reason": schema.StringAttribute{
												Computed: true,
												Optional: true,
												MarkdownDescription: `reason contains a programmatic identifier indicating the reason for the condition's last transition.` + "\n" +
													`Producers of specific condition types may define expected values and meanings for this field,` + "\n" +
													`and whether the values are considered a guaranteed API.` + "\n" +
													`The value should be a CamelCase string.` + "\n" +
													`This field may not be empty.` + "\n" +
													`Not Null`,
												Validators: []validator.String{
													speakeasy_stringvalidators.NotNull(),
													stringvalidator.UTF8LengthBetween(1, 1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`), "must match pattern "+regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`).String()),
												},
											},
											"status": schema.StringAttribute{
												Computed:    true,
												Optional:    true,
												Description: `status of the condition, one of True, False, Unknown. Not Null; must be one of ["True", "False", "Unknown"]`,
												Validators: []validator.String{
													speakeasy_stringvalidators.NotNull(),
													stringvalidator.OneOf(
														"True",
														"False",
														"Unknown",
													),
												},
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Optional:    true,
												Description: `type of condition in CamelCase or in foo.example.com/CamelCase. Not Null`,
												Validators: []validator.String{
													speakeasy_stringvalidators.NotNull(),
													stringvalidator.UTF8LengthAtMost(316),
													stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), "must match pattern "+regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`).String()),
												},
											},
										},
									},
									Description: `Conditions is an array of hostname generator conditions.`,
								},
								"hostname_generator_ref": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"core_name": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Not Null`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
											},
										},
									},
									Description: `Not Null`,
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
								},
							},
						},
					},
					"tls": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `must be one of ["Ready", "NotReady"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Ready",
										"NotReady",
									),
								},
							},
						},
					},
					"vips": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
					},
				},
				Description: `Status is the current status of the Kuma MeshService resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "MeshService"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MeshService",
					),
				},
			},
			"warnings": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				MarkdownDescription: `warnings is a list of warning messages to return to the requesting Kuma API clients.` + "\n" +
					`Warning messages describe a problem the client making the API request should correct or be aware of.`,
			},
		},
	}
}

func (r *MeshServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshServiceResourceModel
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

	meshServiceItem := *data.ToSharedMeshServiceItem()
	request := operations.PutMeshServiceRequest{
		CpID:            cpID,
		Mesh:            mesh,
		Name:            name,
		MeshServiceItem: meshServiceItem,
	}
	res, err := r.client.MeshService.PutMeshService(ctx, request)
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
	if !(res.MeshServiceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshServiceCreateOrUpdateSuccessResponse(res.MeshServiceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshServiceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshService.GetMeshService(ctx, request1)
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
	if !(res1.MeshServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshServiceItem(res1.MeshServiceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshServiceResourceModel
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

	request := operations.GetMeshServiceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshService.GetMeshService(ctx, request)
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
	if !(res.MeshServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshServiceItem(res.MeshServiceItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshServiceResourceModel
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

	meshServiceItem := *data.ToSharedMeshServiceItem()
	request := operations.PutMeshServiceRequest{
		CpID:            cpID,
		Mesh:            mesh,
		Name:            name,
		MeshServiceItem: meshServiceItem,
	}
	res, err := r.client.MeshService.PutMeshService(ctx, request)
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
	if !(res.MeshServiceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshServiceCreateOrUpdateSuccessResponse(res.MeshServiceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshServiceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshService.GetMeshService(ctx, request1)
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
	if !(res1.MeshServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshServiceItem(res1.MeshServiceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshServiceResourceModel
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

	request := operations.DeleteMeshServiceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshService.DeleteMeshService(ctx, request)
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

func (r *MeshServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
