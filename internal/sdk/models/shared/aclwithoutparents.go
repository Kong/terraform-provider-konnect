// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ACLWithoutParents struct {
	Group *string  `json:"group,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

func (o *ACLWithoutParents) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *ACLWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
