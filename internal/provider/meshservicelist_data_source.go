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
var _ datasource.DataSource = &MeshServiceListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshServiceListDataSource{}

func NewMeshServiceListDataSource() datasource.DataSource {
	return &MeshServiceListDataSource{}
}

// MeshServiceListDataSource is the data source implementation.
type MeshServiceListDataSource struct {
	client *sdk.Konnect
}

// MeshServiceListDataSourceModel describes the data model.
type MeshServiceListDataSourceModel struct {
	CpID   types.String              `tfsdk:"cp_id"`
	Items  []tfTypes.MeshServiceItem `tfsdk:"items"`
	Key    types.String              `tfsdk:"key"`
	Mesh   types.String              `tfsdk:"mesh"`
	Next   types.String              `tfsdk:"next"`
	Offset types.Int64               `tfsdk:"offset"`
	Size   types.Int64               `tfsdk:"size"`
	Total  types.Number              `tfsdk:"total"`
	Value  types.String              `tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshServiceListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_service_list"
}

// Schema defines the schema for the data source.
func (r *MeshServiceListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshServiceList DataSource",

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
								"identities": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"ports": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"app_protocol": schema.StringAttribute{
												Computed:    true,
												Description: `Protocol identifies a protocol supported by a service.`,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
											"port": schema.Int64Attribute{
												Computed: true,
											},
											"target_port": schema.SingleNestedAttribute{
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
								},
								"selector": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"dataplane_ref": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"dataplane_tags": schema.MapAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
									},
								},
								"state": schema.StringAttribute{
									Computed: true,
									MarkdownDescription: `State of MeshService. Available if there is at least one healthy endpoint. Otherwise, Unavailable.` + "\n" +
										`It's used for cross zone communication to check if we should send traffic to it, when MeshService is aggregated into MeshMultiZoneService.`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshService resource.`,
						},
						"status": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"addresses": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
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
								},
								"dataplane_proxies": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"connected": schema.Int64Attribute{
											Computed:    true,
											Description: `Number of data plane proxies connected to the zone control plane`,
										},
										"healthy": schema.Int64Attribute{
											Computed:    true,
											Description: `Number of data plane proxies with all healthy inbounds selected by this MeshService.`,
										},
										"total": schema.Int64Attribute{
											Computed:    true,
											Description: `Total number of data plane proxies.`,
										},
									},
									Description: `Data plane proxies statistics selected by this MeshService.`,
								},
								"hostname_generators": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"conditions": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"message": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `message is a human readable message indicating details about the transition.` + "\n" +
																`This may be an empty string.`,
														},
														"reason": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `reason contains a programmatic identifier indicating the reason for the condition's last transition.` + "\n" +
																`Producers of specific condition types may define expected values and meanings for this field,` + "\n" +
																`and whether the values are considered a guaranteed API.` + "\n" +
																`The value should be a CamelCase string.` + "\n" +
																`This field may not be empty.`,
														},
														"status": schema.StringAttribute{
															Computed:    true,
															Description: `status of the condition, one of True, False, Unknown.`,
														},
														"type": schema.StringAttribute{
															Computed:    true,
															Description: `type of condition in CamelCase or in foo.example.com/CamelCase.`,
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
								"tls": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"status": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"vips": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
							Description: `Status is the current status of the Kuma MeshService resource.`,
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

func (r *MeshServiceListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshServiceListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshServiceListDataSourceModel
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
	var filter *operations.GetMeshServiceListQueryParamFilter
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
	filter = &operations.GetMeshServiceListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	var mesh string
	mesh = data.Mesh.ValueString()

	request := operations.GetMeshServiceListRequest{
		CpID:   cpID,
		Offset: offset,
		Size:   size,
		Filter: filter,
		Mesh:   mesh,
	}
	res, err := r.client.MeshService.GetMeshServiceList(ctx, request)
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
	if !(res.MeshServiceList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshServiceList(res.MeshServiceList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
