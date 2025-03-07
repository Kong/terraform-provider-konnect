// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshTLSListDataSourceModel) RefreshFromSharedMeshTLSList(resp *shared.MeshTLSList) {
	if resp != nil {
		r.Items = []tfTypes.MeshTLSItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshTLSItem
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
			items1.Spec.From = []tfTypes.MeshTLSItemFrom{}
			for fromCount, fromItem := range itemsItem.Spec.From {
				var from1 tfTypes.MeshTLSItemFrom
				if fromItem.Default == nil {
					from1.Default = nil
				} else {
					from1.Default = &tfTypes.MeshTLSItemDefault{}
					if fromItem.Default.Mode != nil {
						from1.Default.Mode = types.StringValue(string(*fromItem.Default.Mode))
					} else {
						from1.Default.Mode = types.StringNull()
					}
					from1.Default.TLSCiphers = make([]types.String, 0, len(fromItem.Default.TLSCiphers))
					for _, v := range fromItem.Default.TLSCiphers {
						from1.Default.TLSCiphers = append(from1.Default.TLSCiphers, types.StringValue(string(v)))
					}
					if fromItem.Default.TLSVersion == nil {
						from1.Default.TLSVersion = nil
					} else {
						from1.Default.TLSVersion = &tfTypes.Version{}
						if fromItem.Default.TLSVersion.Max != nil {
							from1.Default.TLSVersion.Max = types.StringValue(string(*fromItem.Default.TLSVersion.Max))
						} else {
							from1.Default.TLSVersion.Max = types.StringNull()
						}
						if fromItem.Default.TLSVersion.Min != nil {
							from1.Default.TLSVersion.Min = types.StringValue(string(*fromItem.Default.TLSVersion.Min))
						} else {
							from1.Default.TLSVersion.Min = types.StringNull()
						}
					}
				}
				from1.TargetRef.Kind = types.StringValue(string(fromItem.TargetRef.Kind))
				if len(fromItem.TargetRef.Labels) > 0 {
					from1.TargetRef.Labels = make(map[string]types.String)
					for key1, value1 := range fromItem.TargetRef.Labels {
						from1.TargetRef.Labels[key1] = types.StringValue(value1)
					}
				}
				from1.TargetRef.Mesh = types.StringPointerValue(fromItem.TargetRef.Mesh)
				from1.TargetRef.Name = types.StringPointerValue(fromItem.TargetRef.Name)
				from1.TargetRef.Namespace = types.StringPointerValue(fromItem.TargetRef.Namespace)
				from1.TargetRef.ProxyTypes = make([]types.String, 0, len(fromItem.TargetRef.ProxyTypes))
				for _, v := range fromItem.TargetRef.ProxyTypes {
					from1.TargetRef.ProxyTypes = append(from1.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				from1.TargetRef.SectionName = types.StringPointerValue(fromItem.TargetRef.SectionName)
				if len(fromItem.TargetRef.Tags) > 0 {
					from1.TargetRef.Tags = make(map[string]types.String)
					for key2, value2 := range fromItem.TargetRef.Tags {
						from1.TargetRef.Tags[key2] = types.StringValue(value2)
					}
				}
				if fromCount+1 > len(items1.Spec.From) {
					items1.Spec.From = append(items1.Spec.From, from1)
				} else {
					items1.Spec.From[fromCount].Default = from1.Default
					items1.Spec.From[fromCount].TargetRef = from1.TargetRef
				}
			}
			items1.Spec.Rules = []tfTypes.MeshTLSItemRules{}
			for rulesCount, rulesItem := range itemsItem.Spec.Rules {
				var rules1 tfTypes.MeshTLSItemRules
				if rulesItem.Default == nil {
					rules1.Default = nil
				} else {
					rules1.Default = &tfTypes.MeshTLSItemDefault{}
					if rulesItem.Default.Mode != nil {
						rules1.Default.Mode = types.StringValue(string(*rulesItem.Default.Mode))
					} else {
						rules1.Default.Mode = types.StringNull()
					}
					rules1.Default.TLSCiphers = make([]types.String, 0, len(rulesItem.Default.TLSCiphers))
					for _, v := range rulesItem.Default.TLSCiphers {
						rules1.Default.TLSCiphers = append(rules1.Default.TLSCiphers, types.StringValue(string(v)))
					}
					if rulesItem.Default.TLSVersion == nil {
						rules1.Default.TLSVersion = nil
					} else {
						rules1.Default.TLSVersion = &tfTypes.Version{}
						if rulesItem.Default.TLSVersion.Max != nil {
							rules1.Default.TLSVersion.Max = types.StringValue(string(*rulesItem.Default.TLSVersion.Max))
						} else {
							rules1.Default.TLSVersion.Max = types.StringNull()
						}
						if rulesItem.Default.TLSVersion.Min != nil {
							rules1.Default.TLSVersion.Min = types.StringValue(string(*rulesItem.Default.TLSVersion.Min))
						} else {
							rules1.Default.TLSVersion.Min = types.StringNull()
						}
					}
				}
				if rulesCount+1 > len(items1.Spec.Rules) {
					items1.Spec.Rules = append(items1.Spec.Rules, rules1)
				} else {
					items1.Spec.Rules[rulesCount].Default = rules1.Default
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items1.Spec.TargetRef = nil
			} else {
				items1.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				items1.Spec.TargetRef.Kind = types.StringValue(string(itemsItem.Spec.TargetRef.Kind))
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items1.Spec.TargetRef.Labels = make(map[string]types.String)
					for key3, value3 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key3] = types.StringValue(value3)
					}
				}
				items1.Spec.TargetRef.Mesh = types.StringPointerValue(itemsItem.Spec.TargetRef.Mesh)
				items1.Spec.TargetRef.Name = types.StringPointerValue(itemsItem.Spec.TargetRef.Name)
				items1.Spec.TargetRef.Namespace = types.StringPointerValue(itemsItem.Spec.TargetRef.Namespace)
				items1.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(itemsItem.Spec.TargetRef.ProxyTypes))
				for _, v := range itemsItem.Spec.TargetRef.ProxyTypes {
					items1.Spec.TargetRef.ProxyTypes = append(items1.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				items1.Spec.TargetRef.SectionName = types.StringPointerValue(itemsItem.Spec.TargetRef.SectionName)
				if len(itemsItem.Spec.TargetRef.Tags) > 0 {
					items1.Spec.TargetRef.Tags = make(map[string]types.String)
					for key4, value4 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key4] = types.StringValue(value4)
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
