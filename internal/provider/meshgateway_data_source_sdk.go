// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *MeshGatewayDataSourceModel) RefreshFromSharedMeshGatewayItem(resp *shared.MeshGatewayItem) {
	if resp != nil {
		if resp.Conf == nil {
			r.Conf = nil
		} else {
			r.Conf = &tfTypes.Conf{}
			r.Conf.Listeners = []tfTypes.Listeners{}
			if len(r.Conf.Listeners) > len(resp.Conf.Listeners) {
				r.Conf.Listeners = r.Conf.Listeners[:len(resp.Conf.Listeners)]
			}
			for listenersCount, listenersItem := range resp.Conf.Listeners {
				var listeners1 tfTypes.Listeners
				listeners1.CrossMesh = types.BoolPointerValue(listenersItem.CrossMesh)
				listeners1.Hostname = types.StringPointerValue(listenersItem.Hostname)
				listeners1.Port = types.Int64PointerValue(listenersItem.Port)
				if listenersItem.Protocol == nil {
					listeners1.Protocol = nil
				} else {
					listeners1.Protocol = &tfTypes.Mode{}
					if listenersItem.Protocol.Str != nil {
						listeners1.Protocol.Str = types.StringPointerValue(listenersItem.Protocol.Str)
					}
					if listenersItem.Protocol.Integer != nil {
						listeners1.Protocol.Integer = types.Int64PointerValue(listenersItem.Protocol.Integer)
					}
				}
				if listenersItem.Resources == nil {
					listeners1.Resources = nil
				} else {
					listeners1.Resources = &tfTypes.Resources{}
					listeners1.Resources.ConnectionLimit = types.Int64PointerValue(listenersItem.Resources.ConnectionLimit)
				}
				if len(listenersItem.Tags) > 0 {
					listeners1.Tags = make(map[string]types.String)
					for key, value := range listenersItem.Tags {
						listeners1.Tags[key] = types.StringValue(value)
					}
				}
				if listenersItem.TLS == nil {
					listeners1.TLS = nil
				} else {
					listeners1.TLS = &tfTypes.MeshGatewayItemTLS{}
					listeners1.TLS.Certificates = []tfTypes.Certificates{}
					for certificatesCount, certificatesItem := range listenersItem.TLS.Certificates {
						var certificates1 tfTypes.Certificates
						typeVarResult, _ := json.Marshal(certificatesItem.Type)
						certificates1.Type = types.StringValue(string(typeVarResult))
						if certificatesCount+1 > len(listeners1.TLS.Certificates) {
							listeners1.TLS.Certificates = append(listeners1.TLS.Certificates, certificates1)
						} else {
							listeners1.TLS.Certificates[certificatesCount].Type = certificates1.Type
						}
					}
					if listenersItem.TLS.Mode == nil {
						listeners1.TLS.Mode = nil
					} else {
						listeners1.TLS.Mode = &tfTypes.Mode{}
						if listenersItem.TLS.Mode.Str != nil {
							listeners1.TLS.Mode.Str = types.StringPointerValue(listenersItem.TLS.Mode.Str)
						}
						if listenersItem.TLS.Mode.Integer != nil {
							listeners1.TLS.Mode.Integer = types.Int64PointerValue(listenersItem.TLS.Mode.Integer)
						}
					}
					if listenersItem.TLS.Options == nil {
						listeners1.TLS.Options = types.StringNull()
					} else {
						optionsVarResult, _ := json.Marshal(listenersItem.TLS.Options)
						listeners1.TLS.Options = types.StringValue(string(optionsVarResult))
					}
				}
				if listenersCount+1 > len(r.Conf.Listeners) {
					r.Conf.Listeners = append(r.Conf.Listeners, listeners1)
				} else {
					r.Conf.Listeners[listenersCount].CrossMesh = listeners1.CrossMesh
					r.Conf.Listeners[listenersCount].Hostname = listeners1.Hostname
					r.Conf.Listeners[listenersCount].Port = listeners1.Port
					r.Conf.Listeners[listenersCount].Protocol = listeners1.Protocol
					r.Conf.Listeners[listenersCount].Resources = listeners1.Resources
					r.Conf.Listeners[listenersCount].Tags = listeners1.Tags
					r.Conf.Listeners[listenersCount].TLS = listeners1.TLS
				}
			}
		}
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
			for key1, value1 := range resp.Labels {
				r.Labels[key1] = types.StringValue(value1)
			}
		}
		r.Mesh = types.StringValue(resp.Mesh)
		r.Name = types.StringValue(resp.Name)
		r.Selectors = []tfTypes.Selectors{}
		if len(r.Selectors) > len(resp.Selectors) {
			r.Selectors = r.Selectors[:len(resp.Selectors)]
		}
		for selectorsCount, selectorsItem := range resp.Selectors {
			var selectors1 tfTypes.Selectors
			if len(selectorsItem.Match) > 0 {
				selectors1.Match = make(map[string]types.String)
				for key2, value2 := range selectorsItem.Match {
					selectors1.Match[key2] = types.StringValue(value2)
				}
			}
			if selectorsCount+1 > len(r.Selectors) {
				r.Selectors = append(r.Selectors, selectors1)
			} else {
				r.Selectors[selectorsCount].Match = selectors1.Match
			}
		}
		if len(resp.Tags) > 0 {
			r.Tags = make(map[string]types.String)
			for key3, value3 := range resp.Tags {
				r.Tags[key3] = types.StringValue(value3)
			}
		}
		r.Type = types.StringValue(resp.Type)
	}
}