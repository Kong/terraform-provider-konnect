// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *SystemAccountTeamResourceModel) ToSharedAddSystemAccountToTeam() *shared.AddSystemAccountToTeam {
	accountID := new(string)
	if !r.AccountID.IsUnknown() && !r.AccountID.IsNull() {
		*accountID = r.AccountID.ValueString()
	} else {
		accountID = nil
	}
	out := shared.AddSystemAccountToTeam{
		AccountID: accountID,
	}
	return &out
}
