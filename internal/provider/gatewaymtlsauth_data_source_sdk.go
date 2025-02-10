// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayMTLSAuthDataSourceModel) RefreshFromSharedMTLSAuth(resp *shared.MTLSAuth) {
	if resp != nil {
		if resp.CaCertificate == nil {
			r.CaCertificate = nil
		} else {
			r.CaCertificate = &tfTypes.ACLWithoutParentsConsumer{}
			r.CaCertificate.ID = types.StringPointerValue(resp.CaCertificate.ID)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.SubjectName = types.StringValue(resp.SubjectName)
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
