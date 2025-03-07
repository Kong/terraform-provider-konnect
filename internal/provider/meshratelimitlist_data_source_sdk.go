// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshRateLimitListDataSourceModel) RefreshFromSharedMeshRateLimitList(resp *shared.MeshRateLimitList) {
	if resp != nil {
		r.Items = []tfTypes.MeshRateLimitItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshRateLimitItem
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
			items1.Spec.From = []tfTypes.MeshRateLimitItemFrom{}
			for fromCount, fromItem := range itemsItem.Spec.From {
				var from1 tfTypes.MeshRateLimitItemFrom
				if fromItem.Default == nil {
					from1.Default = nil
				} else {
					from1.Default = &tfTypes.MeshRateLimitItemDefault{}
					if fromItem.Default.Local == nil {
						from1.Default.Local = nil
					} else {
						from1.Default.Local = &tfTypes.Local{}
						if fromItem.Default.Local.HTTP == nil {
							from1.Default.Local.HTTP = nil
						} else {
							from1.Default.Local.HTTP = &tfTypes.MeshRateLimitItemHTTP{}
							from1.Default.Local.HTTP.Disabled = types.BoolPointerValue(fromItem.Default.Local.HTTP.Disabled)
							if fromItem.Default.Local.HTTP.OnRateLimit == nil {
								from1.Default.Local.HTTP.OnRateLimit = nil
							} else {
								from1.Default.Local.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
								if fromItem.Default.Local.HTTP.OnRateLimit.Headers == nil {
									from1.Default.Local.HTTP.OnRateLimit.Headers = nil
								} else {
									from1.Default.Local.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemSpecHeaders{}
									from1.Default.Local.HTTP.OnRateLimit.Headers.Add = []tfTypes.MeshGlobalRateLimitItemSpecAdd{}
									for addCount, addItem := range fromItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
										var add1 tfTypes.MeshGlobalRateLimitItemSpecAdd
										add1.Name = types.StringValue(addItem.Name)
										add1.Value = types.StringValue(addItem.Value)
										if addCount+1 > len(from1.Default.Local.HTTP.OnRateLimit.Headers.Add) {
											from1.Default.Local.HTTP.OnRateLimit.Headers.Add = append(from1.Default.Local.HTTP.OnRateLimit.Headers.Add, add1)
										} else {
											from1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount].Name = add1.Name
											from1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount].Value = add1.Value
										}
									}
									from1.Default.Local.HTTP.OnRateLimit.Headers.Set = []tfTypes.MeshGlobalRateLimitItemSpecAdd{}
									for setCount, setItem := range fromItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
										var set1 tfTypes.MeshGlobalRateLimitItemSpecAdd
										set1.Name = types.StringValue(setItem.Name)
										set1.Value = types.StringValue(setItem.Value)
										if setCount+1 > len(from1.Default.Local.HTTP.OnRateLimit.Headers.Set) {
											from1.Default.Local.HTTP.OnRateLimit.Headers.Set = append(from1.Default.Local.HTTP.OnRateLimit.Headers.Set, set1)
										} else {
											from1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount].Name = set1.Name
											from1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount].Value = set1.Value
										}
									}
								}
								if fromItem.Default.Local.HTTP.OnRateLimit.Status != nil {
									from1.Default.Local.HTTP.OnRateLimit.Status = types.Int64Value(int64(*fromItem.Default.Local.HTTP.OnRateLimit.Status))
								} else {
									from1.Default.Local.HTTP.OnRateLimit.Status = types.Int64Null()
								}
							}
							if fromItem.Default.Local.HTTP.RequestRate == nil {
								from1.Default.Local.HTTP.RequestRate = nil
							} else {
								from1.Default.Local.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
								from1.Default.Local.HTTP.RequestRate.Interval = types.StringValue(fromItem.Default.Local.HTTP.RequestRate.Interval)
								from1.Default.Local.HTTP.RequestRate.Num = types.Int64Value(int64(fromItem.Default.Local.HTTP.RequestRate.Num))
							}
						}
						if fromItem.Default.Local.TCP == nil {
							from1.Default.Local.TCP = nil
						} else {
							from1.Default.Local.TCP = &tfTypes.MeshRateLimitItemTCP{}
							if fromItem.Default.Local.TCP.ConnectionRate == nil {
								from1.Default.Local.TCP.ConnectionRate = nil
							} else {
								from1.Default.Local.TCP.ConnectionRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
								from1.Default.Local.TCP.ConnectionRate.Interval = types.StringValue(fromItem.Default.Local.TCP.ConnectionRate.Interval)
								from1.Default.Local.TCP.ConnectionRate.Num = types.Int64Value(int64(fromItem.Default.Local.TCP.ConnectionRate.Num))
							}
							from1.Default.Local.TCP.Disabled = types.BoolPointerValue(fromItem.Default.Local.TCP.Disabled)
						}
					}
				}
				from1.TargetRef.Kind = types.StringValue(string(fromItem.TargetRef.Kind))
				if len(fromItem.TargetRef.Labels) > 0 {
					from1.TargetRef.Labels = make(map[string]types.String)
					for key1, value3 := range fromItem.TargetRef.Labels {
						from1.TargetRef.Labels[key1] = types.StringValue(value3)
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
					for key2, value4 := range fromItem.TargetRef.Tags {
						from1.TargetRef.Tags[key2] = types.StringValue(value4)
					}
				}
				if fromCount+1 > len(items1.Spec.From) {
					items1.Spec.From = append(items1.Spec.From, from1)
				} else {
					items1.Spec.From[fromCount].Default = from1.Default
					items1.Spec.From[fromCount].TargetRef = from1.TargetRef
				}
			}
			items1.Spec.Rules = []tfTypes.MeshRateLimitItemRules{}
			for rulesCount, rulesItem := range itemsItem.Spec.Rules {
				var rules1 tfTypes.MeshRateLimitItemRules
				if rulesItem.Default == nil {
					rules1.Default = nil
				} else {
					rules1.Default = &tfTypes.MeshRateLimitItemDefault{}
					if rulesItem.Default.Local == nil {
						rules1.Default.Local = nil
					} else {
						rules1.Default.Local = &tfTypes.Local{}
						if rulesItem.Default.Local.HTTP == nil {
							rules1.Default.Local.HTTP = nil
						} else {
							rules1.Default.Local.HTTP = &tfTypes.MeshRateLimitItemHTTP{}
							rules1.Default.Local.HTTP.Disabled = types.BoolPointerValue(rulesItem.Default.Local.HTTP.Disabled)
							if rulesItem.Default.Local.HTTP.OnRateLimit == nil {
								rules1.Default.Local.HTTP.OnRateLimit = nil
							} else {
								rules1.Default.Local.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
								if rulesItem.Default.Local.HTTP.OnRateLimit.Headers == nil {
									rules1.Default.Local.HTTP.OnRateLimit.Headers = nil
								} else {
									rules1.Default.Local.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemSpecHeaders{}
									rules1.Default.Local.HTTP.OnRateLimit.Headers.Add = []tfTypes.MeshGlobalRateLimitItemSpecAdd{}
									for addCount1, addItem1 := range rulesItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
										var add3 tfTypes.MeshGlobalRateLimitItemSpecAdd
										add3.Name = types.StringValue(addItem1.Name)
										add3.Value = types.StringValue(addItem1.Value)
										if addCount1+1 > len(rules1.Default.Local.HTTP.OnRateLimit.Headers.Add) {
											rules1.Default.Local.HTTP.OnRateLimit.Headers.Add = append(rules1.Default.Local.HTTP.OnRateLimit.Headers.Add, add3)
										} else {
											rules1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount1].Name = add3.Name
											rules1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount1].Value = add3.Value
										}
									}
									rules1.Default.Local.HTTP.OnRateLimit.Headers.Set = []tfTypes.MeshGlobalRateLimitItemSpecAdd{}
									for setCount1, setItem1 := range rulesItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
										var set3 tfTypes.MeshGlobalRateLimitItemSpecAdd
										set3.Name = types.StringValue(setItem1.Name)
										set3.Value = types.StringValue(setItem1.Value)
										if setCount1+1 > len(rules1.Default.Local.HTTP.OnRateLimit.Headers.Set) {
											rules1.Default.Local.HTTP.OnRateLimit.Headers.Set = append(rules1.Default.Local.HTTP.OnRateLimit.Headers.Set, set3)
										} else {
											rules1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount1].Name = set3.Name
											rules1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount1].Value = set3.Value
										}
									}
								}
								if rulesItem.Default.Local.HTTP.OnRateLimit.Status != nil {
									rules1.Default.Local.HTTP.OnRateLimit.Status = types.Int64Value(int64(*rulesItem.Default.Local.HTTP.OnRateLimit.Status))
								} else {
									rules1.Default.Local.HTTP.OnRateLimit.Status = types.Int64Null()
								}
							}
							if rulesItem.Default.Local.HTTP.RequestRate == nil {
								rules1.Default.Local.HTTP.RequestRate = nil
							} else {
								rules1.Default.Local.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
								rules1.Default.Local.HTTP.RequestRate.Interval = types.StringValue(rulesItem.Default.Local.HTTP.RequestRate.Interval)
								rules1.Default.Local.HTTP.RequestRate.Num = types.Int64Value(int64(rulesItem.Default.Local.HTTP.RequestRate.Num))
							}
						}
						if rulesItem.Default.Local.TCP == nil {
							rules1.Default.Local.TCP = nil
						} else {
							rules1.Default.Local.TCP = &tfTypes.MeshRateLimitItemTCP{}
							if rulesItem.Default.Local.TCP.ConnectionRate == nil {
								rules1.Default.Local.TCP.ConnectionRate = nil
							} else {
								rules1.Default.Local.TCP.ConnectionRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
								rules1.Default.Local.TCP.ConnectionRate.Interval = types.StringValue(rulesItem.Default.Local.TCP.ConnectionRate.Interval)
								rules1.Default.Local.TCP.ConnectionRate.Num = types.Int64Value(int64(rulesItem.Default.Local.TCP.ConnectionRate.Num))
							}
							rules1.Default.Local.TCP.Disabled = types.BoolPointerValue(rulesItem.Default.Local.TCP.Disabled)
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
					for key3, value7 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key3] = types.StringValue(value7)
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
					for key4, value8 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key4] = types.StringValue(value8)
					}
				}
			}
			items1.Spec.To = []tfTypes.MeshRateLimitItemFrom{}
			for toCount, toItem := range itemsItem.Spec.To {
				var to1 tfTypes.MeshRateLimitItemFrom
				if toItem.Default == nil {
					to1.Default = nil
				} else {
					to1.Default = &tfTypes.MeshRateLimitItemDefault{}
					if toItem.Default.Local == nil {
						to1.Default.Local = nil
					} else {
						to1.Default.Local = &tfTypes.Local{}
						if toItem.Default.Local.HTTP == nil {
							to1.Default.Local.HTTP = nil
						} else {
							to1.Default.Local.HTTP = &tfTypes.MeshRateLimitItemHTTP{}
							to1.Default.Local.HTTP.Disabled = types.BoolPointerValue(toItem.Default.Local.HTTP.Disabled)
							if toItem.Default.Local.HTTP.OnRateLimit == nil {
								to1.Default.Local.HTTP.OnRateLimit = nil
							} else {
								to1.Default.Local.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
								if toItem.Default.Local.HTTP.OnRateLimit.Headers == nil {
									to1.Default.Local.HTTP.OnRateLimit.Headers = nil
								} else {
									to1.Default.Local.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemSpecHeaders{}
									to1.Default.Local.HTTP.OnRateLimit.Headers.Add = []tfTypes.MeshGlobalRateLimitItemSpecAdd{}
									for addCount2, addItem2 := range toItem.Default.Local.HTTP.OnRateLimit.Headers.Add {
										var add5 tfTypes.MeshGlobalRateLimitItemSpecAdd
										add5.Name = types.StringValue(addItem2.Name)
										add5.Value = types.StringValue(addItem2.Value)
										if addCount2+1 > len(to1.Default.Local.HTTP.OnRateLimit.Headers.Add) {
											to1.Default.Local.HTTP.OnRateLimit.Headers.Add = append(to1.Default.Local.HTTP.OnRateLimit.Headers.Add, add5)
										} else {
											to1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount2].Name = add5.Name
											to1.Default.Local.HTTP.OnRateLimit.Headers.Add[addCount2].Value = add5.Value
										}
									}
									to1.Default.Local.HTTP.OnRateLimit.Headers.Set = []tfTypes.MeshGlobalRateLimitItemSpecAdd{}
									for setCount2, setItem2 := range toItem.Default.Local.HTTP.OnRateLimit.Headers.Set {
										var set5 tfTypes.MeshGlobalRateLimitItemSpecAdd
										set5.Name = types.StringValue(setItem2.Name)
										set5.Value = types.StringValue(setItem2.Value)
										if setCount2+1 > len(to1.Default.Local.HTTP.OnRateLimit.Headers.Set) {
											to1.Default.Local.HTTP.OnRateLimit.Headers.Set = append(to1.Default.Local.HTTP.OnRateLimit.Headers.Set, set5)
										} else {
											to1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount2].Name = set5.Name
											to1.Default.Local.HTTP.OnRateLimit.Headers.Set[setCount2].Value = set5.Value
										}
									}
								}
								if toItem.Default.Local.HTTP.OnRateLimit.Status != nil {
									to1.Default.Local.HTTP.OnRateLimit.Status = types.Int64Value(int64(*toItem.Default.Local.HTTP.OnRateLimit.Status))
								} else {
									to1.Default.Local.HTTP.OnRateLimit.Status = types.Int64Null()
								}
							}
							if toItem.Default.Local.HTTP.RequestRate == nil {
								to1.Default.Local.HTTP.RequestRate = nil
							} else {
								to1.Default.Local.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
								to1.Default.Local.HTTP.RequestRate.Interval = types.StringValue(toItem.Default.Local.HTTP.RequestRate.Interval)
								to1.Default.Local.HTTP.RequestRate.Num = types.Int64Value(int64(toItem.Default.Local.HTTP.RequestRate.Num))
							}
						}
						if toItem.Default.Local.TCP == nil {
							to1.Default.Local.TCP = nil
						} else {
							to1.Default.Local.TCP = &tfTypes.MeshRateLimitItemTCP{}
							if toItem.Default.Local.TCP.ConnectionRate == nil {
								to1.Default.Local.TCP.ConnectionRate = nil
							} else {
								to1.Default.Local.TCP.ConnectionRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
								to1.Default.Local.TCP.ConnectionRate.Interval = types.StringValue(toItem.Default.Local.TCP.ConnectionRate.Interval)
								to1.Default.Local.TCP.ConnectionRate.Num = types.Int64Value(int64(toItem.Default.Local.TCP.ConnectionRate.Num))
							}
							to1.Default.Local.TCP.Disabled = types.BoolPointerValue(toItem.Default.Local.TCP.Disabled)
						}
					}
				}
				to1.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
				if len(toItem.TargetRef.Labels) > 0 {
					to1.TargetRef.Labels = make(map[string]types.String)
					for key5, value11 := range toItem.TargetRef.Labels {
						to1.TargetRef.Labels[key5] = types.StringValue(value11)
					}
				}
				to1.TargetRef.Mesh = types.StringPointerValue(toItem.TargetRef.Mesh)
				to1.TargetRef.Name = types.StringPointerValue(toItem.TargetRef.Name)
				to1.TargetRef.Namespace = types.StringPointerValue(toItem.TargetRef.Namespace)
				to1.TargetRef.ProxyTypes = make([]types.String, 0, len(toItem.TargetRef.ProxyTypes))
				for _, v := range toItem.TargetRef.ProxyTypes {
					to1.TargetRef.ProxyTypes = append(to1.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				to1.TargetRef.SectionName = types.StringPointerValue(toItem.TargetRef.SectionName)
				if len(toItem.TargetRef.Tags) > 0 {
					to1.TargetRef.Tags = make(map[string]types.String)
					for key6, value12 := range toItem.TargetRef.Tags {
						to1.TargetRef.Tags[key6] = types.StringValue(value12)
					}
				}
				if toCount+1 > len(items1.Spec.To) {
					items1.Spec.To = append(items1.Spec.To, to1)
				} else {
					items1.Spec.To[toCount].Default = to1.Default
					items1.Spec.To[toCount].TargetRef = to1.TargetRef
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
