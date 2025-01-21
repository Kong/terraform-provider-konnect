// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshTrafficPermissionDataSourceModel) RefreshFromSharedMeshTrafficPermissionItem(resp *shared.MeshTrafficPermissionItem) {
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
		r.Spec.From = []tfTypes.MeshTrafficPermissionItemFrom{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from1 tfTypes.MeshTrafficPermissionItemFrom
			if fromItem.Default == nil {
				from1.Default = nil
			} else {
				from1.Default = &tfTypes.MeshTrafficPermissionItemDefault{}
				if fromItem.Default.Action != nil {
					from1.Default.Action = types.StringValue(string(*fromItem.Default.Action))
				} else {
					from1.Default.Action = types.StringNull()
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
				for key3, value3 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key3] = types.StringValue(value3)
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
				for key4, value4 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key4] = types.StringValue(value4)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
