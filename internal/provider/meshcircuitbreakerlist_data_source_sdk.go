// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshCircuitBreakerListDataSourceModel) RefreshFromSharedMeshCircuitBreakerList(resp *shared.MeshCircuitBreakerList) {
	if resp != nil {
		r.Items = []tfTypes.MeshCircuitBreakerItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshCircuitBreakerItem
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
			items1.Spec.From = []tfTypes.MeshCircuitBreakerItemFrom{}
			for fromCount, fromItem := range itemsItem.Spec.From {
				var from1 tfTypes.MeshCircuitBreakerItemFrom
				if fromItem.Default == nil {
					from1.Default = nil
				} else {
					from1.Default = &tfTypes.MeshCircuitBreakerItemDefault{}
					if fromItem.Default.ConnectionLimits == nil {
						from1.Default.ConnectionLimits = nil
					} else {
						from1.Default.ConnectionLimits = &tfTypes.ConnectionLimits{}
						if fromItem.Default.ConnectionLimits.MaxConnectionPools != nil {
							from1.Default.ConnectionLimits.MaxConnectionPools = types.Int64Value(int64(*fromItem.Default.ConnectionLimits.MaxConnectionPools))
						} else {
							from1.Default.ConnectionLimits.MaxConnectionPools = types.Int64Null()
						}
						if fromItem.Default.ConnectionLimits.MaxConnections != nil {
							from1.Default.ConnectionLimits.MaxConnections = types.Int64Value(int64(*fromItem.Default.ConnectionLimits.MaxConnections))
						} else {
							from1.Default.ConnectionLimits.MaxConnections = types.Int64Null()
						}
						if fromItem.Default.ConnectionLimits.MaxPendingRequests != nil {
							from1.Default.ConnectionLimits.MaxPendingRequests = types.Int64Value(int64(*fromItem.Default.ConnectionLimits.MaxPendingRequests))
						} else {
							from1.Default.ConnectionLimits.MaxPendingRequests = types.Int64Null()
						}
						if fromItem.Default.ConnectionLimits.MaxRequests != nil {
							from1.Default.ConnectionLimits.MaxRequests = types.Int64Value(int64(*fromItem.Default.ConnectionLimits.MaxRequests))
						} else {
							from1.Default.ConnectionLimits.MaxRequests = types.Int64Null()
						}
						if fromItem.Default.ConnectionLimits.MaxRetries != nil {
							from1.Default.ConnectionLimits.MaxRetries = types.Int64Value(int64(*fromItem.Default.ConnectionLimits.MaxRetries))
						} else {
							from1.Default.ConnectionLimits.MaxRetries = types.Int64Null()
						}
					}
					if fromItem.Default.OutlierDetection == nil {
						from1.Default.OutlierDetection = nil
					} else {
						from1.Default.OutlierDetection = &tfTypes.OutlierDetection{}
						from1.Default.OutlierDetection.BaseEjectionTime = types.StringPointerValue(fromItem.Default.OutlierDetection.BaseEjectionTime)
						if fromItem.Default.OutlierDetection.Detectors == nil {
							from1.Default.OutlierDetection.Detectors = nil
						} else {
							from1.Default.OutlierDetection.Detectors = &tfTypes.Detectors{}
							if fromItem.Default.OutlierDetection.Detectors.FailurePercentage == nil {
								from1.Default.OutlierDetection.Detectors.FailurePercentage = nil
							} else {
								from1.Default.OutlierDetection.Detectors.FailurePercentage = &tfTypes.FailurePercentage{}
								if fromItem.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts != nil {
									from1.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts))
								} else {
									from1.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts = types.Int64Null()
								}
								if fromItem.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume != nil {
									from1.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume))
								} else {
									from1.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume = types.Int64Null()
								}
								if fromItem.Default.OutlierDetection.Detectors.FailurePercentage.Threshold != nil {
									from1.Default.OutlierDetection.Detectors.FailurePercentage.Threshold = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.FailurePercentage.Threshold))
								} else {
									from1.Default.OutlierDetection.Detectors.FailurePercentage.Threshold = types.Int64Null()
								}
							}
							if fromItem.Default.OutlierDetection.Detectors.GatewayFailures == nil {
								from1.Default.OutlierDetection.Detectors.GatewayFailures = nil
							} else {
								from1.Default.OutlierDetection.Detectors.GatewayFailures = &tfTypes.GatewayFailures{}
								if fromItem.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive != nil {
									from1.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive))
								} else {
									from1.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive = types.Int64Null()
								}
							}
							if fromItem.Default.OutlierDetection.Detectors.LocalOriginFailures == nil {
								from1.Default.OutlierDetection.Detectors.LocalOriginFailures = nil
							} else {
								from1.Default.OutlierDetection.Detectors.LocalOriginFailures = &tfTypes.GatewayFailures{}
								if fromItem.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive != nil {
									from1.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive))
								} else {
									from1.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive = types.Int64Null()
								}
							}
							if fromItem.Default.OutlierDetection.Detectors.SuccessRate == nil {
								from1.Default.OutlierDetection.Detectors.SuccessRate = nil
							} else {
								from1.Default.OutlierDetection.Detectors.SuccessRate = &tfTypes.SuccessRate{}
								if fromItem.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts != nil {
									from1.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts))
								} else {
									from1.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts = types.Int64Null()
								}
								if fromItem.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume != nil {
									from1.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume))
								} else {
									from1.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume = types.Int64Null()
								}
								if fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor == nil {
									from1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor = nil
								} else {
									from1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor = &tfTypes.Mode{}
									if fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer != nil {
										from1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer = types.Int64PointerValue(fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer)
									}
									if fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str != nil {
										from1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str = types.StringPointerValue(fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str)
									}
								}
							}
							if fromItem.Default.OutlierDetection.Detectors.TotalFailures == nil {
								from1.Default.OutlierDetection.Detectors.TotalFailures = nil
							} else {
								from1.Default.OutlierDetection.Detectors.TotalFailures = &tfTypes.GatewayFailures{}
								if fromItem.Default.OutlierDetection.Detectors.TotalFailures.Consecutive != nil {
									from1.Default.OutlierDetection.Detectors.TotalFailures.Consecutive = types.Int64Value(int64(*fromItem.Default.OutlierDetection.Detectors.TotalFailures.Consecutive))
								} else {
									from1.Default.OutlierDetection.Detectors.TotalFailures.Consecutive = types.Int64Null()
								}
							}
						}
						from1.Default.OutlierDetection.Disabled = types.BoolPointerValue(fromItem.Default.OutlierDetection.Disabled)
						from1.Default.OutlierDetection.Interval = types.StringPointerValue(fromItem.Default.OutlierDetection.Interval)
						if fromItem.Default.OutlierDetection.MaxEjectionPercent != nil {
							from1.Default.OutlierDetection.MaxEjectionPercent = types.Int64Value(int64(*fromItem.Default.OutlierDetection.MaxEjectionPercent))
						} else {
							from1.Default.OutlierDetection.MaxEjectionPercent = types.Int64Null()
						}
						from1.Default.OutlierDetection.SplitExternalAndLocalErrors = types.BoolPointerValue(fromItem.Default.OutlierDetection.SplitExternalAndLocalErrors)
					}
				}
				if fromItem.TargetRef.Kind != nil {
					from1.TargetRef.Kind = types.StringValue(string(*fromItem.TargetRef.Kind))
				} else {
					from1.TargetRef.Kind = types.StringNull()
				}
				if len(fromItem.TargetRef.Labels) > 0 {
					from1.TargetRef.Labels = make(map[string]types.String)
					for key1, value1 := range fromItem.TargetRef.Labels {
						from1.TargetRef.Labels[key1] = types.StringValue(value1)
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
					for key3, value3 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key3] = types.StringValue(value3)
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
					for key4, value4 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key4] = types.StringValue(value4)
					}
				}
			}
			items1.Spec.To = []tfTypes.MeshCircuitBreakerItemFrom{}
			for toCount, toItem := range itemsItem.Spec.To {
				var to1 tfTypes.MeshCircuitBreakerItemFrom
				if toItem.Default == nil {
					to1.Default = nil
				} else {
					to1.Default = &tfTypes.MeshCircuitBreakerItemDefault{}
					if toItem.Default.ConnectionLimits == nil {
						to1.Default.ConnectionLimits = nil
					} else {
						to1.Default.ConnectionLimits = &tfTypes.ConnectionLimits{}
						if toItem.Default.ConnectionLimits.MaxConnectionPools != nil {
							to1.Default.ConnectionLimits.MaxConnectionPools = types.Int64Value(int64(*toItem.Default.ConnectionLimits.MaxConnectionPools))
						} else {
							to1.Default.ConnectionLimits.MaxConnectionPools = types.Int64Null()
						}
						if toItem.Default.ConnectionLimits.MaxConnections != nil {
							to1.Default.ConnectionLimits.MaxConnections = types.Int64Value(int64(*toItem.Default.ConnectionLimits.MaxConnections))
						} else {
							to1.Default.ConnectionLimits.MaxConnections = types.Int64Null()
						}
						if toItem.Default.ConnectionLimits.MaxPendingRequests != nil {
							to1.Default.ConnectionLimits.MaxPendingRequests = types.Int64Value(int64(*toItem.Default.ConnectionLimits.MaxPendingRequests))
						} else {
							to1.Default.ConnectionLimits.MaxPendingRequests = types.Int64Null()
						}
						if toItem.Default.ConnectionLimits.MaxRequests != nil {
							to1.Default.ConnectionLimits.MaxRequests = types.Int64Value(int64(*toItem.Default.ConnectionLimits.MaxRequests))
						} else {
							to1.Default.ConnectionLimits.MaxRequests = types.Int64Null()
						}
						if toItem.Default.ConnectionLimits.MaxRetries != nil {
							to1.Default.ConnectionLimits.MaxRetries = types.Int64Value(int64(*toItem.Default.ConnectionLimits.MaxRetries))
						} else {
							to1.Default.ConnectionLimits.MaxRetries = types.Int64Null()
						}
					}
					if toItem.Default.OutlierDetection == nil {
						to1.Default.OutlierDetection = nil
					} else {
						to1.Default.OutlierDetection = &tfTypes.OutlierDetection{}
						to1.Default.OutlierDetection.BaseEjectionTime = types.StringPointerValue(toItem.Default.OutlierDetection.BaseEjectionTime)
						if toItem.Default.OutlierDetection.Detectors == nil {
							to1.Default.OutlierDetection.Detectors = nil
						} else {
							to1.Default.OutlierDetection.Detectors = &tfTypes.Detectors{}
							if toItem.Default.OutlierDetection.Detectors.FailurePercentage == nil {
								to1.Default.OutlierDetection.Detectors.FailurePercentage = nil
							} else {
								to1.Default.OutlierDetection.Detectors.FailurePercentage = &tfTypes.FailurePercentage{}
								if toItem.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts != nil {
									to1.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts))
								} else {
									to1.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts = types.Int64Null()
								}
								if toItem.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume != nil {
									to1.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume))
								} else {
									to1.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume = types.Int64Null()
								}
								if toItem.Default.OutlierDetection.Detectors.FailurePercentage.Threshold != nil {
									to1.Default.OutlierDetection.Detectors.FailurePercentage.Threshold = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.FailurePercentage.Threshold))
								} else {
									to1.Default.OutlierDetection.Detectors.FailurePercentage.Threshold = types.Int64Null()
								}
							}
							if toItem.Default.OutlierDetection.Detectors.GatewayFailures == nil {
								to1.Default.OutlierDetection.Detectors.GatewayFailures = nil
							} else {
								to1.Default.OutlierDetection.Detectors.GatewayFailures = &tfTypes.GatewayFailures{}
								if toItem.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive != nil {
									to1.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive))
								} else {
									to1.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive = types.Int64Null()
								}
							}
							if toItem.Default.OutlierDetection.Detectors.LocalOriginFailures == nil {
								to1.Default.OutlierDetection.Detectors.LocalOriginFailures = nil
							} else {
								to1.Default.OutlierDetection.Detectors.LocalOriginFailures = &tfTypes.GatewayFailures{}
								if toItem.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive != nil {
									to1.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive))
								} else {
									to1.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive = types.Int64Null()
								}
							}
							if toItem.Default.OutlierDetection.Detectors.SuccessRate == nil {
								to1.Default.OutlierDetection.Detectors.SuccessRate = nil
							} else {
								to1.Default.OutlierDetection.Detectors.SuccessRate = &tfTypes.SuccessRate{}
								if toItem.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts != nil {
									to1.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts))
								} else {
									to1.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts = types.Int64Null()
								}
								if toItem.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume != nil {
									to1.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume))
								} else {
									to1.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume = types.Int64Null()
								}
								if toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor == nil {
									to1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor = nil
								} else {
									to1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor = &tfTypes.Mode{}
									if toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer != nil {
										to1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer = types.Int64PointerValue(toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer)
									}
									if toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str != nil {
										to1.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str = types.StringPointerValue(toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str)
									}
								}
							}
							if toItem.Default.OutlierDetection.Detectors.TotalFailures == nil {
								to1.Default.OutlierDetection.Detectors.TotalFailures = nil
							} else {
								to1.Default.OutlierDetection.Detectors.TotalFailures = &tfTypes.GatewayFailures{}
								if toItem.Default.OutlierDetection.Detectors.TotalFailures.Consecutive != nil {
									to1.Default.OutlierDetection.Detectors.TotalFailures.Consecutive = types.Int64Value(int64(*toItem.Default.OutlierDetection.Detectors.TotalFailures.Consecutive))
								} else {
									to1.Default.OutlierDetection.Detectors.TotalFailures.Consecutive = types.Int64Null()
								}
							}
						}
						to1.Default.OutlierDetection.Disabled = types.BoolPointerValue(toItem.Default.OutlierDetection.Disabled)
						to1.Default.OutlierDetection.Interval = types.StringPointerValue(toItem.Default.OutlierDetection.Interval)
						if toItem.Default.OutlierDetection.MaxEjectionPercent != nil {
							to1.Default.OutlierDetection.MaxEjectionPercent = types.Int64Value(int64(*toItem.Default.OutlierDetection.MaxEjectionPercent))
						} else {
							to1.Default.OutlierDetection.MaxEjectionPercent = types.Int64Null()
						}
						to1.Default.OutlierDetection.SplitExternalAndLocalErrors = types.BoolPointerValue(toItem.Default.OutlierDetection.SplitExternalAndLocalErrors)
					}
				}
				if toItem.TargetRef.Kind != nil {
					to1.TargetRef.Kind = types.StringValue(string(*toItem.TargetRef.Kind))
				} else {
					to1.TargetRef.Kind = types.StringNull()
				}
				if len(toItem.TargetRef.Labels) > 0 {
					to1.TargetRef.Labels = make(map[string]types.String)
					for key5, value5 := range toItem.TargetRef.Labels {
						to1.TargetRef.Labels[key5] = types.StringValue(value5)
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
					for key6, value6 := range toItem.TargetRef.Tags {
						to1.TargetRef.Tags[key6] = types.StringValue(value6)
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