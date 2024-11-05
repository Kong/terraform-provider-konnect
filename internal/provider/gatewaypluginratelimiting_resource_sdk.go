// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginRateLimitingResourceModel) ToSharedCreateRateLimitingPlugin() *shared.CreateRateLimitingPlugin {
	var config *shared.CreateRateLimitingPluginConfig
	if r.Config != nil {
		day := new(float64)
		if !r.Config.Day.IsUnknown() && !r.Config.Day.IsNull() {
			*day, _ = r.Config.Day.ValueBigFloat().Float64()
		} else {
			day = nil
		}
		errorCode := new(float64)
		if !r.Config.ErrorCode.IsUnknown() && !r.Config.ErrorCode.IsNull() {
			*errorCode, _ = r.Config.ErrorCode.ValueBigFloat().Float64()
		} else {
			errorCode = nil
		}
		errorMessage := new(string)
		if !r.Config.ErrorMessage.IsUnknown() && !r.Config.ErrorMessage.IsNull() {
			*errorMessage = r.Config.ErrorMessage.ValueString()
		} else {
			errorMessage = nil
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
		hour := new(float64)
		if !r.Config.Hour.IsUnknown() && !r.Config.Hour.IsNull() {
			*hour, _ = r.Config.Hour.ValueBigFloat().Float64()
		} else {
			hour = nil
		}
		limitBy := new(shared.CreateRateLimitingPluginLimitBy)
		if !r.Config.LimitBy.IsUnknown() && !r.Config.LimitBy.IsNull() {
			*limitBy = shared.CreateRateLimitingPluginLimitBy(r.Config.LimitBy.ValueString())
		} else {
			limitBy = nil
		}
		minute := new(float64)
		if !r.Config.Minute.IsUnknown() && !r.Config.Minute.IsNull() {
			*minute, _ = r.Config.Minute.ValueBigFloat().Float64()
		} else {
			minute = nil
		}
		month := new(float64)
		if !r.Config.Month.IsUnknown() && !r.Config.Month.IsNull() {
			*month, _ = r.Config.Month.ValueBigFloat().Float64()
		} else {
			month = nil
		}
		path := new(string)
		if !r.Config.Path.IsUnknown() && !r.Config.Path.IsNull() {
			*path = r.Config.Path.ValueString()
		} else {
			path = nil
		}
		policy := new(shared.CreateRateLimitingPluginPolicy)
		if !r.Config.Policy.IsUnknown() && !r.Config.Policy.IsNull() {
			*policy = shared.CreateRateLimitingPluginPolicy(r.Config.Policy.ValueString())
		} else {
			policy = nil
		}
		var redis *shared.CreateRateLimitingPluginRedis
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
			redis = &shared.CreateRateLimitingPluginRedis{
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
		second := new(float64)
		if !r.Config.Second.IsUnknown() && !r.Config.Second.IsNull() {
			*second, _ = r.Config.Second.ValueBigFloat().Float64()
		} else {
			second = nil
		}
		syncRate := new(float64)
		if !r.Config.SyncRate.IsUnknown() && !r.Config.SyncRate.IsNull() {
			*syncRate, _ = r.Config.SyncRate.ValueBigFloat().Float64()
		} else {
			syncRate = nil
		}
		year := new(float64)
		if !r.Config.Year.IsUnknown() && !r.Config.Year.IsNull() {
			*year, _ = r.Config.Year.ValueBigFloat().Float64()
		} else {
			year = nil
		}
		config = &shared.CreateRateLimitingPluginConfig{
			Day:               day,
			ErrorCode:         errorCode,
			ErrorMessage:      errorMessage,
			FaultTolerant:     faultTolerant,
			HeaderName:        headerName,
			HideClientHeaders: hideClientHeaders,
			Hour:              hour,
			LimitBy:           limitBy,
			Minute:            minute,
			Month:             month,
			Path:              path,
			Policy:            policy,
			Redis:             redis,
			Second:            second,
			SyncRate:          syncRate,
			Year:              year,
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
	var ordering *shared.CreateRateLimitingPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateRateLimitingPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateRateLimitingPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateRateLimitingPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateRateLimitingPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateRateLimitingPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateRateLimitingPluginProtocols = []shared.CreateRateLimitingPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateRateLimitingPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateRateLimitingPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateRateLimitingPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateRateLimitingPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateRateLimitingPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateRateLimitingPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateRateLimitingPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateRateLimitingPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateRateLimitingPluginService{
			ID: id3,
		}
	}
	out := shared.CreateRateLimitingPlugin{
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

func (r *GatewayPluginRateLimitingResourceModel) RefreshFromSharedRateLimitingPlugin(resp *shared.RateLimitingPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateRateLimitingPluginConfig{}
			if resp.Config.Day != nil {
				r.Config.Day = types.NumberValue(big.NewFloat(float64(*resp.Config.Day)))
			} else {
				r.Config.Day = types.NumberNull()
			}
			if resp.Config.ErrorCode != nil {
				r.Config.ErrorCode = types.NumberValue(big.NewFloat(float64(*resp.Config.ErrorCode)))
			} else {
				r.Config.ErrorCode = types.NumberNull()
			}
			r.Config.ErrorMessage = types.StringPointerValue(resp.Config.ErrorMessage)
			r.Config.FaultTolerant = types.BoolPointerValue(resp.Config.FaultTolerant)
			r.Config.HeaderName = types.StringPointerValue(resp.Config.HeaderName)
			r.Config.HideClientHeaders = types.BoolPointerValue(resp.Config.HideClientHeaders)
			if resp.Config.Hour != nil {
				r.Config.Hour = types.NumberValue(big.NewFloat(float64(*resp.Config.Hour)))
			} else {
				r.Config.Hour = types.NumberNull()
			}
			if resp.Config.LimitBy != nil {
				r.Config.LimitBy = types.StringValue(string(*resp.Config.LimitBy))
			} else {
				r.Config.LimitBy = types.StringNull()
			}
			if resp.Config.Minute != nil {
				r.Config.Minute = types.NumberValue(big.NewFloat(float64(*resp.Config.Minute)))
			} else {
				r.Config.Minute = types.NumberNull()
			}
			if resp.Config.Month != nil {
				r.Config.Month = types.NumberValue(big.NewFloat(float64(*resp.Config.Month)))
			} else {
				r.Config.Month = types.NumberNull()
			}
			r.Config.Path = types.StringPointerValue(resp.Config.Path)
			if resp.Config.Policy != nil {
				r.Config.Policy = types.StringValue(string(*resp.Config.Policy))
			} else {
				r.Config.Policy = types.StringNull()
			}
			if resp.Config.Redis == nil {
				r.Config.Redis = nil
			} else {
				r.Config.Redis = &tfTypes.CreateRateLimitingPluginRedis{}
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
			if resp.Config.Second != nil {
				r.Config.Second = types.NumberValue(big.NewFloat(float64(*resp.Config.Second)))
			} else {
				r.Config.Second = types.NumberNull()
			}
			if resp.Config.SyncRate != nil {
				r.Config.SyncRate = types.NumberValue(big.NewFloat(float64(*resp.Config.SyncRate)))
			} else {
				r.Config.SyncRate = types.NumberNull()
			}
			if resp.Config.Year != nil {
				r.Config.Year = types.NumberValue(big.NewFloat(float64(*resp.Config.Year)))
			} else {
				r.Config.Year = types.NumberNull()
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