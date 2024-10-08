// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AppearanceFontName - The name of the font to render in the browser.
type AppearanceFontName string

const (
	AppearanceFontNameRoboto        AppearanceFontName = "Roboto"
	AppearanceFontNameInter         AppearanceFontName = "Inter"
	AppearanceFontNameOpenSans      AppearanceFontName = "Open Sans"
	AppearanceFontNameLato          AppearanceFontName = "Lato"
	AppearanceFontNameSlabo27px     AppearanceFontName = "Slabo 27px"
	AppearanceFontNameSlabo13px     AppearanceFontName = "Slabo 13px"
	AppearanceFontNameOswald        AppearanceFontName = "Oswald"
	AppearanceFontNameSourceSansPro AppearanceFontName = "Source Sans Pro"
	AppearanceFontNameMontserrat    AppearanceFontName = "Montserrat"
	AppearanceFontNameRaleway       AppearanceFontName = "Raleway"
	AppearanceFontNamePtSans        AppearanceFontName = "PT Sans"
	AppearanceFontNameLora          AppearanceFontName = "Lora"
	AppearanceFontNameRobotoMono    AppearanceFontName = "Roboto Mono"
	AppearanceFontNameInconsolata   AppearanceFontName = "Inconsolata"
	AppearanceFontNameSourceCodePro AppearanceFontName = "Source Code Pro"
	AppearanceFontNamePtMono        AppearanceFontName = "PT Mono"
	AppearanceFontNameUbuntuMono    AppearanceFontName = "Ubuntu Mono"
	AppearanceFontNameIbmPlexMono   AppearanceFontName = "IBM Plex Mono"
)

func (e AppearanceFontName) ToPointer() *AppearanceFontName {
	return &e
}
func (e *AppearanceFontName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Roboto":
		fallthrough
	case "Inter":
		fallthrough
	case "Open Sans":
		fallthrough
	case "Lato":
		fallthrough
	case "Slabo 27px":
		fallthrough
	case "Slabo 13px":
		fallthrough
	case "Oswald":
		fallthrough
	case "Source Sans Pro":
		fallthrough
	case "Montserrat":
		fallthrough
	case "Raleway":
		fallthrough
	case "PT Sans":
		fallthrough
	case "Lora":
		fallthrough
	case "Roboto Mono":
		fallthrough
	case "Inconsolata":
		fallthrough
	case "Source Code Pro":
		fallthrough
	case "PT Mono":
		fallthrough
	case "Ubuntu Mono":
		fallthrough
	case "IBM Plex Mono":
		*e = AppearanceFontName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppearanceFontName: %v", v)
	}
}
