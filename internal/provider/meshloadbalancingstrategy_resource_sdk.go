// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshLoadBalancingStrategyResourceModel) ToSharedMeshLoadBalancingStrategyItem() *shared.MeshLoadBalancingStrategyItem {
	typeVar := shared.MeshLoadBalancingStrategyItemType(r.Type.ValueString())
	mesh := new(string)
	if !r.Mesh.IsUnknown() && !r.Mesh.IsNull() {
		*mesh = r.Mesh.ValueString()
	} else {
		mesh = nil
	}
	var name string
	name = r.Name.ValueString()

	labels := make(map[string]string)
	for labelsKey, labelsValue := range r.Labels {
		var labelsInst string
		labelsInst = labelsValue.ValueString()

		labels[labelsKey] = labelsInst
	}
	var targetRef *shared.MeshLoadBalancingStrategyItemTargetRef
	if r.Spec.TargetRef != nil {
		kind := new(shared.MeshLoadBalancingStrategyItemKind)
		if !r.Spec.TargetRef.Kind.IsUnknown() && !r.Spec.TargetRef.Kind.IsNull() {
			*kind = shared.MeshLoadBalancingStrategyItemKind(r.Spec.TargetRef.Kind.ValueString())
		} else {
			kind = nil
		}
		labels1 := make(map[string]string)
		for labelsKey1, labelsValue1 := range r.Spec.TargetRef.Labels {
			var labelsInst1 string
			labelsInst1 = labelsValue1.ValueString()

			labels1[labelsKey1] = labelsInst1
		}
		mesh1 := new(string)
		if !r.Spec.TargetRef.Mesh.IsUnknown() && !r.Spec.TargetRef.Mesh.IsNull() {
			*mesh1 = r.Spec.TargetRef.Mesh.ValueString()
		} else {
			mesh1 = nil
		}
		name1 := new(string)
		if !r.Spec.TargetRef.Name.IsUnknown() && !r.Spec.TargetRef.Name.IsNull() {
			*name1 = r.Spec.TargetRef.Name.ValueString()
		} else {
			name1 = nil
		}
		namespace := new(string)
		if !r.Spec.TargetRef.Namespace.IsUnknown() && !r.Spec.TargetRef.Namespace.IsNull() {
			*namespace = r.Spec.TargetRef.Namespace.ValueString()
		} else {
			namespace = nil
		}
		var proxyTypes []shared.MeshLoadBalancingStrategyItemProxyTypes = []shared.MeshLoadBalancingStrategyItemProxyTypes{}
		for _, proxyTypesItem := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshLoadBalancingStrategyItemProxyTypes(proxyTypesItem.ValueString()))
		}
		sectionName := new(string)
		if !r.Spec.TargetRef.SectionName.IsUnknown() && !r.Spec.TargetRef.SectionName.IsNull() {
			*sectionName = r.Spec.TargetRef.SectionName.ValueString()
		} else {
			sectionName = nil
		}
		tags := make(map[string]string)
		for tagsKey, tagsValue := range r.Spec.TargetRef.Tags {
			var tagsInst string
			tagsInst = tagsValue.ValueString()

			tags[tagsKey] = tagsInst
		}
		targetRef = &shared.MeshLoadBalancingStrategyItemTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name1,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags,
		}
	}
	var to []shared.MeshLoadBalancingStrategyItemTo = []shared.MeshLoadBalancingStrategyItemTo{}
	for _, toItem := range r.Spec.To {
		var defaultVar *shared.MeshLoadBalancingStrategyItemDefault
		if toItem.Default != nil {
			var loadBalancer *shared.LoadBalancer
			if toItem.Default.LoadBalancer != nil {
				var leastRequest *shared.LeastRequest
				if toItem.Default.LoadBalancer.LeastRequest != nil {
					var activeRequestBias *shared.ActiveRequestBias
					if toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias != nil {
						integer := new(int64)
						if !toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Integer.IsUnknown() && !toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Integer.IsNull() {
							*integer = toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Integer.ValueInt64()
						} else {
							integer = nil
						}
						if integer != nil {
							activeRequestBias = &shared.ActiveRequestBias{
								Integer: integer,
							}
						}
						str := new(string)
						if !toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Str.IsUnknown() && !toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Str.IsNull() {
							*str = toItem.Default.LoadBalancer.LeastRequest.ActiveRequestBias.Str.ValueString()
						} else {
							str = nil
						}
						if str != nil {
							activeRequestBias = &shared.ActiveRequestBias{
								Str: str,
							}
						}
					}
					choiceCount := new(int)
					if !toItem.Default.LoadBalancer.LeastRequest.ChoiceCount.IsUnknown() && !toItem.Default.LoadBalancer.LeastRequest.ChoiceCount.IsNull() {
						*choiceCount = int(toItem.Default.LoadBalancer.LeastRequest.ChoiceCount.ValueInt64())
					} else {
						choiceCount = nil
					}
					leastRequest = &shared.LeastRequest{
						ActiveRequestBias: activeRequestBias,
						ChoiceCount:       choiceCount,
					}
				}
				var maglev *shared.Maglev
				if toItem.Default.LoadBalancer.Maglev != nil {
					var hashPolicies []shared.HashPolicies = []shared.HashPolicies{}
					for _, hashPoliciesItem := range toItem.Default.LoadBalancer.Maglev.HashPolicies {
						var connection *shared.Connection
						if hashPoliciesItem.Connection != nil {
							sourceIP := new(bool)
							if !hashPoliciesItem.Connection.SourceIP.IsUnknown() && !hashPoliciesItem.Connection.SourceIP.IsNull() {
								*sourceIP = hashPoliciesItem.Connection.SourceIP.ValueBool()
							} else {
								sourceIP = nil
							}
							connection = &shared.Connection{
								SourceIP: sourceIP,
							}
						}
						var cookie *shared.Cookie
						if hashPoliciesItem.Cookie != nil {
							var name2 string
							name2 = hashPoliciesItem.Cookie.Name.ValueString()

							path := new(string)
							if !hashPoliciesItem.Cookie.Path.IsUnknown() && !hashPoliciesItem.Cookie.Path.IsNull() {
								*path = hashPoliciesItem.Cookie.Path.ValueString()
							} else {
								path = nil
							}
							ttl := new(string)
							if !hashPoliciesItem.Cookie.TTL.IsUnknown() && !hashPoliciesItem.Cookie.TTL.IsNull() {
								*ttl = hashPoliciesItem.Cookie.TTL.ValueString()
							} else {
								ttl = nil
							}
							cookie = &shared.Cookie{
								Name: name2,
								Path: path,
								TTL:  ttl,
							}
						}
						var filterState *shared.FilterState
						if hashPoliciesItem.FilterState != nil {
							var key string
							key = hashPoliciesItem.FilterState.Key.ValueString()

							filterState = &shared.FilterState{
								Key: key,
							}
						}
						var header *shared.MeshLoadBalancingStrategyItemSpecHeader
						if hashPoliciesItem.Header != nil {
							var name3 string
							name3 = hashPoliciesItem.Header.Name.ValueString()

							header = &shared.MeshLoadBalancingStrategyItemSpecHeader{
								Name: name3,
							}
						}
						var queryParameter *shared.QueryParameter
						if hashPoliciesItem.QueryParameter != nil {
							var name4 string
							name4 = hashPoliciesItem.QueryParameter.Name.ValueString()

							queryParameter = &shared.QueryParameter{
								Name: name4,
							}
						}
						terminal := new(bool)
						if !hashPoliciesItem.Terminal.IsUnknown() && !hashPoliciesItem.Terminal.IsNull() {
							*terminal = hashPoliciesItem.Terminal.ValueBool()
						} else {
							terminal = nil
						}
						type1 := shared.MeshLoadBalancingStrategyItemSpecToDefaultType(hashPoliciesItem.Type.ValueString())
						hashPolicies = append(hashPolicies, shared.HashPolicies{
							Connection:     connection,
							Cookie:         cookie,
							FilterState:    filterState,
							Header:         header,
							QueryParameter: queryParameter,
							Terminal:       terminal,
							Type:           type1,
						})
					}
					tableSize := new(int)
					if !toItem.Default.LoadBalancer.Maglev.TableSize.IsUnknown() && !toItem.Default.LoadBalancer.Maglev.TableSize.IsNull() {
						*tableSize = int(toItem.Default.LoadBalancer.Maglev.TableSize.ValueInt64())
					} else {
						tableSize = nil
					}
					maglev = &shared.Maglev{
						HashPolicies: hashPolicies,
						TableSize:    tableSize,
					}
				}
				var random *shared.MeshLoadBalancingStrategyItemRandom
				if toItem.Default.LoadBalancer.Random != nil {
					random = &shared.MeshLoadBalancingStrategyItemRandom{}
				}
				var ringHash *shared.RingHash
				if toItem.Default.LoadBalancer.RingHash != nil {
					hashFunction := new(shared.HashFunction)
					if !toItem.Default.LoadBalancer.RingHash.HashFunction.IsUnknown() && !toItem.Default.LoadBalancer.RingHash.HashFunction.IsNull() {
						*hashFunction = shared.HashFunction(toItem.Default.LoadBalancer.RingHash.HashFunction.ValueString())
					} else {
						hashFunction = nil
					}
					var hashPolicies1 []shared.MeshLoadBalancingStrategyItemHashPolicies = []shared.MeshLoadBalancingStrategyItemHashPolicies{}
					for _, hashPoliciesItem1 := range toItem.Default.LoadBalancer.RingHash.HashPolicies {
						var connection1 *shared.MeshLoadBalancingStrategyItemConnection
						if hashPoliciesItem1.Connection != nil {
							sourceIp1 := new(bool)
							if !hashPoliciesItem1.Connection.SourceIP.IsUnknown() && !hashPoliciesItem1.Connection.SourceIP.IsNull() {
								*sourceIp1 = hashPoliciesItem1.Connection.SourceIP.ValueBool()
							} else {
								sourceIp1 = nil
							}
							connection1 = &shared.MeshLoadBalancingStrategyItemConnection{
								SourceIP: sourceIp1,
							}
						}
						var cookie1 *shared.MeshLoadBalancingStrategyItemCookie
						if hashPoliciesItem1.Cookie != nil {
							var name5 string
							name5 = hashPoliciesItem1.Cookie.Name.ValueString()

							path1 := new(string)
							if !hashPoliciesItem1.Cookie.Path.IsUnknown() && !hashPoliciesItem1.Cookie.Path.IsNull() {
								*path1 = hashPoliciesItem1.Cookie.Path.ValueString()
							} else {
								path1 = nil
							}
							ttl1 := new(string)
							if !hashPoliciesItem1.Cookie.TTL.IsUnknown() && !hashPoliciesItem1.Cookie.TTL.IsNull() {
								*ttl1 = hashPoliciesItem1.Cookie.TTL.ValueString()
							} else {
								ttl1 = nil
							}
							cookie1 = &shared.MeshLoadBalancingStrategyItemCookie{
								Name: name5,
								Path: path1,
								TTL:  ttl1,
							}
						}
						var filterState1 *shared.MeshLoadBalancingStrategyItemFilterState
						if hashPoliciesItem1.FilterState != nil {
							var key1 string
							key1 = hashPoliciesItem1.FilterState.Key.ValueString()

							filterState1 = &shared.MeshLoadBalancingStrategyItemFilterState{
								Key: key1,
							}
						}
						var header1 *shared.MeshLoadBalancingStrategyItemHeader
						if hashPoliciesItem1.Header != nil {
							var name6 string
							name6 = hashPoliciesItem1.Header.Name.ValueString()

							header1 = &shared.MeshLoadBalancingStrategyItemHeader{
								Name: name6,
							}
						}
						var queryParameter1 *shared.MeshLoadBalancingStrategyItemQueryParameter
						if hashPoliciesItem1.QueryParameter != nil {
							var name7 string
							name7 = hashPoliciesItem1.QueryParameter.Name.ValueString()

							queryParameter1 = &shared.MeshLoadBalancingStrategyItemQueryParameter{
								Name: name7,
							}
						}
						terminal1 := new(bool)
						if !hashPoliciesItem1.Terminal.IsUnknown() && !hashPoliciesItem1.Terminal.IsNull() {
							*terminal1 = hashPoliciesItem1.Terminal.ValueBool()
						} else {
							terminal1 = nil
						}
						type2 := shared.MeshLoadBalancingStrategyItemSpecToType(hashPoliciesItem1.Type.ValueString())
						hashPolicies1 = append(hashPolicies1, shared.MeshLoadBalancingStrategyItemHashPolicies{
							Connection:     connection1,
							Cookie:         cookie1,
							FilterState:    filterState1,
							Header:         header1,
							QueryParameter: queryParameter1,
							Terminal:       terminal1,
							Type:           type2,
						})
					}
					maxRingSize := new(int)
					if !toItem.Default.LoadBalancer.RingHash.MaxRingSize.IsUnknown() && !toItem.Default.LoadBalancer.RingHash.MaxRingSize.IsNull() {
						*maxRingSize = int(toItem.Default.LoadBalancer.RingHash.MaxRingSize.ValueInt64())
					} else {
						maxRingSize = nil
					}
					minRingSize := new(int)
					if !toItem.Default.LoadBalancer.RingHash.MinRingSize.IsUnknown() && !toItem.Default.LoadBalancer.RingHash.MinRingSize.IsNull() {
						*minRingSize = int(toItem.Default.LoadBalancer.RingHash.MinRingSize.ValueInt64())
					} else {
						minRingSize = nil
					}
					ringHash = &shared.RingHash{
						HashFunction: hashFunction,
						HashPolicies: hashPolicies1,
						MaxRingSize:  maxRingSize,
						MinRingSize:  minRingSize,
					}
				}
				var roundRobin *shared.RoundRobin
				if toItem.Default.LoadBalancer.RoundRobin != nil {
					roundRobin = &shared.RoundRobin{}
				}
				typeVar1 := shared.MeshLoadBalancingStrategyItemSpecType(toItem.Default.LoadBalancer.Type.ValueString())
				loadBalancer = &shared.LoadBalancer{
					LeastRequest: leastRequest,
					Maglev:       maglev,
					Random:       random,
					RingHash:     ringHash,
					RoundRobin:   roundRobin,
					Type:         typeVar1,
				}
			}
			var localityAwareness *shared.LocalityAwareness
			if toItem.Default.LocalityAwareness != nil {
				var crossZone *shared.CrossZone
				if toItem.Default.LocalityAwareness.CrossZone != nil {
					var failover []shared.Failover = []shared.Failover{}
					for _, failoverItem := range toItem.Default.LocalityAwareness.CrossZone.Failover {
						var from *shared.MeshLoadBalancingStrategyItemFrom
						if failoverItem.From != nil {
							var zones []string = []string{}
							for _, zonesItem := range failoverItem.From.Zones {
								zones = append(zones, zonesItem.ValueString())
							}
							from = &shared.MeshLoadBalancingStrategyItemFrom{
								Zones: zones,
							}
						}
						typeVar2 := shared.MeshLoadBalancingStrategyItemSpecToDefaultLocalityAwarenessType(failoverItem.To.Type.ValueString())
						var zones1 []string = []string{}
						for _, zonesItem1 := range failoverItem.To.Zones {
							zones1 = append(zones1, zonesItem1.ValueString())
						}
						to1 := shared.MeshLoadBalancingStrategyItemSpecTo{
							Type:  typeVar2,
							Zones: zones1,
						}
						failover = append(failover, shared.Failover{
							From: from,
							To:   to1,
						})
					}
					var failoverThreshold *shared.FailoverThreshold
					if toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold != nil {
						var percentage shared.MeshLoadBalancingStrategyItemPercentage
						integer1 := new(int64)
						if !toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Integer.IsUnknown() && !toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Integer.IsNull() {
							*integer1 = toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Integer.ValueInt64()
						} else {
							integer1 = nil
						}
						if integer1 != nil {
							percentage = shared.MeshLoadBalancingStrategyItemPercentage{
								Integer: integer1,
							}
						}
						str1 := new(string)
						if !toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Str.IsUnknown() && !toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Str.IsNull() {
							*str1 = toItem.Default.LocalityAwareness.CrossZone.FailoverThreshold.Percentage.Str.ValueString()
						} else {
							str1 = nil
						}
						if str1 != nil {
							percentage = shared.MeshLoadBalancingStrategyItemPercentage{
								Str: str1,
							}
						}
						failoverThreshold = &shared.FailoverThreshold{
							Percentage: percentage,
						}
					}
					crossZone = &shared.CrossZone{
						Failover:          failover,
						FailoverThreshold: failoverThreshold,
					}
				}
				disabled := new(bool)
				if !toItem.Default.LocalityAwareness.Disabled.IsUnknown() && !toItem.Default.LocalityAwareness.Disabled.IsNull() {
					*disabled = toItem.Default.LocalityAwareness.Disabled.ValueBool()
				} else {
					disabled = nil
				}
				var localZone *shared.LocalZone
				if toItem.Default.LocalityAwareness.LocalZone != nil {
					var affinityTags []shared.AffinityTags = []shared.AffinityTags{}
					for _, affinityTagsItem := range toItem.Default.LocalityAwareness.LocalZone.AffinityTags {
						var key2 string
						key2 = affinityTagsItem.Key.ValueString()

						weight := new(int)
						if !affinityTagsItem.Weight.IsUnknown() && !affinityTagsItem.Weight.IsNull() {
							*weight = int(affinityTagsItem.Weight.ValueInt64())
						} else {
							weight = nil
						}
						affinityTags = append(affinityTags, shared.AffinityTags{
							Key:    key2,
							Weight: weight,
						})
					}
					localZone = &shared.LocalZone{
						AffinityTags: affinityTags,
					}
				}
				localityAwareness = &shared.LocalityAwareness{
					CrossZone: crossZone,
					Disabled:  disabled,
					LocalZone: localZone,
				}
			}
			defaultVar = &shared.MeshLoadBalancingStrategyItemDefault{
				LoadBalancer:      loadBalancer,
				LocalityAwareness: localityAwareness,
			}
		}
		kind1 := new(shared.MeshLoadBalancingStrategyItemSpecKind)
		if !toItem.TargetRef.Kind.IsUnknown() && !toItem.TargetRef.Kind.IsNull() {
			*kind1 = shared.MeshLoadBalancingStrategyItemSpecKind(toItem.TargetRef.Kind.ValueString())
		} else {
			kind1 = nil
		}
		labels2 := make(map[string]string)
		for labelsKey2, labelsValue2 := range toItem.TargetRef.Labels {
			var labelsInst2 string
			labelsInst2 = labelsValue2.ValueString()

			labels2[labelsKey2] = labelsInst2
		}
		mesh2 := new(string)
		if !toItem.TargetRef.Mesh.IsUnknown() && !toItem.TargetRef.Mesh.IsNull() {
			*mesh2 = toItem.TargetRef.Mesh.ValueString()
		} else {
			mesh2 = nil
		}
		name8 := new(string)
		if !toItem.TargetRef.Name.IsUnknown() && !toItem.TargetRef.Name.IsNull() {
			*name8 = toItem.TargetRef.Name.ValueString()
		} else {
			name8 = nil
		}
		namespace1 := new(string)
		if !toItem.TargetRef.Namespace.IsUnknown() && !toItem.TargetRef.Namespace.IsNull() {
			*namespace1 = toItem.TargetRef.Namespace.ValueString()
		} else {
			namespace1 = nil
		}
		var proxyTypes1 []shared.MeshLoadBalancingStrategyItemSpecProxyTypes = []shared.MeshLoadBalancingStrategyItemSpecProxyTypes{}
		for _, proxyTypesItem1 := range toItem.TargetRef.ProxyTypes {
			proxyTypes1 = append(proxyTypes1, shared.MeshLoadBalancingStrategyItemSpecProxyTypes(proxyTypesItem1.ValueString()))
		}
		sectionName1 := new(string)
		if !toItem.TargetRef.SectionName.IsUnknown() && !toItem.TargetRef.SectionName.IsNull() {
			*sectionName1 = toItem.TargetRef.SectionName.ValueString()
		} else {
			sectionName1 = nil
		}
		tags1 := make(map[string]string)
		for tagsKey1, tagsValue1 := range toItem.TargetRef.Tags {
			var tagsInst1 string
			tagsInst1 = tagsValue1.ValueString()

			tags1[tagsKey1] = tagsInst1
		}
		targetRef1 := shared.MeshLoadBalancingStrategyItemSpecTargetRef{
			Kind:        kind1,
			Labels:      labels2,
			Mesh:        mesh2,
			Name:        name8,
			Namespace:   namespace1,
			ProxyTypes:  proxyTypes1,
			SectionName: sectionName1,
			Tags:        tags1,
		}
		to = append(to, shared.MeshLoadBalancingStrategyItemTo{
			Default:   defaultVar,
			TargetRef: targetRef1,
		})
	}
	spec := shared.MeshLoadBalancingStrategyItemSpec{
		TargetRef: targetRef,
		To:        to,
	}
	creationTime := new(time.Time)
	if !r.CreationTime.IsUnknown() && !r.CreationTime.IsNull() {
		*creationTime, _ = time.Parse(time.RFC3339Nano, r.CreationTime.ValueString())
	} else {
		creationTime = nil
	}
	modificationTime := new(time.Time)
	if !r.ModificationTime.IsUnknown() && !r.ModificationTime.IsNull() {
		*modificationTime, _ = time.Parse(time.RFC3339Nano, r.ModificationTime.ValueString())
	} else {
		modificationTime = nil
	}
	out := shared.MeshLoadBalancingStrategyItem{
		Type:             typeVar,
		Mesh:             mesh,
		Name:             name,
		Labels:           labels,
		Spec:             spec,
		CreationTime:     creationTime,
		ModificationTime: modificationTime,
	}
	return &out
}

func (r *MeshLoadBalancingStrategyResourceModel) RefreshFromSharedMeshLoadBalancingStrategyCreateOrUpdateSuccessResponse(resp *shared.MeshLoadBalancingStrategyCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = []types.String{}
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshLoadBalancingStrategyResourceModel) RefreshFromSharedMeshLoadBalancingStrategyItem(resp *shared.MeshLoadBalancingStrategyItem) {
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
			if resp.Spec.TargetRef.Kind != nil {
				r.Spec.TargetRef.Kind = types.StringValue(string(*resp.Spec.TargetRef.Kind))
			} else {
				r.Spec.TargetRef.Kind = types.StringNull()
			}
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String)
				for key1, value1 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key1] = types.StringValue(value1)
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
								failover1.From.Zones = []types.String{}
								for _, v := range failoverItem.From.Zones {
									failover1.From.Zones = append(failover1.From.Zones, types.StringValue(v))
								}
							}
							failover1.To.Type = types.StringValue(string(failoverItem.To.Type))
							failover1.To.Zones = []types.String{}
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
			if toItem.TargetRef.Kind != nil {
				to1.TargetRef.Kind = types.StringValue(string(*toItem.TargetRef.Kind))
			} else {
				to1.TargetRef.Kind = types.StringNull()
			}
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String)
				for key6, value3 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key6] = types.StringValue(value3)
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