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

func (r *MeshProxyPatchListDataSourceModel) RefreshFromSharedMeshProxyPatchList(resp *shared.MeshProxyPatchList) {
	if resp != nil {
		r.Items = []tfTypes.MeshProxyPatchItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshProxyPatchItem
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
			items1.Spec.Default.AppendModifications = []tfTypes.AppendModifications{}
			for appendModificationsCount, appendModificationsItem := range itemsItem.Spec.Default.AppendModifications {
				var appendModifications1 tfTypes.AppendModifications
				if appendModificationsItem.Cluster == nil {
					appendModifications1.Cluster = nil
				} else {
					appendModifications1.Cluster = &tfTypes.Cluster{}
					appendModifications1.Cluster.JSONPatches = []tfTypes.JSONPatches{}
					for jsonPatchesCount, jsonPatchesItem := range appendModificationsItem.Cluster.JSONPatches {
						var jsonPatches1 tfTypes.JSONPatches
						jsonPatches1.From = types.StringPointerValue(jsonPatchesItem.From)
						jsonPatches1.Op = types.StringValue(string(jsonPatchesItem.Op))
						jsonPatches1.Path = types.StringValue(jsonPatchesItem.Path)
						if jsonPatchesItem.Value == nil {
							jsonPatches1.Value = types.StringNull()
						} else {
							valueResult, _ := json.Marshal(jsonPatchesItem.Value)
							jsonPatches1.Value = types.StringValue(string(valueResult))
						}
						if jsonPatchesCount+1 > len(appendModifications1.Cluster.JSONPatches) {
							appendModifications1.Cluster.JSONPatches = append(appendModifications1.Cluster.JSONPatches, jsonPatches1)
						} else {
							appendModifications1.Cluster.JSONPatches[jsonPatchesCount].From = jsonPatches1.From
							appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Op = jsonPatches1.Op
							appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Path = jsonPatches1.Path
							appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Value = jsonPatches1.Value
						}
					}
					if appendModificationsItem.Cluster.Match == nil {
						appendModifications1.Cluster.Match = nil
					} else {
						appendModifications1.Cluster.Match = &tfTypes.MeshProxyPatchItemMatch{}
						appendModifications1.Cluster.Match.Name = types.StringPointerValue(appendModificationsItem.Cluster.Match.Name)
						appendModifications1.Cluster.Match.Origin = types.StringPointerValue(appendModificationsItem.Cluster.Match.Origin)
					}
					appendModifications1.Cluster.Operation = types.StringValue(string(appendModificationsItem.Cluster.Operation))
					appendModifications1.Cluster.Value = types.StringPointerValue(appendModificationsItem.Cluster.Value)
				}
				if appendModificationsItem.HTTPFilter == nil {
					appendModifications1.HTTPFilter = nil
				} else {
					appendModifications1.HTTPFilter = &tfTypes.HTTPFilter{}
					appendModifications1.HTTPFilter.JSONPatches = []tfTypes.JSONPatches{}
					for jsonPatchesCount1, jsonPatchesItem1 := range appendModificationsItem.HTTPFilter.JSONPatches {
						var jsonPatches3 tfTypes.JSONPatches
						jsonPatches3.From = types.StringPointerValue(jsonPatchesItem1.From)
						jsonPatches3.Op = types.StringValue(string(jsonPatchesItem1.Op))
						jsonPatches3.Path = types.StringValue(jsonPatchesItem1.Path)
						if jsonPatchesItem1.Value == nil {
							jsonPatches3.Value = types.StringNull()
						} else {
							valueResult1, _ := json.Marshal(jsonPatchesItem1.Value)
							jsonPatches3.Value = types.StringValue(string(valueResult1))
						}
						if jsonPatchesCount1+1 > len(appendModifications1.HTTPFilter.JSONPatches) {
							appendModifications1.HTTPFilter.JSONPatches = append(appendModifications1.HTTPFilter.JSONPatches, jsonPatches3)
						} else {
							appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].From = jsonPatches3.From
							appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Op = jsonPatches3.Op
							appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Path = jsonPatches3.Path
							appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Value = jsonPatches3.Value
						}
					}
					if appendModificationsItem.HTTPFilter.Match == nil {
						appendModifications1.HTTPFilter.Match = nil
					} else {
						appendModifications1.HTTPFilter.Match = &tfTypes.MeshProxyPatchItemSpecMatch{}
						appendModifications1.HTTPFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.ListenerName)
						if len(appendModificationsItem.HTTPFilter.Match.ListenerTags) > 0 {
							appendModifications1.HTTPFilter.Match.ListenerTags = make(map[string]types.String)
							for key1, value4 := range appendModificationsItem.HTTPFilter.Match.ListenerTags {
								appendModifications1.HTTPFilter.Match.ListenerTags[key1] = types.StringValue(value4)
							}
						}
						appendModifications1.HTTPFilter.Match.Name = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Name)
						appendModifications1.HTTPFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Origin)
					}
					appendModifications1.HTTPFilter.Operation = types.StringValue(string(appendModificationsItem.HTTPFilter.Operation))
					appendModifications1.HTTPFilter.Value = types.StringPointerValue(appendModificationsItem.HTTPFilter.Value)
				}
				if appendModificationsItem.Listener == nil {
					appendModifications1.Listener = nil
				} else {
					appendModifications1.Listener = &tfTypes.Listener{}
					appendModifications1.Listener.JSONPatches = []tfTypes.JSONPatches{}
					for jsonPatchesCount2, jsonPatchesItem2 := range appendModificationsItem.Listener.JSONPatches {
						var jsonPatches5 tfTypes.JSONPatches
						jsonPatches5.From = types.StringPointerValue(jsonPatchesItem2.From)
						jsonPatches5.Op = types.StringValue(string(jsonPatchesItem2.Op))
						jsonPatches5.Path = types.StringValue(jsonPatchesItem2.Path)
						if jsonPatchesItem2.Value == nil {
							jsonPatches5.Value = types.StringNull()
						} else {
							valueResult2, _ := json.Marshal(jsonPatchesItem2.Value)
							jsonPatches5.Value = types.StringValue(string(valueResult2))
						}
						if jsonPatchesCount2+1 > len(appendModifications1.Listener.JSONPatches) {
							appendModifications1.Listener.JSONPatches = append(appendModifications1.Listener.JSONPatches, jsonPatches5)
						} else {
							appendModifications1.Listener.JSONPatches[jsonPatchesCount2].From = jsonPatches5.From
							appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Op = jsonPatches5.Op
							appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Path = jsonPatches5.Path
							appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Value = jsonPatches5.Value
						}
					}
					if appendModificationsItem.Listener.Match == nil {
						appendModifications1.Listener.Match = nil
					} else {
						appendModifications1.Listener.Match = &tfTypes.MeshProxyPatchItemSpecDefaultMatch{}
						appendModifications1.Listener.Match.Name = types.StringPointerValue(appendModificationsItem.Listener.Match.Name)
						appendModifications1.Listener.Match.Origin = types.StringPointerValue(appendModificationsItem.Listener.Match.Origin)
						if len(appendModificationsItem.Listener.Match.Tags) > 0 {
							appendModifications1.Listener.Match.Tags = make(map[string]types.String)
							for key2, value7 := range appendModificationsItem.Listener.Match.Tags {
								appendModifications1.Listener.Match.Tags[key2] = types.StringValue(value7)
							}
						}
					}
					appendModifications1.Listener.Operation = types.StringValue(string(appendModificationsItem.Listener.Operation))
					appendModifications1.Listener.Value = types.StringPointerValue(appendModificationsItem.Listener.Value)
				}
				if appendModificationsItem.NetworkFilter == nil {
					appendModifications1.NetworkFilter = nil
				} else {
					appendModifications1.NetworkFilter = &tfTypes.HTTPFilter{}
					appendModifications1.NetworkFilter.JSONPatches = []tfTypes.JSONPatches{}
					for jsonPatchesCount3, jsonPatchesItem3 := range appendModificationsItem.NetworkFilter.JSONPatches {
						var jsonPatches7 tfTypes.JSONPatches
						jsonPatches7.From = types.StringPointerValue(jsonPatchesItem3.From)
						jsonPatches7.Op = types.StringValue(string(jsonPatchesItem3.Op))
						jsonPatches7.Path = types.StringValue(jsonPatchesItem3.Path)
						if jsonPatchesItem3.Value == nil {
							jsonPatches7.Value = types.StringNull()
						} else {
							valueResult3, _ := json.Marshal(jsonPatchesItem3.Value)
							jsonPatches7.Value = types.StringValue(string(valueResult3))
						}
						if jsonPatchesCount3+1 > len(appendModifications1.NetworkFilter.JSONPatches) {
							appendModifications1.NetworkFilter.JSONPatches = append(appendModifications1.NetworkFilter.JSONPatches, jsonPatches7)
						} else {
							appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].From = jsonPatches7.From
							appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Op = jsonPatches7.Op
							appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Path = jsonPatches7.Path
							appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Value = jsonPatches7.Value
						}
					}
					if appendModificationsItem.NetworkFilter.Match == nil {
						appendModifications1.NetworkFilter.Match = nil
					} else {
						appendModifications1.NetworkFilter.Match = &tfTypes.MeshProxyPatchItemSpecMatch{}
						appendModifications1.NetworkFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.ListenerName)
						if len(appendModificationsItem.NetworkFilter.Match.ListenerTags) > 0 {
							appendModifications1.NetworkFilter.Match.ListenerTags = make(map[string]types.String)
							for key3, value10 := range appendModificationsItem.NetworkFilter.Match.ListenerTags {
								appendModifications1.NetworkFilter.Match.ListenerTags[key3] = types.StringValue(value10)
							}
						}
						appendModifications1.NetworkFilter.Match.Name = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Name)
						appendModifications1.NetworkFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Origin)
					}
					appendModifications1.NetworkFilter.Operation = types.StringValue(string(appendModificationsItem.NetworkFilter.Operation))
					appendModifications1.NetworkFilter.Value = types.StringPointerValue(appendModificationsItem.NetworkFilter.Value)
				}
				if appendModificationsItem.VirtualHost == nil {
					appendModifications1.VirtualHost = nil
				} else {
					appendModifications1.VirtualHost = &tfTypes.VirtualHost{}
					appendModifications1.VirtualHost.JSONPatches = []tfTypes.JSONPatches{}
					for jsonPatchesCount4, jsonPatchesItem4 := range appendModificationsItem.VirtualHost.JSONPatches {
						var jsonPatches9 tfTypes.JSONPatches
						jsonPatches9.From = types.StringPointerValue(jsonPatchesItem4.From)
						jsonPatches9.Op = types.StringValue(string(jsonPatchesItem4.Op))
						jsonPatches9.Path = types.StringValue(jsonPatchesItem4.Path)
						if jsonPatchesItem4.Value == nil {
							jsonPatches9.Value = types.StringNull()
						} else {
							valueResult4, _ := json.Marshal(jsonPatchesItem4.Value)
							jsonPatches9.Value = types.StringValue(string(valueResult4))
						}
						if jsonPatchesCount4+1 > len(appendModifications1.VirtualHost.JSONPatches) {
							appendModifications1.VirtualHost.JSONPatches = append(appendModifications1.VirtualHost.JSONPatches, jsonPatches9)
						} else {
							appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].From = jsonPatches9.From
							appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Op = jsonPatches9.Op
							appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Path = jsonPatches9.Path
							appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Value = jsonPatches9.Value
						}
					}
					appendModifications1.VirtualHost.Match.Name = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Name)
					appendModifications1.VirtualHost.Match.Origin = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Origin)
					appendModifications1.VirtualHost.Match.RouteConfigurationName = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.RouteConfigurationName)
					appendModifications1.VirtualHost.Operation = types.StringValue(string(appendModificationsItem.VirtualHost.Operation))
					appendModifications1.VirtualHost.Value = types.StringPointerValue(appendModificationsItem.VirtualHost.Value)
				}
				if appendModificationsCount+1 > len(items1.Spec.Default.AppendModifications) {
					items1.Spec.Default.AppendModifications = append(items1.Spec.Default.AppendModifications, appendModifications1)
				} else {
					items1.Spec.Default.AppendModifications[appendModificationsCount].Cluster = appendModifications1.Cluster
					items1.Spec.Default.AppendModifications[appendModificationsCount].HTTPFilter = appendModifications1.HTTPFilter
					items1.Spec.Default.AppendModifications[appendModificationsCount].Listener = appendModifications1.Listener
					items1.Spec.Default.AppendModifications[appendModificationsCount].NetworkFilter = appendModifications1.NetworkFilter
					items1.Spec.Default.AppendModifications[appendModificationsCount].VirtualHost = appendModifications1.VirtualHost
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items1.Spec.TargetRef = nil
			} else {
				items1.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				if itemsItem.Spec.TargetRef.Kind != nil {
					items1.Spec.TargetRef.Kind = types.StringValue(string(*itemsItem.Spec.TargetRef.Kind))
				} else {
					items1.Spec.TargetRef.Kind = types.StringNull()
				}
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items1.Spec.TargetRef.Labels = make(map[string]types.String)
					for key4, value14 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key4] = types.StringValue(value14)
					}
				}
				items1.Spec.TargetRef.Mesh = types.StringPointerValue(itemsItem.Spec.TargetRef.Mesh)
				items1.Spec.TargetRef.Name = types.StringPointerValue(itemsItem.Spec.TargetRef.Name)
				items1.Spec.TargetRef.Namespace = types.StringPointerValue(itemsItem.Spec.TargetRef.Namespace)
				items1.Spec.TargetRef.ProxyTypes = []types.String{}
				for _, v := range itemsItem.Spec.TargetRef.ProxyTypes {
					items1.Spec.TargetRef.ProxyTypes = append(items1.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				items1.Spec.TargetRef.SectionName = types.StringPointerValue(itemsItem.Spec.TargetRef.SectionName)
				if len(itemsItem.Spec.TargetRef.Tags) > 0 {
					items1.Spec.TargetRef.Tags = make(map[string]types.String)
					for key5, value15 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key5] = types.StringValue(value15)
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