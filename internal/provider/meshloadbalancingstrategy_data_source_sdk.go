// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshLoadBalancingStrategyDataSourceModel) RefreshFromSharedMeshLoadBalancingStrategyItem(resp *shared.MeshLoadBalancingStrategyItem) {
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
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String)
				for key1, value1 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key1] = types.StringValue(value1)
				}
			}
			r.Spec.TargetRef.Mesh = types.StringPointerValue(resp.Spec.TargetRef.Mesh)
			r.Spec.TargetRef.Name = types.StringPointerValue(resp.Spec.TargetRef.Name)
			r.Spec.TargetRef.Namespace = types.StringPointerValue(resp.Spec.TargetRef.Namespace)
			r.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(resp.Spec.TargetRef.ProxyTypes))
			for _, v := range resp.Spec.TargetRef.ProxyTypes {
				r.Spec.TargetRef.ProxyTypes = append(r.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			r.Spec.TargetRef.SectionName = types.StringPointerValue(resp.Spec.TargetRef.SectionName)
			if len(resp.Spec.TargetRef.Tags) > 0 {
				r.Spec.TargetRef.Tags = make(map[string]types.String)
				for key2, value2 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key2] = types.StringValue(value2)
				}
			}
		}
		r.Spec.To = []tfTypes.MeshLoadBalancingStrategyItemTo{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshLoadBalancingStrategyItemTo
			if toItem.Default == nil {
				to1.Default = nil
			} else {
				to1.Default = &tfTypes.MeshLoadBalancingStrategyItemDefault{}
				if toItem.Default.LoadBalancer == nil {
					to1.Default.LoadBalancer = nil
				} else {
					to1.Default.LoadBalancer = &tfTypes.LoadBalancer{}
					if toItem.Default.LoadBalancer.LeastRequest == nil {
						to1.Default.LoadBalancer.LeastRequest = nil
					} else {
						to1.Default.LoadBalancer.LeastRequest = &tfTypes.LeastRequest{}
						if toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias == nil {
							to1.Default.LoadBalancer.LeastRequest.ActiveRequestBias = nil
						} else {
							to1.Default.LoadBalancer.LeastRequest.ActiveRequestBias = &tfTypes.Mode{}
							if toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Integer != nil {
								to1.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Integer = types.Int64PointerValue(toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Integer)
							}
							if toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Str != nil {
								to1.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Str = types.StringPointerValue(toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Str)
							}
						}
						if toItem.Default.LoadBalancer.LeastRequest.ChoiceCount != nil {
							to1.Default.LoadBalancer.LeastRequest.ChoiceCount = types.Int64Value(int64(*toItem.Default.LoadBalancer.LeastRequest.ChoiceCount))
						} else {
							to1.Default.LoadBalancer.LeastRequest.ChoiceCount = types.Int64Null()
						}
					}
					if toItem.Default.LoadBalancer.Maglev == nil {
						to1.Default.LoadBalancer.Maglev = nil
					} else {
						to1.Default.LoadBalancer.Maglev = &tfTypes.Maglev{}
						to1.Default.LoadBalancer.Maglev.HashPolicies = []tfTypes.HashPolicies{}
						for hashPoliciesCount, hashPoliciesItem := range toItem.Default.LoadBalancer.Maglev.HashPolicies {
							var hashPolicies1 tfTypes.HashPolicies
							if hashPoliciesItem.Connection == nil {
								hashPolicies1.Connection = nil
							} else {
								hashPolicies1.Connection = &tfTypes.Connection{}
								hashPolicies1.Connection.SourceIP = types.BoolPointerValue(hashPoliciesItem.Connection.SourceIP)
							}
							if hashPoliciesItem.Cookie == nil {
								hashPolicies1.Cookie = nil
							} else {
								hashPolicies1.Cookie = &tfTypes.Cookie{}
								hashPolicies1.Cookie.Name = types.StringValue(hashPoliciesItem.Cookie.Name)
								hashPolicies1.Cookie.Path = types.StringPointerValue(hashPoliciesItem.Cookie.Path)
								hashPolicies1.Cookie.TTL = types.StringPointerValue(hashPoliciesItem.Cookie.TTL)
							}
							if hashPoliciesItem.FilterState == nil {
								hashPolicies1.FilterState = nil
							} else {
								hashPolicies1.FilterState = &tfTypes.FilterState{}
								hashPolicies1.FilterState.Key = types.StringValue(hashPoliciesItem.FilterState.Key)
							}
							if hashPoliciesItem.Header == nil {
								hashPolicies1.Header = nil
							} else {
								hashPolicies1.Header = &tfTypes.MeshLoadBalancingStrategyItemSpecHeader{}
								hashPolicies1.Header.Name = types.StringValue(hashPoliciesItem.Header.Name)
							}
							if hashPoliciesItem.QueryParameter == nil {
								hashPolicies1.QueryParameter = nil
							} else {
								hashPolicies1.QueryParameter = &tfTypes.MeshLoadBalancingStrategyItemSpecHeader{}
								hashPolicies1.QueryParameter.Name = types.StringValue(hashPoliciesItem.QueryParameter.Name)
							}
							hashPolicies1.Terminal = types.BoolPointerValue(hashPoliciesItem.Terminal)
							hashPolicies1.Type = types.StringValue(string(hashPoliciesItem.Type))
							if hashPoliciesCount+1 > len(to1.Default.LoadBalancer.Maglev.HashPolicies) {
								to1.Default.LoadBalancer.Maglev.HashPolicies = append(to1.Default.LoadBalancer.Maglev.HashPolicies, hashPolicies1)
							} else {
								to1.Default.LoadBalancer.Maglev.HashPolicies[hashPoliciesCount].Connection = hashPolicies1.Connection
								to1.Default.LoadBalancer.Maglev.HashPolicies[hashPoliciesCount].Cookie = hashPolicies1.Cookie
								to1.Default.LoadBalancer.Maglev.HashPolicies[hashPoliciesCount].FilterState = hashPolicies1.FilterState
								to1.Default.LoadBalancer.Maglev.HashPolicies[hashPoliciesCount].Header = hashPolicies1.Header
								to1.Default.LoadBalancer.Maglev.HashPolicies[hashPoliciesCount].QueryParameter = hashPolicies1.QueryParameter
								to1.Default.LoadBalancer.Maglev.HashPolicies[hashPoliciesCount].Terminal = hashPolicies1.Terminal
								to1.Default.LoadBalancer.Maglev.HashPolicies[hashPoliciesCount].Type = hashPolicies1.Type
							}
						}
						if toItem.Default.LoadBalancer.Maglev.TableSize != nil {
							to1.Default.LoadBalancer.Maglev.TableSize = types.Int64Value(int64(*toItem.Default.LoadBalancer.Maglev.TableSize))
						} else {
							to1.Default.LoadBalancer.Maglev.TableSize = types.Int64Null()
						}
					}
					if toItem.Default.LoadBalancer.Random == nil {
						to1.Default.LoadBalancer.Random = nil
					} else {
						to1.Default.LoadBalancer.Random = &tfTypes.Metadata{}
					}
					if toItem.Default.LoadBalancer.RingHash == nil {
						to1.Default.LoadBalancer.RingHash = nil
					} else {
						to1.Default.LoadBalancer.RingHash = &tfTypes.RingHash{}
						if toItem.Default.LoadBalancer.RingHash.HashFunction != nil {
							to1.Default.LoadBalancer.RingHash.HashFunction = types.StringValue(string(*toItem.Default.LoadBalancer.RingHash.HashFunction))
						} else {
							to1.Default.LoadBalancer.RingHash.HashFunction = types.StringNull()
						}
						to1.Default.LoadBalancer.RingHash.HashPolicies = []tfTypes.HashPolicies{}
						for hashPoliciesCount1, hashPoliciesItem1 := range toItem.Default.LoadBalancer.RingHash.HashPolicies {
							var hashPolicies3 tfTypes.HashPolicies
							if hashPoliciesItem1.Connection == nil {
								hashPolicies3.Connection = nil
							} else {
								hashPolicies3.Connection = &tfTypes.Connection{}
								hashPolicies3.Connection.SourceIP = types.BoolPointerValue(hashPoliciesItem1.Connection.SourceIP)
							}
							if hashPoliciesItem1.Cookie == nil {
								hashPolicies3.Cookie = nil
							} else {
								hashPolicies3.Cookie = &tfTypes.Cookie{}
								hashPolicies3.Cookie.Name = types.StringValue(hashPoliciesItem1.Cookie.Name)
								hashPolicies3.Cookie.Path = types.StringPointerValue(hashPoliciesItem1.Cookie.Path)
								hashPolicies3.Cookie.TTL = types.StringPointerValue(hashPoliciesItem1.Cookie.TTL)
							}
							if hashPoliciesItem1.FilterState == nil {
								hashPolicies3.FilterState = nil
							} else {
								hashPolicies3.FilterState = &tfTypes.FilterState{}
								hashPolicies3.FilterState.Key = types.StringValue(hashPoliciesItem1.FilterState.Key)
							}
							if hashPoliciesItem1.Header == nil {
								hashPolicies3.Header = nil
							} else {
								hashPolicies3.Header = &tfTypes.MeshLoadBalancingStrategyItemSpecHeader{}
								hashPolicies3.Header.Name = types.StringValue(hashPoliciesItem1.Header.Name)
							}
							if hashPoliciesItem1.QueryParameter == nil {
								hashPolicies3.QueryParameter = nil
							} else {
								hashPolicies3.QueryParameter = &tfTypes.MeshLoadBalancingStrategyItemSpecHeader{}
								hashPolicies3.QueryParameter.Name = types.StringValue(hashPoliciesItem1.QueryParameter.Name)
							}
							hashPolicies3.Terminal = types.BoolPointerValue(hashPoliciesItem1.Terminal)
							hashPolicies3.Type = types.StringValue(string(hashPoliciesItem1.Type))
							if hashPoliciesCount1+1 > len(to1.Default.LoadBalancer.RingHash.HashPolicies) {
								to1.Default.LoadBalancer.RingHash.HashPolicies = append(to1.Default.LoadBalancer.RingHash.HashPolicies, hashPolicies3)
							} else {
								to1.Default.LoadBalancer.RingHash.HashPolicies[hashPoliciesCount1].Connection = hashPolicies3.Connection
								to1.Default.LoadBalancer.RingHash.HashPolicies[hashPoliciesCount1].Cookie = hashPolicies3.Cookie
								to1.Default.LoadBalancer.RingHash.HashPolicies[hashPoliciesCount1].FilterState = hashPolicies3.FilterState
								to1.Default.LoadBalancer.RingHash.HashPolicies[hashPoliciesCount1].Header = hashPolicies3.Header
								to1.Default.LoadBalancer.RingHash.HashPolicies[hashPoliciesCount1].QueryParameter = hashPolicies3.QueryParameter
								to1.Default.LoadBalancer.RingHash.HashPolicies[hashPoliciesCount1].Terminal = hashPolicies3.Terminal
								to1.Default.LoadBalancer.RingHash.HashPolicies[hashPoliciesCount1].Type = hashPolicies3.Type
							}
						}
						if toItem.Default.LoadBalancer.RingHash.MaxRingSize != nil {
							to1.Default.LoadBalancer.RingHash.MaxRingSize = types.Int64Value(int64(*toItem.Default.LoadBalancer.RingHash.MaxRingSize))
						} else {
							to1.Default.LoadBalancer.RingHash.MaxRingSize = types.Int64Null()
						}
						if toItem.Default.LoadBalancer.RingHash.MinRingSize != nil {
							to1.Default.LoadBalancer.RingHash.MinRingSize = types.Int64Value(int64(*toItem.Default.LoadBalancer.RingHash.MinRingSize))
						} else {
							to1.Default.LoadBalancer.RingHash.MinRingSize = types.Int64Null()
						}
					}
					if toItem.Default.LoadBalancer.RoundRobin == nil {
						to1.Default.LoadBalancer.RoundRobin = nil
					} else {
						to1.Default.LoadBalancer.RoundRobin = &tfTypes.Metadata{}
					}
					to1.Default.LoadBalancer.Type = types.StringValue(string(toItem.Default.LoadBalancer.Type))
				}
				if toItem.Default.LocalityAwareness == nil {
					to1.Default.LocalityAwareness = nil
				} else {
					to1.Default.LocalityAwareness = &tfTypes.LocalityAwareness{}
					if toItem.Default.LocalityAwareness.CrossZone == nil {
						to1.Default.LocalityAwareness.CrossZone = nil
					} else {
						to1.Default.LocalityAwareness.CrossZone = &tfTypes.CrossZone{}
						to1.Default.LocalityAwareness.CrossZone.Failover = []tfTypes.Failover{}
						for failoverCount, failoverItem := range toItem.Default.LocalityAwareness.CrossZone.Failover {
							var failover1 tfTypes.Failover
							if failoverItem.From == nil {
								failover1.From = nil
							} else {
								failover1.From = &tfTypes.MeshLoadBalancingStrategyItemFrom{}
								failover1.From.Zones = make([]types.String, 0, len(failoverItem.From.Zones))
								for _, v := range failoverItem.From.Zones {
									failover1.From.Zones = append(failover1.From.Zones, types.StringValue(v))
								}
							}
							failover1.To.Type = types.StringValue(string(failoverItem.To.Type))
							failover1.To.Zones = make([]types.String, 0, len(failoverItem.To.Zones))
							for _, v := range failoverItem.To.Zones {
								failover1.To.Zones = append(failover1.To.Zones, types.StringValue(v))
							}
							if failoverCount+1 > len(to1.Default.LocalityAwareness.CrossZone.Failover) {
								to1.Default.LocalityAwareness.CrossZone.Failover = append(to1.Default.LocalityAwareness.CrossZone.Failover, failover1)
							} else {
								to1.Default.LocalityAwareness.CrossZone.Failover[failoverCount].From = failover1.From
								to1.Default.LocalityAwareness.CrossZone.Failover[failoverCount].To = failover1.To
							}
						}
						if toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold == nil {
							to1.Default.LocalityAwareness.CrossZone.FailoverThreshold = nil
						} else {
							to1.Default.LocalityAwareness.CrossZone.FailoverThreshold = &tfTypes.FailoverThreshold{}
							if toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Integer != nil {
								to1.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Integer = types.Int64PointerValue(toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Integer)
							}
							if toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Str != nil {
								to1.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Str = types.StringPointerValue(toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Str)
							}
						}
					}
					to1.Default.LocalityAwareness.Disabled = types.BoolPointerValue(toItem.Default.LocalityAwareness.Disabled)
					if toItem.Default.LocalityAwareness.LocalZone == nil {
						to1.Default.LocalityAwareness.LocalZone = nil
					} else {
						to1.Default.LocalityAwareness.LocalZone = &tfTypes.LocalZone{}
						to1.Default.LocalityAwareness.LocalZone.AffinityTags = []tfTypes.AffinityTags{}
						for affinityTagsCount, affinityTagsItem := range toItem.Default.LocalityAwareness.LocalZone.AffinityTags {
							var affinityTags1 tfTypes.AffinityTags
							affinityTags1.Key = types.StringValue(affinityTagsItem.Key)
							if affinityTagsItem.Weight != nil {
								affinityTags1.Weight = types.Int64Value(int64(*affinityTagsItem.Weight))
							} else {
								affinityTags1.Weight = types.Int64Null()
							}
							if affinityTagsCount+1 > len(to1.Default.LocalityAwareness.LocalZone.AffinityTags) {
								to1.Default.LocalityAwareness.LocalZone.AffinityTags = append(to1.Default.LocalityAwareness.LocalZone.AffinityTags, affinityTags1)
							} else {
								to1.Default.LocalityAwareness.LocalZone.AffinityTags[affinityTagsCount].Key = affinityTags1.Key
								to1.Default.LocalityAwareness.LocalZone.AffinityTags[affinityTagsCount].Weight = affinityTags1.Weight
							}
						}
					}
				}
			}
			to1.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String)
				for key6, value3 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key6] = types.StringValue(value3)
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
				for key7, value4 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key7] = types.StringValue(value4)
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
