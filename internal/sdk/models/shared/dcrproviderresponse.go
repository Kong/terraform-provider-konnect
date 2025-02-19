// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/internal/utils"
)

type DcrProviderResponseType string

const (
	DcrProviderResponseTypeAuth0   DcrProviderResponseType = "auth0"
	DcrProviderResponseTypeAzureAd DcrProviderResponseType = "azureAd"
	DcrProviderResponseTypeCurity  DcrProviderResponseType = "curity"
	DcrProviderResponseTypeOkta    DcrProviderResponseType = "okta"
	DcrProviderResponseTypeHTTP    DcrProviderResponseType = "http"
)

// DcrProviderResponse - A response containing a single DCR provider object. Sensitive fields will be removed from the response.
type DcrProviderResponse struct {
	DcrProviderAuth0   *DcrProviderAuth0   `queryParam:"inline"`
	DcrProviderAzureAd *DcrProviderAzureAd `queryParam:"inline"`
	DcrProviderCurity  *DcrProviderCurity  `queryParam:"inline"`
	DcrProviderOkta    *DcrProviderOkta    `queryParam:"inline"`
	DcrProviderHTTP    *DcrProviderHTTP    `queryParam:"inline"`

	Type DcrProviderResponseType
}

func CreateDcrProviderResponseAuth0(auth0 DcrProviderAuth0) DcrProviderResponse {
	typ := DcrProviderResponseTypeAuth0

	typStr := ProviderType(typ)
	auth0.ProviderType = typStr

	return DcrProviderResponse{
		DcrProviderAuth0: &auth0,
		Type:             typ,
	}
}

func CreateDcrProviderResponseAzureAd(azureAd DcrProviderAzureAd) DcrProviderResponse {
	typ := DcrProviderResponseTypeAzureAd

	typStr := DcrProviderAzureAdProviderType(typ)
	azureAd.ProviderType = typStr

	return DcrProviderResponse{
		DcrProviderAzureAd: &azureAd,
		Type:               typ,
	}
}

func CreateDcrProviderResponseCurity(curity DcrProviderCurity) DcrProviderResponse {
	typ := DcrProviderResponseTypeCurity

	typStr := DcrProviderCurityProviderType(typ)
	curity.ProviderType = typStr

	return DcrProviderResponse{
		DcrProviderCurity: &curity,
		Type:              typ,
	}
}

func CreateDcrProviderResponseOkta(okta DcrProviderOkta) DcrProviderResponse {
	typ := DcrProviderResponseTypeOkta

	typStr := DcrProviderOktaProviderType(typ)
	okta.ProviderType = typStr

	return DcrProviderResponse{
		DcrProviderOkta: &okta,
		Type:            typ,
	}
}

func CreateDcrProviderResponseHTTP(http DcrProviderHTTP) DcrProviderResponse {
	typ := DcrProviderResponseTypeHTTP

	typStr := DcrProviderHTTPProviderType(typ)
	http.ProviderType = typStr

	return DcrProviderResponse{
		DcrProviderHTTP: &http,
		Type:            typ,
	}
}

func (u *DcrProviderResponse) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		ProviderType string `json:"provider_type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.ProviderType {
	case "auth0":
		dcrProviderAuth0 := new(DcrProviderAuth0)
		if err := utils.UnmarshalJSON(data, &dcrProviderAuth0, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (ProviderType == auth0) type DcrProviderAuth0 within DcrProviderResponse: %w", string(data), err)
		}

		u.DcrProviderAuth0 = dcrProviderAuth0
		u.Type = DcrProviderResponseTypeAuth0
		return nil
	case "azureAd":
		dcrProviderAzureAd := new(DcrProviderAzureAd)
		if err := utils.UnmarshalJSON(data, &dcrProviderAzureAd, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (ProviderType == azureAd) type DcrProviderAzureAd within DcrProviderResponse: %w", string(data), err)
		}

		u.DcrProviderAzureAd = dcrProviderAzureAd
		u.Type = DcrProviderResponseTypeAzureAd
		return nil
	case "curity":
		dcrProviderCurity := new(DcrProviderCurity)
		if err := utils.UnmarshalJSON(data, &dcrProviderCurity, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (ProviderType == curity) type DcrProviderCurity within DcrProviderResponse: %w", string(data), err)
		}

		u.DcrProviderCurity = dcrProviderCurity
		u.Type = DcrProviderResponseTypeCurity
		return nil
	case "okta":
		dcrProviderOkta := new(DcrProviderOkta)
		if err := utils.UnmarshalJSON(data, &dcrProviderOkta, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (ProviderType == okta) type DcrProviderOkta within DcrProviderResponse: %w", string(data), err)
		}

		u.DcrProviderOkta = dcrProviderOkta
		u.Type = DcrProviderResponseTypeOkta
		return nil
	case "http":
		dcrProviderHTTP := new(DcrProviderHTTP)
		if err := utils.UnmarshalJSON(data, &dcrProviderHTTP, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (ProviderType == http) type DcrProviderHTTP within DcrProviderResponse: %w", string(data), err)
		}

		u.DcrProviderHTTP = dcrProviderHTTP
		u.Type = DcrProviderResponseTypeHTTP
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DcrProviderResponse", string(data))
}

func (u DcrProviderResponse) MarshalJSON() ([]byte, error) {
	if u.DcrProviderAuth0 != nil {
		return utils.MarshalJSON(u.DcrProviderAuth0, "", true)
	}

	if u.DcrProviderAzureAd != nil {
		return utils.MarshalJSON(u.DcrProviderAzureAd, "", true)
	}

	if u.DcrProviderCurity != nil {
		return utils.MarshalJSON(u.DcrProviderCurity, "", true)
	}

	if u.DcrProviderOkta != nil {
		return utils.MarshalJSON(u.DcrProviderOkta, "", true)
	}

	if u.DcrProviderHTTP != nil {
		return utils.MarshalJSON(u.DcrProviderHTTP, "", true)
	}

	return nil, errors.New("could not marshal union type DcrProviderResponse: all fields are null")
}
