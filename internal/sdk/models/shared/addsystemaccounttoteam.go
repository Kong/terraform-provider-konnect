// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AddSystemAccountToTeam - The request schema for adding a system account to a team.
type AddSystemAccountToTeam struct {
	// ID of the system account.
	AccountID *string `json:"id,omitempty"`
}

func (o *AddSystemAccountToTeam) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}