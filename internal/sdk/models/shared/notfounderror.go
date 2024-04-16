// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// NotFoundError - standard error
type NotFoundError struct {
	Status   interface{} `json:"status"`
	Title    interface{} `json:"title"`
	Type     interface{} `json:"type,omitempty"`
	Instance interface{} `json:"instance"`
	Detail   interface{} `json:"detail"`
}

func (o *NotFoundError) GetStatus() interface{} {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *NotFoundError) GetTitle() interface{} {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *NotFoundError) GetType() interface{} {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *NotFoundError) GetInstance() interface{} {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *NotFoundError) GetDetail() interface{} {
	if o == nil {
		return nil
	}
	return o.Detail
}
