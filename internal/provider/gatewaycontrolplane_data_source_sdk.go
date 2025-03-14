// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayControlPlaneDataSourceModel) RefreshFromSharedControlPlane(resp *shared.ControlPlane) {
	if resp != nil {
		r.Config.AuthType = types.StringValue(string(resp.Config.AuthType))
		r.Config.CloudGateway = types.BoolValue(resp.Config.CloudGateway)
		r.Config.ClusterType = types.StringValue(string(resp.Config.ClusterType))
		r.Config.ControlPlaneEndpoint = types.StringValue(resp.Config.ControlPlaneEndpoint)
		r.Config.ProxyUrls = []tfTypes.ProxyURL{}
		if len(r.Config.ProxyUrls) > len(resp.Config.ProxyUrls) {
			r.Config.ProxyUrls = r.Config.ProxyUrls[:len(resp.Config.ProxyUrls)]
		}
		for proxyUrlsCount, proxyUrlsItem := range resp.Config.ProxyUrls {
			var proxyUrls1 tfTypes.ProxyURL
			proxyUrls1.Host = types.StringValue(proxyUrlsItem.Host)
			proxyUrls1.Port = types.Int64Value(proxyUrlsItem.Port)
			proxyUrls1.Protocol = types.StringValue(proxyUrlsItem.Protocol)
			if proxyUrlsCount+1 > len(r.Config.ProxyUrls) {
				r.Config.ProxyUrls = append(r.Config.ProxyUrls, proxyUrls1)
			} else {
				r.Config.ProxyUrls[proxyUrlsCount].Host = proxyUrls1.Host
				r.Config.ProxyUrls[proxyUrlsCount].Port = proxyUrls1.Port
				r.Config.ProxyUrls[proxyUrlsCount].Protocol = proxyUrls1.Protocol
			}
		}
		r.Config.TelemetryEndpoint = types.StringValue(resp.Config.TelemetryEndpoint)
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringValue(resp.ID)
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringPointerValue(value)
			}
		}
		r.Name = types.StringValue(resp.Name)
	}
}
