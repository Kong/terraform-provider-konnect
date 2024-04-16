// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

// ClientCertificate - Certificate to be used as client certificate while TLS handshaking to the upstream server.
type ClientCertificate struct {
	ID *string `json:"id,omitempty"`
}

func (o *ClientCertificate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// Protocol - The protocol used to communicate with the upstream.
type Protocol string

const (
	ProtocolGrpc           Protocol = "grpc"
	ProtocolGrpcs          Protocol = "grpcs"
	ProtocolHTTP           Protocol = "http"
	ProtocolHTTPS          Protocol = "https"
	ProtocolTCP            Protocol = "tcp"
	ProtocolTLS            Protocol = "tls"
	ProtocolTLSPassthrough Protocol = "tls_passthrough"
	ProtocolUDP            Protocol = "udp"
	ProtocolWs             Protocol = "ws"
	ProtocolWss            Protocol = "wss"
)

func (e Protocol) ToPointer() *Protocol {
	return &e
}

func (e *Protocol) UnmarshalJSON(data []byte) error {
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
		*e = Protocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Protocol: %v", v)
	}
}

type CreateService struct {
	// Helper field to set `protocol`, `host`, `port` and `path` using a URL. This field is write-only and is not returned in responses.
	URL *string `json:"url,omitempty"`
	// Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).
	CaCertificates []string `json:"ca_certificates,omitempty"`
	// Certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *ClientCertificate `json:"client_certificate,omitempty"`
	// The timeout in milliseconds for establishing a connection to the upstream server.
	ConnectTimeout *int64 `default:"60000" json:"connect_timeout"`
	// Whether the Service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.
	Enabled *bool `default:"true" json:"enabled"`
	// The host of the upstream server. Note that the host value is case sensitive.
	Host *string `json:"host,omitempty"`
	// The Service name.
	Name *string `json:"name,omitempty"`
	// The path to be used in requests to the upstream server.
	Path *string `json:"path,omitempty"`
	// The upstream server port.
	Port *int64 `default:"80" json:"port"`
	// The protocol used to communicate with the upstream.
	Protocol *Protocol `default:"http" json:"protocol"`
	// The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.
	ReadTimeout *int64 `default:"60000" json:"read_timeout"`
	// The number of retries to execute upon failure to proxy.
	Retries *int64 `default:"5" json:"retries"`
	// An optional set of strings associated with the Service for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.
	TLSVerify *bool `json:"tls_verify,omitempty"`
	// Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.
	TLSVerifyDepth *int64 `json:"tls_verify_depth,omitempty"`
	// The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.
	WriteTimeout *int64 `default:"60000" json:"write_timeout"`
}

func (c CreateService) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateService) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateService) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *CreateService) GetCaCertificates() []string {
	if o == nil {
		return nil
	}
	return o.CaCertificates
}

func (o *CreateService) GetClientCertificate() *ClientCertificate {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *CreateService) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *CreateService) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateService) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateService) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateService) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *CreateService) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateService) GetProtocol() *Protocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *CreateService) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *CreateService) GetRetries() *int64 {
	if o == nil {
		return nil
	}
	return o.Retries
}

func (o *CreateService) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateService) GetTLSVerify() *bool {
	if o == nil {
		return nil
	}
	return o.TLSVerify
}

func (o *CreateService) GetTLSVerifyDepth() *int64 {
	if o == nil {
		return nil
	}
	return o.TLSVerifyDepth
}

func (o *CreateService) GetWriteTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.WriteTimeout
}
