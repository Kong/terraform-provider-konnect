// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateAPIProductVersionDTOPublishStatus - The publish status of the API product version. Applies publish status to all related portal product versions. This field is deprecated: Use [PortalProductVersion.publish_status](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type CreateAPIProductVersionDTOPublishStatus string

const (
	CreateAPIProductVersionDTOPublishStatusUnpublished CreateAPIProductVersionDTOPublishStatus = "unpublished"
	CreateAPIProductVersionDTOPublishStatusPublished   CreateAPIProductVersionDTOPublishStatus = "published"
)

func (e CreateAPIProductVersionDTOPublishStatus) ToPointer() *CreateAPIProductVersionDTOPublishStatus {
	return &e
}
func (e *CreateAPIProductVersionDTOPublishStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unpublished":
		fallthrough
	case "published":
		*e = CreateAPIProductVersionDTOPublishStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAPIProductVersionDTOPublishStatus: %v", v)
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
	PublishStatus *CreateAPIProductVersionDTOPublishStatus `json:"publish_status,omitempty"`
	// Indicates if the version of the API product is deprecated. Applies deprecation or removes deprecation from all related portal product versions. This field is deprecated: Use [PortalProductVersion.deprecated](https://docs.konghq.com/konnect/api/portal-management/latest/#/Portal%20Product%20Versions/create-portal-product-version) instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Deprecated *bool `json:"deprecated,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels         map[string]string      `json:"labels,omitempty"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
}

func (o *CreateAPIProductVersionDTO) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateAPIProductVersionDTO) GetPublishStatus() *CreateAPIProductVersionDTOPublishStatus {
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

func (o *CreateAPIProductVersionDTO) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *CreateAPIProductVersionDTO) GetGatewayService() *GatewayServicePayload {
	if o == nil {
		return nil
	}
	return o.GatewayService
}
