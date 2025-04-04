// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayUpstreamResource{}
var _ resource.ResourceWithImportState = &GatewayUpstreamResource{}

func NewGatewayUpstreamResource() resource.Resource {
	return &GatewayUpstreamResource{}
}

// GatewayUpstreamResource defines the resource implementation.
type GatewayUpstreamResource struct {
	client *sdk.Konnect
}

// GatewayUpstreamResourceModel describes the resource data model.
type GatewayUpstreamResourceModel struct {
	Algorithm              types.String                       `tfsdk:"algorithm"`
	ClientCertificate      *tfTypes.ACLWithoutParentsConsumer `tfsdk:"client_certificate"`
	ControlPlaneID         types.String                       `tfsdk:"control_plane_id"`
	CreatedAt              types.Int64                        `tfsdk:"created_at"`
	HashFallback           types.String                       `tfsdk:"hash_fallback"`
	HashFallbackHeader     types.String                       `tfsdk:"hash_fallback_header"`
	HashFallbackQueryArg   types.String                       `tfsdk:"hash_fallback_query_arg"`
	HashFallbackURICapture types.String                       `tfsdk:"hash_fallback_uri_capture"`
	HashOn                 types.String                       `tfsdk:"hash_on"`
	HashOnCookie           types.String                       `tfsdk:"hash_on_cookie"`
	HashOnCookiePath       types.String                       `tfsdk:"hash_on_cookie_path"`
	HashOnHeader           types.String                       `tfsdk:"hash_on_header"`
	HashOnQueryArg         types.String                       `tfsdk:"hash_on_query_arg"`
	HashOnURICapture       types.String                       `tfsdk:"hash_on_uri_capture"`
	Healthchecks           *tfTypes.Healthchecks              `tfsdk:"healthchecks"`
	HostHeader             types.String                       `tfsdk:"host_header"`
	ID                     types.String                       `tfsdk:"id"`
	Name                   types.String                       `tfsdk:"name"`
	Slots                  types.Int64                        `tfsdk:"slots"`
	Tags                   []types.String                     `tfsdk:"tags"`
	UpdatedAt              types.Int64                        `tfsdk:"updated_at"`
	UseSrvName             types.Bool                         `tfsdk:"use_srv_name"`
}

func (r *GatewayUpstreamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_upstream"
}

func (r *GatewayUpstreamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayUpstream Resource",
		Attributes: map[string]schema.Attribute{
			"algorithm": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Which load balancing algorithm to use. must be one of ["consistent-hashing", "least-connections", "round-robin", "latency"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"consistent-hashing",
						"least-connections",
						"round-robin",
						"latency",
					),
				},
			},
			"client_certificate": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{
					"id": types.StringType,
				})),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.`,
			},
			"control_plane_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"hash_fallback": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `What to use as hashing input if the primary ` + "`" + `hash_on` + "`" + ` does not return a hash (eg. header is missing, or no Consumer identified). Not available if ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `cookie` + "`" + `. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query_arg", "uri_capture"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"none",
						"consumer",
						"ip",
						"header",
						"cookie",
						"path",
						"query_arg",
						"uri_capture",
					),
				},
			},
			"hash_fallback_header": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The header name to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `header` + "`" + `.`,
			},
			"hash_fallback_query_arg": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The name of the query string argument to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `query_arg` + "`" + `.`,
			},
			"hash_fallback_uri_capture": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The name of the route URI capture to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `uri_capture` + "`" + `.`,
			},
			"hash_on": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `What to use as hashing input. Using ` + "`" + `none` + "`" + ` results in a weighted-round-robin scheme with no hashing. must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query_arg", "uri_capture"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"none",
						"consumer",
						"ip",
						"header",
						"cookie",
						"path",
						"query_arg",
						"uri_capture",
					),
				},
			},
			"hash_on_cookie": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The cookie name to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` or ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `cookie` + "`" + `. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.`,
			},
			"hash_on_cookie_path": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The cookie path to set in the response headers. Only required when ` + "`" + `hash_on` + "`" + ` or ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `cookie` + "`" + `.`,
			},
			"hash_on_header": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The header name to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `header` + "`" + `.`,
			},
			"hash_on_query_arg": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The name of the query string argument to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `query_arg` + "`" + `.`,
			},
			"hash_on_uri_capture": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The name of the route URI capture to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `uri_capture` + "`" + `.`,
			},
			"healthchecks": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"concurrency": schema.Int64Attribute{
								Computed: true,
								Optional: true,
							},
							"headers": schema.MapAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"healthy": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										Optional:    true,
										ElementType: types.Int64Type,
									},
									"interval": schema.NumberAttribute{
										Computed: true,
										Optional: true,
									},
									"successes": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
								},
							},
							"http_path": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"https_sni": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"https_verify_certificate": schema.BoolAttribute{
								Computed: true,
								Optional: true,
							},
							"timeout": schema.NumberAttribute{
								Computed: true,
								Optional: true,
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `must be one of ["tcp", "http", "https", "grpc", "grpcs"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"tcp",
										"http",
										"https",
										"grpc",
										"grpcs",
									),
								},
							},
							"unhealthy": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_failures": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										Optional:    true,
										ElementType: types.Int64Type,
									},
									"interval": schema.NumberAttribute{
										Computed: true,
										Optional: true,
									},
									"tcp_failures": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"timeouts": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
								},
							},
						},
					},
					"passive": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"healthy": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										Optional:    true,
										ElementType: types.Int64Type,
									},
									"successes": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
								},
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `must be one of ["tcp", "http", "https", "grpc", "grpcs"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"tcp",
										"http",
										"https",
										"grpc",
										"grpcs",
									),
								},
							},
							"unhealthy": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_failures": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"http_statuses": schema.ListAttribute{
										Computed:    true,
										Optional:    true,
										ElementType: types.Int64Type,
									},
									"tcp_failures": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"timeouts": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
								},
							},
						},
					},
					"threshold": schema.NumberAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"host_header": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The hostname to be used as ` + "`" + `Host` + "`" + ` header when proxying requests through Kong.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `This is a hostname, which must be equal to the ` + "`" + `host` + "`" + ` of a Service.`,
			},
			"slots": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `The number of slots in the load balancer algorithm. If ` + "`" + `algorithm` + "`" + ` is set to ` + "`" + `round-robin` + "`" + `, this setting determines the maximum number of slots. If ` + "`" + `algorithm` + "`" + ` is set to ` + "`" + `consistent-hashing` + "`" + `, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range ` + "`" + `10` + "`" + `-` + "`" + `65536` + "`" + `.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Upstream for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Optional:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
			"use_srv_name": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream ` + "`" + `Host` + "`" + `.`,
			},
		},
	}
}

func (r *GatewayUpstreamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayUpstreamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayUpstreamResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	upstream := *data.ToSharedUpstream()
	request := operations.CreateUpstreamRequest{
		ControlPlaneID: controlPlaneID,
		Upstream:       upstream,
	}
	res, err := r.client.Upstreams.CreateUpstream(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Upstream != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpstream(res.Upstream)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayUpstreamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayUpstreamResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var upstreamID string
	upstreamID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetUpstreamRequest{
		UpstreamID:     upstreamID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Upstreams.GetUpstream(ctx, request)
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
	if !(res.Upstream != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpstream(res.Upstream)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayUpstreamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayUpstreamResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var upstreamID string
	upstreamID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	upstream := *data.ToSharedUpstream()
	request := operations.UpsertUpstreamRequest{
		UpstreamID:     upstreamID,
		ControlPlaneID: controlPlaneID,
		Upstream:       upstream,
	}
	res, err := r.client.Upstreams.UpsertUpstream(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Upstream != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpstream(res.Upstream)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayUpstreamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayUpstreamResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	var upstreamID string
	upstreamID = data.ID.ValueString()

	request := operations.DeleteUpstreamRequest{
		ControlPlaneID: controlPlaneID,
		UpstreamID:     upstreamID,
	}
	res, err := r.client.Upstreams.DeleteUpstream(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GatewayUpstreamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "upstream_id": "426d620c-7058-4ae6-aacc-f85a3204a2c5"}': `+err.Error())
		return
	}

	if len(data.ControlPlaneID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field control_plane_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9524ec7d-36d9-465d-a8c5-83a3c9390458"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("control_plane_id"), data.ControlPlaneID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"426d620c-7058-4ae6-aacc-f85a3204a2c5"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
