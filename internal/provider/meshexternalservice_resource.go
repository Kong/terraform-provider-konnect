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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	custom_listplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/listplanmodifier"
	speakeasy_listplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/listplanmodifier"
	speakeasy_objectplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/objectplanmodifier"
	custom_stringplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/stringplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/v2/internal/planmodifiers/stringplanmodifier"
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
var _ resource.Resource = &MeshExternalServiceResource{}
var _ resource.ResourceWithImportState = &MeshExternalServiceResource{}

func NewMeshExternalServiceResource() resource.Resource {
	return &MeshExternalServiceResource{}
}

// MeshExternalServiceResource defines the resource implementation.
type MeshExternalServiceResource struct {
	client *sdk.Konnect
}

// MeshExternalServiceResourceModel describes the resource data model.
type MeshExternalServiceResourceModel struct {
	CpID             types.String                        `tfsdk:"cp_id"`
	CreationTime     types.String                        `tfsdk:"creation_time"`
	Labels           map[string]types.String             `tfsdk:"labels"`
	Mesh             types.String                        `tfsdk:"mesh"`
	ModificationTime types.String                        `tfsdk:"modification_time"`
	Name             types.String                        `tfsdk:"name"`
	Spec             tfTypes.MeshExternalServiceItemSpec `tfsdk:"spec"`
	Status           *tfTypes.Status                     `tfsdk:"status"`
	Type             types.String                        `tfsdk:"type"`
	Warnings         []types.String                      `tfsdk:"warnings"`
}

func (r *MeshExternalServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_external_service"
}

func (r *MeshExternalServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshExternalService Resource",
		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
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
				Required:    true,
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
				Required:    true,
				Description: `name of the MeshExternalService`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"endpoints": schema.ListNestedAttribute{
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
								"address": schema.StringAttribute{
									Optional:    true,
									Description: `Address defines an address to which a user want to send a request. Is possible to provide ` + "`" + `domain` + "`" + `, ` + "`" + `ip` + "`" + `. Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
										stringvalidator.UTF8LengthAtLeast(1),
									},
								},
								"port": schema.Int64Attribute{
									Optional:    true,
									Description: `Port of the endpoint. Not Null`,
									Validators: []validator.Int64{
										speakeasy_int64validators.NotNull(),
										int64validator.Between(1, 65535),
									},
								},
							},
						},
						Description: `Endpoints defines a list of destinations to send traffic to.`,
					},
					"extension": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"config": schema.StringAttribute{
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									custom_stringplanmodifier.ArbitraryJSONModifier(),
								},
								Description: `Config freeform configuration for the extension. Parsed as JSON.`,
								Validators: []validator.String{
									validators.IsValidJSON(),
								},
							},
							"type": schema.StringAttribute{
								Required:    true,
								Description: `Type of the extension.`,
							},
						},
						Description: `Extension struct for a plugin configuration, in the presence of an extension ` + "`" + `endpoints` + "`" + ` and ` + "`" + `tls` + "`" + ` are not required anymore - it's up to the extension to validate them independently.`,
					},
					"match": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"port": schema.Int64Attribute{
								Required:    true,
								Description: `Port defines a port to which a user does request.`,
								Validators: []validator.Int64{
									int64validator.Between(1, 65535),
								},
							},
							"protocol": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString("tcp"),
								Description: `Protocol defines a protocol of the communication. Possible values: ` + "`" + `tcp` + "`" + `, ` + "`" + `grpc` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `http2` + "`" + `. Default: "tcp"; must be one of ["tcp", "grpc", "http", "http2"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"tcp",
										"grpc",
										"http",
										"http2",
									),
								},
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString("HostnameGenerator"),
								Description: `Type of the match, only ` + "`" + `HostnameGenerator` + "`" + ` is available at the moment. Default: "HostnameGenerator"; must be "HostnameGenerator"`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"HostnameGenerator",
									),
								},
							},
						},
						Description: `Match defines traffic that should be routed through the sidecar.`,
					},
					"tls": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"allow_renegotiation": schema.BoolAttribute{
								Computed: true,
								Optional: true,
								Default:  booldefault.StaticBool(false),
								MarkdownDescription: `AllowRenegotiation defines if TLS sessions will allow renegotiation.` + "\n" +
									`Setting this to true is not recommended for security reasons.` + "\n" +
									`Default: false`,
							},
							"enabled": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
								Description: `Enabled defines if proxy should originate TLS. Default: false`,
							},
							"verification": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"inline": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is inline bytes.`,
											},
											"inline_string": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is inline string` + "`" + ``,
											},
											"secret": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is a secret with given Secret key.`,
											},
										},
										Description: `CaCert defines a certificate of CA.`,
									},
									"client_cert": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"inline": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is inline bytes.`,
											},
											"inline_string": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is inline string` + "`" + ``,
											},
											"secret": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is a secret with given Secret key.`,
											},
										},
										Description: `ClientCert defines a certificate of a client.`,
									},
									"client_key": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"inline": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is inline bytes.`,
											},
											"inline_string": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is inline string` + "`" + ``,
											},
											"secret": schema.StringAttribute{
												Optional:    true,
												Description: `Data source is a secret with given Secret key.`,
											},
										},
										Description: `ClientKey defines a client private key.`,
									},
									"mode": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("Secured"),
										Description: `Mode defines if proxy should skip verification, one of ` + "`" + `SkipSAN` + "`" + `, ` + "`" + `SkipCA` + "`" + `, ` + "`" + `Secured` + "`" + `, ` + "`" + `SkipAll` + "`" + `. Default ` + "`" + `Secured` + "`" + `. Default: "Secured"; must be one of ["SkipSAN", "SkipCA", "Secured", "SkipAll"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SkipSAN",
												"SkipCA",
												"Secured",
												"SkipAll",
											),
										},
									},
									"server_name": schema.StringAttribute{
										Optional:    true,
										Description: `ServerName overrides the default Server Name Indicator set by Kuma.`,
									},
									"subject_alt_names": schema.ListNestedAttribute{
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
												"type": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Default:     stringdefault.StaticString("Exact"),
													Description: `Type specifies matching type, one of ` + "`" + `Exact` + "`" + `, ` + "`" + `Prefix` + "`" + `. Default: ` + "`" + `Exact` + "`" + `. Default: "Exact"; must be one of ["Exact", "Prefix"]`,
													Validators: []validator.String{
														stringvalidator.OneOf(
															"Exact",
															"Prefix",
														),
													},
												},
												"value": schema.StringAttribute{
													Optional:    true,
													Description: `Value to match. Not Null`,
													Validators: []validator.String{
														speakeasy_stringvalidators.NotNull(),
													},
												},
											},
										},
										Description: `SubjectAltNames list of names to verify in the certificate.`,
									},
								},
								Description: `Verification section for providing TLS verification details.`,
							},
							"version": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"max": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("TLSAuto"),
										Description: `Max defines maximum supported version. One of ` + "`" + `TLSAuto` + "`" + `, ` + "`" + `TLS10` + "`" + `, ` + "`" + `TLS11` + "`" + `, ` + "`" + `TLS12` + "`" + `, ` + "`" + `TLS13` + "`" + `. Default: "TLSAuto"; must be one of ["TLSAuto", "TLS10", "TLS11", "TLS12", "TLS13"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"TLSAuto",
												"TLS10",
												"TLS11",
												"TLS12",
												"TLS13",
											),
										},
									},
									"min": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("TLSAuto"),
										Description: `Min defines minimum supported version. One of ` + "`" + `TLSAuto` + "`" + `, ` + "`" + `TLS10` + "`" + `, ` + "`" + `TLS11` + "`" + `, ` + "`" + `TLS12` + "`" + `, ` + "`" + `TLS13` + "`" + `. Default: "TLSAuto"; must be one of ["TLSAuto", "TLS10", "TLS11", "TLS12", "TLS13"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"TLSAuto",
												"TLS10",
												"TLS11",
												"TLS12",
												"TLS13",
											),
										},
									},
								},
								Description: `Version section for providing version specification.`,
							},
						},
						Description: `Tls provides a TLS configuration when proxy is resposible for a TLS origination`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshExternalService resource.`,
			},
			"status": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"addresses": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Computed: true,
								},
								"hostname_generator_ref": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"core_name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"origin": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Description: `Addresses section for generated domains`,
					},
					"hostname_generators": schema.ListNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
							speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
						},
						NestedObject: schema.NestedAttributeObject{
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"conditions": schema.ListNestedAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.List{
										custom_listplanmodifier.SupressZeroNullModifier(),
										speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
									},
									NestedObject: schema.NestedAttributeObject{
										PlanModifiers: []planmodifier.Object{
											speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
										},
										Attributes: map[string]schema.Attribute{
											"message": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `message is a human readable message indicating details about the transition.` + "\n" +
													`This may be an empty string.`,
												Validators: []validator.String{
													stringvalidator.UTF8LengthAtMost(32768),
												},
											},
											"reason": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `reason contains a programmatic identifier indicating the reason for the condition's last transition.` + "\n" +
													`Producers of specific condition types may define expected values and meanings for this field,` + "\n" +
													`and whether the values are considered a guaranteed API.` + "\n" +
													`The value should be a CamelCase string.` + "\n" +
													`This field may not be empty.`,
												Validators: []validator.String{
													stringvalidator.UTF8LengthBetween(1, 1024),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`), "must match pattern "+regexp.MustCompile(`^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$`).String()),
												},
											},
											"status": schema.StringAttribute{
												Computed: true,
												PlanModifiers: []planmodifier.String{
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
												},
												Description: `status of the condition, one of True, False, Unknown. must be one of ["True", "False", "Unknown"]`,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"True",
														"False",
														"Unknown",
													),
												},
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Description: `type of condition in CamelCase or in foo.example.com/CamelCase.`,
												Validators: []validator.String{
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
									Attributes: map[string]schema.Attribute{
										"core_name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"vip": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"ip": schema.StringAttribute{
								Computed:    true,
								Description: `Value allocated IP for a provided domain with ` + "`" + `HostnameGenerator` + "`" + ` type in a match section.`,
							},
						},
						Description: `Vip section for allocated IP`,
					},
				},
				Description: `Status is the current status of the Kuma MeshExternalService resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "MeshExternalService"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MeshExternalService",
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

func (r *MeshExternalServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshExternalServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshExternalServiceResourceModel
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

	meshExternalServiceItem := *data.ToSharedMeshExternalServiceItemInput()
	request := operations.CreateMeshExternalServiceRequest{
		CpID:                    cpID,
		Mesh:                    mesh,
		Name:                    name,
		MeshExternalServiceItem: meshExternalServiceItem,
	}
	res, err := r.client.MeshExternalService.CreateMeshExternalService(ctx, request)
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
	if !(res.MeshExternalServiceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshExternalServiceCreateOrUpdateSuccessResponse(res.MeshExternalServiceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshExternalServiceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshExternalService.GetMeshExternalService(ctx, request1)
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
	if !(res1.MeshExternalServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshExternalServiceItem(res1.MeshExternalServiceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshExternalServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshExternalServiceResourceModel
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

	request := operations.GetMeshExternalServiceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshExternalService.GetMeshExternalService(ctx, request)
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
	if !(res.MeshExternalServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshExternalServiceItem(res.MeshExternalServiceItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshExternalServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshExternalServiceResourceModel
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

	meshExternalServiceItem := *data.ToSharedMeshExternalServiceItemInput()
	request := operations.UpdateMeshExternalServiceRequest{
		CpID:                    cpID,
		Mesh:                    mesh,
		Name:                    name,
		MeshExternalServiceItem: meshExternalServiceItem,
	}
	res, err := r.client.MeshExternalService.UpdateMeshExternalService(ctx, request)
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
	if !(res.MeshExternalServiceCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshExternalServiceCreateOrUpdateSuccessResponse(res.MeshExternalServiceCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var cpId1 string
	cpId1 = data.CpID.ValueString()

	var mesh1 string
	mesh1 = data.Mesh.ValueString()

	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetMeshExternalServiceRequest{
		CpID: cpId1,
		Mesh: mesh1,
		Name: name1,
	}
	res1, err := r.client.MeshExternalService.GetMeshExternalService(ctx, request1)
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
	if !(res1.MeshExternalServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedMeshExternalServiceItem(res1.MeshExternalServiceItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshExternalServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshExternalServiceResourceModel
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

	request := operations.DeleteMeshExternalServiceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshExternalService.DeleteMeshExternalService(ctx, request)
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

func (r *MeshExternalServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
