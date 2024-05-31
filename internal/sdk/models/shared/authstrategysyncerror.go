// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type Value string

const (
	ValuePluginSyncErrorComm               Value = "plugin_sync_error_comm"
	ValuePluginSyncErrorUnknown            Value = "plugin_sync_error_unknown"
	ValuePluginSyncErrorFatal              Value = "plugin_sync_error_fatal"
	ValuePluginSyncErrorUpdatingPluginRefs Value = "plugin_sync_error_updating_plugin_refs"
)

func (e Value) ToPointer() *Value {
	return &e
}
func (e *Value) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plugin_sync_error_comm":
		fallthrough
	case "plugin_sync_error_unknown":
		fallthrough
	case "plugin_sync_error_fatal":
		fallthrough
	case "plugin_sync_error_updating_plugin_refs":
		*e = Value(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Value: %v", v)
	}
}

type Details struct {
	Type                 *string  `json:"type,omitempty"`
	Message              []string `json:"message,omitempty"`
	AdditionalProperties any      `additionalProperties:"true" json:"-"`
}

func (d Details) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Details) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Details) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Details) GetMessage() []string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *Details) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

type Info struct {
	Details              []Details `json:"details,omitempty"`
	AdditionalProperties any       `additionalProperties:"true" json:"-"`
}

func (i Info) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Info) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Info) GetDetails() []Details {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *Info) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

type AuthStrategySyncError struct {
	PluginName *string `json:"plugin_name,omitempty"`
	Value      *Value  `json:"value,omitempty"`
	Message    string  `json:"message"`
	Info       *Info   `json:"info,omitempty"`
}

func (o *AuthStrategySyncError) GetPluginName() *string {
	if o == nil {
		return nil
	}
	return o.PluginName
}

func (o *AuthStrategySyncError) GetValue() *Value {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *AuthStrategySyncError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *AuthStrategySyncError) GetInfo() *Info {
	if o == nil {
		return nil
	}
	return o.Info
}
