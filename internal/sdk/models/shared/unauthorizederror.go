// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UnauthorizedError - standard error
type UnauthorizedError struct {
	Status   any `json:"status"`
	Title    any `json:"title"`
	Type     any `json:"type,omitempty"`
	Instance any `json:"instance"`
	Detail   any `json:"detail"`
}

func (o *UnauthorizedError) GetStatus() any {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnauthorizedError) GetTitle() any {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UnauthorizedError) GetType() any {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UnauthorizedError) GetInstance() any {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *UnauthorizedError) GetDetail() any {
	if o == nil {
		return nil
	}
	return o.Detail
}