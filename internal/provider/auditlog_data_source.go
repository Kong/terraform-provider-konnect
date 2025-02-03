// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &AuditLogDataSource{}
var _ datasource.DataSourceWithConfigure = &AuditLogDataSource{}

func NewAuditLogDataSource() datasource.DataSource {
	return &AuditLogDataSource{}
}

// AuditLogDataSource is the data source implementation.
type AuditLogDataSource struct {
	client *sdk.Konnect
}

// AuditLogDataSourceModel describes the data model.
type AuditLogDataSourceModel struct {
	Enabled             types.Bool   `tfsdk:"enabled"`
	Endpoint            types.String `tfsdk:"endpoint"`
	LogFormat           types.String `tfsdk:"log_format"`
	SkipSslVerification types.Bool   `tfsdk:"skip_ssl_verification"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *AuditLogDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_audit_log"
}

// Schema defines the schema for the data source.
func (r *AuditLogDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AuditLog DataSource",

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Indicates whether audit data should be sent to the webhook.`,
			},
			"endpoint": schema.StringAttribute{
				Computed:    true,
				Description: `The endpoint that will receive audit log messages.`,
			},
			"log_format": schema.StringAttribute{
				Computed:    true,
				Description: `The output format of each log messages.`,
			},
			"skip_ssl_verification": schema.BoolAttribute{
				Computed:    true,
				Description: `Indicates if the SSL certificate verification of the host endpoint should be skipped when delivering payloads.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `Timestamp when this webhook was last updated. Initial value is 0001-01-01T00:00:0Z.`,
			},
		},
	}
}

func (r *AuditLogDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AuditLogDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AuditLogDataSourceModel
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

	res, err := r.client.AuditLogs.GetAuditLogWebhook(ctx)
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
	if !(res.AuditLogWebhook != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAuditLogWebhook(res.AuditLogWebhook)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
