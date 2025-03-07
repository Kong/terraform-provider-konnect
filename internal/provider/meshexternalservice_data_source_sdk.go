// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshExternalServiceDataSourceModel) RefreshFromSharedMeshExternalServiceItem(resp *shared.MeshExternalServiceItem) {
	if resp != nil {
		if resp.CreationTime != nil {
			r.CreationTime = types.StringValue(resp.CreationTime.Format(time.RFC3339Nano))
		} else {
			r.CreationTime = types.StringNull()
		}
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Mesh = types.StringPointerValue(resp.Mesh)
		if resp.ModificationTime != nil {
			r.ModificationTime = types.StringValue(resp.ModificationTime.Format(time.RFC3339Nano))
		} else {
			r.ModificationTime = types.StringNull()
		}
		r.Name = types.StringValue(resp.Name)
		r.Spec.Endpoints = []tfTypes.Endpoints{}
		if len(r.Spec.Endpoints) > len(resp.Spec.Endpoints) {
			r.Spec.Endpoints = r.Spec.Endpoints[:len(resp.Spec.Endpoints)]
		}
		for endpointsCount, endpointsItem := range resp.Spec.Endpoints {
			var endpoints1 tfTypes.Endpoints
			endpoints1.Address = types.StringValue(endpointsItem.Address)
			endpoints1.Port = types.Int64Value(endpointsItem.Port)
			if endpointsCount+1 > len(r.Spec.Endpoints) {
				r.Spec.Endpoints = append(r.Spec.Endpoints, endpoints1)
			} else {
				r.Spec.Endpoints[endpointsCount].Address = endpoints1.Address
				r.Spec.Endpoints[endpointsCount].Port = endpoints1.Port
			}
		}
		if resp.Spec.Extension == nil {
			r.Spec.Extension = nil
		} else {
			r.Spec.Extension = &tfTypes.Extension{}
			if resp.Spec.Extension.Config == nil {
				r.Spec.Extension.Config = types.StringNull()
			} else {
				configResult, _ := json.Marshal(resp.Spec.Extension.Config)
				r.Spec.Extension.Config = types.StringValue(string(configResult))
			}
			r.Spec.Extension.Type = types.StringValue(resp.Spec.Extension.Type)
		}
		r.Spec.Match.Port = types.Int64Value(resp.Spec.Match.Port)
		if resp.Spec.Match.Protocol != nil {
			r.Spec.Match.Protocol = types.StringValue(string(*resp.Spec.Match.Protocol))
		} else {
			r.Spec.Match.Protocol = types.StringNull()
		}
		if resp.Spec.Match.Type != nil {
			r.Spec.Match.Type = types.StringValue(string(*resp.Spec.Match.Type))
		} else {
			r.Spec.Match.Type = types.StringNull()
		}
		if resp.Spec.TLS == nil {
			r.Spec.TLS = nil
		} else {
			r.Spec.TLS = &tfTypes.TLS{}
			r.Spec.TLS.AllowRenegotiation = types.BoolPointerValue(resp.Spec.TLS.AllowRenegotiation)
			r.Spec.TLS.Enabled = types.BoolPointerValue(resp.Spec.TLS.Enabled)
			if resp.Spec.TLS.Verification == nil {
				r.Spec.TLS.Verification = nil
			} else {
				r.Spec.TLS.Verification = &tfTypes.Verification{}
				if resp.Spec.TLS.Verification.CaCert == nil {
					r.Spec.TLS.Verification.CaCert = nil
				} else {
					r.Spec.TLS.Verification.CaCert = &tfTypes.CaCert{}
					r.Spec.TLS.Verification.CaCert.Inline = types.StringPointerValue(resp.Spec.TLS.Verification.CaCert.Inline)
					r.Spec.TLS.Verification.CaCert.InlineString = types.StringPointerValue(resp.Spec.TLS.Verification.CaCert.InlineString)
					r.Spec.TLS.Verification.CaCert.Secret = types.StringPointerValue(resp.Spec.TLS.Verification.CaCert.Secret)
				}
				if resp.Spec.TLS.Verification.ClientCert == nil {
					r.Spec.TLS.Verification.ClientCert = nil
				} else {
					r.Spec.TLS.Verification.ClientCert = &tfTypes.CaCert{}
					r.Spec.TLS.Verification.ClientCert.Inline = types.StringPointerValue(resp.Spec.TLS.Verification.ClientCert.Inline)
					r.Spec.TLS.Verification.ClientCert.InlineString = types.StringPointerValue(resp.Spec.TLS.Verification.ClientCert.InlineString)
					r.Spec.TLS.Verification.ClientCert.Secret = types.StringPointerValue(resp.Spec.TLS.Verification.ClientCert.Secret)
				}
				if resp.Spec.TLS.Verification.ClientKey == nil {
					r.Spec.TLS.Verification.ClientKey = nil
				} else {
					r.Spec.TLS.Verification.ClientKey = &tfTypes.CaCert{}
					r.Spec.TLS.Verification.ClientKey.Inline = types.StringPointerValue(resp.Spec.TLS.Verification.ClientKey.Inline)
					r.Spec.TLS.Verification.ClientKey.InlineString = types.StringPointerValue(resp.Spec.TLS.Verification.ClientKey.InlineString)
					r.Spec.TLS.Verification.ClientKey.Secret = types.StringPointerValue(resp.Spec.TLS.Verification.ClientKey.Secret)
				}
				if resp.Spec.TLS.Verification.Mode != nil {
					r.Spec.TLS.Verification.Mode = types.StringValue(string(*resp.Spec.TLS.Verification.Mode))
				} else {
					r.Spec.TLS.Verification.Mode = types.StringNull()
				}
				r.Spec.TLS.Verification.ServerName = types.StringPointerValue(resp.Spec.TLS.Verification.ServerName)
				r.Spec.TLS.Verification.SubjectAltNames = []tfTypes.SubjectAltNames{}
				if len(r.Spec.TLS.Verification.SubjectAltNames) > len(resp.Spec.TLS.Verification.SubjectAltNames) {
					r.Spec.TLS.Verification.SubjectAltNames = r.Spec.TLS.Verification.SubjectAltNames[:len(resp.Spec.TLS.Verification.SubjectAltNames)]
				}
				for subjectAltNamesCount, subjectAltNamesItem := range resp.Spec.TLS.Verification.SubjectAltNames {
					var subjectAltNames1 tfTypes.SubjectAltNames
					if subjectAltNamesItem.Type != nil {
						subjectAltNames1.Type = types.StringValue(string(*subjectAltNamesItem.Type))
					} else {
						subjectAltNames1.Type = types.StringNull()
					}
					subjectAltNames1.Value = types.StringValue(subjectAltNamesItem.Value)
					if subjectAltNamesCount+1 > len(r.Spec.TLS.Verification.SubjectAltNames) {
						r.Spec.TLS.Verification.SubjectAltNames = append(r.Spec.TLS.Verification.SubjectAltNames, subjectAltNames1)
					} else {
						r.Spec.TLS.Verification.SubjectAltNames[subjectAltNamesCount].Type = subjectAltNames1.Type
						r.Spec.TLS.Verification.SubjectAltNames[subjectAltNamesCount].Value = subjectAltNames1.Value
					}
				}
			}
			if resp.Spec.TLS.Version == nil {
				r.Spec.TLS.Version = nil
			} else {
				r.Spec.TLS.Version = &tfTypes.MeshExternalServiceItemVersion{}
				if resp.Spec.TLS.Version.Max != nil {
					r.Spec.TLS.Version.Max = types.StringValue(string(*resp.Spec.TLS.Version.Max))
				} else {
					r.Spec.TLS.Version.Max = types.StringNull()
				}
				if resp.Spec.TLS.Version.Min != nil {
					r.Spec.TLS.Version.Min = types.StringValue(string(*resp.Spec.TLS.Version.Min))
				} else {
					r.Spec.TLS.Version.Min = types.StringNull()
				}
			}
		}
		if resp.Status == nil {
			r.Status = nil
		} else {
			r.Status = &tfTypes.Status{}
			r.Status.Addresses = []tfTypes.Addresses{}
			if len(r.Status.Addresses) > len(resp.Status.Addresses) {
				r.Status.Addresses = r.Status.Addresses[:len(resp.Status.Addresses)]
			}
			for addressesCount, addressesItem := range resp.Status.Addresses {
				var addresses1 tfTypes.Addresses
				addresses1.Hostname = types.StringPointerValue(addressesItem.Hostname)
				if addressesItem.HostnameGeneratorRef == nil {
					addresses1.HostnameGeneratorRef = nil
				} else {
					addresses1.HostnameGeneratorRef = &tfTypes.HostnameGeneratorRef{}
					addresses1.HostnameGeneratorRef.CoreName = types.StringValue(addressesItem.HostnameGeneratorRef.CoreName)
				}
				addresses1.Origin = types.StringPointerValue(addressesItem.Origin)
				if addressesCount+1 > len(r.Status.Addresses) {
					r.Status.Addresses = append(r.Status.Addresses, addresses1)
				} else {
					r.Status.Addresses[addressesCount].Hostname = addresses1.Hostname
					r.Status.Addresses[addressesCount].HostnameGeneratorRef = addresses1.HostnameGeneratorRef
					r.Status.Addresses[addressesCount].Origin = addresses1.Origin
				}
			}
			r.Status.HostnameGenerators = []tfTypes.HostnameGenerators{}
			if len(r.Status.HostnameGenerators) > len(resp.Status.HostnameGenerators) {
				r.Status.HostnameGenerators = r.Status.HostnameGenerators[:len(resp.Status.HostnameGenerators)]
			}
			for hostnameGeneratorsCount, hostnameGeneratorsItem := range resp.Status.HostnameGenerators {
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
				if hostnameGeneratorsCount+1 > len(r.Status.HostnameGenerators) {
					r.Status.HostnameGenerators = append(r.Status.HostnameGenerators, hostnameGenerators1)
				} else {
					r.Status.HostnameGenerators[hostnameGeneratorsCount].Conditions = hostnameGenerators1.Conditions
					r.Status.HostnameGenerators[hostnameGeneratorsCount].HostnameGeneratorRef = hostnameGenerators1.HostnameGeneratorRef
				}
			}
			if resp.Status.Vip == nil {
				r.Status.Vip = nil
			} else {
				r.Status.Vip = &tfTypes.Vip{}
				r.Status.Vip.IP = types.StringPointerValue(resp.Status.Vip.IP)
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
