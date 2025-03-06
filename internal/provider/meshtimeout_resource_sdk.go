// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshTimeoutResourceModel) ToSharedMeshTimeoutItemInput() *shared.MeshTimeoutItemInput {
	typeVar := shared.MeshTimeoutItemType(r.Type.ValueString())
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
	var from []shared.MeshTimeoutItemFrom = []shared.MeshTimeoutItemFrom{}
	for _, fromItem := range r.Spec.From {
		var defaultVar *shared.MeshTimeoutItemDefault
		if fromItem.Default != nil {
			connectionTimeout := new(string)
			if !fromItem.Default.ConnectionTimeout.IsUnknown() && !fromItem.Default.ConnectionTimeout.IsNull() {
				*connectionTimeout = fromItem.Default.ConnectionTimeout.ValueString()
			} else {
				connectionTimeout = nil
			}
			var http *shared.MeshTimeoutItemHTTP
			if fromItem.Default.HTTP != nil {
				maxConnectionDuration := new(string)
				if !fromItem.Default.HTTP.MaxConnectionDuration.IsUnknown() && !fromItem.Default.HTTP.MaxConnectionDuration.IsNull() {
					*maxConnectionDuration = fromItem.Default.HTTP.MaxConnectionDuration.ValueString()
				} else {
					maxConnectionDuration = nil
				}
				maxStreamDuration := new(string)
				if !fromItem.Default.HTTP.MaxStreamDuration.IsUnknown() && !fromItem.Default.HTTP.MaxStreamDuration.IsNull() {
					*maxStreamDuration = fromItem.Default.HTTP.MaxStreamDuration.ValueString()
				} else {
					maxStreamDuration = nil
				}
				requestHeadersTimeout := new(string)
				if !fromItem.Default.HTTP.RequestHeadersTimeout.IsUnknown() && !fromItem.Default.HTTP.RequestHeadersTimeout.IsNull() {
					*requestHeadersTimeout = fromItem.Default.HTTP.RequestHeadersTimeout.ValueString()
				} else {
					requestHeadersTimeout = nil
				}
				requestTimeout := new(string)
				if !fromItem.Default.HTTP.RequestTimeout.IsUnknown() && !fromItem.Default.HTTP.RequestTimeout.IsNull() {
					*requestTimeout = fromItem.Default.HTTP.RequestTimeout.ValueString()
				} else {
					requestTimeout = nil
				}
				streamIdleTimeout := new(string)
				if !fromItem.Default.HTTP.StreamIdleTimeout.IsUnknown() && !fromItem.Default.HTTP.StreamIdleTimeout.IsNull() {
					*streamIdleTimeout = fromItem.Default.HTTP.StreamIdleTimeout.ValueString()
				} else {
					streamIdleTimeout = nil
				}
				http = &shared.MeshTimeoutItemHTTP{
					MaxConnectionDuration: maxConnectionDuration,
					MaxStreamDuration:     maxStreamDuration,
					RequestHeadersTimeout: requestHeadersTimeout,
					RequestTimeout:        requestTimeout,
					StreamIdleTimeout:     streamIdleTimeout,
				}
			}
			idleTimeout := new(string)
			if !fromItem.Default.IdleTimeout.IsUnknown() && !fromItem.Default.IdleTimeout.IsNull() {
				*idleTimeout = fromItem.Default.IdleTimeout.ValueString()
			} else {
				idleTimeout = nil
			}
			defaultVar = &shared.MeshTimeoutItemDefault{
				ConnectionTimeout: connectionTimeout,
				HTTP:              http,
				IdleTimeout:       idleTimeout,
			}
		}
		kind := shared.MeshTimeoutItemSpecKind(fromItem.TargetRef.Kind.ValueString())
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
		var proxyTypes []shared.MeshTimeoutItemSpecProxyTypes = []shared.MeshTimeoutItemSpecProxyTypes{}
		for _, proxyTypesItem := range fromItem.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshTimeoutItemSpecProxyTypes(proxyTypesItem.ValueString()))
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
		targetRef := shared.MeshTimeoutItemSpecTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name1,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags,
		}
		from = append(from, shared.MeshTimeoutItemFrom{
			Default:   defaultVar,
			TargetRef: targetRef,
		})
	}
	var rules []shared.MeshTimeoutItemRules = []shared.MeshTimeoutItemRules{}
	for _, rulesItem := range r.Spec.Rules {
		var default1 *shared.MeshTimeoutItemSpecDefault
		if rulesItem.Default != nil {
			connectionTimeout1 := new(string)
			if !rulesItem.Default.ConnectionTimeout.IsUnknown() && !rulesItem.Default.ConnectionTimeout.IsNull() {
				*connectionTimeout1 = rulesItem.Default.ConnectionTimeout.ValueString()
			} else {
				connectionTimeout1 = nil
			}
			var http1 *shared.MeshTimeoutItemSpecHTTP
			if rulesItem.Default.HTTP != nil {
				maxConnectionDuration1 := new(string)
				if !rulesItem.Default.HTTP.MaxConnectionDuration.IsUnknown() && !rulesItem.Default.HTTP.MaxConnectionDuration.IsNull() {
					*maxConnectionDuration1 = rulesItem.Default.HTTP.MaxConnectionDuration.ValueString()
				} else {
					maxConnectionDuration1 = nil
				}
				maxStreamDuration1 := new(string)
				if !rulesItem.Default.HTTP.MaxStreamDuration.IsUnknown() && !rulesItem.Default.HTTP.MaxStreamDuration.IsNull() {
					*maxStreamDuration1 = rulesItem.Default.HTTP.MaxStreamDuration.ValueString()
				} else {
					maxStreamDuration1 = nil
				}
				requestHeadersTimeout1 := new(string)
				if !rulesItem.Default.HTTP.RequestHeadersTimeout.IsUnknown() && !rulesItem.Default.HTTP.RequestHeadersTimeout.IsNull() {
					*requestHeadersTimeout1 = rulesItem.Default.HTTP.RequestHeadersTimeout.ValueString()
				} else {
					requestHeadersTimeout1 = nil
				}
				requestTimeout1 := new(string)
				if !rulesItem.Default.HTTP.RequestTimeout.IsUnknown() && !rulesItem.Default.HTTP.RequestTimeout.IsNull() {
					*requestTimeout1 = rulesItem.Default.HTTP.RequestTimeout.ValueString()
				} else {
					requestTimeout1 = nil
				}
				streamIdleTimeout1 := new(string)
				if !rulesItem.Default.HTTP.StreamIdleTimeout.IsUnknown() && !rulesItem.Default.HTTP.StreamIdleTimeout.IsNull() {
					*streamIdleTimeout1 = rulesItem.Default.HTTP.StreamIdleTimeout.ValueString()
				} else {
					streamIdleTimeout1 = nil
				}
				http1 = &shared.MeshTimeoutItemSpecHTTP{
					MaxConnectionDuration: maxConnectionDuration1,
					MaxStreamDuration:     maxStreamDuration1,
					RequestHeadersTimeout: requestHeadersTimeout1,
					RequestTimeout:        requestTimeout1,
					StreamIdleTimeout:     streamIdleTimeout1,
				}
			}
			idleTimeout1 := new(string)
			if !rulesItem.Default.IdleTimeout.IsUnknown() && !rulesItem.Default.IdleTimeout.IsNull() {
				*idleTimeout1 = rulesItem.Default.IdleTimeout.ValueString()
			} else {
				idleTimeout1 = nil
			}
			default1 = &shared.MeshTimeoutItemSpecDefault{
				ConnectionTimeout: connectionTimeout1,
				HTTP:              http1,
				IdleTimeout:       idleTimeout1,
			}
		}
		rules = append(rules, shared.MeshTimeoutItemRules{
			Default: default1,
		})
	}
	var targetRef1 *shared.MeshTimeoutItemTargetRef
	if r.Spec.TargetRef != nil {
		kind1 := shared.MeshTimeoutItemKind(r.Spec.TargetRef.Kind.ValueString())
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
		var proxyTypes1 []shared.MeshTimeoutItemProxyTypes = []shared.MeshTimeoutItemProxyTypes{}
		for _, proxyTypesItem1 := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes1 = append(proxyTypes1, shared.MeshTimeoutItemProxyTypes(proxyTypesItem1.ValueString()))
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
		targetRef1 = &shared.MeshTimeoutItemTargetRef{
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
	var to []shared.MeshTimeoutItemTo = []shared.MeshTimeoutItemTo{}
	for _, toItem := range r.Spec.To {
		var default2 *shared.MeshTimeoutItemSpecToDefault
		if toItem.Default != nil {
			connectionTimeout2 := new(string)
			if !toItem.Default.ConnectionTimeout.IsUnknown() && !toItem.Default.ConnectionTimeout.IsNull() {
				*connectionTimeout2 = toItem.Default.ConnectionTimeout.ValueString()
			} else {
				connectionTimeout2 = nil
			}
			var http2 *shared.MeshTimeoutItemSpecToHTTP
			if toItem.Default.HTTP != nil {
				maxConnectionDuration2 := new(string)
				if !toItem.Default.HTTP.MaxConnectionDuration.IsUnknown() && !toItem.Default.HTTP.MaxConnectionDuration.IsNull() {
					*maxConnectionDuration2 = toItem.Default.HTTP.MaxConnectionDuration.ValueString()
				} else {
					maxConnectionDuration2 = nil
				}
				maxStreamDuration2 := new(string)
				if !toItem.Default.HTTP.MaxStreamDuration.IsUnknown() && !toItem.Default.HTTP.MaxStreamDuration.IsNull() {
					*maxStreamDuration2 = toItem.Default.HTTP.MaxStreamDuration.ValueString()
				} else {
					maxStreamDuration2 = nil
				}
				requestHeadersTimeout2 := new(string)
				if !toItem.Default.HTTP.RequestHeadersTimeout.IsUnknown() && !toItem.Default.HTTP.RequestHeadersTimeout.IsNull() {
					*requestHeadersTimeout2 = toItem.Default.HTTP.RequestHeadersTimeout.ValueString()
				} else {
					requestHeadersTimeout2 = nil
				}
				requestTimeout2 := new(string)
				if !toItem.Default.HTTP.RequestTimeout.IsUnknown() && !toItem.Default.HTTP.RequestTimeout.IsNull() {
					*requestTimeout2 = toItem.Default.HTTP.RequestTimeout.ValueString()
				} else {
					requestTimeout2 = nil
				}
				streamIdleTimeout2 := new(string)
				if !toItem.Default.HTTP.StreamIdleTimeout.IsUnknown() && !toItem.Default.HTTP.StreamIdleTimeout.IsNull() {
					*streamIdleTimeout2 = toItem.Default.HTTP.StreamIdleTimeout.ValueString()
				} else {
					streamIdleTimeout2 = nil
				}
				http2 = &shared.MeshTimeoutItemSpecToHTTP{
					MaxConnectionDuration: maxConnectionDuration2,
					MaxStreamDuration:     maxStreamDuration2,
					RequestHeadersTimeout: requestHeadersTimeout2,
					RequestTimeout:        requestTimeout2,
					StreamIdleTimeout:     streamIdleTimeout2,
				}
			}
			idleTimeout2 := new(string)
			if !toItem.Default.IdleTimeout.IsUnknown() && !toItem.Default.IdleTimeout.IsNull() {
				*idleTimeout2 = toItem.Default.IdleTimeout.ValueString()
			} else {
				idleTimeout2 = nil
			}
			default2 = &shared.MeshTimeoutItemSpecToDefault{
				ConnectionTimeout: connectionTimeout2,
				HTTP:              http2,
				IdleTimeout:       idleTimeout2,
			}
		}
		kind2 := shared.MeshTimeoutItemSpecToKind(toItem.TargetRef.Kind.ValueString())
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
		var proxyTypes2 []shared.MeshTimeoutItemSpecToProxyTypes = []shared.MeshTimeoutItemSpecToProxyTypes{}
		for _, proxyTypesItem2 := range toItem.TargetRef.ProxyTypes {
			proxyTypes2 = append(proxyTypes2, shared.MeshTimeoutItemSpecToProxyTypes(proxyTypesItem2.ValueString()))
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
		targetRef2 := shared.MeshTimeoutItemSpecToTargetRef{
			Kind:        kind2,
			Labels:      labels3,
			Mesh:        mesh3,
			Name:        name3,
			Namespace:   namespace2,
			ProxyTypes:  proxyTypes2,
			SectionName: sectionName2,
			Tags:        tags2,
		}
		to = append(to, shared.MeshTimeoutItemTo{
			Default:   default2,
			TargetRef: targetRef2,
		})
	}
	spec := shared.MeshTimeoutItemSpec{
		From:      from,
		Rules:     rules,
		TargetRef: targetRef1,
		To:        to,
	}
	out := shared.MeshTimeoutItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}
	return &out
}

func (r *MeshTimeoutResourceModel) RefreshFromSharedMeshTimeoutCreateOrUpdateSuccessResponse(resp *shared.MeshTimeoutCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshTimeoutResourceModel) RefreshFromSharedMeshTimeoutItem(resp *shared.MeshTimeoutItem) {
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
		r.Spec.From = []tfTypes.MeshTimeoutItemFrom{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from1 tfTypes.MeshTimeoutItemFrom
			if fromItem.Default == nil {
				from1.Default = nil
			} else {
				from1.Default = &tfTypes.MeshTimeoutItemDefault{}
				from1.Default.ConnectionTimeout = types.StringPointerValue(fromItem.Default.ConnectionTimeout)
				if fromItem.Default.HTTP == nil {
					from1.Default.HTTP = nil
				} else {
					from1.Default.HTTP = &tfTypes.MeshTimeoutItemHTTP{}
					from1.Default.HTTP.MaxConnectionDuration = types.StringPointerValue(fromItem.Default.HTTP.MaxConnectionDuration)
					from1.Default.HTTP.MaxStreamDuration = types.StringPointerValue(fromItem.Default.HTTP.MaxStreamDuration)
					from1.Default.HTTP.RequestHeadersTimeout = types.StringPointerValue(fromItem.Default.HTTP.RequestHeadersTimeout)
					from1.Default.HTTP.RequestTimeout = types.StringPointerValue(fromItem.Default.HTTP.RequestTimeout)
					from1.Default.HTTP.StreamIdleTimeout = types.StringPointerValue(fromItem.Default.HTTP.StreamIdleTimeout)
				}
				from1.Default.IdleTimeout = types.StringPointerValue(fromItem.Default.IdleTimeout)
			}
			from1.TargetRef.Kind = types.StringValue(string(fromItem.TargetRef.Kind))
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
		r.Spec.Rules = []tfTypes.MeshTimeoutItemRules{}
		if len(r.Spec.Rules) > len(resp.Spec.Rules) {
			r.Spec.Rules = r.Spec.Rules[:len(resp.Spec.Rules)]
		}
		for rulesCount, rulesItem := range resp.Spec.Rules {
			var rules1 tfTypes.MeshTimeoutItemRules
			if rulesItem.Default == nil {
				rules1.Default = nil
			} else {
				rules1.Default = &tfTypes.MeshTimeoutItemDefault{}
				rules1.Default.ConnectionTimeout = types.StringPointerValue(rulesItem.Default.ConnectionTimeout)
				if rulesItem.Default.HTTP == nil {
					rules1.Default.HTTP = nil
				} else {
					rules1.Default.HTTP = &tfTypes.MeshTimeoutItemHTTP{}
					rules1.Default.HTTP.MaxConnectionDuration = types.StringPointerValue(rulesItem.Default.HTTP.MaxConnectionDuration)
					rules1.Default.HTTP.MaxStreamDuration = types.StringPointerValue(rulesItem.Default.HTTP.MaxStreamDuration)
					rules1.Default.HTTP.RequestHeadersTimeout = types.StringPointerValue(rulesItem.Default.HTTP.RequestHeadersTimeout)
					rules1.Default.HTTP.RequestTimeout = types.StringPointerValue(rulesItem.Default.HTTP.RequestTimeout)
					rules1.Default.HTTP.StreamIdleTimeout = types.StringPointerValue(rulesItem.Default.HTTP.StreamIdleTimeout)
				}
				rules1.Default.IdleTimeout = types.StringPointerValue(rulesItem.Default.IdleTimeout)
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
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
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
		r.Spec.To = []tfTypes.MeshTimeoutItemFrom{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to1 tfTypes.MeshTimeoutItemFrom
			if toItem.Default == nil {
				to1.Default = nil
			} else {
				to1.Default = &tfTypes.MeshTimeoutItemDefault{}
				to1.Default.ConnectionTimeout = types.StringPointerValue(toItem.Default.ConnectionTimeout)
				if toItem.Default.HTTP == nil {
					to1.Default.HTTP = nil
				} else {
					to1.Default.HTTP = &tfTypes.MeshTimeoutItemHTTP{}
					to1.Default.HTTP.MaxConnectionDuration = types.StringPointerValue(toItem.Default.HTTP.MaxConnectionDuration)
					to1.Default.HTTP.MaxStreamDuration = types.StringPointerValue(toItem.Default.HTTP.MaxStreamDuration)
					to1.Default.HTTP.RequestHeadersTimeout = types.StringPointerValue(toItem.Default.HTTP.RequestHeadersTimeout)
					to1.Default.HTTP.RequestTimeout = types.StringPointerValue(toItem.Default.HTTP.RequestTimeout)
					to1.Default.HTTP.StreamIdleTimeout = types.StringPointerValue(toItem.Default.HTTP.StreamIdleTimeout)
				}
				to1.Default.IdleTimeout = types.StringPointerValue(toItem.Default.IdleTimeout)
			}
			to1.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
			if len(toItem.TargetRef.Labels) > 0 {
				to1.TargetRef.Labels = make(map[string]types.String)
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
				to1.TargetRef.Tags = make(map[string]types.String)
				for key6, value6 := range toItem.TargetRef.Tags {
					to1.TargetRef.Tags[key6] = types.StringValue(value6)
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
