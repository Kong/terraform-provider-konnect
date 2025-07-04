// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type ConsumerRealmCreateRequest struct {
	Name string `json:"name"`
	// List of control plane ids that are allowed to use the realm. This is used when `allow_all_control_planes` value
	// is set to false.
	//
	AllowedControlPlanes []string `json:"allowed_control_planes,omitempty"`
	// Allow all control planes to use the realm. When this value is set it takes precedence on `allowed_control_planes`.
	//
	AllowAllControlPlanes *bool `json:"allow_all_control_planes,omitempty"`
	// The time in minutes that the Consumer will be cached in memory on a running Gateway if it is
	// successfully loaded from Konnect.
	//
	TTL *int64 `default:"10" json:"ttl"`
	// If a running Gateway triggers a lookup for a Consumer that cannot be authenticated this realm, a `negative_ttl`
	// is set. The Gateway will not try to lookup the Consumer in Konnect for `negative_ttl` minutes.
	//
	// A Consumer will be cached for `negative_ttl` if they do not exist in the Realm, or if the provided credentials
	// are invalid.
	//
	NegativeTTL *int64 `default:"10" json:"negative_ttl"`
	// A list of consumer groups to automatically add to any consumers created within this Realm.
	// If `consumer_groups` are provided on the Consumer object _and_ on the Realm, the Consumer will be placed in all defined consumer groups.
	//
	ConsumerGroups []string `json:"consumer_groups,omitempty"`
}

func (c ConsumerRealmCreateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConsumerRealmCreateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConsumerRealmCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConsumerRealmCreateRequest) GetAllowedControlPlanes() []string {
	if o == nil {
		return nil
	}
	return o.AllowedControlPlanes
}

func (o *ConsumerRealmCreateRequest) GetAllowAllControlPlanes() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAllControlPlanes
}

func (o *ConsumerRealmCreateRequest) GetTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *ConsumerRealmCreateRequest) GetNegativeTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.NegativeTTL
}

func (o *ConsumerRealmCreateRequest) GetConsumerGroups() []string {
	if o == nil {
		return nil
	}
	return o.ConsumerGroups
}
