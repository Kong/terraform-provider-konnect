// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginXMLThreatProtectionResourceModel) ToSharedXMLThreatProtectionPluginInput() *shared.XMLThreatProtectionPluginInput {
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
	var ordering *shared.XMLThreatProtectionPluginOrdering
	if r.Ordering != nil {
		var after *shared.XMLThreatProtectionPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.XMLThreatProtectionPluginAfter{
				Access: access,
			}
		}
		var before *shared.XMLThreatProtectionPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.XMLThreatProtectionPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.XMLThreatProtectionPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	allowDtd := new(bool)
	if !r.Config.AllowDtd.IsUnknown() && !r.Config.AllowDtd.IsNull() {
		*allowDtd = r.Config.AllowDtd.ValueBool()
	} else {
		allowDtd = nil
	}
	var allowedContentTypes []string = []string{}
	for _, allowedContentTypesItem := range r.Config.AllowedContentTypes {
		allowedContentTypes = append(allowedContentTypes, allowedContentTypesItem.ValueString())
	}
	attribute := new(int64)
	if !r.Config.Attribute.IsUnknown() && !r.Config.Attribute.IsNull() {
		*attribute = r.Config.Attribute.ValueInt64()
	} else {
		attribute = nil
	}
	blaMaxAmplification := new(float64)
	if !r.Config.BlaMaxAmplification.IsUnknown() && !r.Config.BlaMaxAmplification.IsNull() {
		*blaMaxAmplification, _ = r.Config.BlaMaxAmplification.ValueBigFloat().Float64()
	} else {
		blaMaxAmplification = nil
	}
	blaThreshold := new(int64)
	if !r.Config.BlaThreshold.IsUnknown() && !r.Config.BlaThreshold.IsNull() {
		*blaThreshold = r.Config.BlaThreshold.ValueInt64()
	} else {
		blaThreshold = nil
	}
	buffer := new(int64)
	if !r.Config.Buffer.IsUnknown() && !r.Config.Buffer.IsNull() {
		*buffer = r.Config.Buffer.ValueInt64()
	} else {
		buffer = nil
	}
	var checkedContentTypes []string = []string{}
	for _, checkedContentTypesItem := range r.Config.CheckedContentTypes {
		checkedContentTypes = append(checkedContentTypes, checkedContentTypesItem.ValueString())
	}
	comment := new(int64)
	if !r.Config.Comment.IsUnknown() && !r.Config.Comment.IsNull() {
		*comment = r.Config.Comment.ValueInt64()
	} else {
		comment = nil
	}
	document := new(int64)
	if !r.Config.Document.IsUnknown() && !r.Config.Document.IsNull() {
		*document = r.Config.Document.ValueInt64()
	} else {
		document = nil
	}
	entity := new(int64)
	if !r.Config.Entity.IsUnknown() && !r.Config.Entity.IsNull() {
		*entity = r.Config.Entity.ValueInt64()
	} else {
		entity = nil
	}
	entityname := new(int64)
	if !r.Config.Entityname.IsUnknown() && !r.Config.Entityname.IsNull() {
		*entityname = r.Config.Entityname.ValueInt64()
	} else {
		entityname = nil
	}
	entityproperty := new(int64)
	if !r.Config.Entityproperty.IsUnknown() && !r.Config.Entityproperty.IsNull() {
		*entityproperty = r.Config.Entityproperty.ValueInt64()
	} else {
		entityproperty = nil
	}
	localname := new(int64)
	if !r.Config.Localname.IsUnknown() && !r.Config.Localname.IsNull() {
		*localname = r.Config.Localname.ValueInt64()
	} else {
		localname = nil
	}
	maxAttributes := new(int64)
	if !r.Config.MaxAttributes.IsUnknown() && !r.Config.MaxAttributes.IsNull() {
		*maxAttributes = r.Config.MaxAttributes.ValueInt64()
	} else {
		maxAttributes = nil
	}
	maxChildren := new(int64)
	if !r.Config.MaxChildren.IsUnknown() && !r.Config.MaxChildren.IsNull() {
		*maxChildren = r.Config.MaxChildren.ValueInt64()
	} else {
		maxChildren = nil
	}
	maxDepth := new(int64)
	if !r.Config.MaxDepth.IsUnknown() && !r.Config.MaxDepth.IsNull() {
		*maxDepth = r.Config.MaxDepth.ValueInt64()
	} else {
		maxDepth = nil
	}
	maxNamespaces := new(int64)
	if !r.Config.MaxNamespaces.IsUnknown() && !r.Config.MaxNamespaces.IsNull() {
		*maxNamespaces = r.Config.MaxNamespaces.ValueInt64()
	} else {
		maxNamespaces = nil
	}
	namespaceAware := new(bool)
	if !r.Config.NamespaceAware.IsUnknown() && !r.Config.NamespaceAware.IsNull() {
		*namespaceAware = r.Config.NamespaceAware.ValueBool()
	} else {
		namespaceAware = nil
	}
	namespaceuri := new(int64)
	if !r.Config.Namespaceuri.IsUnknown() && !r.Config.Namespaceuri.IsNull() {
		*namespaceuri = r.Config.Namespaceuri.ValueInt64()
	} else {
		namespaceuri = nil
	}
	pidata := new(int64)
	if !r.Config.Pidata.IsUnknown() && !r.Config.Pidata.IsNull() {
		*pidata = r.Config.Pidata.ValueInt64()
	} else {
		pidata = nil
	}
	pitarget := new(int64)
	if !r.Config.Pitarget.IsUnknown() && !r.Config.Pitarget.IsNull() {
		*pitarget = r.Config.Pitarget.ValueInt64()
	} else {
		pitarget = nil
	}
	prefix := new(int64)
	if !r.Config.Prefix.IsUnknown() && !r.Config.Prefix.IsNull() {
		*prefix = r.Config.Prefix.ValueInt64()
	} else {
		prefix = nil
	}
	text := new(int64)
	if !r.Config.Text.IsUnknown() && !r.Config.Text.IsNull() {
		*text = r.Config.Text.ValueInt64()
	} else {
		text = nil
	}
	config := shared.XMLThreatProtectionPluginConfig{
		AllowDtd:            allowDtd,
		AllowedContentTypes: allowedContentTypes,
		Attribute:           attribute,
		BlaMaxAmplification: blaMaxAmplification,
		BlaThreshold:        blaThreshold,
		Buffer:              buffer,
		CheckedContentTypes: checkedContentTypes,
		Comment:             comment,
		Document:            document,
		Entity:              entity,
		Entityname:          entityname,
		Entityproperty:      entityproperty,
		Localname:           localname,
		MaxAttributes:       maxAttributes,
		MaxChildren:         maxChildren,
		MaxDepth:            maxDepth,
		MaxNamespaces:       maxNamespaces,
		NamespaceAware:      namespaceAware,
		Namespaceuri:        namespaceuri,
		Pidata:              pidata,
		Pitarget:            pitarget,
		Prefix:              prefix,
		Text:                text,
	}
	var consumer *shared.XMLThreatProtectionPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.XMLThreatProtectionPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.XMLThreatProtectionPluginProtocols = []shared.XMLThreatProtectionPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.XMLThreatProtectionPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.XMLThreatProtectionPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.XMLThreatProtectionPluginRoute{
			ID: id2,
		}
	}
	var service *shared.XMLThreatProtectionPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.XMLThreatProtectionPluginService{
			ID: id3,
		}
	}
	out := shared.XMLThreatProtectionPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginXMLThreatProtectionResourceModel) RefreshFromSharedXMLThreatProtectionPlugin(resp *shared.XMLThreatProtectionPlugin) {
	if resp != nil {
		r.Config.AllowDtd = types.BoolPointerValue(resp.Config.AllowDtd)
		r.Config.AllowedContentTypes = make([]types.String, 0, len(resp.Config.AllowedContentTypes))
		for _, v := range resp.Config.AllowedContentTypes {
			r.Config.AllowedContentTypes = append(r.Config.AllowedContentTypes, types.StringValue(v))
		}
		r.Config.Attribute = types.Int64PointerValue(resp.Config.Attribute)
		if resp.Config.BlaMaxAmplification != nil {
			r.Config.BlaMaxAmplification = types.NumberValue(big.NewFloat(float64(*resp.Config.BlaMaxAmplification)))
		} else {
			r.Config.BlaMaxAmplification = types.NumberNull()
		}
		r.Config.BlaThreshold = types.Int64PointerValue(resp.Config.BlaThreshold)
		r.Config.Buffer = types.Int64PointerValue(resp.Config.Buffer)
		r.Config.CheckedContentTypes = make([]types.String, 0, len(resp.Config.CheckedContentTypes))
		for _, v := range resp.Config.CheckedContentTypes {
			r.Config.CheckedContentTypes = append(r.Config.CheckedContentTypes, types.StringValue(v))
		}
		r.Config.Comment = types.Int64PointerValue(resp.Config.Comment)
		r.Config.Document = types.Int64PointerValue(resp.Config.Document)
		r.Config.Entity = types.Int64PointerValue(resp.Config.Entity)
		r.Config.Entityname = types.Int64PointerValue(resp.Config.Entityname)
		r.Config.Entityproperty = types.Int64PointerValue(resp.Config.Entityproperty)
		r.Config.Localname = types.Int64PointerValue(resp.Config.Localname)
		r.Config.MaxAttributes = types.Int64PointerValue(resp.Config.MaxAttributes)
		r.Config.MaxChildren = types.Int64PointerValue(resp.Config.MaxChildren)
		r.Config.MaxDepth = types.Int64PointerValue(resp.Config.MaxDepth)
		r.Config.MaxNamespaces = types.Int64PointerValue(resp.Config.MaxNamespaces)
		r.Config.NamespaceAware = types.BoolPointerValue(resp.Config.NamespaceAware)
		r.Config.Namespaceuri = types.Int64PointerValue(resp.Config.Namespaceuri)
		r.Config.Pidata = types.Int64PointerValue(resp.Config.Pidata)
		r.Config.Pitarget = types.Int64PointerValue(resp.Config.Pitarget)
		r.Config.Prefix = types.Int64PointerValue(resp.Config.Prefix)
		r.Config.Text = types.Int64PointerValue(resp.Config.Text)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
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
