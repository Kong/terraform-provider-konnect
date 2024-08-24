// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ACLConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *ACLConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type ACL struct {
	// Unix epoch when the resource was created.
	CreatedAt *int64       `json:"created_at,omitempty"`
	Group     *string      `json:"group,omitempty"`
	ID        *string      `json:"id,omitempty"`
	Tags      []string     `json:"tags,omitempty"`
	Consumer  *ACLConsumer `json:"consumer,omitempty"`
}

func (o *ACL) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ACL) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *ACL) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ACL) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ACL) GetConsumer() *ACLConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}
