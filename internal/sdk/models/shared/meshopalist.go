// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// MeshOPAList - List
type MeshOPAList struct {
	Items []MeshOPAItem `json:"items,omitempty"`
	// The total number of entities
	Total *float64 `json:"total,omitempty"`
	// URL to the next page
	Next *string `json:"next,omitempty"`
}

func (o *MeshOPAList) GetItems() []MeshOPAItem {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *MeshOPAList) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *MeshOPAList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}