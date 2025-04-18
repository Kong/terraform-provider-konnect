// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayPluginResponseRatelimitingResourceModel) ToSharedResponseRatelimitingPlugin() *shared.ResponseRatelimitingPlugin {
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
	var ordering *shared.ResponseRatelimitingPluginOrdering
	if r.Ordering != nil {
		var after *shared.ResponseRatelimitingPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.ResponseRatelimitingPluginAfter{
				Access: access,
			}
		}
		var before *shared.ResponseRatelimitingPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.ResponseRatelimitingPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.ResponseRatelimitingPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	updatedAt := new(int64)
	if !r.UpdatedAt.IsUnknown() && !r.UpdatedAt.IsNull() {
		*updatedAt = r.UpdatedAt.ValueInt64()
	} else {
		updatedAt = nil
	}
	var config *shared.ResponseRatelimitingPluginConfig
	if r.Config != nil {
		blockOnFirstViolation := new(bool)
		if !r.Config.BlockOnFirstViolation.IsUnknown() && !r.Config.BlockOnFirstViolation.IsNull() {
			*blockOnFirstViolation = r.Config.BlockOnFirstViolation.ValueBool()
		} else {
			blockOnFirstViolation = nil
		}
		faultTolerant := new(bool)
		if !r.Config.FaultTolerant.IsUnknown() && !r.Config.FaultTolerant.IsNull() {
			*faultTolerant = r.Config.FaultTolerant.ValueBool()
		} else {
			faultTolerant = nil
		}
		headerName := new(string)
		if !r.Config.HeaderName.IsUnknown() && !r.Config.HeaderName.IsNull() {
			*headerName = r.Config.HeaderName.ValueString()
		} else {
			headerName = nil
		}
		hideClientHeaders := new(bool)
		if !r.Config.HideClientHeaders.IsUnknown() && !r.Config.HideClientHeaders.IsNull() {
			*hideClientHeaders = r.Config.HideClientHeaders.ValueBool()
		} else {
			hideClientHeaders = nil
		}
		limitBy := new(shared.ResponseRatelimitingPluginLimitBy)
		if !r.Config.LimitBy.IsUnknown() && !r.Config.LimitBy.IsNull() {
			*limitBy = shared.ResponseRatelimitingPluginLimitBy(r.Config.LimitBy.ValueString())
		} else {
			limitBy = nil
		}
		limits := make(map[string]interface{})
		for limitsKey, limitsValue := range r.Config.Limits {
			var limitsInst interface{}
			_ = json.Unmarshal([]byte(limitsValue.ValueString()), &limitsInst)
			limits[limitsKey] = limitsInst
		}
		policy := new(shared.ResponseRatelimitingPluginPolicy)
		if !r.Config.Policy.IsUnknown() && !r.Config.Policy.IsNull() {
			*policy = shared.ResponseRatelimitingPluginPolicy(r.Config.Policy.ValueString())
		} else {
			policy = nil
		}
		var redis *shared.ResponseRatelimitingPluginRedis
		if r.Config.Redis != nil {
			database := new(int64)
			if !r.Config.Redis.Database.IsUnknown() && !r.Config.Redis.Database.IsNull() {
				*database = r.Config.Redis.Database.ValueInt64()
			} else {
				database = nil
			}
			host := new(string)
			if !r.Config.Redis.Host.IsUnknown() && !r.Config.Redis.Host.IsNull() {
				*host = r.Config.Redis.Host.ValueString()
			} else {
				host = nil
			}
			password := new(string)
			if !r.Config.Redis.Password.IsUnknown() && !r.Config.Redis.Password.IsNull() {
				*password = r.Config.Redis.Password.ValueString()
			} else {
				password = nil
			}
			port := new(int64)
			if !r.Config.Redis.Port.IsUnknown() && !r.Config.Redis.Port.IsNull() {
				*port = r.Config.Redis.Port.ValueInt64()
			} else {
				port = nil
			}
			serverName := new(string)
			if !r.Config.Redis.ServerName.IsUnknown() && !r.Config.Redis.ServerName.IsNull() {
				*serverName = r.Config.Redis.ServerName.ValueString()
			} else {
				serverName = nil
			}
			ssl := new(bool)
			if !r.Config.Redis.Ssl.IsUnknown() && !r.Config.Redis.Ssl.IsNull() {
				*ssl = r.Config.Redis.Ssl.ValueBool()
			} else {
				ssl = nil
			}
			sslVerify := new(bool)
			if !r.Config.Redis.SslVerify.IsUnknown() && !r.Config.Redis.SslVerify.IsNull() {
				*sslVerify = r.Config.Redis.SslVerify.ValueBool()
			} else {
				sslVerify = nil
			}
			timeout := new(int64)
			if !r.Config.Redis.Timeout.IsUnknown() && !r.Config.Redis.Timeout.IsNull() {
				*timeout = r.Config.Redis.Timeout.ValueInt64()
			} else {
				timeout = nil
			}
			username := new(string)
			if !r.Config.Redis.Username.IsUnknown() && !r.Config.Redis.Username.IsNull() {
				*username = r.Config.Redis.Username.ValueString()
			} else {
				username = nil
			}
			redis = &shared.ResponseRatelimitingPluginRedis{
				Database:   database,
				Host:       host,
				Password:   password,
				Port:       port,
				ServerName: serverName,
				Ssl:        ssl,
				SslVerify:  sslVerify,
				Timeout:    timeout,
				Username:   username,
			}
		}
		config = &shared.ResponseRatelimitingPluginConfig{
			BlockOnFirstViolation: blockOnFirstViolation,
			FaultTolerant:         faultTolerant,
			HeaderName:            headerName,
			HideClientHeaders:     hideClientHeaders,
			LimitBy:               limitBy,
			Limits:                limits,
			Policy:                policy,
			Redis:                 redis,
		}
	}
	var consumer *shared.ResponseRatelimitingPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.ResponseRatelimitingPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.ResponseRatelimitingPluginProtocols = []shared.ResponseRatelimitingPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.ResponseRatelimitingPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.ResponseRatelimitingPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.ResponseRatelimitingPluginRoute{
			ID: id2,
		}
	}
	var service *shared.ResponseRatelimitingPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.ResponseRatelimitingPluginService{
			ID: id3,
		}
	}
	out := shared.ResponseRatelimitingPlugin{
		CreatedAt:    createdAt,
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		UpdatedAt:    updatedAt,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginResponseRatelimitingResourceModel) RefreshFromSharedResponseRatelimitingPlugin(ctx context.Context, resp *shared.ResponseRatelimitingPlugin) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.ResponseRatelimitingPluginConfig{}
			r.Config.BlockOnFirstViolation = types.BoolPointerValue(resp.Config.BlockOnFirstViolation)
			r.Config.FaultTolerant = types.BoolPointerValue(resp.Config.FaultTolerant)
			r.Config.HeaderName = types.StringPointerValue(resp.Config.HeaderName)
			r.Config.HideClientHeaders = types.BoolPointerValue(resp.Config.HideClientHeaders)
			if resp.Config.LimitBy != nil {
				r.Config.LimitBy = types.StringValue(string(*resp.Config.LimitBy))
			} else {
				r.Config.LimitBy = types.StringNull()
			}
			if len(resp.Config.Limits) > 0 {
				r.Config.Limits = make(map[string]types.String, len(resp.Config.Limits))
				for key, value := range resp.Config.Limits {
					result, _ := json.Marshal(value)
					r.Config.Limits[key] = types.StringValue(string(result))
				}
			}
			if resp.Config.Policy != nil {
				r.Config.Policy = types.StringValue(string(*resp.Config.Policy))
			} else {
				r.Config.Policy = types.StringNull()
			}
			if resp.Config.Redis == nil {
				r.Config.Redis = nil
			} else {
				r.Config.Redis = &tfTypes.RateLimitingPluginRedis{}
				r.Config.Redis.Database = types.Int64PointerValue(resp.Config.Redis.Database)
				r.Config.Redis.Host = types.StringPointerValue(resp.Config.Redis.Host)
				r.Config.Redis.Password = types.StringPointerValue(resp.Config.Redis.Password)
				r.Config.Redis.Port = types.Int64PointerValue(resp.Config.Redis.Port)
				r.Config.Redis.ServerName = types.StringPointerValue(resp.Config.Redis.ServerName)
				r.Config.Redis.Ssl = types.BoolPointerValue(resp.Config.Redis.Ssl)
				r.Config.Redis.SslVerify = types.BoolPointerValue(resp.Config.Redis.SslVerify)
				r.Config.Redis.Timeout = types.Int64PointerValue(resp.Config.Redis.Timeout)
				r.Config.Redis.Username = types.StringPointerValue(resp.Config.Redis.Username)
			}
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

	return diags
}
