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

func (r *GatewayPluginAiSemanticCacheDataSourceModel) ToOperationsGetAisemanticcachePluginRequest(ctx context.Context) (*operations.GetAisemanticcachePluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetAisemanticcachePluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginAiSemanticCacheDataSourceModel) RefreshFromSharedAiSemanticCachePlugin(ctx context.Context, resp *shared.AiSemanticCachePlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.AiSemanticCachePluginConfig{}
			r.Config.CacheControl = types.BoolPointerValue(resp.Config.CacheControl)
			r.Config.CacheTTL = types.Int64PointerValue(resp.Config.CacheTTL)
			if resp.Config.Embeddings == nil {
				r.Config.Embeddings = nil
			} else {
				r.Config.Embeddings = &tfTypes.AiRagInjectorPluginEmbeddings{}
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
				if resp.Config.Embeddings.Model == nil {
					r.Config.Embeddings.Model = nil
				} else {
					r.Config.Embeddings.Model = &tfTypes.AiRagInjectorPluginModel{}
					r.Config.Embeddings.Model.Name = types.StringPointerValue(resp.Config.Embeddings.Model.Name)
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
					if resp.Config.Embeddings.Model.Provider != nil {
						r.Config.Embeddings.Model.Provider = types.StringValue(string(*resp.Config.Embeddings.Model.Provider))
					} else {
						r.Config.Embeddings.Model.Provider = types.StringNull()
					}
				}
			}
			r.Config.ExactCaching = types.BoolPointerValue(resp.Config.ExactCaching)
			r.Config.IgnoreAssistantPrompts = types.BoolPointerValue(resp.Config.IgnoreAssistantPrompts)
			r.Config.IgnoreSystemPrompts = types.BoolPointerValue(resp.Config.IgnoreSystemPrompts)
			r.Config.IgnoreToolPrompts = types.BoolPointerValue(resp.Config.IgnoreToolPrompts)
			if resp.Config.LlmFormat != nil {
				r.Config.LlmFormat = types.StringValue(string(*resp.Config.LlmFormat))
			} else {
				r.Config.LlmFormat = types.StringNull()
			}
			r.Config.MessageCountback = types.Float64PointerValue(resp.Config.MessageCountback)
			r.Config.StopOnFailure = types.BoolPointerValue(resp.Config.StopOnFailure)
			if resp.Config.Vectordb == nil {
				r.Config.Vectordb = nil
			} else {
				r.Config.Vectordb = &tfTypes.AiSemanticCachePluginVectordb{}
				r.Config.Vectordb.Dimensions = types.Int64PointerValue(resp.Config.Vectordb.Dimensions)
				if resp.Config.Vectordb.DistanceMetric != nil {
					r.Config.Vectordb.DistanceMetric = types.StringValue(string(*resp.Config.Vectordb.DistanceMetric))
				} else {
					r.Config.Vectordb.DistanceMetric = types.StringNull()
				}
				if resp.Config.Vectordb.Pgvector == nil {
					r.Config.Vectordb.Pgvector = nil
				} else {
					r.Config.Vectordb.Pgvector = &tfTypes.Pgvector{}
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
				}
				if resp.Config.Vectordb.Redis == nil {
					r.Config.Vectordb.Redis = nil
				} else {
					r.Config.Vectordb.Redis = &tfTypes.PartialRedisEEConfig{}
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
				}
				if resp.Config.Vectordb.Strategy != nil {
					r.Config.Vectordb.Strategy = types.StringValue(string(*resp.Config.Vectordb.Strategy))
				} else {
					r.Config.Vectordb.Strategy = types.StringNull()
				}
				r.Config.Vectordb.Threshold = types.Float64PointerValue(resp.Config.Vectordb.Threshold)
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
