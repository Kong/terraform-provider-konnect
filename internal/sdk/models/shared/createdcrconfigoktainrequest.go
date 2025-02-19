// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CreateDcrConfigOktaInRequest - Payload to create an Okta DCR provider.
type CreateDcrConfigOktaInRequest struct {
	// This secret should be copied from your identity provider's settings after you create a client
	// and assign it as the management client for DCR for this developer portal
	//
	DcrToken string `json:"dcr_token"`
}

func (o *CreateDcrConfigOktaInRequest) GetDcrToken() string {
	if o == nil {
		return ""
	}
	return o.DcrToken
}
