// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshProxyPatchResourceModel) ToSharedMeshProxyPatchItem() *shared.MeshProxyPatchItem {
	typeVar := shared.MeshProxyPatchItemType(r.Type.ValueString())
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
	var appendModifications []shared.AppendModifications = []shared.AppendModifications{}
	for _, appendModificationsItem := range r.Spec.Default.AppendModifications {
		var cluster *shared.Cluster
		if appendModificationsItem.Cluster != nil {
			var jsonPatches []shared.JSONPatches = []shared.JSONPatches{}
			for _, jsonPatchesItem := range appendModificationsItem.Cluster.JSONPatches {
				from := new(string)
				if !jsonPatchesItem.From.IsUnknown() && !jsonPatchesItem.From.IsNull() {
					*from = jsonPatchesItem.From.ValueString()
				} else {
					from = nil
				}
				op := shared.Op(jsonPatchesItem.Op.ValueString())
				var path string
				path = jsonPatchesItem.Path.ValueString()

				var value interface{}
				if !jsonPatchesItem.Value.IsUnknown() && !jsonPatchesItem.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem.Value.ValueString()), &value)
				}
				jsonPatches = append(jsonPatches, shared.JSONPatches{
					From:  from,
					Op:    op,
					Path:  path,
					Value: value,
				})
			}
			var match *shared.MeshProxyPatchItemMatch
			if appendModificationsItem.Cluster.Match != nil {
				name1 := new(string)
				if !appendModificationsItem.Cluster.Match.Name.IsUnknown() && !appendModificationsItem.Cluster.Match.Name.IsNull() {
					*name1 = appendModificationsItem.Cluster.Match.Name.ValueString()
				} else {
					name1 = nil
				}
				origin := new(string)
				if !appendModificationsItem.Cluster.Match.Origin.IsUnknown() && !appendModificationsItem.Cluster.Match.Origin.IsNull() {
					*origin = appendModificationsItem.Cluster.Match.Origin.ValueString()
				} else {
					origin = nil
				}
				match = &shared.MeshProxyPatchItemMatch{
					Name:   name1,
					Origin: origin,
				}
			}
			operation := shared.Operation(appendModificationsItem.Cluster.Operation.ValueString())
			value1 := new(string)
			if !appendModificationsItem.Cluster.Value.IsUnknown() && !appendModificationsItem.Cluster.Value.IsNull() {
				*value1 = appendModificationsItem.Cluster.Value.ValueString()
			} else {
				value1 = nil
			}
			cluster = &shared.Cluster{
				JSONPatches: jsonPatches,
				Match:       match,
				Operation:   operation,
				Value:       value1,
			}
		}
		var httpFilter *shared.HTTPFilter
		if appendModificationsItem.HTTPFilter != nil {
			var jsonPatches1 []shared.MeshProxyPatchItemJSONPatches = []shared.MeshProxyPatchItemJSONPatches{}
			for _, jsonPatchesItem1 := range appendModificationsItem.HTTPFilter.JSONPatches {
				from1 := new(string)
				if !jsonPatchesItem1.From.IsUnknown() && !jsonPatchesItem1.From.IsNull() {
					*from1 = jsonPatchesItem1.From.ValueString()
				} else {
					from1 = nil
				}
				op1 := shared.MeshProxyPatchItemOp(jsonPatchesItem1.Op.ValueString())
				var path1 string
				path1 = jsonPatchesItem1.Path.ValueString()

				var value2 interface{}
				if !jsonPatchesItem1.Value.IsUnknown() && !jsonPatchesItem1.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem1.Value.ValueString()), &value2)
				}
				jsonPatches1 = append(jsonPatches1, shared.MeshProxyPatchItemJSONPatches{
					From:  from1,
					Op:    op1,
					Path:  path1,
					Value: value2,
				})
			}
			var match1 *shared.MeshProxyPatchItemSpecMatch
			if appendModificationsItem.HTTPFilter.Match != nil {
				listenerName := new(string)
				if !appendModificationsItem.HTTPFilter.Match.ListenerName.IsUnknown() && !appendModificationsItem.HTTPFilter.Match.ListenerName.IsNull() {
					*listenerName = appendModificationsItem.HTTPFilter.Match.ListenerName.ValueString()
				} else {
					listenerName = nil
				}
				listenerTags := make(map[string]string)
				for listenerTagsKey, listenerTagsValue := range appendModificationsItem.HTTPFilter.Match.ListenerTags {
					var listenerTagsInst string
					listenerTagsInst = listenerTagsValue.ValueString()

					listenerTags[listenerTagsKey] = listenerTagsInst
				}
				name2 := new(string)
				if !appendModificationsItem.HTTPFilter.Match.Name.IsUnknown() && !appendModificationsItem.HTTPFilter.Match.Name.IsNull() {
					*name2 = appendModificationsItem.HTTPFilter.Match.Name.ValueString()
				} else {
					name2 = nil
				}
				origin1 := new(string)
				if !appendModificationsItem.HTTPFilter.Match.Origin.IsUnknown() && !appendModificationsItem.HTTPFilter.Match.Origin.IsNull() {
					*origin1 = appendModificationsItem.HTTPFilter.Match.Origin.ValueString()
				} else {
					origin1 = nil
				}
				match1 = &shared.MeshProxyPatchItemSpecMatch{
					ListenerName: listenerName,
					ListenerTags: listenerTags,
					Name:         name2,
					Origin:       origin1,
				}
			}
			operation1 := shared.MeshProxyPatchItemOperation(appendModificationsItem.HTTPFilter.Operation.ValueString())
			value3 := new(string)
			if !appendModificationsItem.HTTPFilter.Value.IsUnknown() && !appendModificationsItem.HTTPFilter.Value.IsNull() {
				*value3 = appendModificationsItem.HTTPFilter.Value.ValueString()
			} else {
				value3 = nil
			}
			httpFilter = &shared.HTTPFilter{
				JSONPatches: jsonPatches1,
				Match:       match1,
				Operation:   operation1,
				Value:       value3,
			}
		}
		var listener *shared.Listener
		if appendModificationsItem.Listener != nil {
			var jsonPatches2 []shared.MeshProxyPatchItemSpecJSONPatches = []shared.MeshProxyPatchItemSpecJSONPatches{}
			for _, jsonPatchesItem2 := range appendModificationsItem.Listener.JSONPatches {
				from2 := new(string)
				if !jsonPatchesItem2.From.IsUnknown() && !jsonPatchesItem2.From.IsNull() {
					*from2 = jsonPatchesItem2.From.ValueString()
				} else {
					from2 = nil
				}
				op2 := shared.MeshProxyPatchItemSpecOp(jsonPatchesItem2.Op.ValueString())
				var path2 string
				path2 = jsonPatchesItem2.Path.ValueString()

				var value4 interface{}
				if !jsonPatchesItem2.Value.IsUnknown() && !jsonPatchesItem2.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem2.Value.ValueString()), &value4)
				}
				jsonPatches2 = append(jsonPatches2, shared.MeshProxyPatchItemSpecJSONPatches{
					From:  from2,
					Op:    op2,
					Path:  path2,
					Value: value4,
				})
			}
			var match2 *shared.MeshProxyPatchItemSpecDefaultMatch
			if appendModificationsItem.Listener.Match != nil {
				name3 := new(string)
				if !appendModificationsItem.Listener.Match.Name.IsUnknown() && !appendModificationsItem.Listener.Match.Name.IsNull() {
					*name3 = appendModificationsItem.Listener.Match.Name.ValueString()
				} else {
					name3 = nil
				}
				origin2 := new(string)
				if !appendModificationsItem.Listener.Match.Origin.IsUnknown() && !appendModificationsItem.Listener.Match.Origin.IsNull() {
					*origin2 = appendModificationsItem.Listener.Match.Origin.ValueString()
				} else {
					origin2 = nil
				}
				tags := make(map[string]string)
				for tagsKey, tagsValue := range appendModificationsItem.Listener.Match.Tags {
					var tagsInst string
					tagsInst = tagsValue.ValueString()

					tags[tagsKey] = tagsInst
				}
				match2 = &shared.MeshProxyPatchItemSpecDefaultMatch{
					Name:   name3,
					Origin: origin2,
					Tags:   tags,
				}
			}
			operation2 := shared.MeshProxyPatchItemSpecOperation(appendModificationsItem.Listener.Operation.ValueString())
			value5 := new(string)
			if !appendModificationsItem.Listener.Value.IsUnknown() && !appendModificationsItem.Listener.Value.IsNull() {
				*value5 = appendModificationsItem.Listener.Value.ValueString()
			} else {
				value5 = nil
			}
			listener = &shared.Listener{
				JSONPatches: jsonPatches2,
				Match:       match2,
				Operation:   operation2,
				Value:       value5,
			}
		}
		var networkFilter *shared.NetworkFilter
		if appendModificationsItem.NetworkFilter != nil {
			var jsonPatches3 []shared.MeshProxyPatchItemSpecDefaultJSONPatches = []shared.MeshProxyPatchItemSpecDefaultJSONPatches{}
			for _, jsonPatchesItem3 := range appendModificationsItem.NetworkFilter.JSONPatches {
				from3 := new(string)
				if !jsonPatchesItem3.From.IsUnknown() && !jsonPatchesItem3.From.IsNull() {
					*from3 = jsonPatchesItem3.From.ValueString()
				} else {
					from3 = nil
				}
				op3 := shared.MeshProxyPatchItemSpecDefaultOp(jsonPatchesItem3.Op.ValueString())
				var path3 string
				path3 = jsonPatchesItem3.Path.ValueString()

				var value6 interface{}
				if !jsonPatchesItem3.Value.IsUnknown() && !jsonPatchesItem3.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem3.Value.ValueString()), &value6)
				}
				jsonPatches3 = append(jsonPatches3, shared.MeshProxyPatchItemSpecDefaultJSONPatches{
					From:  from3,
					Op:    op3,
					Path:  path3,
					Value: value6,
				})
			}
			var match3 *shared.MeshProxyPatchItemSpecDefaultAppendModificationsMatch
			if appendModificationsItem.NetworkFilter.Match != nil {
				listenerName1 := new(string)
				if !appendModificationsItem.NetworkFilter.Match.ListenerName.IsUnknown() && !appendModificationsItem.NetworkFilter.Match.ListenerName.IsNull() {
					*listenerName1 = appendModificationsItem.NetworkFilter.Match.ListenerName.ValueString()
				} else {
					listenerName1 = nil
				}
				listenerTags1 := make(map[string]string)
				for listenerTagsKey1, listenerTagsValue1 := range appendModificationsItem.NetworkFilter.Match.ListenerTags {
					var listenerTagsInst1 string
					listenerTagsInst1 = listenerTagsValue1.ValueString()

					listenerTags1[listenerTagsKey1] = listenerTagsInst1
				}
				name4 := new(string)
				if !appendModificationsItem.NetworkFilter.Match.Name.IsUnknown() && !appendModificationsItem.NetworkFilter.Match.Name.IsNull() {
					*name4 = appendModificationsItem.NetworkFilter.Match.Name.ValueString()
				} else {
					name4 = nil
				}
				origin3 := new(string)
				if !appendModificationsItem.NetworkFilter.Match.Origin.IsUnknown() && !appendModificationsItem.NetworkFilter.Match.Origin.IsNull() {
					*origin3 = appendModificationsItem.NetworkFilter.Match.Origin.ValueString()
				} else {
					origin3 = nil
				}
				match3 = &shared.MeshProxyPatchItemSpecDefaultAppendModificationsMatch{
					ListenerName: listenerName1,
					ListenerTags: listenerTags1,
					Name:         name4,
					Origin:       origin3,
				}
			}
			operation3 := shared.MeshProxyPatchItemSpecDefaultOperation(appendModificationsItem.NetworkFilter.Operation.ValueString())
			value7 := new(string)
			if !appendModificationsItem.NetworkFilter.Value.IsUnknown() && !appendModificationsItem.NetworkFilter.Value.IsNull() {
				*value7 = appendModificationsItem.NetworkFilter.Value.ValueString()
			} else {
				value7 = nil
			}
			networkFilter = &shared.NetworkFilter{
				JSONPatches: jsonPatches3,
				Match:       match3,
				Operation:   operation3,
				Value:       value7,
			}
		}
		var virtualHost *shared.VirtualHost
		if appendModificationsItem.VirtualHost != nil {
			var jsonPatches4 []shared.MeshProxyPatchItemSpecDefaultAppendModificationsJSONPatches = []shared.MeshProxyPatchItemSpecDefaultAppendModificationsJSONPatches{}
			for _, jsonPatchesItem4 := range appendModificationsItem.VirtualHost.JSONPatches {
				from4 := new(string)
				if !jsonPatchesItem4.From.IsUnknown() && !jsonPatchesItem4.From.IsNull() {
					*from4 = jsonPatchesItem4.From.ValueString()
				} else {
					from4 = nil
				}
				op4 := shared.MeshProxyPatchItemSpecDefaultAppendModificationsOp(jsonPatchesItem4.Op.ValueString())
				var path4 string
				path4 = jsonPatchesItem4.Path.ValueString()

				var value8 interface{}
				if !jsonPatchesItem4.Value.IsUnknown() && !jsonPatchesItem4.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem4.Value.ValueString()), &value8)
				}
				jsonPatches4 = append(jsonPatches4, shared.MeshProxyPatchItemSpecDefaultAppendModificationsJSONPatches{
					From:  from4,
					Op:    op4,
					Path:  path4,
					Value: value8,
				})
			}
			name5 := new(string)
			if !appendModificationsItem.VirtualHost.Match.Name.IsUnknown() && !appendModificationsItem.VirtualHost.Match.Name.IsNull() {
				*name5 = appendModificationsItem.VirtualHost.Match.Name.ValueString()
			} else {
				name5 = nil
			}
			origin4 := new(string)
			if !appendModificationsItem.VirtualHost.Match.Origin.IsUnknown() && !appendModificationsItem.VirtualHost.Match.Origin.IsNull() {
				*origin4 = appendModificationsItem.VirtualHost.Match.Origin.ValueString()
			} else {
				origin4 = nil
			}
			routeConfigurationName := new(string)
			if !appendModificationsItem.VirtualHost.Match.RouteConfigurationName.IsUnknown() && !appendModificationsItem.VirtualHost.Match.RouteConfigurationName.IsNull() {
				*routeConfigurationName = appendModificationsItem.VirtualHost.Match.RouteConfigurationName.ValueString()
			} else {
				routeConfigurationName = nil
			}
			match4 := shared.MeshProxyPatchItemSpecDefaultAppendModificationsVirtualHostMatch{
				Name:                   name5,
				Origin:                 origin4,
				RouteConfigurationName: routeConfigurationName,
			}
			operation4 := shared.MeshProxyPatchItemSpecDefaultAppendModificationsOperation(appendModificationsItem.VirtualHost.Operation.ValueString())
			value9 := new(string)
			if !appendModificationsItem.VirtualHost.Value.IsUnknown() && !appendModificationsItem.VirtualHost.Value.IsNull() {
				*value9 = appendModificationsItem.VirtualHost.Value.ValueString()
			} else {
				value9 = nil
			}
			virtualHost = &shared.VirtualHost{
				JSONPatches: jsonPatches4,
				Match:       match4,
				Operation:   operation4,
				Value:       value9,
			}
		}
		appendModifications = append(appendModifications, shared.AppendModifications{
			Cluster:       cluster,
			HTTPFilter:    httpFilter,
			Listener:      listener,
			NetworkFilter: networkFilter,
			VirtualHost:   virtualHost,
		})
	}
	defaultVar := shared.MeshProxyPatchItemDefault{
		AppendModifications: appendModifications,
	}
	var targetRef *shared.MeshProxyPatchItemTargetRef
	if r.Spec.TargetRef != nil {
		kind := new(shared.MeshProxyPatchItemKind)
		if !r.Spec.TargetRef.Kind.IsUnknown() && !r.Spec.TargetRef.Kind.IsNull() {
			*kind = shared.MeshProxyPatchItemKind(r.Spec.TargetRef.Kind.ValueString())
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
		name6 := new(string)
		if !r.Spec.TargetRef.Name.IsUnknown() && !r.Spec.TargetRef.Name.IsNull() {
			*name6 = r.Spec.TargetRef.Name.ValueString()
		} else {
			name6 = nil
		}
		namespace := new(string)
		if !r.Spec.TargetRef.Namespace.IsUnknown() && !r.Spec.TargetRef.Namespace.IsNull() {
			*namespace = r.Spec.TargetRef.Namespace.ValueString()
		} else {
			namespace = nil
		}
		var proxyTypes []shared.MeshProxyPatchItemProxyTypes = []shared.MeshProxyPatchItemProxyTypes{}
		for _, proxyTypesItem := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshProxyPatchItemProxyTypes(proxyTypesItem.ValueString()))
		}
		sectionName := new(string)
		if !r.Spec.TargetRef.SectionName.IsUnknown() && !r.Spec.TargetRef.SectionName.IsNull() {
			*sectionName = r.Spec.TargetRef.SectionName.ValueString()
		} else {
			sectionName = nil
		}
		tags1 := make(map[string]string)
		for tagsKey1, tagsValue1 := range r.Spec.TargetRef.Tags {
			var tagsInst1 string
			tagsInst1 = tagsValue1.ValueString()

			tags1[tagsKey1] = tagsInst1
		}
		targetRef = &shared.MeshProxyPatchItemTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name6,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags1,
		}
	}
	spec := shared.MeshProxyPatchItemSpec{
		Default:   defaultVar,
		TargetRef: targetRef,
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
	out := shared.MeshProxyPatchItem{
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

func (r *MeshProxyPatchResourceModel) RefreshFromSharedMeshProxyPatchCreateOrUpdateSuccessResponse(resp *shared.MeshProxyPatchCreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = []types.String{}
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshProxyPatchResourceModel) RefreshFromSharedMeshProxyPatchItem(resp *shared.MeshProxyPatchItem) {
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
		r.Spec.Default.AppendModifications = []tfTypes.AppendModifications{}
		if len(r.Spec.Default.AppendModifications) > len(resp.Spec.Default.AppendModifications) {
			r.Spec.Default.AppendModifications = r.Spec.Default.AppendModifications[:len(resp.Spec.Default.AppendModifications)]
		}
		for appendModificationsCount, appendModificationsItem := range resp.Spec.Default.AppendModifications {
			var appendModifications1 tfTypes.AppendModifications
			if appendModificationsItem.Cluster == nil {
				appendModifications1.Cluster = nil
			} else {
				appendModifications1.Cluster = &tfTypes.Cluster{}
				appendModifications1.Cluster.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount, jsonPatchesItem := range appendModificationsItem.Cluster.JSONPatches {
					var jsonPatches1 tfTypes.JSONPatches
					jsonPatches1.From = types.StringPointerValue(jsonPatchesItem.From)
					jsonPatches1.Op = types.StringValue(string(jsonPatchesItem.Op))
					jsonPatches1.Path = types.StringValue(jsonPatchesItem.Path)
					if jsonPatchesItem.Value == nil {
						jsonPatches1.Value = types.StringNull()
					} else {
						valueResult, _ := json.Marshal(jsonPatchesItem.Value)
						jsonPatches1.Value = types.StringValue(string(valueResult))
					}
					if jsonPatchesCount+1 > len(appendModifications1.Cluster.JSONPatches) {
						appendModifications1.Cluster.JSONPatches = append(appendModifications1.Cluster.JSONPatches, jsonPatches1)
					} else {
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].From = jsonPatches1.From
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Op = jsonPatches1.Op
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Path = jsonPatches1.Path
						appendModifications1.Cluster.JSONPatches[jsonPatchesCount].Value = jsonPatches1.Value
					}
				}
				if appendModificationsItem.Cluster.Match == nil {
					appendModifications1.Cluster.Match = nil
				} else {
					appendModifications1.Cluster.Match = &tfTypes.MeshProxyPatchItemMatch{}
					appendModifications1.Cluster.Match.Name = types.StringPointerValue(appendModificationsItem.Cluster.Match.Name)
					appendModifications1.Cluster.Match.Origin = types.StringPointerValue(appendModificationsItem.Cluster.Match.Origin)
				}
				appendModifications1.Cluster.Operation = types.StringValue(string(appendModificationsItem.Cluster.Operation))
				appendModifications1.Cluster.Value = types.StringPointerValue(appendModificationsItem.Cluster.Value)
			}
			if appendModificationsItem.HTTPFilter == nil {
				appendModifications1.HTTPFilter = nil
			} else {
				appendModifications1.HTTPFilter = &tfTypes.HTTPFilter{}
				appendModifications1.HTTPFilter.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount1, jsonPatchesItem1 := range appendModificationsItem.HTTPFilter.JSONPatches {
					var jsonPatches3 tfTypes.JSONPatches
					jsonPatches3.From = types.StringPointerValue(jsonPatchesItem1.From)
					jsonPatches3.Op = types.StringValue(string(jsonPatchesItem1.Op))
					jsonPatches3.Path = types.StringValue(jsonPatchesItem1.Path)
					if jsonPatchesItem1.Value == nil {
						jsonPatches3.Value = types.StringNull()
					} else {
						valueResult1, _ := json.Marshal(jsonPatchesItem1.Value)
						jsonPatches3.Value = types.StringValue(string(valueResult1))
					}
					if jsonPatchesCount1+1 > len(appendModifications1.HTTPFilter.JSONPatches) {
						appendModifications1.HTTPFilter.JSONPatches = append(appendModifications1.HTTPFilter.JSONPatches, jsonPatches3)
					} else {
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].From = jsonPatches3.From
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Op = jsonPatches3.Op
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Path = jsonPatches3.Path
						appendModifications1.HTTPFilter.JSONPatches[jsonPatchesCount1].Value = jsonPatches3.Value
					}
				}
				if appendModificationsItem.HTTPFilter.Match == nil {
					appendModifications1.HTTPFilter.Match = nil
				} else {
					appendModifications1.HTTPFilter.Match = &tfTypes.MeshProxyPatchItemSpecMatch{}
					appendModifications1.HTTPFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.ListenerName)
					if len(appendModificationsItem.HTTPFilter.Match.ListenerTags) > 0 {
						appendModifications1.HTTPFilter.Match.ListenerTags = make(map[string]types.String)
						for key1, value4 := range appendModificationsItem.HTTPFilter.Match.ListenerTags {
							appendModifications1.HTTPFilter.Match.ListenerTags[key1] = types.StringValue(value4)
						}
					}
					appendModifications1.HTTPFilter.Match.Name = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Name)
					appendModifications1.HTTPFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Origin)
				}
				appendModifications1.HTTPFilter.Operation = types.StringValue(string(appendModificationsItem.HTTPFilter.Operation))
				appendModifications1.HTTPFilter.Value = types.StringPointerValue(appendModificationsItem.HTTPFilter.Value)
			}
			if appendModificationsItem.Listener == nil {
				appendModifications1.Listener = nil
			} else {
				appendModifications1.Listener = &tfTypes.Listener{}
				appendModifications1.Listener.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount2, jsonPatchesItem2 := range appendModificationsItem.Listener.JSONPatches {
					var jsonPatches5 tfTypes.JSONPatches
					jsonPatches5.From = types.StringPointerValue(jsonPatchesItem2.From)
					jsonPatches5.Op = types.StringValue(string(jsonPatchesItem2.Op))
					jsonPatches5.Path = types.StringValue(jsonPatchesItem2.Path)
					if jsonPatchesItem2.Value == nil {
						jsonPatches5.Value = types.StringNull()
					} else {
						valueResult2, _ := json.Marshal(jsonPatchesItem2.Value)
						jsonPatches5.Value = types.StringValue(string(valueResult2))
					}
					if jsonPatchesCount2+1 > len(appendModifications1.Listener.JSONPatches) {
						appendModifications1.Listener.JSONPatches = append(appendModifications1.Listener.JSONPatches, jsonPatches5)
					} else {
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].From = jsonPatches5.From
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Op = jsonPatches5.Op
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Path = jsonPatches5.Path
						appendModifications1.Listener.JSONPatches[jsonPatchesCount2].Value = jsonPatches5.Value
					}
				}
				if appendModificationsItem.Listener.Match == nil {
					appendModifications1.Listener.Match = nil
				} else {
					appendModifications1.Listener.Match = &tfTypes.MeshProxyPatchItemSpecDefaultMatch{}
					appendModifications1.Listener.Match.Name = types.StringPointerValue(appendModificationsItem.Listener.Match.Name)
					appendModifications1.Listener.Match.Origin = types.StringPointerValue(appendModificationsItem.Listener.Match.Origin)
					if len(appendModificationsItem.Listener.Match.Tags) > 0 {
						appendModifications1.Listener.Match.Tags = make(map[string]types.String)
						for key2, value7 := range appendModificationsItem.Listener.Match.Tags {
							appendModifications1.Listener.Match.Tags[key2] = types.StringValue(value7)
						}
					}
				}
				appendModifications1.Listener.Operation = types.StringValue(string(appendModificationsItem.Listener.Operation))
				appendModifications1.Listener.Value = types.StringPointerValue(appendModificationsItem.Listener.Value)
			}
			if appendModificationsItem.NetworkFilter == nil {
				appendModifications1.NetworkFilter = nil
			} else {
				appendModifications1.NetworkFilter = &tfTypes.HTTPFilter{}
				appendModifications1.NetworkFilter.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount3, jsonPatchesItem3 := range appendModificationsItem.NetworkFilter.JSONPatches {
					var jsonPatches7 tfTypes.JSONPatches
					jsonPatches7.From = types.StringPointerValue(jsonPatchesItem3.From)
					jsonPatches7.Op = types.StringValue(string(jsonPatchesItem3.Op))
					jsonPatches7.Path = types.StringValue(jsonPatchesItem3.Path)
					if jsonPatchesItem3.Value == nil {
						jsonPatches7.Value = types.StringNull()
					} else {
						valueResult3, _ := json.Marshal(jsonPatchesItem3.Value)
						jsonPatches7.Value = types.StringValue(string(valueResult3))
					}
					if jsonPatchesCount3+1 > len(appendModifications1.NetworkFilter.JSONPatches) {
						appendModifications1.NetworkFilter.JSONPatches = append(appendModifications1.NetworkFilter.JSONPatches, jsonPatches7)
					} else {
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].From = jsonPatches7.From
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Op = jsonPatches7.Op
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Path = jsonPatches7.Path
						appendModifications1.NetworkFilter.JSONPatches[jsonPatchesCount3].Value = jsonPatches7.Value
					}
				}
				if appendModificationsItem.NetworkFilter.Match == nil {
					appendModifications1.NetworkFilter.Match = nil
				} else {
					appendModifications1.NetworkFilter.Match = &tfTypes.MeshProxyPatchItemSpecMatch{}
					appendModifications1.NetworkFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.ListenerName)
					if len(appendModificationsItem.NetworkFilter.Match.ListenerTags) > 0 {
						appendModifications1.NetworkFilter.Match.ListenerTags = make(map[string]types.String)
						for key3, value10 := range appendModificationsItem.NetworkFilter.Match.ListenerTags {
							appendModifications1.NetworkFilter.Match.ListenerTags[key3] = types.StringValue(value10)
						}
					}
					appendModifications1.NetworkFilter.Match.Name = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Name)
					appendModifications1.NetworkFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Origin)
				}
				appendModifications1.NetworkFilter.Operation = types.StringValue(string(appendModificationsItem.NetworkFilter.Operation))
				appendModifications1.NetworkFilter.Value = types.StringPointerValue(appendModificationsItem.NetworkFilter.Value)
			}
			if appendModificationsItem.VirtualHost == nil {
				appendModifications1.VirtualHost = nil
			} else {
				appendModifications1.VirtualHost = &tfTypes.VirtualHost{}
				appendModifications1.VirtualHost.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount4, jsonPatchesItem4 := range appendModificationsItem.VirtualHost.JSONPatches {
					var jsonPatches9 tfTypes.JSONPatches
					jsonPatches9.From = types.StringPointerValue(jsonPatchesItem4.From)
					jsonPatches9.Op = types.StringValue(string(jsonPatchesItem4.Op))
					jsonPatches9.Path = types.StringValue(jsonPatchesItem4.Path)
					if jsonPatchesItem4.Value == nil {
						jsonPatches9.Value = types.StringNull()
					} else {
						valueResult4, _ := json.Marshal(jsonPatchesItem4.Value)
						jsonPatches9.Value = types.StringValue(string(valueResult4))
					}
					if jsonPatchesCount4+1 > len(appendModifications1.VirtualHost.JSONPatches) {
						appendModifications1.VirtualHost.JSONPatches = append(appendModifications1.VirtualHost.JSONPatches, jsonPatches9)
					} else {
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].From = jsonPatches9.From
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Op = jsonPatches9.Op
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Path = jsonPatches9.Path
						appendModifications1.VirtualHost.JSONPatches[jsonPatchesCount4].Value = jsonPatches9.Value
					}
				}
				appendModifications1.VirtualHost.Match.Name = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Name)
				appendModifications1.VirtualHost.Match.Origin = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Origin)
				appendModifications1.VirtualHost.Match.RouteConfigurationName = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.RouteConfigurationName)
				appendModifications1.VirtualHost.Operation = types.StringValue(string(appendModificationsItem.VirtualHost.Operation))
				appendModifications1.VirtualHost.Value = types.StringPointerValue(appendModificationsItem.VirtualHost.Value)
			}
			if appendModificationsCount+1 > len(r.Spec.Default.AppendModifications) {
				r.Spec.Default.AppendModifications = append(r.Spec.Default.AppendModifications, appendModifications1)
			} else {
				r.Spec.Default.AppendModifications[appendModificationsCount].Cluster = appendModifications1.Cluster
				r.Spec.Default.AppendModifications[appendModificationsCount].HTTPFilter = appendModifications1.HTTPFilter
				r.Spec.Default.AppendModifications[appendModificationsCount].Listener = appendModifications1.Listener
				r.Spec.Default.AppendModifications[appendModificationsCount].NetworkFilter = appendModifications1.NetworkFilter
				r.Spec.Default.AppendModifications[appendModificationsCount].VirtualHost = appendModifications1.VirtualHost
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
				for key4, value14 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key4] = types.StringValue(value14)
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
				for key5, value15 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key5] = types.StringValue(value15)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
