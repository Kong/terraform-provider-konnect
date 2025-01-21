// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type MeshGatewayItemProtocolType string

const (
	MeshGatewayItemProtocolTypeStr     MeshGatewayItemProtocolType = "str"
	MeshGatewayItemProtocolTypeInteger MeshGatewayItemProtocolType = "integer"
)

// MeshGatewayItemProtocol - Protocol specifies the network protocol this listener expects to receive.
type MeshGatewayItemProtocol struct {
	Str     *string
	Integer *int64

	Type MeshGatewayItemProtocolType
}

func CreateMeshGatewayItemProtocolStr(str string) MeshGatewayItemProtocol {
	typ := MeshGatewayItemProtocolTypeStr

	return MeshGatewayItemProtocol{
		Str:  &str,
		Type: typ,
	}
}

func CreateMeshGatewayItemProtocolInteger(integer int64) MeshGatewayItemProtocol {
	typ := MeshGatewayItemProtocolTypeInteger

	return MeshGatewayItemProtocol{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *MeshGatewayItemProtocol) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = MeshGatewayItemProtocolTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = MeshGatewayItemProtocolTypeInteger
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MeshGatewayItemProtocol", string(data))
}

func (u MeshGatewayItemProtocol) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type MeshGatewayItemProtocol: all fields are null")
}

// Resources is used to specify listener-specific resource settings.
type Resources struct {
	ConnectionLimit *int64 `json:"connection_limit,omitempty"`
}

func (o *Resources) GetConnectionLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectionLimit
}

// Certificates - DataSource defines the source of bytes to use.
type Certificates struct {
	// Types that are assignable to Type:
	//
	// 	*DataSource_Secret
	// 	*DataSource_File
	// 	*DataSource_Inline
	// 	*DataSource_InlineString
	Type any `json:"Type"`
}

func (o *Certificates) GetType() any {
	if o == nil {
		return nil
	}
	return o.Type
}

type MeshGatewayItemModeType string

const (
	MeshGatewayItemModeTypeStr     MeshGatewayItemModeType = "str"
	MeshGatewayItemModeTypeInteger MeshGatewayItemModeType = "integer"
)

// MeshGatewayItemMode - Mode defines the TLS behavior for the TLS session initiated
// by the client.
type MeshGatewayItemMode struct {
	Str     *string
	Integer *int64

	Type MeshGatewayItemModeType
}

func CreateMeshGatewayItemModeStr(str string) MeshGatewayItemMode {
	typ := MeshGatewayItemModeTypeStr

	return MeshGatewayItemMode{
		Str:  &str,
		Type: typ,
	}
}

func CreateMeshGatewayItemModeInteger(integer int64) MeshGatewayItemMode {
	typ := MeshGatewayItemModeTypeInteger

	return MeshGatewayItemMode{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *MeshGatewayItemMode) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = MeshGatewayItemModeTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = MeshGatewayItemModeTypeInteger
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MeshGatewayItemMode", string(data))
}

func (u MeshGatewayItemMode) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type MeshGatewayItemMode: all fields are null")
}

// MeshGatewayItemOptions - Options should eventually configure how TLS is configured. This
// is where cipher suite and version configuration can be specified,
// client certificates enforced, and so on.
type MeshGatewayItemOptions struct {
}

// MeshGatewayItemTLS - TLS is the TLS configuration for the Listener. This field
// is required if the Protocol field is "HTTPS" or "TLS" and
// ignored otherwise.
type MeshGatewayItemTLS struct {
	// Certificates is an array of datasources that contain TLS
	// certificates and private keys.  Each datasource must contain a
	// sequence of PEM-encoded objects. The server certificate and private
	// key are required, but additional certificates are allowed and will
	// be added to the certificate chain.  The server certificate must
	// be the first certificate in the datasource.
	//
	// When multiple certificate datasources are configured, they must have
	// different key types. In practice, this means that one datasource
	// should contain an RSA key and certificate, and the other an
	// ECDSA key and certificate.
	Certificates []Certificates `json:"certificates,omitempty"`
	// Mode defines the TLS behavior for the TLS session initiated
	// by the client.
	Mode *MeshGatewayItemMode `json:"mode,omitempty"`
	// Options should eventually configure how TLS is configured. This
	// is where cipher suite and version configuration can be specified,
	// client certificates enforced, and so on.
	Options *MeshGatewayItemOptions `json:"options,omitempty"`
}

func (o *MeshGatewayItemTLS) GetCertificates() []Certificates {
	if o == nil {
		return nil
	}
	return o.Certificates
}

func (o *MeshGatewayItemTLS) GetMode() *MeshGatewayItemMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *MeshGatewayItemTLS) GetOptions() *MeshGatewayItemOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

type Listeners struct {
	// CrossMesh enables traffic to flow to this listener only from other
	// meshes.
	CrossMesh *bool `json:"crossMesh,omitempty"`
	// Hostname specifies the virtual hostname to match for protocol types that
	// define this concept. When unspecified, "", or `*`, all hostnames are
	// matched. This field can be omitted for protocols that don't require
	// hostname based matching.
	Hostname *string `json:"hostname,omitempty"`
	// Port is the network port. Multiple listeners may use the
	// same port, subject to the Listener compatibility rules.
	Port *int64 `json:"port,omitempty"`
	// Protocol specifies the network protocol this listener expects to receive.
	Protocol *MeshGatewayItemProtocol `json:"protocol,omitempty"`
	// Resources is used to specify listener-specific resource settings.
	Resources *Resources `json:"resources,omitempty"`
	// Tags specifies a unique combination of tags that routes can use
	// to match themselves to this listener.
	//
	// When matching routes to listeners, the control plane constructs a
	// set of matching tags for each listener by forming the union of the
	// gateway tags and the listener tags. A route will be attached to the
	// listener if all of the route's tags are preset in the matching tags
	Tags map[string]string `json:"tags,omitempty"`
	// TLS is the TLS configuration for the Listener. This field
	// is required if the Protocol field is "HTTPS" or "TLS" and
	// ignored otherwise.
	TLS *MeshGatewayItemTLS `json:"tls,omitempty"`
}

func (o *Listeners) GetCrossMesh() *bool {
	if o == nil {
		return nil
	}
	return o.CrossMesh
}

func (o *Listeners) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *Listeners) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Listeners) GetProtocol() *MeshGatewayItemProtocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *Listeners) GetResources() *Resources {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *Listeners) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Listeners) GetTLS() *MeshGatewayItemTLS {
	if o == nil {
		return nil
	}
	return o.TLS
}

// Conf - The desired configuration of the MeshGateway.
type Conf struct {
	// Listeners define logical endpoints that are bound on this MeshGateway's
	// address(es).
	Listeners []Listeners `json:"listeners,omitempty"`
}

func (o *Conf) GetListeners() []Listeners {
	if o == nil {
		return nil
	}
	return o.Listeners
}

// Selectors - Selector defines structure for selecting tags for given dataplane
type Selectors struct {
	// Tags to match, can be used for both source and destinations
	Match map[string]string `json:"match,omitempty"`
}

func (o *Selectors) GetMatch() map[string]string {
	if o == nil {
		return nil
	}
	return o.Match
}

type MeshGatewayItem struct {
	// The desired configuration of the MeshGateway.
	Conf   *Conf             `json:"conf,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	Mesh   string            `json:"mesh"`
	Name   string            `json:"name"`
	// Selectors is a list of selectors that are used to match builtin
	// gateway dataplanes that will receive this MeshGateway configuration.
	Selectors []Selectors `json:"selectors,omitempty"`
	// Tags is the set of tags common to all of the gateway's listeners.
	//
	// This field must not include a `kuma.io/service` tag (the service is always
	// defined on the dataplanes).
	Tags map[string]string `json:"tags,omitempty"`
	Type string            `json:"type"`
}

func (o *MeshGatewayItem) GetConf() *Conf {
	if o == nil {
		return nil
	}
	return o.Conf
}

func (o *MeshGatewayItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshGatewayItem) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *MeshGatewayItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshGatewayItem) GetSelectors() []Selectors {
	if o == nil {
		return nil
	}
	return o.Selectors
}

func (o *MeshGatewayItem) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *MeshGatewayItem) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
