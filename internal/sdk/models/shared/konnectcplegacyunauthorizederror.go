// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// KonnectCPLegacyUnauthorizedError - standard error
type KonnectCPLegacyUnauthorizedError struct {
	Message any `json:"message,omitempty"`
}

func (o *KonnectCPLegacyUnauthorizedError) GetMessage() any {
	if o == nil {
		return nil
	}
	return o.Message
}
