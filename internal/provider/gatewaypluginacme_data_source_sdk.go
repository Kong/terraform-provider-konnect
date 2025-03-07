// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginAcmeDataSourceModel) RefreshFromSharedAcmePlugin(resp *shared.AcmePlugin) {
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
		r.Config.Domains = make([]types.String, 0, len(resp.Config.Domains))
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
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
