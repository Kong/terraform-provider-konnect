// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/provider/typeconvert"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *APIProductDocumentDataSourceModel) ToOperationsGetAPIProductDocumentRequest(ctx context.Context) (*operations.GetAPIProductDocumentRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var apiProductID string
	apiProductID = r.APIProductID.ValueString()

	var id string
	id = r.ID.ValueString()

	out := operations.GetAPIProductDocumentRequest{
		APIProductID: apiProductID,
		ID:           id,
	}

	return &out, diags
}

func (r *APIProductDocumentDataSourceModel) RefreshFromSharedAPIProductDocument(ctx context.Context, resp *shared.APIProductDocument) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Content = types.StringValue(resp.Content)
		r.CreatedAt = types.StringValue(typeconvert.TimeToString(resp.CreatedAt))
		r.ID = types.StringValue(resp.ID)
		r.ParentDocumentID = types.StringPointerValue(resp.ParentDocumentID)
		r.Slug = types.StringValue(resp.Slug)
		r.Status = types.StringValue(string(resp.Status))
		r.Title = types.StringValue(resp.Title)
		r.UpdatedAt = types.StringValue(typeconvert.TimeToString(resp.UpdatedAt))
	}

	return diags
}
