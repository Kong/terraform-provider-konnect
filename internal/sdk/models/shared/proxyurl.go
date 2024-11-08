// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ProxyURL - Proxy URL associated with reaching the data-planes connected to a control-plane.
type ProxyURL struct {
	// Hostname of the proxy URL.
	Host string `json:"host"`
	// Port of the proxy URL.
	Port int64 `json:"port"`
	// Protocol of the proxy URL.
	Protocol string `json:"protocol"`
}

func (o *ProxyURL) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *ProxyURL) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ProxyURL) GetProtocol() string {
	if o == nil {
		return ""
	}
	return o.Protocol
}