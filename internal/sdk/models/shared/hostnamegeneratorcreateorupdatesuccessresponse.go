// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HostnameGeneratorCreateOrUpdateSuccessResponse struct {
	// warnings is a list of warning messages to return to the requesting Kuma API clients.
	// Warning messages describe a problem the client making the API request should correct or be aware of.
	//
	Warnings []string `json:"warnings,omitempty"`
}

func (o *HostnameGeneratorCreateOrUpdateSuccessResponse) GetWarnings() []string {
	if o == nil {
		return nil
	}
	return o.Warnings
}
