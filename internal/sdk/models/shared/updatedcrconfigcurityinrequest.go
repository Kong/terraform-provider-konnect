// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateDcrConfigCurityInRequest - Payload to update a Curity DCR provider.
type UpdateDcrConfigCurityInRequest struct {
	// This ID should be copied from your identity provider's settings after you create a client
	// and assign it as the management client for DCR for this developer portal
	//
	InitialClientID *string `json:"initial_client_id,omitempty"`
	// This secret should be copied from your identity provider's settings after you create a client
	// and assign it as the management client for DCR for this developer portal
	//
	InitialClientSecret *string `json:"initial_client_secret,omitempty"`
}

func (o *UpdateDcrConfigCurityInRequest) GetInitialClientID() *string {
	if o == nil {
		return nil
	}
	return o.InitialClientID
}

func (o *UpdateDcrConfigCurityInRequest) GetInitialClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.InitialClientSecret
}
