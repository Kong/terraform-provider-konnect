// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type TLSMetadataHeadersPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *TLSMetadataHeadersPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type TLSMetadataHeadersPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *TLSMetadataHeadersPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type TLSMetadataHeadersPluginOrdering struct {
	After  *TLSMetadataHeadersPluginAfter  `json:"after,omitempty"`
	Before *TLSMetadataHeadersPluginBefore `json:"before,omitempty"`
}

func (o *TLSMetadataHeadersPluginOrdering) GetAfter() *TLSMetadataHeadersPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *TLSMetadataHeadersPluginOrdering) GetBefore() *TLSMetadataHeadersPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type TLSMetadataHeadersPluginConfig struct {
	// Define the HTTP header name used for the SHA1 fingerprint of the client certificate.
	ClientCertFingerprintHeaderName *string `json:"client_cert_fingerprint_header_name,omitempty"`
	// Define the HTTP header name used for the PEM format URL encoded client certificate.
	ClientCertHeaderName *string `json:"client_cert_header_name,omitempty"`
	// Define the HTTP header name used for the issuer DN of the client certificate.
	ClientCertIssuerDnHeaderName *string `json:"client_cert_issuer_dn_header_name,omitempty"`
	// Define the HTTP header name used for the subject DN of the client certificate.
	ClientCertSubjectDnHeaderName *string `json:"client_cert_subject_dn_header_name,omitempty"`
	// Define the HTTP header name used for the serial number of the client certificate.
	ClientSerialHeaderName *string `json:"client_serial_header_name,omitempty"`
	// Enables TLS client certificate metadata values to be injected into HTTP headers.
	InjectClientCertDetails *bool `json:"inject_client_cert_details,omitempty"`
}

func (o *TLSMetadataHeadersPluginConfig) GetClientCertFingerprintHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertFingerprintHeaderName
}

func (o *TLSMetadataHeadersPluginConfig) GetClientCertHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertHeaderName
}

func (o *TLSMetadataHeadersPluginConfig) GetClientCertIssuerDnHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertIssuerDnHeaderName
}

func (o *TLSMetadataHeadersPluginConfig) GetClientCertSubjectDnHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertSubjectDnHeaderName
}

func (o *TLSMetadataHeadersPluginConfig) GetClientSerialHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.ClientSerialHeaderName
}

func (o *TLSMetadataHeadersPluginConfig) GetInjectClientCertDetails() *bool {
	if o == nil {
		return nil
	}
	return o.InjectClientCertDetails
}

type TLSMetadataHeadersPluginProtocols string

const (
	TLSMetadataHeadersPluginProtocolsGrpcs TLSMetadataHeadersPluginProtocols = "grpcs"
	TLSMetadataHeadersPluginProtocolsHTTPS TLSMetadataHeadersPluginProtocols = "https"
	TLSMetadataHeadersPluginProtocolsTLS   TLSMetadataHeadersPluginProtocols = "tls"
)

func (e TLSMetadataHeadersPluginProtocols) ToPointer() *TLSMetadataHeadersPluginProtocols {
	return &e
}
func (e *TLSMetadataHeadersPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpcs":
		fallthrough
	case "https":
		fallthrough
	case "tls":
		*e = TLSMetadataHeadersPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TLSMetadataHeadersPluginProtocols: %v", v)
	}
}

// TLSMetadataHeadersPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
type TLSMetadataHeadersPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *TLSMetadataHeadersPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TLSMetadataHeadersPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type TLSMetadataHeadersPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *TLSMetadataHeadersPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// TLSMetadataHeadersPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type TLSMetadataHeadersPlugin struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                             `json:"enabled,omitempty"`
	ID           *string                           `json:"id,omitempty"`
	InstanceName *string                           `json:"instance_name,omitempty"`
	name         string                            `const:"tls-metadata-headers" json:"name"`
	Ordering     *TLSMetadataHeadersPluginOrdering `json:"ordering,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64                          `json:"updated_at,omitempty"`
	Config    *TLSMetadataHeadersPluginConfig `json:"config,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support tcp and tls.
	Protocols []TLSMetadataHeadersPluginProtocols `json:"protocols,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the route being used.
	Route *TLSMetadataHeadersPluginRoute `json:"route"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *TLSMetadataHeadersPluginService `json:"service"`
}

func (t TLSMetadataHeadersPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSMetadataHeadersPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSMetadataHeadersPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TLSMetadataHeadersPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *TLSMetadataHeadersPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TLSMetadataHeadersPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *TLSMetadataHeadersPlugin) GetName() string {
	return "tls-metadata-headers"
}

func (o *TLSMetadataHeadersPlugin) GetOrdering() *TLSMetadataHeadersPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *TLSMetadataHeadersPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *TLSMetadataHeadersPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TLSMetadataHeadersPlugin) GetConfig() *TLSMetadataHeadersPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *TLSMetadataHeadersPlugin) GetProtocols() []TLSMetadataHeadersPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *TLSMetadataHeadersPlugin) GetRoute() *TLSMetadataHeadersPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *TLSMetadataHeadersPlugin) GetService() *TLSMetadataHeadersPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
