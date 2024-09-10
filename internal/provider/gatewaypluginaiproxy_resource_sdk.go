// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginAIProxyResourceModel) ToSharedCreateAIProxyPlugin() *shared.CreateAIProxyPlugin {
	var config *shared.CreateAIProxyPluginConfig
	if r.Config != nil {
		var auth *shared.CreateAIProxyPluginAuth
		if r.Config.Auth != nil {
			azureClientID := new(string)
			if !r.Config.Auth.AzureClientID.IsUnknown() && !r.Config.Auth.AzureClientID.IsNull() {
				*azureClientID = r.Config.Auth.AzureClientID.ValueString()
			} else {
				azureClientID = nil
			}
			azureClientSecret := new(string)
			if !r.Config.Auth.AzureClientSecret.IsUnknown() && !r.Config.Auth.AzureClientSecret.IsNull() {
				*azureClientSecret = r.Config.Auth.AzureClientSecret.ValueString()
			} else {
				azureClientSecret = nil
			}
			azureTenantID := new(string)
			if !r.Config.Auth.AzureTenantID.IsUnknown() && !r.Config.Auth.AzureTenantID.IsNull() {
				*azureTenantID = r.Config.Auth.AzureTenantID.ValueString()
			} else {
				azureTenantID = nil
			}
			azureUseManagedIdentity := new(bool)
			if !r.Config.Auth.AzureUseManagedIdentity.IsUnknown() && !r.Config.Auth.AzureUseManagedIdentity.IsNull() {
				*azureUseManagedIdentity = r.Config.Auth.AzureUseManagedIdentity.ValueBool()
			} else {
				azureUseManagedIdentity = nil
			}
			headerName := new(string)
			if !r.Config.Auth.HeaderName.IsUnknown() && !r.Config.Auth.HeaderName.IsNull() {
				*headerName = r.Config.Auth.HeaderName.ValueString()
			} else {
				headerName = nil
			}
			headerValue := new(string)
			if !r.Config.Auth.HeaderValue.IsUnknown() && !r.Config.Auth.HeaderValue.IsNull() {
				*headerValue = r.Config.Auth.HeaderValue.ValueString()
			} else {
				headerValue = nil
			}
			paramLocation := new(shared.CreateAIProxyPluginParamLocation)
			if !r.Config.Auth.ParamLocation.IsUnknown() && !r.Config.Auth.ParamLocation.IsNull() {
				*paramLocation = shared.CreateAIProxyPluginParamLocation(r.Config.Auth.ParamLocation.ValueString())
			} else {
				paramLocation = nil
			}
			paramName := new(string)
			if !r.Config.Auth.ParamName.IsUnknown() && !r.Config.Auth.ParamName.IsNull() {
				*paramName = r.Config.Auth.ParamName.ValueString()
			} else {
				paramName = nil
			}
			paramValue := new(string)
			if !r.Config.Auth.ParamValue.IsUnknown() && !r.Config.Auth.ParamValue.IsNull() {
				*paramValue = r.Config.Auth.ParamValue.ValueString()
			} else {
				paramValue = nil
			}
			auth = &shared.CreateAIProxyPluginAuth{
				AzureClientID:           azureClientID,
				AzureClientSecret:       azureClientSecret,
				AzureTenantID:           azureTenantID,
				AzureUseManagedIdentity: azureUseManagedIdentity,
				HeaderName:              headerName,
				HeaderValue:             headerValue,
				ParamLocation:           paramLocation,
				ParamName:               paramName,
				ParamValue:              paramValue,
			}
		}
		var logging *shared.CreateAIProxyPluginLogging
		if r.Config.Logging != nil {
			logPayloads := new(bool)
			if !r.Config.Logging.LogPayloads.IsUnknown() && !r.Config.Logging.LogPayloads.IsNull() {
				*logPayloads = r.Config.Logging.LogPayloads.ValueBool()
			} else {
				logPayloads = nil
			}
			logStatistics := new(bool)
			if !r.Config.Logging.LogStatistics.IsUnknown() && !r.Config.Logging.LogStatistics.IsNull() {
				*logStatistics = r.Config.Logging.LogStatistics.ValueBool()
			} else {
				logStatistics = nil
			}
			logging = &shared.CreateAIProxyPluginLogging{
				LogPayloads:   logPayloads,
				LogStatistics: logStatistics,
			}
		}
		var model *shared.CreateAIProxyPluginModel
		if r.Config.Model != nil {
			name := new(string)
			if !r.Config.Model.Name.IsUnknown() && !r.Config.Model.Name.IsNull() {
				*name = r.Config.Model.Name.ValueString()
			} else {
				name = nil
			}
			var optionsVar *shared.CreateAIProxyPluginOptions
			if r.Config.Model.Options != nil {
				anthropicVersion := new(string)
				if !r.Config.Model.Options.AnthropicVersion.IsUnknown() && !r.Config.Model.Options.AnthropicVersion.IsNull() {
					*anthropicVersion = r.Config.Model.Options.AnthropicVersion.ValueString()
				} else {
					anthropicVersion = nil
				}
				azureAPIVersion := new(string)
				if !r.Config.Model.Options.AzureAPIVersion.IsUnknown() && !r.Config.Model.Options.AzureAPIVersion.IsNull() {
					*azureAPIVersion = r.Config.Model.Options.AzureAPIVersion.ValueString()
				} else {
					azureAPIVersion = nil
				}
				azureDeploymentID := new(string)
				if !r.Config.Model.Options.AzureDeploymentID.IsUnknown() && !r.Config.Model.Options.AzureDeploymentID.IsNull() {
					*azureDeploymentID = r.Config.Model.Options.AzureDeploymentID.ValueString()
				} else {
					azureDeploymentID = nil
				}
				azureInstance := new(string)
				if !r.Config.Model.Options.AzureInstance.IsUnknown() && !r.Config.Model.Options.AzureInstance.IsNull() {
					*azureInstance = r.Config.Model.Options.AzureInstance.ValueString()
				} else {
					azureInstance = nil
				}
				llama2Format := new(shared.CreateAIProxyPluginLlama2Format)
				if !r.Config.Model.Options.Llama2Format.IsUnknown() && !r.Config.Model.Options.Llama2Format.IsNull() {
					*llama2Format = shared.CreateAIProxyPluginLlama2Format(r.Config.Model.Options.Llama2Format.ValueString())
				} else {
					llama2Format = nil
				}
				maxTokens := new(int64)
				if !r.Config.Model.Options.MaxTokens.IsUnknown() && !r.Config.Model.Options.MaxTokens.IsNull() {
					*maxTokens = r.Config.Model.Options.MaxTokens.ValueInt64()
				} else {
					maxTokens = nil
				}
				mistralFormat := new(shared.CreateAIProxyPluginMistralFormat)
				if !r.Config.Model.Options.MistralFormat.IsUnknown() && !r.Config.Model.Options.MistralFormat.IsNull() {
					*mistralFormat = shared.CreateAIProxyPluginMistralFormat(r.Config.Model.Options.MistralFormat.ValueString())
				} else {
					mistralFormat = nil
				}
				temperature := new(float64)
				if !r.Config.Model.Options.Temperature.IsUnknown() && !r.Config.Model.Options.Temperature.IsNull() {
					*temperature, _ = r.Config.Model.Options.Temperature.ValueBigFloat().Float64()
				} else {
					temperature = nil
				}
				topK := new(int64)
				if !r.Config.Model.Options.TopK.IsUnknown() && !r.Config.Model.Options.TopK.IsNull() {
					*topK = r.Config.Model.Options.TopK.ValueInt64()
				} else {
					topK = nil
				}
				topP := new(float64)
				if !r.Config.Model.Options.TopP.IsUnknown() && !r.Config.Model.Options.TopP.IsNull() {
					*topP, _ = r.Config.Model.Options.TopP.ValueBigFloat().Float64()
				} else {
					topP = nil
				}
				upstreamPath := new(string)
				if !r.Config.Model.Options.UpstreamPath.IsUnknown() && !r.Config.Model.Options.UpstreamPath.IsNull() {
					*upstreamPath = r.Config.Model.Options.UpstreamPath.ValueString()
				} else {
					upstreamPath = nil
				}
				upstreamURL := new(string)
				if !r.Config.Model.Options.UpstreamURL.IsUnknown() && !r.Config.Model.Options.UpstreamURL.IsNull() {
					*upstreamURL = r.Config.Model.Options.UpstreamURL.ValueString()
				} else {
					upstreamURL = nil
				}
				optionsVar = &shared.CreateAIProxyPluginOptions{
					AnthropicVersion:  anthropicVersion,
					AzureAPIVersion:   azureAPIVersion,
					AzureDeploymentID: azureDeploymentID,
					AzureInstance:     azureInstance,
					Llama2Format:      llama2Format,
					MaxTokens:         maxTokens,
					MistralFormat:     mistralFormat,
					Temperature:       temperature,
					TopK:              topK,
					TopP:              topP,
					UpstreamPath:      upstreamPath,
					UpstreamURL:       upstreamURL,
				}
			}
			provider := new(shared.CreateAIProxyPluginProvider)
			if !r.Config.Model.Provider.IsUnknown() && !r.Config.Model.Provider.IsNull() {
				*provider = shared.CreateAIProxyPluginProvider(r.Config.Model.Provider.ValueString())
			} else {
				provider = nil
			}
			model = &shared.CreateAIProxyPluginModel{
				Name:     name,
				Options:  optionsVar,
				Provider: provider,
			}
		}
		responseStreaming := new(shared.CreateAIProxyPluginResponseStreaming)
		if !r.Config.ResponseStreaming.IsUnknown() && !r.Config.ResponseStreaming.IsNull() {
			*responseStreaming = shared.CreateAIProxyPluginResponseStreaming(r.Config.ResponseStreaming.ValueString())
		} else {
			responseStreaming = nil
		}
		routeType := new(shared.CreateAIProxyPluginRouteType)
		if !r.Config.RouteType.IsUnknown() && !r.Config.RouteType.IsNull() {
			*routeType = shared.CreateAIProxyPluginRouteType(r.Config.RouteType.ValueString())
		} else {
			routeType = nil
		}
		config = &shared.CreateAIProxyPluginConfig{
			Auth:              auth,
			Logging:           logging,
			Model:             model,
			ResponseStreaming: responseStreaming,
			RouteType:         routeType,
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
	var ordering *shared.CreateAIProxyPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateAIProxyPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateAIProxyPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateAIProxyPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateAIProxyPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateAIProxyPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateAIProxyPluginProtocols = []shared.CreateAIProxyPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateAIProxyPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateAIProxyPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateAIProxyPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateAIProxyPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateAIProxyPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateAIProxyPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateAIProxyPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateAIProxyPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateAIProxyPluginService{
			ID: id3,
		}
	}
	out := shared.CreateAIProxyPlugin{
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

func (r *GatewayPluginAIProxyResourceModel) RefreshFromSharedAIProxyPlugin(resp *shared.AIProxyPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAIProxyPluginConfig{}
			if resp.Config.Auth == nil {
				r.Config.Auth = nil
			} else {
				r.Config.Auth = &tfTypes.CreateAIProxyPluginAuth{}
				r.Config.Auth.AzureClientID = types.StringPointerValue(resp.Config.Auth.AzureClientID)
				r.Config.Auth.AzureClientSecret = types.StringPointerValue(resp.Config.Auth.AzureClientSecret)
				r.Config.Auth.AzureTenantID = types.StringPointerValue(resp.Config.Auth.AzureTenantID)
				r.Config.Auth.AzureUseManagedIdentity = types.BoolPointerValue(resp.Config.Auth.AzureUseManagedIdentity)
				r.Config.Auth.HeaderName = types.StringPointerValue(resp.Config.Auth.HeaderName)
				r.Config.Auth.HeaderValue = types.StringPointerValue(resp.Config.Auth.HeaderValue)
				if resp.Config.Auth.ParamLocation != nil {
					r.Config.Auth.ParamLocation = types.StringValue(string(*resp.Config.Auth.ParamLocation))
				} else {
					r.Config.Auth.ParamLocation = types.StringNull()
				}
				r.Config.Auth.ParamName = types.StringPointerValue(resp.Config.Auth.ParamName)
				r.Config.Auth.ParamValue = types.StringPointerValue(resp.Config.Auth.ParamValue)
			}
			if resp.Config.Logging == nil {
				r.Config.Logging = nil
			} else {
				r.Config.Logging = &tfTypes.CreateAIProxyPluginLogging{}
				r.Config.Logging.LogPayloads = types.BoolPointerValue(resp.Config.Logging.LogPayloads)
				r.Config.Logging.LogStatistics = types.BoolPointerValue(resp.Config.Logging.LogStatistics)
			}
			if resp.Config.Model == nil {
				r.Config.Model = nil
			} else {
				r.Config.Model = &tfTypes.CreateAIProxyPluginModel{}
				r.Config.Model.Name = types.StringPointerValue(resp.Config.Model.Name)
				if resp.Config.Model.Options == nil {
					r.Config.Model.Options = nil
				} else {
					r.Config.Model.Options = &tfTypes.CreateAIProxyPluginOptions{}
					r.Config.Model.Options.AnthropicVersion = types.StringPointerValue(resp.Config.Model.Options.AnthropicVersion)
					r.Config.Model.Options.AzureAPIVersion = types.StringPointerValue(resp.Config.Model.Options.AzureAPIVersion)
					r.Config.Model.Options.AzureDeploymentID = types.StringPointerValue(resp.Config.Model.Options.AzureDeploymentID)
					r.Config.Model.Options.AzureInstance = types.StringPointerValue(resp.Config.Model.Options.AzureInstance)
					if resp.Config.Model.Options.Llama2Format != nil {
						r.Config.Model.Options.Llama2Format = types.StringValue(string(*resp.Config.Model.Options.Llama2Format))
					} else {
						r.Config.Model.Options.Llama2Format = types.StringNull()
					}
					r.Config.Model.Options.MaxTokens = types.Int64PointerValue(resp.Config.Model.Options.MaxTokens)
					if resp.Config.Model.Options.MistralFormat != nil {
						r.Config.Model.Options.MistralFormat = types.StringValue(string(*resp.Config.Model.Options.MistralFormat))
					} else {
						r.Config.Model.Options.MistralFormat = types.StringNull()
					}
					if resp.Config.Model.Options.Temperature != nil {
						r.Config.Model.Options.Temperature = types.NumberValue(big.NewFloat(float64(*resp.Config.Model.Options.Temperature)))
					} else {
						r.Config.Model.Options.Temperature = types.NumberNull()
					}
					r.Config.Model.Options.TopK = types.Int64PointerValue(resp.Config.Model.Options.TopK)
					if resp.Config.Model.Options.TopP != nil {
						r.Config.Model.Options.TopP = types.NumberValue(big.NewFloat(float64(*resp.Config.Model.Options.TopP)))
					} else {
						r.Config.Model.Options.TopP = types.NumberNull()
					}
					r.Config.Model.Options.UpstreamPath = types.StringPointerValue(resp.Config.Model.Options.UpstreamPath)
					r.Config.Model.Options.UpstreamURL = types.StringPointerValue(resp.Config.Model.Options.UpstreamURL)
				}
				if resp.Config.Model.Provider != nil {
					r.Config.Model.Provider = types.StringValue(string(*resp.Config.Model.Provider))
				} else {
					r.Config.Model.Provider = types.StringNull()
				}
			}
			if resp.Config.ResponseStreaming != nil {
				r.Config.ResponseStreaming = types.StringValue(string(*resp.Config.ResponseStreaming))
			} else {
				r.Config.ResponseStreaming = types.StringNull()
			}
			if resp.Config.RouteType != nil {
				r.Config.RouteType = types.StringValue(string(*resp.Config.RouteType))
			} else {
				r.Config.RouteType = types.StringNull()
			}
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
