// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshExternalServiceListDataSourceModel) RefreshFromSharedMeshExternalServiceList(resp *shared.MeshExternalServiceList) {
	if resp != nil {
		r.Items = []tfTypes.MeshExternalServiceItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshExternalServiceItem
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
			items1.Spec.Endpoints = []tfTypes.Endpoints{}
			for endpointsCount, endpointsItem := range itemsItem.Spec.Endpoints {
				var endpoints1 tfTypes.Endpoints
				endpoints1.Address = types.StringValue(endpointsItem.Address)
				endpoints1.Port = types.Int64Value(endpointsItem.Port)
				if endpointsCount+1 > len(items1.Spec.Endpoints) {
					items1.Spec.Endpoints = append(items1.Spec.Endpoints, endpoints1)
				} else {
					items1.Spec.Endpoints[endpointsCount].Address = endpoints1.Address
					items1.Spec.Endpoints[endpointsCount].Port = endpoints1.Port
				}
			}
			if itemsItem.Spec.Extension == nil {
				items1.Spec.Extension = nil
			} else {
				items1.Spec.Extension = &tfTypes.Extension{}
				if itemsItem.Spec.Extension.Config == nil {
					items1.Spec.Extension.Config = types.StringNull()
				} else {
					configResult, _ := json.Marshal(itemsItem.Spec.Extension.Config)
					items1.Spec.Extension.Config = types.StringValue(string(configResult))
				}
				items1.Spec.Extension.Type = types.StringValue(itemsItem.Spec.Extension.Type)
			}
			items1.Spec.Match.Port = types.Int64Value(itemsItem.Spec.Match.Port)
			if itemsItem.Spec.Match.Protocol != nil {
				items1.Spec.Match.Protocol = types.StringValue(string(*itemsItem.Spec.Match.Protocol))
			} else {
				items1.Spec.Match.Protocol = types.StringNull()
			}
			if itemsItem.Spec.Match.Type != nil {
				items1.Spec.Match.Type = types.StringValue(string(*itemsItem.Spec.Match.Type))
			} else {
				items1.Spec.Match.Type = types.StringNull()
			}
			if itemsItem.Spec.TLS == nil {
				items1.Spec.TLS = nil
			} else {
				items1.Spec.TLS = &tfTypes.TLS{}
				items1.Spec.TLS.AllowRenegotiation = types.BoolPointerValue(itemsItem.Spec.TLS.AllowRenegotiation)
				items1.Spec.TLS.Enabled = types.BoolPointerValue(itemsItem.Spec.TLS.Enabled)
				if itemsItem.Spec.TLS.Verification == nil {
					items1.Spec.TLS.Verification = nil
				} else {
					items1.Spec.TLS.Verification = &tfTypes.Verification{}
					if itemsItem.Spec.TLS.Verification.CaCert == nil {
						items1.Spec.TLS.Verification.CaCert = nil
					} else {
						items1.Spec.TLS.Verification.CaCert = &tfTypes.CaCert{}
						items1.Spec.TLS.Verification.CaCert.Inline = types.StringPointerValue(itemsItem.Spec.TLS.Verification.CaCert.Inline)
						items1.Spec.TLS.Verification.CaCert.InlineString = types.StringPointerValue(itemsItem.Spec.TLS.Verification.CaCert.InlineString)
						items1.Spec.TLS.Verification.CaCert.Secret = types.StringPointerValue(itemsItem.Spec.TLS.Verification.CaCert.Secret)
					}
					if itemsItem.Spec.TLS.Verification.ClientCert == nil {
						items1.Spec.TLS.Verification.ClientCert = nil
					} else {
						items1.Spec.TLS.Verification.ClientCert = &tfTypes.CaCert{}
						items1.Spec.TLS.Verification.ClientCert.Inline = types.StringPointerValue(itemsItem.Spec.TLS.Verification.ClientCert.Inline)
						items1.Spec.TLS.Verification.ClientCert.InlineString = types.StringPointerValue(itemsItem.Spec.TLS.Verification.ClientCert.InlineString)
						items1.Spec.TLS.Verification.ClientCert.Secret = types.StringPointerValue(itemsItem.Spec.TLS.Verification.ClientCert.Secret)
					}
					if itemsItem.Spec.TLS.Verification.ClientKey == nil {
						items1.Spec.TLS.Verification.ClientKey = nil
					} else {
						items1.Spec.TLS.Verification.ClientKey = &tfTypes.CaCert{}
						items1.Spec.TLS.Verification.ClientKey.Inline = types.StringPointerValue(itemsItem.Spec.TLS.Verification.ClientKey.Inline)
						items1.Spec.TLS.Verification.ClientKey.InlineString = types.StringPointerValue(itemsItem.Spec.TLS.Verification.ClientKey.InlineString)
						items1.Spec.TLS.Verification.ClientKey.Secret = types.StringPointerValue(itemsItem.Spec.TLS.Verification.ClientKey.Secret)
					}
					if itemsItem.Spec.TLS.Verification.Mode != nil {
						items1.Spec.TLS.Verification.Mode = types.StringValue(string(*itemsItem.Spec.TLS.Verification.Mode))
					} else {
						items1.Spec.TLS.Verification.Mode = types.StringNull()
					}
					items1.Spec.TLS.Verification.ServerName = types.StringPointerValue(itemsItem.Spec.TLS.Verification.ServerName)
					items1.Spec.TLS.Verification.SubjectAltNames = []tfTypes.SubjectAltNames{}
					for subjectAltNamesCount, subjectAltNamesItem := range itemsItem.Spec.TLS.Verification.SubjectAltNames {
						var subjectAltNames1 tfTypes.SubjectAltNames
						if subjectAltNamesItem.Type != nil {
							subjectAltNames1.Type = types.StringValue(string(*subjectAltNamesItem.Type))
						} else {
							subjectAltNames1.Type = types.StringNull()
						}
						subjectAltNames1.Value = types.StringValue(subjectAltNamesItem.Value)
						if subjectAltNamesCount+1 > len(items1.Spec.TLS.Verification.SubjectAltNames) {
							items1.Spec.TLS.Verification.SubjectAltNames = append(items1.Spec.TLS.Verification.SubjectAltNames, subjectAltNames1)
						} else {
							items1.Spec.TLS.Verification.SubjectAltNames[subjectAltNamesCount].Type = subjectAltNames1.Type
							items1.Spec.TLS.Verification.SubjectAltNames[subjectAltNamesCount].Value = subjectAltNames1.Value
						}
					}
				}
				if itemsItem.Spec.TLS.Version == nil {
					items1.Spec.TLS.Version = nil
				} else {
					items1.Spec.TLS.Version = &tfTypes.MeshExternalServiceItemVersion{}
					if itemsItem.Spec.TLS.Version.Max != nil {
						items1.Spec.TLS.Version.Max = types.StringValue(string(*itemsItem.Spec.TLS.Version.Max))
					} else {
						items1.Spec.TLS.Version.Max = types.StringNull()
					}
					if itemsItem.Spec.TLS.Version.Min != nil {
						items1.Spec.TLS.Version.Min = types.StringValue(string(*itemsItem.Spec.TLS.Version.Min))
					} else {
						items1.Spec.TLS.Version.Min = types.StringNull()
					}
				}
			}
			if itemsItem.Status == nil {
				items1.Status = nil
			} else {
				items1.Status = &tfTypes.Status{}
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
				if itemsItem.Status.Vip == nil {
					items1.Status.Vip = nil
				} else {
					items1.Status.Vip = &tfTypes.Vip{}
					items1.Status.Vip.IP = types.StringPointerValue(itemsItem.Status.Vip.IP)
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
