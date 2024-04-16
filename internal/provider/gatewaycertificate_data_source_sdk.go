// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayCertificateDataSourceModel) RefreshFromSharedCertificate(resp *shared.Certificate) {
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
