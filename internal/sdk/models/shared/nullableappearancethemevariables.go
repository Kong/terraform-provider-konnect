// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// NullableAppearanceThemeVariables - Groups of variables for configuring visual details of the portal user interface. Set theme_name to 'custom' to use custom values for theme variables.
type NullableAppearanceThemeVariables struct {
	Colors AppearanceThemeColorVariables `json:"colors"`
}

func (o *NullableAppearanceThemeVariables) GetColors() AppearanceThemeColorVariables {
	if o == nil {
		return AppearanceThemeColorVariables{}
	}
	return o.Colors
}
