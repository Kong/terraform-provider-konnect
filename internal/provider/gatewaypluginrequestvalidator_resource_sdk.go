// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginRequestValidatorResourceModel) ToSharedRequestValidatorPluginInput() *shared.RequestValidatorPluginInput {
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
	var ordering *shared.RequestValidatorPluginOrdering
	if r.Ordering != nil {
		var after *shared.RequestValidatorPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.RequestValidatorPluginAfter{
				Access: access,
			}
		}
		var before *shared.RequestValidatorPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.RequestValidatorPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.RequestValidatorPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var allowedContentTypes []string = []string{}
	for _, allowedContentTypesItem := range r.Config.AllowedContentTypes {
		allowedContentTypes = append(allowedContentTypes, allowedContentTypesItem.ValueString())
	}
	bodySchema := new(string)
	if !r.Config.BodySchema.IsUnknown() && !r.Config.BodySchema.IsNull() {
		*bodySchema = r.Config.BodySchema.ValueString()
	} else {
		bodySchema = nil
	}
	contentTypeParameterValidation := new(bool)
	if !r.Config.ContentTypeParameterValidation.IsUnknown() && !r.Config.ContentTypeParameterValidation.IsNull() {
		*contentTypeParameterValidation = r.Config.ContentTypeParameterValidation.ValueBool()
	} else {
		contentTypeParameterValidation = nil
	}
	var parameterSchema []shared.ParameterSchema = []shared.ParameterSchema{}
	for _, parameterSchemaItem := range r.Config.ParameterSchema {
		explode := new(bool)
		if !parameterSchemaItem.Explode.IsUnknown() && !parameterSchemaItem.Explode.IsNull() {
			*explode = parameterSchemaItem.Explode.ValueBool()
		} else {
			explode = nil
		}
		in := shared.In(parameterSchemaItem.In.ValueString())
		var name string
		name = parameterSchemaItem.Name.ValueString()

		var required bool
		required = parameterSchemaItem.Required.ValueBool()

		schema := new(string)
		if !parameterSchemaItem.Schema.IsUnknown() && !parameterSchemaItem.Schema.IsNull() {
			*schema = parameterSchemaItem.Schema.ValueString()
		} else {
			schema = nil
		}
		style := new(shared.Style)
		if !parameterSchemaItem.Style.IsUnknown() && !parameterSchemaItem.Style.IsNull() {
			*style = shared.Style(parameterSchemaItem.Style.ValueString())
		} else {
			style = nil
		}
		parameterSchema = append(parameterSchema, shared.ParameterSchema{
			Explode:  explode,
			In:       in,
			Name:     name,
			Required: required,
			Schema:   schema,
			Style:    style,
		})
	}
	verboseResponse := new(bool)
	if !r.Config.VerboseResponse.IsUnknown() && !r.Config.VerboseResponse.IsNull() {
		*verboseResponse = r.Config.VerboseResponse.ValueBool()
	} else {
		verboseResponse = nil
	}
	version := new(shared.Version)
	if !r.Config.Version.IsUnknown() && !r.Config.Version.IsNull() {
		*version = shared.Version(r.Config.Version.ValueString())
	} else {
		version = nil
	}
	config := shared.RequestValidatorPluginConfig{
		AllowedContentTypes:            allowedContentTypes,
		BodySchema:                     bodySchema,
		ContentTypeParameterValidation: contentTypeParameterValidation,
		ParameterSchema:                parameterSchema,
		VerboseResponse:                verboseResponse,
		Version:                        version,
	}
	var consumer *shared.RequestValidatorPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.RequestValidatorPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.RequestValidatorPluginProtocols = []shared.RequestValidatorPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.RequestValidatorPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.RequestValidatorPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.RequestValidatorPluginRoute{
			ID: id2,
		}
	}
	var service *shared.RequestValidatorPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.RequestValidatorPluginService{
			ID: id3,
		}
	}
	out := shared.RequestValidatorPluginInput{
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

func (r *GatewayPluginRequestValidatorResourceModel) RefreshFromSharedRequestValidatorPlugin(resp *shared.RequestValidatorPlugin) {
	if resp != nil {
		r.Config.AllowedContentTypes = []types.String{}
		for _, v := range resp.Config.AllowedContentTypes {
			r.Config.AllowedContentTypes = append(r.Config.AllowedContentTypes, types.StringValue(v))
		}
		r.Config.BodySchema = types.StringPointerValue(resp.Config.BodySchema)
		r.Config.ContentTypeParameterValidation = types.BoolPointerValue(resp.Config.ContentTypeParameterValidation)
		r.Config.ParameterSchema = []tfTypes.ParameterSchema{}
		if len(r.Config.ParameterSchema) > len(resp.Config.ParameterSchema) {
			r.Config.ParameterSchema = r.Config.ParameterSchema[:len(resp.Config.ParameterSchema)]
		}
		for parameterSchemaCount, parameterSchemaItem := range resp.Config.ParameterSchema {
			var parameterSchema1 tfTypes.ParameterSchema
			parameterSchema1.Explode = types.BoolPointerValue(parameterSchemaItem.Explode)
			parameterSchema1.In = types.StringValue(string(parameterSchemaItem.In))
			parameterSchema1.Name = types.StringValue(parameterSchemaItem.Name)
			parameterSchema1.Required = types.BoolValue(parameterSchemaItem.Required)
			parameterSchema1.Schema = types.StringPointerValue(parameterSchemaItem.Schema)
			if parameterSchemaItem.Style != nil {
				parameterSchema1.Style = types.StringValue(string(*parameterSchemaItem.Style))
			} else {
				parameterSchema1.Style = types.StringNull()
			}
			if parameterSchemaCount+1 > len(r.Config.ParameterSchema) {
				r.Config.ParameterSchema = append(r.Config.ParameterSchema, parameterSchema1)
			} else {
				r.Config.ParameterSchema[parameterSchemaCount].Explode = parameterSchema1.Explode
				r.Config.ParameterSchema[parameterSchemaCount].In = parameterSchema1.In
				r.Config.ParameterSchema[parameterSchemaCount].Name = parameterSchema1.Name
				r.Config.ParameterSchema[parameterSchemaCount].Required = parameterSchema1.Required
				r.Config.ParameterSchema[parameterSchemaCount].Schema = parameterSchema1.Schema
				r.Config.ParameterSchema[parameterSchemaCount].Style = parameterSchema1.Style
			}
		}
		r.Config.VerboseResponse = types.BoolPointerValue(resp.Config.VerboseResponse)
		if resp.Config.Version != nil {
			r.Config.Version = types.StringValue(string(*resp.Config.Version))
		} else {
			r.Config.Version = types.StringNull()
		}
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
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = []types.String{}
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = []types.String{}
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
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
