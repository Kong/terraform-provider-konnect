// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayRouteDataSourceModel) RefreshFromSharedRoute(resp *shared.Route) {
	if resp != nil {
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Destinations = []tfTypes.ClusterNodes{}
		if len(r.Destinations) > len(resp.Destinations) {
			r.Destinations = r.Destinations[:len(resp.Destinations)]
		}
		for destinationsCount, destinationsItem := range resp.Destinations {
			var destinations1 tfTypes.ClusterNodes
			destinations1.IP = types.StringPointerValue(destinationsItem.IP)
			destinations1.Port = types.Int64PointerValue(destinationsItem.Port)
			if destinationsCount+1 > len(r.Destinations) {
				r.Destinations = append(r.Destinations, destinations1)
			} else {
				r.Destinations[destinationsCount].IP = destinations1.IP
				r.Destinations[destinationsCount].Port = destinations1.Port
			}
		}
		if len(resp.Headers) > 0 {
			r.Headers = make(map[string]types.String)
			for key, value := range resp.Headers {
				r.Headers[key] = types.StringValue(value)
			}
		}
		r.Hosts = []types.String{}
		for _, v := range resp.Hosts {
			r.Hosts = append(r.Hosts, types.StringValue(v))
		}
		if resp.HTTPSRedirectStatusCode != nil {
			r.HTTPSRedirectStatusCode = types.Int64Value(int64(*resp.HTTPSRedirectStatusCode))
		} else {
			r.HTTPSRedirectStatusCode = types.Int64Null()
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Methods = []types.String{}
		for _, v := range resp.Methods {
			r.Methods = append(r.Methods, types.StringValue(v))
		}
		r.Name = types.StringPointerValue(resp.Name)
		if resp.PathHandling != nil {
			r.PathHandling = types.StringValue(string(*resp.PathHandling))
		} else {
			r.PathHandling = types.StringNull()
		}
		r.Paths = []types.String{}
		for _, v := range resp.Paths {
			r.Paths = append(r.Paths, types.StringValue(v))
		}
		r.PreserveHost = types.BoolPointerValue(resp.PreserveHost)
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		r.RegexPriority = types.Int64PointerValue(resp.RegexPriority)
		r.RequestBuffering = types.BoolPointerValue(resp.RequestBuffering)
		r.ResponseBuffering = types.BoolPointerValue(resp.ResponseBuffering)
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Snis = []types.String{}
		for _, v := range resp.Snis {
			r.Snis = append(r.Snis, types.StringValue(v))
		}
		r.Sources = []tfTypes.ClusterNodes{}
		if len(r.Sources) > len(resp.Sources) {
			r.Sources = r.Sources[:len(resp.Sources)]
		}
		for sourcesCount, sourcesItem := range resp.Sources {
			var sources1 tfTypes.ClusterNodes
			sources1.IP = types.StringPointerValue(sourcesItem.IP)
			sources1.Port = types.Int64PointerValue(sourcesItem.Port)
			if sourcesCount+1 > len(r.Sources) {
				r.Sources = append(r.Sources, sources1)
			} else {
				r.Sources[sourcesCount].IP = sources1.IP
				r.Sources[sourcesCount].Port = sources1.Port
			}
		}
		r.StripPath = types.BoolPointerValue(resp.StripPath)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
