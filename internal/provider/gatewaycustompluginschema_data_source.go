// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayCustomPluginSchemaDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayCustomPluginSchemaDataSource{}

func NewGatewayCustomPluginSchemaDataSource() datasource.DataSource {
	return &GatewayCustomPluginSchemaDataSource{}
}

// GatewayCustomPluginSchemaDataSource is the data source implementation.
type GatewayCustomPluginSchemaDataSource struct {
	client *sdk.Konnect
}

// GatewayCustomPluginSchemaDataSourceModel describes the data model.
type GatewayCustomPluginSchemaDataSourceModel struct {
	ControlPlaneID types.String `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64  `tfsdk:"created_at"`
	LuaSchema      types.String `tfsdk:"lua_schema"`
	Name           types.String `tfsdk:"name"`
	UpdatedAt      types.Int64  `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayCustomPluginSchemaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_custom_plugin_schema"
}

// Schema defines the schema for the data source.
func (r *GatewayCustomPluginSchemaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayCustomPluginSchema DataSource",

		Attributes: map[string]schema.Attribute{
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `An ISO-8604 timestamp representation of custom plugin schema creation date.`,
			},
			"lua_schema": schema.StringAttribute{
				Computed:    true,
				Description: `The custom plugin schema; ` + "`" + `jq -Rs '.' schema.lua` + "`" + `.`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The custom plugin name`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `An ISO-8604 timestamp representation of custom plugin schema update date.`,
			},
		},
	}
}

func (r *GatewayCustomPluginSchemaDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayCustomPluginSchemaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayCustomPluginSchemaDataSourceModel
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

	controlPlaneID := data.ControlPlaneID.ValueString()
	name := data.Name.ValueString()
	request := operations.GetPluginSchemaRequest{
		ControlPlaneID: controlPlaneID,
		Name:           name,
	}
	res, err := r.client.CustomPluginSchemas.GetPluginSchema(ctx, request)
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
	if !(res.PluginSchemas != nil && res.PluginSchemas.Item != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedItem(res.PluginSchemas.Item)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
