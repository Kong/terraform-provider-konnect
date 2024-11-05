// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Kind string

const (
	KindStatic Kind = "static"
)

func (e Kind) ToPointer() *Kind {
	return &e
}
func (e *Kind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		*e = Kind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Kind: %v", v)
	}
}

// ConfigurationDataPlaneGroupAutoscaleStatic - Object that describes the static autoscaling strategy.
type ConfigurationDataPlaneGroupAutoscaleStatic struct {
	Kind Kind `json:"kind"`
	// Instance type name to indicate capacity.
	InstanceType InstanceTypeName `json:"instance_type"`
	// Number of data-planes the deployment target will contain.
	RequestedInstances int64 `json:"requested_instances"`
}

func (o *ConfigurationDataPlaneGroupAutoscaleStatic) GetKind() Kind {
	if o == nil {
		return Kind("")
	}
	return o.Kind
}

func (o *ConfigurationDataPlaneGroupAutoscaleStatic) GetInstanceType() InstanceTypeName {
	if o == nil {
		return InstanceTypeName("")
	}
	return o.InstanceType
}

func (o *ConfigurationDataPlaneGroupAutoscaleStatic) GetRequestedInstances() int64 {
	if o == nil {
		return 0
	}
	return o.RequestedInstances
}