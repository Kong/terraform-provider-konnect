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
var _ datasource.DataSource = &MeshDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshDataSource{}

func NewMeshDataSource() datasource.DataSource {
	return &MeshDataSource{}
}

// MeshDataSource is the data source implementation.
type MeshDataSource struct {
	client *sdk.Konnect
}

// MeshDataSourceModel describes the data model.
type MeshDataSourceModel struct {
	Constraints                 *tfTypes.Constraints    `tfsdk:"constraints"`
	CpID                        types.String            `tfsdk:"cp_id"`
	Labels                      map[string]types.String `tfsdk:"labels"`
	Logging                     *tfTypes.Logging        `tfsdk:"logging"`
	MeshServices                *tfTypes.MeshServices   `tfsdk:"mesh_services"`
	Metrics                     *tfTypes.Metrics        `tfsdk:"metrics"`
	Mtls                        *tfTypes.Mtls           `tfsdk:"mtls"`
	Name                        types.String            `tfsdk:"name"`
	Networking                  *tfTypes.Networking     `tfsdk:"networking"`
	Routing                     *tfTypes.Routing        `tfsdk:"routing"`
	SkipCreatingInitialPolicies []types.String          `tfsdk:"skip_creating_initial_policies"`
	Tracing                     *tfTypes.Tracing        `tfsdk:"tracing"`
	Type                        types.String            `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *MeshDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh"
}

// Schema defines the schema for the data source.
func (r *MeshDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Mesh DataSource",

		Attributes: map[string]schema.Attribute{
			"constraints": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"dataplane_proxy": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"requirements": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"tags": schema.MapAttribute{
											Computed:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Tags defines set of required tags. You can specify '*' in value to` + "\n" +
												`require non empty value of tag`,
										},
									},
								},
								MarkdownDescription: `Requirements defines a set of requirements that data plane proxies must` + "\n" +
									`fulfill in order to join the mesh. A data plane proxy must fulfill at` + "\n" +
									`least one requirement in order to join the mesh. Empty list of allowed` + "\n" +
									`requirements means that any proxy that is not explicitly denied can join.`,
							},
							"restrictions": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"tags": schema.MapAttribute{
											Computed:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Tags defines set of required tags. You can specify '*' in value to` + "\n" +
												`require non empty value of tag`,
										},
									},
								},
								MarkdownDescription: `Restrictions defines a set of restrictions that data plane proxies cannot` + "\n" +
									`fulfill in order to join the mesh. A data plane proxy cannot fulfill any` + "\n" +
									`requirement in order to join the mesh.` + "\n" +
									`Restrictions takes precedence over requirements.`,
							},
						},
						MarkdownDescription: `DataplaneProxyMembership defines a set of requirements for data plane` + "\n" +
							`proxies to be a member of the mesh.`,
					},
				},
				Description: `Constraints that applies to the mesh and its entities`,
			},
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"logging": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"backends": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conf": schema.StringAttribute{
									Computed:    true,
									Description: `Configuration of the backend. Parsed as JSON.`,
								},
								"format": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: `Format of access logs. Placeholders available on` + "\n" +
										`https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log`,
								},
								"name": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: `Name of the backend, can be then used in Mesh.logging.defaultBackend or in` + "\n" +
										`TrafficLogging`,
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `Type of the backend (Kuma ships with 'tcp' and 'file')`,
								},
							},
						},
						Description: `List of available logging backends`,
					},
					"default_backend": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the default backend`,
					},
				},
				MarkdownDescription: `Logging settings.` + "\n" +
					`+optional`,
			},
			"mesh_services": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"mode": schema.SingleNestedAttribute{
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
			},
			"metrics": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"backends": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conf": schema.StringAttribute{
									Computed:    true,
									Description: `Configuration of the backend. Parsed as JSON.`,
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `Name of the backend, can be then used in Mesh.metrics.enabledBackend`,
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `Type of the backend (Kuma ships with 'prometheus')`,
								},
							},
						},
						Description: `List of available Metrics backends`,
					},
					"enabled_backend": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the enabled backend`,
					},
				},
				MarkdownDescription: `Configuration for metrics collected and exposed by dataplanes.` + "\n" +
					`` + "\n" +
					`Settings defined here become defaults for every dataplane in a given Mesh.` + "\n" +
					`Additionally, it is also possible to further customize this configuration` + "\n" +
					`for each dataplane individually using Dataplane resource.` + "\n" +
					`+optional`,
			},
			"mtls": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"backends": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conf": schema.StringAttribute{
									Computed:    true,
									Description: `Configuration of the backend. Parsed as JSON.`,
								},
								"dp_cert": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"request_timeout": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"nanos": schema.Int64Attribute{
													Computed: true,
												},
												"seconds": schema.Int64Attribute{
													Computed: true,
												},
											},
											Description: `Timeout on request to CA for DP certificate generation and retrieval`,
										},
										"rotation": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"expiration": schema.StringAttribute{
													Computed:    true,
													Description: `Time after which generated certificate for Dataplane will expire`,
												},
											},
											Description: `Rotation settings`,
										},
									},
									Description: `Dataplane certificate settings`,
								},
								"mode": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"integer": schema.Int64Attribute{
											Computed: true,
										},
										"str": schema.StringAttribute{
											Computed: true,
										},
									},
									MarkdownDescription: `Mode defines the behaviour of inbound listeners with regard to traffic` + "\n" +
										`encryption`,
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `Name of the backend`,
								},
								"root_chain": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"request_timeout": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"nanos": schema.Int64Attribute{
													Computed: true,
												},
												"seconds": schema.Int64Attribute{
													Computed: true,
												},
											},
											MarkdownDescription: `Timeout on request for to CA for root certificate chain.` + "\n" +
												`If not specified, defaults to 10s.`,
										},
									},
								},
								"type": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: `Type of the backend. Has to be one of the loaded plugins (Kuma ships with` + "\n" +
										`builtin and provided)`,
								},
							},
						},
						Description: `List of available Certificate Authority backends`,
					},
					"enabled_backend": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the enabled backend`,
					},
					"skip_validation": schema.BoolAttribute{
						Computed:    true,
						Description: `If enabled, skips CA validation.`,
					},
				},
				MarkdownDescription: `mTLS settings.` + "\n" +
					`+optional`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `name of the Mesh`,
			},
			"networking": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"outbound": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"passthrough": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"value": schema.BoolAttribute{
										Computed: true,
									},
								},
								Description: `Control the passthrough cluster`,
							},
						},
						Description: `Outbound settings`,
					},
				},
				Description: `Networking settings of the mesh`,
			},
			"routing": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"default_forbid_mesh_external_service_access": schema.BoolAttribute{
						Computed: true,
						MarkdownDescription: `If true, blocks traffic to MeshExternalServices.` + "\n" +
							`Default: false`,
					},
					"locality_aware_load_balancing": schema.BoolAttribute{
						Computed:    true,
						Description: `Enable the Locality Aware Load Balancing`,
					},
					"zone_egress": schema.BoolAttribute{
						Computed: true,
						MarkdownDescription: `Enable routing traffic to services in other zone or external services` + "\n" +
							`through ZoneEgress. Default: false`,
					},
				},
				Description: `Routing settings of the mesh`,
			},
			"skip_creating_initial_policies": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				MarkdownDescription: `List of policies to skip creating by default when the mesh is created.` + "\n" +
					`e.g. TrafficPermission, MeshRetry, etc. An '*' can be used to skip all` + "\n" +
					`policies.`,
			},
			"tracing": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"backends": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conf": schema.StringAttribute{
									Computed:    true,
									Description: `Configuration of the backend. Parsed as JSON.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: `Name of the backend, can be then used in Mesh.tracing.defaultBackend or in` + "\n" +
										`TrafficTrace`,
								},
								"sampling": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"value": schema.NumberAttribute{
											Computed: true,
										},
									},
									MarkdownDescription: `Percentage of traces that will be sent to the backend (range 0.0 - 100.0).` + "\n" +
										`Empty value defaults to 100.0%`,
								},
								"type": schema.StringAttribute{
									Computed:    true,
									Description: `Type of the backend (Kuma ships with 'zipkin')`,
								},
							},
						},
						Description: `List of available tracing backends`,
					},
					"default_backend": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the default backend`,
					},
				},
				MarkdownDescription: `Tracing settings.` + "\n" +
					`+optional`,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *MeshDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshDataSourceModel
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

	var name string
	name = data.Name.ValueString()

	request := operations.GetMeshRequest{
		CpID: cpID,
		Name: name,
	}
	res, err := r.client.Mesh.GetMesh(ctx, request)
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
	if !(res.MeshItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshItem(res.MeshItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}