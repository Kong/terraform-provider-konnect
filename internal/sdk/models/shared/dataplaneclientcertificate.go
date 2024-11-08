// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DataPlaneClientCertificateItem struct {
	// Unique ID of the certificate entity.
	ID *string `json:"id,omitempty"`
	// Date certificate was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Date certificate was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// JSON escaped string of the certificate.
	Cert *string `json:"cert,omitempty"`
}

func (o *DataPlaneClientCertificateItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DataPlaneClientCertificateItem) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *DataPlaneClientCertificateItem) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *DataPlaneClientCertificateItem) GetCert() *string {
	if o == nil {
		return nil
	}
	return o.Cert
}

// DataPlaneClientCertificate - Response body for retrieving a dp-client-certificate.
type DataPlaneClientCertificate struct {
	Item *DataPlaneClientCertificateItem `json:"item,omitempty"`
}

func (o *DataPlaneClientCertificate) GetItem() *DataPlaneClientCertificateItem {
	if o == nil {
		return nil
	}
	return o.Item
}
