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
var _ datasource.DataSource = &MeshHTTPRouteListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshHTTPRouteListDataSource{}

func NewMeshHTTPRouteListDataSource() datasource.DataSource {
	return &MeshHTTPRouteListDataSource{}
}

// MeshHTTPRouteListDataSource is the data source implementation.
type MeshHTTPRouteListDataSource struct {
	client *sdk.Konnect
}

// MeshHTTPRouteListDataSourceModel describes the data model.
type MeshHTTPRouteListDataSourceModel struct {
	CpID   types.String                `tfsdk:"cp_id"`
	Items  []tfTypes.MeshHTTPRouteItem `tfsdk:"items"`
	Key    types.String                `tfsdk:"key"`
	Mesh   types.String                `tfsdk:"mesh"`
	Next   types.String                `tfsdk:"next"`
	Offset types.Int64                 `tfsdk:"offset"`
	Size   types.Int64                 `tfsdk:"size"`
	Total  types.Number                `tfsdk:"total"`
	Value  types.String                `tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshHTTPRouteListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_http_route_list"
}

// Schema defines the schema for the data source.
func (r *MeshHTTPRouteListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshHTTPRouteList DataSource",

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
										`defined inplace.`,
								},
								"to": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"hostnames": schema.ListAttribute{
												Computed:    true,
												ElementType: types.StringType,
												MarkdownDescription: `Hostnames is only valid when targeting MeshGateway and limits the` + "\n" +
													`effects of the rules to requests to this hostname.` + "\n" +
													`Given hostnames must intersect with the hostname of the listeners the` + "\n" +
													`route attaches to.`,
											},
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
																"filters": schema.ListNestedAttribute{
																	Computed: true,
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"request_header_modifier": schema.SingleNestedAttribute{
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
																					"remove": schema.ListAttribute{
																						Computed:    true,
																						ElementType: types.StringType,
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
																				MarkdownDescription: `Only one action is supported per header name.` + "\n" +
																					`Configuration to set or add multiple values for a header must use RFC 7230` + "\n" +
																					`header value formatting, separating each value with a comma.`,
																			},
																			"request_mirror": schema.SingleNestedAttribute{
																				Computed: true,
																				Attributes: map[string]schema.Attribute{
																					"backend_ref": schema.SingleNestedAttribute{
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
																						Description: `BackendRef defines where to forward traffic.`,
																					},
																					"percentage": schema.SingleNestedAttribute{
																						Computed: true,
																						Attributes: map[string]schema.Attribute{
																							"integer": schema.Int64Attribute{
																								Computed: true,
																							},
																							"str": schema.StringAttribute{
																								Computed: true,
																							},
																						},
																						MarkdownDescription: `Percentage of requests to mirror. If not specified, all requests` + "\n" +
																							`to the target cluster will be mirrored.`,
																					},
																				},
																			},
																			"request_redirect": schema.SingleNestedAttribute{
																				Computed: true,
																				Attributes: map[string]schema.Attribute{
																					"hostname": schema.StringAttribute{
																						Computed: true,
																						MarkdownDescription: `PreciseHostname is the fully qualified domain name of a network host. This` + "\n" +
																							`matches the RFC 1123 definition of a hostname with 1 notable exception that` + "\n" +
																							`numeric IP addresses are not allowed.` + "\n" +
																							`` + "\n" +
																							`Note that as per RFC1035 and RFC1123, a *label* must consist of lower case` + "\n" +
																							`alphanumeric characters or '-', and must start and end with an alphanumeric` + "\n" +
																							`character. No other punctuation is allowed.`,
																					},
																					"path": schema.SingleNestedAttribute{
																						Computed: true,
																						Attributes: map[string]schema.Attribute{
																							"replace_full_path": schema.StringAttribute{
																								Computed: true,
																							},
																							"replace_prefix_match": schema.StringAttribute{
																								Computed: true,
																							},
																							"type": schema.StringAttribute{
																								Computed: true,
																							},
																						},
																						MarkdownDescription: `Path defines parameters used to modify the path of the incoming request.` + "\n" +
																							`The modified path is then used to construct the location header.` + "\n" +
																							`When empty, the request path is used as-is.`,
																					},
																					"port": schema.Int64Attribute{
																						Computed: true,
																						MarkdownDescription: `Port is the port to be used in the value of the ` + "`" + `Location` + "`" + `` + "\n" +
																							`header in the response.` + "\n" +
																							`When empty, port (if specified) of the request is used.`,
																					},
																					"scheme": schema.StringAttribute{
																						Computed: true,
																					},
																					"status_code": schema.Int64Attribute{
																						Computed:    true,
																						Description: `StatusCode is the HTTP status code to be used in response.`,
																					},
																				},
																			},
																			"response_header_modifier": schema.SingleNestedAttribute{
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
																					"remove": schema.ListAttribute{
																						Computed:    true,
																						ElementType: types.StringType,
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
																				MarkdownDescription: `Only one action is supported per header name.` + "\n" +
																					`Configuration to set or add multiple values for a header must use RFC 7230` + "\n" +
																					`header value formatting, separating each value with a comma.`,
																			},
																			"type": schema.StringAttribute{
																				Computed: true,
																			},
																			"url_rewrite": schema.SingleNestedAttribute{
																				Computed: true,
																				Attributes: map[string]schema.Attribute{
																					"host_to_backend_hostname": schema.BoolAttribute{
																						Computed: true,
																						MarkdownDescription: `HostToBackendHostname rewrites the hostname to the hostname of the` + "\n" +
																							`upstream host. This option is only available when targeting MeshGateways.`,
																					},
																					"hostname": schema.StringAttribute{
																						Computed:    true,
																						Description: `Hostname is the value to be used to replace the host header value during forwarding.`,
																					},
																					"path": schema.SingleNestedAttribute{
																						Computed: true,
																						Attributes: map[string]schema.Attribute{
																							"replace_full_path": schema.StringAttribute{
																								Computed: true,
																							},
																							"replace_prefix_match": schema.StringAttribute{
																								Computed: true,
																							},
																							"type": schema.StringAttribute{
																								Computed: true,
																							},
																						},
																						Description: `Path defines a path rewrite.`,
																					},
																				},
																			},
																		},
																	},
																},
															},
															MarkdownDescription: `Default holds routing rules that can be merged with rules from other` + "\n" +
																`policies.`,
														},
														"matches": schema.ListNestedAttribute{
															Computed: true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"headers": schema.ListNestedAttribute{
																		Computed: true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Computed: true,
																					MarkdownDescription: `Name is the name of the HTTP Header to be matched. Name MUST be lower case` + "\n" +
																						`as they will be handled with case insensitivity (See https://tools.ietf.org/html/rfc7230#section-3.2).`,
																				},
																				"type": schema.StringAttribute{
																					Computed:    true,
																					Description: `Type specifies how to match against the value of the header.`,
																				},
																				"value": schema.StringAttribute{
																					Computed:    true,
																					Description: `Value is the value of HTTP Header to be matched.`,
																				},
																			},
																		},
																	},
																	"method": schema.StringAttribute{
																		Computed: true,
																	},
																	"path": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"type": schema.StringAttribute{
																				Computed: true,
																			},
																			"value": schema.StringAttribute{
																				Computed: true,
																				MarkdownDescription: `Exact or prefix matches must be an absolute path. A prefix matches only` + "\n" +
																					`if separated by a slash or the entire path.`,
																			},
																		},
																	},
																	"query_params": schema.ListNestedAttribute{
																		Computed: true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Computed: true,
																				},
																				"type": schema.StringAttribute{
																					Computed: true,
																				},
																				"value": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																		},
																		MarkdownDescription: `QueryParams matches based on HTTP URL query parameters. Multiple matches` + "\n" +
																			`are ANDed together such that all listed matches must succeed.`,
																	},
																},
															},
															MarkdownDescription: `Matches describes how to match HTTP requests this rule should be applied` + "\n" +
																`to.`,
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
													`request destinations.`,
											},
										},
									},
									Description: `To matches destination services of requests and holds configuration.`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshHTTPRoute resource.`,
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
			"offset": schema.Int64Attribute{
				Optional:    true,
				Description: `offset in the list of entities`,
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Description: `the number of items per page`,
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

func (r *MeshHTTPRouteListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshHTTPRouteListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshHTTPRouteListDataSourceModel
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

	offset := new(int64)
	if !data.Offset.IsUnknown() && !data.Offset.IsNull() {
		*offset = data.Offset.ValueInt64()
	} else {
		offset = nil
	}
	size := new(int64)
	if !data.Size.IsUnknown() && !data.Size.IsNull() {
		*size = data.Size.ValueInt64()
	} else {
		size = nil
	}
	var filter *operations.GetMeshHTTPRouteListQueryParamFilter
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
	filter = &operations.GetMeshHTTPRouteListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	var mesh string
	mesh = data.Mesh.ValueString()

	request := operations.GetMeshHTTPRouteListRequest{
		CpID:   cpID,
		Offset: offset,
		Size:   size,
		Filter: filter,
		Mesh:   mesh,
	}
	res, err := r.client.MeshHTTPRoute.GetMeshHTTPRouteList(ctx, request)
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
	if !(res.MeshHTTPRouteList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshHTTPRouteList(res.MeshHTTPRouteList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
