// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MeshOPADataSource{}
var _ datasource.DataSourceWithConfigure = &MeshOPADataSource{}

func NewMeshOPADataSource() datasource.DataSource {
	return &MeshOPADataSource{}
}

// MeshOPADataSource is the data source implementation.
type MeshOPADataSource struct {
	client *sdk.Konnect
}

// MeshOPADataSourceModel describes the data model.
type MeshOPADataSourceModel struct {
	CpID             types.String            `tfsdk:"cp_id"`
	CreationTime     types.String            `tfsdk:"creation_time"`
	Labels           map[string]types.String `tfsdk:"labels"`
	Mesh             types.String            `tfsdk:"mesh"`
	ModificationTime types.String            `tfsdk:"modification_time"`
	Name             types.String            `tfsdk:"name"`
	Spec             tfTypes.MeshOPAItemSpec `tfsdk:"spec"`
	Type             types.String            `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *MeshOPADataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_opa"
}

// Schema defines the schema for the data source.
func (r *MeshOPADataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshOPA DataSource",

		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"creation_time": schema.StringAttribute{
				Computed:    true,
				Description: `Time at which the resource was created`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The labels to help identity resources`,
			},
			"mesh": schema.StringAttribute{
				Required:    true,
				Description: `name of the mesh`,
			},
			"modification_time": schema.StringAttribute{
				Computed:    true,
				Description: `Time at which the resource was updated`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `name of the MeshOPA`,
			},
			"spec": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"agent_config": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"inline": schema.StringAttribute{
										Computed:    true,
										Description: `Data source is inline bytes.`,
									},
									"inline_string": schema.StringAttribute{
										Computed:    true,
										Description: `Data source is inline string` + "`" + ``,
									},
									"secret": schema.StringAttribute{
										Computed:    true,
										Description: `Data source is a secret with given Secret key.`,
									},
								},
								Description: `AgentConfig defines bootstrap OPA agent configuration.`,
							},
							"append_policies": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ignore_decision": schema.BoolAttribute{
											Computed:    true,
											Description: `If true, then policy won't be taken into account when making a decision.`,
										},
										"rego": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"inline": schema.StringAttribute{
													Computed:    true,
													Description: `Data source is inline bytes.`,
												},
												"inline_string": schema.StringAttribute{
													Computed:    true,
													Description: `Data source is inline string` + "`" + ``,
												},
												"secret": schema.StringAttribute{
													Computed:    true,
													Description: `Data source is a secret with given Secret key.`,
												},
											},
											Description: `OPA Policy written in Rego. Available values: secret, inline, inlineString.`,
										},
									},
								},
								Description: `Policies define OPA policies that will be applied on OPA Agent.`,
							},
							"auth_config": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"on_agent_failure": schema.StringAttribute{
										Computed: true,
										MarkdownDescription: `OnAgentFailure either 'allow' or 'deny' (default to deny) whether` + "\n" +
											`to allow requests when the authorization agent failed.`,
									},
									"request_body": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"max_size": schema.Int32Attribute{
												Computed: true,
												MarkdownDescription: `MaxSize defines the maximum payload size sent to authorization agent. If the payload` + "\n" +
													`is larger it will be truncated and there will be a header` + "\n" +
													`` + "`" + `x-envoy-auth-partial-body: true` + "`" + `. If it is set to 0 no body will be` + "\n" +
													`sent to the agent.`,
											},
											"send_raw_body": schema.BoolAttribute{
												Computed:    true,
												Description: `SendRawBody enable sending raw body instead of the body encoded into UTF-8`,
											},
										},
										MarkdownDescription: `RequestBody configuration to apply on the request body sent to the` + "\n" +
											`authorization agent (if absent, the body is not sent).`,
									},
									"status_on_error": schema.Int32Attribute{
										Computed: true,
										MarkdownDescription: `StatusOnError is the http status to return when there's a connection` + "\n" +
											`failure between the dataplane and the authorization agent`,
									},
									"timeout": schema.StringAttribute{
										Computed:    true,
										Description: `Timeout for the single gRPC request from Envoy to OPA Agent.`,
									},
								},
								Description: `AuthConfig are configurations specific to the filter.`,
							},
						},
					},
					"target_ref": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Computed:    true,
								Description: `Kind of the referenced resource`,
							},
							"labels": schema.MapAttribute{
								Computed:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
									`Name and Namespace can be used.`,
							},
							"mesh": schema.StringAttribute{
								Computed:    true,
								Description: `Mesh is reserved for future use to identify cross mesh resources.`,
							},
							"name": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
									`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
							},
							"namespace": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
									`will be targeted.`,
							},
							"proxy_types": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
								MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
									`all data plane types are targeted by the policy.`,
							},
							"section_name": schema.StringAttribute{
								Computed: true,
								MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
									`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
							},
							"tags": schema.MapAttribute{
								Computed:    true,
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
				Computed:    true,
				Description: `the type of the resource`,
			},
		},
	}
}

func (r *MeshOPADataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *MeshOPADataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshOPADataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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
