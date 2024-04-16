// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Status - The HTTP status code.
type Status int64

const (
	StatusFiveHundred Status = 500
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 500:
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// InternalServerError - The error response object.
type InternalServerError struct {
	// The HTTP status code.
	Status Status `json:"status"`
	// The error response code.
	Title string `json:"title"`
	// The Konnect traceback code
	Instance string `json:"instance"`
	// Details about the error.
	Detail *string `json:"detail,omitempty"`
}

func (o *InternalServerError) GetStatus() Status {
	if o == nil {
		return Status(0)
	}
	return o.Status
}

func (o *InternalServerError) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *InternalServerError) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *InternalServerError) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}
