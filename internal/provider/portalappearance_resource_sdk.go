// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *PortalAppearanceResourceModel) ToSharedUpdatePortalAppearanceRequest() *shared.UpdatePortalAppearanceRequest {
	themeName := new(shared.PortalTheme)
	if !r.ThemeName.IsUnknown() && !r.ThemeName.IsNull() {
		*themeName = shared.PortalTheme(r.ThemeName.ValueString())
	} else {
		themeName = nil
	}
	var customTheme *shared.NullableAppearanceThemeVariables
	if r.CustomTheme != nil {
		var value string
		value = r.CustomTheme.Colors.Section.Header.Value.ValueString()

		description := new(string)
		if !r.CustomTheme.Colors.Section.Header.Description.IsUnknown() && !r.CustomTheme.Colors.Section.Header.Description.IsNull() {
			*description = r.CustomTheme.Colors.Section.Header.Description.ValueString()
		} else {
			description = nil
		}
		header := shared.AppearanceColorVariable{
			Value:       value,
			Description: description,
		}
		var value1 string
		value1 = r.CustomTheme.Colors.Section.Body.Value.ValueString()

		description1 := new(string)
		if !r.CustomTheme.Colors.Section.Body.Description.IsUnknown() && !r.CustomTheme.Colors.Section.Body.Description.IsNull() {
			*description1 = r.CustomTheme.Colors.Section.Body.Description.ValueString()
		} else {
			description1 = nil
		}
		body := shared.AppearanceColorVariable{
			Value:       value1,
			Description: description1,
		}
		var value2 string
		value2 = r.CustomTheme.Colors.Section.Hero.Value.ValueString()

		description2 := new(string)
		if !r.CustomTheme.Colors.Section.Hero.Description.IsUnknown() && !r.CustomTheme.Colors.Section.Hero.Description.IsNull() {
			*description2 = r.CustomTheme.Colors.Section.Hero.Description.ValueString()
		} else {
			description2 = nil
		}
		hero := shared.AppearanceColorVariable{
			Value:       value2,
			Description: description2,
		}
		var value3 string
		value3 = r.CustomTheme.Colors.Section.Accent.Value.ValueString()

		description3 := new(string)
		if !r.CustomTheme.Colors.Section.Accent.Description.IsUnknown() && !r.CustomTheme.Colors.Section.Accent.Description.IsNull() {
			*description3 = r.CustomTheme.Colors.Section.Accent.Description.ValueString()
		} else {
			description3 = nil
		}
		accent := shared.AppearanceColorVariable{
			Value:       value3,
			Description: description3,
		}
		var value4 string
		value4 = r.CustomTheme.Colors.Section.Tertiary.Value.ValueString()

		description4 := new(string)
		if !r.CustomTheme.Colors.Section.Tertiary.Description.IsUnknown() && !r.CustomTheme.Colors.Section.Tertiary.Description.IsNull() {
			*description4 = r.CustomTheme.Colors.Section.Tertiary.Description.ValueString()
		} else {
			description4 = nil
		}
		tertiary := shared.AppearanceColorVariable{
			Value:       value4,
			Description: description4,
		}
		var value5 string
		value5 = r.CustomTheme.Colors.Section.Stroke.Value.ValueString()

		description5 := new(string)
		if !r.CustomTheme.Colors.Section.Stroke.Description.IsUnknown() && !r.CustomTheme.Colors.Section.Stroke.Description.IsNull() {
			*description5 = r.CustomTheme.Colors.Section.Stroke.Description.ValueString()
		} else {
			description5 = nil
		}
		stroke := shared.AppearanceColorVariable{
			Value:       value5,
			Description: description5,
		}
		var value6 string
		value6 = r.CustomTheme.Colors.Section.Footer.Value.ValueString()

		description6 := new(string)
		if !r.CustomTheme.Colors.Section.Footer.Description.IsUnknown() && !r.CustomTheme.Colors.Section.Footer.Description.IsNull() {
			*description6 = r.CustomTheme.Colors.Section.Footer.Description.ValueString()
		} else {
			description6 = nil
		}
		footer := shared.AppearanceColorVariable{
			Value:       value6,
			Description: description6,
		}
		section := shared.Section{
			Header:   header,
			Body:     body,
			Hero:     hero,
			Accent:   accent,
			Tertiary: tertiary,
			Stroke:   stroke,
			Footer:   footer,
		}
		var value7 string
		value7 = r.CustomTheme.Colors.Text.Header.Value.ValueString()

		description7 := new(string)
		if !r.CustomTheme.Colors.Text.Header.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Header.Description.IsNull() {
			*description7 = r.CustomTheme.Colors.Text.Header.Description.ValueString()
		} else {
			description7 = nil
		}
		header1 := shared.AppearanceColorVariable{
			Value:       value7,
			Description: description7,
		}
		var value8 string
		value8 = r.CustomTheme.Colors.Text.Hero.Value.ValueString()

		description8 := new(string)
		if !r.CustomTheme.Colors.Text.Hero.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Hero.Description.IsNull() {
			*description8 = r.CustomTheme.Colors.Text.Hero.Description.ValueString()
		} else {
			description8 = nil
		}
		hero1 := shared.AppearanceColorVariable{
			Value:       value8,
			Description: description8,
		}
		var value9 string
		value9 = r.CustomTheme.Colors.Text.Headings.Value.ValueString()

		description9 := new(string)
		if !r.CustomTheme.Colors.Text.Headings.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Headings.Description.IsNull() {
			*description9 = r.CustomTheme.Colors.Text.Headings.Description.ValueString()
		} else {
			description9 = nil
		}
		headings := shared.AppearanceColorVariable{
			Value:       value9,
			Description: description9,
		}
		var value10 string
		value10 = r.CustomTheme.Colors.Text.Primary.Value.ValueString()

		description10 := new(string)
		if !r.CustomTheme.Colors.Text.Primary.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Primary.Description.IsNull() {
			*description10 = r.CustomTheme.Colors.Text.Primary.Description.ValueString()
		} else {
			description10 = nil
		}
		primary := shared.AppearanceColorVariable{
			Value:       value10,
			Description: description10,
		}
		var value11 string
		value11 = r.CustomTheme.Colors.Text.Secondary.Value.ValueString()

		description11 := new(string)
		if !r.CustomTheme.Colors.Text.Secondary.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Secondary.Description.IsNull() {
			*description11 = r.CustomTheme.Colors.Text.Secondary.Description.ValueString()
		} else {
			description11 = nil
		}
		secondary := shared.AppearanceColorVariable{
			Value:       value11,
			Description: description11,
		}
		var value12 string
		value12 = r.CustomTheme.Colors.Text.Accent.Value.ValueString()

		description12 := new(string)
		if !r.CustomTheme.Colors.Text.Accent.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Accent.Description.IsNull() {
			*description12 = r.CustomTheme.Colors.Text.Accent.Description.ValueString()
		} else {
			description12 = nil
		}
		accent1 := shared.AppearanceColorVariable{
			Value:       value12,
			Description: description12,
		}
		var value13 string
		value13 = r.CustomTheme.Colors.Text.Link.Value.ValueString()

		description13 := new(string)
		if !r.CustomTheme.Colors.Text.Link.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Link.Description.IsNull() {
			*description13 = r.CustomTheme.Colors.Text.Link.Description.ValueString()
		} else {
			description13 = nil
		}
		link := shared.AppearanceColorVariable{
			Value:       value13,
			Description: description13,
		}
		var value14 string
		value14 = r.CustomTheme.Colors.Text.Footer.Value.ValueString()

		description14 := new(string)
		if !r.CustomTheme.Colors.Text.Footer.Description.IsUnknown() && !r.CustomTheme.Colors.Text.Footer.Description.IsNull() {
			*description14 = r.CustomTheme.Colors.Text.Footer.Description.ValueString()
		} else {
			description14 = nil
		}
		footer1 := shared.AppearanceColorVariable{
			Value:       value14,
			Description: description14,
		}
		text := shared.Text{
			Header:    header1,
			Hero:      hero1,
			Headings:  headings,
			Primary:   primary,
			Secondary: secondary,
			Accent:    accent1,
			Link:      link,
			Footer:    footer1,
		}
		var value15 string
		value15 = r.CustomTheme.Colors.Button.PrimaryFill.Value.ValueString()

		description15 := new(string)
		if !r.CustomTheme.Colors.Button.PrimaryFill.Description.IsUnknown() && !r.CustomTheme.Colors.Button.PrimaryFill.Description.IsNull() {
			*description15 = r.CustomTheme.Colors.Button.PrimaryFill.Description.ValueString()
		} else {
			description15 = nil
		}
		primaryFill := shared.AppearanceColorVariable{
			Value:       value15,
			Description: description15,
		}
		var value16 string
		value16 = r.CustomTheme.Colors.Button.PrimaryText.Value.ValueString()

		description16 := new(string)
		if !r.CustomTheme.Colors.Button.PrimaryText.Description.IsUnknown() && !r.CustomTheme.Colors.Button.PrimaryText.Description.IsNull() {
			*description16 = r.CustomTheme.Colors.Button.PrimaryText.Description.ValueString()
		} else {
			description16 = nil
		}
		primaryText := shared.AppearanceColorVariable{
			Value:       value16,
			Description: description16,
		}
		button := shared.Button{
			PrimaryFill: primaryFill,
			PrimaryText: primaryText,
		}
		colors := shared.AppearanceThemeColorVariables{
			Section: section,
			Text:    text,
			Button:  button,
		}
		customTheme = &shared.NullableAppearanceThemeVariables{
			Colors: colors,
		}
	}
	var customFonts *shared.NullableAppearanceFonts
	if r.CustomFonts != nil {
		base := shared.AppearanceFontName(r.CustomFonts.Base.ValueString())
		code := shared.AppearanceFontName(r.CustomFonts.Code.ValueString())
		headings1 := shared.AppearanceFontName(r.CustomFonts.Headings.ValueString())
		customFonts = &shared.NullableAppearanceFonts{
			Base:     base,
			Code:     code,
			Headings: headings1,
		}
	}
	useCustomFonts := new(bool)
	if !r.UseCustomFonts.IsUnknown() && !r.UseCustomFonts.IsNull() {
		*useCustomFonts = r.UseCustomFonts.ValueBool()
	} else {
		useCustomFonts = nil
	}
	var text1 *shared.NullableAppearanceTextVariables
	if r.Text != nil {
		var welcomeMessage string
		welcomeMessage = r.Text.Catalog.WelcomeMessage.ValueString()

		var primaryHeader string
		primaryHeader = r.Text.Catalog.PrimaryHeader.ValueString()

		catalog := shared.Catalog{
			WelcomeMessage: welcomeMessage,
			PrimaryHeader:  primaryHeader,
		}
		text1 = &shared.NullableAppearanceTextVariables{
			Catalog: catalog,
		}
	}
	var images *shared.AppearanceImages
	if r.Images != nil {
		var logo *shared.AppearanceImage
		if r.Images.Logo != nil {
			var data string
			data = r.Images.Logo.Data.ValueString()

			filename := new(string)
			if !r.Images.Logo.Filename.IsUnknown() && !r.Images.Logo.Filename.IsNull() {
				*filename = r.Images.Logo.Filename.ValueString()
			} else {
				filename = nil
			}
			logo = &shared.AppearanceImage{
				Data:     data,
				Filename: filename,
			}
		}
		var favicon *shared.AppearanceImage
		if r.Images.Favicon != nil {
			var data1 string
			data1 = r.Images.Favicon.Data.ValueString()

			filename1 := new(string)
			if !r.Images.Favicon.Filename.IsUnknown() && !r.Images.Favicon.Filename.IsNull() {
				*filename1 = r.Images.Favicon.Filename.ValueString()
			} else {
				filename1 = nil
			}
			favicon = &shared.AppearanceImage{
				Data:     data1,
				Filename: filename1,
			}
		}
		var catalogCover *shared.AppearanceImage
		if r.Images.CatalogCover != nil {
			var data2 string
			data2 = r.Images.CatalogCover.Data.ValueString()

			filename2 := new(string)
			if !r.Images.CatalogCover.Filename.IsUnknown() && !r.Images.CatalogCover.Filename.IsNull() {
				*filename2 = r.Images.CatalogCover.Filename.ValueString()
			} else {
				filename2 = nil
			}
			catalogCover = &shared.AppearanceImage{
				Data:     data2,
				Filename: filename2,
			}
		}
		images = &shared.AppearanceImages{
			Logo:         logo,
			Favicon:      favicon,
			CatalogCover: catalogCover,
		}
	}
	out := shared.UpdatePortalAppearanceRequest{
		ThemeName:      themeName,
		CustomTheme:    customTheme,
		CustomFonts:    customFonts,
		UseCustomFonts: useCustomFonts,
		Text:           text1,
		Images:         images,
	}
	return &out
}

func (r *PortalAppearanceResourceModel) RefreshFromSharedUpdatePortalAppearanceResponse(resp *shared.UpdatePortalAppearanceResponse) {
	if resp != nil {
		if resp.CustomFonts == nil {
			r.CustomFonts = nil
		} else {
			r.CustomFonts = &tfTypes.NullableAppearanceFonts{}
			r.CustomFonts.Base = types.StringValue(string(resp.CustomFonts.Base))
			r.CustomFonts.Code = types.StringValue(string(resp.CustomFonts.Code))
			r.CustomFonts.Headings = types.StringValue(string(resp.CustomFonts.Headings))
		}
		if resp.CustomTheme == nil {
			r.CustomTheme = nil
		} else {
			r.CustomTheme = &tfTypes.NullableAppearanceThemeVariables{}
			r.CustomTheme.Colors.Button.PrimaryFill.Description = types.StringPointerValue(resp.CustomTheme.Colors.Button.PrimaryFill.Description)
			r.CustomTheme.Colors.Button.PrimaryFill.Value = types.StringValue(resp.CustomTheme.Colors.Button.PrimaryFill.Value)
			r.CustomTheme.Colors.Button.PrimaryText.Description = types.StringPointerValue(resp.CustomTheme.Colors.Button.PrimaryText.Description)
			r.CustomTheme.Colors.Button.PrimaryText.Value = types.StringValue(resp.CustomTheme.Colors.Button.PrimaryText.Value)
			r.CustomTheme.Colors.Section.Accent.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Accent.Description)
			r.CustomTheme.Colors.Section.Accent.Value = types.StringValue(resp.CustomTheme.Colors.Section.Accent.Value)
			r.CustomTheme.Colors.Section.Body.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Body.Description)
			r.CustomTheme.Colors.Section.Body.Value = types.StringValue(resp.CustomTheme.Colors.Section.Body.Value)
			r.CustomTheme.Colors.Section.Footer.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Footer.Description)
			r.CustomTheme.Colors.Section.Footer.Value = types.StringValue(resp.CustomTheme.Colors.Section.Footer.Value)
			r.CustomTheme.Colors.Section.Header.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Header.Description)
			r.CustomTheme.Colors.Section.Header.Value = types.StringValue(resp.CustomTheme.Colors.Section.Header.Value)
			r.CustomTheme.Colors.Section.Hero.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Hero.Description)
			r.CustomTheme.Colors.Section.Hero.Value = types.StringValue(resp.CustomTheme.Colors.Section.Hero.Value)
			r.CustomTheme.Colors.Section.Stroke.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Stroke.Description)
			r.CustomTheme.Colors.Section.Stroke.Value = types.StringValue(resp.CustomTheme.Colors.Section.Stroke.Value)
			r.CustomTheme.Colors.Section.Tertiary.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Tertiary.Description)
			r.CustomTheme.Colors.Section.Tertiary.Value = types.StringValue(resp.CustomTheme.Colors.Section.Tertiary.Value)
			r.CustomTheme.Colors.Text.Accent.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Accent.Description)
			r.CustomTheme.Colors.Text.Accent.Value = types.StringValue(resp.CustomTheme.Colors.Text.Accent.Value)
			r.CustomTheme.Colors.Text.Footer.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Footer.Description)
			r.CustomTheme.Colors.Text.Footer.Value = types.StringValue(resp.CustomTheme.Colors.Text.Footer.Value)
			r.CustomTheme.Colors.Text.Header.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Header.Description)
			r.CustomTheme.Colors.Text.Header.Value = types.StringValue(resp.CustomTheme.Colors.Text.Header.Value)
			r.CustomTheme.Colors.Text.Headings.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Headings.Description)
			r.CustomTheme.Colors.Text.Headings.Value = types.StringValue(resp.CustomTheme.Colors.Text.Headings.Value)
			r.CustomTheme.Colors.Text.Hero.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Hero.Description)
			r.CustomTheme.Colors.Text.Hero.Value = types.StringValue(resp.CustomTheme.Colors.Text.Hero.Value)
			r.CustomTheme.Colors.Text.Link.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Link.Description)
			r.CustomTheme.Colors.Text.Link.Value = types.StringValue(resp.CustomTheme.Colors.Text.Link.Value)
			r.CustomTheme.Colors.Text.Primary.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Primary.Description)
			r.CustomTheme.Colors.Text.Primary.Value = types.StringValue(resp.CustomTheme.Colors.Text.Primary.Value)
			r.CustomTheme.Colors.Text.Secondary.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Secondary.Description)
			r.CustomTheme.Colors.Text.Secondary.Value = types.StringValue(resp.CustomTheme.Colors.Text.Secondary.Value)
		}
		if resp.Images == nil {
			r.Images = nil
		} else {
			r.Images = &tfTypes.AppearanceImages{}
			if resp.Images.CatalogCover == nil {
				r.Images.CatalogCover = nil
			} else {
				r.Images.CatalogCover = &tfTypes.AppearanceImage{}
				r.Images.CatalogCover.Data = types.StringValue(resp.Images.CatalogCover.Data)
				r.Images.CatalogCover.Filename = types.StringPointerValue(resp.Images.CatalogCover.Filename)
			}
			if resp.Images.Favicon == nil {
				r.Images.Favicon = nil
			} else {
				r.Images.Favicon = &tfTypes.AppearanceImage{}
				r.Images.Favicon.Data = types.StringValue(resp.Images.Favicon.Data)
				r.Images.Favicon.Filename = types.StringPointerValue(resp.Images.Favicon.Filename)
			}
			if resp.Images.Logo == nil {
				r.Images.Logo = nil
			} else {
				r.Images.Logo = &tfTypes.AppearanceImage{}
				r.Images.Logo.Data = types.StringValue(resp.Images.Logo.Data)
				r.Images.Logo.Filename = types.StringPointerValue(resp.Images.Logo.Filename)
			}
		}
		if resp.Text == nil {
			r.Text = nil
		} else {
			r.Text = &tfTypes.NullableAppearanceTextVariables{}
			r.Text.Catalog.PrimaryHeader = types.StringValue(resp.Text.Catalog.PrimaryHeader)
			r.Text.Catalog.WelcomeMessage = types.StringValue(resp.Text.Catalog.WelcomeMessage)
		}
		r.ThemeName = types.StringValue(string(resp.ThemeName))
		r.UseCustomFonts = types.BoolValue(resp.UseCustomFonts)
	}
}

func (r *PortalAppearanceResourceModel) RefreshFromSharedGetPortalAppearanceResponse(resp *shared.GetPortalAppearanceResponse) {
	if resp != nil {
		if resp.CustomFonts == nil {
			r.CustomFonts = nil
		} else {
			r.CustomFonts = &tfTypes.NullableAppearanceFonts{}
			r.CustomFonts.Base = types.StringValue(string(resp.CustomFonts.Base))
			r.CustomFonts.Code = types.StringValue(string(resp.CustomFonts.Code))
			r.CustomFonts.Headings = types.StringValue(string(resp.CustomFonts.Headings))
		}
		if resp.CustomTheme == nil {
			r.CustomTheme = nil
		} else {
			r.CustomTheme = &tfTypes.NullableAppearanceThemeVariables{}
			r.CustomTheme.Colors.Button.PrimaryFill.Description = types.StringPointerValue(resp.CustomTheme.Colors.Button.PrimaryFill.Description)
			r.CustomTheme.Colors.Button.PrimaryFill.Value = types.StringValue(resp.CustomTheme.Colors.Button.PrimaryFill.Value)
			r.CustomTheme.Colors.Button.PrimaryText.Description = types.StringPointerValue(resp.CustomTheme.Colors.Button.PrimaryText.Description)
			r.CustomTheme.Colors.Button.PrimaryText.Value = types.StringValue(resp.CustomTheme.Colors.Button.PrimaryText.Value)
			r.CustomTheme.Colors.Section.Accent.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Accent.Description)
			r.CustomTheme.Colors.Section.Accent.Value = types.StringValue(resp.CustomTheme.Colors.Section.Accent.Value)
			r.CustomTheme.Colors.Section.Body.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Body.Description)
			r.CustomTheme.Colors.Section.Body.Value = types.StringValue(resp.CustomTheme.Colors.Section.Body.Value)
			r.CustomTheme.Colors.Section.Footer.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Footer.Description)
			r.CustomTheme.Colors.Section.Footer.Value = types.StringValue(resp.CustomTheme.Colors.Section.Footer.Value)
			r.CustomTheme.Colors.Section.Header.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Header.Description)
			r.CustomTheme.Colors.Section.Header.Value = types.StringValue(resp.CustomTheme.Colors.Section.Header.Value)
			r.CustomTheme.Colors.Section.Hero.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Hero.Description)
			r.CustomTheme.Colors.Section.Hero.Value = types.StringValue(resp.CustomTheme.Colors.Section.Hero.Value)
			r.CustomTheme.Colors.Section.Stroke.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Stroke.Description)
			r.CustomTheme.Colors.Section.Stroke.Value = types.StringValue(resp.CustomTheme.Colors.Section.Stroke.Value)
			r.CustomTheme.Colors.Section.Tertiary.Description = types.StringPointerValue(resp.CustomTheme.Colors.Section.Tertiary.Description)
			r.CustomTheme.Colors.Section.Tertiary.Value = types.StringValue(resp.CustomTheme.Colors.Section.Tertiary.Value)
			r.CustomTheme.Colors.Text.Accent.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Accent.Description)
			r.CustomTheme.Colors.Text.Accent.Value = types.StringValue(resp.CustomTheme.Colors.Text.Accent.Value)
			r.CustomTheme.Colors.Text.Footer.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Footer.Description)
			r.CustomTheme.Colors.Text.Footer.Value = types.StringValue(resp.CustomTheme.Colors.Text.Footer.Value)
			r.CustomTheme.Colors.Text.Header.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Header.Description)
			r.CustomTheme.Colors.Text.Header.Value = types.StringValue(resp.CustomTheme.Colors.Text.Header.Value)
			r.CustomTheme.Colors.Text.Headings.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Headings.Description)
			r.CustomTheme.Colors.Text.Headings.Value = types.StringValue(resp.CustomTheme.Colors.Text.Headings.Value)
			r.CustomTheme.Colors.Text.Hero.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Hero.Description)
			r.CustomTheme.Colors.Text.Hero.Value = types.StringValue(resp.CustomTheme.Colors.Text.Hero.Value)
			r.CustomTheme.Colors.Text.Link.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Link.Description)
			r.CustomTheme.Colors.Text.Link.Value = types.StringValue(resp.CustomTheme.Colors.Text.Link.Value)
			r.CustomTheme.Colors.Text.Primary.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Primary.Description)
			r.CustomTheme.Colors.Text.Primary.Value = types.StringValue(resp.CustomTheme.Colors.Text.Primary.Value)
			r.CustomTheme.Colors.Text.Secondary.Description = types.StringPointerValue(resp.CustomTheme.Colors.Text.Secondary.Description)
			r.CustomTheme.Colors.Text.Secondary.Value = types.StringValue(resp.CustomTheme.Colors.Text.Secondary.Value)
		}
		if resp.Images == nil {
			r.Images = nil
		} else {
			r.Images = &tfTypes.AppearanceImages{}
			if resp.Images.CatalogCover == nil {
				r.Images.CatalogCover = nil
			} else {
				r.Images.CatalogCover = &tfTypes.AppearanceImage{}
				r.Images.CatalogCover.Data = types.StringValue(resp.Images.CatalogCover.Data)
				r.Images.CatalogCover.Filename = types.StringPointerValue(resp.Images.CatalogCover.Filename)
			}
			if resp.Images.Favicon == nil {
				r.Images.Favicon = nil
			} else {
				r.Images.Favicon = &tfTypes.AppearanceImage{}
				r.Images.Favicon.Data = types.StringValue(resp.Images.Favicon.Data)
				r.Images.Favicon.Filename = types.StringPointerValue(resp.Images.Favicon.Filename)
			}
			if resp.Images.Logo == nil {
				r.Images.Logo = nil
			} else {
				r.Images.Logo = &tfTypes.AppearanceImage{}
				r.Images.Logo.Data = types.StringValue(resp.Images.Logo.Data)
				r.Images.Logo.Filename = types.StringPointerValue(resp.Images.Logo.Filename)
			}
		}
		if resp.Text == nil {
			r.Text = nil
		} else {
			r.Text = &tfTypes.NullableAppearanceTextVariables{}
			r.Text.Catalog.PrimaryHeader = types.StringValue(resp.Text.Catalog.PrimaryHeader)
			r.Text.Catalog.WelcomeMessage = types.StringValue(resp.Text.Catalog.WelcomeMessage)
		}
		r.ThemeName = types.StringValue(string(resp.ThemeName))
		r.UseCustomFonts = types.BoolValue(resp.UseCustomFonts)
	}
}