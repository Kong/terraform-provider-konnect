// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateAPIProductDTO - The request schema for updating an API product.
type UpdateAPIProductDTO struct {
	// The name for the API product.
	Name *string `json:"name,omitempty"`
	// The description of the API product.
	Description *string `json:"description,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Labels are intended to store **INTERNAL** metadata.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]*string `json:"labels,omitempty"`
	// Public labels store information about an entity that can be used for filtering a list of objects.
	//
	// Public labels are intended to store **PUBLIC** metadata.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	PublicLabels map[string]*string `json:"public_labels,omitempty"`
	// The list of portal identifiers which this API product should be published to
	PortalIds []string `json:"portal_ids"`
}

func (o *UpdateAPIProductDTO) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateAPIProductDTO) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateAPIProductDTO) GetLabels() map[string]*string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *UpdateAPIProductDTO) GetPublicLabels() map[string]*string {
	if o == nil {
		return nil
	}
	return o.PublicLabels
}

func (o *UpdateAPIProductDTO) GetPortalIds() []string {
	if o == nil {
		return []string{}
	}
	return o.PortalIds
}
