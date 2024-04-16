// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayCertificateResourceModel) ToSharedCreateCertificate() *shared.CreateCertificate {
	cert := r.Cert.ValueString()
	certAlt := new(string)
	if !r.CertAlt.IsUnknown() && !r.CertAlt.IsNull() {
		*certAlt = r.CertAlt.ValueString()
	} else {
		certAlt = nil
	}
	key := r.Key.ValueString()
	keyAlt := new(string)
	if !r.KeyAlt.IsUnknown() && !r.KeyAlt.IsNull() {
		*keyAlt = r.KeyAlt.ValueString()
	} else {
		keyAlt = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.CreateCertificate{
		Cert:    cert,
		CertAlt: certAlt,
		Key:     key,
		KeyAlt:  keyAlt,
		Tags:    tags,
	}
	return &out
}

func (r *GatewayCertificateResourceModel) RefreshFromSharedCertificate(resp *shared.Certificate) {
	if resp != nil {
		r.Cert = types.StringValue(resp.Cert)
		r.CertAlt = types.StringPointerValue(resp.CertAlt)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringValue(resp.Key)
		r.KeyAlt = types.StringPointerValue(resp.KeyAlt)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
