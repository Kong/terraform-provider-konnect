// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginKafkaLogDataSourceModel) ToOperationsGetKafkalogPluginRequest(ctx context.Context) (*operations.GetKafkalogPluginRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var pluginID string
	pluginID = r.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = r.ControlPlaneID.ValueString()

	out := operations.GetKafkalogPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}

	return &out, diags
}

func (r *GatewayPluginKafkaLogDataSourceModel) RefreshFromSharedKafkaLogPlugin(ctx context.Context, resp *shared.KafkaLogPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.KafkaLogPluginConfig{}
			if resp.Config.Authentication == nil {
				r.Config.Authentication = nil
			} else {
				r.Config.Authentication = &tfTypes.Authentication{}
				if resp.Config.Authentication.Mechanism != nil {
					r.Config.Authentication.Mechanism = types.StringValue(string(*resp.Config.Authentication.Mechanism))
				} else {
					r.Config.Authentication.Mechanism = types.StringNull()
				}
				r.Config.Authentication.Password = types.StringPointerValue(resp.Config.Authentication.Password)
				if resp.Config.Authentication.Strategy != nil {
					r.Config.Authentication.Strategy = types.StringValue(string(*resp.Config.Authentication.Strategy))
				} else {
					r.Config.Authentication.Strategy = types.StringNull()
				}
				r.Config.Authentication.Tokenauth = types.BoolPointerValue(resp.Config.Authentication.Tokenauth)
				r.Config.Authentication.User = types.StringPointerValue(resp.Config.Authentication.User)
			}
			r.Config.BootstrapServers = []tfTypes.BootstrapServers{}
			if len(r.Config.BootstrapServers) > len(resp.Config.BootstrapServers) {
				r.Config.BootstrapServers = r.Config.BootstrapServers[:len(resp.Config.BootstrapServers)]
			}
			for bootstrapServersCount, bootstrapServersItem := range resp.Config.BootstrapServers {
				var bootstrapServers tfTypes.BootstrapServers
				bootstrapServers.Host = types.StringValue(bootstrapServersItem.Host)
				bootstrapServers.Port = types.Int64Value(bootstrapServersItem.Port)
				if bootstrapServersCount+1 > len(r.Config.BootstrapServers) {
					r.Config.BootstrapServers = append(r.Config.BootstrapServers, bootstrapServers)
				} else {
					r.Config.BootstrapServers[bootstrapServersCount].Host = bootstrapServers.Host
					r.Config.BootstrapServers[bootstrapServersCount].Port = bootstrapServers.Port
				}
			}
			r.Config.ClusterName = types.StringPointerValue(resp.Config.ClusterName)
			if len(resp.Config.CustomFieldsByLua) > 0 {
				r.Config.CustomFieldsByLua = make(map[string]types.String, len(resp.Config.CustomFieldsByLua))
				for key, value := range resp.Config.CustomFieldsByLua {
					result, _ := json.Marshal(value)
					r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
				}
			}
			r.Config.Keepalive = types.Int64PointerValue(resp.Config.Keepalive)
			r.Config.KeepaliveEnabled = types.BoolPointerValue(resp.Config.KeepaliveEnabled)
			r.Config.ProducerAsync = types.BoolPointerValue(resp.Config.ProducerAsync)
			r.Config.ProducerAsyncBufferingLimitsMessagesInMemory = types.Int64PointerValue(resp.Config.ProducerAsyncBufferingLimitsMessagesInMemory)
			r.Config.ProducerAsyncFlushTimeout = types.Int64PointerValue(resp.Config.ProducerAsyncFlushTimeout)
			if resp.Config.ProducerRequestAcks != nil {
				r.Config.ProducerRequestAcks = types.Int64Value(int64(*resp.Config.ProducerRequestAcks))
			} else {
				r.Config.ProducerRequestAcks = types.Int64Null()
			}
			r.Config.ProducerRequestLimitsBytesPerRequest = types.Int64PointerValue(resp.Config.ProducerRequestLimitsBytesPerRequest)
			r.Config.ProducerRequestLimitsMessagesPerRequest = types.Int64PointerValue(resp.Config.ProducerRequestLimitsMessagesPerRequest)
			r.Config.ProducerRequestRetriesBackoffTimeout = types.Int64PointerValue(resp.Config.ProducerRequestRetriesBackoffTimeout)
			r.Config.ProducerRequestRetriesMaxAttempts = types.Int64PointerValue(resp.Config.ProducerRequestRetriesMaxAttempts)
			r.Config.ProducerRequestTimeout = types.Int64PointerValue(resp.Config.ProducerRequestTimeout)
			if resp.Config.Security == nil {
				r.Config.Security = nil
			} else {
				r.Config.Security = &tfTypes.KafkaConsumePluginSecurity{}
				r.Config.Security.CertificateID = types.StringPointerValue(resp.Config.Security.CertificateID)
				r.Config.Security.Ssl = types.BoolPointerValue(resp.Config.Security.Ssl)
			}
			r.Config.Timeout = types.Int64PointerValue(resp.Config.Timeout)
			r.Config.Topic = types.StringPointerValue(resp.Config.Topic)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.Set{}
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
