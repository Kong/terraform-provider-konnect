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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshOPAResource{}
var _ resource.ResourceWithImportState = &MeshOPAResource{}

func NewMeshOPAResource() resource.Resource {
	return &MeshOPAResource{}
}

// MeshOPAResource defines the resource implementation.
type MeshOPAResource struct {
	client *sdk.Konnect
}

// MeshOPAResourceModel describes the resource data model.
type MeshOPAResourceModel struct {
	CpID             types.String            `tfsdk:"cp_id"`
	CreationTime     types.String            `tfsdk:"creation_time"`
	Labels           map[string]types.String `tfsdk:"labels"`
	Mesh             types.String            `tfsdk:"mesh"`
	ModificationTime types.String            `tfsdk:"modification_time"`
	Name             types.String            `tfsdk:"name"`
	Spec             tfTypes.MeshOPAItemSpec `tfsdk:"spec"`
	Type             types.String            `tfsdk:"type"`
	Warnings         []types.String          `tfsdk:"warnings"`
}

func (r *MeshOPAResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_opa"
}

func (r *MeshOPAResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshOPA Resource",
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
				Description: `name of the MeshOPA`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"agent_config": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"inline": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Data source is inline bytes.`,
									},
									"inline_string": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Data source is inline string` + "`" + ``,
									},
									"secret": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Data source is a secret with given Secret key.`,
									},
								},
								Description: `AgentConfig defines bootstrap OPA agent configuration.`,
							},
							"append_policies": schema.ListNestedAttribute{
								Computed: true,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
									Attributes: map[string]schema.Attribute{
										"ignore_decision": schema.BoolAttribute{
											Computed:    true,
											Optional:    true,
											Description: `If true, then policy won't be taken into account when making a decision.`,
										},
										"rego": schema.SingleNestedAttribute{
											Computed: true,
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"inline": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Description: `Data source is inline bytes.`,
												},
												"inline_string": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Description: `Data source is inline string` + "`" + ``,
												},
												"secret": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Description: `Data source is a secret with given Secret key.`,
												},
											},
											Description: `OPA Policy written in Rego. Available values: secret, inline, inlineString. Not Null`,
											Validators: []validator.Object{
												speakeasy_objectvalidators.NotNull(),
											},
										},
									},
								},
								Description: `Policies define OPA policies that will be applied on OPA Agent.`,
							},
							"auth_config": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"on_agent_failure": schema.StringAttribute{
										Computed: true,
										Optional: true,
										MarkdownDescription: `OnAgentFailure either 'allow' or 'deny' (default to deny) whether` + "\n" +
											`to allow requests when the authorization agent failed.` + "\n" +
											`must be one of ["Allow", "Deny"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Allow",
												"Deny",
											),
										},
									},
									"request_body": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"max_size": schema.Int64Attribute{
												Computed: true,
												Optional: true,
												MarkdownDescription: `MaxSize defines the maximum payload size sent to authorization agent. If the payload` + "\n" +
													`is larger it will be truncated and there will be a header` + "\n" +
													`` + "`" + `x-envoy-auth-partial-body: true` + "`" + `. If it is set to 0 no body will be` + "\n" +
													`sent to the agent.`,
											},
											"send_raw_body": schema.BoolAttribute{
												Computed:    true,
												Optional:    true,
												Description: `SendRawBody enable sending raw body instead of the body encoded into UTF-8`,
											},
										},
										MarkdownDescription: `RequestBody configuration to apply on the request body sent to the` + "\n" +
											`authorization agent (if absent, the body is not sent).`,
									},
									"status_on_error": schema.Int64Attribute{
										Computed: true,
										Optional: true,
										MarkdownDescription: `StatusOnError is the http status to return when there's a connection` + "\n" +
											`failure between the dataplane and the authorization agent`,
									},
									"timeout": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `Timeout for the single gRPC request from Envoy to OPA Agent.`,
									},
								},
								Description: `AuthConfig are configurations specific to the filter.`,
							},
						},
					},
					"target_ref": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
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
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
									`Name and Namespace can be used.`,
							},
							"mesh": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Mesh is reserved for future use to identify cross mesh resources.`,
							},
							"name": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
									`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
							},
							"namespace": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
									`will be targeted.`,
							},
							"proxy_types": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
									`all data plane types are targeted by the policy.`,
								Validators: []validator.List{
									listvalidator.SizeAtLeast(1),
								},
							},
							"section_name": schema.StringAttribute{
								Computed: true,
								Optional: true,
								MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
									`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
							},
							"tags": schema.MapAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
									`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
							},
						},
						MarkdownDescription: `TargetRef is a reference to the resource the policy takes an effect on.` + "\n" +
							`The resource could be either a real store object or virtual resource` + "\n" +
							`defined inplace.`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshOPA resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "MeshOPA"`,
				Validators: []validator.String{
					stringvalidator.OneOf("MeshOPA"),
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

func (r *MeshOPAResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshOPAResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshOPAResourceModel
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

	meshOPAItem := *data.ToSharedMeshOPAItem()
	request := operations.PutMeshOPARequest{
		CpID:        cpID,
		Mesh:        mesh,
		Name:        name,
		MeshOPAItem: meshOPAItem,
	}
	res, err := r.client.MeshOPA.PutMeshOPA(ctx, request)
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
	if !(res.MeshOPACreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshOPACreateOrUpdateSuccessResponse(res.MeshOPACreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshOPARequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshOPA.GetMeshOPA(ctx, request1)
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
	if !(res1.MeshOPAItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshOPAItem(res1.MeshOPAItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshOPAResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshOPAResourceModel
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

	request := operations.GetMeshOPARequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshOPA.GetMeshOPA(ctx, request)
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
	if !(res.MeshOPAItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshOPAItem(res.MeshOPAItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshOPAResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshOPAResourceModel
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

	meshOPAItem := *data.ToSharedMeshOPAItem()
	request := operations.PutMeshOPARequest{
		CpID:        cpID,
		Mesh:        mesh,
		Name:        name,
		MeshOPAItem: meshOPAItem,
	}
	res, err := r.client.MeshOPA.PutMeshOPA(ctx, request)
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
	if !(res.MeshOPACreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshOPACreateOrUpdateSuccessResponse(res.MeshOPACreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshOPARequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshOPA.GetMeshOPA(ctx, request1)
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
	if !(res1.MeshOPAItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshOPAItem(res1.MeshOPAItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshOPAResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshOPAResourceModel
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

	request := operations.DeleteMeshOPARequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshOPA.DeleteMeshOPA(ctx, request)
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

func (r *MeshOPAResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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