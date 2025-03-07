// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshPassthroughResourceModel) ToSharedMeshPassthroughItemInput() *shared.MeshPassthroughItemInput {
	typeVar := shared.MeshPassthroughItemType(r.Type.ValueString())
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
	var defaultVar *shared.MeshPassthroughItemDefault
	if r.Spec.Default != nil {
		var appendMatch []shared.AppendMatch = []shared.AppendMatch{}
		for _, appendMatchItem := range r.Spec.Default.AppendMatch {
			port := new(int)
			if !appendMatchItem.Port.IsUnknown() && !appendMatchItem.Port.IsNull() {
				*port = int(appendMatchItem.Port.ValueInt32())
			} else {
				port = nil
			}
			protocol := new(shared.MeshPassthroughItemProtocol)
			if !appendMatchItem.Protocol.IsUnknown() && !appendMatchItem.Protocol.IsNull() {
				*protocol = shared.MeshPassthroughItemProtocol(appendMatchItem.Protocol.ValueString())
			} else {
				protocol = nil
			}
			type1 := shared.MeshPassthroughItemSpecType(appendMatchItem.Type.ValueString())
			var value string
			value = appendMatchItem.Value.ValueString()

			appendMatch = append(appendMatch, shared.AppendMatch{
				Port:     port,
				Protocol: protocol,
				Type:     type1,
				Value:    value,
			})
		}
		passthroughMode := new(shared.PassthroughMode)
		if !r.Spec.Default.PassthroughMode.IsUnknown() && !r.Spec.Default.PassthroughMode.IsNull() {
			*passthroughMode = shared.PassthroughMode(r.Spec.Default.PassthroughMode.ValueString())
		} else {
			passthroughMode = nil
		}
		defaultVar = &shared.MeshPassthroughItemDefault{
			AppendMatch:     appendMatch,
			PassthroughMode: passthroughMode,
		}
	}
	var targetRef *shared.MeshPassthroughItemTargetRef
	if r.Spec.TargetRef != nil {
		kind := shared.MeshPassthroughItemKind(r.Spec.TargetRef.Kind.ValueString())
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
		var proxyTypes []shared.MeshPassthroughItemProxyTypes = []shared.MeshPassthroughItemProxyTypes{}
		for _, proxyTypesItem := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshPassthroughItemProxyTypes(proxyTypesItem.ValueString()))
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
		targetRef = &shared.MeshPassthroughItemTargetRef{
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
	spec := shared.MeshPassthroughItemSpec{
		Default:   defaultVar,
		TargetRef: targetRef,
	}
	out := shared.MeshPassthroughItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshPassthroughResourceModel) RefreshFromSharedMeshPassthroughCreateOrUpdateSuccessResponse(resp *shared.MeshPassthroughCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshPassthroughResourceModel) RefreshFromSharedMeshPassthroughItem(resp *shared.MeshPassthroughItem) {
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
		if resp.Spec.Default == nil {
			r.Spec.Default = nil
		} else {
			r.Spec.Default = &tfTypes.MeshPassthroughItemDefault{}
			r.Spec.Default.AppendMatch = []tfTypes.AppendMatch{}
			if len(r.Spec.Default.AppendMatch) > len(resp.Spec.Default.AppendMatch) {
				r.Spec.Default.AppendMatch = r.Spec.Default.AppendMatch[:len(resp.Spec.Default.AppendMatch)]
			}
			for appendMatchCount, appendMatchItem := range resp.Spec.Default.AppendMatch {
				var appendMatch1 tfTypes.AppendMatch
				if appendMatchItem.Port != nil {
					appendMatch1.Port = types.Int32Value(int32(*appendMatchItem.Port))
				} else {
					appendMatch1.Port = types.Int32Null()
				}
				if appendMatchItem.Protocol != nil {
					appendMatch1.Protocol = types.StringValue(string(*appendMatchItem.Protocol))
				} else {
					appendMatch1.Protocol = types.StringNull()
				}
				appendMatch1.Type = types.StringValue(string(appendMatchItem.Type))
				appendMatch1.Value = types.StringValue(appendMatchItem.Value)
				if appendMatchCount+1 > len(r.Spec.Default.AppendMatch) {
					r.Spec.Default.AppendMatch = append(r.Spec.Default.AppendMatch, appendMatch1)
				} else {
					r.Spec.Default.AppendMatch[appendMatchCount].Port = appendMatch1.Port
					r.Spec.Default.AppendMatch[appendMatchCount].Protocol = appendMatch1.Protocol
					r.Spec.Default.AppendMatch[appendMatchCount].Type = appendMatch1.Type
					r.Spec.Default.AppendMatch[appendMatchCount].Value = appendMatch1.Value
				}
			}
			if resp.Spec.Default.PassthroughMode != nil {
				r.Spec.Default.PassthroughMode = types.StringValue(string(*resp.Spec.Default.PassthroughMode))
			} else {
				r.Spec.Default.PassthroughMode = types.StringNull()
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
				for key1, value2 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key1] = types.StringValue(value2)
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
				for key2, value3 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key2] = types.StringValue(value3)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
