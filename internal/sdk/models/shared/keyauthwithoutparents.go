// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type KeyAuthWithoutParentsConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthWithoutParentsConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KeyAuthWithoutParents struct {
	Consumer *KeyAuthWithoutParentsConsumer `json:"consumer"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Key       *string  `json:"key,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	// key-auth ttl in seconds
	TTL *int64 `json:"ttl,omitempty"`
}

func (o *KeyAuthWithoutParents) GetConsumer() *KeyAuthWithoutParentsConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *KeyAuthWithoutParents) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *KeyAuthWithoutParents) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeyAuthWithoutParents) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *KeyAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *KeyAuthWithoutParents) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}
