// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AddUserToTeam struct {
	// The user ID for the user being added to a team.
	UserID string `json:"id"`
}

func (o *AddUserToTeam) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}