// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type TransitGatewayDNSConfig struct {
	DomainProxyList            []types.String `tfsdk:"domain_proxy_list"`
	RemoteDNSServerIPAddresses []types.String `tfsdk:"remote_dns_server_ip_addresses"`
}
