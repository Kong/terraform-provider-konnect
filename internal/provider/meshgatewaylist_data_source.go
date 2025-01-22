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
var _ datasource.DataSource = &MeshGatewayListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshGatewayListDataSource{}

func NewMeshGatewayListDataSource() datasource.DataSource {
	return &MeshGatewayListDataSource{}
}

// MeshGatewayListDataSource is the data source implementation.
type MeshGatewayListDataSource struct {
	client *sdk.Konnect
}

// MeshGatewayListDataSourceModel describes the data model.
type MeshGatewayListDataSourceModel struct {
	CpID  types.String              `tfsdk:"cp_id"`
	Items []tfTypes.MeshGatewayItem `tfsdk:"items"`
	Mesh  types.String              `tfsdk:"mesh"`
	Next  types.String              `tfsdk:"next"`
	Total types.Number              `tfsdk:"total"`
}

// Metadata returns the data source type name.
func (r *MeshGatewayListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_gateway_list"
}

// Schema defines the schema for the data source.
func (r *MeshGatewayListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshGatewayList DataSource",

		Attributes: map[string]schema.Attribute{
			"cp_id": schema.StringAttribute{
				Required:    true,
				Description: `Id of the Konnect resource`,
			},
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"conf": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"listeners": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cross_mesh": schema.BoolAttribute{
												Computed: true,
												MarkdownDescription: `CrossMesh enables traffic to flow to this listener only from other` + "\n" +
													`meshes.`,
											},
											"hostname": schema.StringAttribute{
												Computed: true,
												MarkdownDescription: `Hostname specifies the virtual hostname to match for protocol types that` + "\n" +
													`define this concept. When unspecified, "", or ` + "`" + `*` + "`" + `, all hostnames are` + "\n" +
													`matched. This field can be omitted for protocols that don't require` + "\n" +
													`hostname based matching.`,
											},
											"port": schema.Int64Attribute{
												Computed: true,
												MarkdownDescription: `Port is the network port. Multiple listeners may use the` + "\n" +
													`same port, subject to the Listener compatibility rules.`,
											},
											"protocol": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"integer": schema.Int64Attribute{
														Computed: true,
													},
													"str": schema.StringAttribute{
														Computed: true,
													},
												},
												Description: `Protocol specifies the network protocol this listener expects to receive.`,
											},
											"resources": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"connection_limit": schema.Int64Attribute{
														Computed: true,
													},
												},
												Description: `Resources is used to specify listener-specific resource settings.`,
											},
											"tags": schema.MapAttribute{
												Computed:    true,
												ElementType: types.StringType,
												MarkdownDescription: `Tags specifies a unique combination of tags that routes can use` + "\n" +
													`to match themselves to this listener.` + "\n" +
													`` + "\n" +
													`When matching routes to listeners, the control plane constructs a` + "\n" +
													`set of matching tags for each listener by forming the union of the` + "\n" +
													`gateway tags and the listener tags. A route will be attached to the` + "\n" +
													`listener if all of the route's tags are preset in the matching tags`,
											},
											"tls": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"certificates": schema.ListNestedAttribute{
														Computed: true,
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"type": schema.StringAttribute{
																	Computed: true,
																	MarkdownDescription: `Types that are assignable to Type:` + "\n" +
																		`` + "\n" +
																		`	*DataSource_Secret` + "\n" +
																		`	*DataSource_File` + "\n" +
																		`	*DataSource_Inline` + "\n" +
																		`	*DataSource_InlineString` + "\n" +
																		`Parsed as JSON.`,
																},
															},
														},
														MarkdownDescription: `Certificates is an array of datasources that contain TLS` + "\n" +
															`certificates and private keys.  Each datasource must contain a` + "\n" +
															`sequence of PEM-encoded objects. The server certificate and private` + "\n" +
															`key are required, but additional certificates are allowed and will` + "\n" +
															`be added to the certificate chain.  The server certificate must` + "\n" +
															`be the first certificate in the datasource.` + "\n" +
															`` + "\n" +
															`When multiple certificate datasources are configured, they must have` + "\n" +
															`different key types. In practice, this means that one datasource` + "\n" +
															`should contain an RSA key and certificate, and the other an` + "\n" +
															`ECDSA key and certificate.`,
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
														MarkdownDescription: `Mode defines the TLS behavior for the TLS session initiated` + "\n" +
															`by the client.`,
													},
													"options": schema.SingleNestedAttribute{
														Computed: true,
														MarkdownDescription: `Options should eventually configure how TLS is configured. This` + "\n" +
															`is where cipher suite and version configuration can be specified,` + "\n" +
															`client certificates enforced, and so on.`,
													},
												},
												MarkdownDescription: `TLS is the TLS configuration for the Listener. This field` + "\n" +
													`is required if the Protocol field is "HTTPS" or "TLS" and` + "\n" +
													`ignored otherwise.`,
											},
										},
									},
									MarkdownDescription: `Listeners define logical endpoints that are bound on this MeshGateway's` + "\n" +
										`address(es).`,
								},
							},
							Description: `The desired configuration of the MeshGateway.`,
						},
						"labels": schema.MapAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"mesh": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"selectors": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"match": schema.MapAttribute{
										Computed:    true,
										ElementType: types.StringType,
										Description: `Tags to match, can be used for both source and destinations`,
									},
								},
							},
							MarkdownDescription: `Selectors is a list of selectors that are used to match builtin` + "\n" +
								`gateway dataplanes that will receive this MeshGateway configuration.`,
						},
						"tags": schema.MapAttribute{
							Computed:    true,
							ElementType: types.StringType,
							MarkdownDescription: `Tags is the set of tags common to all of the gateway's listeners.` + "\n" +
								`` + "\n" +
								`This field must not include a ` + "`" + `kuma.io/service` + "`" + ` tag (the service is always` + "\n" +
								`defined on the dataplanes).`,
						},
						"type": schema.StringAttribute{
							Computed: true,
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

func (r *MeshGatewayListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshGatewayListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshGatewayListDataSourceModel
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

	request := operations.GetMeshGatewayListRequest{
		CpID: cpID,
		Mesh: mesh,
	}
	res, err := r.client.MeshGateway.GetMeshGatewayList(ctx, request)
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
	if !(res.MeshGatewayList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshGatewayList(res.MeshGatewayList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
