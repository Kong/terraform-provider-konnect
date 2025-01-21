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
var _ datasource.DataSource = &MeshMetricDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshMetricDataSource{}

func NewMeshMetricDataSource() datasource.DataSource {
	return &MeshMetricDataSource{}
}

// MeshMetricDataSource is the data source implementation.
type MeshMetricDataSource struct {
	client *sdk.Konnect
}

// MeshMetricDataSourceModel describes the data model.
type MeshMetricDataSourceModel struct {
	CreationTime     types.String               `tfsdk:"creation_time"`
	ID               types.String               `tfsdk:"id"`
	Labels           map[string]types.String    `tfsdk:"labels"`
	Mesh             types.String               `tfsdk:"mesh"`
	ModificationTime types.String               `tfsdk:"modification_time"`
	Name             types.String               `tfsdk:"name"`
	Spec             tfTypes.MeshMetricItemSpec `tfsdk:"spec"`
	Type             types.String               `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *MeshMetricDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_metric"
}

// Schema defines the schema for the data source.
func (r *MeshMetricDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshMetric DataSource",

		Attributes: map[string]schema.Attribute{
			"creation_time": schema.StringAttribute{
				Computed:    true,
				Description: `Time at which the resource was created`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
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
				Description: `name of the MeshMetric`,
			},
			"spec": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"default": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"applications": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Computed:    true,
											Description: `Address on which an application listens.`,
										},
										"name": schema.StringAttribute{
											Computed:    true,
											Description: `Name of the application to scrape`,
										},
										"path": schema.StringAttribute{
											Computed:    true,
											Description: `Path on which an application expose HTTP endpoint with metrics.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Description: `Port on which an application expose HTTP endpoint with metrics.`,
										},
									},
								},
								Description: `Applications is a list of application that Dataplane Proxy will scrape`,
							},
							"backends": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"open_telemetry": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"endpoint": schema.StringAttribute{
													Computed:    true,
													Description: `Endpoint for OpenTelemetry collector`,
												},
												"refresh_interval": schema.StringAttribute{
													Computed:    true,
													Description: `RefreshInterval defines how frequent metrics should be pushed to collector`,
												},
											},
											Description: `OpenTelemetry backend configuration`,
										},
										"prometheus": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"client_id": schema.StringAttribute{
													Computed:    true,
													Description: `ClientId of the Prometheus backend. Needed when using MADS for DP discovery.`,
												},
												"path": schema.StringAttribute{
													Computed:    true,
													Description: `Path on which a dataplane should expose HTTP endpoint with Prometheus metrics.`,
												},
												"port": schema.Int64Attribute{
													Computed:    true,
													Description: `Port on which a dataplane should expose HTTP endpoint with Prometheus metrics.`,
												},
												"tls": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"mode": schema.StringAttribute{
															Computed:    true,
															Description: `Configuration of TLS for Prometheus listener.`,
														},
													},
													Description: `Configuration of TLS for prometheus listener.`,
												},
											},
											Description: `Prometheus backend configuration.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `Type of the backend that will be used to collect metrics. At the moment only Prometheus backend is available.`,
										},
									},
								},
								Description: `Backends list that will be used to collect metrics.`,
							},
							"sidecar": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"include_unused": schema.BoolAttribute{
										Computed: true,
										MarkdownDescription: `IncludeUnused if false will scrape only metrics that has been by sidecar (counters incremented` + "\n" +
											`at least once, gauges changed at least once, and histograms added to at` + "\n" +
											`least once). If true will scrape all metrics (even the ones with zeros).`,
									},
									"profiles": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"append_profiles": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Computed:    true,
															Description: `Name of the predefined profile, one of: all, basic, none`,
														},
													},
												},
												Description: `AppendProfiles allows to combine the metrics from multiple predefined profiles.`,
											},
											"exclude": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"match": schema.StringAttribute{
															Computed:    true,
															Description: `Match is the value used to match using particular Type`,
														},
														"type": schema.StringAttribute{
															Computed:    true,
															Description: `Type defined the type of selector, one of: prefix, regex, exact`,
														},
													},
												},
												MarkdownDescription: `Exclude makes it possible to exclude groups of metrics from a resulting profile.` + "\n" +
													`Exclude is subordinate to Include.`,
											},
											"include": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"match": schema.StringAttribute{
															Computed:    true,
															Description: `Match is the value used to match using particular Type`,
														},
														"type": schema.StringAttribute{
															Computed:    true,
															Description: `Type defined the type of selector, one of: prefix, regex, exact`,
														},
													},
												},
												MarkdownDescription: `Include makes it possible to include additional metrics in a selected profiles.` + "\n" +
													`Include takes precedence over Exclude.`,
											},
										},
										Description: `Profiles allows to customize which metrics are published.`,
									},
								},
								Description: `Sidecar metrics collection configuration`,
							},
						},
						Description: `MeshMetric configuration.`,
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
							`defined in-place.`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshMetric resource.`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `the type of the resource`,
			},
		},
	}
}

func (r *MeshMetricDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshMetricDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshMetricDataSourceModel
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

	var id string
	id = data.ID.ValueString()

	var mesh string
	mesh = data.Mesh.ValueString()

	var name string
	name = data.Name.ValueString()

	request := operations.GetMeshMetricRequest{
		ID:   id,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshMetric.GetMeshMetric(ctx, request)
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
	if !(res.MeshMetricItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMetricItem(res.MeshMetricItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
