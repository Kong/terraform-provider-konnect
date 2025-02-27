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
var _ datasource.DataSource = &MeshMultiZoneServiceDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshMultiZoneServiceDataSource{}

func NewMeshMultiZoneServiceDataSource() datasource.DataSource {
	return &MeshMultiZoneServiceDataSource{}
}

// MeshMultiZoneServiceDataSource is the data source implementation.
type MeshMultiZoneServiceDataSource struct {
	client *sdk.Konnect
}

// MeshMultiZoneServiceDataSourceModel describes the data model.
type MeshMultiZoneServiceDataSourceModel struct {
	CpID             types.String                            `tfsdk:"cp_id"`
	CreationTime     types.String                            `tfsdk:"creation_time"`
	Labels           map[string]types.String                 `tfsdk:"labels"`
	Mesh             types.String                            `tfsdk:"mesh"`
	ModificationTime types.String                            `tfsdk:"modification_time"`
	Name             types.String                            `tfsdk:"name"`
	Spec             tfTypes.MeshMultiZoneServiceItemSpec    `tfsdk:"spec"`
	Status           *tfTypes.MeshMultiZoneServiceItemStatus `tfsdk:"status"`
	Type             types.String                            `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *MeshMultiZoneServiceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_multi_zone_service"
}

// Schema defines the schema for the data source.
func (r *MeshMultiZoneServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshMultiZoneService DataSource",

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
				Description: `name of the MeshMultiZoneService`,
			},
			"spec": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
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
							},
						},
						Description: `Ports is a list of ports from selected MeshServices`,
					},
					"selector": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"mesh_service": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Computed:    true,
										ElementType: types.StringType,
										Description: `MatchLabels matches multiple MeshServices by labels`,
									},
								},
								Description: `MeshService selects MeshServices`,
							},
						},
						Description: `Selector is a way to select multiple MeshServices`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshMultiZoneService resource.`,
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
						Description: `Addresses is a list of addresses generated by HostnameGenerator`,
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
						Description: `Status of hostnames generator applied on this resource`,
					},
					"mesh_services": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mesh": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Description: `Name is a core name of MeshService`,
								},
								"namespace": schema.StringAttribute{
									Computed: true,
								},
								"zone": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Description: `MeshServices is a list of matched MeshServices`,
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
						Description: `VIPs is a list of assigned Kuma VIPs.`,
					},
				},
				Description: `Status is the current status of the Kuma MeshMultiZoneService resource.`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `the type of the resource`,
			},
		},
	}
}

func (r *MeshMultiZoneServiceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshMultiZoneServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshMultiZoneServiceDataSourceModel
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

	request := operations.GetMeshMultiZoneServiceRequest{
		CpID: cpID,
		Mesh: mesh,
		Name: name,
	}
	res, err := r.client.MeshMultiZoneService.GetMeshMultiZoneService(ctx, request)
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
	if !(res.MeshMultiZoneServiceItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMultiZoneServiceItem(res.MeshMultiZoneServiceItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
