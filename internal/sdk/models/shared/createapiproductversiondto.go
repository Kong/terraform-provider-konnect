// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PublishStatus - The publish status of the API product version. Applies publish status to all related portal product versions. This field is deprecated: Use [PortalProductVersion.publish_status](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type PublishStatus string

const (
	PublishStatusUnpublished PublishStatus = "unpublished"
	PublishStatusPublished   PublishStatus = "published"
)

func (e PublishStatus) ToPointer() *PublishStatus {
	return &e
}
func (e *PublishStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unpublished":
		fallthrough
	case "published":
		*e = PublishStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PublishStatus: %v", v)
	}
}

// CreateAPIProductVersionDTO - The request schema to create a version of an API product.
// Note that the `publish_status` and `deprecated` fields are deprecated:  Use [PortalProductVersion.publish_status](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
type CreateAPIProductVersionDTO struct {
	// The version name of the API product version.
	Name string `json:"name"`
	// The publish status of the API product version. Applies publish status to all related portal product versions. This field is deprecated: Use [PortalProductVersion.publish_status](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	PublishStatus *PublishStatus `json:"publish_status,omitempty"`
	// Indicates if the version of the API product is deprecated. Applies deprecation or removes deprecation from all related portal product versions. This field is deprecated: Use [PortalProductVersion.deprecated](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Deprecated     *bool                  `json:"deprecated,omitempty"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
}

func (o *CreateAPIProductVersionDTO) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateAPIProductVersionDTO) GetPublishStatus() *PublishStatus {
	if o == nil {
		return nil
	}
	return o.PublishStatus
}

func (o *CreateAPIProductVersionDTO) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *CreateAPIProductVersionDTO) GetGatewayService() *GatewayServicePayload {
	if o == nil {
		return nil
	}
	return o.GatewayService
}