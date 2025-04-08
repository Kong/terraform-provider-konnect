// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ID - Returns entities that exact match any of the comma-delimited phrases in the filter string.
type ID struct {
	Eq  *string `queryParam:"name=eq"`
	Oeq *string `queryParam:"name=oeq"`
}

func (o *ID) GetEq() *string {
	if o == nil {
		return nil
	}
	return o.Eq
}

func (o *ID) GetOeq() *string {
	if o == nil {
		return nil
	}
	return o.Oeq
}

// Name - Filters on the given string field value by exact match inequality.
type Name struct {
	Eq       *string `queryParam:"name=eq"`
	Contains *string `queryParam:"name=contains"`
	Neq      *string `queryParam:"name=neq"`
	Oeq      *string `queryParam:"name=oeq"`
}

func (o *Name) GetEq() *string {
	if o == nil {
		return nil
	}
	return o.Eq
}

func (o *Name) GetContains() *string {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *Name) GetNeq() *string {
	if o == nil {
		return nil
	}
	return o.Neq
}

func (o *Name) GetOeq() *string {
	if o == nil {
		return nil
	}
	return o.Oeq
}

// ClusterType - Filters on the given string field value by exact match inequality.
type ClusterType struct {
	Eq  *string `queryParam:"name=eq"`
	Neq *string `queryParam:"name=neq"`
	Oeq *string `queryParam:"name=oeq"`
}

func (o *ClusterType) GetEq() *string {
	if o == nil {
		return nil
	}
	return o.Eq
}

func (o *ClusterType) GetNeq() *string {
	if o == nil {
		return nil
	}
	return o.Neq
}

func (o *ClusterType) GetOeq() *string {
	if o == nil {
		return nil
	}
	return o.Oeq
}

type ControlPlaneFilterParameters struct {
	// Returns entities that exact match any of the comma-delimited phrases in the filter string.
	ID *ID `queryParam:"name=id"`
	// Filters on the given string field value by exact match inequality.
	Name *Name `queryParam:"name=name"`
	// Filters on the given string field value by exact match inequality.
	ClusterType *ClusterType `queryParam:"name=cluster_type"`
	// Filter by a boolean value (true/false).
	CloudGateway *bool `queryParam:"name=cloud_gateway"`
}

func (o *ControlPlaneFilterParameters) GetID() *ID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ControlPlaneFilterParameters) GetName() *Name {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ControlPlaneFilterParameters) GetClusterType() *ClusterType {
	if o == nil {
		return nil
	}
	return o.ClusterType
}

func (o *ControlPlaneFilterParameters) GetCloudGateway() *bool {
	if o == nil {
		return nil
	}
	return o.CloudGateway
}
