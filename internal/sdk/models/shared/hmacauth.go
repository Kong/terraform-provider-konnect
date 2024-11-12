// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HMACAuthConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *HMACAuthConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type HMACAuth struct {
	Consumer *HMACAuthConsumer `json:"consumer"`
	// Unix epoch when the resource was created.
	CreatedAt *int64   `json:"created_at,omitempty"`
	ID        *string  `json:"id,omitempty"`
	Secret    *string  `json:"secret,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	Username  *string  `json:"username,omitempty"`
}

func (o *HMACAuth) GetConsumer() *HMACAuthConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *HMACAuth) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HMACAuth) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HMACAuth) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *HMACAuth) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *HMACAuth) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
