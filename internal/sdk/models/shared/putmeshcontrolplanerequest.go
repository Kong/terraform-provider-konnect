// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PutMeshControlPlaneRequest - a payload to update the whole control plane
type PutMeshControlPlaneRequest struct {
	// The name of the control plane.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Labels to facilitate tagged search on control planes. Keys must be of length 1-63 characters.
	Labels map[string]string `json:"labels,omitempty"`
}

func (o *PutMeshControlPlaneRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PutMeshControlPlaneRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PutMeshControlPlaneRequest) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}
