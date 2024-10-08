// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginOasValidationResourceModel) ToSharedCreateOasValidationPlugin() *shared.CreateOasValidationPlugin {
	var config *shared.CreateOasValidationPluginConfig
	if r.Config != nil {
		allowedHeaderParameters := new(string)
		if !r.Config.AllowedHeaderParameters.IsUnknown() && !r.Config.AllowedHeaderParameters.IsNull() {
			*allowedHeaderParameters = r.Config.AllowedHeaderParameters.ValueString()
		} else {
			allowedHeaderParameters = nil
		}
		apiSpec := new(string)
		if !r.Config.APISpec.IsUnknown() && !r.Config.APISpec.IsNull() {
			*apiSpec = r.Config.APISpec.ValueString()
		} else {
			apiSpec = nil
		}
		apiSpecEncoded := new(bool)
		if !r.Config.APISpecEncoded.IsUnknown() && !r.Config.APISpecEncoded.IsNull() {
			*apiSpecEncoded = r.Config.APISpecEncoded.ValueBool()
		} else {
			apiSpecEncoded = nil
		}
		customBasePath := new(string)
		if !r.Config.CustomBasePath.IsUnknown() && !r.Config.CustomBasePath.IsNull() {
			*customBasePath = r.Config.CustomBasePath.ValueString()
		} else {
			customBasePath = nil
		}
		headerParameterCheck := new(bool)
		if !r.Config.HeaderParameterCheck.IsUnknown() && !r.Config.HeaderParameterCheck.IsNull() {
			*headerParameterCheck = r.Config.HeaderParameterCheck.ValueBool()
		} else {
			headerParameterCheck = nil
		}
		includeBasePath := new(bool)
		if !r.Config.IncludeBasePath.IsUnknown() && !r.Config.IncludeBasePath.IsNull() {
			*includeBasePath = r.Config.IncludeBasePath.ValueBool()
		} else {
			includeBasePath = nil
		}
		notifyOnlyRequestValidationFailure := new(bool)
		if !r.Config.NotifyOnlyRequestValidationFailure.IsUnknown() && !r.Config.NotifyOnlyRequestValidationFailure.IsNull() {
			*notifyOnlyRequestValidationFailure = r.Config.NotifyOnlyRequestValidationFailure.ValueBool()
		} else {
			notifyOnlyRequestValidationFailure = nil
		}
		notifyOnlyResponseBodyValidationFailure := new(bool)
		if !r.Config.NotifyOnlyResponseBodyValidationFailure.IsUnknown() && !r.Config.NotifyOnlyResponseBodyValidationFailure.IsNull() {
			*notifyOnlyResponseBodyValidationFailure = r.Config.NotifyOnlyResponseBodyValidationFailure.ValueBool()
		} else {
			notifyOnlyResponseBodyValidationFailure = nil
		}
		queryParameterCheck := new(bool)
		if !r.Config.QueryParameterCheck.IsUnknown() && !r.Config.QueryParameterCheck.IsNull() {
			*queryParameterCheck = r.Config.QueryParameterCheck.ValueBool()
		} else {
			queryParameterCheck = nil
		}
		validateRequestBody := new(bool)
		if !r.Config.ValidateRequestBody.IsUnknown() && !r.Config.ValidateRequestBody.IsNull() {
			*validateRequestBody = r.Config.ValidateRequestBody.ValueBool()
		} else {
			validateRequestBody = nil
		}
		validateRequestHeaderParams := new(bool)
		if !r.Config.ValidateRequestHeaderParams.IsUnknown() && !r.Config.ValidateRequestHeaderParams.IsNull() {
			*validateRequestHeaderParams = r.Config.ValidateRequestHeaderParams.ValueBool()
		} else {
			validateRequestHeaderParams = nil
		}
		validateRequestQueryParams := new(bool)
		if !r.Config.ValidateRequestQueryParams.IsUnknown() && !r.Config.ValidateRequestQueryParams.IsNull() {
			*validateRequestQueryParams = r.Config.ValidateRequestQueryParams.ValueBool()
		} else {
			validateRequestQueryParams = nil
		}
		validateRequestURIParams := new(bool)
		if !r.Config.ValidateRequestURIParams.IsUnknown() && !r.Config.ValidateRequestURIParams.IsNull() {
			*validateRequestURIParams = r.Config.ValidateRequestURIParams.ValueBool()
		} else {
			validateRequestURIParams = nil
		}
		validateResponseBody := new(bool)
		if !r.Config.ValidateResponseBody.IsUnknown() && !r.Config.ValidateResponseBody.IsNull() {
			*validateResponseBody = r.Config.ValidateResponseBody.ValueBool()
		} else {
			validateResponseBody = nil
		}
		verboseResponse := new(bool)
		if !r.Config.VerboseResponse.IsUnknown() && !r.Config.VerboseResponse.IsNull() {
			*verboseResponse = r.Config.VerboseResponse.ValueBool()
		} else {
			verboseResponse = nil
		}
		config = &shared.CreateOasValidationPluginConfig{
			AllowedHeaderParameters:                 allowedHeaderParameters,
			APISpec:                                 apiSpec,
			APISpecEncoded:                          apiSpecEncoded,
			CustomBasePath:                          customBasePath,
			HeaderParameterCheck:                    headerParameterCheck,
			IncludeBasePath:                         includeBasePath,
			NotifyOnlyRequestValidationFailure:      notifyOnlyRequestValidationFailure,
			NotifyOnlyResponseBodyValidationFailure: notifyOnlyResponseBodyValidationFailure,
			QueryParameterCheck:                     queryParameterCheck,
			ValidateRequestBody:                     validateRequestBody,
			ValidateRequestHeaderParams:             validateRequestHeaderParams,
			ValidateRequestQueryParams:              validateRequestQueryParams,
			ValidateRequestURIParams:                validateRequestURIParams,
			ValidateResponseBody:                    validateResponseBody,
			VerboseResponse:                         verboseResponse,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.CreateOasValidationPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateOasValidationPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateOasValidationPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateOasValidationPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateOasValidationPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateOasValidationPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateOasValidationPluginProtocols = []shared.CreateOasValidationPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateOasValidationPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateOasValidationPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateOasValidationPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateOasValidationPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateOasValidationPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateOasValidationPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateOasValidationPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateOasValidationPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateOasValidationPluginService{
			ID: id3,
		}
	}
	out := shared.CreateOasValidationPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginOasValidationResourceModel) RefreshFromSharedOasValidationPlugin(resp *shared.OasValidationPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateOasValidationPluginConfig{}
			r.Config.AllowedHeaderParameters = types.StringPointerValue(resp.Config.AllowedHeaderParameters)
			r.Config.APISpec = types.StringPointerValue(resp.Config.APISpec)
			r.Config.APISpecEncoded = types.BoolPointerValue(resp.Config.APISpecEncoded)
			r.Config.CustomBasePath = types.StringPointerValue(resp.Config.CustomBasePath)
			r.Config.HeaderParameterCheck = types.BoolPointerValue(resp.Config.HeaderParameterCheck)
			r.Config.IncludeBasePath = types.BoolPointerValue(resp.Config.IncludeBasePath)
			r.Config.NotifyOnlyRequestValidationFailure = types.BoolPointerValue(resp.Config.NotifyOnlyRequestValidationFailure)
			r.Config.NotifyOnlyResponseBodyValidationFailure = types.BoolPointerValue(resp.Config.NotifyOnlyResponseBodyValidationFailure)
			r.Config.QueryParameterCheck = types.BoolPointerValue(resp.Config.QueryParameterCheck)
			r.Config.ValidateRequestBody = types.BoolPointerValue(resp.Config.ValidateRequestBody)
			r.Config.ValidateRequestHeaderParams = types.BoolPointerValue(resp.Config.ValidateRequestHeaderParams)
			r.Config.ValidateRequestQueryParams = types.BoolPointerValue(resp.Config.ValidateRequestQueryParams)
			r.Config.ValidateRequestURIParams = types.BoolPointerValue(resp.Config.ValidateRequestURIParams)
			r.Config.ValidateResponseBody = types.BoolPointerValue(resp.Config.ValidateResponseBody)
			r.Config.VerboseResponse = types.BoolPointerValue(resp.Config.VerboseResponse)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.CreateACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.CreateACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.CreateACLPluginAfter{}
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
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
