// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"math/big"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginAWSLambdaResource{}
var _ resource.ResourceWithImportState = &GatewayPluginAWSLambdaResource{}

func NewGatewayPluginAWSLambdaResource() resource.Resource {
	return &GatewayPluginAWSLambdaResource{}
}

// GatewayPluginAWSLambdaResource defines the resource implementation.
type GatewayPluginAWSLambdaResource struct {
	client *sdk.Konnect
}

// GatewayPluginAWSLambdaResourceModel describes the resource data model.
type GatewayPluginAWSLambdaResourceModel struct {
	Config         tfTypes.CreateAWSLambdaPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                `tfsdk:"consumer"`
	ControlPlaneID types.String                        `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                         `tfsdk:"created_at"`
	Enabled        types.Bool                          `tfsdk:"enabled"`
	ID             types.String                        `tfsdk:"id"`
	Protocols      []types.String                      `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                `tfsdk:"service"`
	Tags           []types.String                      `tfsdk:"tags"`
}

func (r *GatewayPluginAWSLambdaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_aws_lambda"
}

func (r *GatewayPluginAWSLambdaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginAWSLambda Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"aws_assume_role_arn": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The target AWS IAM role ARN used to invoke the Lambda function.`,
					},
					"aws_imds_protocol_version": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("v1"),
						Description: `Identifier to select the IMDS protocol version to use: ` + "`" + `v1` + "`" + ` or ` + "`" + `v2` + "`" + `. must be one of ["v1", "v2"]; Default: "v1"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"v1",
								"v2",
							),
						},
					},
					"aws_key": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The AWS key credential to be used when invoking the function.`,
					},
					"aws_region": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"aws_role_session_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("kong"),
						Description: `The identifier of the assumed role session. Default: "kong"`,
					},
					"aws_secret": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The AWS secret credential to be used when invoking the function. `,
					},
					"awsgateway_compatible": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `An optional value that defines whether the plugin should wrap requests into the Amazon API gateway. Default: false`,
					},
					"base64_encode_body": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
						Description: `An optional value that Base64-encodes the request body. Default: true`,
					},
					"disable_https": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Default: false`,
					},
					"forward_request_body": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `An optional value that defines whether the request body is sent in the request_body field of the JSON-encoded request. If the body arguments can be parsed, they are sent in the separate request_body_args field of the request. . Default: false`,
					},
					"forward_request_headers": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `An optional value that defines whether the original HTTP request headers are sent as a map in the request_headers field of the JSON-encoded request. Default: false`,
					},
					"forward_request_method": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `An optional value that defines whether the original HTTP request method verb is sent in the request_method field of the JSON-encoded request. Default: false`,
					},
					"forward_request_uri": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `An optional value that defines whether the original HTTP request URI is sent in the request_uri field of the JSON-encoded request. Default: false`,
					},
					"function_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The AWS Lambda function name to invoke.`,
					},
					"host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"invocation_type": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("RequestResponse"),
						Description: `The InvocationType to use when invoking the function. Available types are RequestResponse, Event, DryRun. must be one of ["RequestResponse", "Event", "DryRun"]; Default: "RequestResponse"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"RequestResponse",
								"Event",
								"DryRun",
							),
						},
					},
					"is_proxy_integration": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `An optional value that defines whether the response format to receive from the Lambda to this format. Default: false`,
					},
					"keepalive": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(60000)),
						Description: `An optional value in milliseconds that defines how long an idle connection lives before being closed. Default: 60000`,
					},
					"log_type": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("Tail"),
						Description: `The LogType to use when invoking the function. By default, None and Tail are supported. must be one of ["Tail", "None"]; Default: "Tail"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Tail",
								"None",
							),
						},
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(443),
						Description: `An integer representing a port number between 0 and 65535, inclusive. Default: 443`,
					},
					"proxy_url": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"qualifier": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The qualifier to use when invoking the function.`,
					},
					"skip_large_bodies": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
						Description: `An optional value that defines whether Kong should send large bodies that are buffered to disk. Default: true`,
					},
					"timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(60000)),
						Description: `An optional timeout in milliseconds when invoking the function. Default: 60000`,
					},
					"unhandled_status": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The response status code to use (instead of the default 200, 202, or 204) in the case of an Unhandled Function Error.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Description: `Whether the plugin is applied. Default: true`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the Plugin to lookup`,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
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
		},
	}
}

func (r *GatewayPluginAWSLambdaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginAWSLambdaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginAWSLambdaResourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	createAWSLambdaPlugin := data.ToSharedCreateAWSLambdaPlugin()
	request := operations.CreateAwslambdaPluginRequest{
		ControlPlaneID:        controlPlaneID,
		CreateAWSLambdaPlugin: createAWSLambdaPlugin,
	}
	res, err := r.client.Plugins.CreateAwslambdaPlugin(ctx, request)
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
	if res.AWSLambdaPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAWSLambdaPlugin(res.AWSLambdaPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginAWSLambdaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginAWSLambdaResourceModel
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

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.GetAwslambdaPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetAwslambdaPlugin(ctx, request)
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
	if res.AWSLambdaPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAWSLambdaPlugin(res.AWSLambdaPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginAWSLambdaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginAWSLambdaResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	createAWSLambdaPlugin := data.ToSharedCreateAWSLambdaPlugin()
	request := operations.UpdateAwslambdaPluginRequest{
		PluginID:              pluginID,
		ControlPlaneID:        controlPlaneID,
		CreateAWSLambdaPlugin: createAWSLambdaPlugin,
	}
	res, err := r.client.Plugins.UpdateAwslambdaPlugin(ctx, request)
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
	if res.AWSLambdaPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAWSLambdaPlugin(res.AWSLambdaPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginAWSLambdaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginAWSLambdaResourceModel
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

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.DeleteAwslambdaPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteAwslambdaPlugin(ctx, request)
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

func (r *GatewayPluginAWSLambdaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_plugin_aws_lambda. Reason: composite imports strings not supported.")
}
