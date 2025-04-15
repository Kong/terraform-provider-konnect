// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginLdapAuthAdvancedResourceModel) ToSharedLdapAuthAdvancedPlugin() *shared.LdapAuthAdvancedPlugin {
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.LdapAuthAdvancedPluginOrdering
	if r.Ordering != nil {
		var after *shared.LdapAuthAdvancedPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.LdapAuthAdvancedPluginAfter{
				Access: access,
			}
		}
		var before *shared.LdapAuthAdvancedPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.LdapAuthAdvancedPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.LdapAuthAdvancedPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.LdapAuthAdvancedPluginConfig
	if r.Config != nil {
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		attribute := new(string)
		if !r.Config.Attribute.IsUnknown() && !r.Config.Attribute.IsNull() {
			*attribute = r.Config.Attribute.ValueString()
		} else {
			attribute = nil
		}
		baseDn := new(string)
		if !r.Config.BaseDn.IsUnknown() && !r.Config.BaseDn.IsNull() {
			*baseDn = r.Config.BaseDn.ValueString()
		} else {
			baseDn = nil
		}
		bindDn := new(string)
		if !r.Config.BindDn.IsUnknown() && !r.Config.BindDn.IsNull() {
			*bindDn = r.Config.BindDn.ValueString()
		} else {
			bindDn = nil
		}
		cacheTTL := new(float64)
		if !r.Config.CacheTTL.IsUnknown() && !r.Config.CacheTTL.IsNull() {
			*cacheTTL = r.Config.CacheTTL.ValueFloat64()
		} else {
			cacheTTL = nil
		}
		var consumerBy []shared.LdapAuthAdvancedPluginConsumerBy = []shared.LdapAuthAdvancedPluginConsumerBy{}
		for _, consumerByItem := range r.Config.ConsumerBy {
			consumerBy = append(consumerBy, shared.LdapAuthAdvancedPluginConsumerBy(consumerByItem.ValueString()))
		}
		consumerOptional := new(bool)
		if !r.Config.ConsumerOptional.IsUnknown() && !r.Config.ConsumerOptional.IsNull() {
			*consumerOptional = r.Config.ConsumerOptional.ValueBool()
		} else {
			consumerOptional = nil
		}
		groupBaseDn := new(string)
		if !r.Config.GroupBaseDn.IsUnknown() && !r.Config.GroupBaseDn.IsNull() {
			*groupBaseDn = r.Config.GroupBaseDn.ValueString()
		} else {
			groupBaseDn = nil
		}
		groupMemberAttribute := new(string)
		if !r.Config.GroupMemberAttribute.IsUnknown() && !r.Config.GroupMemberAttribute.IsNull() {
			*groupMemberAttribute = r.Config.GroupMemberAttribute.ValueString()
		} else {
			groupMemberAttribute = nil
		}
		groupNameAttribute := new(string)
		if !r.Config.GroupNameAttribute.IsUnknown() && !r.Config.GroupNameAttribute.IsNull() {
			*groupNameAttribute = r.Config.GroupNameAttribute.ValueString()
		} else {
			groupNameAttribute = nil
		}
		var groupsRequired []string = []string{}
		for _, groupsRequiredItem := range r.Config.GroupsRequired {
			groupsRequired = append(groupsRequired, groupsRequiredItem.ValueString())
		}
		headerType := new(string)
		if !r.Config.HeaderType.IsUnknown() && !r.Config.HeaderType.IsNull() {
			*headerType = r.Config.HeaderType.ValueString()
		} else {
			headerType = nil
		}
		hideCredentials := new(bool)
		if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
			*hideCredentials = r.Config.HideCredentials.ValueBool()
		} else {
			hideCredentials = nil
		}
		keepalive := new(float64)
		if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
			*keepalive = r.Config.Keepalive.ValueFloat64()
		} else {
			keepalive = nil
		}
		ldapHost := new(string)
		if !r.Config.LdapHost.IsUnknown() && !r.Config.LdapHost.IsNull() {
			*ldapHost = r.Config.LdapHost.ValueString()
		} else {
			ldapHost = nil
		}
		ldapPassword := new(string)
		if !r.Config.LdapPassword.IsUnknown() && !r.Config.LdapPassword.IsNull() {
			*ldapPassword = r.Config.LdapPassword.ValueString()
		} else {
			ldapPassword = nil
		}
		ldapPort := new(float64)
		if !r.Config.LdapPort.IsUnknown() && !r.Config.LdapPort.IsNull() {
			*ldapPort = r.Config.LdapPort.ValueFloat64()
		} else {
			ldapPort = nil
		}
		ldaps := new(bool)
		if !r.Config.Ldaps.IsUnknown() && !r.Config.Ldaps.IsNull() {
			*ldaps = r.Config.Ldaps.ValueBool()
		} else {
			ldaps = nil
		}
		logSearchResults := new(bool)
		if !r.Config.LogSearchResults.IsUnknown() && !r.Config.LogSearchResults.IsNull() {
			*logSearchResults = r.Config.LogSearchResults.ValueBool()
		} else {
			logSearchResults = nil
		}
		realm := new(string)
		if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
			*realm = r.Config.Realm.ValueString()
		} else {
			realm = nil
		}
		startTLS := new(bool)
		if !r.Config.StartTLS.IsUnknown() && !r.Config.StartTLS.IsNull() {
			*startTLS = r.Config.StartTLS.ValueBool()
		} else {
			startTLS = nil
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout = r.Config.Timeout.ValueFloat64()
		} else {
			timeout = nil
		}
		verifyLdapHost := new(bool)
		if !r.Config.VerifyLdapHost.IsUnknown() && !r.Config.VerifyLdapHost.IsNull() {
			*verifyLdapHost = r.Config.VerifyLdapHost.ValueBool()
		} else {
			verifyLdapHost = nil
		}
		config = &shared.LdapAuthAdvancedPluginConfig{
			Anonymous:            anonymous,
			Attribute:            attribute,
			BaseDn:               baseDn,
			BindDn:               bindDn,
			CacheTTL:             cacheTTL,
			ConsumerBy:           consumerBy,
			ConsumerOptional:     consumerOptional,
			GroupBaseDn:          groupBaseDn,
			GroupMemberAttribute: groupMemberAttribute,
			GroupNameAttribute:   groupNameAttribute,
			GroupsRequired:       groupsRequired,
			HeaderType:           headerType,
			HideCredentials:      hideCredentials,
			Keepalive:            keepalive,
			LdapHost:             ldapHost,
			LdapPassword:         ldapPassword,
			LdapPort:             ldapPort,
			Ldaps:                ldaps,
			LogSearchResults:     logSearchResults,
			Realm:                realm,
			StartTLS:             startTLS,
			Timeout:              timeout,
			VerifyLdapHost:       verifyLdapHost,
		}
	}
	var protocols []shared.LdapAuthAdvancedPluginProtocols = []shared.LdapAuthAdvancedPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.LdapAuthAdvancedPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.LdapAuthAdvancedPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.LdapAuthAdvancedPluginRoute{
			ID: id1,
		}
	}
	var service *shared.LdapAuthAdvancedPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.LdapAuthAdvancedPluginService{
			ID: id2,
		}
	}
	out := shared.LdapAuthAdvancedPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginLdapAuthAdvancedResourceModel) RefreshFromSharedLdapAuthAdvancedPlugin(ctx context.Context, resp *shared.LdapAuthAdvancedPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.LdapAuthAdvancedPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.Attribute = types.StringPointerValue(resp.Config.Attribute)
			r.Config.BaseDn = types.StringPointerValue(resp.Config.BaseDn)
			r.Config.BindDn = types.StringPointerValue(resp.Config.BindDn)
			r.Config.CacheTTL = types.Float64PointerValue(resp.Config.CacheTTL)
			r.Config.ConsumerBy = make([]types.String, 0, len(resp.Config.ConsumerBy))
			for _, v := range resp.Config.ConsumerBy {
				r.Config.ConsumerBy = append(r.Config.ConsumerBy, types.StringValue(string(v)))
			}
			r.Config.ConsumerOptional = types.BoolPointerValue(resp.Config.ConsumerOptional)
			r.Config.GroupBaseDn = types.StringPointerValue(resp.Config.GroupBaseDn)
			r.Config.GroupMemberAttribute = types.StringPointerValue(resp.Config.GroupMemberAttribute)
			r.Config.GroupNameAttribute = types.StringPointerValue(resp.Config.GroupNameAttribute)
			r.Config.GroupsRequired = make([]types.String, 0, len(resp.Config.GroupsRequired))
			for _, v := range resp.Config.GroupsRequired {
				r.Config.GroupsRequired = append(r.Config.GroupsRequired, types.StringValue(v))
			}
			r.Config.HeaderType = types.StringPointerValue(resp.Config.HeaderType)
			r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
			r.Config.Keepalive = types.Float64PointerValue(resp.Config.Keepalive)
			r.Config.LdapHost = types.StringPointerValue(resp.Config.LdapHost)
			r.Config.LdapPassword = types.StringPointerValue(resp.Config.LdapPassword)
			r.Config.LdapPort = types.Float64PointerValue(resp.Config.LdapPort)
			r.Config.Ldaps = types.BoolPointerValue(resp.Config.Ldaps)
			r.Config.LogSearchResults = types.BoolPointerValue(resp.Config.LogSearchResults)
			r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
			r.Config.StartTLS = types.BoolPointerValue(resp.Config.StartTLS)
			r.Config.Timeout = types.Float64PointerValue(resp.Config.Timeout)
			r.Config.VerifyLdapHost = types.BoolPointerValue(resp.Config.VerifyLdapHost)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.ACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.ACLPluginAfter{}
				r.Ordering.After.Access = make([]types.String, 0, len(resp.Ordering.After.Access))
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = make([]types.String, 0, len(resp.Ordering.Before.Access))
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}

	return diags
}
