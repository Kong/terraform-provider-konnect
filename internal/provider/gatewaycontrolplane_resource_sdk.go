// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayControlPlaneResourceModel) ToSharedCreateControlPlaneRequest() *shared.CreateControlPlaneRequest {
	name := r.Name.ValueString()
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	clusterType := new(shared.ClusterType)
	if !r.ClusterType.IsUnknown() && !r.ClusterType.IsNull() {
		*clusterType = shared.ClusterType(r.ClusterType.ValueString())
	} else {
		clusterType = nil
	}
	authType := new(shared.AuthType)
	if !r.AuthType.IsUnknown() && !r.AuthType.IsNull() {
		*authType = shared.AuthType(r.AuthType.ValueString())
	} else {
		authType = nil
	}
	cloudGateway := new(bool)
	if !r.CloudGateway.IsUnknown() && !r.CloudGateway.IsNull() {
		*cloudGateway = r.CloudGateway.ValueBool()
	} else {
		cloudGateway = nil
	}
	var proxyUrls []shared.ProxyURL = []shared.ProxyURL{}
	for _, proxyUrlsItem := range r.ProxyUrls {
		host := proxyUrlsItem.Host.ValueString()
		port := proxyUrlsItem.Port.ValueInt64()
		protocol := proxyUrlsItem.Protocol.ValueString()
		proxyUrls = append(proxyUrls, shared.ProxyURL{
			Host:     host,
			Port:     port,
			Protocol: protocol,
		})
	}
	var labels interface{}
	if !r.Labels.IsUnknown() && !r.Labels.IsNull() {
		_ = json.Unmarshal([]byte(r.Labels.ValueString()), &labels)
	}
	out := shared.CreateControlPlaneRequest{
		Name:         name,
		Description:  description,
		ClusterType:  clusterType,
		AuthType:     authType,
		CloudGateway: cloudGateway,
		ProxyUrls:    proxyUrls,
		Labels:       labels,
	}
	return &out
}

func (r *GatewayControlPlaneResourceModel) RefreshFromSharedControlPlane(resp *shared.ControlPlane) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.Config{}
			r.Config.ControlPlaneEndpoint = types.StringPointerValue(resp.Config.ControlPlaneEndpoint)
			r.Config.TelemetryEndpoint = types.StringPointerValue(resp.Config.TelemetryEndpoint)
		}
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		if resp.Labels == nil {
			r.Labels = types.StringNull()
		} else {
			labelsResult, _ := json.Marshal(resp.Labels)
			r.Labels = types.StringValue(string(labelsResult))
		}
		r.Name = types.StringPointerValue(resp.Name)
	}
}

func (r *GatewayControlPlaneResourceModel) ToSharedUpdateControlPlaneRequest() *shared.UpdateControlPlaneRequest {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	authType := new(shared.UpdateControlPlaneRequestAuthType)
	if !r.AuthType.IsUnknown() && !r.AuthType.IsNull() {
		*authType = shared.UpdateControlPlaneRequestAuthType(r.AuthType.ValueString())
	} else {
		authType = nil
	}
	var proxyUrls []shared.ProxyURL = []shared.ProxyURL{}
	for _, proxyUrlsItem := range r.ProxyUrls {
		host := proxyUrlsItem.Host.ValueString()
		port := proxyUrlsItem.Port.ValueInt64()
		protocol := proxyUrlsItem.Protocol.ValueString()
		proxyUrls = append(proxyUrls, shared.ProxyURL{
			Host:     host,
			Port:     port,
			Protocol: protocol,
		})
	}
	var labels interface{}
	if !r.Labels.IsUnknown() && !r.Labels.IsNull() {
		_ = json.Unmarshal([]byte(r.Labels.ValueString()), &labels)
	}
	out := shared.UpdateControlPlaneRequest{
		Name:        name,
		Description: description,
		AuthType:    authType,
		ProxyUrls:   proxyUrls,
		Labels:      labels,
	}
	return &out
}
