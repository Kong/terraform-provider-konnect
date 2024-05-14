// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"time"
)

type AppAuthStrategyOpenIDConnectResponseStrategyType string

const (
	AppAuthStrategyOpenIDConnectResponseStrategyTypeOpenidConnect AppAuthStrategyOpenIDConnectResponseStrategyType = "openid_connect"
)

func (e AppAuthStrategyOpenIDConnectResponseStrategyType) ToPointer() *AppAuthStrategyOpenIDConnectResponseStrategyType {
	return &e
}

func (e *AppAuthStrategyOpenIDConnectResponseStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openid_connect":
		*e = AppAuthStrategyOpenIDConnectResponseStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppAuthStrategyOpenIDConnectResponseStrategyType: %v", v)
	}
}

// AppAuthStrategyOpenIDConnectResponseConfigs - JSON-B object containing the configuration for the OIDC strategy
type AppAuthStrategyOpenIDConnectResponseConfigs struct {
	// A more advanced mode to configure an API Product Version’s Application Auth Strategy.
	// Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests.
	// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy.
	// An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.
	//
	OpenidConnect AppAuthStrategyConfigOpenIDConnect `json:"openid-connect"`
}

func (o *AppAuthStrategyOpenIDConnectResponseConfigs) GetOpenidConnect() AppAuthStrategyConfigOpenIDConnect {
	if o == nil {
		return AppAuthStrategyConfigOpenIDConnect{}
	}
	return o.OpenidConnect
}

type AppAuthStrategyOpenIDConnectResponseDcrProvider struct {
	// Contains a unique identifier used by the API for this resource.
	ID   string `json:"id"`
	Name string `json:"name"`
	// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
	//
	DisplayName *string `json:"display_name,omitempty"`
	// The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http
	ProviderType DcrProviderType `json:"provider_type"`
}

func (o *AppAuthStrategyOpenIDConnectResponseDcrProvider) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppAuthStrategyOpenIDConnectResponseDcrProvider) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppAuthStrategyOpenIDConnectResponseDcrProvider) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppAuthStrategyOpenIDConnectResponseDcrProvider) GetProviderType() DcrProviderType {
	if o == nil {
		return DcrProviderType("")
	}
	return o.ProviderType
}

// AppAuthStrategyOpenIDConnectResponse - Response payload from creating an OIDC Application Auth Strategy
type AppAuthStrategyOpenIDConnectResponse struct {
	// Contains a unique identifier used by the API for this resource.
	ID string `json:"id"`
	// The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
	//
	Name string `json:"name"`
	// The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
	//
	DisplayName  string                                           `json:"display_name"`
	StrategyType AppAuthStrategyOpenIDConnectResponseStrategyType `json:"strategy_type"`
	// JSON-B object containing the configuration for the OIDC strategy
	Configs AppAuthStrategyOpenIDConnectResponseConfigs `json:"configs"`
	// At least one published product version is using this auth strategy.
	Active      bool                                             `json:"active"`
	DcrProvider *AppAuthStrategyOpenIDConnectResponseDcrProvider `json:"dcr_provider"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a AppAuthStrategyOpenIDConnectResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppAuthStrategyOpenIDConnectResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetStrategyType() AppAuthStrategyOpenIDConnectResponseStrategyType {
	if o == nil {
		return AppAuthStrategyOpenIDConnectResponseStrategyType("")
	}
	return o.StrategyType
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetConfigs() AppAuthStrategyOpenIDConnectResponseConfigs {
	if o == nil {
		return AppAuthStrategyOpenIDConnectResponseConfigs{}
	}
	return o.Configs
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetDcrProvider() *AppAuthStrategyOpenIDConnectResponseDcrProvider {
	if o == nil {
		return nil
	}
	return o.DcrProvider
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AppAuthStrategyOpenIDConnectResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type AppAuthStrategyKeyAuthResponseStrategyType string

const (
	AppAuthStrategyKeyAuthResponseStrategyTypeKeyAuth AppAuthStrategyKeyAuthResponseStrategyType = "key_auth"
)

func (e AppAuthStrategyKeyAuthResponseStrategyType) ToPointer() *AppAuthStrategyKeyAuthResponseStrategyType {
	return &e
}

func (e *AppAuthStrategyKeyAuthResponseStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "key_auth":
		*e = AppAuthStrategyKeyAuthResponseStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppAuthStrategyKeyAuthResponseStrategyType: %v", v)
	}
}

// AppAuthStrategyKeyAuthResponseConfigs - JSON-B object containing the configuration for the Key Auth strategy
type AppAuthStrategyKeyAuthResponseConfigs struct {
	// The most basic mode to configure an Application Auth Strategy for an API Product Version.
	// Using this mode will allow developers to generate API keys that will authenticate their application requests.
	// Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for Key Auth.
	//
	KeyAuth AppAuthStrategyConfigKeyAuth `json:"key-auth"`
}

func (o *AppAuthStrategyKeyAuthResponseConfigs) GetKeyAuth() AppAuthStrategyConfigKeyAuth {
	if o == nil {
		return AppAuthStrategyConfigKeyAuth{}
	}
	return o.KeyAuth
}

type DcrProvider struct {
	// Contains a unique identifier used by the API for this resource.
	ID   string `json:"id"`
	Name string `json:"name"`
	// The display name of the DCR provider. This is used to identify the DCR provider in the Portal UI.
	//
	DisplayName *string `json:"display_name,omitempty"`
	// The type of DCR provider. Can be one of the following - auth0, azureAd, curity, okta, http
	ProviderType DcrProviderType `json:"provider_type"`
}

func (o *DcrProvider) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DcrProvider) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DcrProvider) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *DcrProvider) GetProviderType() DcrProviderType {
	if o == nil {
		return DcrProviderType("")
	}
	return o.ProviderType
}

// AppAuthStrategyKeyAuthResponse - Response payload from creating or updating a Key Auth Application Auth Strategy
type AppAuthStrategyKeyAuthResponse struct {
	// Contains a unique identifier used by the API for this resource.
	ID string `json:"id"`
	// The name of the auth strategy. This is used to identify the auth strategy in the Konnect UI.
	//
	Name string `json:"name"`
	// The display name of the Auth strategy. This is used to identify the Auth strategy in the Portal UI.
	//
	DisplayName  string                                     `json:"display_name"`
	StrategyType AppAuthStrategyKeyAuthResponseStrategyType `json:"strategy_type"`
	// JSON-B object containing the configuration for the Key Auth strategy
	Configs AppAuthStrategyKeyAuthResponseConfigs `json:"configs"`
	// At least one published product version is using this auth strategy.
	Active      bool         `json:"active"`
	DcrProvider *DcrProvider `json:"dcr_provider"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt time.Time `json:"created_at"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt time.Time `json:"updated_at"`
}

func (a AppAuthStrategyKeyAuthResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppAuthStrategyKeyAuthResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppAuthStrategyKeyAuthResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppAuthStrategyKeyAuthResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppAuthStrategyKeyAuthResponse) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *AppAuthStrategyKeyAuthResponse) GetStrategyType() AppAuthStrategyKeyAuthResponseStrategyType {
	if o == nil {
		return AppAuthStrategyKeyAuthResponseStrategyType("")
	}
	return o.StrategyType
}

func (o *AppAuthStrategyKeyAuthResponse) GetConfigs() AppAuthStrategyKeyAuthResponseConfigs {
	if o == nil {
		return AppAuthStrategyKeyAuthResponseConfigs{}
	}
	return o.Configs
}

func (o *AppAuthStrategyKeyAuthResponse) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *AppAuthStrategyKeyAuthResponse) GetDcrProvider() *DcrProvider {
	if o == nil {
		return nil
	}
	return o.DcrProvider
}

func (o *AppAuthStrategyKeyAuthResponse) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AppAuthStrategyKeyAuthResponse) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

type CreateAppAuthStrategyResponseType string

const (
	CreateAppAuthStrategyResponseTypeKeyAuth       CreateAppAuthStrategyResponseType = "key_auth"
	CreateAppAuthStrategyResponseTypeOpenidConnect CreateAppAuthStrategyResponseType = "openid_connect"
)

// CreateAppAuthStrategyResponse - A set of plugin configurations that represent how the gateway will perform authentication and authorization for a Product Version. Called “Auth Strategy” for short in the context of portals/applications. The plugins are synced to any Gateway Service that is currently linked or becomes linked to the Product Version.
type CreateAppAuthStrategyResponse struct {
	AppAuthStrategyKeyAuthResponse       *AppAuthStrategyKeyAuthResponse
	AppAuthStrategyOpenIDConnectResponse *AppAuthStrategyOpenIDConnectResponse

	Type CreateAppAuthStrategyResponseType
}

func CreateCreateAppAuthStrategyResponseKeyAuth(keyAuth AppAuthStrategyKeyAuthResponse) CreateAppAuthStrategyResponse {
	typ := CreateAppAuthStrategyResponseTypeKeyAuth

	typStr := AppAuthStrategyKeyAuthResponseStrategyType(typ)
	keyAuth.StrategyType = typStr

	return CreateAppAuthStrategyResponse{
		AppAuthStrategyKeyAuthResponse: &keyAuth,
		Type:                           typ,
	}
}

func CreateCreateAppAuthStrategyResponseOpenidConnect(openidConnect AppAuthStrategyOpenIDConnectResponse) CreateAppAuthStrategyResponse {
	typ := CreateAppAuthStrategyResponseTypeOpenidConnect

	typStr := AppAuthStrategyOpenIDConnectResponseStrategyType(typ)
	openidConnect.StrategyType = typStr

	return CreateAppAuthStrategyResponse{
		AppAuthStrategyOpenIDConnectResponse: &openidConnect,
		Type:                                 typ,
	}
}

func (u *CreateAppAuthStrategyResponse) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		StrategyType string `json:"strategy_type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.StrategyType {
	case "key_auth":
		appAuthStrategyKeyAuthResponse := new(AppAuthStrategyKeyAuthResponse)
		if err := utils.UnmarshalJSON(data, &appAuthStrategyKeyAuthResponse, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AppAuthStrategyKeyAuthResponse = appAuthStrategyKeyAuthResponse
		u.Type = CreateAppAuthStrategyResponseTypeKeyAuth
		return nil
	case "openid_connect":
		appAuthStrategyOpenIDConnectResponse := new(AppAuthStrategyOpenIDConnectResponse)
		if err := utils.UnmarshalJSON(data, &appAuthStrategyOpenIDConnectResponse, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AppAuthStrategyOpenIDConnectResponse = appAuthStrategyOpenIDConnectResponse
		u.Type = CreateAppAuthStrategyResponseTypeOpenidConnect
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateAppAuthStrategyResponse) MarshalJSON() ([]byte, error) {
	if u.AppAuthStrategyKeyAuthResponse != nil {
		return utils.MarshalJSON(u.AppAuthStrategyKeyAuthResponse, "", true)
	}

	if u.AppAuthStrategyOpenIDConnectResponse != nil {
		return utils.MarshalJSON(u.AppAuthStrategyOpenIDConnectResponse, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
