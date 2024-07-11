// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SlugConflict struct {
	Instance string  `json:"instance"`
	Status   float64 `json:"status"`
	Title    string  `json:"title"`
	Type     *string `json:"type,omitempty"`
}

func (o *SlugConflict) GetInstance() string {
	if o == nil {
		return ""
	}
	return o.Instance
}

func (o *SlugConflict) GetStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.Status
}

func (o *SlugConflict) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *SlugConflict) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
