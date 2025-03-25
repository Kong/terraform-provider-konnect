// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

// AccountKey - The private key associated with the account.
type AccountKey struct {
	// The Key ID.
	KeyID string `json:"key_id"`
	// The ID of the key set to associate the Key ID with.
	KeySet *string `json:"key_set,omitempty"`
}

func (o *AccountKey) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *AccountKey) GetKeySet() *string {
	if o == nil {
		return nil
	}
	return o.KeySet
}

// CertType - The certificate type to create. The possible values are `rsa` for RSA certificate or `ecc` for EC certificate.
type CertType string

const (
	CertTypeEcc CertType = "ecc"
	CertTypeRsa CertType = "rsa"
)

func (e CertType) ToPointer() *CertType {
	return &e
}
func (e *CertType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ecc":
		fallthrough
	case "rsa":
		*e = CertType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CertType: %v", v)
	}
}

// RsaKeySize - RSA private key size for the certificate. The possible values are 2048, 3072, or 4096.
type RsaKeySize int64

const (
	RsaKeySizeTwoThousandAndFortyEight   RsaKeySize = 2048
	RsaKeySizeThreeThousandAndSeventyTwo RsaKeySize = 3072
	RsaKeySizeFourThousandAndNinetySix   RsaKeySize = 4096
)

func (e RsaKeySize) ToPointer() *RsaKeySize {
	return &e
}
func (e *RsaKeySize) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 2048:
		fallthrough
	case 3072:
		fallthrough
	case 4096:
		*e = RsaKeySize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RsaKeySize: %v", v)
	}
}

// Storage - The backend storage type to use. In DB-less mode and Konnect, `kong` storage is unavailable. In hybrid mode and Konnect, `shm` storage is unavailable. `shm` storage does not persist during Kong restarts and does not work for Kong running on different machines, so consider using one of `kong`, `redis`, `consul`, or `vault` in production.
type Storage string

const (
	StorageConsul Storage = "consul"
	StorageKong   Storage = "kong"
	StorageRedis  Storage = "redis"
	StorageShm    Storage = "shm"
	StorageVault  Storage = "vault"
)

func (e Storage) ToPointer() *Storage {
	return &e
}
func (e *Storage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consul":
		fallthrough
	case "kong":
		fallthrough
	case "redis":
		fallthrough
	case "shm":
		fallthrough
	case "vault":
		*e = Storage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Storage: %v", v)
	}
}

type Consul struct {
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Boolean representation of https.
	HTTPS *bool `json:"https,omitempty"`
	// KV prefix path.
	KvPath *string `json:"kv_path,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// Timeout in milliseconds.
	Timeout *float64 `json:"timeout,omitempty"`
	// Consul ACL token.
	Token *string `json:"token,omitempty"`
}

func (o *Consul) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *Consul) GetHTTPS() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPS
}

func (o *Consul) GetKvPath() *string {
	if o == nil {
		return nil
	}
	return o.KvPath
}

func (o *Consul) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Consul) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *Consul) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// ExtraOptions - Custom ACME Redis options
type ExtraOptions struct {
	// A namespace to prepend to all keys stored in Redis.
	Namespace *string `json:"namespace,omitempty"`
	// The number of keys to return in Redis SCAN calls.
	ScanCount *float64 `json:"scan_count,omitempty"`
}

func (o *ExtraOptions) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *ExtraOptions) GetScanCount() *float64 {
	if o == nil {
		return nil
	}
	return o.ScanCount
}

type AcmePluginRedis struct {
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// Custom ACME Redis options
	ExtraOptions *ExtraOptions `json:"extra_options,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	Timeout *int64 `json:"timeout,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *AcmePluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *AcmePluginRedis) GetExtraOptions() *ExtraOptions {
	if o == nil {
		return nil
	}
	return o.ExtraOptions
}

func (o *AcmePluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *AcmePluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *AcmePluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *AcmePluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *AcmePluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *AcmePluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *AcmePluginRedis) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *AcmePluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type Shm struct {
	// Name of shared memory zone used for Kong API gateway storage
	ShmName *string `json:"shm_name,omitempty"`
}

func (o *Shm) GetShmName() *string {
	if o == nil {
		return nil
	}
	return o.ShmName
}

// AcmePluginAuthMethod - Auth Method, default to token, can be 'token' or 'kubernetes'.
type AcmePluginAuthMethod string

const (
	AcmePluginAuthMethodKubernetes AcmePluginAuthMethod = "kubernetes"
	AcmePluginAuthMethodToken      AcmePluginAuthMethod = "token"
)

func (e AcmePluginAuthMethod) ToPointer() *AcmePluginAuthMethod {
	return &e
}
func (e *AcmePluginAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kubernetes":
		fallthrough
	case "token":
		*e = AcmePluginAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AcmePluginAuthMethod: %v", v)
	}
}

type AcmePluginVault struct {
	// Auth Method, default to token, can be 'token' or 'kubernetes'.
	AuthMethod *AcmePluginAuthMethod `json:"auth_method,omitempty"`
	// Vault's authentication path to use.
	AuthPath *string `json:"auth_path,omitempty"`
	// The role to try and assign.
	AuthRole *string `json:"auth_role,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Boolean representation of https.
	HTTPS *bool `json:"https,omitempty"`
	// The path to the JWT.
	JwtPath *string `json:"jwt_path,omitempty"`
	// KV prefix path.
	KvPath *string `json:"kv_path,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// Timeout in milliseconds.
	Timeout *float64 `json:"timeout,omitempty"`
	// SNI used in request, default to host if omitted.
	TLSServerName *string `json:"tls_server_name,omitempty"`
	// Turn on TLS verification.
	TLSVerify *bool `json:"tls_verify,omitempty"`
	// Consul ACL token.
	Token *string `json:"token,omitempty"`
}

func (o *AcmePluginVault) GetAuthMethod() *AcmePluginAuthMethod {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *AcmePluginVault) GetAuthPath() *string {
	if o == nil {
		return nil
	}
	return o.AuthPath
}

func (o *AcmePluginVault) GetAuthRole() *string {
	if o == nil {
		return nil
	}
	return o.AuthRole
}

func (o *AcmePluginVault) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *AcmePluginVault) GetHTTPS() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPS
}

func (o *AcmePluginVault) GetJwtPath() *string {
	if o == nil {
		return nil
	}
	return o.JwtPath
}

func (o *AcmePluginVault) GetKvPath() *string {
	if o == nil {
		return nil
	}
	return o.KvPath
}

func (o *AcmePluginVault) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *AcmePluginVault) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *AcmePluginVault) GetTLSServerName() *string {
	if o == nil {
		return nil
	}
	return o.TLSServerName
}

func (o *AcmePluginVault) GetTLSVerify() *bool {
	if o == nil {
		return nil
	}
	return o.TLSVerify
}

func (o *AcmePluginVault) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

type StorageConfig struct {
	Consul *Consul          `json:"consul,omitempty"`
	Kong   map[string]any   `json:"kong,omitempty"`
	Redis  *AcmePluginRedis `json:"redis,omitempty"`
	Shm    *Shm             `json:"shm,omitempty"`
	Vault  *AcmePluginVault `json:"vault,omitempty"`
}

func (o *StorageConfig) GetConsul() *Consul {
	if o == nil {
		return nil
	}
	return o.Consul
}

func (o *StorageConfig) GetKong() map[string]any {
	if o == nil {
		return nil
	}
	return o.Kong
}

func (o *StorageConfig) GetRedis() *AcmePluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *StorageConfig) GetShm() *Shm {
	if o == nil {
		return nil
	}
	return o.Shm
}

func (o *StorageConfig) GetVault() *AcmePluginVault {
	if o == nil {
		return nil
	}
	return o.Vault
}

type AcmePluginConfig struct {
	// The account identifier. Can be reused in a different plugin instance.
	AccountEmail *string `json:"account_email,omitempty"`
	// The private key associated with the account.
	AccountKey *AccountKey `json:"account_key,omitempty"`
	// If set to `true`, the plugin allows all domains and ignores any values in the `domains` list.
	AllowAnyDomain *bool `json:"allow_any_domain,omitempty"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	APIURI *string `json:"api_uri,omitempty"`
	// The certificate type to create. The possible values are `rsa` for RSA certificate or `ecc` for EC certificate.
	CertType *CertType `json:"cert_type,omitempty"`
	// An array of strings representing hosts. A valid host is a string containing one or more labels separated by periods, with at most one wildcard label ('*')
	Domains []string `json:"domains,omitempty"`
	// External account binding (EAB) base64-encoded URL string of the HMAC key. You usually don't need to set this unless it is explicitly required by the CA.
	EabHmacKey *string `json:"eab_hmac_key,omitempty"`
	// External account binding (EAB) key id. You usually don't need to set this unless it is explicitly required by the CA.
	EabKid *string `json:"eab_kid,omitempty"`
	// A boolean value that controls whether to include the IPv4 address in the common name field of generated certificates.
	EnableIpv4CommonName *bool `json:"enable_ipv4_common_name,omitempty"`
	// Minutes to wait for each domain that fails to create a certificate. This applies to both a
	// new certificate and a renewal certificate.
	FailBackoffMinutes *float64 `json:"fail_backoff_minutes,omitempty"`
	// A string value that specifies the preferred certificate chain to use when generating certificates.
	PreferredChain *string `json:"preferred_chain,omitempty"`
	// Days remaining to renew the certificate before it expires.
	RenewThresholdDays *float64 `json:"renew_threshold_days,omitempty"`
	// RSA private key size for the certificate. The possible values are 2048, 3072, or 4096.
	RsaKeySize *RsaKeySize `json:"rsa_key_size,omitempty"`
	// The backend storage type to use. In DB-less mode and Konnect, `kong` storage is unavailable. In hybrid mode and Konnect, `shm` storage is unavailable. `shm` storage does not persist during Kong restarts and does not work for Kong running on different machines, so consider using one of `kong`, `redis`, `consul`, or `vault` in production.
	Storage       *Storage       `json:"storage,omitempty"`
	StorageConfig *StorageConfig `json:"storage_config,omitempty"`
	// If you are using Let's Encrypt, you must set this to `true` to agree the terms of service.
	TosAccepted *bool `json:"tos_accepted,omitempty"`
}

func (o *AcmePluginConfig) GetAccountEmail() *string {
	if o == nil {
		return nil
	}
	return o.AccountEmail
}

func (o *AcmePluginConfig) GetAccountKey() *AccountKey {
	if o == nil {
		return nil
	}
	return o.AccountKey
}

func (o *AcmePluginConfig) GetAllowAnyDomain() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAnyDomain
}

func (o *AcmePluginConfig) GetAPIURI() *string {
	if o == nil {
		return nil
	}
	return o.APIURI
}

func (o *AcmePluginConfig) GetCertType() *CertType {
	if o == nil {
		return nil
	}
	return o.CertType
}

func (o *AcmePluginConfig) GetDomains() []string {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *AcmePluginConfig) GetEabHmacKey() *string {
	if o == nil {
		return nil
	}
	return o.EabHmacKey
}

func (o *AcmePluginConfig) GetEabKid() *string {
	if o == nil {
		return nil
	}
	return o.EabKid
}

func (o *AcmePluginConfig) GetEnableIpv4CommonName() *bool {
	if o == nil {
		return nil
	}
	return o.EnableIpv4CommonName
}

func (o *AcmePluginConfig) GetFailBackoffMinutes() *float64 {
	if o == nil {
		return nil
	}
	return o.FailBackoffMinutes
}

func (o *AcmePluginConfig) GetPreferredChain() *string {
	if o == nil {
		return nil
	}
	return o.PreferredChain
}

func (o *AcmePluginConfig) GetRenewThresholdDays() *float64 {
	if o == nil {
		return nil
	}
	return o.RenewThresholdDays
}

func (o *AcmePluginConfig) GetRsaKeySize() *RsaKeySize {
	if o == nil {
		return nil
	}
	return o.RsaKeySize
}

func (o *AcmePluginConfig) GetStorage() *Storage {
	if o == nil {
		return nil
	}
	return o.Storage
}

func (o *AcmePluginConfig) GetStorageConfig() *StorageConfig {
	if o == nil {
		return nil
	}
	return o.StorageConfig
}

func (o *AcmePluginConfig) GetTosAccepted() *bool {
	if o == nil {
		return nil
	}
	return o.TosAccepted
}

type AcmePluginProtocols string

const (
	AcmePluginProtocolsGrpc  AcmePluginProtocols = "grpc"
	AcmePluginProtocolsGrpcs AcmePluginProtocols = "grpcs"
	AcmePluginProtocolsHTTP  AcmePluginProtocols = "http"
	AcmePluginProtocolsHTTPS AcmePluginProtocols = "https"
)

func (e AcmePluginProtocols) ToPointer() *AcmePluginProtocols {
	return &e
}
func (e *AcmePluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		*e = AcmePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AcmePluginProtocols: %v", v)
	}
}

// AcmePlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AcmePlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"acme" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64           `json:"updated_at,omitempty"`
	Config    AcmePluginConfig `json:"config"`
	// A set of strings representing HTTP protocols.
	Protocols []AcmePluginProtocols `json:"protocols,omitempty"`
}

func (a AcmePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AcmePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AcmePlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AcmePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AcmePlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AcmePlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AcmePlugin) GetName() string {
	return "acme"
}

func (o *AcmePlugin) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AcmePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AcmePlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AcmePlugin) GetConfig() AcmePluginConfig {
	if o == nil {
		return AcmePluginConfig{}
	}
	return o.Config
}

func (o *AcmePlugin) GetProtocols() []AcmePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

// AcmePluginInput - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type AcmePluginInput struct {
	// Whether the plugin is applied.
	Enabled      *bool             `json:"enabled,omitempty"`
	ID           *string           `json:"id,omitempty"`
	InstanceName *string           `json:"instance_name,omitempty"`
	name         string            `const:"acme" json:"name"`
	Ordering     map[string]string `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags   []string         `json:"tags,omitempty"`
	Config AcmePluginConfig `json:"config"`
	// A set of strings representing HTTP protocols.
	Protocols []AcmePluginProtocols `json:"protocols,omitempty"`
}

func (a AcmePluginInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AcmePluginInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AcmePluginInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AcmePluginInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AcmePluginInput) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AcmePluginInput) GetName() string {
	return "acme"
}

func (o *AcmePluginInput) GetOrdering() map[string]string {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AcmePluginInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AcmePluginInput) GetConfig() AcmePluginConfig {
	if o == nil {
		return AcmePluginConfig{}
	}
	return o.Config
}

func (o *AcmePluginInput) GetProtocols() []AcmePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}
