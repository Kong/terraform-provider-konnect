// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *MeshOPAResourceModel) ToSharedMeshOPAItem() *shared.MeshOPAItem {
	typeVar := shared.MeshOPAItemType(r.Type.ValueString())
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
	var defaultVar *shared.MeshOPAItemDefault
	if r.Spec.Default != nil {
		var agentConfig *shared.AgentConfig
		if r.Spec.Default.AgentConfig != nil {
			inline := new(string)
			if !r.Spec.Default.AgentConfig.Inline.IsUnknown() && !r.Spec.Default.AgentConfig.Inline.IsNull() {
				*inline = r.Spec.Default.AgentConfig.Inline.ValueString()
			} else {
				inline = nil
			}
			inlineString := new(string)
			if !r.Spec.Default.AgentConfig.InlineString.IsUnknown() && !r.Spec.Default.AgentConfig.InlineString.IsNull() {
				*inlineString = r.Spec.Default.AgentConfig.InlineString.ValueString()
			} else {
				inlineString = nil
			}
			secret := new(string)
			if !r.Spec.Default.AgentConfig.Secret.IsUnknown() && !r.Spec.Default.AgentConfig.Secret.IsNull() {
				*secret = r.Spec.Default.AgentConfig.Secret.ValueString()
			} else {
				secret = nil
			}
			agentConfig = &shared.AgentConfig{
				Inline:       inline,
				InlineString: inlineString,
				Secret:       secret,
			}
		}
		var appendPolicies []shared.AppendPolicies = []shared.AppendPolicies{}
		for _, appendPoliciesItem := range r.Spec.Default.AppendPolicies {
			ignoreDecision := new(bool)
			if !appendPoliciesItem.IgnoreDecision.IsUnknown() && !appendPoliciesItem.IgnoreDecision.IsNull() {
				*ignoreDecision = appendPoliciesItem.IgnoreDecision.ValueBool()
			} else {
				ignoreDecision = nil
			}
			inline1 := new(string)
			if !appendPoliciesItem.Rego.Inline.IsUnknown() && !appendPoliciesItem.Rego.Inline.IsNull() {
				*inline1 = appendPoliciesItem.Rego.Inline.ValueString()
			} else {
				inline1 = nil
			}
			inlineString1 := new(string)
			if !appendPoliciesItem.Rego.InlineString.IsUnknown() && !appendPoliciesItem.Rego.InlineString.IsNull() {
				*inlineString1 = appendPoliciesItem.Rego.InlineString.ValueString()
			} else {
				inlineString1 = nil
			}
			secret1 := new(string)
			if !appendPoliciesItem.Rego.Secret.IsUnknown() && !appendPoliciesItem.Rego.Secret.IsNull() {
				*secret1 = appendPoliciesItem.Rego.Secret.ValueString()
			} else {
				secret1 = nil
			}
			rego := shared.Rego{
				Inline:       inline1,
				InlineString: inlineString1,
				Secret:       secret1,
			}
			appendPolicies = append(appendPolicies, shared.AppendPolicies{
				IgnoreDecision: ignoreDecision,
				Rego:           rego,
			})
		}
		var authConfig *shared.AuthConfig
		if r.Spec.Default.AuthConfig != nil {
			onAgentFailure := new(shared.OnAgentFailure)
			if !r.Spec.Default.AuthConfig.OnAgentFailure.IsUnknown() && !r.Spec.Default.AuthConfig.OnAgentFailure.IsNull() {
				*onAgentFailure = shared.OnAgentFailure(r.Spec.Default.AuthConfig.OnAgentFailure.ValueString())
			} else {
				onAgentFailure = nil
			}
			var requestBody *shared.RequestBody
			if r.Spec.Default.AuthConfig.RequestBody != nil {
				maxSize := new(int)
				if !r.Spec.Default.AuthConfig.RequestBody.MaxSize.IsUnknown() && !r.Spec.Default.AuthConfig.RequestBody.MaxSize.IsNull() {
					*maxSize = int(r.Spec.Default.AuthConfig.RequestBody.MaxSize.ValueInt64())
				} else {
					maxSize = nil
				}
				sendRawBody := new(bool)
				if !r.Spec.Default.AuthConfig.RequestBody.SendRawBody.IsUnknown() && !r.Spec.Default.AuthConfig.RequestBody.SendRawBody.IsNull() {
					*sendRawBody = r.Spec.Default.AuthConfig.RequestBody.SendRawBody.ValueBool()
				} else {
					sendRawBody = nil
				}
				requestBody = &shared.RequestBody{
					MaxSize:     maxSize,
					SendRawBody: sendRawBody,
				}
			}
			statusOnError := new(int)
			if !r.Spec.Default.AuthConfig.StatusOnError.IsUnknown() && !r.Spec.Default.AuthConfig.StatusOnError.IsNull() {
				*statusOnError = int(r.Spec.Default.AuthConfig.StatusOnError.ValueInt64())
			} else {
				statusOnError = nil
			}
			timeout := new(string)
			if !r.Spec.Default.AuthConfig.Timeout.IsUnknown() && !r.Spec.Default.AuthConfig.Timeout.IsNull() {
				*timeout = r.Spec.Default.AuthConfig.Timeout.ValueString()
			} else {
				timeout = nil
			}
			authConfig = &shared.AuthConfig{
				OnAgentFailure: onAgentFailure,
				RequestBody:    requestBody,
				StatusOnError:  statusOnError,
				Timeout:        timeout,
			}
		}
		defaultVar = &shared.MeshOPAItemDefault{
			AgentConfig:    agentConfig,
			AppendPolicies: appendPolicies,
			AuthConfig:     authConfig,
		}
	}
	var targetRef *shared.MeshOPAItemTargetRef
	if r.Spec.TargetRef != nil {
		kind := new(shared.MeshOPAItemKind)
		if !r.Spec.TargetRef.Kind.IsUnknown() && !r.Spec.TargetRef.Kind.IsNull() {
			*kind = shared.MeshOPAItemKind(r.Spec.TargetRef.Kind.ValueString())
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
		var proxyTypes []shared.MeshOPAItemProxyTypes = []shared.MeshOPAItemProxyTypes{}
		for _, proxyTypesItem := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshOPAItemProxyTypes(proxyTypesItem.ValueString()))
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
		targetRef = &shared.MeshOPAItemTargetRef{
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
	spec := shared.MeshOPAItemSpec{
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
	out := shared.MeshOPAItem{
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

func (r *MeshOPAResourceModel) RefreshFromSharedMeshOPACreateOrUpdateSuccessResponse(resp *shared.MeshOPACreateOrUpdateSuccessResponse) {
	if resp != nil {
		r.Warnings = []types.String{}
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}
}

func (r *MeshOPAResourceModel) RefreshFromSharedMeshOPAItem(resp *shared.MeshOPAItem) {
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
		if resp.Spec.Default == nil {
			r.Spec.Default = nil
		} else {
			r.Spec.Default = &tfTypes.MeshOPAItemDefault{}
			if resp.Spec.Default.AgentConfig == nil {
				r.Spec.Default.AgentConfig = nil
			} else {
				r.Spec.Default.AgentConfig = &tfTypes.CaCert{}
				r.Spec.Default.AgentConfig.Inline = types.StringPointerValue(resp.Spec.Default.AgentConfig.Inline)
				r.Spec.Default.AgentConfig.InlineString = types.StringPointerValue(resp.Spec.Default.AgentConfig.InlineString)
				r.Spec.Default.AgentConfig.Secret = types.StringPointerValue(resp.Spec.Default.AgentConfig.Secret)
			}
			r.Spec.Default.AppendPolicies = []tfTypes.AppendPolicies{}
			if len(r.Spec.Default.AppendPolicies) > len(resp.Spec.Default.AppendPolicies) {
				r.Spec.Default.AppendPolicies = r.Spec.Default.AppendPolicies[:len(resp.Spec.Default.AppendPolicies)]
			}
			for appendPoliciesCount, appendPoliciesItem := range resp.Spec.Default.AppendPolicies {
				var appendPolicies1 tfTypes.AppendPolicies
				appendPolicies1.IgnoreDecision = types.BoolPointerValue(appendPoliciesItem.IgnoreDecision)
				appendPolicies1.Rego.Inline = types.StringPointerValue(appendPoliciesItem.Rego.Inline)
				appendPolicies1.Rego.InlineString = types.StringPointerValue(appendPoliciesItem.Rego.InlineString)
				appendPolicies1.Rego.Secret = types.StringPointerValue(appendPoliciesItem.Rego.Secret)
				if appendPoliciesCount+1 > len(r.Spec.Default.AppendPolicies) {
					r.Spec.Default.AppendPolicies = append(r.Spec.Default.AppendPolicies, appendPolicies1)
				} else {
					r.Spec.Default.AppendPolicies[appendPoliciesCount].IgnoreDecision = appendPolicies1.IgnoreDecision
					r.Spec.Default.AppendPolicies[appendPoliciesCount].Rego = appendPolicies1.Rego
				}
			}
			if resp.Spec.Default.AuthConfig == nil {
				r.Spec.Default.AuthConfig = nil
			} else {
				r.Spec.Default.AuthConfig = &tfTypes.AuthConfig{}
				if resp.Spec.Default.AuthConfig.OnAgentFailure != nil {
					r.Spec.Default.AuthConfig.OnAgentFailure = types.StringValue(string(*resp.Spec.Default.AuthConfig.OnAgentFailure))
				} else {
					r.Spec.Default.AuthConfig.OnAgentFailure = types.StringNull()
				}
				if resp.Spec.Default.AuthConfig.RequestBody == nil {
					r.Spec.Default.AuthConfig.RequestBody = nil
				} else {
					r.Spec.Default.AuthConfig.RequestBody = &tfTypes.RequestBody{}
					if resp.Spec.Default.AuthConfig.RequestBody.MaxSize != nil {
						r.Spec.Default.AuthConfig.RequestBody.MaxSize = types.Int64Value(int64(*resp.Spec.Default.AuthConfig.RequestBody.MaxSize))
					} else {
						r.Spec.Default.AuthConfig.RequestBody.MaxSize = types.Int64Null()
					}
					r.Spec.Default.AuthConfig.RequestBody.SendRawBody = types.BoolPointerValue(resp.Spec.Default.AuthConfig.RequestBody.SendRawBody)
				}
				if resp.Spec.Default.AuthConfig.StatusOnError != nil {
					r.Spec.Default.AuthConfig.StatusOnError = types.Int64Value(int64(*resp.Spec.Default.AuthConfig.StatusOnError))
				} else {
					r.Spec.Default.AuthConfig.StatusOnError = types.Int64Null()
				}
				r.Spec.Default.AuthConfig.Timeout = types.StringPointerValue(resp.Spec.Default.AuthConfig.Timeout)
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
		r.Type = types.StringValue(string(resp.Type))
	}
}