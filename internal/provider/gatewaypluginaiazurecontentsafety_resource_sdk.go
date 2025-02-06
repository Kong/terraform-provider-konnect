// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginAiAzureContentSafetyResourceModel) ToSharedAiAzureContentSafetyPluginInput() *shared.AiAzureContentSafetyPluginInput {
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
	var ordering *shared.AiAzureContentSafetyPluginOrdering
	if r.Ordering != nil {
		var after *shared.AiAzureContentSafetyPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AiAzureContentSafetyPluginAfter{
				Access: access,
			}
		}
		var before *shared.AiAzureContentSafetyPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AiAzureContentSafetyPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AiAzureContentSafetyPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	azureAPIVersion := new(string)
	if !r.Config.AzureAPIVersion.IsUnknown() && !r.Config.AzureAPIVersion.IsNull() {
		*azureAPIVersion = r.Config.AzureAPIVersion.ValueString()
	} else {
		azureAPIVersion = nil
	}
	azureClientID := new(string)
	if !r.Config.AzureClientID.IsUnknown() && !r.Config.AzureClientID.IsNull() {
		*azureClientID = r.Config.AzureClientID.ValueString()
	} else {
		azureClientID = nil
	}
	azureClientSecret := new(string)
	if !r.Config.AzureClientSecret.IsUnknown() && !r.Config.AzureClientSecret.IsNull() {
		*azureClientSecret = r.Config.AzureClientSecret.ValueString()
	} else {
		azureClientSecret = nil
	}
	azureTenantID := new(string)
	if !r.Config.AzureTenantID.IsUnknown() && !r.Config.AzureTenantID.IsNull() {
		*azureTenantID = r.Config.AzureTenantID.ValueString()
	} else {
		azureTenantID = nil
	}
	azureUseManagedIdentity := new(bool)
	if !r.Config.AzureUseManagedIdentity.IsUnknown() && !r.Config.AzureUseManagedIdentity.IsNull() {
		*azureUseManagedIdentity = r.Config.AzureUseManagedIdentity.ValueBool()
	} else {
		azureUseManagedIdentity = nil
	}
	var blocklistNames []string = []string{}
	for _, blocklistNamesItem := range r.Config.BlocklistNames {
		blocklistNames = append(blocklistNames, blocklistNamesItem.ValueString())
	}
	var categories []shared.Categories = []shared.Categories{}
	for _, categoriesItem := range r.Config.Categories {
		var name string
		name = categoriesItem.Name.ValueString()

		var rejectionLevel int64
		rejectionLevel = categoriesItem.RejectionLevel.ValueInt64()

		categories = append(categories, shared.Categories{
			Name:           name,
			RejectionLevel: rejectionLevel,
		})
	}
	contentSafetyKey := new(string)
	if !r.Config.ContentSafetyKey.IsUnknown() && !r.Config.ContentSafetyKey.IsNull() {
		*contentSafetyKey = r.Config.ContentSafetyKey.ValueString()
	} else {
		contentSafetyKey = nil
	}
	contentSafetyURL := new(string)
	if !r.Config.ContentSafetyURL.IsUnknown() && !r.Config.ContentSafetyURL.IsNull() {
		*contentSafetyURL = r.Config.ContentSafetyURL.ValueString()
	} else {
		contentSafetyURL = nil
	}
	haltOnBlocklistHit := new(bool)
	if !r.Config.HaltOnBlocklistHit.IsUnknown() && !r.Config.HaltOnBlocklistHit.IsNull() {
		*haltOnBlocklistHit = r.Config.HaltOnBlocklistHit.ValueBool()
	} else {
		haltOnBlocklistHit = nil
	}
	outputType := new(shared.OutputType)
	if !r.Config.OutputType.IsUnknown() && !r.Config.OutputType.IsNull() {
		*outputType = shared.OutputType(r.Config.OutputType.ValueString())
	} else {
		outputType = nil
	}
	revealFailureReason := new(bool)
	if !r.Config.RevealFailureReason.IsUnknown() && !r.Config.RevealFailureReason.IsNull() {
		*revealFailureReason = r.Config.RevealFailureReason.ValueBool()
	} else {
		revealFailureReason = nil
	}
	textSource := new(shared.TextSource)
	if !r.Config.TextSource.IsUnknown() && !r.Config.TextSource.IsNull() {
		*textSource = shared.TextSource(r.Config.TextSource.ValueString())
	} else {
		textSource = nil
	}
	config := shared.AiAzureContentSafetyPluginConfig{
		AzureAPIVersion:         azureAPIVersion,
		AzureClientID:           azureClientID,
		AzureClientSecret:       azureClientSecret,
		AzureTenantID:           azureTenantID,
		AzureUseManagedIdentity: azureUseManagedIdentity,
		BlocklistNames:          blocklistNames,
		Categories:              categories,
		ContentSafetyKey:        contentSafetyKey,
		ContentSafetyURL:        contentSafetyURL,
		HaltOnBlocklistHit:      haltOnBlocklistHit,
		OutputType:              outputType,
		RevealFailureReason:     revealFailureReason,
		TextSource:              textSource,
	}
	var protocols []shared.AiAzureContentSafetyPluginProtocols = []shared.AiAzureContentSafetyPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AiAzureContentSafetyPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AiAzureContentSafetyPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.AiAzureContentSafetyPluginRoute{
			ID: id1,
		}
	}
	var service *shared.AiAzureContentSafetyPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.AiAzureContentSafetyPluginService{
			ID: id2,
		}
	}
	out := shared.AiAzureContentSafetyPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginAiAzureContentSafetyResourceModel) RefreshFromSharedAiAzureContentSafetyPlugin(resp *shared.AiAzureContentSafetyPlugin) {
	if resp != nil {
		r.Config.AzureAPIVersion = types.StringPointerValue(resp.Config.AzureAPIVersion)
		r.Config.AzureClientID = types.StringPointerValue(resp.Config.AzureClientID)
		r.Config.AzureClientSecret = types.StringPointerValue(resp.Config.AzureClientSecret)
		r.Config.AzureTenantID = types.StringPointerValue(resp.Config.AzureTenantID)
		r.Config.AzureUseManagedIdentity = types.BoolPointerValue(resp.Config.AzureUseManagedIdentity)
		r.Config.BlocklistNames = []types.String{}
		for _, v := range resp.Config.BlocklistNames {
			r.Config.BlocklistNames = append(r.Config.BlocklistNames, types.StringValue(v))
		}
		r.Config.Categories = []tfTypes.Categories{}
		if len(r.Config.Categories) > len(resp.Config.Categories) {
			r.Config.Categories = r.Config.Categories[:len(resp.Config.Categories)]
		}
		for categoriesCount, categoriesItem := range resp.Config.Categories {
			var categories1 tfTypes.Categories
			categories1.Name = types.StringValue(categoriesItem.Name)
			categories1.RejectionLevel = types.Int64Value(categoriesItem.RejectionLevel)
			if categoriesCount+1 > len(r.Config.Categories) {
				r.Config.Categories = append(r.Config.Categories, categories1)
			} else {
				r.Config.Categories[categoriesCount].Name = categories1.Name
				r.Config.Categories[categoriesCount].RejectionLevel = categories1.RejectionLevel
			}
		}
		r.Config.ContentSafetyKey = types.StringPointerValue(resp.Config.ContentSafetyKey)
		r.Config.ContentSafetyURL = types.StringPointerValue(resp.Config.ContentSafetyURL)
		r.Config.HaltOnBlocklistHit = types.BoolPointerValue(resp.Config.HaltOnBlocklistHit)
		if resp.Config.OutputType != nil {
			r.Config.OutputType = types.StringValue(string(*resp.Config.OutputType))
		} else {
			r.Config.OutputType = types.StringNull()
		}
		r.Config.RevealFailureReason = types.BoolPointerValue(resp.Config.RevealFailureReason)
		if resp.Config.TextSource != nil {
			r.Config.TextSource = types.StringValue(string(*resp.Config.TextSource))
		} else {
			r.Config.TextSource = types.StringNull()
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
