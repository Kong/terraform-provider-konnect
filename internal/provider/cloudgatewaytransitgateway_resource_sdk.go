// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *CloudGatewayTransitGatewayResourceModel) ToSharedCreateTransitGatewayRequest() *shared.CreateTransitGatewayRequest {
	var out shared.CreateTransitGatewayRequest
	var awsTransitGateway *shared.AWSTransitGateway
	if r.AWSTransitGateway != nil {
		var name string
		name = r.AWSTransitGateway.Name.ValueString()

		var dnsConfig []shared.TransitGatewayDNSConfig = []shared.TransitGatewayDNSConfig{}
		for _, dnsConfigItem := range r.AWSTransitGateway.DNSConfig {
			var remoteDNSServerIPAddresses []string = []string{}
			for _, remoteDNSServerIPAddressesItem := range dnsConfigItem.RemoteDNSServerIPAddresses {
				remoteDNSServerIPAddresses = append(remoteDNSServerIPAddresses, remoteDNSServerIPAddressesItem.ValueString())
			}
			var domainProxyList []string = []string{}
			for _, domainProxyListItem := range dnsConfigItem.DomainProxyList {
				domainProxyList = append(domainProxyList, domainProxyListItem.ValueString())
			}
			dnsConfig = append(dnsConfig, shared.TransitGatewayDNSConfig{
				RemoteDNSServerIPAddresses: remoteDNSServerIPAddresses,
				DomainProxyList:            domainProxyList,
			})
		}
		var cidrBlocks []string = []string{}
		for _, cidrBlocksItem := range r.AWSTransitGateway.CidrBlocks {
			cidrBlocks = append(cidrBlocks, cidrBlocksItem.ValueString())
		}
		kind := shared.AWSTransitGatewayAttachmentType(r.AWSTransitGateway.TransitGatewayAttachmentConfig.Kind.ValueString())
		var transitGatewayID string
		transitGatewayID = r.AWSTransitGateway.TransitGatewayAttachmentConfig.TransitGatewayID.ValueString()

		var ramShareArn string
		ramShareArn = r.AWSTransitGateway.TransitGatewayAttachmentConfig.RAMShareArn.ValueString()

		transitGatewayAttachmentConfig := shared.AwsTransitGatewayAttachmentConfig{
			Kind:             kind,
			TransitGatewayID: transitGatewayID,
			RAMShareArn:      ramShareArn,
		}
		awsTransitGateway = &shared.AWSTransitGateway{
			Name:                           name,
			DNSConfig:                      dnsConfig,
			CidrBlocks:                     cidrBlocks,
			TransitGatewayAttachmentConfig: transitGatewayAttachmentConfig,
		}
	}
	if awsTransitGateway != nil {
		out = shared.CreateTransitGatewayRequest{
			AWSTransitGateway: awsTransitGateway,
		}
	}
	var azureTransitGateway *shared.AzureTransitGateway
	if r.AzureTransitGateway != nil {
		var name1 string
		name1 = r.AzureTransitGateway.Name.ValueString()

		var dnsConfig1 []shared.TransitGatewayDNSConfig = []shared.TransitGatewayDNSConfig{}
		for _, dnsConfigItem1 := range r.AzureTransitGateway.DNSConfig {
			var remoteDNSServerIPAddresses1 []string = []string{}
			for _, remoteDNSServerIPAddressesItem1 := range dnsConfigItem1.RemoteDNSServerIPAddresses {
				remoteDNSServerIPAddresses1 = append(remoteDNSServerIPAddresses1, remoteDNSServerIPAddressesItem1.ValueString())
			}
			var domainProxyList1 []string = []string{}
			for _, domainProxyListItem1 := range dnsConfigItem1.DomainProxyList {
				domainProxyList1 = append(domainProxyList1, domainProxyListItem1.ValueString())
			}
			dnsConfig1 = append(dnsConfig1, shared.TransitGatewayDNSConfig{
				RemoteDNSServerIPAddresses: remoteDNSServerIPAddresses1,
				DomainProxyList:            domainProxyList1,
			})
		}
		kind1 := shared.AzureVNETPeeringAttachmentType(r.AzureTransitGateway.TransitGatewayAttachmentConfig.Kind.ValueString())
		var tenantID string
		tenantID = r.AzureTransitGateway.TransitGatewayAttachmentConfig.TenantID.ValueString()

		var subscriptionID string
		subscriptionID = r.AzureTransitGateway.TransitGatewayAttachmentConfig.SubscriptionID.ValueString()

		var resourceGroupName string
		resourceGroupName = r.AzureTransitGateway.TransitGatewayAttachmentConfig.ResourceGroupName.ValueString()

		var vnetName string
		vnetName = r.AzureTransitGateway.TransitGatewayAttachmentConfig.VnetName.ValueString()

		transitGatewayAttachmentConfig1 := shared.AzureVNETPeeringAttachmentConfig{
			Kind:              kind1,
			TenantID:          tenantID,
			SubscriptionID:    subscriptionID,
			ResourceGroupName: resourceGroupName,
			VnetName:          vnetName,
		}
		azureTransitGateway = &shared.AzureTransitGateway{
			Name:                           name1,
			DNSConfig:                      dnsConfig1,
			TransitGatewayAttachmentConfig: transitGatewayAttachmentConfig1,
		}
	}
	if azureTransitGateway != nil {
		out = shared.CreateTransitGatewayRequest{
			AzureTransitGateway: azureTransitGateway,
		}
	}
	return &out
}

func (r *CloudGatewayTransitGatewayResourceModel) RefreshFromSharedTransitGatewayResponse(resp *shared.TransitGatewayResponse) {
	if resp != nil {
		if resp.AwsTransitGatewayResponse != nil {
			r.AwsTransitGatewayResponse = &tfTypes.AwsTransitGatewayResponse{}
			r.AwsTransitGatewayResponse.CidrBlocks = make([]types.String, 0, len(resp.AwsTransitGatewayResponse.CidrBlocks))
			for _, v := range resp.AwsTransitGatewayResponse.CidrBlocks {
				r.AwsTransitGatewayResponse.CidrBlocks = append(r.AwsTransitGatewayResponse.CidrBlocks, types.StringValue(v))
			}
			r.AwsTransitGatewayResponse.CreatedAt = types.StringValue(resp.AwsTransitGatewayResponse.CreatedAt.Format(time.RFC3339Nano))
			r.AwsTransitGatewayResponse.DNSConfig = []tfTypes.TransitGatewayDNSConfig{}
			if len(r.AwsTransitGatewayResponse.DNSConfig) > len(resp.AwsTransitGatewayResponse.DNSConfig) {
				r.AwsTransitGatewayResponse.DNSConfig = r.AwsTransitGatewayResponse.DNSConfig[:len(resp.AwsTransitGatewayResponse.DNSConfig)]
			}
			for dnsConfigCount, dnsConfigItem := range resp.AwsTransitGatewayResponse.DNSConfig {
				var dnsConfig1 tfTypes.TransitGatewayDNSConfig
				dnsConfig1.DomainProxyList = make([]types.String, 0, len(dnsConfigItem.DomainProxyList))
				for _, v := range dnsConfigItem.DomainProxyList {
					dnsConfig1.DomainProxyList = append(dnsConfig1.DomainProxyList, types.StringValue(v))
				}
				dnsConfig1.RemoteDNSServerIPAddresses = make([]types.String, 0, len(dnsConfigItem.RemoteDNSServerIPAddresses))
				for _, v := range dnsConfigItem.RemoteDNSServerIPAddresses {
					dnsConfig1.RemoteDNSServerIPAddresses = append(dnsConfig1.RemoteDNSServerIPAddresses, types.StringValue(v))
				}
				if dnsConfigCount+1 > len(r.AwsTransitGatewayResponse.DNSConfig) {
					r.AwsTransitGatewayResponse.DNSConfig = append(r.AwsTransitGatewayResponse.DNSConfig, dnsConfig1)
				} else {
					r.AwsTransitGatewayResponse.DNSConfig[dnsConfigCount].DomainProxyList = dnsConfig1.DomainProxyList
					r.AwsTransitGatewayResponse.DNSConfig[dnsConfigCount].RemoteDNSServerIPAddresses = dnsConfig1.RemoteDNSServerIPAddresses
				}
			}
			r.AwsTransitGatewayResponse.EntityVersion = types.Int64Value(resp.AwsTransitGatewayResponse.EntityVersion)
			r.EntityVersion = r.AwsTransitGatewayResponse.EntityVersion
			r.AwsTransitGatewayResponse.ID = types.StringValue(resp.AwsTransitGatewayResponse.ID)
			r.ID = r.AwsTransitGatewayResponse.ID
			r.AwsTransitGatewayResponse.Name = types.StringValue(resp.AwsTransitGatewayResponse.Name)
			r.Name = r.AwsTransitGatewayResponse.Name
			r.AwsTransitGatewayResponse.State = types.StringValue(string(resp.AwsTransitGatewayResponse.State))
			r.AwsTransitGatewayResponse.TransitGatewayAttachmentConfig.Kind = types.StringValue(string(resp.AwsTransitGatewayResponse.TransitGatewayAttachmentConfig.Kind))
			r.AwsTransitGatewayResponse.TransitGatewayAttachmentConfig.RAMShareArn = types.StringValue(resp.AwsTransitGatewayResponse.TransitGatewayAttachmentConfig.RAMShareArn)
			r.AwsTransitGatewayResponse.TransitGatewayAttachmentConfig.TransitGatewayID = types.StringValue(resp.AwsTransitGatewayResponse.TransitGatewayAttachmentConfig.TransitGatewayID)
			r.AwsTransitGatewayResponse.UpdatedAt = types.StringValue(resp.AwsTransitGatewayResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.AzureTransitGatewayResponse != nil {
			r.AzureTransitGatewayResponse = &tfTypes.AzureTransitGatewayResponse{}
			r.AzureTransitGatewayResponse.CreatedAt = types.StringValue(resp.AzureTransitGatewayResponse.CreatedAt.Format(time.RFC3339Nano))
			r.AzureTransitGatewayResponse.DNSConfig = []tfTypes.TransitGatewayDNSConfig{}
			if len(r.AzureTransitGatewayResponse.DNSConfig) > len(resp.AzureTransitGatewayResponse.DNSConfig) {
				r.AzureTransitGatewayResponse.DNSConfig = r.AzureTransitGatewayResponse.DNSConfig[:len(resp.AzureTransitGatewayResponse.DNSConfig)]
			}
			for dnsConfigCount1, dnsConfigItem1 := range resp.AzureTransitGatewayResponse.DNSConfig {
				var dnsConfig3 tfTypes.TransitGatewayDNSConfig
				dnsConfig3.DomainProxyList = make([]types.String, 0, len(dnsConfigItem1.DomainProxyList))
				for _, v := range dnsConfigItem1.DomainProxyList {
					dnsConfig3.DomainProxyList = append(dnsConfig3.DomainProxyList, types.StringValue(v))
				}
				dnsConfig3.RemoteDNSServerIPAddresses = make([]types.String, 0, len(dnsConfigItem1.RemoteDNSServerIPAddresses))
				for _, v := range dnsConfigItem1.RemoteDNSServerIPAddresses {
					dnsConfig3.RemoteDNSServerIPAddresses = append(dnsConfig3.RemoteDNSServerIPAddresses, types.StringValue(v))
				}
				if dnsConfigCount1+1 > len(r.AzureTransitGatewayResponse.DNSConfig) {
					r.AzureTransitGatewayResponse.DNSConfig = append(r.AzureTransitGatewayResponse.DNSConfig, dnsConfig3)
				} else {
					r.AzureTransitGatewayResponse.DNSConfig[dnsConfigCount1].DomainProxyList = dnsConfig3.DomainProxyList
					r.AzureTransitGatewayResponse.DNSConfig[dnsConfigCount1].RemoteDNSServerIPAddresses = dnsConfig3.RemoteDNSServerIPAddresses
				}
			}
			r.AzureTransitGatewayResponse.EntityVersion = types.Int64Value(resp.AzureTransitGatewayResponse.EntityVersion)
			r.EntityVersion = r.AzureTransitGatewayResponse.EntityVersion
			r.AzureTransitGatewayResponse.ID = types.StringValue(resp.AzureTransitGatewayResponse.ID)
			r.ID = r.AzureTransitGatewayResponse.ID
			r.AzureTransitGatewayResponse.Name = types.StringValue(resp.AzureTransitGatewayResponse.Name)
			r.Name = r.AzureTransitGatewayResponse.Name
			r.AzureTransitGatewayResponse.State = types.StringValue(string(resp.AzureTransitGatewayResponse.State))
			r.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.Kind = types.StringValue(string(resp.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.Kind))
			r.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.ResourceGroupName = types.StringValue(resp.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.ResourceGroupName)
			r.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.SubscriptionID = types.StringValue(resp.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.SubscriptionID)
			r.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.TenantID = types.StringValue(resp.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.TenantID)
			r.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.VnetName = types.StringValue(resp.AzureTransitGatewayResponse.TransitGatewayAttachmentConfig.VnetName)
			r.AzureTransitGatewayResponse.UpdatedAt = types.StringValue(resp.AzureTransitGatewayResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
	}
}
