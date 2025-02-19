// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// DcrConfigHTTPInResponse - A DCR provider configuration for HTTP
type DcrConfigHTTPInResponse struct {
	DcrBaseURL string `json:"dcr_base_url"`
	// This flag disables all the event-hooks on the application flow for the DCR provider.
	DisableEventHooks *bool `json:"disable_event_hooks,omitempty"`
	// This flag disable the refresh-secret endpoint on the application flow for the DCR provider.
	DisableRefreshSecret *bool `json:"disable_refresh_secret,omitempty"`
}

func (o *DcrConfigHTTPInResponse) GetDcrBaseURL() string {
	if o == nil {
		return ""
	}
	return o.DcrBaseURL
}

func (o *DcrConfigHTTPInResponse) GetDisableEventHooks() *bool {
	if o == nil {
		return nil
	}
	return o.DisableEventHooks
}

func (o *DcrConfigHTTPInResponse) GetDisableRefreshSecret() *bool {
	if o == nil {
		return nil
	}
	return o.DisableRefreshSecret
}
