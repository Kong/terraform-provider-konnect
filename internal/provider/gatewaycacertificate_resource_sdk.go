// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayCACertificateResourceModel) ToSharedCACertificateInput() *shared.CACertificateInput {
	var cert string
	cert = r.Cert.ValueString()

	certDigest := new(string)
	if !r.CertDigest.IsUnknown() && !r.CertDigest.IsNull() {
		*certDigest = r.CertDigest.ValueString()
	} else {
		certDigest = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.CACertificateInput{
		Cert:       cert,
		CertDigest: certDigest,
		ID:         id,
		Tags:       tags,
	}
	return &out
}

func (r *GatewayCACertificateResourceModel) RefreshFromSharedCACertificate(resp *shared.CACertificate) {
	if resp != nil {
		r.Cert = types.StringValue(resp.Cert)
		r.CertDigest = types.StringPointerValue(resp.CertDigest)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
