// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &GatewayPluginHeaderCertAuthResource{}
var _ resource.ResourceWithImportState = &GatewayPluginHeaderCertAuthResource{}

func NewGatewayPluginHeaderCertAuthResource() resource.Resource {
	return &GatewayPluginHeaderCertAuthResource{}
}

// GatewayPluginHeaderCertAuthResource defines the resource implementation.
type GatewayPluginHeaderCertAuthResource struct {
	client *sdk.Konnect
}

// GatewayPluginHeaderCertAuthResourceModel describes the resource data model.
type GatewayPluginHeaderCertAuthResourceModel struct {
	Config         tfTypes.HeaderCertAuthPluginConfig `tfsdk:"config"`
	ControlPlaneID types.String                       `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                        `tfsdk:"created_at"`
	Enabled        types.Bool                         `tfsdk:"enabled"`
	ID             types.String                       `tfsdk:"id"`
	InstanceName   types.String                       `tfsdk:"instance_name"`
	Ordering       map[string]types.String            `tfsdk:"ordering"`
	Protocols      []types.String                     `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer `tfsdk:"route"`
	Service        *tfTypes.ACLWithoutParentsConsumer `tfsdk:"service"`
	Tags           []types.String                     `tfsdk:"tags"`
	UpdatedAt      types.Int64                        `tfsdk:"updated_at"`
}

func (r *GatewayPluginHeaderCertAuthResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_header_cert_auth"
}

func (r *GatewayPluginHeaderCertAuthResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginHeaderCertAuth Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"allow_partial_chain": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Allow certificate verification with only an intermediate certificate. When this is enabled, you don't need to upload the full chain to Kong Certificates.`,
					},
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure ` + "`" + `4xx` + "`" + `. Note that this value must refer to the consumer ` + "`" + `id` + "`" + ` or ` + "`" + `username` + "`" + ` attribute, and **not** its ` + "`" + `custom_id` + "`" + `.`,
					},
					"authenticated_group_by": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Certificate property to use as the authenticated group. Valid values are ` + "`" + `CN` + "`" + ` (Common Name) or ` + "`" + `DN` + "`" + ` (Distinguished Name). Once ` + "`" + `skip_consumer_lookup` + "`" + ` is applied, any client with a valid certificate can access the Service/API. To restrict usage to only some of the authenticated users, also add the ACL plugin (not covered here) and create allowed or denied groups of users. must be one of ["CN", "DN"]`,
						Validators: []validator.String{
							stringvalidator.OneOf("CN", "DN"),
						},
					},
					"ca_certificates": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of CA Certificates strings to use as Certificate Authorities (CA) when validating a client certificate. At least one is required but you can specify as many as needed. The value of this array is comprised of primary keys (` + "`" + `id` + "`" + `).`,
					},
					"cache_ttl": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Cache expiry time in seconds.`,
					},
					"cert_cache_ttl": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The length of time in milliseconds between refreshes of the revocation check status cache.`,
					},
					"certificate_header_format": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Format of the certificate header. Supported formats: ` + "`" + `base64_encoded` + "`" + `, ` + "`" + `url_encoded` + "`" + `. must be one of ["base64_encoded", "url_encoded"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"base64_encoded",
								"url_encoded",
							),
						},
					},
					"certificate_header_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Name of the header that contains the certificate, received from the WAF or other L7 downstream proxy.`,
					},
					"consumer_by": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `Whether to match the subject name of the client-supplied certificate against consumer's ` + "`" + `username` + "`" + ` and/or ` + "`" + `custom_id` + "`" + ` attribute. If set to ` + "`" + `[]` + "`" + ` (the empty array), then auto-matching is disabled.`,
					},
					"default_consumer": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The UUID or username of the consumer to use when a trusted client certificate is presented but no consumer matches. Note that this value must refer to the consumer ` + "`" + `id` + "`" + ` or ` + "`" + `username` + "`" + ` attribute, and **not** its ` + "`" + `custom_id` + "`" + `.`,
					},
					"http_proxy_host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"http_proxy_port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
						Validators: []validator.Int64{
							int64validator.AtMost(65535),
						},
					},
					"http_timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `HTTP timeout threshold in milliseconds when communicating with the OCSP server or downloading CRL.`,
					},
					"https_proxy_host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"https_proxy_port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `An integer representing a port number between 0 and 65535, inclusive.`,
						Validators: []validator.Int64{
							int64validator.AtMost(65535),
						},
					},
					"revocation_check_mode": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Controls client certificate revocation check behavior. If set to ` + "`" + `SKIP` + "`" + `, no revocation check is performed. If set to ` + "`" + `IGNORE_CA_ERROR` + "`" + `, the plugin respects the revocation status when either OCSP or CRL URL is set, and doesn't fail on network issues. If set to ` + "`" + `STRICT` + "`" + `, the plugin only treats the certificate as valid when it's able to verify the revocation status. must be one of ["IGNORE_CA_ERROR", "SKIP", "STRICT"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"IGNORE_CA_ERROR",
								"SKIP",
								"STRICT",
							),
						},
					},
					"secure_source": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Whether to secure the source of the request. If set to ` + "`" + `true` + "`" + `, the plugin will only allow requests from trusted IPs (configured by the ` + "`" + `trusted_ips` + "`" + ` config option).`,
					},
					"skip_consumer_lookup": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Skip consumer lookup once certificate is trusted against the configured CA list.`,
					},
				},
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
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.MapAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A set of strings representing HTTP protocols.`,
			},
			"route": schema.SingleNestedAttribute{
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
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.`,
			},
			"service": schema.SingleNestedAttribute{
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
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayPluginHeaderCertAuthResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginHeaderCertAuthResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginHeaderCertAuthResourceModel
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

	headerCertAuthPlugin := *data.ToSharedHeaderCertAuthPluginInput()
	request := operations.CreateHeadercertauthPluginRequest{
		ControlPlaneID:       controlPlaneID,
		HeaderCertAuthPlugin: headerCertAuthPlugin,
	}
	res, err := r.client.Plugins.CreateHeadercertauthPlugin(ctx, request)
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
	if !(res.HeaderCertAuthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedHeaderCertAuthPlugin(res.HeaderCertAuthPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginHeaderCertAuthResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginHeaderCertAuthResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetHeadercertauthPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetHeadercertauthPlugin(ctx, request)
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
	if !(res.HeaderCertAuthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedHeaderCertAuthPlugin(res.HeaderCertAuthPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginHeaderCertAuthResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginHeaderCertAuthResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	headerCertAuthPlugin := *data.ToSharedHeaderCertAuthPluginInput()
	request := operations.UpdateHeadercertauthPluginRequest{
		PluginID:             pluginID,
		ControlPlaneID:       controlPlaneID,
		HeaderCertAuthPlugin: headerCertAuthPlugin,
	}
	res, err := r.client.Plugins.UpdateHeadercertauthPlugin(ctx, request)
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
	if !(res.HeaderCertAuthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedHeaderCertAuthPlugin(res.HeaderCertAuthPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginHeaderCertAuthResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginHeaderCertAuthResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.DeleteHeadercertauthPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteHeadercertauthPlugin(ctx, request)
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

func (r *GatewayPluginHeaderCertAuthResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "plugin_id": "3473c251-5b6c-4f45-b1ff-7ede735a366d"}': `+err.Error())
		return
	}

	if len(data.ControlPlaneID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field control_plane_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9524ec7d-36d9-465d-a8c5-83a3c9390458"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("control_plane_id"), data.ControlPlaneID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"3473c251-5b6c-4f45-b1ff-7ede735a366d"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}
