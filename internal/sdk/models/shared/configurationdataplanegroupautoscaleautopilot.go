// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConfigurationDataPlaneGroupAutoscaleAutopilotKind string

const (
	ConfigurationDataPlaneGroupAutoscaleAutopilotKindAutopilot ConfigurationDataPlaneGroupAutoscaleAutopilotKind = "autopilot"
)

func (e ConfigurationDataPlaneGroupAutoscaleAutopilotKind) ToPointer() *ConfigurationDataPlaneGroupAutoscaleAutopilotKind {
	return &e
}
func (e *ConfigurationDataPlaneGroupAutoscaleAutopilotKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "autopilot":
		*e = ConfigurationDataPlaneGroupAutoscaleAutopilotKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConfigurationDataPlaneGroupAutoscaleAutopilotKind: %v", v)
	}
}

// ConfigurationDataPlaneGroupAutoscaleAutopilot - Object that describes the autopilot autoscaling strategy.
type ConfigurationDataPlaneGroupAutoscaleAutopilot struct {
	// Base number of requests per second that the deployment target should support.
	BaseRps int64                                             `json:"base_rps"`
	Kind    ConfigurationDataPlaneGroupAutoscaleAutopilotKind `json:"kind"`
	// Max number of requests per second that the deployment target should support. If not set, this defaults to 10x base_rps.
	MaxRps *int64 `json:"max_rps,omitempty"`
}

func (o *ConfigurationDataPlaneGroupAutoscaleAutopilot) GetBaseRps() int64 {
	if o == nil {
		return 0
	}
	return o.BaseRps
}

func (o *ConfigurationDataPlaneGroupAutoscaleAutopilot) GetKind() ConfigurationDataPlaneGroupAutoscaleAutopilotKind {
	if o == nil {
		return ConfigurationDataPlaneGroupAutoscaleAutopilotKind("")
	}
	return o.Kind
}

func (o *ConfigurationDataPlaneGroupAutoscaleAutopilot) GetMaxRps() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRps
}
