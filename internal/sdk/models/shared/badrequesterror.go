// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// BadRequestError - standard error
type BadRequestError struct {
	// The HTTP status code of the error. Useful when passing the response
	// body to child properties in a frontend UI. Must be returned as an integer.
	//
	Status int64 `json:"status"`
	// A short, human-readable summary of the problem. It should not
	// change between occurences of a problem, except for localization.
	// Should be provided as "Sentence case" for direct use in the UI.
	//
	Title string `json:"title"`
	// The error type.
	Type *string `json:"type,omitempty"`
	// Used to return the correlation ID back to the user, in the format
	// kong:trace:<correlation_id>. This helps us find the relevant logs
	// when a customer reports an issue.
	//
	Instance string `json:"instance"`
	// A human readable explanation specific to this occurence of the problem.
	// This field may contain request/entity data to help the user understand
	// what went wrong. Enclose variable values in square brackets. Should be
	// provided as "Sentence case" for direct use in the UI.
	//
	Detail string `json:"detail"`
	// invalid parameters
	InvalidParameters []InvalidParameters `json:"invalid_parameters"`
}

func (o *BadRequestError) GetStatus() int64 {
	if o == nil {
		return 0
	}
	return o.Status
}

func (o *BadRequestError) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *BadRequestError) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *BadRequestError) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *BadRequestError) GetDetail() string {
	if o == nil {
		return ""
	}
	return o.Detail
}

func (o *BadRequestError) GetInvalidParameters() []InvalidParameters {
	if o == nil {
		return []InvalidParameters{}
	}
	return o.InvalidParameters
}