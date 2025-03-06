// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshGlobalRateLimitResourceModel) ToSharedMeshGlobalRateLimitItemInput() *shared.MeshGlobalRateLimitItemInput {
	typeVar := shared.MeshGlobalRateLimitItemType(r.Type.ValueString())
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
	var from []shared.MeshGlobalRateLimitItemFrom = []shared.MeshGlobalRateLimitItemFrom{}
	for _, fromItem := range r.Spec.From {
		var defaultVar *shared.MeshGlobalRateLimitItemDefault
		if fromItem.Default != nil {
			limitOnServiceFail := new(bool)
			if !fromItem.Default.Backend.RateLimitService.LimitOnServiceFail.IsUnknown() && !fromItem.Default.Backend.RateLimitService.LimitOnServiceFail.IsNull() {
				*limitOnServiceFail = fromItem.Default.Backend.RateLimitService.LimitOnServiceFail.ValueBool()
			} else {
				limitOnServiceFail = nil
			}
			timeout := new(string)
			if !fromItem.Default.Backend.RateLimitService.Timeout.IsUnknown() && !fromItem.Default.Backend.RateLimitService.Timeout.IsNull() {
				*timeout = fromItem.Default.Backend.RateLimitService.Timeout.ValueString()
			} else {
				timeout = nil
			}
			url := new(string)
			if !fromItem.Default.Backend.RateLimitService.URL.IsUnknown() && !fromItem.Default.Backend.RateLimitService.URL.IsNull() {
				*url = fromItem.Default.Backend.RateLimitService.URL.ValueString()
			} else {
				url = nil
			}
			rateLimitService := shared.RateLimitService{
				LimitOnServiceFail: limitOnServiceFail,
				Timeout:            timeout,
				URL:                url,
			}
			backend := shared.Backend{
				RateLimitService: rateLimitService,
			}
			disabled := new(bool)
			if !fromItem.Default.HTTP.Disabled.IsUnknown() && !fromItem.Default.HTTP.Disabled.IsNull() {
				*disabled = fromItem.Default.HTTP.Disabled.ValueBool()
			} else {
				disabled = nil
			}
			var onRateLimit *shared.OnRateLimit
			if fromItem.Default.HTTP.OnRateLimit != nil {
				var headers *shared.MeshGlobalRateLimitItemSpecHeaders
				if fromItem.Default.HTTP.OnRateLimit.Headers != nil {
					var add []shared.MeshGlobalRateLimitItemSpecAdd = []shared.MeshGlobalRateLimitItemSpecAdd{}
					for _, addItem := range fromItem.Default.HTTP.OnRateLimit.Headers.Add {
						var name1 string
						name1 = addItem.Name.ValueString()

						var value string
						value = addItem.Value.ValueString()

						add = append(add, shared.MeshGlobalRateLimitItemSpecAdd{
							Name:  name1,
							Value: value,
						})
					}
					var set []shared.MeshGlobalRateLimitItemSpecSet = []shared.MeshGlobalRateLimitItemSpecSet{}
					for _, setItem := range fromItem.Default.HTTP.OnRateLimit.Headers.Set {
						var name2 string
						name2 = setItem.Name.ValueString()

						var value1 string
						value1 = setItem.Value.ValueString()

						set = append(set, shared.MeshGlobalRateLimitItemSpecSet{
							Name:  name2,
							Value: value1,
						})
					}
					headers = &shared.MeshGlobalRateLimitItemSpecHeaders{
						Add: add,
						Set: set,
					}
				}
				status := new(int)
				if !fromItem.Default.HTTP.OnRateLimit.Status.IsUnknown() && !fromItem.Default.HTTP.OnRateLimit.Status.IsNull() {
					*status = int(fromItem.Default.HTTP.OnRateLimit.Status.ValueInt64())
				} else {
					status = nil
				}
				onRateLimit = &shared.OnRateLimit{
					Headers: headers,
					Status:  status,
				}
			}
			var ratelimitOnRequest []shared.RatelimitOnRequest = []shared.RatelimitOnRequest{}
			for _, ratelimitOnRequestItem := range fromItem.Default.HTTP.RatelimitOnRequest {
				kind := shared.MeshGlobalRateLimitItemSpecFromKind(ratelimitOnRequestItem.Kind.ValueString())
				var limits []shared.Limits = []shared.Limits{}
				for _, limitsItem := range ratelimitOnRequestItem.Limits {
					var requestRate *shared.MeshGlobalRateLimitItemSpecFromRequestRate
					if limitsItem.RequestRate != nil {
						var interval string
						interval = limitsItem.RequestRate.Interval.ValueString()

						var num int
						num = int(limitsItem.RequestRate.Num.ValueInt64())

						requestRate = &shared.MeshGlobalRateLimitItemSpecFromRequestRate{
							Interval: interval,
							Num:      num,
						}
					}
					var value2 string
					value2 = limitsItem.Value.ValueString()

					limits = append(limits, shared.Limits{
						RequestRate: requestRate,
						Value:       value2,
					})
				}
				var name3 string
				name3 = ratelimitOnRequestItem.Name.ValueString()

				ratelimitOnRequest = append(ratelimitOnRequest, shared.RatelimitOnRequest{
					Kind:   kind,
					Limits: limits,
					Name:   name3,
				})
			}
			var requestRate1 *shared.RequestRate
			if fromItem.Default.HTTP.RequestRate != nil {
				var interval1 string
				interval1 = fromItem.Default.HTTP.RequestRate.Interval.ValueString()

				var num1 int
				num1 = int(fromItem.Default.HTTP.RequestRate.Num.ValueInt64())

				requestRate1 = &shared.RequestRate{
					Interval: interval1,
					Num:      num1,
				}
			}
			http := shared.MeshGlobalRateLimitItemHTTP{
				Disabled:           disabled,
				OnRateLimit:        onRateLimit,
				RatelimitOnRequest: ratelimitOnRequest,
				RequestRate:        requestRate1,
			}
			mode := new(shared.MeshGlobalRateLimitItemMode)
			if !fromItem.Default.Mode.IsUnknown() && !fromItem.Default.Mode.IsNull() {
				*mode = shared.MeshGlobalRateLimitItemMode(fromItem.Default.Mode.ValueString())
			} else {
				mode = nil
			}
			defaultVar = &shared.MeshGlobalRateLimitItemDefault{
				Backend: backend,
				HTTP:    http,
				Mode:    mode,
			}
		}
		kind1 := shared.MeshGlobalRateLimitItemSpecKind(fromItem.TargetRef.Kind.ValueString())
		labels1 := make(map[string]string)
		for labelsKey1, labelsValue1 := range fromItem.TargetRef.Labels {
			var labelsInst1 string
			labelsInst1 = labelsValue1.ValueString()

			labels1[labelsKey1] = labelsInst1
		}
		mesh1 := new(string)
		if !fromItem.TargetRef.Mesh.IsUnknown() && !fromItem.TargetRef.Mesh.IsNull() {
			*mesh1 = fromItem.TargetRef.Mesh.ValueString()
		} else {
			mesh1 = nil
		}
		name4 := new(string)
		if !fromItem.TargetRef.Name.IsUnknown() && !fromItem.TargetRef.Name.IsNull() {
			*name4 = fromItem.TargetRef.Name.ValueString()
		} else {
			name4 = nil
		}
		namespace := new(string)
		if !fromItem.TargetRef.Namespace.IsUnknown() && !fromItem.TargetRef.Namespace.IsNull() {
			*namespace = fromItem.TargetRef.Namespace.ValueString()
		} else {
			namespace = nil
		}
		var proxyTypes []shared.MeshGlobalRateLimitItemSpecProxyTypes = []shared.MeshGlobalRateLimitItemSpecProxyTypes{}
		for _, proxyTypesItem := range fromItem.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshGlobalRateLimitItemSpecProxyTypes(proxyTypesItem.ValueString()))
		}
		sectionName := new(string)
		if !fromItem.TargetRef.SectionName.IsUnknown() && !fromItem.TargetRef.SectionName.IsNull() {
			*sectionName = fromItem.TargetRef.SectionName.ValueString()
		} else {
			sectionName = nil
		}
		tags := make(map[string]string)
		for tagsKey, tagsValue := range fromItem.TargetRef.Tags {
			var tagsInst string
			tagsInst = tagsValue.ValueString()

			tags[tagsKey] = tagsInst
		}
		targetRef := shared.MeshGlobalRateLimitItemSpecTargetRef{
			Kind:        kind1,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name4,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags,
		}
		from = append(from, shared.MeshGlobalRateLimitItemFrom{
			Default:   defaultVar,
			TargetRef: targetRef,
		})
	}
	var targetRef1 *shared.MeshGlobalRateLimitItemTargetRef
	if r.Spec.TargetRef != nil {
		kind2 := shared.MeshGlobalRateLimitItemKind(r.Spec.TargetRef.Kind.ValueString())
		labels2 := make(map[string]string)
		for labelsKey2, labelsValue2 := range r.Spec.TargetRef.Labels {
			var labelsInst2 string
			labelsInst2 = labelsValue2.ValueString()

			labels2[labelsKey2] = labelsInst2
		}
		mesh2 := new(string)
		if !r.Spec.TargetRef.Mesh.IsUnknown() && !r.Spec.TargetRef.Mesh.IsNull() {
			*mesh2 = r.Spec.TargetRef.Mesh.ValueString()
		} else {
			mesh2 = nil
		}
		name5 := new(string)
		if !r.Spec.TargetRef.Name.IsUnknown() && !r.Spec.TargetRef.Name.IsNull() {
			*name5 = r.Spec.TargetRef.Name.ValueString()
		} else {
			name5 = nil
		}
		namespace1 := new(string)
		if !r.Spec.TargetRef.Namespace.IsUnknown() && !r.Spec.TargetRef.Namespace.IsNull() {
			*namespace1 = r.Spec.TargetRef.Namespace.ValueString()
		} else {
			namespace1 = nil
		}
		var proxyTypes1 []shared.MeshGlobalRateLimitItemProxyTypes = []shared.MeshGlobalRateLimitItemProxyTypes{}
		for _, proxyTypesItem1 := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes1 = append(proxyTypes1, shared.MeshGlobalRateLimitItemProxyTypes(proxyTypesItem1.ValueString()))
		}
		sectionName1 := new(string)
		if !r.Spec.TargetRef.SectionName.IsUnknown() && !r.Spec.TargetRef.SectionName.IsNull() {
			*sectionName1 = r.Spec.TargetRef.SectionName.ValueString()
		} else {
			sectionName1 = nil
		}
		tags1 := make(map[string]string)
		for tagsKey1, tagsValue1 := range r.Spec.TargetRef.Tags {
			var tagsInst1 string
			tagsInst1 = tagsValue1.ValueString()

			tags1[tagsKey1] = tagsInst1
		}
		targetRef1 = &shared.MeshGlobalRateLimitItemTargetRef{
			Kind:        kind2,
			Labels:      labels2,
			Mesh:        mesh2,
			Name:        name5,
			Namespace:   namespace1,
			ProxyTypes:  proxyTypes1,
			SectionName: sectionName1,
			Tags:        tags1,
		}
	}
	var to []shared.MeshGlobalRateLimitItemTo = []shared.MeshGlobalRateLimitItemTo{}
	for _, toItem := range r.Spec.To {
		var default1 *shared.MeshGlobalRateLimitItemSpecDefault
		if toItem.Default != nil {
			limitOnServiceFail1 := new(bool)
			if !toItem.Default.Backend.RateLimitService.LimitOnServiceFail.IsUnknown() && !toItem.Default.Backend.RateLimitService.LimitOnServiceFail.IsNull() {
				*limitOnServiceFail1 = toItem.Default.Backend.RateLimitService.LimitOnServiceFail.ValueBool()
			} else {
				limitOnServiceFail1 = nil
			}
			timeout1 := new(string)
			if !toItem.Default.Backend.RateLimitService.Timeout.IsUnknown() && !toItem.Default.Backend.RateLimitService.Timeout.IsNull() {
				*timeout1 = toItem.Default.Backend.RateLimitService.Timeout.ValueString()
			} else {
				timeout1 = nil
			}
			url1 := new(string)
			if !toItem.Default.Backend.RateLimitService.URL.IsUnknown() && !toItem.Default.Backend.RateLimitService.URL.IsNull() {
				*url1 = toItem.Default.Backend.RateLimitService.URL.ValueString()
			} else {
				url1 = nil
			}
			rateLimitService1 := shared.MeshGlobalRateLimitItemRateLimitService{
				LimitOnServiceFail: limitOnServiceFail1,
				Timeout:            timeout1,
				URL:                url1,
			}
			backend1 := shared.MeshGlobalRateLimitItemBackend{
				RateLimitService: rateLimitService1,
			}
			disabled1 := new(bool)
			if !toItem.Default.HTTP.Disabled.IsUnknown() && !toItem.Default.HTTP.Disabled.IsNull() {
				*disabled1 = toItem.Default.HTTP.Disabled.ValueBool()
			} else {
				disabled1 = nil
			}
			var onRateLimit1 *shared.MeshGlobalRateLimitItemOnRateLimit
			if toItem.Default.HTTP.OnRateLimit != nil {
				var headers1 *shared.MeshGlobalRateLimitItemHeaders
				if toItem.Default.HTTP.OnRateLimit.Headers != nil {
					var add1 []shared.MeshGlobalRateLimitItemAdd = []shared.MeshGlobalRateLimitItemAdd{}
					for _, addItem1 := range toItem.Default.HTTP.OnRateLimit.Headers.Add {
						var name6 string
						name6 = addItem1.Name.ValueString()

						var value3 string
						value3 = addItem1.Value.ValueString()

						add1 = append(add1, shared.MeshGlobalRateLimitItemAdd{
							Name:  name6,
							Value: value3,
						})
					}
					var set1 []shared.MeshGlobalRateLimitItemSet = []shared.MeshGlobalRateLimitItemSet{}
					for _, setItem1 := range toItem.Default.HTTP.OnRateLimit.Headers.Set {
						var name7 string
						name7 = setItem1.Name.ValueString()

						var value4 string
						value4 = setItem1.Value.ValueString()

						set1 = append(set1, shared.MeshGlobalRateLimitItemSet{
							Name:  name7,
							Value: value4,
						})
					}
					headers1 = &shared.MeshGlobalRateLimitItemHeaders{
						Add: add1,
						Set: set1,
					}
				}
				status1 := new(int)
				if !toItem.Default.HTTP.OnRateLimit.Status.IsUnknown() && !toItem.Default.HTTP.OnRateLimit.Status.IsNull() {
					*status1 = int(toItem.Default.HTTP.OnRateLimit.Status.ValueInt64())
				} else {
					status1 = nil
				}
				onRateLimit1 = &shared.MeshGlobalRateLimitItemOnRateLimit{
					Headers: headers1,
					Status:  status1,
				}
			}
			var ratelimitOnRequest1 []shared.MeshGlobalRateLimitItemRatelimitOnRequest = []shared.MeshGlobalRateLimitItemRatelimitOnRequest{}
			for _, ratelimitOnRequestItem1 := range toItem.Default.HTTP.RatelimitOnRequest {
				kind3 := shared.MeshGlobalRateLimitItemSpecToDefaultKind(ratelimitOnRequestItem1.Kind.ValueString())
				var limits1 []shared.MeshGlobalRateLimitItemLimits = []shared.MeshGlobalRateLimitItemLimits{}
				for _, limitsItem1 := range ratelimitOnRequestItem1.Limits {
					var requestRate2 *shared.MeshGlobalRateLimitItemSpecRequestRate
					if limitsItem1.RequestRate != nil {
						var interval2 string
						interval2 = limitsItem1.RequestRate.Interval.ValueString()

						var num2 int
						num2 = int(limitsItem1.RequestRate.Num.ValueInt64())

						requestRate2 = &shared.MeshGlobalRateLimitItemSpecRequestRate{
							Interval: interval2,
							Num:      num2,
						}
					}
					var value5 string
					value5 = limitsItem1.Value.ValueString()

					limits1 = append(limits1, shared.MeshGlobalRateLimitItemLimits{
						RequestRate: requestRate2,
						Value:       value5,
					})
				}
				var name8 string
				name8 = ratelimitOnRequestItem1.Name.ValueString()

				ratelimitOnRequest1 = append(ratelimitOnRequest1, shared.MeshGlobalRateLimitItemRatelimitOnRequest{
					Kind:   kind3,
					Limits: limits1,
					Name:   name8,
				})
			}
			var requestRate3 *shared.MeshGlobalRateLimitItemRequestRate
			if toItem.Default.HTTP.RequestRate != nil {
				var interval3 string
				interval3 = toItem.Default.HTTP.RequestRate.Interval.ValueString()

				var num3 int
				num3 = int(toItem.Default.HTTP.RequestRate.Num.ValueInt64())

				requestRate3 = &shared.MeshGlobalRateLimitItemRequestRate{
					Interval: interval3,
					Num:      num3,
				}
			}
			http1 := shared.MeshGlobalRateLimitItemSpecHTTP{
				Disabled:           disabled1,
				OnRateLimit:        onRateLimit1,
				RatelimitOnRequest: ratelimitOnRequest1,
				RequestRate:        requestRate3,
			}
			mode1 := new(shared.MeshGlobalRateLimitItemSpecMode)
			if !toItem.Default.Mode.IsUnknown() && !toItem.Default.Mode.IsNull() {
				*mode1 = shared.MeshGlobalRateLimitItemSpecMode(toItem.Default.Mode.ValueString())
			} else {
				mode1 = nil
			}
			default1 = &shared.MeshGlobalRateLimitItemSpecDefault{
				Backend: backend1,
				HTTP:    http1,
				Mode:    mode1,
			}
		}
		kind4 := shared.MeshGlobalRateLimitItemSpecToKind(toItem.TargetRef.Kind.ValueString())
		labels3 := make(map[string]string)
		for labelsKey3, labelsValue3 := range toItem.TargetRef.Labels {
			var labelsInst3 string
			labelsInst3 = labelsValue3.ValueString()

			labels3[labelsKey3] = labelsInst3
		}
		mesh3 := new(string)
		if !toItem.TargetRef.Mesh.IsUnknown() && !toItem.TargetRef.Mesh.IsNull() {
			*mesh3 = toItem.TargetRef.Mesh.ValueString()
		} else {
			mesh3 = nil
		}
		name9 := new(string)
		if !toItem.TargetRef.Name.IsUnknown() && !toItem.TargetRef.Name.IsNull() {
			*name9 = toItem.TargetRef.Name.ValueString()
		} else {
			name9 = nil
		}
		namespace2 := new(string)
		if !toItem.TargetRef.Namespace.IsUnknown() && !toItem.TargetRef.Namespace.IsNull() {
			*namespace2 = toItem.TargetRef.Namespace.ValueString()
		} else {
			namespace2 = nil
		}
		var proxyTypes2 []shared.MeshGlobalRateLimitItemSpecToProxyTypes = []shared.MeshGlobalRateLimitItemSpecToProxyTypes{}
		for _, proxyTypesItem2 := range toItem.TargetRef.ProxyTypes {
			proxyTypes2 = append(proxyTypes2, shared.MeshGlobalRateLimitItemSpecToProxyTypes(proxyTypesItem2.ValueString()))
		}
		sectionName2 := new(string)
		if !toItem.TargetRef.SectionName.IsUnknown() && !toItem.TargetRef.SectionName.IsNull() {
			*sectionName2 = toItem.TargetRef.SectionName.ValueString()
		} else {
			sectionName2 = nil
		}
		tags2 := make(map[string]string)
		for tagsKey2, tagsValue2 := range toItem.TargetRef.Tags {
			var tagsInst2 string
			tagsInst2 = tagsValue2.ValueString()

			tags2[tagsKey2] = tagsInst2
		}
		targetRef2 := shared.MeshGlobalRateLimitItemSpecToTargetRef{
			Kind:        kind4,
			Labels:      labels3,
			Mesh:        mesh3,
			Name:        name9,
			Namespace:   namespace2,
			ProxyTypes:  proxyTypes2,
			SectionName: sectionName2,
			Tags:        tags2,
		}
		to = append(to, shared.MeshGlobalRateLimitItemTo{
			Default:   default1,
			TargetRef: targetRef2,
		})
	}
	spec := shared.MeshGlobalRateLimitItemSpec{
		From:      from,
		TargetRef: targetRef1,
		To:        to,
	}
	out := shared.MeshGlobalRateLimitItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshGlobalRateLimitResourceModel) RefreshFromSharedMeshGlobalRateLimitCreateOrUpdateSuccessResponse(resp *shared.MeshGlobalRateLimitCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshGlobalRateLimitResourceModel) RefreshFromSharedMeshGlobalRateLimitItem(resp *shared.MeshGlobalRateLimitItem) {
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
		r.Spec.From = []tfTypes.MeshGlobalRateLimitItemFrom{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from1 tfTypes.MeshGlobalRateLimitItemFrom
			if fromItem.Default == nil {
				from1.Default = nil
			} else {
				from1.Default = &tfTypes.MeshGlobalRateLimitItemDefault{}
				from1.Default.Backend.RateLimitService.LimitOnServiceFail = types.BoolPointerValue(fromItem.Default.Backend.RateLimitService.LimitOnServiceFail)
				from1.Default.Backend.RateLimitService.Timeout = types.StringPointerValue(fromItem.Default.Backend.RateLimitService.Timeout)
				from1.Default.Backend.RateLimitService.URL = types.StringPointerValue(fromItem.Default.Backend.RateLimitService.URL)
				from1.Default.HTTP.Disabled = types.BoolPointerValue(fromItem.Default.HTTP.Disabled)
				if fromItem.Default.HTTP.OnRateLimit == nil {
					from1.Default.HTTP.OnRateLimit = nil
				} else {
					from1.Default.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
					if fromItem.Default.HTTP.OnRateLimit.Headers == nil {
						from1.Default.HTTP.OnRateLimit.Headers = nil
					} else {
						from1.Default.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemSpecHeaders{}
						from1.Default.HTTP.OnRateLimit.Headers.Add = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for addCount, addItem := range fromItem.Default.HTTP.OnRateLimit.Headers.Add {
							var add1 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							add1.Name = types.StringValue(addItem.Name)
							add1.Value = types.StringValue(addItem.Value)
							if addCount+1 > len(from1.Default.HTTP.OnRateLimit.Headers.Add) {
								from1.Default.HTTP.OnRateLimit.Headers.Add = append(from1.Default.HTTP.OnRateLimit.Headers.Add, add1)
							} else {
								from1.Default.HTTP.OnRateLimit.Headers.Add[addCount].Name = add1.Name
								from1.Default.HTTP.OnRateLimit.Headers.Add[addCount].Value = add1.Value
							}
						}
						from1.Default.HTTP.OnRateLimit.Headers.Set = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for setCount, setItem := range fromItem.Default.HTTP.OnRateLimit.Headers.Set {
							var set1 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							set1.Name = types.StringValue(setItem.Name)
							set1.Value = types.StringValue(setItem.Value)
							if setCount+1 > len(from1.Default.HTTP.OnRateLimit.Headers.Set) {
								from1.Default.HTTP.OnRateLimit.Headers.Set = append(from1.Default.HTTP.OnRateLimit.Headers.Set, set1)
							} else {
								from1.Default.HTTP.OnRateLimit.Headers.Set[setCount].Name = set1.Name
								from1.Default.HTTP.OnRateLimit.Headers.Set[setCount].Value = set1.Value
							}
						}
					}
					if fromItem.Default.HTTP.OnRateLimit.Status != nil {
						from1.Default.HTTP.OnRateLimit.Status = types.Int64Value(int64(*fromItem.Default.HTTP.OnRateLimit.Status))
					} else {
						from1.Default.HTTP.OnRateLimit.Status = types.Int64Null()
					}
				}
				from1.Default.HTTP.RatelimitOnRequest = []tfTypes.RatelimitOnRequest{}
				for ratelimitOnRequestCount, ratelimitOnRequestItem := range fromItem.Default.HTTP.RatelimitOnRequest {
					var ratelimitOnRequest1 tfTypes.RatelimitOnRequest
					ratelimitOnRequest1.Kind = types.StringValue(string(ratelimitOnRequestItem.Kind))
					ratelimitOnRequest1.Limits = []tfTypes.Limits{}
					for limitsCount, limitsItem := range ratelimitOnRequestItem.Limits {
						var limits1 tfTypes.Limits
						if limitsItem.RequestRate == nil {
							limits1.RequestRate = nil
						} else {
							limits1.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							limits1.RequestRate.Interval = types.StringValue(limitsItem.RequestRate.Interval)
							limits1.RequestRate.Num = types.Int64Value(int64(limitsItem.RequestRate.Num))
						}
						limits1.Value = types.StringValue(limitsItem.Value)
						if limitsCount+1 > len(ratelimitOnRequest1.Limits) {
							ratelimitOnRequest1.Limits = append(ratelimitOnRequest1.Limits, limits1)
						} else {
							ratelimitOnRequest1.Limits[limitsCount].RequestRate = limits1.RequestRate
							ratelimitOnRequest1.Limits[limitsCount].Value = limits1.Value
						}
					}
					ratelimitOnRequest1.Name = types.StringValue(ratelimitOnRequestItem.Name)
					if ratelimitOnRequestCount+1 > len(from1.Default.HTTP.RatelimitOnRequest) {
						from1.Default.HTTP.RatelimitOnRequest = append(from1.Default.HTTP.RatelimitOnRequest, ratelimitOnRequest1)
					} else {
						from1.Default.HTTP.RatelimitOnRequest[ratelimitOnRequestCount].Kind = ratelimitOnRequest1.Kind
						from1.Default.HTTP.RatelimitOnRequest[ratelimitOnRequestCount].Limits = ratelimitOnRequest1.Limits
						from1.Default.HTTP.RatelimitOnRequest[ratelimitOnRequestCount].Name = ratelimitOnRequest1.Name
					}
				}
				if fromItem.Default.HTTP.RequestRate == nil {
					from1.Default.HTTP.RequestRate = nil
				} else {
					from1.Default.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
					from1.Default.HTTP.RequestRate.Interval = types.StringValue(fromItem.Default.HTTP.RequestRate.Interval)
					from1.Default.HTTP.RequestRate.Num = types.Int64Value(int64(fromItem.Default.HTTP.RequestRate.Num))
				}
				if fromItem.Default.Mode != nil {
					from1.Default.Mode = types.StringValue(string(*fromItem.Default.Mode))
				} else {
					from1.Default.Mode = types.StringNull()
				}
			}
			from1.TargetRef.Kind = types.StringValue(string(fromItem.TargetRef.Kind))
			if len(fromItem.TargetRef.Labels) > 0 {
				from1.TargetRef.Labels = make(map[string]types.String)
				for key1, value4 := range fromItem.TargetRef.Labels {
					from1.TargetRef.Labels[key1] = types.StringValue(value4)
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
				for key2, value5 := range fromItem.TargetRef.Tags {
					from1.TargetRef.Tags[key2] = types.StringValue(value5)
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
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String)
				for key3, value6 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key3] = types.StringValue(value6)
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
				for key4, value7 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key4] = types.StringValue(value7)
				}
			}
		}
		r.Spec.To = []tfTypes.MeshGlobalRateLimitItemFrom{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshGlobalRateLimitItemFrom
			if toItem.Default == nil {
				to1.Default = nil
			} else {
				to1.Default = &tfTypes.MeshGlobalRateLimitItemDefault{}
				to1.Default.Backend.RateLimitService.LimitOnServiceFail = types.BoolPointerValue(toItem.Default.Backend.RateLimitService.LimitOnServiceFail)
				to1.Default.Backend.RateLimitService.Timeout = types.StringPointerValue(toItem.Default.Backend.RateLimitService.Timeout)
				to1.Default.Backend.RateLimitService.URL = types.StringPointerValue(toItem.Default.Backend.RateLimitService.URL)
				to1.Default.HTTP.Disabled = types.BoolPointerValue(toItem.Default.HTTP.Disabled)
				if toItem.Default.HTTP.OnRateLimit == nil {
					to1.Default.HTTP.OnRateLimit = nil
				} else {
					to1.Default.HTTP.OnRateLimit = &tfTypes.OnRateLimit{}
					if toItem.Default.HTTP.OnRateLimit.Headers == nil {
						to1.Default.HTTP.OnRateLimit.Headers = nil
					} else {
						to1.Default.HTTP.OnRateLimit.Headers = &tfTypes.MeshGlobalRateLimitItemSpecHeaders{}
						to1.Default.HTTP.OnRateLimit.Headers.Add = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for addCount1, addItem1 := range toItem.Default.HTTP.OnRateLimit.Headers.Add {
							var add3 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							add3.Name = types.StringValue(addItem1.Name)
							add3.Value = types.StringValue(addItem1.Value)
							if addCount1+1 > len(to1.Default.HTTP.OnRateLimit.Headers.Add) {
								to1.Default.HTTP.OnRateLimit.Headers.Add = append(to1.Default.HTTP.OnRateLimit.Headers.Add, add3)
							} else {
								to1.Default.HTTP.OnRateLimit.Headers.Add[addCount1].Name = add3.Name
								to1.Default.HTTP.OnRateLimit.Headers.Add[addCount1].Value = add3.Value
							}
						}
						to1.Default.HTTP.OnRateLimit.Headers.Set = []tfTypes.ConfigurationDataPlaneGroupEnvironmentField{}
						for setCount1, setItem1 := range toItem.Default.HTTP.OnRateLimit.Headers.Set {
							var set3 tfTypes.ConfigurationDataPlaneGroupEnvironmentField
							set3.Name = types.StringValue(setItem1.Name)
							set3.Value = types.StringValue(setItem1.Value)
							if setCount1+1 > len(to1.Default.HTTP.OnRateLimit.Headers.Set) {
								to1.Default.HTTP.OnRateLimit.Headers.Set = append(to1.Default.HTTP.OnRateLimit.Headers.Set, set3)
							} else {
								to1.Default.HTTP.OnRateLimit.Headers.Set[setCount1].Name = set3.Name
								to1.Default.HTTP.OnRateLimit.Headers.Set[setCount1].Value = set3.Value
							}
						}
					}
					if toItem.Default.HTTP.OnRateLimit.Status != nil {
						to1.Default.HTTP.OnRateLimit.Status = types.Int64Value(int64(*toItem.Default.HTTP.OnRateLimit.Status))
					} else {
						to1.Default.HTTP.OnRateLimit.Status = types.Int64Null()
					}
				}
				to1.Default.HTTP.RatelimitOnRequest = []tfTypes.RatelimitOnRequest{}
				for ratelimitOnRequestCount1, ratelimitOnRequestItem1 := range toItem.Default.HTTP.RatelimitOnRequest {
					var ratelimitOnRequest3 tfTypes.RatelimitOnRequest
					ratelimitOnRequest3.Kind = types.StringValue(string(ratelimitOnRequestItem1.Kind))
					ratelimitOnRequest3.Limits = []tfTypes.Limits{}
					for limitsCount1, limitsItem1 := range ratelimitOnRequestItem1.Limits {
						var limits3 tfTypes.Limits
						if limitsItem1.RequestRate == nil {
							limits3.RequestRate = nil
						} else {
							limits3.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
							limits3.RequestRate.Interval = types.StringValue(limitsItem1.RequestRate.Interval)
							limits3.RequestRate.Num = types.Int64Value(int64(limitsItem1.RequestRate.Num))
						}
						limits3.Value = types.StringValue(limitsItem1.Value)
						if limitsCount1+1 > len(ratelimitOnRequest3.Limits) {
							ratelimitOnRequest3.Limits = append(ratelimitOnRequest3.Limits, limits3)
						} else {
							ratelimitOnRequest3.Limits[limitsCount1].RequestRate = limits3.RequestRate
							ratelimitOnRequest3.Limits[limitsCount1].Value = limits3.Value
						}
					}
					ratelimitOnRequest3.Name = types.StringValue(ratelimitOnRequestItem1.Name)
					if ratelimitOnRequestCount1+1 > len(to1.Default.HTTP.RatelimitOnRequest) {
						to1.Default.HTTP.RatelimitOnRequest = append(to1.Default.HTTP.RatelimitOnRequest, ratelimitOnRequest3)
					} else {
						to1.Default.HTTP.RatelimitOnRequest[ratelimitOnRequestCount1].Kind = ratelimitOnRequest3.Kind
						to1.Default.HTTP.RatelimitOnRequest[ratelimitOnRequestCount1].Limits = ratelimitOnRequest3.Limits
						to1.Default.HTTP.RatelimitOnRequest[ratelimitOnRequestCount1].Name = ratelimitOnRequest3.Name
					}
				}
				if toItem.Default.HTTP.RequestRate == nil {
					to1.Default.HTTP.RequestRate = nil
				} else {
					to1.Default.HTTP.RequestRate = &tfTypes.MeshGlobalRateLimitItemSpecFromRequestRate{}
					to1.Default.HTTP.RequestRate.Interval = types.StringValue(toItem.Default.HTTP.RequestRate.Interval)
					to1.Default.HTTP.RequestRate.Num = types.Int64Value(int64(toItem.Default.HTTP.RequestRate.Num))
				}
				if toItem.Default.Mode != nil {
					to1.Default.Mode = types.StringValue(string(*toItem.Default.Mode))
				} else {
					to1.Default.Mode = types.StringNull()
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
