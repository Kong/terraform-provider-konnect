// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginAcmeResourceModel) ToSharedAcmePluginInput() *shared.AcmePluginInput {
	accountEmail := new(string)
	if !r.Config.AccountEmail.IsUnknown() && !r.Config.AccountEmail.IsNull() {
		*accountEmail = r.Config.AccountEmail.ValueString()
	} else {
		accountEmail = nil
	}
	var accountKey *shared.AccountKey
	if r.Config.AccountKey != nil {
		var keyID string
		keyID = r.Config.AccountKey.KeyID.ValueString()

		keySet := new(string)
		if !r.Config.AccountKey.KeySet.IsUnknown() && !r.Config.AccountKey.KeySet.IsNull() {
			*keySet = r.Config.AccountKey.KeySet.ValueString()
		} else {
			keySet = nil
		}
		accountKey = &shared.AccountKey{
			KeyID:  keyID,
			KeySet: keySet,
		}
	}
	allowAnyDomain := new(bool)
	if !r.Config.AllowAnyDomain.IsUnknown() && !r.Config.AllowAnyDomain.IsNull() {
		*allowAnyDomain = r.Config.AllowAnyDomain.ValueBool()
	} else {
		allowAnyDomain = nil
	}
	apiURI := new(string)
	if !r.Config.APIURI.IsUnknown() && !r.Config.APIURI.IsNull() {
		*apiURI = r.Config.APIURI.ValueString()
	} else {
		apiURI = nil
	}
	certType := new(shared.CertType)
	if !r.Config.CertType.IsUnknown() && !r.Config.CertType.IsNull() {
		*certType = shared.CertType(r.Config.CertType.ValueString())
	} else {
		certType = nil
	}
	var domains []string = []string{}
	for _, domainsItem := range r.Config.Domains {
		domains = append(domains, domainsItem.ValueString())
	}
	eabHmacKey := new(string)
	if !r.Config.EabHmacKey.IsUnknown() && !r.Config.EabHmacKey.IsNull() {
		*eabHmacKey = r.Config.EabHmacKey.ValueString()
	} else {
		eabHmacKey = nil
	}
	eabKid := new(string)
	if !r.Config.EabKid.IsUnknown() && !r.Config.EabKid.IsNull() {
		*eabKid = r.Config.EabKid.ValueString()
	} else {
		eabKid = nil
	}
	enableIpv4CommonName := new(bool)
	if !r.Config.EnableIpv4CommonName.IsUnknown() && !r.Config.EnableIpv4CommonName.IsNull() {
		*enableIpv4CommonName = r.Config.EnableIpv4CommonName.ValueBool()
	} else {
		enableIpv4CommonName = nil
	}
	failBackoffMinutes := new(float64)
	if !r.Config.FailBackoffMinutes.IsUnknown() && !r.Config.FailBackoffMinutes.IsNull() {
		*failBackoffMinutes, _ = r.Config.FailBackoffMinutes.ValueBigFloat().Float64()
	} else {
		failBackoffMinutes = nil
	}
	preferredChain := new(string)
	if !r.Config.PreferredChain.IsUnknown() && !r.Config.PreferredChain.IsNull() {
		*preferredChain = r.Config.PreferredChain.ValueString()
	} else {
		preferredChain = nil
	}
	renewThresholdDays := new(float64)
	if !r.Config.RenewThresholdDays.IsUnknown() && !r.Config.RenewThresholdDays.IsNull() {
		*renewThresholdDays, _ = r.Config.RenewThresholdDays.ValueBigFloat().Float64()
	} else {
		renewThresholdDays = nil
	}
	rsaKeySize := new(shared.RsaKeySize)
	if !r.Config.RsaKeySize.IsUnknown() && !r.Config.RsaKeySize.IsNull() {
		*rsaKeySize = shared.RsaKeySize(r.Config.RsaKeySize.ValueInt64())
	} else {
		rsaKeySize = nil
	}
	storage := new(shared.Storage)
	if !r.Config.Storage.IsUnknown() && !r.Config.Storage.IsNull() {
		*storage = shared.Storage(r.Config.Storage.ValueString())
	} else {
		storage = nil
	}
	var storageConfig *shared.StorageConfig
	if r.Config.StorageConfig != nil {
		var consul *shared.Consul
		if r.Config.StorageConfig.Consul != nil {
			host := new(string)
			if !r.Config.StorageConfig.Consul.Host.IsUnknown() && !r.Config.StorageConfig.Consul.Host.IsNull() {
				*host = r.Config.StorageConfig.Consul.Host.ValueString()
			} else {
				host = nil
			}
			https := new(bool)
			if !r.Config.StorageConfig.Consul.HTTPS.IsUnknown() && !r.Config.StorageConfig.Consul.HTTPS.IsNull() {
				*https = r.Config.StorageConfig.Consul.HTTPS.ValueBool()
			} else {
				https = nil
			}
			kvPath := new(string)
			if !r.Config.StorageConfig.Consul.KvPath.IsUnknown() && !r.Config.StorageConfig.Consul.KvPath.IsNull() {
				*kvPath = r.Config.StorageConfig.Consul.KvPath.ValueString()
			} else {
				kvPath = nil
			}
			port := new(int64)
			if !r.Config.StorageConfig.Consul.Port.IsUnknown() && !r.Config.StorageConfig.Consul.Port.IsNull() {
				*port = r.Config.StorageConfig.Consul.Port.ValueInt64()
			} else {
				port = nil
			}
			timeout := new(float64)
			if !r.Config.StorageConfig.Consul.Timeout.IsUnknown() && !r.Config.StorageConfig.Consul.Timeout.IsNull() {
				*timeout, _ = r.Config.StorageConfig.Consul.Timeout.ValueBigFloat().Float64()
			} else {
				timeout = nil
			}
			token := new(string)
			if !r.Config.StorageConfig.Consul.Token.IsUnknown() && !r.Config.StorageConfig.Consul.Token.IsNull() {
				*token = r.Config.StorageConfig.Consul.Token.ValueString()
			} else {
				token = nil
			}
			consul = &shared.Consul{
				Host:    host,
				HTTPS:   https,
				KvPath:  kvPath,
				Port:    port,
				Timeout: timeout,
				Token:   token,
			}
		}
		kong := make(map[string]interface{})
		for kongKey, kongValue := range r.Config.StorageConfig.Kong {
			var kongInst interface{}
			_ = json.Unmarshal([]byte(kongValue.ValueString()), &kongInst)
			kong[kongKey] = kongInst
		}
		var redis *shared.AcmePluginRedis
		if r.Config.StorageConfig.Redis != nil {
			database := new(int64)
			if !r.Config.StorageConfig.Redis.Database.IsUnknown() && !r.Config.StorageConfig.Redis.Database.IsNull() {
				*database = r.Config.StorageConfig.Redis.Database.ValueInt64()
			} else {
				database = nil
			}
			var extraOptions *shared.ExtraOptions
			if r.Config.StorageConfig.Redis.ExtraOptions != nil {
				namespace := new(string)
				if !r.Config.StorageConfig.Redis.ExtraOptions.Namespace.IsUnknown() && !r.Config.StorageConfig.Redis.ExtraOptions.Namespace.IsNull() {
					*namespace = r.Config.StorageConfig.Redis.ExtraOptions.Namespace.ValueString()
				} else {
					namespace = nil
				}
				scanCount := new(float64)
				if !r.Config.StorageConfig.Redis.ExtraOptions.ScanCount.IsUnknown() && !r.Config.StorageConfig.Redis.ExtraOptions.ScanCount.IsNull() {
					*scanCount, _ = r.Config.StorageConfig.Redis.ExtraOptions.ScanCount.ValueBigFloat().Float64()
				} else {
					scanCount = nil
				}
				extraOptions = &shared.ExtraOptions{
					Namespace: namespace,
					ScanCount: scanCount,
				}
			}
			host1 := new(string)
			if !r.Config.StorageConfig.Redis.Host.IsUnknown() && !r.Config.StorageConfig.Redis.Host.IsNull() {
				*host1 = r.Config.StorageConfig.Redis.Host.ValueString()
			} else {
				host1 = nil
			}
			password := new(string)
			if !r.Config.StorageConfig.Redis.Password.IsUnknown() && !r.Config.StorageConfig.Redis.Password.IsNull() {
				*password = r.Config.StorageConfig.Redis.Password.ValueString()
			} else {
				password = nil
			}
			port1 := new(int64)
			if !r.Config.StorageConfig.Redis.Port.IsUnknown() && !r.Config.StorageConfig.Redis.Port.IsNull() {
				*port1 = r.Config.StorageConfig.Redis.Port.ValueInt64()
			} else {
				port1 = nil
			}
			serverName := new(string)
			if !r.Config.StorageConfig.Redis.ServerName.IsUnknown() && !r.Config.StorageConfig.Redis.ServerName.IsNull() {
				*serverName = r.Config.StorageConfig.Redis.ServerName.ValueString()
			} else {
				serverName = nil
			}
			ssl := new(bool)
			if !r.Config.StorageConfig.Redis.Ssl.IsUnknown() && !r.Config.StorageConfig.Redis.Ssl.IsNull() {
				*ssl = r.Config.StorageConfig.Redis.Ssl.ValueBool()
			} else {
				ssl = nil
			}
			sslVerify := new(bool)
			if !r.Config.StorageConfig.Redis.SslVerify.IsUnknown() && !r.Config.StorageConfig.Redis.SslVerify.IsNull() {
				*sslVerify = r.Config.StorageConfig.Redis.SslVerify.ValueBool()
			} else {
				sslVerify = nil
			}
			timeout1 := new(int64)
			if !r.Config.StorageConfig.Redis.Timeout.IsUnknown() && !r.Config.StorageConfig.Redis.Timeout.IsNull() {
				*timeout1 = r.Config.StorageConfig.Redis.Timeout.ValueInt64()
			} else {
				timeout1 = nil
			}
			username := new(string)
			if !r.Config.StorageConfig.Redis.Username.IsUnknown() && !r.Config.StorageConfig.Redis.Username.IsNull() {
				*username = r.Config.StorageConfig.Redis.Username.ValueString()
			} else {
				username = nil
			}
			redis = &shared.AcmePluginRedis{
				Database:     database,
				ExtraOptions: extraOptions,
				Host:         host1,
				Password:     password,
				Port:         port1,
				ServerName:   serverName,
				Ssl:          ssl,
				SslVerify:    sslVerify,
				Timeout:      timeout1,
				Username:     username,
			}
		}
		var shm *shared.Shm
		if r.Config.StorageConfig.Shm != nil {
			shmName := new(string)
			if !r.Config.StorageConfig.Shm.ShmName.IsUnknown() && !r.Config.StorageConfig.Shm.ShmName.IsNull() {
				*shmName = r.Config.StorageConfig.Shm.ShmName.ValueString()
			} else {
				shmName = nil
			}
			shm = &shared.Shm{
				ShmName: shmName,
			}
		}
		var vault *shared.AcmePluginVault
		if r.Config.StorageConfig.Vault != nil {
			authMethod := new(shared.AuthMethod)
			if !r.Config.StorageConfig.Vault.AuthMethod.IsUnknown() && !r.Config.StorageConfig.Vault.AuthMethod.IsNull() {
				*authMethod = shared.AuthMethod(r.Config.StorageConfig.Vault.AuthMethod.ValueString())
			} else {
				authMethod = nil
			}
			authPath := new(string)
			if !r.Config.StorageConfig.Vault.AuthPath.IsUnknown() && !r.Config.StorageConfig.Vault.AuthPath.IsNull() {
				*authPath = r.Config.StorageConfig.Vault.AuthPath.ValueString()
			} else {
				authPath = nil
			}
			authRole := new(string)
			if !r.Config.StorageConfig.Vault.AuthRole.IsUnknown() && !r.Config.StorageConfig.Vault.AuthRole.IsNull() {
				*authRole = r.Config.StorageConfig.Vault.AuthRole.ValueString()
			} else {
				authRole = nil
			}
			host2 := new(string)
			if !r.Config.StorageConfig.Vault.Host.IsUnknown() && !r.Config.StorageConfig.Vault.Host.IsNull() {
				*host2 = r.Config.StorageConfig.Vault.Host.ValueString()
			} else {
				host2 = nil
			}
			https1 := new(bool)
			if !r.Config.StorageConfig.Vault.HTTPS.IsUnknown() && !r.Config.StorageConfig.Vault.HTTPS.IsNull() {
				*https1 = r.Config.StorageConfig.Vault.HTTPS.ValueBool()
			} else {
				https1 = nil
			}
			jwtPath := new(string)
			if !r.Config.StorageConfig.Vault.JwtPath.IsUnknown() && !r.Config.StorageConfig.Vault.JwtPath.IsNull() {
				*jwtPath = r.Config.StorageConfig.Vault.JwtPath.ValueString()
			} else {
				jwtPath = nil
			}
			kvPath1 := new(string)
			if !r.Config.StorageConfig.Vault.KvPath.IsUnknown() && !r.Config.StorageConfig.Vault.KvPath.IsNull() {
				*kvPath1 = r.Config.StorageConfig.Vault.KvPath.ValueString()
			} else {
				kvPath1 = nil
			}
			port2 := new(int64)
			if !r.Config.StorageConfig.Vault.Port.IsUnknown() && !r.Config.StorageConfig.Vault.Port.IsNull() {
				*port2 = r.Config.StorageConfig.Vault.Port.ValueInt64()
			} else {
				port2 = nil
			}
			timeout2 := new(float64)
			if !r.Config.StorageConfig.Vault.Timeout.IsUnknown() && !r.Config.StorageConfig.Vault.Timeout.IsNull() {
				*timeout2, _ = r.Config.StorageConfig.Vault.Timeout.ValueBigFloat().Float64()
			} else {
				timeout2 = nil
			}
			tlsServerName := new(string)
			if !r.Config.StorageConfig.Vault.TLSServerName.IsUnknown() && !r.Config.StorageConfig.Vault.TLSServerName.IsNull() {
				*tlsServerName = r.Config.StorageConfig.Vault.TLSServerName.ValueString()
			} else {
				tlsServerName = nil
			}
			tlsVerify := new(bool)
			if !r.Config.StorageConfig.Vault.TLSVerify.IsUnknown() && !r.Config.StorageConfig.Vault.TLSVerify.IsNull() {
				*tlsVerify = r.Config.StorageConfig.Vault.TLSVerify.ValueBool()
			} else {
				tlsVerify = nil
			}
			token1 := new(string)
			if !r.Config.StorageConfig.Vault.Token.IsUnknown() && !r.Config.StorageConfig.Vault.Token.IsNull() {
				*token1 = r.Config.StorageConfig.Vault.Token.ValueString()
			} else {
				token1 = nil
			}
			vault = &shared.AcmePluginVault{
				AuthMethod:    authMethod,
				AuthPath:      authPath,
				AuthRole:      authRole,
				Host:          host2,
				HTTPS:         https1,
				JwtPath:       jwtPath,
				KvPath:        kvPath1,
				Port:          port2,
				Timeout:       timeout2,
				TLSServerName: tlsServerName,
				TLSVerify:     tlsVerify,
				Token:         token1,
			}
		}
		storageConfig = &shared.StorageConfig{
			Consul: consul,
			Kong:   kong,
			Redis:  redis,
			Shm:    shm,
			Vault:  vault,
		}
	}
	tosAccepted := new(bool)
	if !r.Config.TosAccepted.IsUnknown() && !r.Config.TosAccepted.IsNull() {
		*tosAccepted = r.Config.TosAccepted.ValueBool()
	} else {
		tosAccepted = nil
	}
	config := shared.AcmePluginConfig{
		AccountEmail:         accountEmail,
		AccountKey:           accountKey,
		AllowAnyDomain:       allowAnyDomain,
		APIURI:               apiURI,
		CertType:             certType,
		Domains:              domains,
		EabHmacKey:           eabHmacKey,
		EabKid:               eabKid,
		EnableIpv4CommonName: enableIpv4CommonName,
		FailBackoffMinutes:   failBackoffMinutes,
		PreferredChain:       preferredChain,
		RenewThresholdDays:   renewThresholdDays,
		RsaKeySize:           rsaKeySize,
		Storage:              storage,
		StorageConfig:        storageConfig,
		TosAccepted:          tosAccepted,
	}
	var consumer *shared.AcmePluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.AcmePluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.AcmePluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.AcmePluginConsumerGroup{
			ID: id1,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id2 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id2 = r.ID.ValueString()
	} else {
		id2 = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.AcmePluginOrdering
	if r.Ordering != nil {
		var after *shared.AcmePluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AcmePluginAfter{
				Access: access,
			}
		}
		var before *shared.AcmePluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AcmePluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AcmePluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.AcmePluginProtocols = []shared.AcmePluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AcmePluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AcmePluginRoute
	if r.Route != nil {
		id3 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id3 = r.Route.ID.ValueString()
		} else {
			id3 = nil
		}
		route = &shared.AcmePluginRoute{
			ID: id3,
		}
	}
	var service *shared.AcmePluginService
	if r.Service != nil {
		id4 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id4 = r.Service.ID.ValueString()
		} else {
			id4 = nil
		}
		service = &shared.AcmePluginService{
			ID: id4,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.AcmePluginInput{
		Config:        config,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Enabled:       enabled,
		ID:            id2,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Route:         route,
		Service:       service,
		Tags:          tags,
	}
	return &out
}

func (r *GatewayPluginAcmeResourceModel) RefreshFromSharedAcmePlugin(resp *shared.AcmePlugin) {
	if resp != nil {
		r.Config.AccountEmail = types.StringPointerValue(resp.Config.AccountEmail)
		if resp.Config.AccountKey == nil {
			r.Config.AccountKey = nil
		} else {
			r.Config.AccountKey = &tfTypes.AccountKey{}
			r.Config.AccountKey.KeyID = types.StringValue(resp.Config.AccountKey.KeyID)
			r.Config.AccountKey.KeySet = types.StringPointerValue(resp.Config.AccountKey.KeySet)
		}
		r.Config.AllowAnyDomain = types.BoolPointerValue(resp.Config.AllowAnyDomain)
		r.Config.APIURI = types.StringPointerValue(resp.Config.APIURI)
		if resp.Config.CertType != nil {
			r.Config.CertType = types.StringValue(string(*resp.Config.CertType))
		} else {
			r.Config.CertType = types.StringNull()
		}
		r.Config.Domains = []types.String{}
		for _, v := range resp.Config.Domains {
			r.Config.Domains = append(r.Config.Domains, types.StringValue(v))
		}
		r.Config.EabHmacKey = types.StringPointerValue(resp.Config.EabHmacKey)
		r.Config.EabKid = types.StringPointerValue(resp.Config.EabKid)
		r.Config.EnableIpv4CommonName = types.BoolPointerValue(resp.Config.EnableIpv4CommonName)
		if resp.Config.FailBackoffMinutes != nil {
			r.Config.FailBackoffMinutes = types.NumberValue(big.NewFloat(float64(*resp.Config.FailBackoffMinutes)))
		} else {
			r.Config.FailBackoffMinutes = types.NumberNull()
		}
		r.Config.PreferredChain = types.StringPointerValue(resp.Config.PreferredChain)
		if resp.Config.RenewThresholdDays != nil {
			r.Config.RenewThresholdDays = types.NumberValue(big.NewFloat(float64(*resp.Config.RenewThresholdDays)))
		} else {
			r.Config.RenewThresholdDays = types.NumberNull()
		}
		if resp.Config.RsaKeySize != nil {
			r.Config.RsaKeySize = types.Int64Value(int64(*resp.Config.RsaKeySize))
		} else {
			r.Config.RsaKeySize = types.Int64Null()
		}
		if resp.Config.Storage != nil {
			r.Config.Storage = types.StringValue(string(*resp.Config.Storage))
		} else {
			r.Config.Storage = types.StringNull()
		}
		if resp.Config.StorageConfig == nil {
			r.Config.StorageConfig = nil
		} else {
			r.Config.StorageConfig = &tfTypes.StorageConfig{}
			if resp.Config.StorageConfig.Consul == nil {
				r.Config.StorageConfig.Consul = nil
			} else {
				r.Config.StorageConfig.Consul = &tfTypes.Consul{}
				r.Config.StorageConfig.Consul.Host = types.StringPointerValue(resp.Config.StorageConfig.Consul.Host)
				r.Config.StorageConfig.Consul.HTTPS = types.BoolPointerValue(resp.Config.StorageConfig.Consul.HTTPS)
				r.Config.StorageConfig.Consul.KvPath = types.StringPointerValue(resp.Config.StorageConfig.Consul.KvPath)
				r.Config.StorageConfig.Consul.Port = types.Int64PointerValue(resp.Config.StorageConfig.Consul.Port)
				if resp.Config.StorageConfig.Consul.Timeout != nil {
					r.Config.StorageConfig.Consul.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.StorageConfig.Consul.Timeout)))
				} else {
					r.Config.StorageConfig.Consul.Timeout = types.NumberNull()
				}
				r.Config.StorageConfig.Consul.Token = types.StringPointerValue(resp.Config.StorageConfig.Consul.Token)
			}
			if len(resp.Config.StorageConfig.Kong) > 0 {
				r.Config.StorageConfig.Kong = make(map[string]types.String)
				for key, value := range resp.Config.StorageConfig.Kong {
					result, _ := json.Marshal(value)
					r.Config.StorageConfig.Kong[key] = types.StringValue(string(result))
				}
			}
			if resp.Config.StorageConfig.Redis == nil {
				r.Config.StorageConfig.Redis = nil
			} else {
				r.Config.StorageConfig.Redis = &tfTypes.AcmePluginRedis{}
				r.Config.StorageConfig.Redis.Database = types.Int64PointerValue(resp.Config.StorageConfig.Redis.Database)
				if resp.Config.StorageConfig.Redis.ExtraOptions == nil {
					r.Config.StorageConfig.Redis.ExtraOptions = nil
				} else {
					r.Config.StorageConfig.Redis.ExtraOptions = &tfTypes.ExtraOptions{}
					r.Config.StorageConfig.Redis.ExtraOptions.Namespace = types.StringPointerValue(resp.Config.StorageConfig.Redis.ExtraOptions.Namespace)
					if resp.Config.StorageConfig.Redis.ExtraOptions.ScanCount != nil {
						r.Config.StorageConfig.Redis.ExtraOptions.ScanCount = types.NumberValue(big.NewFloat(float64(*resp.Config.StorageConfig.Redis.ExtraOptions.ScanCount)))
					} else {
						r.Config.StorageConfig.Redis.ExtraOptions.ScanCount = types.NumberNull()
					}
				}
				r.Config.StorageConfig.Redis.Host = types.StringPointerValue(resp.Config.StorageConfig.Redis.Host)
				r.Config.StorageConfig.Redis.Password = types.StringPointerValue(resp.Config.StorageConfig.Redis.Password)
				r.Config.StorageConfig.Redis.Port = types.Int64PointerValue(resp.Config.StorageConfig.Redis.Port)
				r.Config.StorageConfig.Redis.ServerName = types.StringPointerValue(resp.Config.StorageConfig.Redis.ServerName)
				r.Config.StorageConfig.Redis.Ssl = types.BoolPointerValue(resp.Config.StorageConfig.Redis.Ssl)
				r.Config.StorageConfig.Redis.SslVerify = types.BoolPointerValue(resp.Config.StorageConfig.Redis.SslVerify)
				r.Config.StorageConfig.Redis.Timeout = types.Int64PointerValue(resp.Config.StorageConfig.Redis.Timeout)
				r.Config.StorageConfig.Redis.Username = types.StringPointerValue(resp.Config.StorageConfig.Redis.Username)
			}
			if resp.Config.StorageConfig.Shm == nil {
				r.Config.StorageConfig.Shm = nil
			} else {
				r.Config.StorageConfig.Shm = &tfTypes.Shm{}
				r.Config.StorageConfig.Shm.ShmName = types.StringPointerValue(resp.Config.StorageConfig.Shm.ShmName)
			}
			if resp.Config.StorageConfig.Vault == nil {
				r.Config.StorageConfig.Vault = nil
			} else {
				r.Config.StorageConfig.Vault = &tfTypes.AcmePluginVault{}
				if resp.Config.StorageConfig.Vault.AuthMethod != nil {
					r.Config.StorageConfig.Vault.AuthMethod = types.StringValue(string(*resp.Config.StorageConfig.Vault.AuthMethod))
				} else {
					r.Config.StorageConfig.Vault.AuthMethod = types.StringNull()
				}
				r.Config.StorageConfig.Vault.AuthPath = types.StringPointerValue(resp.Config.StorageConfig.Vault.AuthPath)
				r.Config.StorageConfig.Vault.AuthRole = types.StringPointerValue(resp.Config.StorageConfig.Vault.AuthRole)
				r.Config.StorageConfig.Vault.Host = types.StringPointerValue(resp.Config.StorageConfig.Vault.Host)
				r.Config.StorageConfig.Vault.HTTPS = types.BoolPointerValue(resp.Config.StorageConfig.Vault.HTTPS)
				r.Config.StorageConfig.Vault.JwtPath = types.StringPointerValue(resp.Config.StorageConfig.Vault.JwtPath)
				r.Config.StorageConfig.Vault.KvPath = types.StringPointerValue(resp.Config.StorageConfig.Vault.KvPath)
				r.Config.StorageConfig.Vault.Port = types.Int64PointerValue(resp.Config.StorageConfig.Vault.Port)
				if resp.Config.StorageConfig.Vault.Timeout != nil {
					r.Config.StorageConfig.Vault.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.StorageConfig.Vault.Timeout)))
				} else {
					r.Config.StorageConfig.Vault.Timeout = types.NumberNull()
				}
				r.Config.StorageConfig.Vault.TLSServerName = types.StringPointerValue(resp.Config.StorageConfig.Vault.TLSServerName)
				r.Config.StorageConfig.Vault.TLSVerify = types.BoolPointerValue(resp.Config.StorageConfig.Vault.TLSVerify)
				r.Config.StorageConfig.Vault.Token = types.StringPointerValue(resp.Config.StorageConfig.Vault.Token)
			}
		}
		r.Config.TosAccepted = types.BoolPointerValue(resp.Config.TosAccepted)
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
