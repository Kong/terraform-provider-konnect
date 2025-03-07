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
var _ datasource.DataSource = &MeshLoadBalancingStrategyListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshLoadBalancingStrategyListDataSource{}

func NewMeshLoadBalancingStrategyListDataSource() datasource.DataSource {
	return &MeshLoadBalancingStrategyListDataSource{}
}

// MeshLoadBalancingStrategyListDataSource is the data source implementation.
type MeshLoadBalancingStrategyListDataSource struct {
	client *sdk.Konnect
}

// MeshLoadBalancingStrategyListDataSourceModel describes the data model.
type MeshLoadBalancingStrategyListDataSourceModel struct {
	CpID   types.String                            `tfsdk:"cp_id"`
	Items  []tfTypes.MeshLoadBalancingStrategyItem `tfsdk:"items"`
	Key    types.String                            `queryParam:"name=key" tfsdk:"key"`
	Mesh   types.String                            `tfsdk:"mesh"`
	Next   types.String                            `tfsdk:"next"`
	Offset types.Int64                             `queryParam:"style=form,explode=true,name=offset" tfsdk:"offset"`
	Size   types.Int64                             `queryParam:"style=form,explode=true,name=size" tfsdk:"size"`
	Total  types.Number                            `tfsdk:"total"`
	Value  types.String                            `queryParam:"name=value" tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshLoadBalancingStrategyListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_load_balancing_strategy_list"
}

// Schema defines the schema for the data source.
func (r *MeshLoadBalancingStrategyListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshLoadBalancingStrategyList DataSource",

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
											"default": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"load_balancer": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"least_request": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"active_request_bias": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"integer": schema.Int64Attribute{
																				Computed: true,
																			},
																			"str": schema.StringAttribute{
																				Computed: true,
																			},
																		},
																		MarkdownDescription: `ActiveRequestBias refers to dynamic weights applied when hosts have varying load` + "\n" +
																			`balancing weights. A higher value here aggressively reduces the weight of endpoints` + "\n" +
																			`that are currently handling active requests. In essence, the higher the ActiveRequestBias` + "\n" +
																			`value, the more forcefully it reduces the load balancing weight of endpoints that are` + "\n" +
																			`actively serving requests.`,
																	},
																	"choice_count": schema.Int32Attribute{
																		Computed: true,
																		MarkdownDescription: `ChoiceCount is the number of random healthy hosts from which the host with` + "\n" +
																			`the fewest active requests will be chosen. Defaults to 2 so that Envoy performs` + "\n" +
																			`two-choice selection if the field is not set.`,
																	},
																},
																MarkdownDescription: `LeastRequest selects N random available hosts as specified in 'choiceCount' (2 by default)` + "\n" +
																	`and picks the host which has the fewest active requests`,
															},
															"maglev": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"hash_policies": schema.ListNestedAttribute{
																		Computed: true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"connection": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"source_ip": schema.BoolAttribute{
																							Computed:    true,
																							Description: `Hash on source IP address.`,
																						},
																					},
																				},
																				"cookie": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Computed:    true,
																							Description: `The name of the cookie that will be used to obtain the hash key.`,
																						},
																						"path": schema.StringAttribute{
																							Computed:    true,
																							Description: `The name of the path for the cookie.`,
																						},
																						"ttl": schema.StringAttribute{
																							Computed:    true,
																							Description: `If specified, a cookie with the TTL will be generated if the cookie is not present.`,
																						},
																					},
																				},
																				"filter_state": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Computed: true,
																							MarkdownDescription: `The name of the Object in the per-request filterState, which is` + "\n" +
																								`an Envoy::Hashable object. If there is no data associated with the key,` + "\n" +
																								`or the stored object is not Envoy::Hashable, no hash will be produced.`,
																						},
																					},
																				},
																				"header": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Computed:    true,
																							Description: `The name of the request header that will be used to obtain the hash key.`,
																						},
																					},
																				},
																				"query_parameter": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Computed: true,
																							MarkdownDescription: `The name of the URL query parameter that will be used to obtain the hash key.` + "\n" +
																								`If the parameter is not present, no hash will be produced. Query parameter names` + "\n" +
																								`are case-sensitive.`,
																						},
																					},
																				},
																				"terminal": schema.BoolAttribute{
																					Computed: true,
																					MarkdownDescription: `Terminal is a flag that short-circuits the hash computing. This field provides` + "\n" +
																						`a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback` + "\n" +
																						`to rest of the policy list”, it saves time when the terminal policy works.` + "\n" +
																						`If true, and there is already a hash computed, ignore rest of the list of hash polices.`,
																				},
																				"type": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																		},
																		MarkdownDescription: `HashPolicies specify a list of request/connection properties that are used to calculate a hash.` + "\n" +
																			`These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute` + "\n" +
																			`set to true, and there is already a hash generated, the hash is returned immediately,` + "\n" +
																			`ignoring the rest of the hash policy list.`,
																	},
																	"table_size": schema.Int32Attribute{
																		Computed: true,
																		MarkdownDescription: `The table size for Maglev hashing. Maglev aims for “minimal disruption”` + "\n" +
																			`rather than an absolute guarantee. Minimal disruption means that when` + "\n" +
																			`the set of upstream hosts change, a connection will likely be sent` + "\n" +
																			`to the same upstream as it was before. Increasing the table size reduces` + "\n" +
																			`the amount of disruption. The table size must be prime number limited to 5000011.` + "\n" +
																			`If it is not specified, the default is 65537.`,
																	},
																},
																MarkdownDescription: `Maglev implements consistent hashing to upstream hosts. Maglev can be used as` + "\n" +
																	`a drop in replacement for the ring hash load balancer any place in which` + "\n" +
																	`consistent hashing is desired.`,
															},
															"random": schema.SingleNestedAttribute{
																Computed: true,
																MarkdownDescription: `Random selects a random available host. The random load balancer generally` + "\n" +
																	`performs better than round-robin if no health checking policy is configured.` + "\n" +
																	`Random selection avoids bias towards the host in the set that comes after a failed host.`,
															},
															"ring_hash": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"hash_function": schema.StringAttribute{
																		Computed: true,
																		MarkdownDescription: `HashFunction is a function used to hash hosts onto the ketama ring.` + "\n" +
																			`The value defaults to XX_HASH. Available values – XX_HASH, MURMUR_HASH_2.`,
																	},
																	"hash_policies": schema.ListNestedAttribute{
																		Computed: true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"connection": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"source_ip": schema.BoolAttribute{
																							Computed:    true,
																							Description: `Hash on source IP address.`,
																						},
																					},
																				},
																				"cookie": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Computed:    true,
																							Description: `The name of the cookie that will be used to obtain the hash key.`,
																						},
																						"path": schema.StringAttribute{
																							Computed:    true,
																							Description: `The name of the path for the cookie.`,
																						},
																						"ttl": schema.StringAttribute{
																							Computed:    true,
																							Description: `If specified, a cookie with the TTL will be generated if the cookie is not present.`,
																						},
																					},
																				},
																				"filter_state": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Computed: true,
																							MarkdownDescription: `The name of the Object in the per-request filterState, which is` + "\n" +
																								`an Envoy::Hashable object. If there is no data associated with the key,` + "\n" +
																								`or the stored object is not Envoy::Hashable, no hash will be produced.`,
																						},
																					},
																				},
																				"header": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Computed:    true,
																							Description: `The name of the request header that will be used to obtain the hash key.`,
																						},
																					},
																				},
																				"query_parameter": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Computed: true,
																							MarkdownDescription: `The name of the URL query parameter that will be used to obtain the hash key.` + "\n" +
																								`If the parameter is not present, no hash will be produced. Query parameter names` + "\n" +
																								`are case-sensitive.`,
																						},
																					},
																				},
																				"terminal": schema.BoolAttribute{
																					Computed: true,
																					MarkdownDescription: `Terminal is a flag that short-circuits the hash computing. This field provides` + "\n" +
																						`a ‘fallback’ style of configuration: “if a terminal policy doesn’t work, fallback` + "\n" +
																						`to rest of the policy list”, it saves time when the terminal policy works.` + "\n" +
																						`If true, and there is already a hash computed, ignore rest of the list of hash polices.`,
																				},
																				"type": schema.StringAttribute{
																					Computed: true,
																				},
																			},
																		},
																		MarkdownDescription: `HashPolicies specify a list of request/connection properties that are used to calculate a hash.` + "\n" +
																			`These hash policies are executed in the specified order. If a hash policy has the “terminal” attribute` + "\n" +
																			`set to true, and there is already a hash generated, the hash is returned immediately,` + "\n" +
																			`ignoring the rest of the hash policy list.`,
																	},
																	"max_ring_size": schema.Int32Attribute{
																		Computed: true,
																		MarkdownDescription: `Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries,` + "\n" +
																			`but can be lowered to further constrain resource use.`,
																	},
																	"min_ring_size": schema.Int32Attribute{
																		Computed: true,
																		MarkdownDescription: `Minimum hash ring size. The larger the ring is (that is,` + "\n" +
																			`the more hashes there are for each provided host) the better the request distribution` + "\n" +
																			`will reflect the desired weights. Defaults to 1024 entries, and limited to 8M entries.`,
																	},
																},
																MarkdownDescription: `RingHash  implements consistent hashing to upstream hosts. Each host is mapped` + "\n" +
																	`onto a circle (the “ring”) by hashing its address; each request is then routed` + "\n" +
																	`to a host by hashing some property of the request, and finding the nearest` + "\n" +
																	`corresponding host clockwise around the ring.`,
															},
															"round_robin": schema.SingleNestedAttribute{
																Computed: true,
																MarkdownDescription: `RoundRobin is a load balancing algorithm that distributes requests` + "\n" +
																	`across available upstream hosts in round-robin order.`,
															},
															"type": schema.StringAttribute{
																Computed: true,
															},
														},
														Description: `LoadBalancer allows to specify load balancing algorithm.`,
													},
													"locality_awareness": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"cross_zone": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"failover": schema.ListNestedAttribute{
																		Computed: true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"from": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"zones": schema.ListAttribute{
																							Computed:    true,
																							ElementType: types.StringType,
																						},
																					},
																					Description: `From defines the list of zones to which the rule applies`,
																				},
																				"to": schema.SingleNestedAttribute{
																					Computed: true,
																					Attributes: map[string]schema.Attribute{
																						"type": schema.StringAttribute{
																							Computed:    true,
																							Description: `Type defines how target zones will be picked from available zones`,
																						},
																						"zones": schema.ListAttribute{
																							Computed:    true,
																							ElementType: types.StringType,
																						},
																					},
																					Description: `To defines to which zones the traffic should be load balanced`,
																				},
																			},
																		},
																		Description: `Failover defines list of load balancing rules in order of priority`,
																	},
																	"failover_threshold": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
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
																			},
																		},
																		MarkdownDescription: `FailoverThreshold defines the percentage of live destination dataplane proxies below which load balancing to the` + "\n" +
																			`next priority starts.` + "\n" +
																			`Example: If you configure failoverThreshold to 70, and you have deployed 10 destination dataplane proxies.` + "\n" +
																			`Load balancing to next priority will start when number of live destination dataplane proxies drops below 7.` + "\n" +
																			`Default 50`,
																	},
																},
																MarkdownDescription: `CrossZone defines locality aware load balancing priorities when dataplane proxies inside local zone` + "\n" +
																	`are unavailable`,
															},
															"disabled": schema.BoolAttribute{
																Computed: true,
																MarkdownDescription: `Disabled allows to disable locality-aware load balancing.` + "\n" +
																	`When disabled requests are distributed across all endpoints regardless of locality.`,
															},
															"local_zone": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"affinity_tags": schema.ListNestedAttribute{
																		Computed: true,
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Computed:    true,
																					Description: `Key defines tag for which affinity is configured`,
																				},
																				"weight": schema.Int32Attribute{
																					Computed: true,
																					MarkdownDescription: `Weight of the tag used for load balancing. The bigger the weight the bigger the priority.` + "\n" +
																						`Percentage of local traffic load balanced to tag is computed by dividing weight by sum of weights from all tags.` + "\n" +
																						`For example with two affinity tags first with weight 80 and second with weight 20,` + "\n" +
																						`then 80% of traffic will be redirected to the first tag, and 20% of traffic will be redirected to second one.` + "\n" +
																						`Setting weights is not mandatory. When weights are not set control plane will compute default weight based on list order.` + "\n" +
																						`Default: If you do not specify weight we will adjust them so that 90% traffic goes to first tag, 9% to next, and 1% to third and so on.`,
																				},
																			},
																		},
																		Description: `AffinityTags list of tags for local zone load balancing.`,
																	},
																},
																Description: `LocalZone defines locality aware load balancing priorities between dataplane proxies inside a zone`,
															},
														},
														Description: `LocalityAwareness contains configuration for locality aware load balancing.`,
													},
												},
												MarkdownDescription: `Default is a configuration specific to the group of destinations referenced in` + "\n" +
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
													`destinations.`,
											},
										},
									},
									Description: `To list makes a match between the consumed services and corresponding configurations`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshLoadBalancingStrategy resource.`,
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

func (r *MeshLoadBalancingStrategyListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshLoadBalancingStrategyListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshLoadBalancingStrategyListDataSourceModel
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
	var filter *operations.GetMeshLoadBalancingStrategyListQueryParamFilter
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
	filter = &operations.GetMeshLoadBalancingStrategyListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	var mesh string
	mesh = data.Mesh.ValueString()

	request := operations.GetMeshLoadBalancingStrategyListRequest{
		CpID:   cpID,
		Offset: offset,
		Size:   size,
		Filter: filter,
		Mesh:   mesh,
	}
	res, err := r.client.MeshLoadBalancingStrategy.GetMeshLoadBalancingStrategyList(ctx, request)
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
	if !(res.MeshLoadBalancingStrategyList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshLoadBalancingStrategyList(res.MeshLoadBalancingStrategyList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
