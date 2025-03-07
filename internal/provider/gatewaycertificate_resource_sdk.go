// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayCertificateResourceModel) ToSharedCertificateInput() *shared.CertificateInput {
	var cert string
	cert = r.Cert.ValueString()

	certAlt := new(string)
	if !r.CertAlt.IsUnknown() && !r.CertAlt.IsNull() {
		*certAlt = r.CertAlt.ValueString()
	} else {
		certAlt = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var key string
	key = r.Key.ValueString()

	keyAlt := new(string)
	if !r.KeyAlt.IsUnknown() && !r.KeyAlt.IsNull() {
		*keyAlt = r.KeyAlt.ValueString()
	} else {
		keyAlt = nil
	}
	var snis []string = []string{}
	for _, snisItem := range r.Snis {
		snis = append(snis, snisItem.ValueString())
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.CertificateInput{
		Cert:    cert,
		CertAlt: certAlt,
		ID:      id,
		Key:     key,
		KeyAlt:  keyAlt,
		Snis:    snis,
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
		if resp.Snis != nil {
			r.Snis = make([]types.String, 0, len(resp.Snis))
			for _, v := range resp.Snis {
				r.Snis = append(r.Snis, types.StringValue(v))
			}
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
