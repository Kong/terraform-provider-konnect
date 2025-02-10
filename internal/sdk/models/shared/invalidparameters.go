// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type InvalidParametersType string

const (
	InvalidParametersTypeInvalidParameterStandard      InvalidParametersType = "InvalidParameterStandard"
	InvalidParametersTypeInvalidParameterMinimumLength InvalidParametersType = "InvalidParameterMinimumLength"
	InvalidParametersTypeInvalidParameterMaximumLength InvalidParametersType = "InvalidParameterMaximumLength"
	InvalidParametersTypeInvalidParameterChoiceItem    InvalidParametersType = "InvalidParameterChoiceItem"
	InvalidParametersTypeInvalidParameterDependentItem InvalidParametersType = "InvalidParameterDependentItem"
)

type InvalidParameters struct {
	InvalidParameterStandard      *InvalidParameterStandard      `queryParam:"inline"`
	InvalidParameterMinimumLength *InvalidParameterMinimumLength `queryParam:"inline"`
	InvalidParameterMaximumLength *InvalidParameterMaximumLength `queryParam:"inline"`
	InvalidParameterChoiceItem    *InvalidParameterChoiceItem    `queryParam:"inline"`
	InvalidParameterDependentItem *InvalidParameterDependentItem `queryParam:"inline"`

	Type InvalidParametersType
}

func CreateInvalidParametersInvalidParameterStandard(invalidParameterStandard InvalidParameterStandard) InvalidParameters {
	typ := InvalidParametersTypeInvalidParameterStandard

	return InvalidParameters{
		InvalidParameterStandard: &invalidParameterStandard,
		Type:                     typ,
	}
}

func CreateInvalidParametersInvalidParameterMinimumLength(invalidParameterMinimumLength InvalidParameterMinimumLength) InvalidParameters {
	typ := InvalidParametersTypeInvalidParameterMinimumLength

	return InvalidParameters{
		InvalidParameterMinimumLength: &invalidParameterMinimumLength,
		Type:                          typ,
	}
}

func CreateInvalidParametersInvalidParameterMaximumLength(invalidParameterMaximumLength InvalidParameterMaximumLength) InvalidParameters {
	typ := InvalidParametersTypeInvalidParameterMaximumLength

	return InvalidParameters{
		InvalidParameterMaximumLength: &invalidParameterMaximumLength,
		Type:                          typ,
	}
}

func CreateInvalidParametersInvalidParameterChoiceItem(invalidParameterChoiceItem InvalidParameterChoiceItem) InvalidParameters {
	typ := InvalidParametersTypeInvalidParameterChoiceItem

	return InvalidParameters{
		InvalidParameterChoiceItem: &invalidParameterChoiceItem,
		Type:                       typ,
	}
}

func CreateInvalidParametersInvalidParameterDependentItem(invalidParameterDependentItem InvalidParameterDependentItem) InvalidParameters {
	typ := InvalidParametersTypeInvalidParameterDependentItem

	return InvalidParameters{
		InvalidParameterDependentItem: &invalidParameterDependentItem,
		Type:                          typ,
	}
}

func (u *InvalidParameters) UnmarshalJSON(data []byte) error {

	var invalidParameterStandard InvalidParameterStandard = InvalidParameterStandard{}
	if err := utils.UnmarshalJSON(data, &invalidParameterStandard, "", true, true); err == nil {
		u.InvalidParameterStandard = &invalidParameterStandard
		u.Type = InvalidParametersTypeInvalidParameterStandard
		return nil
	}

	var invalidParameterMinimumLength InvalidParameterMinimumLength = InvalidParameterMinimumLength{}
	if err := utils.UnmarshalJSON(data, &invalidParameterMinimumLength, "", true, true); err == nil {
		u.InvalidParameterMinimumLength = &invalidParameterMinimumLength
		u.Type = InvalidParametersTypeInvalidParameterMinimumLength
		return nil
	}

	var invalidParameterMaximumLength InvalidParameterMaximumLength = InvalidParameterMaximumLength{}
	if err := utils.UnmarshalJSON(data, &invalidParameterMaximumLength, "", true, true); err == nil {
		u.InvalidParameterMaximumLength = &invalidParameterMaximumLength
		u.Type = InvalidParametersTypeInvalidParameterMaximumLength
		return nil
	}

	var invalidParameterChoiceItem InvalidParameterChoiceItem = InvalidParameterChoiceItem{}
	if err := utils.UnmarshalJSON(data, &invalidParameterChoiceItem, "", true, true); err == nil {
		u.InvalidParameterChoiceItem = &invalidParameterChoiceItem
		u.Type = InvalidParametersTypeInvalidParameterChoiceItem
		return nil
	}

	var invalidParameterDependentItem InvalidParameterDependentItem = InvalidParameterDependentItem{}
	if err := utils.UnmarshalJSON(data, &invalidParameterDependentItem, "", true, true); err == nil {
		u.InvalidParameterDependentItem = &invalidParameterDependentItem
		u.Type = InvalidParametersTypeInvalidParameterDependentItem
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for InvalidParameters", string(data))
}

func (u InvalidParameters) MarshalJSON() ([]byte, error) {
	if u.InvalidParameterStandard != nil {
		return utils.MarshalJSON(u.InvalidParameterStandard, "", true)
	}

	if u.InvalidParameterMinimumLength != nil {
		return utils.MarshalJSON(u.InvalidParameterMinimumLength, "", true)
	}

	if u.InvalidParameterMaximumLength != nil {
		return utils.MarshalJSON(u.InvalidParameterMaximumLength, "", true)
	}

	if u.InvalidParameterChoiceItem != nil {
		return utils.MarshalJSON(u.InvalidParameterChoiceItem, "", true)
	}

	if u.InvalidParameterDependentItem != nil {
		return utils.MarshalJSON(u.InvalidParameterDependentItem, "", true)
	}

	return nil, errors.New("could not marshal union type InvalidParameters: all fields are null")
}
