// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginRequestTerminationResourceModel) ToSharedRequestTerminationPluginInput() *shared.RequestTerminationPluginInput {
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
	var ordering *shared.RequestTerminationPluginOrdering
	if r.Ordering != nil {
		var after *shared.RequestTerminationPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.RequestTerminationPluginAfter{
				Access: access,
			}
		}
		var before *shared.RequestTerminationPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.RequestTerminationPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.RequestTerminationPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	body := new(string)
	if !r.Config.Body.IsUnknown() && !r.Config.Body.IsNull() {
		*body = r.Config.Body.ValueString()
	} else {
		body = nil
	}
	contentType := new(string)
	if !r.Config.ContentType.IsUnknown() && !r.Config.ContentType.IsNull() {
		*contentType = r.Config.ContentType.ValueString()
	} else {
		contentType = nil
	}
	echo := new(bool)
	if !r.Config.Echo.IsUnknown() && !r.Config.Echo.IsNull() {
		*echo = r.Config.Echo.ValueBool()
	} else {
		echo = nil
	}
	message := new(string)
	if !r.Config.Message.IsUnknown() && !r.Config.Message.IsNull() {
		*message = r.Config.Message.ValueString()
	} else {
		message = nil
	}
	statusCode := new(int64)
	if !r.Config.StatusCode.IsUnknown() && !r.Config.StatusCode.IsNull() {
		*statusCode = r.Config.StatusCode.ValueInt64()
	} else {
		statusCode = nil
	}
	trigger := new(string)
	if !r.Config.Trigger.IsUnknown() && !r.Config.Trigger.IsNull() {
		*trigger = r.Config.Trigger.ValueString()
	} else {
		trigger = nil
	}
	config := shared.RequestTerminationPluginConfig{
		Body:        body,
		ContentType: contentType,
		Echo:        echo,
		Message:     message,
		StatusCode:  statusCode,
		Trigger:     trigger,
	}
	var consumer *shared.RequestTerminationPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.RequestTerminationPluginConsumer{
			ID: id1,
		}
	}
	var consumerGroup *shared.RequestTerminationPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id2 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id2 = r.ConsumerGroup.ID.ValueString()
		} else {
			id2 = nil
		}
		consumerGroup = &shared.RequestTerminationPluginConsumerGroup{
			ID: id2,
		}
	}
	var protocols []shared.RequestTerminationPluginProtocols = []shared.RequestTerminationPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.RequestTerminationPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.RequestTerminationPluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.RequestTerminationPluginRoute{
			ID: id3,
		}
	}
	var service *shared.RequestTerminationPluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.RequestTerminationPluginService{
			ID: id4,
		}
	}
	out := shared.RequestTerminationPluginInput{
		Enabled:       enabled,
		ID:            id,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Tags:          tags,
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginRequestTerminationResourceModel) RefreshFromSharedRequestTerminationPlugin(resp *shared.RequestTerminationPlugin) {
	if resp != nil {
		r.Config.Body = types.StringPointerValue(resp.Config.Body)
		r.Config.ContentType = types.StringPointerValue(resp.Config.ContentType)
		r.Config.Echo = types.BoolPointerValue(resp.Config.Echo)
		r.Config.Message = types.StringPointerValue(resp.Config.Message)
		r.Config.StatusCode = types.Int64PointerValue(resp.Config.StatusCode)
		r.Config.Trigger = types.StringPointerValue(resp.Config.Trigger)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLWithoutParentsConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
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
}
