// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshAccessLogDataSourceModel) RefreshFromSharedMeshAccessLogItem(resp *shared.MeshAccessLogItem) {
	if resp != nil {
		if resp.CreationTime != nil {
			r.CreationTime = types.StringValue(resp.CreationTime.Format(time.RFC3339Nano))
		} else {
			r.CreationTime = types.StringNull()
		}
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String)
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
		r.Spec.From = []tfTypes.From{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from1 tfTypes.From
			if fromItem.Default == nil {
				from1.Default = nil
			} else {
				from1.Default = &tfTypes.MeshAccessLogItemSpecDefault{}
				from1.Default.Backends = []tfTypes.MeshAccessLogItemBackends{}
				for backendsCount, backendsItem := range fromItem.Default.Backends {
					var backends1 tfTypes.MeshAccessLogItemBackends
					if backendsItem.File == nil {
						backends1.File = nil
					} else {
						backends1.File = &tfTypes.File{}
						if backendsItem.File.Format == nil {
							backends1.File.Format = nil
						} else {
							backends1.File.Format = &tfTypes.Format{}
							backends1.File.Format.JSON = []tfTypes.JSON{}
							for jsonCount, jsonItem := range backendsItem.File.Format.JSON {
								var json1 tfTypes.JSON
								json1.Key = types.StringPointerValue(jsonItem.Key)
								json1.Value = types.StringPointerValue(jsonItem.Value)
								if jsonCount+1 > len(backends1.File.Format.JSON) {
									backends1.File.Format.JSON = append(backends1.File.Format.JSON, json1)
								} else {
									backends1.File.Format.JSON[jsonCount].Key = json1.Key
									backends1.File.Format.JSON[jsonCount].Value = json1.Value
								}
							}
							backends1.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem.File.Format.OmitEmptyValues)
							backends1.File.Format.Plain = types.StringPointerValue(backendsItem.File.Format.Plain)
							backends1.File.Format.Type = types.StringValue(string(backendsItem.File.Format.Type))
						}
						backends1.File.Path = types.StringValue(backendsItem.File.Path)
					}
					if backendsItem.OpenTelemetry == nil {
						backends1.OpenTelemetry = nil
					} else {
						backends1.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecOpenTelemetry{}
						backends1.OpenTelemetry.Attributes = []tfTypes.JSON{}
						for attributesCount, attributesItem := range backendsItem.OpenTelemetry.Attributes {
							var attributes1 tfTypes.JSON
							attributes1.Key = types.StringPointerValue(attributesItem.Key)
							attributes1.Value = types.StringPointerValue(attributesItem.Value)
							if attributesCount+1 > len(backends1.OpenTelemetry.Attributes) {
								backends1.OpenTelemetry.Attributes = append(backends1.OpenTelemetry.Attributes, attributes1)
							} else {
								backends1.OpenTelemetry.Attributes[attributesCount].Key = attributes1.Key
								backends1.OpenTelemetry.Attributes[attributesCount].Value = attributes1.Value
							}
						}
						if backendsItem.OpenTelemetry.Body == nil {
							backends1.OpenTelemetry.Body = types.StringNull()
						} else {
							bodyResult, _ := json.Marshal(backendsItem.OpenTelemetry.Body)
							backends1.OpenTelemetry.Body = types.StringValue(string(bodyResult))
						}
						backends1.OpenTelemetry.Endpoint = types.StringValue(backendsItem.OpenTelemetry.Endpoint)
					}
					if backendsItem.TCP == nil {
						backends1.TCP = nil
					} else {
						backends1.TCP = &tfTypes.MeshAccessLogItemSpecTCP{}
						backends1.TCP.Address = types.StringValue(backendsItem.TCP.Address)
						if backendsItem.TCP.Format == nil {
							backends1.TCP.Format = nil
						} else {
							backends1.TCP.Format = &tfTypes.Format{}
							backends1.TCP.Format.JSON = []tfTypes.JSON{}
							for jsonCount1, jsonItem1 := range backendsItem.TCP.Format.JSON {
								var json3 tfTypes.JSON
								json3.Key = types.StringPointerValue(jsonItem1.Key)
								json3.Value = types.StringPointerValue(jsonItem1.Value)
								if jsonCount1+1 > len(backends1.TCP.Format.JSON) {
									backends1.TCP.Format.JSON = append(backends1.TCP.Format.JSON, json3)
								} else {
									backends1.TCP.Format.JSON[jsonCount1].Key = json3.Key
									backends1.TCP.Format.JSON[jsonCount1].Value = json3.Value
								}
							}
							backends1.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem.TCP.Format.OmitEmptyValues)
							backends1.TCP.Format.Plain = types.StringPointerValue(backendsItem.TCP.Format.Plain)
							backends1.TCP.Format.Type = types.StringValue(string(backendsItem.TCP.Format.Type))
						}
					}
					backends1.Type = types.StringValue(string(backendsItem.Type))
					if backendsCount+1 > len(from1.Default.Backends) {
						from1.Default.Backends = append(from1.Default.Backends, backends1)
					} else {
						from1.Default.Backends[backendsCount].File = backends1.File
						from1.Default.Backends[backendsCount].OpenTelemetry = backends1.OpenTelemetry
						from1.Default.Backends[backendsCount].TCP = backends1.TCP
						from1.Default.Backends[backendsCount].Type = backends1.Type
					}
				}
			}
			if fromItem.TargetRef.Kind != nil {
				from1.TargetRef.Kind = types.StringValue(string(*fromItem.TargetRef.Kind))
			} else {
				from1.TargetRef.Kind = types.StringNull()
			}
			if len(fromItem.TargetRef.Labels) > 0 {
				from1.TargetRef.Labels = make(map[string]types.String)
				for key4, value4 := range fromItem.TargetRef.Labels {
					from1.TargetRef.Labels[key4] = types.StringValue(value4)
				}
			}
			from1.TargetRef.Mesh = types.StringPointerValue(fromItem.TargetRef.Mesh)
			from1.TargetRef.Name = types.StringPointerValue(fromItem.TargetRef.Name)
			from1.TargetRef.Namespace = types.StringPointerValue(fromItem.TargetRef.Namespace)
			from1.TargetRef.ProxyTypes = []types.String{}
			for _, v := range fromItem.TargetRef.ProxyTypes {
				from1.TargetRef.ProxyTypes = append(from1.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			from1.TargetRef.SectionName = types.StringPointerValue(fromItem.TargetRef.SectionName)
			if len(fromItem.TargetRef.Tags) > 0 {
				from1.TargetRef.Tags = make(map[string]types.String)
				for key5, value5 := range fromItem.TargetRef.Tags {
					from1.TargetRef.Tags[key5] = types.StringValue(value5)
				}
			}
			if fromCount+1 > len(r.Spec.From) {
				r.Spec.From = append(r.Spec.From, from1)
			} else {
				r.Spec.From[fromCount].Default = from1.Default
				r.Spec.From[fromCount].TargetRef = from1.TargetRef
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			if resp.Spec.TargetRef.Kind != nil {
				r.Spec.TargetRef.Kind = types.StringValue(string(*resp.Spec.TargetRef.Kind))
			} else {
				r.Spec.TargetRef.Kind = types.StringNull()
			}
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String)
				for key6, value6 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key6] = types.StringValue(value6)
				}
			}
			r.Spec.TargetRef.Mesh = types.StringPointerValue(resp.Spec.TargetRef.Mesh)
			r.Spec.TargetRef.Name = types.StringPointerValue(resp.Spec.TargetRef.Name)
			r.Spec.TargetRef.Namespace = types.StringPointerValue(resp.Spec.TargetRef.Namespace)
			r.Spec.TargetRef.ProxyTypes = []types.String{}
			for _, v := range resp.Spec.TargetRef.ProxyTypes {
				r.Spec.TargetRef.ProxyTypes = append(r.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			r.Spec.TargetRef.SectionName = types.StringPointerValue(resp.Spec.TargetRef.SectionName)
			if len(resp.Spec.TargetRef.Tags) > 0 {
				r.Spec.TargetRef.Tags = make(map[string]types.String)
				for key7, value7 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key7] = types.StringValue(value7)
				}
			}
		}
		r.Spec.To = []tfTypes.From{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.From
			if toItem.Default == nil {
				to1.Default = nil
			} else {
				to1.Default = &tfTypes.MeshAccessLogItemSpecDefault{}
				to1.Default.Backends = []tfTypes.MeshAccessLogItemBackends{}
				for backendsCount1, backendsItem1 := range toItem.Default.Backends {
					var backends3 tfTypes.MeshAccessLogItemBackends
					if backendsItem1.File == nil {
						backends3.File = nil
					} else {
						backends3.File = &tfTypes.File{}
						if backendsItem1.File.Format == nil {
							backends3.File.Format = nil
						} else {
							backends3.File.Format = &tfTypes.Format{}
							backends3.File.Format.JSON = []tfTypes.JSON{}
							for jsonCount2, jsonItem2 := range backendsItem1.File.Format.JSON {
								var json5 tfTypes.JSON
								json5.Key = types.StringPointerValue(jsonItem2.Key)
								json5.Value = types.StringPointerValue(jsonItem2.Value)
								if jsonCount2+1 > len(backends3.File.Format.JSON) {
									backends3.File.Format.JSON = append(backends3.File.Format.JSON, json5)
								} else {
									backends3.File.Format.JSON[jsonCount2].Key = json5.Key
									backends3.File.Format.JSON[jsonCount2].Value = json5.Value
								}
							}
							backends3.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem1.File.Format.OmitEmptyValues)
							backends3.File.Format.Plain = types.StringPointerValue(backendsItem1.File.Format.Plain)
							backends3.File.Format.Type = types.StringValue(string(backendsItem1.File.Format.Type))
						}
						backends3.File.Path = types.StringValue(backendsItem1.File.Path)
					}
					if backendsItem1.OpenTelemetry == nil {
						backends3.OpenTelemetry = nil
					} else {
						backends3.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecOpenTelemetry{}
						backends3.OpenTelemetry.Attributes = []tfTypes.JSON{}
						for attributesCount1, attributesItem1 := range backendsItem1.OpenTelemetry.Attributes {
							var attributes3 tfTypes.JSON
							attributes3.Key = types.StringPointerValue(attributesItem1.Key)
							attributes3.Value = types.StringPointerValue(attributesItem1.Value)
							if attributesCount1+1 > len(backends3.OpenTelemetry.Attributes) {
								backends3.OpenTelemetry.Attributes = append(backends3.OpenTelemetry.Attributes, attributes3)
							} else {
								backends3.OpenTelemetry.Attributes[attributesCount1].Key = attributes3.Key
								backends3.OpenTelemetry.Attributes[attributesCount1].Value = attributes3.Value
							}
						}
						if backendsItem1.OpenTelemetry.Body == nil {
							backends3.OpenTelemetry.Body = types.StringNull()
						} else {
							bodyResult1, _ := json.Marshal(backendsItem1.OpenTelemetry.Body)
							backends3.OpenTelemetry.Body = types.StringValue(string(bodyResult1))
						}
						backends3.OpenTelemetry.Endpoint = types.StringValue(backendsItem1.OpenTelemetry.Endpoint)
					}
					if backendsItem1.TCP == nil {
						backends3.TCP = nil
					} else {
						backends3.TCP = &tfTypes.MeshAccessLogItemSpecTCP{}
						backends3.TCP.Address = types.StringValue(backendsItem1.TCP.Address)
						if backendsItem1.TCP.Format == nil {
							backends3.TCP.Format = nil
						} else {
							backends3.TCP.Format = &tfTypes.Format{}
							backends3.TCP.Format.JSON = []tfTypes.JSON{}
							for jsonCount3, jsonItem3 := range backendsItem1.TCP.Format.JSON {
								var json7 tfTypes.JSON
								json7.Key = types.StringPointerValue(jsonItem3.Key)
								json7.Value = types.StringPointerValue(jsonItem3.Value)
								if jsonCount3+1 > len(backends3.TCP.Format.JSON) {
									backends3.TCP.Format.JSON = append(backends3.TCP.Format.JSON, json7)
								} else {
									backends3.TCP.Format.JSON[jsonCount3].Key = json7.Key
									backends3.TCP.Format.JSON[jsonCount3].Value = json7.Value
								}
							}
							backends3.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem1.TCP.Format.OmitEmptyValues)
							backends3.TCP.Format.Plain = types.StringPointerValue(backendsItem1.TCP.Format.Plain)
							backends3.TCP.Format.Type = types.StringValue(string(backendsItem1.TCP.Format.Type))
						}
					}
					backends3.Type = types.StringValue(string(backendsItem1.Type))
					if backendsCount1+1 > len(to1.Default.Backends) {
						to1.Default.Backends = append(to1.Default.Backends, backends3)
					} else {
						to1.Default.Backends[backendsCount1].File = backends3.File
						to1.Default.Backends[backendsCount1].OpenTelemetry = backends3.OpenTelemetry
						to1.Default.Backends[backendsCount1].TCP = backends3.TCP
						to1.Default.Backends[backendsCount1].Type = backends3.Type
					}
				}
			}
			if toItem.TargetRef.Kind != nil {
				to1.TargetRef.Kind = types.StringValue(string(*toItem.TargetRef.Kind))
			} else {
				to1.TargetRef.Kind = types.StringNull()
			}
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String)
				for key11, value11 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key11] = types.StringValue(value11)
				}
			}
			to1.TargetRef.Mesh = types.StringPointerValue(toItem.TargetRef.Mesh)
			to1.TargetRef.Name = types.StringPointerValue(toItem.TargetRef.Name)
			to1.TargetRef.Namespace = types.StringPointerValue(toItem.TargetRef.Namespace)
			to1.TargetRef.ProxyTypes = []types.String{}
			for _, v := range toItem.TargetRef.ProxyTypes {
				to1.TargetRef.ProxyTypes = append(to1.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			to1.TargetRef.SectionName = types.StringPointerValue(toItem.TargetRef.SectionName)
			if len(toItem.TargetRef.Tags) > 0 {
				to1.TargetRef.Tags = make(map[string]types.String)
				for key12, value12 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key12] = types.StringValue(value12)
				}
			}
			if toCount+1 > len(r.Spec.To) {
				r.Spec.To = append(r.Spec.To, to1)
			} else {
				r.Spec.To[toCount].Default = to1.Default
				r.Spec.To[toCount].TargetRef = to1.TargetRef
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}