// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *CloudGatewayCustomDomainDataSourceModel) RefreshFromSharedCustomDomain(resp *shared.CustomDomain) {
	if resp != nil {
		r.CertificateID = types.StringPointerValue(resp.CertificateID)
		r.ControlPlaneGeo = types.StringValue(string(resp.ControlPlaneGeo))
		r.ControlPlaneID = types.StringValue(resp.ControlPlaneID)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.Domain = types.StringValue(resp.Domain)
		r.EntityVersion = types.Int64Value(resp.EntityVersion)
		r.ID = types.StringValue(resp.ID)
		r.SniID = types.StringPointerValue(resp.SniID)
		r.State = types.StringValue(string(resp.State))
		r.StateMetadata.Reason = types.StringPointerValue(resp.StateMetadata.Reason)
		r.StateMetadata.ReportedStatus = types.StringPointerValue(resp.StateMetadata.ReportedStatus)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}
