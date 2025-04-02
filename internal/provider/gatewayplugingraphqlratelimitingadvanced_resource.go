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
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/v2/internal/validators/objectvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginGraphqlRateLimitingAdvancedResource{}
var _ resource.ResourceWithImportState = &GatewayPluginGraphqlRateLimitingAdvancedResource{}

func NewGatewayPluginGraphqlRateLimitingAdvancedResource() resource.Resource {
	return &GatewayPluginGraphqlRateLimitingAdvancedResource{}
}

// GatewayPluginGraphqlRateLimitingAdvancedResource defines the resource implementation.
type GatewayPluginGraphqlRateLimitingAdvancedResource struct {
	client *sdk.Konnect
}

// GatewayPluginGraphqlRateLimitingAdvancedResourceModel describes the resource data model.
type GatewayPluginGraphqlRateLimitingAdvancedResourceModel struct {
	Config         *tfTypes.GraphqlRateLimitingAdvancedPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLWithoutParentsConsumer               `tfsdk:"consumer"`
	ControlPlaneID types.String                                     `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                                      `tfsdk:"created_at"`
	Enabled        types.Bool                                       `tfsdk:"enabled"`
	ID             types.String                                     `tfsdk:"id"`
	InstanceName   types.String                                     `tfsdk:"instance_name"`
	Ordering       *tfTypes.ACLPluginOrdering                       `tfsdk:"ordering"`
	Protocols      []types.String                                   `tfsdk:"protocols"`
	Route          *tfTypes.ACLWithoutParentsConsumer               `tfsdk:"route"`
	Service        *tfTypes.ACLWithoutParentsConsumer               `tfsdk:"service"`
	Tags           []types.String                                   `tfsdk:"tags"`
	UpdatedAt      types.Int64                                      `tfsdk:"updated_at"`
}

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_graphql_rate_limiting_advanced"
}

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginGraphqlRateLimitingAdvanced Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"cost_strategy": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Strategy to use to evaluate query costs. Either ` + "`" + `default` + "`" + ` or ` + "`" + `node_quantifier` + "`" + `. must be one of ["default", "node_quantifier"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"default",
								"node_quantifier",
							),
						},
					},
					"dictionary_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The shared dictionary where counters will be stored until the next sync cycle.`,
					},
					"hide_client_headers": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Optionally hide informative response headers. Available options: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
					},
					"identifier": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `How to define the rate limit key. Can be ` + "`" + `ip` + "`" + `, ` + "`" + `credential` + "`" + `, ` + "`" + `consumer` + "`" + `. must be one of ["consumer", "credential", "ip"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"consumer",
								"credential",
								"ip",
							),
						},
					},
					"limit": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.NumberType,
						Description: `One or more requests-per-window limits to apply.`,
					},
					"max_cost": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A defined maximum cost per query. 0 means unlimited.`,
					},
					"namespace": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The rate limiting namespace to use for this plugin instance. This namespace is used to share rate limiting counters across different instances. If it is not provided, a random UUID is generated. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. ` + "`" + `strategy` + "`" + `, ` + "`" + `redis` + "`" + `, ` + "`" + `sync_rate` + "`" + `, ` + "`" + `window_size` + "`" + `, ` + "`" + `dictionary_name` + "`" + `, need to be the same.`,
					},
					"redis": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"cluster_max_redirections": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Maximum retry attempts for redirection.`,
							},
							"cluster_nodes": schema.ListNestedAttribute{
								Computed: true,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `A string representing a host name, such as example.com.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Optional:    true,
											Description: `An integer representing a port number between 0 and 65535, inclusive.`,
											Validators: []validator.Int64{
												int64validator.AtMost(65535),
											},
										},
									},
								},
								Description: `Cluster addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.`,
							},
							"connect_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"connection_is_proxied": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If the connection to Redis is proxied (e.g. Envoy), set it ` + "`" + `true` + "`" + `. Set the ` + "`" + `host` + "`" + ` and ` + "`" + `port` + "`" + ` to point to the proxy address.`,
							},
							"database": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy`,
							},
							"host": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing a host name, such as example.com.`,
							},
							"keepalive_backlog": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return ` + "`" + `nil` + "`" + `. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than ` + "`" + `keepalive_pool_size` + "`" + `. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than ` + "`" + `keepalive_pool_size` + "`" + `.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"keepalive_pool_size": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither ` + "`" + `keepalive_pool_size` + "`" + ` nor ` + "`" + `keepalive_backlog` + "`" + ` is specified, no pool is created. If ` + "`" + `keepalive_pool_size` + "`" + ` isn't specified but ` + "`" + `keepalive_backlog` + "`" + ` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.`,
								Validators: []validator.Int64{
									int64validator.Between(1, 2147483646),
								},
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
							},
							"port": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a port number between 0 and 65535, inclusive.`,
								Validators: []validator.Int64{
									int64validator.AtMost(65535),
								},
							},
							"read_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"send_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"sentinel_master": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.`,
							},
							"sentinel_nodes": schema.ListNestedAttribute{
								Computed: true,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `A string representing a host name, such as example.com.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Optional:    true,
											Description: `An integer representing a port number between 0 and 65535, inclusive.`,
											Validators: []validator.Int64{
												int64validator.AtMost(65535),
											},
										},
									},
								},
								Description: `Sentinel node addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.`,
							},
							"sentinel_password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.`,
							},
							"sentinel_role": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel role to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Sentinel. must be one of ["any", "master", "slave"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"any",
										"master",
										"slave",
									),
								},
							},
							"sentinel_username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.`,
							},
							"server_name": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing an SNI (server name indication) value for TLS.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If set to true, uses SSL to connect to Redis.`,
							},
							"ssl_verify": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` in ` + "`" + `kong.conf` + "`" + ` to specify the CA (or server) certificate used by your Redis server. You may also need to configure ` + "`" + `lua_ssl_verify_depth` + "`" + ` accordingly.`,
							},
							"username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
							},
						},
					},
					"score_factor": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A scoring factor to multiply (or divide) the cost. The ` + "`" + `score_factor` + "`" + ` must always be greater than 0.`,
					},
					"strategy": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The rate-limiting strategy to use for retrieving and incrementing the limits. must be one of ["cluster", "redis"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"cluster",
								"redis",
							),
						},
					},
					"sync_rate": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 syncs the counters in that many number of seconds.`,
					},
					"window_size": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.NumberType,
						Description: `One or more window sizes to apply a limit to (defined in seconds).`,
					},
					"window_type": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Sets the time window to either ` + "`" + `sliding` + "`" + ` or ` + "`" + `fixed` + "`" + `. must be one of ["fixed", "sliding"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"fixed",
								"sliding",
							),
						},
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
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
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
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
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
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
				Optional:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginGraphqlRateLimitingAdvancedResourceModel
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

	graphqlRateLimitingAdvancedPlugin := *data.ToSharedGraphqlRateLimitingAdvancedPlugin()
	request := operations.CreateGraphqlratelimitingadvancedPluginRequest{
		ControlPlaneID:                    controlPlaneID,
		GraphqlRateLimitingAdvancedPlugin: graphqlRateLimitingAdvancedPlugin,
	}
	res, err := r.client.Plugins.CreateGraphqlratelimitingadvancedPlugin(ctx, request)
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
	if !(res.GraphqlRateLimitingAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGraphqlRateLimitingAdvancedPlugin(res.GraphqlRateLimitingAdvancedPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginGraphqlRateLimitingAdvancedResourceModel
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

	request := operations.GetGraphqlratelimitingadvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetGraphqlratelimitingadvancedPlugin(ctx, request)
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
	if !(res.GraphqlRateLimitingAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGraphqlRateLimitingAdvancedPlugin(res.GraphqlRateLimitingAdvancedPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginGraphqlRateLimitingAdvancedResourceModel
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

	graphqlRateLimitingAdvancedPlugin := *data.ToSharedGraphqlRateLimitingAdvancedPlugin()
	request := operations.UpdateGraphqlratelimitingadvancedPluginRequest{
		PluginID:                          pluginID,
		ControlPlaneID:                    controlPlaneID,
		GraphqlRateLimitingAdvancedPlugin: graphqlRateLimitingAdvancedPlugin,
	}
	res, err := r.client.Plugins.UpdateGraphqlratelimitingadvancedPlugin(ctx, request)
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
	if !(res.GraphqlRateLimitingAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGraphqlRateLimitingAdvancedPlugin(res.GraphqlRateLimitingAdvancedPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginGraphqlRateLimitingAdvancedResourceModel
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

	request := operations.DeleteGraphqlratelimitingadvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteGraphqlratelimitingadvancedPlugin(ctx, request)
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

func (r *GatewayPluginGraphqlRateLimitingAdvancedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
