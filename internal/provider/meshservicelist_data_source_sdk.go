// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshServiceListDataSourceModel) RefreshFromSharedMeshServiceList(resp *shared.MeshServiceList) {
	if resp != nil {
		r.Items = []tfTypes.MeshServiceItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshServiceItem
			if itemsItem.CreationTime != nil {
				items1.CreationTime = types.StringValue(itemsItem.CreationTime.Format(time.RFC3339Nano))
			} else {
				items1.CreationTime = types.StringNull()
			}
			if len(itemsItem.Labels) > 0 {
				items1.Labels = make(map[string]types.String)
				for key, value := range itemsItem.Labels {
					items1.Labels[key] = types.StringValue(value)
				}
			}
			items1.Mesh = types.StringPointerValue(itemsItem.Mesh)
			if itemsItem.ModificationTime != nil {
				items1.ModificationTime = types.StringValue(itemsItem.ModificationTime.Format(time.RFC3339Nano))
			} else {
				items1.ModificationTime = types.StringNull()
			}
			items1.Name = types.StringValue(itemsItem.Name)
			items1.Spec.Identities = []tfTypes.Identities{}
			for identitiesCount, identitiesItem := range itemsItem.Spec.Identities {
				var identities1 tfTypes.Identities
				identities1.Type = types.StringValue(string(identitiesItem.Type))
				identities1.Value = types.StringValue(identitiesItem.Value)
				if identitiesCount+1 > len(items1.Spec.Identities) {
					items1.Spec.Identities = append(items1.Spec.Identities, identities1)
				} else {
					items1.Spec.Identities[identitiesCount].Type = identities1.Type
					items1.Spec.Identities[identitiesCount].Value = identities1.Value
				}
			}
			items1.Spec.Ports = []tfTypes.MeshServiceItemPorts{}
			for portsCount, portsItem := range itemsItem.Spec.Ports {
				var ports1 tfTypes.MeshServiceItemPorts
				ports1.AppProtocol = types.StringPointerValue(portsItem.AppProtocol)
				ports1.Name = types.StringPointerValue(portsItem.Name)
				ports1.Port = types.Int64Value(int64(portsItem.Port))
				if portsItem.TargetPort == nil {
					ports1.TargetPort = nil
				} else {
					ports1.TargetPort = &tfTypes.Mode{}
					if portsItem.TargetPort.Integer != nil {
						ports1.TargetPort.Integer = types.Int64PointerValue(portsItem.TargetPort.Integer)
					}
					if portsItem.TargetPort.Str != nil {
						ports1.TargetPort.Str = types.StringPointerValue(portsItem.TargetPort.Str)
					}
				}
				if portsCount+1 > len(items1.Spec.Ports) {
					items1.Spec.Ports = append(items1.Spec.Ports, ports1)
				} else {
					items1.Spec.Ports[portsCount].AppProtocol = ports1.AppProtocol
					items1.Spec.Ports[portsCount].Name = ports1.Name
					items1.Spec.Ports[portsCount].Port = ports1.Port
					items1.Spec.Ports[portsCount].TargetPort = ports1.TargetPort
				}
			}
			if itemsItem.Spec.Selector == nil {
				items1.Spec.Selector = nil
			} else {
				items1.Spec.Selector = &tfTypes.MeshServiceItemSelector{}
				if itemsItem.Spec.Selector.DataplaneRef == nil {
					items1.Spec.Selector.DataplaneRef = nil
				} else {
					items1.Spec.Selector.DataplaneRef = &tfTypes.DataplaneRef{}
					items1.Spec.Selector.DataplaneRef.Name = types.StringPointerValue(itemsItem.Spec.Selector.DataplaneRef.Name)
				}
				if len(itemsItem.Spec.Selector.DataplaneTags) > 0 {
					items1.Spec.Selector.DataplaneTags = make(map[string]types.String)
					for key1, value2 := range itemsItem.Spec.Selector.DataplaneTags {
						items1.Spec.Selector.DataplaneTags[key1] = types.StringValue(value2)
					}
				}
			}
			if itemsItem.Spec.State != nil {
				items1.Spec.State = types.StringValue(string(*itemsItem.Spec.State))
			} else {
				items1.Spec.State = types.StringNull()
			}
			if itemsItem.Status == nil {
				items1.Status = nil
			} else {
				items1.Status = &tfTypes.MeshServiceItemStatus{}
				items1.Status.Addresses = []tfTypes.Addresses{}
				for addressesCount, addressesItem := range itemsItem.Status.Addresses {
					var addresses1 tfTypes.Addresses
					addresses1.Hostname = types.StringPointerValue(addressesItem.Hostname)
					if addressesItem.HostnameGeneratorRef == nil {
						addresses1.HostnameGeneratorRef = nil
					} else {
						addresses1.HostnameGeneratorRef = &tfTypes.HostnameGeneratorRef{}
						addresses1.HostnameGeneratorRef.CoreName = types.StringValue(addressesItem.HostnameGeneratorRef.CoreName)
					}
					addresses1.Origin = types.StringPointerValue(addressesItem.Origin)
					if addressesCount+1 > len(items1.Status.Addresses) {
						items1.Status.Addresses = append(items1.Status.Addresses, addresses1)
					} else {
						items1.Status.Addresses[addressesCount].Hostname = addresses1.Hostname
						items1.Status.Addresses[addressesCount].HostnameGeneratorRef = addresses1.HostnameGeneratorRef
						items1.Status.Addresses[addressesCount].Origin = addresses1.Origin
					}
				}
				if itemsItem.Status.DataplaneProxies == nil {
					items1.Status.DataplaneProxies = nil
				} else {
					items1.Status.DataplaneProxies = &tfTypes.DataplaneProxies{}
					items1.Status.DataplaneProxies.Connected = types.Int64PointerValue(itemsItem.Status.DataplaneProxies.Connected)
					items1.Status.DataplaneProxies.Healthy = types.Int64PointerValue(itemsItem.Status.DataplaneProxies.Healthy)
					items1.Status.DataplaneProxies.Total = types.Int64PointerValue(itemsItem.Status.DataplaneProxies.Total)
				}
				items1.Status.HostnameGenerators = []tfTypes.HostnameGenerators{}
				for hostnameGeneratorsCount, hostnameGeneratorsItem := range itemsItem.Status.HostnameGenerators {
					var hostnameGenerators1 tfTypes.HostnameGenerators
					hostnameGenerators1.Conditions = []tfTypes.Conditions{}
					for conditionsCount, conditionsItem := range hostnameGeneratorsItem.Conditions {
						var conditions1 tfTypes.Conditions
						conditions1.Message = types.StringValue(conditionsItem.Message)
						conditions1.Reason = types.StringValue(conditionsItem.Reason)
						conditions1.Status = types.StringValue(string(conditionsItem.Status))
						conditions1.Type = types.StringValue(conditionsItem.Type)
						if conditionsCount+1 > len(hostnameGenerators1.Conditions) {
							hostnameGenerators1.Conditions = append(hostnameGenerators1.Conditions, conditions1)
						} else {
							hostnameGenerators1.Conditions[conditionsCount].Message = conditions1.Message
							hostnameGenerators1.Conditions[conditionsCount].Reason = conditions1.Reason
							hostnameGenerators1.Conditions[conditionsCount].Status = conditions1.Status
							hostnameGenerators1.Conditions[conditionsCount].Type = conditions1.Type
						}
					}
					hostnameGenerators1.HostnameGeneratorRef.CoreName = types.StringValue(hostnameGeneratorsItem.HostnameGeneratorRef.CoreName)
					if hostnameGeneratorsCount+1 > len(items1.Status.HostnameGenerators) {
						items1.Status.HostnameGenerators = append(items1.Status.HostnameGenerators, hostnameGenerators1)
					} else {
						items1.Status.HostnameGenerators[hostnameGeneratorsCount].Conditions = hostnameGenerators1.Conditions
						items1.Status.HostnameGenerators[hostnameGeneratorsCount].HostnameGeneratorRef = hostnameGenerators1.HostnameGeneratorRef
					}
				}
				if itemsItem.Status.TLS == nil {
					items1.Status.TLS = nil
				} else {
					items1.Status.TLS = &tfTypes.MeshServiceItemTLS{}
					if itemsItem.Status.TLS.Status != nil {
						items1.Status.TLS.Status = types.StringValue(string(*itemsItem.Status.TLS.Status))
					} else {
						items1.Status.TLS.Status = types.StringNull()
					}
				}
				items1.Status.Vips = []tfTypes.Vip{}
				for vipsCount, vipsItem := range itemsItem.Status.Vips {
					var vips1 tfTypes.Vip
					vips1.IP = types.StringPointerValue(vipsItem.IP)
					if vipsCount+1 > len(items1.Status.Vips) {
						items1.Status.Vips = append(items1.Status.Vips, vips1)
					} else {
						items1.Status.Vips[vipsCount].IP = vips1.IP
					}
				}
			}
			items1.Type = types.StringValue(string(itemsItem.Type))
			if itemsCount+1 > len(r.Items) {
				r.Items = append(r.Items, items1)
			} else {
				r.Items[itemsCount].CreationTime = items1.CreationTime
				r.Items[itemsCount].Labels = items1.Labels
				r.Items[itemsCount].Mesh = items1.Mesh
				r.Items[itemsCount].ModificationTime = items1.ModificationTime
				r.Items[itemsCount].Name = items1.Name
				r.Items[itemsCount].Spec = items1.Spec
				r.Items[itemsCount].Status = items1.Status
				r.Items[itemsCount].Type = items1.Type
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		if resp.Total != nil {
			r.Total = types.NumberValue(big.NewFloat(float64(*resp.Total)))
		} else {
			r.Total = types.NumberNull()
		}
	}
}
