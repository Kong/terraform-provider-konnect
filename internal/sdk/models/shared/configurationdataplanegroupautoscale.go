// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type ConfigurationDataPlaneGroupAutoscaleType string

const (
	ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleStatic    ConfigurationDataPlaneGroupAutoscaleType = "ConfigurationDataPlaneGroupAutoscaleStatic"
	ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleAutopilot ConfigurationDataPlaneGroupAutoscaleType = "ConfigurationDataPlaneGroupAutoscaleAutopilot"
)

type ConfigurationDataPlaneGroupAutoscale struct {
	ConfigurationDataPlaneGroupAutoscaleStatic    *ConfigurationDataPlaneGroupAutoscaleStatic
	ConfigurationDataPlaneGroupAutoscaleAutopilot *ConfigurationDataPlaneGroupAutoscaleAutopilot

	Type ConfigurationDataPlaneGroupAutoscaleType
}

func CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleStatic(configurationDataPlaneGroupAutoscaleStatic ConfigurationDataPlaneGroupAutoscaleStatic) ConfigurationDataPlaneGroupAutoscale {
	typ := ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleStatic

	return ConfigurationDataPlaneGroupAutoscale{
		ConfigurationDataPlaneGroupAutoscaleStatic: &configurationDataPlaneGroupAutoscaleStatic,
		Type: typ,
	}
}

func CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(configurationDataPlaneGroupAutoscaleAutopilot ConfigurationDataPlaneGroupAutoscaleAutopilot) ConfigurationDataPlaneGroupAutoscale {
	typ := ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleAutopilot

	return ConfigurationDataPlaneGroupAutoscale{
		ConfigurationDataPlaneGroupAutoscaleAutopilot: &configurationDataPlaneGroupAutoscaleAutopilot,
		Type: typ,
	}
}

func (u *ConfigurationDataPlaneGroupAutoscale) UnmarshalJSON(data []byte) error {

	var configurationDataPlaneGroupAutoscaleStatic ConfigurationDataPlaneGroupAutoscaleStatic = ConfigurationDataPlaneGroupAutoscaleStatic{}
	if err := utils.UnmarshalJSON(data, &configurationDataPlaneGroupAutoscaleStatic, "", true, true); err == nil {
		u.ConfigurationDataPlaneGroupAutoscaleStatic = &configurationDataPlaneGroupAutoscaleStatic
		u.Type = ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleStatic
		return nil
	}

	var configurationDataPlaneGroupAutoscaleAutopilot ConfigurationDataPlaneGroupAutoscaleAutopilot = ConfigurationDataPlaneGroupAutoscaleAutopilot{}
	if err := utils.UnmarshalJSON(data, &configurationDataPlaneGroupAutoscaleAutopilot, "", true, true); err == nil {
		u.ConfigurationDataPlaneGroupAutoscaleAutopilot = &configurationDataPlaneGroupAutoscaleAutopilot
		u.Type = ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleAutopilot
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ConfigurationDataPlaneGroupAutoscale", string(data))
}

func (u ConfigurationDataPlaneGroupAutoscale) MarshalJSON() ([]byte, error) {
	if u.ConfigurationDataPlaneGroupAutoscaleStatic != nil {
		return utils.MarshalJSON(u.ConfigurationDataPlaneGroupAutoscaleStatic, "", true)
	}

	if u.ConfigurationDataPlaneGroupAutoscaleAutopilot != nil {
		return utils.MarshalJSON(u.ConfigurationDataPlaneGroupAutoscaleAutopilot, "", true)
	}

	return nil, errors.New("could not marshal union type ConfigurationDataPlaneGroupAutoscale: all fields are null")
}
