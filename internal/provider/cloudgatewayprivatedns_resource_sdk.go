// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/provider/typeconvert"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *CloudGatewayPrivateDNSResourceModel) ToSharedCreatePrivateDNSRequest(ctx context.Context) (*shared.CreatePrivateDNSRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	var privateDNSAttachmentConfig *shared.PrivateDNSAttachmentConfig
	if r.PrivateDNSAttachmentConfig != nil {
		var awsPrivateHostedZoneAttachmentConfig *shared.AwsPrivateHostedZoneAttachmentConfig
		if r.PrivateDNSAttachmentConfig.AwsPrivateHostedZoneAttachmentConfig != nil {
			kind := shared.AWSPrivateHostedZoneType(r.PrivateDNSAttachmentConfig.AwsPrivateHostedZoneAttachmentConfig.Kind.ValueString())
			var hostedZoneID string
			hostedZoneID = r.PrivateDNSAttachmentConfig.AwsPrivateHostedZoneAttachmentConfig.HostedZoneID.ValueString()

			awsPrivateHostedZoneAttachmentConfig = &shared.AwsPrivateHostedZoneAttachmentConfig{
				Kind:         kind,
				HostedZoneID: hostedZoneID,
			}
		}
		if awsPrivateHostedZoneAttachmentConfig != nil {
			privateDNSAttachmentConfig = &shared.PrivateDNSAttachmentConfig{
				AwsPrivateHostedZoneAttachmentConfig: awsPrivateHostedZoneAttachmentConfig,
			}
		}
		var awsPrivateDNSResolverAttachmentConfig *shared.AwsPrivateDNSResolverAttachmentConfig
		if r.PrivateDNSAttachmentConfig.AwsPrivateDNSResolverAttachmentConfig != nil {
			kind1 := shared.AWSPrivateDNSResolverType(r.PrivateDNSAttachmentConfig.AwsPrivateDNSResolverAttachmentConfig.Kind.ValueString())
			dnsConfig := make(map[string]shared.PrivateDNSResolverConfig)
			for dnsConfigKey, dnsConfigValue := range r.PrivateDNSAttachmentConfig.AwsPrivateDNSResolverAttachmentConfig.DNSConfig {
				remoteDNSServerIPAddresses := make([]string, 0, len(dnsConfigValue.RemoteDNSServerIPAddresses))
				for _, remoteDNSServerIPAddressesItem := range dnsConfigValue.RemoteDNSServerIPAddresses {
					remoteDNSServerIPAddresses = append(remoteDNSServerIPAddresses, remoteDNSServerIPAddressesItem.ValueString())
				}
				dnsConfigInst := shared.PrivateDNSResolverConfig{
					RemoteDNSServerIPAddresses: remoteDNSServerIPAddresses,
				}
				dnsConfig[dnsConfigKey] = dnsConfigInst
			}
			awsPrivateDNSResolverAttachmentConfig = &shared.AwsPrivateDNSResolverAttachmentConfig{
				Kind:      kind1,
				DNSConfig: dnsConfig,
			}
		}
		if awsPrivateDNSResolverAttachmentConfig != nil {
			privateDNSAttachmentConfig = &shared.PrivateDNSAttachmentConfig{
				AwsPrivateDNSResolverAttachmentConfig: awsPrivateDNSResolverAttachmentConfig,
			}
		}
	}
	out := shared.CreatePrivateDNSRequest{
		Name:                       name,
		PrivateDNSAttachmentConfig: privateDNSAttachmentConfig,
	}

	return &out, diags
}

func (r *CloudGatewayPrivateDNSResourceModel) ToOperationsCreatePrivateDNSRequest(ctx context.Context) (*operations.CreatePrivateDNSRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var networkID string
	networkID = r.NetworkID.ValueString()

	createPrivateDNSRequest, createPrivateDNSRequestDiags := r.ToSharedCreatePrivateDNSRequest(ctx)
	diags.Append(createPrivateDNSRequestDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreatePrivateDNSRequest{
		NetworkID:               networkID,
		CreatePrivateDNSRequest: *createPrivateDNSRequest,
	}

	return &out, diags
}

func (r *CloudGatewayPrivateDNSResourceModel) ToSharedPatchPrivateDNSRequest(ctx context.Context) (*shared.PatchPrivateDNSRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var out shared.PatchPrivateDNSRequest

	return &out, diags
}

func (r *CloudGatewayPrivateDNSResourceModel) ToOperationsUpdatePrivateDNSRequest(ctx context.Context) (*operations.UpdatePrivateDNSRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var networkID string
	networkID = r.NetworkID.ValueString()

	var privateDNSID string
	privateDNSID = r.ID.ValueString()

	patchPrivateDNSRequest, patchPrivateDNSRequestDiags := r.ToSharedPatchPrivateDNSRequest(ctx)
	diags.Append(patchPrivateDNSRequestDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdatePrivateDNSRequest{
		NetworkID:              networkID,
		PrivateDNSID:           privateDNSID,
		PatchPrivateDNSRequest: *patchPrivateDNSRequest,
	}

	return &out, diags
}

func (r *CloudGatewayPrivateDNSResourceModel) ToOperationsGetPrivateDNSRequest(ctx context.Context) (*operations.GetPrivateDNSRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var networkID string
	networkID = r.NetworkID.ValueString()

	var privateDNSID string
	privateDNSID = r.ID.ValueString()

	out := operations.GetPrivateDNSRequest{
		NetworkID:    networkID,
		PrivateDNSID: privateDNSID,
	}

	return &out, diags
}

func (r *CloudGatewayPrivateDNSResourceModel) ToOperationsDeletePrivateDNSRequest(ctx context.Context) (*operations.DeletePrivateDNSRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var networkID string
	networkID = r.NetworkID.ValueString()

	var privateDNSID string
	privateDNSID = r.ID.ValueString()

	out := operations.DeletePrivateDNSRequest{
		NetworkID:    networkID,
		PrivateDNSID: privateDNSID,
	}

	return &out, diags
}

func (r *CloudGatewayPrivateDNSResourceModel) RefreshFromSharedPrivateDNSResponse(ctx context.Context, resp *shared.PrivateDNSResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.AwsPrivateDNSResolverResponse != nil {
			r.AwsPrivateDNSResolverResponse = &tfTypes.AwsPrivateDNSResolverResponse{}
			r.AwsPrivateDNSResolverResponse.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.AwsPrivateDNSResolverResponse.CreatedAt))
			r.AwsPrivateDNSResolverResponse.EntityVersion = types.Int64Value(resp.AwsPrivateDNSResolverResponse.EntityVersion)
			r.EntityVersion = r.AwsPrivateDNSResolverResponse.EntityVersion
			r.AwsPrivateDNSResolverResponse.ID = types.StringValue(resp.AwsPrivateDNSResolverResponse.ID)
			r.ID = r.AwsPrivateDNSResolverResponse.ID
			r.AwsPrivateDNSResolverResponse.Name = types.StringValue(resp.AwsPrivateDNSResolverResponse.Name)
			r.Name = r.AwsPrivateDNSResolverResponse.Name
			if len(resp.AwsPrivateDNSResolverResponse.PrivateDNSAttachmentConfig.DNSConfig) > 0 {
				r.AwsPrivateDNSResolverResponse.PrivateDNSAttachmentConfig.DNSConfig = make(map[string]tfTypes.PrivateDNSResolverConfig, len(resp.AwsPrivateDNSResolverResponse.PrivateDNSAttachmentConfig.DNSConfig))
				for privateDNSResolverConfigKey, privateDNSResolverConfigValue := range resp.AwsPrivateDNSResolverResponse.PrivateDNSAttachmentConfig.DNSConfig {
					var privateDNSResolverConfigResult tfTypes.PrivateDNSResolverConfig
					privateDNSResolverConfigResult.RemoteDNSServerIPAddresses = make([]types.String, 0, len(privateDNSResolverConfigValue.RemoteDNSServerIPAddresses))
					for _, v := range privateDNSResolverConfigValue.RemoteDNSServerIPAddresses {
						privateDNSResolverConfigResult.RemoteDNSServerIPAddresses = append(privateDNSResolverConfigResult.RemoteDNSServerIPAddresses, types.StringValue(v))
					}

					r.AwsPrivateDNSResolverResponse.PrivateDNSAttachmentConfig.DNSConfig[privateDNSResolverConfigKey] = privateDNSResolverConfigResult
				}
			}
			r.AwsPrivateDNSResolverResponse.PrivateDNSAttachmentConfig.Kind = types.StringValue(string(resp.AwsPrivateDNSResolverResponse.PrivateDNSAttachmentConfig.Kind))
			r.AwsPrivateDNSResolverResponse.State = types.StringValue(string(resp.AwsPrivateDNSResolverResponse.State))
			r.AwsPrivateDNSResolverResponse.StateMetadata.Reason = types.StringPointerValue(resp.AwsPrivateDNSResolverResponse.StateMetadata.Reason)
			r.AwsPrivateDNSResolverResponse.StateMetadata.ReportedStatus = types.StringPointerValue(resp.AwsPrivateDNSResolverResponse.StateMetadata.ReportedStatus)
			r.AwsPrivateDNSResolverResponse.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.AwsPrivateDNSResolverResponse.UpdatedAt))
		}
		if resp.AwsPrivateHostedZoneResponse != nil {
			r.AwsPrivateHostedZoneResponse = &tfTypes.AwsPrivateHostedZoneResponse{}
			r.AwsPrivateHostedZoneResponse.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.AwsPrivateHostedZoneResponse.CreatedAt))
			r.AwsPrivateHostedZoneResponse.EntityVersion = types.Int64Value(resp.AwsPrivateHostedZoneResponse.EntityVersion)
			r.EntityVersion = r.AwsPrivateHostedZoneResponse.EntityVersion
			r.AwsPrivateHostedZoneResponse.ID = types.StringValue(resp.AwsPrivateHostedZoneResponse.ID)
			r.ID = r.AwsPrivateHostedZoneResponse.ID
			r.AwsPrivateHostedZoneResponse.Name = types.StringValue(resp.AwsPrivateHostedZoneResponse.Name)
			r.Name = r.AwsPrivateHostedZoneResponse.Name
			r.AwsPrivateHostedZoneResponse.PrivateDNSAttachmentConfig.HostedZoneID = types.StringValue(resp.AwsPrivateHostedZoneResponse.PrivateDNSAttachmentConfig.HostedZoneID)
			r.AwsPrivateHostedZoneResponse.PrivateDNSAttachmentConfig.Kind = types.StringValue(string(resp.AwsPrivateHostedZoneResponse.PrivateDNSAttachmentConfig.Kind))
			r.AwsPrivateHostedZoneResponse.State = types.StringValue(string(resp.AwsPrivateHostedZoneResponse.State))
			r.AwsPrivateHostedZoneResponse.StateMetadata.Reason = types.StringPointerValue(resp.AwsPrivateHostedZoneResponse.StateMetadata.Reason)
			r.AwsPrivateHostedZoneResponse.StateMetadata.ReportedStatus = types.StringPointerValue(resp.AwsPrivateHostedZoneResponse.StateMetadata.ReportedStatus)
			r.AwsPrivateHostedZoneResponse.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.AwsPrivateHostedZoneResponse.UpdatedAt))
		}
	}

	return diags
}
