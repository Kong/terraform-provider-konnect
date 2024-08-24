// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ControlPlaneGeo - Set of control-plane geos supported for deploying cloud-gateways configurations.
type ControlPlaneGeo string

const (
	ControlPlaneGeoUs ControlPlaneGeo = "us"
	ControlPlaneGeoEu ControlPlaneGeo = "eu"
	ControlPlaneGeoAu ControlPlaneGeo = "au"
)

func (e ControlPlaneGeo) ToPointer() *ControlPlaneGeo {
	return &e
}
func (e *ControlPlaneGeo) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us":
		fallthrough
	case "eu":
		fallthrough
	case "au":
		*e = ControlPlaneGeo(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ControlPlaneGeo: %v", v)
	}
}
