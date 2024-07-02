// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayDataPlaneClientCertificateResourceModel) ToSharedDataPlaneClientCertificateRequest() *shared.DataPlaneClientCertificateRequest {
	cert := r.Cert.ValueString()
	out := shared.DataPlaneClientCertificateRequest{
		Cert: cert,
	}
	return &out
}

func (r *GatewayDataPlaneClientCertificateResourceModel) RefreshFromSharedDataPlaneClientCertificateItem(resp *shared.DataPlaneClientCertificateItem) {
	if resp != nil {
		r.Cert = types.StringPointerValue(resp.Cert)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}