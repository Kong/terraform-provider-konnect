// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"net/http"
)

type UpdatePortalTeamRequest struct {
	// ID of the team.
	TeamID string `pathParam:"style=simple,explode=false,name=teamId"`
	// ID of the portal.
	PortalID string `pathParam:"style=simple,explode=false,name=portalId"`
	// Update a team in a portal.
	PortalUpdateTeamRequest *shared.PortalUpdateTeamRequest `request:"mediaType=application/json"`
}

func (o *UpdatePortalTeamRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *UpdatePortalTeamRequest) GetPortalID() string {
	if o == nil {
		return ""
	}
	return o.PortalID
}

func (o *UpdatePortalTeamRequest) GetPortalUpdateTeamRequest() *shared.PortalUpdateTeamRequest {
	if o == nil {
		return nil
	}
	return o.PortalUpdateTeamRequest
}

type UpdatePortalTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Details about a team of developers in a portal.
	PortalTeamResponse *shared.PortalTeamResponse
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *UpdatePortalTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePortalTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePortalTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePortalTeamResponse) GetPortalTeamResponse() *shared.PortalTeamResponse {
	if o == nil {
		return nil
	}
	return o.PortalTeamResponse
}

func (o *UpdatePortalTeamResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *UpdatePortalTeamResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *UpdatePortalTeamResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
