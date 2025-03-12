// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
	"time"
)

// Team - The team object contains information about a group of users.
type Team struct {
	// The team ID.
	ID *string `json:"id,omitempty"`
	// The name of the team.
	Name *string `json:"name,omitempty"`
	// The team description in Konnect.
	Description *string `json:"description,omitempty"`
	// Returns True if a user belongs to a `system_team`. System teams are teams that can manage Konnect objects, like "Organization Admin", or "Service"
	SystemTeam *bool `json:"system_team,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]*string `json:"labels,omitempty"`
	// A Unix timestamp representation of team creation.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A Unix timestamp representation of the most recent change to the team object in Konnect.
	//
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (t Team) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Team) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Team) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Team) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Team) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Team) GetSystemTeam() *bool {
	if o == nil {
		return nil
	}
	return o.SystemTeam
}

func (o *Team) GetLabels() map[string]*string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *Team) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Team) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
