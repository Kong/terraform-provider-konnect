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

func (r *GatewayPluginAiProxyAdvancedDataSourceModel) ToOperationsGetAiproxyadvancedPluginRequest(ctx context.Context) (*operations.GetAiproxyadvancedPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetAiproxyadvancedPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginAiProxyAdvancedDataSourceModel) RefreshFromSharedAiProxyAdvancedPlugin(ctx context.Context, resp *shared.AiProxyAdvancedPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.AiProxyAdvancedPluginConfig{}
			if resp.Config.Balancer == nil {
				r.Config.Balancer = nil
			} else {
				r.Config.Balancer = &tfTypes.Balancer{}
				if resp.Config.Balancer.Algorithm != nil {
					r.Config.Balancer.Algorithm = types.StringValue(string(*resp.Config.Balancer.Algorithm))
				} else {
					r.Config.Balancer.Algorithm = types.StringNull()
				}
				r.Config.Balancer.ConnectTimeout = types.Int64PointerValue(resp.Config.Balancer.ConnectTimeout)
				r.Config.Balancer.FailoverCriteria = make([]types.String, 0, len(resp.Config.Balancer.FailoverCriteria))
				for _, v := range resp.Config.Balancer.FailoverCriteria {
					r.Config.Balancer.FailoverCriteria = append(r.Config.Balancer.FailoverCriteria, types.StringValue(string(v)))
				}
				r.Config.Balancer.HashOnHeader = types.StringPointerValue(resp.Config.Balancer.HashOnHeader)
				if resp.Config.Balancer.LatencyStrategy != nil {
					r.Config.Balancer.LatencyStrategy = types.StringValue(string(*resp.Config.Balancer.LatencyStrategy))
				} else {
					r.Config.Balancer.LatencyStrategy = types.StringNull()
				}
				r.Config.Balancer.ReadTimeout = types.Int64PointerValue(resp.Config.Balancer.ReadTimeout)
				r.Config.Balancer.Retries = types.Int64PointerValue(resp.Config.Balancer.Retries)
				r.Config.Balancer.Slots = types.Int64PointerValue(resp.Config.Balancer.Slots)
				if resp.Config.Balancer.TokensCountStrategy != nil {
					r.Config.Balancer.TokensCountStrategy = types.StringValue(string(*resp.Config.Balancer.TokensCountStrategy))
				} else {
					r.Config.Balancer.TokensCountStrategy = types.StringNull()
				}
				r.Config.Balancer.WriteTimeout = types.Int64PointerValue(resp.Config.Balancer.WriteTimeout)
			}
			if resp.Config.Embeddings == nil {
				r.Config.Embeddings = nil
			} else {
				r.Config.Embeddings = &tfTypes.Embeddings{}
				if resp.Config.Embeddings.Auth == nil {
					r.Config.Embeddings.Auth = nil
				} else {
					r.Config.Embeddings.Auth = &tfTypes.Auth{}
					r.Config.Embeddings.Auth.AllowOverride = types.BoolPointerValue(resp.Config.Embeddings.Auth.AllowOverride)
					r.Config.Embeddings.Auth.AwsAccessKeyID = types.StringPointerValue(resp.Config.Embeddings.Auth.AwsAccessKeyID)
					r.Config.Embeddings.Auth.AwsSecretAccessKey = types.StringPointerValue(resp.Config.Embeddings.Auth.AwsSecretAccessKey)
					r.Config.Embeddings.Auth.AzureClientID = types.StringPointerValue(resp.Config.Embeddings.Auth.AzureClientID)
					r.Config.Embeddings.Auth.AzureClientSecret = types.StringPointerValue(resp.Config.Embeddings.Auth.AzureClientSecret)
					r.Config.Embeddings.Auth.AzureTenantID = types.StringPointerValue(resp.Config.Embeddings.Auth.AzureTenantID)
					r.Config.Embeddings.Auth.AzureUseManagedIdentity = types.BoolPointerValue(resp.Config.Embeddings.Auth.AzureUseManagedIdentity)
					r.Config.Embeddings.Auth.GcpServiceAccountJSON = types.StringPointerValue(resp.Config.Embeddings.Auth.GcpServiceAccountJSON)
					r.Config.Embeddings.Auth.GcpUseServiceAccount = types.BoolPointerValue(resp.Config.Embeddings.Auth.GcpUseServiceAccount)
					r.Config.Embeddings.Auth.HeaderName = types.StringPointerValue(resp.Config.Embeddings.Auth.HeaderName)
					r.Config.Embeddings.Auth.HeaderValue = types.StringPointerValue(resp.Config.Embeddings.Auth.HeaderValue)
					if resp.Config.Embeddings.Auth.ParamLocation != nil {
						r.Config.Embeddings.Auth.ParamLocation = types.StringValue(string(*resp.Config.Embeddings.Auth.ParamLocation))
					} else {
						r.Config.Embeddings.Auth.ParamLocation = types.StringNull()
					}
					r.Config.Embeddings.Auth.ParamName = types.StringPointerValue(resp.Config.Embeddings.Auth.ParamName)
					r.Config.Embeddings.Auth.ParamValue = types.StringPointerValue(resp.Config.Embeddings.Auth.ParamValue)
				}
				r.Config.Embeddings.Model.Name = types.StringValue(resp.Config.Embeddings.Model.Name)
				if resp.Config.Embeddings.Model.Options == nil {
					r.Config.Embeddings.Model.Options = nil
				} else {
					r.Config.Embeddings.Model.Options = &tfTypes.AiProxyAdvancedPluginOptions{}
					r.Config.Embeddings.Model.Options.Azure.APIVersion = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Azure.APIVersion)
					r.Config.Embeddings.Model.Options.Azure.DeploymentID = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Azure.DeploymentID)
					r.Config.Embeddings.Model.Options.Azure.Instance = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Azure.Instance)
					if resp.Config.Embeddings.Model.Options.Bedrock == nil {
						r.Config.Embeddings.Model.Options.Bedrock = nil
					} else {
						r.Config.Embeddings.Model.Options.Bedrock = &tfTypes.Bedrock{}
						r.Config.Embeddings.Model.Options.Bedrock.AwsAssumeRoleArn = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Bedrock.AwsAssumeRoleArn)
						r.Config.Embeddings.Model.Options.Bedrock.AwsRegion = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Bedrock.AwsRegion)
						r.Config.Embeddings.Model.Options.Bedrock.AwsRoleSessionName = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Bedrock.AwsRoleSessionName)
						r.Config.Embeddings.Model.Options.Bedrock.AwsStsEndpointURL = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Bedrock.AwsStsEndpointURL)
					}
					if resp.Config.Embeddings.Model.Options.Gemini == nil {
						r.Config.Embeddings.Model.Options.Gemini = nil
					} else {
						r.Config.Embeddings.Model.Options.Gemini = &tfTypes.Gemini{}
						r.Config.Embeddings.Model.Options.Gemini.APIEndpoint = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Gemini.APIEndpoint)
						r.Config.Embeddings.Model.Options.Gemini.LocationID = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Gemini.LocationID)
						r.Config.Embeddings.Model.Options.Gemini.ProjectID = types.StringPointerValue(resp.Config.Embeddings.Model.Options.Gemini.ProjectID)
					}
					if resp.Config.Embeddings.Model.Options.Huggingface == nil {
						r.Config.Embeddings.Model.Options.Huggingface = nil
					} else {
						r.Config.Embeddings.Model.Options.Huggingface = &tfTypes.Huggingface{}
						r.Config.Embeddings.Model.Options.Huggingface.UseCache = types.BoolPointerValue(resp.Config.Embeddings.Model.Options.Huggingface.UseCache)
						r.Config.Embeddings.Model.Options.Huggingface.WaitForModel = types.BoolPointerValue(resp.Config.Embeddings.Model.Options.Huggingface.WaitForModel)
					}
					r.Config.Embeddings.Model.Options.UpstreamURL = types.StringPointerValue(resp.Config.Embeddings.Model.Options.UpstreamURL)
				}
				r.Config.Embeddings.Model.Provider = types.StringValue(string(resp.Config.Embeddings.Model.Provider))
			}
			if resp.Config.LlmFormat != nil {
				r.Config.LlmFormat = types.StringValue(string(*resp.Config.LlmFormat))
			} else {
				r.Config.LlmFormat = types.StringNull()
			}
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			r.Config.ModelNameHeader = types.BoolPointerValue(resp.Config.ModelNameHeader)
			if resp.Config.ResponseStreaming != nil {
				r.Config.ResponseStreaming = types.StringValue(string(*resp.Config.ResponseStreaming))
			} else {
				r.Config.ResponseStreaming = types.StringNull()
			}
			r.Config.Targets = []tfTypes.Targets{}
			if len(r.Config.Targets) > len(resp.Config.Targets) {
				r.Config.Targets = r.Config.Targets[:len(resp.Config.Targets)]
			}
			for targetsCount, targetsItem := range resp.Config.Targets {
				var targets tfTypes.Targets
				if targetsItem.Auth == nil {
					targets.Auth = nil
				} else {
					targets.Auth = &tfTypes.Auth{}
					targets.Auth.AllowOverride = types.BoolPointerValue(targetsItem.Auth.AllowOverride)
					targets.Auth.AwsAccessKeyID = types.StringPointerValue(targetsItem.Auth.AwsAccessKeyID)
					targets.Auth.AwsSecretAccessKey = types.StringPointerValue(targetsItem.Auth.AwsSecretAccessKey)
					targets.Auth.AzureClientID = types.StringPointerValue(targetsItem.Auth.AzureClientID)
					targets.Auth.AzureClientSecret = types.StringPointerValue(targetsItem.Auth.AzureClientSecret)
					targets.Auth.AzureTenantID = types.StringPointerValue(targetsItem.Auth.AzureTenantID)
					targets.Auth.AzureUseManagedIdentity = types.BoolPointerValue(targetsItem.Auth.AzureUseManagedIdentity)
					targets.Auth.GcpServiceAccountJSON = types.StringPointerValue(targetsItem.Auth.GcpServiceAccountJSON)
					targets.Auth.GcpUseServiceAccount = types.BoolPointerValue(targetsItem.Auth.GcpUseServiceAccount)
					targets.Auth.HeaderName = types.StringPointerValue(targetsItem.Auth.HeaderName)
					targets.Auth.HeaderValue = types.StringPointerValue(targetsItem.Auth.HeaderValue)
					if targetsItem.Auth.ParamLocation != nil {
						targets.Auth.ParamLocation = types.StringValue(string(*targetsItem.Auth.ParamLocation))
					} else {
						targets.Auth.ParamLocation = types.StringNull()
					}
					targets.Auth.ParamName = types.StringPointerValue(targetsItem.Auth.ParamName)
					targets.Auth.ParamValue = types.StringPointerValue(targetsItem.Auth.ParamValue)
				}
				targets.Description = types.StringPointerValue(targetsItem.Description)
				targets.Logging.LogPayloads = types.BoolPointerValue(targetsItem.Logging.LogPayloads)
				targets.Logging.LogStatistics = types.BoolPointerValue(targetsItem.Logging.LogStatistics)
				targets.Model.Name = types.StringPointerValue(targetsItem.Model.Name)
				if targetsItem.Model.Options == nil {
					targets.Model.Options = nil
				} else {
					targets.Model.Options = &tfTypes.OptionsObj{}
					targets.Model.Options.AnthropicVersion = types.StringPointerValue(targetsItem.Model.Options.AnthropicVersion)
					targets.Model.Options.AzureAPIVersion = types.StringPointerValue(targetsItem.Model.Options.AzureAPIVersion)
					targets.Model.Options.AzureDeploymentID = types.StringPointerValue(targetsItem.Model.Options.AzureDeploymentID)
					targets.Model.Options.AzureInstance = types.StringPointerValue(targetsItem.Model.Options.AzureInstance)
					if targetsItem.Model.Options.Bedrock == nil {
						targets.Model.Options.Bedrock = nil
					} else {
						targets.Model.Options.Bedrock = &tfTypes.Bedrock{}
						targets.Model.Options.Bedrock.AwsAssumeRoleArn = types.StringPointerValue(targetsItem.Model.Options.Bedrock.AwsAssumeRoleArn)
						targets.Model.Options.Bedrock.AwsRegion = types.StringPointerValue(targetsItem.Model.Options.Bedrock.AwsRegion)
						targets.Model.Options.Bedrock.AwsRoleSessionName = types.StringPointerValue(targetsItem.Model.Options.Bedrock.AwsRoleSessionName)
						targets.Model.Options.Bedrock.AwsStsEndpointURL = types.StringPointerValue(targetsItem.Model.Options.Bedrock.AwsStsEndpointURL)
					}
					if targetsItem.Model.Options.Gemini == nil {
						targets.Model.Options.Gemini = nil
					} else {
						targets.Model.Options.Gemini = &tfTypes.Gemini{}
						targets.Model.Options.Gemini.APIEndpoint = types.StringPointerValue(targetsItem.Model.Options.Gemini.APIEndpoint)
						targets.Model.Options.Gemini.LocationID = types.StringPointerValue(targetsItem.Model.Options.Gemini.LocationID)
						targets.Model.Options.Gemini.ProjectID = types.StringPointerValue(targetsItem.Model.Options.Gemini.ProjectID)
					}
					if targetsItem.Model.Options.Huggingface == nil {
						targets.Model.Options.Huggingface = nil
					} else {
						targets.Model.Options.Huggingface = &tfTypes.Huggingface{}
						targets.Model.Options.Huggingface.UseCache = types.BoolPointerValue(targetsItem.Model.Options.Huggingface.UseCache)
						targets.Model.Options.Huggingface.WaitForModel = types.BoolPointerValue(targetsItem.Model.Options.Huggingface.WaitForModel)
					}
					targets.Model.Options.InputCost = types.Float64PointerValue(targetsItem.Model.Options.InputCost)
					if targetsItem.Model.Options.Llama2Format != nil {
						targets.Model.Options.Llama2Format = types.StringValue(string(*targetsItem.Model.Options.Llama2Format))
					} else {
						targets.Model.Options.Llama2Format = types.StringNull()
					}
					targets.Model.Options.MaxTokens = types.Int64PointerValue(targetsItem.Model.Options.MaxTokens)
					if targetsItem.Model.Options.MistralFormat != nil {
						targets.Model.Options.MistralFormat = types.StringValue(string(*targetsItem.Model.Options.MistralFormat))
					} else {
						targets.Model.Options.MistralFormat = types.StringNull()
					}
					targets.Model.Options.OutputCost = types.Float64PointerValue(targetsItem.Model.Options.OutputCost)
					targets.Model.Options.Temperature = types.Float64PointerValue(targetsItem.Model.Options.Temperature)
					targets.Model.Options.TopK = types.Int64PointerValue(targetsItem.Model.Options.TopK)
					targets.Model.Options.TopP = types.Float64PointerValue(targetsItem.Model.Options.TopP)
					targets.Model.Options.UpstreamPath = types.StringPointerValue(targetsItem.Model.Options.UpstreamPath)
					targets.Model.Options.UpstreamURL = types.StringPointerValue(targetsItem.Model.Options.UpstreamURL)
				}
				targets.Model.Provider = types.StringValue(string(targetsItem.Model.Provider))
				targets.RouteType = types.StringValue(string(targetsItem.RouteType))
				targets.Weight = types.Int64PointerValue(targetsItem.Weight)
				if targetsCount+1 > len(r.Config.Targets) {
					r.Config.Targets = append(r.Config.Targets, targets)
				} else {
					r.Config.Targets[targetsCount].Auth = targets.Auth
					r.Config.Targets[targetsCount].Description = targets.Description
					r.Config.Targets[targetsCount].Logging = targets.Logging
					r.Config.Targets[targetsCount].Model = targets.Model
					r.Config.Targets[targetsCount].RouteType = targets.RouteType
					r.Config.Targets[targetsCount].Weight = targets.Weight
				}
			}
			if resp.Config.Vectordb == nil {
				r.Config.Vectordb = nil
			} else {
				r.Config.Vectordb = &tfTypes.Vectordb{}
				r.Config.Vectordb.Dimensions = types.Int64Value(resp.Config.Vectordb.Dimensions)
				r.Config.Vectordb.DistanceMetric = types.StringValue(string(resp.Config.Vectordb.DistanceMetric))
				r.Config.Vectordb.Pgvector.Database = types.StringPointerValue(resp.Config.Vectordb.Pgvector.Database)
				r.Config.Vectordb.Pgvector.Host = types.StringPointerValue(resp.Config.Vectordb.Pgvector.Host)
				r.Config.Vectordb.Pgvector.Password = types.StringPointerValue(resp.Config.Vectordb.Pgvector.Password)
				r.Config.Vectordb.Pgvector.Port = types.Int64PointerValue(resp.Config.Vectordb.Pgvector.Port)
				r.Config.Vectordb.Pgvector.Ssl = types.BoolPointerValue(resp.Config.Vectordb.Pgvector.Ssl)
				r.Config.Vectordb.Pgvector.SslCert = types.StringPointerValue(resp.Config.Vectordb.Pgvector.SslCert)
				r.Config.Vectordb.Pgvector.SslCertKey = types.StringPointerValue(resp.Config.Vectordb.Pgvector.SslCertKey)
				r.Config.Vectordb.Pgvector.SslRequired = types.BoolPointerValue(resp.Config.Vectordb.Pgvector.SslRequired)
				r.Config.Vectordb.Pgvector.SslVerify = types.BoolPointerValue(resp.Config.Vectordb.Pgvector.SslVerify)
				if resp.Config.Vectordb.Pgvector.SslVersion != nil {
					r.Config.Vectordb.Pgvector.SslVersion = types.StringValue(string(*resp.Config.Vectordb.Pgvector.SslVersion))
				} else {
					r.Config.Vectordb.Pgvector.SslVersion = types.StringNull()
				}
				r.Config.Vectordb.Pgvector.Timeout = types.Float64PointerValue(resp.Config.Vectordb.Pgvector.Timeout)
				r.Config.Vectordb.Pgvector.User = types.StringPointerValue(resp.Config.Vectordb.Pgvector.User)
				r.Config.Vectordb.Redis.ClusterMaxRedirections = types.Int64PointerValue(resp.Config.Vectordb.Redis.ClusterMaxRedirections)
				r.Config.Vectordb.Redis.ClusterNodes = []tfTypes.PartialRedisEEClusterNodes{}
				if len(r.Config.Vectordb.Redis.ClusterNodes) > len(resp.Config.Vectordb.Redis.ClusterNodes) {
					r.Config.Vectordb.Redis.ClusterNodes = r.Config.Vectordb.Redis.ClusterNodes[:len(resp.Config.Vectordb.Redis.ClusterNodes)]
				}
				for clusterNodesCount, clusterNodesItem := range resp.Config.Vectordb.Redis.ClusterNodes {
					var clusterNodes tfTypes.PartialRedisEEClusterNodes
					clusterNodes.IP = types.StringPointerValue(clusterNodesItem.IP)
					clusterNodes.Port = types.Int64PointerValue(clusterNodesItem.Port)
					if clusterNodesCount+1 > len(r.Config.Vectordb.Redis.ClusterNodes) {
						r.Config.Vectordb.Redis.ClusterNodes = append(r.Config.Vectordb.Redis.ClusterNodes, clusterNodes)
					} else {
						r.Config.Vectordb.Redis.ClusterNodes[clusterNodesCount].IP = clusterNodes.IP
						r.Config.Vectordb.Redis.ClusterNodes[clusterNodesCount].Port = clusterNodes.Port
					}
				}
				r.Config.Vectordb.Redis.ConnectTimeout = types.Int64PointerValue(resp.Config.Vectordb.Redis.ConnectTimeout)
				r.Config.Vectordb.Redis.ConnectionIsProxied = types.BoolPointerValue(resp.Config.Vectordb.Redis.ConnectionIsProxied)
				r.Config.Vectordb.Redis.Database = types.Int64PointerValue(resp.Config.Vectordb.Redis.Database)
				r.Config.Vectordb.Redis.Host = types.StringPointerValue(resp.Config.Vectordb.Redis.Host)
				r.Config.Vectordb.Redis.KeepaliveBacklog = types.Int64PointerValue(resp.Config.Vectordb.Redis.KeepaliveBacklog)
				r.Config.Vectordb.Redis.KeepalivePoolSize = types.Int64PointerValue(resp.Config.Vectordb.Redis.KeepalivePoolSize)
				r.Config.Vectordb.Redis.Password = types.StringPointerValue(resp.Config.Vectordb.Redis.Password)
				r.Config.Vectordb.Redis.Port = types.Int64PointerValue(resp.Config.Vectordb.Redis.Port)
				r.Config.Vectordb.Redis.ReadTimeout = types.Int64PointerValue(resp.Config.Vectordb.Redis.ReadTimeout)
				r.Config.Vectordb.Redis.SendTimeout = types.Int64PointerValue(resp.Config.Vectordb.Redis.SendTimeout)
				r.Config.Vectordb.Redis.SentinelMaster = types.StringPointerValue(resp.Config.Vectordb.Redis.SentinelMaster)
				r.Config.Vectordb.Redis.SentinelNodes = []tfTypes.PartialRedisEESentinelNodes{}
				if len(r.Config.Vectordb.Redis.SentinelNodes) > len(resp.Config.Vectordb.Redis.SentinelNodes) {
					r.Config.Vectordb.Redis.SentinelNodes = r.Config.Vectordb.Redis.SentinelNodes[:len(resp.Config.Vectordb.Redis.SentinelNodes)]
				}
				for sentinelNodesCount, sentinelNodesItem := range resp.Config.Vectordb.Redis.SentinelNodes {
					var sentinelNodes tfTypes.PartialRedisEESentinelNodes
					sentinelNodes.Host = types.StringPointerValue(sentinelNodesItem.Host)
					sentinelNodes.Port = types.Int64PointerValue(sentinelNodesItem.Port)
					if sentinelNodesCount+1 > len(r.Config.Vectordb.Redis.SentinelNodes) {
						r.Config.Vectordb.Redis.SentinelNodes = append(r.Config.Vectordb.Redis.SentinelNodes, sentinelNodes)
					} else {
						r.Config.Vectordb.Redis.SentinelNodes[sentinelNodesCount].Host = sentinelNodes.Host
						r.Config.Vectordb.Redis.SentinelNodes[sentinelNodesCount].Port = sentinelNodes.Port
					}
				}
				r.Config.Vectordb.Redis.SentinelPassword = types.StringPointerValue(resp.Config.Vectordb.Redis.SentinelPassword)
				if resp.Config.Vectordb.Redis.SentinelRole != nil {
					r.Config.Vectordb.Redis.SentinelRole = types.StringValue(string(*resp.Config.Vectordb.Redis.SentinelRole))
				} else {
					r.Config.Vectordb.Redis.SentinelRole = types.StringNull()
				}
				r.Config.Vectordb.Redis.SentinelUsername = types.StringPointerValue(resp.Config.Vectordb.Redis.SentinelUsername)
				r.Config.Vectordb.Redis.ServerName = types.StringPointerValue(resp.Config.Vectordb.Redis.ServerName)
				r.Config.Vectordb.Redis.Ssl = types.BoolPointerValue(resp.Config.Vectordb.Redis.Ssl)
				r.Config.Vectordb.Redis.SslVerify = types.BoolPointerValue(resp.Config.Vectordb.Redis.SslVerify)
				r.Config.Vectordb.Redis.Username = types.StringPointerValue(resp.Config.Vectordb.Redis.Username)
				r.Config.Vectordb.Strategy = types.StringValue(string(resp.Config.Vectordb.Strategy))
				r.Config.Vectordb.Threshold = types.Float64Value(resp.Config.Vectordb.Threshold)
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
