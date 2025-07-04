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

func (r *GatewayPluginAiRequestTransformerDataSourceModel) ToOperationsGetAirequesttransformerPluginRequest(ctx context.Context) (*operations.GetAirequesttransformerPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetAirequesttransformerPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginAiRequestTransformerDataSourceModel) RefreshFromSharedAiRequestTransformerPlugin(ctx context.Context, resp *shared.AiRequestTransformerPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.AiRequestTransformerPluginConfig{}
			r.Config.HTTPProxyHost = types.StringPointerValue(resp.Config.HTTPProxyHost)
			r.Config.HTTPProxyPort = types.Int64PointerValue(resp.Config.HTTPProxyPort)
			r.Config.HTTPTimeout = types.Int64PointerValue(resp.Config.HTTPTimeout)
			r.Config.HTTPSProxyHost = types.StringPointerValue(resp.Config.HTTPSProxyHost)
			r.Config.HTTPSProxyPort = types.Int64PointerValue(resp.Config.HTTPSProxyPort)
			r.Config.HTTPSVerify = types.BoolPointerValue(resp.Config.HTTPSVerify)
			if resp.Config.Llm == nil {
				r.Config.Llm = nil
			} else {
				r.Config.Llm = &tfTypes.Llm{}
				if resp.Config.Llm.Auth == nil {
					r.Config.Llm.Auth = nil
				} else {
					r.Config.Llm.Auth = &tfTypes.Auth{}
					r.Config.Llm.Auth.AllowOverride = types.BoolPointerValue(resp.Config.Llm.Auth.AllowOverride)
					r.Config.Llm.Auth.AwsAccessKeyID = types.StringPointerValue(resp.Config.Llm.Auth.AwsAccessKeyID)
					r.Config.Llm.Auth.AwsSecretAccessKey = types.StringPointerValue(resp.Config.Llm.Auth.AwsSecretAccessKey)
					r.Config.Llm.Auth.AzureClientID = types.StringPointerValue(resp.Config.Llm.Auth.AzureClientID)
					r.Config.Llm.Auth.AzureClientSecret = types.StringPointerValue(resp.Config.Llm.Auth.AzureClientSecret)
					r.Config.Llm.Auth.AzureTenantID = types.StringPointerValue(resp.Config.Llm.Auth.AzureTenantID)
					r.Config.Llm.Auth.AzureUseManagedIdentity = types.BoolPointerValue(resp.Config.Llm.Auth.AzureUseManagedIdentity)
					r.Config.Llm.Auth.GcpServiceAccountJSON = types.StringPointerValue(resp.Config.Llm.Auth.GcpServiceAccountJSON)
					r.Config.Llm.Auth.GcpUseServiceAccount = types.BoolPointerValue(resp.Config.Llm.Auth.GcpUseServiceAccount)
					r.Config.Llm.Auth.HeaderName = types.StringPointerValue(resp.Config.Llm.Auth.HeaderName)
					r.Config.Llm.Auth.HeaderValue = types.StringPointerValue(resp.Config.Llm.Auth.HeaderValue)
					if resp.Config.Llm.Auth.ParamLocation != nil {
						r.Config.Llm.Auth.ParamLocation = types.StringValue(string(*resp.Config.Llm.Auth.ParamLocation))
					} else {
						r.Config.Llm.Auth.ParamLocation = types.StringNull()
					}
					r.Config.Llm.Auth.ParamName = types.StringPointerValue(resp.Config.Llm.Auth.ParamName)
					r.Config.Llm.Auth.ParamValue = types.StringPointerValue(resp.Config.Llm.Auth.ParamValue)
				}
				if resp.Config.Llm.Logging == nil {
					r.Config.Llm.Logging = nil
				} else {
					r.Config.Llm.Logging = &tfTypes.Logging{}
					r.Config.Llm.Logging.LogPayloads = types.BoolPointerValue(resp.Config.Llm.Logging.LogPayloads)
					r.Config.Llm.Logging.LogStatistics = types.BoolPointerValue(resp.Config.Llm.Logging.LogStatistics)
				}
				if resp.Config.Llm.Model == nil {
					r.Config.Llm.Model = nil
				} else {
					r.Config.Llm.Model = &tfTypes.Model{}
					r.Config.Llm.Model.Name = types.StringPointerValue(resp.Config.Llm.Model.Name)
					if resp.Config.Llm.Model.Options == nil {
						r.Config.Llm.Model.Options = nil
					} else {
						r.Config.Llm.Model.Options = &tfTypes.OptionsObj{}
						r.Config.Llm.Model.Options.AnthropicVersion = types.StringPointerValue(resp.Config.Llm.Model.Options.AnthropicVersion)
						r.Config.Llm.Model.Options.AzureAPIVersion = types.StringPointerValue(resp.Config.Llm.Model.Options.AzureAPIVersion)
						r.Config.Llm.Model.Options.AzureDeploymentID = types.StringPointerValue(resp.Config.Llm.Model.Options.AzureDeploymentID)
						r.Config.Llm.Model.Options.AzureInstance = types.StringPointerValue(resp.Config.Llm.Model.Options.AzureInstance)
						if resp.Config.Llm.Model.Options.Bedrock == nil {
							r.Config.Llm.Model.Options.Bedrock = nil
						} else {
							r.Config.Llm.Model.Options.Bedrock = &tfTypes.Bedrock{}
							r.Config.Llm.Model.Options.Bedrock.AwsAssumeRoleArn = types.StringPointerValue(resp.Config.Llm.Model.Options.Bedrock.AwsAssumeRoleArn)
							r.Config.Llm.Model.Options.Bedrock.AwsRegion = types.StringPointerValue(resp.Config.Llm.Model.Options.Bedrock.AwsRegion)
							r.Config.Llm.Model.Options.Bedrock.AwsRoleSessionName = types.StringPointerValue(resp.Config.Llm.Model.Options.Bedrock.AwsRoleSessionName)
							r.Config.Llm.Model.Options.Bedrock.AwsStsEndpointURL = types.StringPointerValue(resp.Config.Llm.Model.Options.Bedrock.AwsStsEndpointURL)
						}
						if resp.Config.Llm.Model.Options.Gemini == nil {
							r.Config.Llm.Model.Options.Gemini = nil
						} else {
							r.Config.Llm.Model.Options.Gemini = &tfTypes.Gemini{}
							r.Config.Llm.Model.Options.Gemini.APIEndpoint = types.StringPointerValue(resp.Config.Llm.Model.Options.Gemini.APIEndpoint)
							r.Config.Llm.Model.Options.Gemini.LocationID = types.StringPointerValue(resp.Config.Llm.Model.Options.Gemini.LocationID)
							r.Config.Llm.Model.Options.Gemini.ProjectID = types.StringPointerValue(resp.Config.Llm.Model.Options.Gemini.ProjectID)
						}
						if resp.Config.Llm.Model.Options.Huggingface == nil {
							r.Config.Llm.Model.Options.Huggingface = nil
						} else {
							r.Config.Llm.Model.Options.Huggingface = &tfTypes.Huggingface{}
							r.Config.Llm.Model.Options.Huggingface.UseCache = types.BoolPointerValue(resp.Config.Llm.Model.Options.Huggingface.UseCache)
							r.Config.Llm.Model.Options.Huggingface.WaitForModel = types.BoolPointerValue(resp.Config.Llm.Model.Options.Huggingface.WaitForModel)
						}
						r.Config.Llm.Model.Options.InputCost = types.Float64PointerValue(resp.Config.Llm.Model.Options.InputCost)
						if resp.Config.Llm.Model.Options.Llama2Format != nil {
							r.Config.Llm.Model.Options.Llama2Format = types.StringValue(string(*resp.Config.Llm.Model.Options.Llama2Format))
						} else {
							r.Config.Llm.Model.Options.Llama2Format = types.StringNull()
						}
						r.Config.Llm.Model.Options.MaxTokens = types.Int64PointerValue(resp.Config.Llm.Model.Options.MaxTokens)
						if resp.Config.Llm.Model.Options.MistralFormat != nil {
							r.Config.Llm.Model.Options.MistralFormat = types.StringValue(string(*resp.Config.Llm.Model.Options.MistralFormat))
						} else {
							r.Config.Llm.Model.Options.MistralFormat = types.StringNull()
						}
						r.Config.Llm.Model.Options.OutputCost = types.Float64PointerValue(resp.Config.Llm.Model.Options.OutputCost)
						r.Config.Llm.Model.Options.Temperature = types.Float64PointerValue(resp.Config.Llm.Model.Options.Temperature)
						r.Config.Llm.Model.Options.TopK = types.Int64PointerValue(resp.Config.Llm.Model.Options.TopK)
						r.Config.Llm.Model.Options.TopP = types.Float64PointerValue(resp.Config.Llm.Model.Options.TopP)
						r.Config.Llm.Model.Options.UpstreamPath = types.StringPointerValue(resp.Config.Llm.Model.Options.UpstreamPath)
						r.Config.Llm.Model.Options.UpstreamURL = types.StringPointerValue(resp.Config.Llm.Model.Options.UpstreamURL)
					}
					if resp.Config.Llm.Model.Provider != nil {
						r.Config.Llm.Model.Provider = types.StringValue(string(*resp.Config.Llm.Model.Provider))
					} else {
						r.Config.Llm.Model.Provider = types.StringNull()
					}
				}
				if resp.Config.Llm.RouteType != nil {
					r.Config.Llm.RouteType = types.StringValue(string(*resp.Config.Llm.RouteType))
				} else {
					r.Config.Llm.RouteType = types.StringNull()
				}
			}
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			r.Config.Prompt = types.StringPointerValue(resp.Config.Prompt)
			r.Config.TransformationExtractPattern = types.StringPointerValue(resp.Config.TransformationExtractPattern)
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
