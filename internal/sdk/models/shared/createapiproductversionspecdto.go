// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateAPIProductVersionSpecDTO - The request schema to create a specification for a version of an API product.
type CreateAPIProductVersionSpecDTO struct {
	// The base64 encoded contents of the API product version specification
	Content string `json:"content"`
	// The name of the API product version specification
	Name string `json:"name"`
}

func (o *CreateAPIProductVersionSpecDTO) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *CreateAPIProductVersionSpecDTO) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
