// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CertificateInput - A certificate object represents a public certificate, and can be optionally paired with the corresponding private key. These objects are used by Kong to handle SSL/TLS termination for encrypted requests, or for use as a trusted CA store when validating peer certificate of client/service. Certificates are optionally associated with SNI objects to tie a cert/key pair to one or more hostnames. If intermediate certificates are required in addition to the main certificate, they should be concatenated together into one string according to the following order: main certificate on the top, followed by any intermediates.
type CertificateInput struct {
	// PEM-encoded public certificate chain of the SSL key pair. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	Cert string `json:"cert"`
	// PEM-encoded public certificate chain of the alternate SSL key pair. This should only be set if you have both RSA and ECDSA types of certificate available and would like Kong to prefer serving using ECDSA certs when client advertises support for it. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	CertAlt *string `json:"cert_alt,omitempty"`
	ID      *string `json:"id,omitempty"`
	// PEM-encoded private key of the SSL key pair. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	Key string `json:"key"`
	// PEM-encoded private key of the alternate SSL key pair. This should only be set if you have both RSA and ECDSA types of certificate available and would like Kong to prefer serving using ECDSA certs when client advertises support for it. This field is _referenceable_, which means it can be securely stored as a [secret](/gateway/latest/plan-and-deploy/security/secrets-management/getting-started) in a vault. References must follow a [specific format](/gateway/latest/plan-and-deploy/security/secrets-management/reference-format).
	KeyAlt *string  `json:"key_alt,omitempty"`
	Snis   []string `json:"snis,omitempty"`
	// An optional set of strings associated with the Certificate for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
}

func (o *CertificateInput) GetCert() string {
	if o == nil {
		return ""
	}
	return o.Cert
}

func (o *CertificateInput) GetCertAlt() *string {
	if o == nil {
		return nil
	}
	return o.CertAlt
}

func (o *CertificateInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CertificateInput) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *CertificateInput) GetKeyAlt() *string {
	if o == nil {
		return nil
	}
	return o.KeyAlt
}

func (o *CertificateInput) GetSnis() []string {
	if o == nil {
		return nil
	}
	return o.Snis
}

func (o *CertificateInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
