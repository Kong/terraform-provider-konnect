// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type GetPortalAppearanceResponse struct {
	// Select a pre-existing default theme or specify 'custom' to use custom_theme variables.
	ThemeName PortalTheme `json:"theme_name"`
	// Groups of variables for configuring visual details of the portal user interface. Set theme_name to 'custom' to use custom values for theme variables.
	CustomTheme *NullableAppearanceThemeVariables `json:"custom_theme"`
	// Font selections to render text in the portal user interface. Must set use_custom_fonts to true to enable using custom font values.
	CustomFonts *NullableAppearanceFonts `json:"custom_fonts"`
	// If true, fonts in custom_fonts will be used over the theme's default fonts
	UseCustomFonts bool `json:"use_custom_fonts"`
	// Values to display for customizable text in the portal user interface
	Text *NullableAppearanceTextVariables `json:"text"`
	// A collection of binary image data to customize images in the portal
	Images *AppearanceImages `json:"images"`
}

func (o *GetPortalAppearanceResponse) GetThemeName() PortalTheme {
	if o == nil {
		return PortalTheme("")
	}
	return o.ThemeName
}

func (o *GetPortalAppearanceResponse) GetCustomTheme() *NullableAppearanceThemeVariables {
	if o == nil {
		return nil
	}
	return o.CustomTheme
}

func (o *GetPortalAppearanceResponse) GetCustomFonts() *NullableAppearanceFonts {
	if o == nil {
		return nil
	}
	return o.CustomFonts
}

func (o *GetPortalAppearanceResponse) GetUseCustomFonts() bool {
	if o == nil {
		return false
	}
	return o.UseCustomFonts
}

func (o *GetPortalAppearanceResponse) GetText() *NullableAppearanceTextVariables {
	if o == nil {
		return nil
	}
	return o.Text
}

func (o *GetPortalAppearanceResponse) GetImages() *AppearanceImages {
	if o == nil {
		return nil
	}
	return o.Images
}
