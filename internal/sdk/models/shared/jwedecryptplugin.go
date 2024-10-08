// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type JweDecryptPluginConfig struct {
	// The name of the header that is used to set the decrypted value.
	ForwardHeaderName *string `json:"forward_header_name,omitempty"`
	// Denote the name or names of all Key Sets that should be inspected when trying to find a suitable key to decrypt the JWE token.
	KeySets []string `json:"key_sets,omitempty"`
	// The name of the header to look for the JWE token.
	LookupHeaderName *string `json:"lookup_header_name,omitempty"`
	// Defines how the plugin behaves in cases where no token was found in the request. When using `strict` mode, the request requires a token to be present and subsequently raise an error if none could be found.
	Strict *bool `json:"strict,omitempty"`
}

func (o *JweDecryptPluginConfig) GetForwardHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.ForwardHeaderName
}

func (o *JweDecryptPluginConfig) GetKeySets() []string {
	if o == nil {
		return nil
	}
	return o.KeySets
}

func (o *JweDecryptPluginConfig) GetLookupHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.LookupHeaderName
}

func (o *JweDecryptPluginConfig) GetStrict() *bool {
	if o == nil {
		return nil
	}
	return o.Strict
}

type JweDecryptPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *JweDecryptPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JweDecryptPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *JweDecryptPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type JweDecryptPluginOrdering struct {
	After  *JweDecryptPluginAfter  `json:"after,omitempty"`
	Before *JweDecryptPluginBefore `json:"before,omitempty"`
}

func (o *JweDecryptPluginOrdering) GetAfter() *JweDecryptPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *JweDecryptPluginOrdering) GetBefore() *JweDecryptPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type JweDecryptPluginProtocols string

const (
	JweDecryptPluginProtocolsGrpc           JweDecryptPluginProtocols = "grpc"
	JweDecryptPluginProtocolsGrpcs          JweDecryptPluginProtocols = "grpcs"
	JweDecryptPluginProtocolsHTTP           JweDecryptPluginProtocols = "http"
	JweDecryptPluginProtocolsHTTPS          JweDecryptPluginProtocols = "https"
	JweDecryptPluginProtocolsTCP            JweDecryptPluginProtocols = "tcp"
	JweDecryptPluginProtocolsTLS            JweDecryptPluginProtocols = "tls"
	JweDecryptPluginProtocolsTLSPassthrough JweDecryptPluginProtocols = "tls_passthrough"
	JweDecryptPluginProtocolsUDP            JweDecryptPluginProtocols = "udp"
	JweDecryptPluginProtocolsWs             JweDecryptPluginProtocols = "ws"
	JweDecryptPluginProtocolsWss            JweDecryptPluginProtocols = "wss"
)

func (e JweDecryptPluginProtocols) ToPointer() *JweDecryptPluginProtocols {
	return &e
}
func (e *JweDecryptPluginProtocols) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = JweDecryptPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JweDecryptPluginProtocols: %v", v)
	}
}

// JweDecryptPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type JweDecryptPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *JweDecryptPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JweDecryptPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *JweDecryptPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JweDecryptPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type JweDecryptPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *JweDecryptPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JweDecryptPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type JweDecryptPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *JweDecryptPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JweDecryptPlugin struct {
	Config *JweDecryptPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                     `json:"enabled,omitempty"`
	ID           *string                   `json:"id,omitempty"`
	InstanceName *string                   `json:"instance_name,omitempty"`
	name         *string                   `const:"jwe-decrypt" json:"name,omitempty"`
	Ordering     *JweDecryptPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []JweDecryptPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *JweDecryptPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *JweDecryptPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *JweDecryptPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *JweDecryptPluginService `json:"service,omitempty"`
}

func (j JweDecryptPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JweDecryptPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JweDecryptPlugin) GetConfig() *JweDecryptPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *JweDecryptPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *JweDecryptPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *JweDecryptPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JweDecryptPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *JweDecryptPlugin) GetName() *string {
	return types.String("jwe-decrypt")
}

func (o *JweDecryptPlugin) GetOrdering() *JweDecryptPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *JweDecryptPlugin) GetProtocols() []JweDecryptPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *JweDecryptPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *JweDecryptPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *JweDecryptPlugin) GetConsumer() *JweDecryptPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *JweDecryptPlugin) GetConsumerGroup() *JweDecryptPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *JweDecryptPlugin) GetRoute() *JweDecryptPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *JweDecryptPlugin) GetService() *JweDecryptPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
