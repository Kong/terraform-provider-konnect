// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ServiceUnavailableStatus - The HTTP status code.
type ServiceUnavailableStatus int64

const (
	ServiceUnavailableStatusFiveHundredAndThree ServiceUnavailableStatus = 503
)

func (e ServiceUnavailableStatus) ToPointer() *ServiceUnavailableStatus {
	return &e
}

func (e *ServiceUnavailableStatus) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 503:
		*e = ServiceUnavailableStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServiceUnavailableStatus: %v", v)
	}
}

// ServiceUnavailable - Error response for temporary service unavailability.
type ServiceUnavailable struct {
	// The HTTP status code.
	Status ServiceUnavailableStatus `json:"status"`
	// The error response code.
	Title string `json:"title"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
}

func (o *ServiceUnavailable) GetStatus() ServiceUnavailableStatus {
	if o == nil {
		return ServiceUnavailableStatus(0)
	}
	return o.Status
}

func (o *ServiceUnavailable) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *ServiceUnavailable) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *ServiceUnavailable) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}
