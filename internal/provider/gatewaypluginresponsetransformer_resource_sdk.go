// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginResponseTransformerResourceModel) ToSharedResponseTransformerPlugin(ctx context.Context) (*shared.ResponseTransformerPlugin, diag.Diagnostics) {
	var diags diag.Diagnostics

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
	var ordering *shared.ResponseTransformerPluginOrdering
	if r.Ordering != nil {
		var after *shared.ResponseTransformerPluginAfter
		if r.Ordering.After != nil {
			access := make([]string, 0, len(r.Ordering.After.Access))
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ResponseTransformerPluginAfter{
				Access: access,
			}
		}
		var before *shared.ResponseTransformerPluginBefore
		if r.Ordering.Before != nil {
			access1 := make([]string, 0, len(r.Ordering.Before.Access))
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ResponseTransformerPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ResponseTransformerPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var partials []shared.ResponseTransformerPluginPartials
	if r.Partials != nil {
		partials = make([]shared.ResponseTransformerPluginPartials, 0, len(r.Partials))
		for _, partialsItem := range r.Partials {
			id1 := new(string)
			if !partialsItem.ID.IsUnknown() && !partialsItem.ID.IsNull() {
				*id1 = partialsItem.ID.ValueString()
			} else {
				id1 = nil
			}
			name := new(string)
			if !partialsItem.Name.IsUnknown() && !partialsItem.Name.IsNull() {
				*name = partialsItem.Name.ValueString()
			} else {
				name = nil
			}
			path := new(string)
			if !partialsItem.Path.IsUnknown() && !partialsItem.Path.IsNull() {
				*path = partialsItem.Path.ValueString()
			} else {
				path = nil
			}
			partials = append(partials, shared.ResponseTransformerPluginPartials{
				ID:   id1,
				Name: name,
				Path: path,
			})
		}
	}
	tags := make([]string, 0, len(r.Tags))
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.ResponseTransformerPluginConfig
	if r.Config != nil {
		var add *shared.ResponseTransformerPluginAdd
		if r.Config.Add != nil {
			headers := make([]string, 0, len(r.Config.Add.Headers))
			for _, headersItem := range r.Config.Add.Headers {
				headers = append(headers, headersItem.ValueString())
			}
			jsonVar := make([]string, 0, len(r.Config.Add.JSON))
			for _, jsonItem := range r.Config.Add.JSON {
				jsonVar = append(jsonVar, jsonItem.ValueString())
			}
			jsonTypes := make([]shared.ResponseTransformerPluginJSONTypes, 0, len(r.Config.Add.JSONTypes))
			for _, jsonTypesItem := range r.Config.Add.JSONTypes {
				jsonTypes = append(jsonTypes, shared.ResponseTransformerPluginJSONTypes(jsonTypesItem.ValueString()))
			}
			add = &shared.ResponseTransformerPluginAdd{
				Headers:   headers,
				JSON:      jsonVar,
				JSONTypes: jsonTypes,
			}
		}
		var append1 *shared.ResponseTransformerPluginAppend
		if r.Config.Append != nil {
			headers1 := make([]string, 0, len(r.Config.Append.Headers))
			for _, headersItem1 := range r.Config.Append.Headers {
				headers1 = append(headers1, headersItem1.ValueString())
			}
			jsonVar1 := make([]string, 0, len(r.Config.Append.JSON))
			for _, jsonItem1 := range r.Config.Append.JSON {
				jsonVar1 = append(jsonVar1, jsonItem1.ValueString())
			}
			jsonTypes1 := make([]shared.ResponseTransformerPluginConfigJSONTypes, 0, len(r.Config.Append.JSONTypes))
			for _, jsonTypesItem1 := range r.Config.Append.JSONTypes {
				jsonTypes1 = append(jsonTypes1, shared.ResponseTransformerPluginConfigJSONTypes(jsonTypesItem1.ValueString()))
			}
			append1 = &shared.ResponseTransformerPluginAppend{
				Headers:   headers1,
				JSON:      jsonVar1,
				JSONTypes: jsonTypes1,
			}
		}
		var remove *shared.ResponseTransformerPluginRemove
		if r.Config.Remove != nil {
			headers2 := make([]string, 0, len(r.Config.Remove.Headers))
			for _, headersItem2 := range r.Config.Remove.Headers {
				headers2 = append(headers2, headersItem2.ValueString())
			}
			jsonVar2 := make([]string, 0, len(r.Config.Remove.JSON))
			for _, jsonItem2 := range r.Config.Remove.JSON {
				jsonVar2 = append(jsonVar2, jsonItem2.ValueString())
			}
			remove = &shared.ResponseTransformerPluginRemove{
				Headers: headers2,
				JSON:    jsonVar2,
			}
		}
		var rename *shared.ResponseTransformerPluginRename
		if r.Config.Rename != nil {
			headers3 := make([]string, 0, len(r.Config.Rename.Headers))
			for _, headersItem3 := range r.Config.Rename.Headers {
				headers3 = append(headers3, headersItem3.ValueString())
			}
			jsonVar3 := make([]string, 0, len(r.Config.Rename.JSON))
			for _, jsonItem3 := range r.Config.Rename.JSON {
				jsonVar3 = append(jsonVar3, jsonItem3.ValueString())
			}
			rename = &shared.ResponseTransformerPluginRename{
				Headers: headers3,
				JSON:    jsonVar3,
			}
		}
		var replace *shared.ResponseTransformerPluginReplace
		if r.Config.Replace != nil {
			headers4 := make([]string, 0, len(r.Config.Replace.Headers))
			for _, headersItem4 := range r.Config.Replace.Headers {
				headers4 = append(headers4, headersItem4.ValueString())
			}
			jsonVar4 := make([]string, 0, len(r.Config.Replace.JSON))
			for _, jsonItem4 := range r.Config.Replace.JSON {
				jsonVar4 = append(jsonVar4, jsonItem4.ValueString())
			}
			jsonTypes2 := make([]shared.ResponseTransformerPluginConfigReplaceJSONTypes, 0, len(r.Config.Replace.JSONTypes))
			for _, jsonTypesItem2 := range r.Config.Replace.JSONTypes {
				jsonTypes2 = append(jsonTypes2, shared.ResponseTransformerPluginConfigReplaceJSONTypes(jsonTypesItem2.ValueString()))
			}
			replace = &shared.ResponseTransformerPluginReplace{
				Headers:   headers4,
				JSON:      jsonVar4,
				JSONTypes: jsonTypes2,
			}
		}
		config = &shared.ResponseTransformerPluginConfig{
			Add:     add,
			Append:  append1,
			Remove:  remove,
			Rename:  rename,
			Replace: replace,
		}
	}
	var consumer *shared.ResponseTransformerPluginConsumer
	if r.Consumer != nil {
		id2 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id2 = r.Consumer.ID.ValueString()
		} else {
			id2 = nil
		}
		consumer = &shared.ResponseTransformerPluginConsumer{
			ID: id2,
		}
	}
	var consumerGroup *shared.ResponseTransformerPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id3 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id3 = r.ConsumerGroup.ID.ValueString()
		} else {
			id3 = nil
		}
		consumerGroup = &shared.ResponseTransformerPluginConsumerGroup{
			ID: id3,
		}
	}
	protocols := make([]shared.ResponseTransformerPluginProtocols, 0, len(r.Protocols))
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ResponseTransformerPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ResponseTransformerPluginRoute
	if r.Route != nil {
		id4 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id4 = r.Route.ID.ValueString()
		} else {
			id4 = nil
		}
		route = &shared.ResponseTransformerPluginRoute{
			ID: id4,
		}
	}
	var service *shared.ResponseTransformerPluginService
	if r.Service != nil {
		id5 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id5 = r.Service.ID.ValueString()
		} else {
			id5 = nil
		}
		service = &shared.ResponseTransformerPluginService{
			ID: id5,
		}
	}
	out := shared.ResponseTransformerPlugin{
		CreatedAt:     createdAt,
		Enabled:       enabled,
		ID:            id,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Partials:      partials,
		Tags:          tags,
		UpdatedAt:     updatedAt,
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
	}

	return &out, diags
}

func (r *GatewayPluginResponseTransformerResourceModel) ToOperationsCreateResponsetransformerPluginRequest(ctx context.Context) (*operations.CreateResponsetransformerPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	responseTransformerPlugin, responseTransformerPluginDiags := r.ToSharedResponseTransformerPlugin(ctx)
	diags.Append(responseTransformerPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateResponsetransformerPluginRequest{
		ControlPlaneID:            controlPlaneID,
		ResponseTransformerPlugin: *responseTransformerPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginResponseTransformerResourceModel) ToOperationsUpdateResponsetransformerPluginRequest(ctx context.Context) (*operations.UpdateResponsetransformerPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	responseTransformerPlugin, responseTransformerPluginDiags := r.ToSharedResponseTransformerPlugin(ctx)
	diags.Append(responseTransformerPluginDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateResponsetransformerPluginRequest{
		PluginID:                  pluginID,
		ControlPlaneID:            controlPlaneID,
		ResponseTransformerPlugin: *responseTransformerPlugin,
	}

	return &out, diags
}

func (r *GatewayPluginResponseTransformerResourceModel) ToOperationsGetResponsetransformerPluginRequest(ctx context.Context) (*operations.GetResponsetransformerPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetResponsetransformerPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginResponseTransformerResourceModel) ToOperationsDeleteResponsetransformerPluginRequest(ctx context.Context) (*operations.DeleteResponsetransformerPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.DeleteResponsetransformerPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginResponseTransformerResourceModel) RefreshFromSharedResponseTransformerPlugin(ctx context.Context, resp *shared.ResponseTransformerPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.ResponseTransformerPluginConfig{}
			if resp.Config.Add == nil {
				r.Config.Add = nil
			} else {
				r.Config.Add = &tfTypes.ResponseTransformerPluginAdd{}
				r.Config.Add.Headers = make([]types.String, 0, len(resp.Config.Add.Headers))
				for _, v := range resp.Config.Add.Headers {
					r.Config.Add.Headers = append(r.Config.Add.Headers, types.StringValue(v))
				}
				r.Config.Add.JSON = make([]types.String, 0, len(resp.Config.Add.JSON))
				for _, v := range resp.Config.Add.JSON {
					r.Config.Add.JSON = append(r.Config.Add.JSON, types.StringValue(v))
				}
				r.Config.Add.JSONTypes = make([]types.String, 0, len(resp.Config.Add.JSONTypes))
				for _, v := range resp.Config.Add.JSONTypes {
					r.Config.Add.JSONTypes = append(r.Config.Add.JSONTypes, types.StringValue(string(v)))
				}
			}
			if resp.Config.Append == nil {
				r.Config.Append = nil
			} else {
				r.Config.Append = &tfTypes.ResponseTransformerPluginAdd{}
				r.Config.Append.Headers = make([]types.String, 0, len(resp.Config.Append.Headers))
				for _, v := range resp.Config.Append.Headers {
					r.Config.Append.Headers = append(r.Config.Append.Headers, types.StringValue(v))
				}
				r.Config.Append.JSON = make([]types.String, 0, len(resp.Config.Append.JSON))
				for _, v := range resp.Config.Append.JSON {
					r.Config.Append.JSON = append(r.Config.Append.JSON, types.StringValue(v))
				}
				r.Config.Append.JSONTypes = make([]types.String, 0, len(resp.Config.Append.JSONTypes))
				for _, v := range resp.Config.Append.JSONTypes {
					r.Config.Append.JSONTypes = append(r.Config.Append.JSONTypes, types.StringValue(string(v)))
				}
			}
			if resp.Config.Remove == nil {
				r.Config.Remove = nil
			} else {
				r.Config.Remove = &tfTypes.ResponseTransformerPluginRemove{}
				r.Config.Remove.Headers = make([]types.String, 0, len(resp.Config.Remove.Headers))
				for _, v := range resp.Config.Remove.Headers {
					r.Config.Remove.Headers = append(r.Config.Remove.Headers, types.StringValue(v))
				}
				r.Config.Remove.JSON = make([]types.String, 0, len(resp.Config.Remove.JSON))
				for _, v := range resp.Config.Remove.JSON {
					r.Config.Remove.JSON = append(r.Config.Remove.JSON, types.StringValue(v))
				}
			}
			if resp.Config.Rename == nil {
				r.Config.Rename = nil
			} else {
				r.Config.Rename = &tfTypes.ResponseTransformerPluginRemove{}
				r.Config.Rename.Headers = make([]types.String, 0, len(resp.Config.Rename.Headers))
				for _, v := range resp.Config.Rename.Headers {
					r.Config.Rename.Headers = append(r.Config.Rename.Headers, types.StringValue(v))
				}
				r.Config.Rename.JSON = make([]types.String, 0, len(resp.Config.Rename.JSON))
				for _, v := range resp.Config.Rename.JSON {
					r.Config.Rename.JSON = append(r.Config.Rename.JSON, types.StringValue(v))
				}
			}
			if resp.Config.Replace == nil {
				r.Config.Replace = nil
			} else {
				r.Config.Replace = &tfTypes.ResponseTransformerPluginAdd{}
				r.Config.Replace.Headers = make([]types.String, 0, len(resp.Config.Replace.Headers))
				for _, v := range resp.Config.Replace.Headers {
					r.Config.Replace.Headers = append(r.Config.Replace.Headers, types.StringValue(v))
				}
				r.Config.Replace.JSON = make([]types.String, 0, len(resp.Config.Replace.JSON))
				for _, v := range resp.Config.Replace.JSON {
					r.Config.Replace.JSON = append(r.Config.Replace.JSON, types.StringValue(v))
				}
				r.Config.Replace.JSONTypes = make([]types.String, 0, len(resp.Config.Replace.JSONTypes))
				for _, v := range resp.Config.Replace.JSONTypes {
					r.Config.Replace.JSONTypes = append(r.Config.Replace.JSONTypes, types.StringValue(string(v)))
				}
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.Set{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.Set{}
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
		if resp.Partials != nil {
			r.Partials = []tfTypes.Partials{}
			if len(r.Partials) > len(resp.Partials) {
				r.Partials = r.Partials[:len(resp.Partials)]
			}
			for partialsCount, partialsItem := range resp.Partials {
				var partials tfTypes.Partials
				partials.ID = types.StringPointerValue(partialsItem.ID)
				partials.Name = types.StringPointerValue(partialsItem.Name)
				partials.Path = types.StringPointerValue(partialsItem.Path)
				if partialsCount+1 > len(r.Partials) {
					r.Partials = append(r.Partials, partials)
				} else {
					r.Partials[partialsCount].ID = partials.ID
					r.Partials[partialsCount].Name = partials.Name
					r.Partials[partialsCount].Path = partials.Path
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
			r.Route = &tfTypes.Set{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.Set{}
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
