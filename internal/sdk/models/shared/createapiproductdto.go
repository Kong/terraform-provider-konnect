// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CreateAPIProductDTO - The request schema to create an API product.
type CreateAPIProductDTO struct {
	// The name of the API product.
	Name string `json:"name"`
	// The description of the API product.
	Description *string `json:"description,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
	// Public labels store information about an entity that can be used for filtering a list of objects.
	//
	// Public labels are intended to store **PUBLIC** metadata.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	PublicLabels map[string]string `json:"public_labels,omitempty"`
	// The list of portal identifiers which this API product should be published to
	PortalIds []string `json:"portal_ids"`
}

func (o *CreateAPIProductDTO) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateAPIProductDTO) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateAPIProductDTO) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *CreateAPIProductDTO) GetPublicLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.PublicLabels
}

func (o *CreateAPIProductDTO) GetPortalIds() []string {
	if o == nil {
		return []string{}
	}
	return o.PortalIds
}
