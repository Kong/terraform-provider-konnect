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
var _ datasource.DataSource = &MeshTCPRouteListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshTCPRouteListDataSource{}

func NewMeshTCPRouteListDataSource() datasource.DataSource {
	return &MeshTCPRouteListDataSource{}
}

// MeshTCPRouteListDataSource is the data source implementation.
type MeshTCPRouteListDataSource struct {
	client *sdk.Konnect
}

// MeshTCPRouteListDataSourceModel describes the data model.
type MeshTCPRouteListDataSourceModel struct {
	CpID  types.String               `tfsdk:"cp_id"`
	Items []tfTypes.MeshTCPRouteItem `tfsdk:"items"`
	Key   types.String               `tfsdk:"key"`
	Mesh  types.String               `tfsdk:"mesh"`
	Next  types.String               `tfsdk:"next"`
	Total types.Number               `tfsdk:"total"`
	Value types.String               `tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshTCPRouteListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_tcp_route_list"
}

// Schema defines the schema for the data source.
func (r *MeshTCPRouteListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshTCPRouteList DataSource",

		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
							Computed:    true,
							Description: `Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.`,
						},
						"modification_time": schema.StringAttribute{
							Computed:    true,
							Description: `Time at which the resource was updated`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `Name of the Kuma resource`,
						},
						"spec": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
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
										`defined in-place.`,
								},
								"to": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"rules": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"default": schema.SingleNestedAttribute{
															Computed: true,
															Attributes: map[string]schema.Attribute{
																"backend_refs": schema.ListNestedAttribute{
																	Computed: true,
																	NestedObject: schema.NestedAttributeObject{
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
																			"port": schema.Int64Attribute{
																				Computed:    true,
																				Description: `Port is only supported when this ref refers to a real MeshService object`,
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
																			"weight": schema.Int64Attribute{
																				Computed: true,
																			},
																		},
																	},
																},
															},
															MarkdownDescription: `Default holds routing rules that can be merged with rules from other` + "\n" +
																`policies.`,
														},
													},
												},
												MarkdownDescription: `Rules contains the routing rules applies to a combination of top-level` + "\n" +
													`targetRef and the targetRef in this entry.`,
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
												MarkdownDescription: `TargetRef is a reference to the resource that represents a group of` + "\n" +
													`destinations.`,
											},
										},
									},
									MarkdownDescription: `To list makes a match between the consumed services and corresponding` + "\n" +
										`configurations`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshTCPRoute resource.`,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: `the type of the resource`,
						},
					},
				},
			},
			"key": schema.StringAttribute{
				Optional: true,
			},
			"mesh": schema.StringAttribute{
				Required:    true,
				Description: `name of the mesh`,
			},
			"next": schema.StringAttribute{
				Computed:    true,
				Description: `URL to the next page`,
			},
			"total": schema.NumberAttribute{
				Computed:    true,
				Description: `The total number of entities`,
			},
			"value": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *MeshTCPRouteListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshTCPRouteListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshTCPRouteListDataSourceModel
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

	var filter *operations.GetMeshTCPRouteListQueryParamFilter
	key := new(string)
	if !data.Key.IsUnknown() && !data.Key.IsNull() {
		*key = data.Key.ValueString()
	} else {
		key = nil
	}
	value := new(string)
	if !data.Value.IsUnknown() && !data.Value.IsNull() {
		*value = data.Value.ValueString()
	} else {
		value = nil
	}
	filter = &operations.GetMeshTCPRouteListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	var mesh string
	mesh = data.Mesh.ValueString()

	request := operations.GetMeshTCPRouteListRequest{
		CpID:   cpID,
		Filter: filter,
		Mesh:   mesh,
	}
	res, err := r.client.MeshTCPRoute.GetMeshTCPRouteList(ctx, request)
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
	if !(res.MeshTCPRouteList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshTCPRouteList(res.MeshTCPRouteList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
