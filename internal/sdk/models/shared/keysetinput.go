// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type KeySetInput struct {
	ID   *string  `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

func (o *KeySetInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeySetInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *KeySetInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
