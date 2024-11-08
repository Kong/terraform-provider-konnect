// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateJweDecryptPluginConfig struct {
	// The name of the header that is used to set the decrypted value.
	ForwardHeaderName *string `json:"forward_header_name,omitempty"`
	// Denote the name or names of all Key Sets that should be inspected when trying to find a suitable key to decrypt the JWE token.
	KeySets []string `json:"key_sets,omitempty"`
	// The name of the header to look for the JWE token.
	LookupHeaderName *string `json:"lookup_header_name,omitempty"`
	// Defines how the plugin behaves in cases where no token was found in the request. When using `strict` mode, the request requires a token to be present and subsequently raise an error if none could be found.
	Strict *bool `json:"strict,omitempty"`
}

func (o *CreateJweDecryptPluginConfig) GetForwardHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.ForwardHeaderName
}

func (o *CreateJweDecryptPluginConfig) GetKeySets() []string {
	if o == nil {
		return nil
	}
	return o.KeySets
}

func (o *CreateJweDecryptPluginConfig) GetLookupHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.LookupHeaderName
}

func (o *CreateJweDecryptPluginConfig) GetStrict() *bool {
	if o == nil {
		return nil
	}
	return o.Strict
}

type CreateJweDecryptPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateJweDecryptPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateJweDecryptPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateJweDecryptPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateJweDecryptPluginOrdering struct {
	After  *CreateJweDecryptPluginAfter  `json:"after,omitempty"`
	Before *CreateJweDecryptPluginBefore `json:"before,omitempty"`
}

func (o *CreateJweDecryptPluginOrdering) GetAfter() *CreateJweDecryptPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateJweDecryptPluginOrdering) GetBefore() *CreateJweDecryptPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateJweDecryptPluginProtocols string

const (
	CreateJweDecryptPluginProtocolsGrpc           CreateJweDecryptPluginProtocols = "grpc"
	CreateJweDecryptPluginProtocolsGrpcs          CreateJweDecryptPluginProtocols = "grpcs"
	CreateJweDecryptPluginProtocolsHTTP           CreateJweDecryptPluginProtocols = "http"
	CreateJweDecryptPluginProtocolsHTTPS          CreateJweDecryptPluginProtocols = "https"
	CreateJweDecryptPluginProtocolsTCP            CreateJweDecryptPluginProtocols = "tcp"
	CreateJweDecryptPluginProtocolsTLS            CreateJweDecryptPluginProtocols = "tls"
	CreateJweDecryptPluginProtocolsTLSPassthrough CreateJweDecryptPluginProtocols = "tls_passthrough"
	CreateJweDecryptPluginProtocolsUDP            CreateJweDecryptPluginProtocols = "udp"
	CreateJweDecryptPluginProtocolsWs             CreateJweDecryptPluginProtocols = "ws"
	CreateJweDecryptPluginProtocolsWss            CreateJweDecryptPluginProtocols = "wss"
)

func (e CreateJweDecryptPluginProtocols) ToPointer() *CreateJweDecryptPluginProtocols {
	return &e
}
func (e *CreateJweDecryptPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateJweDecryptPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateJweDecryptPluginProtocols: %v", v)
	}
}

// CreateJweDecryptPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateJweDecryptPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJweDecryptPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateJweDecryptPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJweDecryptPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJweDecryptPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateJweDecryptPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJweDecryptPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateJweDecryptPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateJweDecryptPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateJweDecryptPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateJweDecryptPlugin struct {
	Config *CreateJweDecryptPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                           `json:"enabled,omitempty"`
	InstanceName *string                         `json:"instance_name,omitempty"`
	name         *string                         `const:"jwe-decrypt" json:"name,omitempty"`
	Ordering     *CreateJweDecryptPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateJweDecryptPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateJweDecryptPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateJweDecryptPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateJweDecryptPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateJweDecryptPluginService `json:"service,omitempty"`
}

func (c CreateJweDecryptPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateJweDecryptPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateJweDecryptPlugin) GetConfig() *CreateJweDecryptPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateJweDecryptPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateJweDecryptPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateJweDecryptPlugin) GetName() *string {
	return types.String("jwe-decrypt")
}

func (o *CreateJweDecryptPlugin) GetOrdering() *CreateJweDecryptPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateJweDecryptPlugin) GetProtocols() []CreateJweDecryptPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateJweDecryptPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateJweDecryptPlugin) GetConsumer() *CreateJweDecryptPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateJweDecryptPlugin) GetConsumerGroup() *CreateJweDecryptPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateJweDecryptPlugin) GetRoute() *CreateJweDecryptPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateJweDecryptPlugin) GetService() *CreateJweDecryptPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}