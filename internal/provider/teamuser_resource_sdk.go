// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
)

func (r *TeamUserResourceModel) ToSharedAddUserToTeam() *shared.AddUserToTeam {
	var userID string
	userID = r.UserID.ValueString()

	out := shared.AddUserToTeam{
		UserID: userID,
	}
	return &out
}
