// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// NotFoundError - standard error
type NotFoundError struct {
	Detail   any `json:"detail"`
	Instance any `json:"instance"`
	Status   any `json:"status"`
	Title    any `json:"title"`
	Type     any `json:"type,omitempty"`
}

func (o *NotFoundError) GetDetail() any {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *NotFoundError) GetInstance() any {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *NotFoundError) GetStatus() any {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *NotFoundError) GetTitle() any {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *NotFoundError) GetType() any {
	if o == nil {
		return nil
	}
	return o.Type
}
