// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// MeshTCPRouteList - List
type MeshTCPRouteList struct {
	Items []MeshTCPRouteItem `json:"items,omitempty"`
	// The total number of entities
	Total *float64 `json:"total,omitempty"`
	// URL to the next page
	Next *string `json:"next,omitempty"`
}

func (o *MeshTCPRouteList) GetItems() []MeshTCPRouteItem {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *MeshTCPRouteList) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *MeshTCPRouteList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}