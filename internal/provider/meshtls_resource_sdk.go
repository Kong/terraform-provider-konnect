// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshTLSResourceModel) ToSharedMeshTLSItemInput() *shared.MeshTLSItemInput {
	typeVar := shared.MeshTLSItemType(r.Type.ValueString())
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
	var from []shared.MeshTLSItemFrom = []shared.MeshTLSItemFrom{}
	for _, fromItem := range r.Spec.From {
		var defaultVar *shared.MeshTLSItemDefault
		if fromItem.Default != nil {
			mode := new(shared.MeshTLSItemMode)
			if !fromItem.Default.Mode.IsUnknown() && !fromItem.Default.Mode.IsNull() {
				*mode = shared.MeshTLSItemMode(fromItem.Default.Mode.ValueString())
			} else {
				mode = nil
			}
			var tlsCiphers []shared.TLSCiphers = []shared.TLSCiphers{}
			for _, tlsCiphersItem := range fromItem.Default.TLSCiphers {
				tlsCiphers = append(tlsCiphers, shared.TLSCiphers(tlsCiphersItem.ValueString()))
			}
			var tlsVersion *shared.TLSVersion
			if fromItem.Default.TLSVersion != nil {
				max := new(shared.MeshTLSItemMax)
				if !fromItem.Default.TLSVersion.Max.IsUnknown() && !fromItem.Default.TLSVersion.Max.IsNull() {
					*max = shared.MeshTLSItemMax(fromItem.Default.TLSVersion.Max.ValueString())
				} else {
					max = nil
				}
				min := new(shared.MeshTLSItemMin)
				if !fromItem.Default.TLSVersion.Min.IsUnknown() && !fromItem.Default.TLSVersion.Min.IsNull() {
					*min = shared.MeshTLSItemMin(fromItem.Default.TLSVersion.Min.ValueString())
				} else {
					min = nil
				}
				tlsVersion = &shared.TLSVersion{
					Max: max,
					Min: min,
				}
			}
			defaultVar = &shared.MeshTLSItemDefault{
				Mode:       mode,
				TLSCiphers: tlsCiphers,
				TLSVersion: tlsVersion,
			}
		}
		kind := new(shared.MeshTLSItemSpecKind)
		if !fromItem.TargetRef.Kind.IsUnknown() && !fromItem.TargetRef.Kind.IsNull() {
			*kind = shared.MeshTLSItemSpecKind(fromItem.TargetRef.Kind.ValueString())
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
		var proxyTypes []shared.MeshTLSItemSpecProxyTypes = []shared.MeshTLSItemSpecProxyTypes{}
		for _, proxyTypesItem := range fromItem.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshTLSItemSpecProxyTypes(proxyTypesItem.ValueString()))
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
		targetRef := shared.MeshTLSItemSpecTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name1,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags,
		}
		from = append(from, shared.MeshTLSItemFrom{
			Default:   defaultVar,
			TargetRef: targetRef,
		})
	}
	var rules []shared.MeshTLSItemRules = []shared.MeshTLSItemRules{}
	for _, rulesItem := range r.Spec.Rules {
		var default1 *shared.MeshTLSItemSpecDefault
		if rulesItem.Default != nil {
			mode1 := new(shared.MeshTLSItemSpecMode)
			if !rulesItem.Default.Mode.IsUnknown() && !rulesItem.Default.Mode.IsNull() {
				*mode1 = shared.MeshTLSItemSpecMode(rulesItem.Default.Mode.ValueString())
			} else {
				mode1 = nil
			}
			var tlsCiphers1 []shared.MeshTLSItemTLSCiphers = []shared.MeshTLSItemTLSCiphers{}
			for _, tlsCiphersItem1 := range rulesItem.Default.TLSCiphers {
				tlsCiphers1 = append(tlsCiphers1, shared.MeshTLSItemTLSCiphers(tlsCiphersItem1.ValueString()))
			}
			var tlsVersion1 *shared.MeshTLSItemTLSVersion
			if rulesItem.Default.TLSVersion != nil {
				max1 := new(shared.MeshTLSItemSpecMax)
				if !rulesItem.Default.TLSVersion.Max.IsUnknown() && !rulesItem.Default.TLSVersion.Max.IsNull() {
					*max1 = shared.MeshTLSItemSpecMax(rulesItem.Default.TLSVersion.Max.ValueString())
				} else {
					max1 = nil
				}
				min1 := new(shared.MeshTLSItemSpecMin)
				if !rulesItem.Default.TLSVersion.Min.IsUnknown() && !rulesItem.Default.TLSVersion.Min.IsNull() {
					*min1 = shared.MeshTLSItemSpecMin(rulesItem.Default.TLSVersion.Min.ValueString())
				} else {
					min1 = nil
				}
				tlsVersion1 = &shared.MeshTLSItemTLSVersion{
					Max: max1,
					Min: min1,
				}
			}
			default1 = &shared.MeshTLSItemSpecDefault{
				Mode:       mode1,
				TLSCiphers: tlsCiphers1,
				TLSVersion: tlsVersion1,
			}
		}
		rules = append(rules, shared.MeshTLSItemRules{
			Default: default1,
		})
	}
	var targetRef1 *shared.MeshTLSItemTargetRef
	if r.Spec.TargetRef != nil {
		kind1 := new(shared.MeshTLSItemKind)
		if !r.Spec.TargetRef.Kind.IsUnknown() && !r.Spec.TargetRef.Kind.IsNull() {
			*kind1 = shared.MeshTLSItemKind(r.Spec.TargetRef.Kind.ValueString())
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
		var proxyTypes1 []shared.MeshTLSItemProxyTypes = []shared.MeshTLSItemProxyTypes{}
		for _, proxyTypesItem1 := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes1 = append(proxyTypes1, shared.MeshTLSItemProxyTypes(proxyTypesItem1.ValueString()))
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
		targetRef1 = &shared.MeshTLSItemTargetRef{
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
	spec := shared.MeshTLSItemSpec{
		From:      from,
		Rules:     rules,
		TargetRef: targetRef1,
	}
	out := shared.MeshTLSItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshTLSResourceModel) RefreshFromSharedMeshTLSCreateOrUpdateSuccessResponse(resp *shared.MeshTLSCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshTLSResourceModel) RefreshFromSharedMeshTLSItem(resp *shared.MeshTLSItem) {
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
		r.Spec.From = []tfTypes.MeshTLSItemFrom{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
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
					from1.Default.TLSVersion = &tfTypes.MeshExternalServiceItemVersion{}
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
			if fromCount+1 > len(r.Spec.From) {
				r.Spec.From = append(r.Spec.From, from1)
			} else {
				r.Spec.From[fromCount].Default = from1.Default
				r.Spec.From[fromCount].TargetRef = from1.TargetRef
			}
		}
		r.Spec.Rules = []tfTypes.MeshTLSItemRules{}
		if len(r.Spec.Rules) > len(resp.Spec.Rules) {
			r.Spec.Rules = r.Spec.Rules[:len(resp.Spec.Rules)]
		}
		for rulesCount, rulesItem := range resp.Spec.Rules {
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
					rules1.Default.TLSVersion = &tfTypes.MeshExternalServiceItemVersion{}
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
			if rulesCount+1 > len(r.Spec.Rules) {
				r.Spec.Rules = append(r.Spec.Rules, rules1)
			} else {
				r.Spec.Rules[rulesCount].Default = rules1.Default
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
				for key3, value3 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key3] = types.StringValue(value3)
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
				for key4, value4 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key4] = types.StringValue(value4)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
