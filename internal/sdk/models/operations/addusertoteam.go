// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

var AddUserToTeamServerList = []string{
	"https://global.api.konghq.com/",
}

type AddUserToTeamRequest struct {
	// ID of the team.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	// The request schema for adding a user to a team.
	AddUserToTeam *shared.AddUserToTeam `request:"mediaType=application/json"`
}

func (o *AddUserToTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *AddUserToTeamRequest) GetAddUserToTeam() *shared.AddUserToTeam {
	if o == nil {
		return nil
	}
	return o.AddUserToTeam
}

type AddUserToTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Not Found
	NotFoundError *shared.NotFoundError
	// Conflict
	ConflictError *shared.ConflictError
}

func (o *AddUserToTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddUserToTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddUserToTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddUserToTeamResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *AddUserToTeamResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}

func (o *AddUserToTeamResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}
