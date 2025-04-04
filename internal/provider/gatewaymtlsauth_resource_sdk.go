// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *GatewayMTLSAuthResourceModel) ToSharedMTLSAuthWithoutParents() *shared.MTLSAuthWithoutParents {
	var caCertificate *shared.MTLSAuthWithoutParentsCaCertificate
	if r.CaCertificate != nil {
		id := new(string)
		if !r.CaCertificate.ID.IsUnknown() && !r.CaCertificate.ID.IsNull() {
			*id = r.CaCertificate.ID.ValueString()
		} else {
			id = nil
		}
		caCertificate = &shared.MTLSAuthWithoutParentsCaCertificate{
			ID: id,
		}
	}
	var consumer *shared.MTLSAuthWithoutParentsConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.MTLSAuthWithoutParentsConsumer{
			ID: id1,
		}
	}
	createdAt := new(int64)
	if !r.CreatedAt.IsUnknown() && !r.CreatedAt.IsNull() {
		*createdAt = r.CreatedAt.ValueInt64()
	} else {
		createdAt = nil
	}
	id2 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id2 = r.ID.ValueString()
	} else {
		id2 = nil
	}
	var subjectName string
	subjectName = r.SubjectName.ValueString()

	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.MTLSAuthWithoutParents{
		CaCertificate: caCertificate,
		Consumer:      consumer,
		CreatedAt:     createdAt,
		ID:            id2,
		SubjectName:   subjectName,
		Tags:          tags,
	}
	return &out
}

func (r *GatewayMTLSAuthResourceModel) RefreshFromSharedMTLSAuth(resp *shared.MTLSAuth) {
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
