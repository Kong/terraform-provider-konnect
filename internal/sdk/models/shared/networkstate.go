// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// NetworkState - State of the network.
type NetworkState string

const (
	NetworkStateCreated      NetworkState = "created"
	NetworkStateInitializing NetworkState = "initializing"
	NetworkStateOffline      NetworkState = "offline"
	NetworkStateReady        NetworkState = "ready"
	NetworkStateTerminating  NetworkState = "terminating"
	NetworkStateTerminated   NetworkState = "terminated"
)

func (e NetworkState) ToPointer() *NetworkState {
	return &e
}
func (e *NetworkState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created":
		fallthrough
	case "initializing":
		fallthrough
	case "offline":
		fallthrough
	case "ready":
		fallthrough
	case "terminating":
		fallthrough
	case "terminated":
		*e = NetworkState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NetworkState: %v", v)
	}
}
