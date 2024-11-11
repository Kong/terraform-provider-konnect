// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayCertificateDataSourceModel) RefreshFromSharedCertificate(resp *shared.Certificate) {
	if resp != nil {
		r.Cert = types.StringValue(resp.Cert)
		r.CertAlt = types.StringPointerValue(resp.CertAlt)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringValue(resp.Key)
		r.KeyAlt = types.StringPointerValue(resp.KeyAlt)
		r.Snis = []types.String{}
		for _, v := range resp.Snis {
			r.Snis = append(r.Snis, types.StringValue(v))
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
