// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshFaultInjectionResourceModel) ToSharedMeshFaultInjectionItem() *shared.MeshFaultInjectionItem {
	typeVar := shared.MeshFaultInjectionItemType(r.Type.ValueString())
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
	var from []shared.MeshFaultInjectionItemFrom = []shared.MeshFaultInjectionItemFrom{}
	for _, fromItem := range r.Spec.From {
		var defaultVar *shared.MeshFaultInjectionItemDefault
		if fromItem.Default != nil {
			var http []shared.HTTP = []shared.HTTP{}
			for _, httpItem := range fromItem.Default.HTTP {
				var abort *shared.Abort
				if httpItem.Abort != nil {
					var httpStatus int
					httpStatus = int(httpItem.Abort.HTTPStatus.ValueInt64())

					var percentage shared.Percentage
					integer := new(int64)
					if !httpItem.Abort.Percentage.Integer.IsUnknown() && !httpItem.Abort.Percentage.Integer.IsNull() {
						*integer = httpItem.Abort.Percentage.Integer.ValueInt64()
					} else {
						integer = nil
					}
					if integer != nil {
						percentage = shared.Percentage{
							Integer: integer,
						}
					}
					str := new(string)
					if !httpItem.Abort.Percentage.Str.IsUnknown() && !httpItem.Abort.Percentage.Str.IsNull() {
						*str = httpItem.Abort.Percentage.Str.ValueString()
					} else {
						str = nil
					}
					if str != nil {
						percentage = shared.Percentage{
							Str: str,
						}
					}
					abort = &shared.Abort{
						HTTPStatus: httpStatus,
						Percentage: percentage,
					}
				}
				var delay *shared.Delay
				if httpItem.Delay != nil {
					var percentage1 shared.MeshFaultInjectionItemPercentage
					integer1 := new(int64)
					if !httpItem.Delay.Percentage.Integer.IsUnknown() && !httpItem.Delay.Percentage.Integer.IsNull() {
						*integer1 = httpItem.Delay.Percentage.Integer.ValueInt64()
					} else {
						integer1 = nil
					}
					if integer1 != nil {
						percentage1 = shared.MeshFaultInjectionItemPercentage{
							Integer: integer1,
						}
					}
					str1 := new(string)
					if !httpItem.Delay.Percentage.Str.IsUnknown() && !httpItem.Delay.Percentage.Str.IsNull() {
						*str1 = httpItem.Delay.Percentage.Str.ValueString()
					} else {
						str1 = nil
					}
					if str1 != nil {
						percentage1 = shared.MeshFaultInjectionItemPercentage{
							Str: str1,
						}
					}
					var value string
					value = httpItem.Delay.Value.ValueString()

					delay = &shared.Delay{
						Percentage: percentage1,
						Value:      value,
					}
				}
				var responseBandwidth *shared.ResponseBandwidth
				if httpItem.ResponseBandwidth != nil {
					var limit string
					limit = httpItem.ResponseBandwidth.Limit.ValueString()

					var percentage2 shared.MeshFaultInjectionItemSpecPercentage
					integer2 := new(int64)
					if !httpItem.ResponseBandwidth.Percentage.Integer.IsUnknown() && !httpItem.ResponseBandwidth.Percentage.Integer.IsNull() {
						*integer2 = httpItem.ResponseBandwidth.Percentage.Integer.ValueInt64()
					} else {
						integer2 = nil
					}
					if integer2 != nil {
						percentage2 = shared.MeshFaultInjectionItemSpecPercentage{
							Integer: integer2,
						}
					}
					str2 := new(string)
					if !httpItem.ResponseBandwidth.Percentage.Str.IsUnknown() && !httpItem.ResponseBandwidth.Percentage.Str.IsNull() {
						*str2 = httpItem.ResponseBandwidth.Percentage.Str.ValueString()
					} else {
						str2 = nil
					}
					if str2 != nil {
						percentage2 = shared.MeshFaultInjectionItemSpecPercentage{
							Str: str2,
						}
					}
					responseBandwidth = &shared.ResponseBandwidth{
						Limit:      limit,
						Percentage: percentage2,
					}
				}
				http = append(http, shared.HTTP{
					Abort:             abort,
					Delay:             delay,
					ResponseBandwidth: responseBandwidth,
				})
			}
			defaultVar = &shared.MeshFaultInjectionItemDefault{
				HTTP: http,
			}
		}
		kind := new(shared.MeshFaultInjectionItemSpecKind)
		if !fromItem.TargetRef.Kind.IsUnknown() && !fromItem.TargetRef.Kind.IsNull() {
			*kind = shared.MeshFaultInjectionItemSpecKind(fromItem.TargetRef.Kind.ValueString())
		} else {
			kind = nil
		}
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
		name1 := new(string)
		if !fromItem.TargetRef.Name.IsUnknown() && !fromItem.TargetRef.Name.IsNull() {
			*name1 = fromItem.TargetRef.Name.ValueString()
		} else {
			name1 = nil
		}
		namespace := new(string)
		if !fromItem.TargetRef.Namespace.IsUnknown() && !fromItem.TargetRef.Namespace.IsNull() {
			*namespace = fromItem.TargetRef.Namespace.ValueString()
		} else {
			namespace = nil
		}
		var proxyTypes []shared.MeshFaultInjectionItemSpecProxyTypes = []shared.MeshFaultInjectionItemSpecProxyTypes{}
		for _, proxyTypesItem := range fromItem.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshFaultInjectionItemSpecProxyTypes(proxyTypesItem.ValueString()))
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
		targetRef := shared.MeshFaultInjectionItemSpecTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name1,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags,
		}
		from = append(from, shared.MeshFaultInjectionItemFrom{
			Default:   defaultVar,
			TargetRef: targetRef,
		})
	}
	var targetRef1 *shared.MeshFaultInjectionItemTargetRef
	if r.Spec.TargetRef != nil {
		kind1 := new(shared.MeshFaultInjectionItemKind)
		if !r.Spec.TargetRef.Kind.IsUnknown() && !r.Spec.TargetRef.Kind.IsNull() {
			*kind1 = shared.MeshFaultInjectionItemKind(r.Spec.TargetRef.Kind.ValueString())
		} else {
			kind1 = nil
		}
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
		name2 := new(string)
		if !r.Spec.TargetRef.Name.IsUnknown() && !r.Spec.TargetRef.Name.IsNull() {
			*name2 = r.Spec.TargetRef.Name.ValueString()
		} else {
			name2 = nil
		}
		namespace1 := new(string)
		if !r.Spec.TargetRef.Namespace.IsUnknown() && !r.Spec.TargetRef.Namespace.IsNull() {
			*namespace1 = r.Spec.TargetRef.Namespace.ValueString()
		} else {
			namespace1 = nil
		}
		var proxyTypes1 []shared.MeshFaultInjectionItemProxyTypes = []shared.MeshFaultInjectionItemProxyTypes{}
		for _, proxyTypesItem1 := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes1 = append(proxyTypes1, shared.MeshFaultInjectionItemProxyTypes(proxyTypesItem1.ValueString()))
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
		targetRef1 = &shared.MeshFaultInjectionItemTargetRef{
			Kind:        kind1,
			Labels:      labels2,
			Mesh:        mesh2,
			Name:        name2,
			Namespace:   namespace1,
			ProxyTypes:  proxyTypes1,
			SectionName: sectionName1,
			Tags:        tags1,
		}
	}
	var to []shared.MeshFaultInjectionItemTo = []shared.MeshFaultInjectionItemTo{}
	for _, toItem := range r.Spec.To {
		var default1 *shared.MeshFaultInjectionItemSpecDefault
		if toItem.Default != nil {
			var http1 []shared.MeshFaultInjectionItemHTTP = []shared.MeshFaultInjectionItemHTTP{}
			for _, httpItem1 := range toItem.Default.HTTP {
				var abort1 *shared.MeshFaultInjectionItemAbort
				if httpItem1.Abort != nil {
					var httpStatus1 int
					httpStatus1 = int(httpItem1.Abort.HTTPStatus.ValueInt64())

					var percentage3 shared.MeshFaultInjectionItemSpecToPercentage
					integer3 := new(int64)
					if !httpItem1.Abort.Percentage.Integer.IsUnknown() && !httpItem1.Abort.Percentage.Integer.IsNull() {
						*integer3 = httpItem1.Abort.Percentage.Integer.ValueInt64()
					} else {
						integer3 = nil
					}
					if integer3 != nil {
						percentage3 = shared.MeshFaultInjectionItemSpecToPercentage{
							Integer: integer3,
						}
					}
					str3 := new(string)
					if !httpItem1.Abort.Percentage.Str.IsUnknown() && !httpItem1.Abort.Percentage.Str.IsNull() {
						*str3 = httpItem1.Abort.Percentage.Str.ValueString()
					} else {
						str3 = nil
					}
					if str3 != nil {
						percentage3 = shared.MeshFaultInjectionItemSpecToPercentage{
							Str: str3,
						}
					}
					abort1 = &shared.MeshFaultInjectionItemAbort{
						HTTPStatus: httpStatus1,
						Percentage: percentage3,
					}
				}
				var delay1 *shared.MeshFaultInjectionItemDelay
				if httpItem1.Delay != nil {
					var percentage4 shared.MeshFaultInjectionItemSpecToDefaultPercentage
					integer4 := new(int64)
					if !httpItem1.Delay.Percentage.Integer.IsUnknown() && !httpItem1.Delay.Percentage.Integer.IsNull() {
						*integer4 = httpItem1.Delay.Percentage.Integer.ValueInt64()
					} else {
						integer4 = nil
					}
					if integer4 != nil {
						percentage4 = shared.MeshFaultInjectionItemSpecToDefaultPercentage{
							Integer: integer4,
						}
					}
					str4 := new(string)
					if !httpItem1.Delay.Percentage.Str.IsUnknown() && !httpItem1.Delay.Percentage.Str.IsNull() {
						*str4 = httpItem1.Delay.Percentage.Str.ValueString()
					} else {
						str4 = nil
					}
					if str4 != nil {
						percentage4 = shared.MeshFaultInjectionItemSpecToDefaultPercentage{
							Str: str4,
						}
					}
					var value1 string
					value1 = httpItem1.Delay.Value.ValueString()

					delay1 = &shared.MeshFaultInjectionItemDelay{
						Percentage: percentage4,
						Value:      value1,
					}
				}
				var responseBandwidth1 *shared.MeshFaultInjectionItemResponseBandwidth
				if httpItem1.ResponseBandwidth != nil {
					var limit1 string
					limit1 = httpItem1.ResponseBandwidth.Limit.ValueString()

					var percentage5 shared.MeshFaultInjectionItemSpecToDefaultHTTPPercentage
					integer5 := new(int64)
					if !httpItem1.ResponseBandwidth.Percentage.Integer.IsUnknown() && !httpItem1.ResponseBandwidth.Percentage.Integer.IsNull() {
						*integer5 = httpItem1.ResponseBandwidth.Percentage.Integer.ValueInt64()
					} else {
						integer5 = nil
					}
					if integer5 != nil {
						percentage5 = shared.MeshFaultInjectionItemSpecToDefaultHTTPPercentage{
							Integer: integer5,
						}
					}
					str5 := new(string)
					if !httpItem1.ResponseBandwidth.Percentage.Str.IsUnknown() && !httpItem1.ResponseBandwidth.Percentage.Str.IsNull() {
						*str5 = httpItem1.ResponseBandwidth.Percentage.Str.ValueString()
					} else {
						str5 = nil
					}
					if str5 != nil {
						percentage5 = shared.MeshFaultInjectionItemSpecToDefaultHTTPPercentage{
							Str: str5,
						}
					}
					responseBandwidth1 = &shared.MeshFaultInjectionItemResponseBandwidth{
						Limit:      limit1,
						Percentage: percentage5,
					}
				}
				http1 = append(http1, shared.MeshFaultInjectionItemHTTP{
					Abort:             abort1,
					Delay:             delay1,
					ResponseBandwidth: responseBandwidth1,
				})
			}
			default1 = &shared.MeshFaultInjectionItemSpecDefault{
				HTTP: http1,
			}
		}
		kind2 := new(shared.MeshFaultInjectionItemSpecToKind)
		if !toItem.TargetRef.Kind.IsUnknown() && !toItem.TargetRef.Kind.IsNull() {
			*kind2 = shared.MeshFaultInjectionItemSpecToKind(toItem.TargetRef.Kind.ValueString())
		} else {
			kind2 = nil
		}
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
		name3 := new(string)
		if !toItem.TargetRef.Name.IsUnknown() && !toItem.TargetRef.Name.IsNull() {
			*name3 = toItem.TargetRef.Name.ValueString()
		} else {
			name3 = nil
		}
		namespace2 := new(string)
		if !toItem.TargetRef.Namespace.IsUnknown() && !toItem.TargetRef.Namespace.IsNull() {
			*namespace2 = toItem.TargetRef.Namespace.ValueString()
		} else {
			namespace2 = nil
		}
		var proxyTypes2 []shared.MeshFaultInjectionItemSpecToProxyTypes = []shared.MeshFaultInjectionItemSpecToProxyTypes{}
		for _, proxyTypesItem2 := range toItem.TargetRef.ProxyTypes {
			proxyTypes2 = append(proxyTypes2, shared.MeshFaultInjectionItemSpecToProxyTypes(proxyTypesItem2.ValueString()))
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
		targetRef2 := shared.MeshFaultInjectionItemSpecToTargetRef{
			Kind:        kind2,
			Labels:      labels3,
			Mesh:        mesh3,
			Name:        name3,
			Namespace:   namespace2,
			ProxyTypes:  proxyTypes2,
			SectionName: sectionName2,
			Tags:        tags2,
		}
		to = append(to, shared.MeshFaultInjectionItemTo{
			Default:   default1,
			TargetRef: targetRef2,
		})
	}
	spec := shared.MeshFaultInjectionItemSpec{
		From:      from,
		TargetRef: targetRef1,
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
	out := shared.MeshFaultInjectionItem{
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

func (r *MeshFaultInjectionResourceModel) RefreshFromSharedMeshFaultInjectionCreateOrUpdateSuccessResponse(resp *shared.MeshFaultInjectionCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = []types.String{}
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshFaultInjectionResourceModel) RefreshFromSharedMeshFaultInjectionItem(resp *shared.MeshFaultInjectionItem) {
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
		r.Spec.From = []tfTypes.MeshFaultInjectionItemFrom{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from1 tfTypes.MeshFaultInjectionItemFrom
			if fromItem.Default == nil {
				from1.Default = nil
			} else {
				from1.Default = &tfTypes.MeshFaultInjectionItemDefault{}
				from1.Default.HTTP = []tfTypes.HTTP{}
				for httpCount, httpItem := range fromItem.Default.HTTP {
					var http1 tfTypes.HTTP
					if httpItem.Abort == nil {
						http1.Abort = nil
					} else {
						http1.Abort = &tfTypes.Abort{}
						http1.Abort.HTTPStatus = types.Int64Value(int64(httpItem.Abort.HTTPStatus))
						if httpItem.Abort.Percentage.Integer != nil {
							http1.Abort.Percentage.Integer = types.Int64PointerValue(httpItem.Abort.Percentage.Integer)
						}
						if httpItem.Abort.Percentage.Str != nil {
							http1.Abort.Percentage.Str = types.StringPointerValue(httpItem.Abort.Percentage.Str)
						}
					}
					if httpItem.Delay == nil {
						http1.Delay = nil
					} else {
						http1.Delay = &tfTypes.Delay{}
						if httpItem.Delay.Percentage.Integer != nil {
							http1.Delay.Percentage.Integer = types.Int64PointerValue(httpItem.Delay.Percentage.Integer)
						}
						if httpItem.Delay.Percentage.Str != nil {
							http1.Delay.Percentage.Str = types.StringPointerValue(httpItem.Delay.Percentage.Str)
						}
						http1.Delay.Value = types.StringValue(httpItem.Delay.Value)
					}
					if httpItem.ResponseBandwidth == nil {
						http1.ResponseBandwidth = nil
					} else {
						http1.ResponseBandwidth = &tfTypes.ResponseBandwidth{}
						http1.ResponseBandwidth.Limit = types.StringValue(httpItem.ResponseBandwidth.Limit)
						if httpItem.ResponseBandwidth.Percentage.Integer != nil {
							http1.ResponseBandwidth.Percentage.Integer = types.Int64PointerValue(httpItem.ResponseBandwidth.Percentage.Integer)
						}
						if httpItem.ResponseBandwidth.Percentage.Str != nil {
							http1.ResponseBandwidth.Percentage.Str = types.StringPointerValue(httpItem.ResponseBandwidth.Percentage.Str)
						}
					}
					if httpCount+1 > len(from1.Default.HTTP) {
						from1.Default.HTTP = append(from1.Default.HTTP, http1)
					} else {
						from1.Default.HTTP[httpCount].Abort = http1.Abort
						from1.Default.HTTP[httpCount].Delay = http1.Delay
						from1.Default.HTTP[httpCount].ResponseBandwidth = http1.ResponseBandwidth
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
				for key1, value2 := range fromItem.TargetRef.Labels {
					from1.TargetRef.Labels[key1] = types.StringValue(value2)
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
				for key2, value3 := range fromItem.TargetRef.Tags {
					from1.TargetRef.Tags[key2] = types.StringValue(value3)
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
				for key3, value4 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key3] = types.StringValue(value4)
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
				for key4, value5 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key4] = types.StringValue(value5)
				}
			}
		}
		r.Spec.To = []tfTypes.MeshFaultInjectionItemFrom{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshFaultInjectionItemFrom
			if toItem.Default == nil {
				to1.Default = nil
			} else {
				to1.Default = &tfTypes.MeshFaultInjectionItemDefault{}
				to1.Default.HTTP = []tfTypes.HTTP{}
				for httpCount1, httpItem1 := range toItem.Default.HTTP {
					var http3 tfTypes.HTTP
					if httpItem1.Abort == nil {
						http3.Abort = nil
					} else {
						http3.Abort = &tfTypes.Abort{}
						http3.Abort.HTTPStatus = types.Int64Value(int64(httpItem1.Abort.HTTPStatus))
						if httpItem1.Abort.Percentage.Integer != nil {
							http3.Abort.Percentage.Integer = types.Int64PointerValue(httpItem1.Abort.Percentage.Integer)
						}
						if httpItem1.Abort.Percentage.Str != nil {
							http3.Abort.Percentage.Str = types.StringPointerValue(httpItem1.Abort.Percentage.Str)
						}
					}
					if httpItem1.Delay == nil {
						http3.Delay = nil
					} else {
						http3.Delay = &tfTypes.Delay{}
						if httpItem1.Delay.Percentage.Integer != nil {
							http3.Delay.Percentage.Integer = types.Int64PointerValue(httpItem1.Delay.Percentage.Integer)
						}
						if httpItem1.Delay.Percentage.Str != nil {
							http3.Delay.Percentage.Str = types.StringPointerValue(httpItem1.Delay.Percentage.Str)
						}
						http3.Delay.Value = types.StringValue(httpItem1.Delay.Value)
					}
					if httpItem1.ResponseBandwidth == nil {
						http3.ResponseBandwidth = nil
					} else {
						http3.ResponseBandwidth = &tfTypes.ResponseBandwidth{}
						http3.ResponseBandwidth.Limit = types.StringValue(httpItem1.ResponseBandwidth.Limit)
						if httpItem1.ResponseBandwidth.Percentage.Integer != nil {
							http3.ResponseBandwidth.Percentage.Integer = types.Int64PointerValue(httpItem1.ResponseBandwidth.Percentage.Integer)
						}
						if httpItem1.ResponseBandwidth.Percentage.Str != nil {
							http3.ResponseBandwidth.Percentage.Str = types.StringPointerValue(httpItem1.ResponseBandwidth.Percentage.Str)
						}
					}
					if httpCount1+1 > len(to1.Default.HTTP) {
						to1.Default.HTTP = append(to1.Default.HTTP, http3)
					} else {
						to1.Default.HTTP[httpCount1].Abort = http3.Abort
						to1.Default.HTTP[httpCount1].Delay = http3.Delay
						to1.Default.HTTP[httpCount1].ResponseBandwidth = http3.ResponseBandwidth
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
				for key5, value7 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key5] = types.StringValue(value7)
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
				for key6, value8 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key6] = types.StringValue(value8)
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