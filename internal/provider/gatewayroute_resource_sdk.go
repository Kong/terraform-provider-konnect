// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayRouteResourceModel) ToSharedRouteInput() *shared.RouteInput {
	var destinations []shared.Destinations = []shared.Destinations{}
	for _, destinationsItem := range r.Destinations {
		ip := new(string)
		if !destinationsItem.IP.IsUnknown() && !destinationsItem.IP.IsNull() {
			*ip = destinationsItem.IP.ValueString()
		} else {
			ip = nil
		}
		port := new(int64)
		if !destinationsItem.Port.IsUnknown() && !destinationsItem.Port.IsNull() {
			*port = destinationsItem.Port.ValueInt64()
		} else {
			port = nil
		}
		destinations = append(destinations, shared.Destinations{
			IP:   ip,
			Port: port,
		})
	}
	headers := make(map[string]string)
	for headersKey, headersValue := range r.Headers {
		var headersInst string
		headersInst = headersValue.ValueString()

		headers[headersKey] = headersInst
	}
	var hosts []string = []string{}
	for _, hostsItem := range r.Hosts {
		hosts = append(hosts, hostsItem.ValueString())
	}
	httpsRedirectStatusCode := new(shared.HTTPSRedirectStatusCode)
	if !r.HTTPSRedirectStatusCode.IsUnknown() && !r.HTTPSRedirectStatusCode.IsNull() {
		*httpsRedirectStatusCode = shared.HTTPSRedirectStatusCode(r.HTTPSRedirectStatusCode.ValueInt64())
	} else {
		httpsRedirectStatusCode = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var methods []string = []string{}
	for _, methodsItem := range r.Methods {
		methods = append(methods, methodsItem.ValueString())
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	pathHandling := new(shared.PathHandling)
	if !r.PathHandling.IsUnknown() && !r.PathHandling.IsNull() {
		*pathHandling = shared.PathHandling(r.PathHandling.ValueString())
	} else {
		pathHandling = nil
	}
	var paths []string = []string{}
	for _, pathsItem := range r.Paths {
		paths = append(paths, pathsItem.ValueString())
	}
	preserveHost := new(bool)
	if !r.PreserveHost.IsUnknown() && !r.PreserveHost.IsNull() {
		*preserveHost = r.PreserveHost.ValueBool()
	} else {
		preserveHost = nil
	}
	var protocols []shared.RouteProtocols = []shared.RouteProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.RouteProtocols(protocolsItem.ValueString()))
	}
	regexPriority := new(int64)
	if !r.RegexPriority.IsUnknown() && !r.RegexPriority.IsNull() {
		*regexPriority = r.RegexPriority.ValueInt64()
	} else {
		regexPriority = nil
	}
	requestBuffering := new(bool)
	if !r.RequestBuffering.IsUnknown() && !r.RequestBuffering.IsNull() {
		*requestBuffering = r.RequestBuffering.ValueBool()
	} else {
		requestBuffering = nil
	}
	responseBuffering := new(bool)
	if !r.ResponseBuffering.IsUnknown() && !r.ResponseBuffering.IsNull() {
		*responseBuffering = r.ResponseBuffering.ValueBool()
	} else {
		responseBuffering = nil
	}
	var service *shared.RouteService
	if r.Service != nil {
		id1 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id1 = r.Service.ID.ValueString()
		} else {
			id1 = nil
		}
		service = &shared.RouteService{
			ID: id1,
		}
	}
	var snis []string = []string{}
	for _, snisItem := range r.Snis {
		snis = append(snis, snisItem.ValueString())
	}
	var sources []shared.Sources = []shared.Sources{}
	for _, sourcesItem := range r.Sources {
		ip1 := new(string)
		if !sourcesItem.IP.IsUnknown() && !sourcesItem.IP.IsNull() {
			*ip1 = sourcesItem.IP.ValueString()
		} else {
			ip1 = nil
		}
		port1 := new(int64)
		if !sourcesItem.Port.IsUnknown() && !sourcesItem.Port.IsNull() {
			*port1 = sourcesItem.Port.ValueInt64()
		} else {
			port1 = nil
		}
		sources = append(sources, shared.Sources{
			IP:   ip1,
			Port: port1,
		})
	}
	stripPath := new(bool)
	if !r.StripPath.IsUnknown() && !r.StripPath.IsNull() {
		*stripPath = r.StripPath.ValueBool()
	} else {
		stripPath = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.RouteInput{
		Destinations:            destinations,
		Headers:                 headers,
		Hosts:                   hosts,
		HTTPSRedirectStatusCode: httpsRedirectStatusCode,
		ID:                      id,
		Methods:                 methods,
		Name:                    name,
		PathHandling:            pathHandling,
		Paths:                   paths,
		PreserveHost:            preserveHost,
		Protocols:               protocols,
		RegexPriority:           regexPriority,
		RequestBuffering:        requestBuffering,
		ResponseBuffering:       responseBuffering,
		Service:                 service,
		Snis:                    snis,
		Sources:                 sources,
		StripPath:               stripPath,
		Tags:                    tags,
	}
	return &out
}

func (r *GatewayRouteResourceModel) RefreshFromSharedRoute(resp *shared.Route) {
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
