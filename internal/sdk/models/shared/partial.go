// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type PartialRedisEEClusterNodes struct {
	// Cluster node IP.
	IP *string `json:"ip,omitempty"`
	// Cluster node port.
	Port *int64 `json:"port,omitempty"`
}

func (o *PartialRedisEEClusterNodes) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *PartialRedisEEClusterNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type PartialRedisEESentinelNodes struct {
	// Sentinel node hostname.
	Host *string `json:"host,omitempty"`
	// Sentinel node port.
	Port *int64 `json:"port,omitempty"`
}

func (o *PartialRedisEESentinelNodes) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *PartialRedisEESentinelNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

// PartialRedisEEConfig - Redis-EE configuration
type PartialRedisEEConfig struct {
	// Maximum retry attempts for redirection.
	ClusterMaxRedirections *int64 `json:"cluster_max_redirections,omitempty"`
	// Cluster addresses for Redis connections using the `redis` strategy.
	ClusterNodes []PartialRedisEEClusterNodes `json:"cluster_nodes,omitempty"`
	// Connect timeout.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// If the connection to Redis is proxied, e.g., Envoy.
	ConnectionIsProxied *bool `json:"connection_is_proxied,omitempty"`
	// Database index.
	Database *int64 `json:"database,omitempty"`
	// Redis host.
	Host *string `json:"host,omitempty"`
	// Limits the total number of opened connections for a pool.
	KeepaliveBacklog *int64 `json:"keepalive_backlog,omitempty"`
	// Size limit for cosocket connection pool per worker process.
	KeepalivePoolSize *int64 `json:"keepalive_pool_size,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// The port is only used when the host is set.
	Port *int64 `json:"port,omitempty"`
	// Read timeout.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// Send timeout.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// Sentinel master to use for Redis connections. Defining this implies using Redis Sentinel.
	SentinelMaster *string `json:"sentinel_master,omitempty"`
	// Sentinel addresses for Redis connections using the `redis` strategy. Array must have at least 1 element.
	SentinelNodes []PartialRedisEESentinelNodes `json:"sentinel_nodes,omitempty"`
	// Sentinel password to authenticate with a Redis Sentinel instance.
	SentinelPassword *string `json:"sentinel_password,omitempty"`
	// Sentinel role to use for Redis connections when `redis` strategy is used, implies using Redis Sentinel.
	SentinelRole *string `json:"sentinel_role,omitempty"`
	// Sentinel username to authenticate with a Redis Sentinel instance. Requires Redis v6.2.0+.
	SentinelUsername *string `json:"sentinel_username,omitempty"`
	// Server name for SSL verification.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. Requires Redis v6.0.0+.
	Username *string `json:"username,omitempty"`
}

func (o *PartialRedisEEConfig) GetClusterMaxRedirections() *int64 {
	if o == nil {
		return nil
	}
	return o.ClusterMaxRedirections
}

func (o *PartialRedisEEConfig) GetClusterNodes() []PartialRedisEEClusterNodes {
	if o == nil {
		return nil
	}
	return o.ClusterNodes
}

func (o *PartialRedisEEConfig) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *PartialRedisEEConfig) GetConnectionIsProxied() *bool {
	if o == nil {
		return nil
	}
	return o.ConnectionIsProxied
}

func (o *PartialRedisEEConfig) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *PartialRedisEEConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *PartialRedisEEConfig) GetKeepaliveBacklog() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepaliveBacklog
}

func (o *PartialRedisEEConfig) GetKeepalivePoolSize() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepalivePoolSize
}

func (o *PartialRedisEEConfig) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *PartialRedisEEConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *PartialRedisEEConfig) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *PartialRedisEEConfig) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

func (o *PartialRedisEEConfig) GetSentinelMaster() *string {
	if o == nil {
		return nil
	}
	return o.SentinelMaster
}

func (o *PartialRedisEEConfig) GetSentinelNodes() []PartialRedisEESentinelNodes {
	if o == nil {
		return nil
	}
	return o.SentinelNodes
}

func (o *PartialRedisEEConfig) GetSentinelPassword() *string {
	if o == nil {
		return nil
	}
	return o.SentinelPassword
}

func (o *PartialRedisEEConfig) GetSentinelRole() *string {
	if o == nil {
		return nil
	}
	return o.SentinelRole
}

func (o *PartialRedisEEConfig) GetSentinelUsername() *string {
	if o == nil {
		return nil
	}
	return o.SentinelUsername
}

func (o *PartialRedisEEConfig) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *PartialRedisEEConfig) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *PartialRedisEEConfig) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *PartialRedisEEConfig) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type PartialRedisEE struct {
	// Redis-EE configuration
	Config PartialRedisEEConfig `json:"config"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	type_     string   `const:"redis-ee" json:"type"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (p PartialRedisEE) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PartialRedisEE) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PartialRedisEE) GetConfig() PartialRedisEEConfig {
	if o == nil {
		return PartialRedisEEConfig{}
	}
	return o.Config
}

func (o *PartialRedisEE) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PartialRedisEE) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PartialRedisEE) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PartialRedisEE) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PartialRedisEE) GetType() string {
	return "redis-ee"
}

func (o *PartialRedisEE) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// PartialRedisCEConfig - Redis-CE configuration
type PartialRedisCEConfig struct {
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// Redis host.
	Host *string `json:"host,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// Redis port.
	Port *int64 `json:"port,omitempty"`
	// Server name for SSL verification.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// Connection timeout.
	Timeout *int64 `json:"timeout,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. Requires Redis v6.0.0+.
	Username *string `json:"username,omitempty"`
}

func (o *PartialRedisCEConfig) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *PartialRedisCEConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *PartialRedisCEConfig) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *PartialRedisCEConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *PartialRedisCEConfig) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *PartialRedisCEConfig) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *PartialRedisCEConfig) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *PartialRedisCEConfig) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *PartialRedisCEConfig) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type PartialRedisCE struct {
	// Redis-CE configuration
	Config PartialRedisCEConfig `json:"config"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	type_     string   `const:"redis-ce" json:"type"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (p PartialRedisCE) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PartialRedisCE) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PartialRedisCE) GetConfig() PartialRedisCEConfig {
	if o == nil {
		return PartialRedisCEConfig{}
	}
	return o.Config
}

func (o *PartialRedisCE) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PartialRedisCE) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PartialRedisCE) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PartialRedisCE) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PartialRedisCE) GetType() string {
	return "redis-ce"
}

func (o *PartialRedisCE) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type PartialType string

const (
	PartialTypeRedisCe PartialType = "redis-ce"
	PartialTypeRedisEe PartialType = "redis-ee"
)

type Partial struct {
	PartialRedisCE *PartialRedisCE `queryParam:"inline"`
	PartialRedisEE *PartialRedisEE `queryParam:"inline"`

	Type PartialType
}

func CreatePartialRedisCe(redisCe PartialRedisCE) Partial {
	typ := PartialTypeRedisCe

	return Partial{
		PartialRedisCE: &redisCe,
		Type:           typ,
	}
}

func CreatePartialRedisEe(redisEe PartialRedisEE) Partial {
	typ := PartialTypeRedisEe

	return Partial{
		PartialRedisEE: &redisEe,
		Type:           typ,
	}
}

func (u *Partial) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Type string `json:"type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "redis-ce":
		partialRedisCE := new(PartialRedisCE)
		if err := utils.UnmarshalJSON(data, &partialRedisCE, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == redis-ce) type PartialRedisCE within Partial: %w", string(data), err)
		}

		u.PartialRedisCE = partialRedisCE
		u.Type = PartialTypeRedisCe
		return nil
	case "redis-ee":
		partialRedisEE := new(PartialRedisEE)
		if err := utils.UnmarshalJSON(data, &partialRedisEE, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == redis-ee) type PartialRedisEE within Partial: %w", string(data), err)
		}

		u.PartialRedisEE = partialRedisEE
		u.Type = PartialTypeRedisEe
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Partial", string(data))
}

func (u Partial) MarshalJSON() ([]byte, error) {
	if u.PartialRedisCE != nil {
		return utils.MarshalJSON(u.PartialRedisCE, "", true)
	}

	if u.PartialRedisEE != nil {
		return utils.MarshalJSON(u.PartialRedisEE, "", true)
	}

	return nil, errors.New("could not marshal union type Partial: all fields are null")
}
