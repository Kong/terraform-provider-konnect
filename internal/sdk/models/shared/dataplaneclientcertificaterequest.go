// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DataPlaneClientCertificateRequest struct {
	// JSON escaped string of the certificate.
	Cert string `json:"cert"`
}

func (o *DataPlaneClientCertificateRequest) GetCert() string {
	if o == nil {
		return ""
	}
	return o.Cert
}
