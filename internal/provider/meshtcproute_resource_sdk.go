// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshTCPRouteResourceModel) ToSharedMeshTCPRouteItemInput() *shared.MeshTCPRouteItemInput {
	typeVar := shared.MeshTCPRouteItemType(r.Type.ValueString())
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
	var targetRef *shared.MeshTCPRouteItemTargetRef
	if r.Spec.TargetRef != nil {
		kind := shared.MeshTCPRouteItemKind(r.Spec.TargetRef.Kind.ValueString())
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
		var proxyTypes []shared.MeshTCPRouteItemProxyTypes = []shared.MeshTCPRouteItemProxyTypes{}
		for _, proxyTypesItem := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshTCPRouteItemProxyTypes(proxyTypesItem.ValueString()))
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
		targetRef = &shared.MeshTCPRouteItemTargetRef{
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
	var to []shared.MeshTCPRouteItemTo = []shared.MeshTCPRouteItemTo{}
	for _, toItem := range r.Spec.To {
		var rules []shared.MeshTCPRouteItemRules = []shared.MeshTCPRouteItemRules{}
		for _, rulesItem := range toItem.Rules {
			var backendRefs []shared.MeshTCPRouteItemBackendRefs = []shared.MeshTCPRouteItemBackendRefs{}
			for _, backendRefsItem := range rulesItem.Default.BackendRefs {
				kind1 := shared.MeshTCPRouteItemSpecToKind(backendRefsItem.Kind.ValueString())
				labels2 := make(map[string]string)
				for labelsKey2, labelsValue2 := range backendRefsItem.Labels {
					var labelsInst2 string
					labelsInst2 = labelsValue2.ValueString()

					labels2[labelsKey2] = labelsInst2
				}
				mesh2 := new(string)
				if !backendRefsItem.Mesh.IsUnknown() && !backendRefsItem.Mesh.IsNull() {
					*mesh2 = backendRefsItem.Mesh.ValueString()
				} else {
					mesh2 = nil
				}
				name2 := new(string)
				if !backendRefsItem.Name.IsUnknown() && !backendRefsItem.Name.IsNull() {
					*name2 = backendRefsItem.Name.ValueString()
				} else {
					name2 = nil
				}
				namespace1 := new(string)
				if !backendRefsItem.Namespace.IsUnknown() && !backendRefsItem.Namespace.IsNull() {
					*namespace1 = backendRefsItem.Namespace.ValueString()
				} else {
					namespace1 = nil
				}
				port := new(int)
				if !backendRefsItem.Port.IsUnknown() && !backendRefsItem.Port.IsNull() {
					*port = int(backendRefsItem.Port.ValueInt32())
				} else {
					port = nil
				}
				var proxyTypes1 []shared.MeshTCPRouteItemSpecToProxyTypes = []shared.MeshTCPRouteItemSpecToProxyTypes{}
				for _, proxyTypesItem1 := range backendRefsItem.ProxyTypes {
					proxyTypes1 = append(proxyTypes1, shared.MeshTCPRouteItemSpecToProxyTypes(proxyTypesItem1.ValueString()))
				}
				sectionName1 := new(string)
				if !backendRefsItem.SectionName.IsUnknown() && !backendRefsItem.SectionName.IsNull() {
					*sectionName1 = backendRefsItem.SectionName.ValueString()
				} else {
					sectionName1 = nil
				}
				tags1 := make(map[string]string)
				for tagsKey1, tagsValue1 := range backendRefsItem.Tags {
					var tagsInst1 string
					tagsInst1 = tagsValue1.ValueString()

					tags1[tagsKey1] = tagsInst1
				}
				weight := new(int64)
				if !backendRefsItem.Weight.IsUnknown() && !backendRefsItem.Weight.IsNull() {
					*weight = backendRefsItem.Weight.ValueInt64()
				} else {
					weight = nil
				}
				backendRefs = append(backendRefs, shared.MeshTCPRouteItemBackendRefs{
					Kind:        kind1,
					Labels:      labels2,
					Mesh:        mesh2,
					Name:        name2,
					Namespace:   namespace1,
					Port:        port,
					ProxyTypes:  proxyTypes1,
					SectionName: sectionName1,
					Tags:        tags1,
					Weight:      weight,
				})
			}
			defaultVar := shared.MeshTCPRouteItemDefault{
				BackendRefs: backendRefs,
			}
			rules = append(rules, shared.MeshTCPRouteItemRules{
				Default: defaultVar,
			})
		}
		kind2 := shared.MeshTCPRouteItemSpecKind(toItem.TargetRef.Kind.ValueString())
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
		var proxyTypes2 []shared.MeshTCPRouteItemSpecProxyTypes = []shared.MeshTCPRouteItemSpecProxyTypes{}
		for _, proxyTypesItem2 := range toItem.TargetRef.ProxyTypes {
			proxyTypes2 = append(proxyTypes2, shared.MeshTCPRouteItemSpecProxyTypes(proxyTypesItem2.ValueString()))
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
		targetRef1 := shared.MeshTCPRouteItemSpecTargetRef{
			Kind:        kind2,
			Labels:      labels3,
			Mesh:        mesh3,
			Name:        name3,
			Namespace:   namespace2,
			ProxyTypes:  proxyTypes2,
			SectionName: sectionName2,
			Tags:        tags2,
		}
		to = append(to, shared.MeshTCPRouteItemTo{
			Rules:     rules,
			TargetRef: targetRef1,
		})
	}
	spec := shared.MeshTCPRouteItemSpec{
		TargetRef: targetRef,
		To:        to,
	}
	out := shared.MeshTCPRouteItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshTCPRouteResourceModel) RefreshFromSharedMeshTCPRouteCreateOrUpdateSuccessResponse(resp *shared.MeshTCPRouteCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshTCPRouteResourceModel) RefreshFromSharedMeshTCPRouteItem(resp *shared.MeshTCPRouteItem) {
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
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
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
				r.Spec.TargetRef.Tags = make(map[string]types.String, len(resp.Spec.TargetRef.Tags))
				for key2, value2 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key2] = types.StringValue(value2)
				}
			}
		}
		r.Spec.To = []tfTypes.MeshTCPRouteItemTo{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshTCPRouteItemTo
			to1.Rules = []tfTypes.MeshTCPRouteItemRules{}
			for rulesCount, rulesItem := range toItem.Rules {
				var rules1 tfTypes.MeshTCPRouteItemRules
				rules1.Default.BackendRefs = []tfTypes.BackendRefs{}
				for backendRefsCount, backendRefsItem := range rulesItem.Default.BackendRefs {
					var backendRefs1 tfTypes.BackendRefs
					backendRefs1.Kind = types.StringValue(string(backendRefsItem.Kind))
					if len(backendRefsItem.Labels) > 0 {
						backendRefs1.Labels = make(map[string]types.String, len(backendRefsItem.Labels))
						for key3, value3 := range backendRefsItem.Labels {
							backendRefs1.Labels[key3] = types.StringValue(value3)
						}
					}
					backendRefs1.Mesh = types.StringPointerValue(backendRefsItem.Mesh)
					backendRefs1.Name = types.StringPointerValue(backendRefsItem.Name)
					backendRefs1.Namespace = types.StringPointerValue(backendRefsItem.Namespace)
					if backendRefsItem.Port != nil {
						backendRefs1.Port = types.Int32Value(int32(*backendRefsItem.Port))
					} else {
						backendRefs1.Port = types.Int32Null()
					}
					backendRefs1.ProxyTypes = make([]types.String, 0, len(backendRefsItem.ProxyTypes))
					for _, v := range backendRefsItem.ProxyTypes {
						backendRefs1.ProxyTypes = append(backendRefs1.ProxyTypes, types.StringValue(string(v)))
					}
					backendRefs1.SectionName = types.StringPointerValue(backendRefsItem.SectionName)
					if len(backendRefsItem.Tags) > 0 {
						backendRefs1.Tags = make(map[string]types.String, len(backendRefsItem.Tags))
						for key4, value4 := range backendRefsItem.Tags {
							backendRefs1.Tags[key4] = types.StringValue(value4)
						}
					}
					backendRefs1.Weight = types.Int64PointerValue(backendRefsItem.Weight)
					if backendRefsCount+1 > len(rules1.Default.BackendRefs) {
						rules1.Default.BackendRefs = append(rules1.Default.BackendRefs, backendRefs1)
					} else {
						rules1.Default.BackendRefs[backendRefsCount].Kind = backendRefs1.Kind
						rules1.Default.BackendRefs[backendRefsCount].Labels = backendRefs1.Labels
						rules1.Default.BackendRefs[backendRefsCount].Mesh = backendRefs1.Mesh
						rules1.Default.BackendRefs[backendRefsCount].Name = backendRefs1.Name
						rules1.Default.BackendRefs[backendRefsCount].Namespace = backendRefs1.Namespace
						rules1.Default.BackendRefs[backendRefsCount].Port = backendRefs1.Port
						rules1.Default.BackendRefs[backendRefsCount].ProxyTypes = backendRefs1.ProxyTypes
						rules1.Default.BackendRefs[backendRefsCount].SectionName = backendRefs1.SectionName
						rules1.Default.BackendRefs[backendRefsCount].Tags = backendRefs1.Tags
						rules1.Default.BackendRefs[backendRefsCount].Weight = backendRefs1.Weight
					}
				}
				if rulesCount+1 > len(to1.Rules) {
					to1.Rules = append(to1.Rules, rules1)
				} else {
					to1.Rules[rulesCount].Default = rules1.Default
				}
			}
			to1.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String, len(toItem.TargetRef.Labels))
				for key5, value5 := range toItem.TargetRef.Labels {
					to1.TargetRef.Labels[key5] = types.StringValue(value5)
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
				to1.TargetRef.Tags = make(map[string]types.String, len(toItem.TargetRef.Tags))
				for key6, value6 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key6] = types.StringValue(value6)
				}
			}
			if toCount+1 > len(r.Spec.To) {
				r.Spec.To = append(r.Spec.To, to1)
			} else {
				r.Spec.To[toCount].Rules = to1.Rules
				r.Spec.To[toCount].TargetRef = to1.TargetRef
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
