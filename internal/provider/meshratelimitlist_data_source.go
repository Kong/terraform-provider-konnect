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
var _ datasource.DataSource = &MeshRateLimitListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshRateLimitListDataSource{}

func NewMeshRateLimitListDataSource() datasource.DataSource {
	return &MeshRateLimitListDataSource{}
}

// MeshRateLimitListDataSource is the data source implementation.
type MeshRateLimitListDataSource struct {
	client *sdk.Konnect
}

// MeshRateLimitListDataSourceModel describes the data model.
type MeshRateLimitListDataSourceModel struct {
	CpID  types.String                `tfsdk:"cp_id"`
	Items []tfTypes.MeshRateLimitItem `tfsdk:"items"`
	Mesh  types.String                `tfsdk:"mesh"`
	Next  types.String                `tfsdk:"next"`
	Total types.Number                `tfsdk:"total"`
}

// Metadata returns the data source type name.
func (r *MeshRateLimitListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_rate_limit_list"
}

// Schema defines the schema for the data source.
func (r *MeshRateLimitListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshRateLimitList DataSource",

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
								"from": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"local": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"http": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"disabled": schema.BoolAttribute{
																		Computed:    true,
																		Description: `Define if rate limiting should be disabled.`,
																	},
																	"on_rate_limit": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"headers": schema.SingleNestedAttribute{
																				Computed: true,
																				Attributes: map[string]schema.Attribute{
																					"add": schema.ListNestedAttribute{
																						Computed: true,
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Computed: true,
																								},
																								"value": schema.StringAttribute{
																									Computed: true,
																								},
																							},
																						},
																					},
																					"set": schema.ListNestedAttribute{
																						Computed: true,
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Computed: true,
																								},
																								"value": schema.StringAttribute{
																									Computed: true,
																								},
																							},
																						},
																					},
																				},
																				Description: `The Headers to be added to the HTTP response on a rate limit event`,
																			},
																			"status": schema.Int64Attribute{
																				Computed:    true,
																				Description: `The HTTP status code to be set on a rate limit event`,
																			},
																		},
																		Description: `Describes the actions to take on a rate limit event`,
																	},
																	"request_rate": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"interval": schema.StringAttribute{
																				Computed:    true,
																				Description: `The interval the number of units is accounted for.`,
																			},
																			"num": schema.Int64Attribute{
																				Computed: true,
																				MarkdownDescription: `Number of units per interval (depending on usage it can be a number of requests,` + "\n" +
																					`or a number of connections).`,
																			},
																		},
																		Description: `Defines how many requests are allowed per interval.`,
																	},
																},
																MarkdownDescription: `LocalHTTP defines configuration of local HTTP rate limiting` + "\n" +
																	`https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter`,
															},
															"tcp": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"connection_rate": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"interval": schema.StringAttribute{
																				Computed:    true,
																				Description: `The interval the number of units is accounted for.`,
																			},
																			"num": schema.Int64Attribute{
																				Computed: true,
																				MarkdownDescription: `Number of units per interval (depending on usage it can be a number of requests,` + "\n" +
																					`or a number of connections).`,
																			},
																		},
																		Description: `Defines how many connections are allowed per interval.`,
																	},
																	"disabled": schema.BoolAttribute{
																		Computed: true,
																		MarkdownDescription: `Define if rate limiting should be disabled.` + "\n" +
																			`Default: false`,
																	},
																},
																MarkdownDescription: `LocalTCP defines confguration of local TCP rate limiting` + "\n" +
																	`https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter`,
															},
														},
														Description: `LocalConf defines local http or/and tcp rate limit configuration`,
													},
												},
												MarkdownDescription: `Default is a configuration specific to the group of clients referenced in` + "\n" +
													`'targetRef'`,
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
													`clients.`,
											},
										},
									},
									Description: `From list makes a match between clients and corresponding configurations`,
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
								"to": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"default": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"local": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"http": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"disabled": schema.BoolAttribute{
																		Computed:    true,
																		Description: `Define if rate limiting should be disabled.`,
																	},
																	"on_rate_limit": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"headers": schema.SingleNestedAttribute{
																				Computed: true,
																				Attributes: map[string]schema.Attribute{
																					"add": schema.ListNestedAttribute{
																						Computed: true,
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Computed: true,
																								},
																								"value": schema.StringAttribute{
																									Computed: true,
																								},
																							},
																						},
																					},
																					"set": schema.ListNestedAttribute{
																						Computed: true,
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Computed: true,
																								},
																								"value": schema.StringAttribute{
																									Computed: true,
																								},
																							},
																						},
																					},
																				},
																				Description: `The Headers to be added to the HTTP response on a rate limit event`,
																			},
																			"status": schema.Int64Attribute{
																				Computed:    true,
																				Description: `The HTTP status code to be set on a rate limit event`,
																			},
																		},
																		Description: `Describes the actions to take on a rate limit event`,
																	},
																	"request_rate": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"interval": schema.StringAttribute{
																				Computed:    true,
																				Description: `The interval the number of units is accounted for.`,
																			},
																			"num": schema.Int64Attribute{
																				Computed: true,
																				MarkdownDescription: `Number of units per interval (depending on usage it can be a number of requests,` + "\n" +
																					`or a number of connections).`,
																			},
																		},
																		Description: `Defines how many requests are allowed per interval.`,
																	},
																},
																MarkdownDescription: `LocalHTTP defines configuration of local HTTP rate limiting` + "\n" +
																	`https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter`,
															},
															"tcp": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"connection_rate": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"interval": schema.StringAttribute{
																				Computed:    true,
																				Description: `The interval the number of units is accounted for.`,
																			},
																			"num": schema.Int64Attribute{
																				Computed: true,
																				MarkdownDescription: `Number of units per interval (depending on usage it can be a number of requests,` + "\n" +
																					`or a number of connections).`,
																			},
																		},
																		Description: `Defines how many connections are allowed per interval.`,
																	},
																	"disabled": schema.BoolAttribute{
																		Computed: true,
																		MarkdownDescription: `Define if rate limiting should be disabled.` + "\n" +
																			`Default: false`,
																	},
																},
																MarkdownDescription: `LocalTCP defines confguration of local TCP rate limiting` + "\n" +
																	`https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter`,
															},
														},
														Description: `LocalConf defines local http or/and tcp rate limit configuration`,
													},
												},
												MarkdownDescription: `Default is a configuration specific to the group of clients referenced in` + "\n" +
													`'targetRef'`,
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
													`clients.`,
											},
										},
									},
									Description: `To list makes a match between clients and corresponding configurations`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshRateLimit resource.`,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: `the type of the resource`,
						},
					},
				},
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
		},
	}
}

func (r *MeshRateLimitListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshRateLimitListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshRateLimitListDataSourceModel
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

	request := operations.GetMeshRateLimitListRequest{
		CpID: cpID,
		Mesh: mesh,
	}
	res, err := r.client.MeshRateLimit.GetMeshRateLimitList(ctx, request)
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
	if !(res.MeshRateLimitList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshRateLimitList(res.MeshRateLimitList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}