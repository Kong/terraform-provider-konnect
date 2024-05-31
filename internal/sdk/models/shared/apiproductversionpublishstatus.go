// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type APIProductVersionPublishStatus string

const (
	APIProductVersionPublishStatusPublished   APIProductVersionPublishStatus = "published"
	APIProductVersionPublishStatusUnpublished APIProductVersionPublishStatus = "unpublished"
)

func (e APIProductVersionPublishStatus) ToPointer() *APIProductVersionPublishStatus {
	return &e
}
func (e *APIProductVersionPublishStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "published":
		fallthrough
	case "unpublished":
		*e = APIProductVersionPublishStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APIProductVersionPublishStatus: %v", v)
	}
}
