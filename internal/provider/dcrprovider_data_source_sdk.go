// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"time"
)

func (r *DcrProviderDataSourceModel) RefreshFromSharedDcrProviderResponse(resp *shared.DcrProviderResponse) {
	if resp != nil {
		if resp.DCRProviderAuth0DCRProviderAuth0 != nil {
			r.Auth0 = &tfTypes.DCRProviderAuth0DCRProviderAuth0{}
			r.Auth0.Active = types.BoolValue(resp.DCRProviderAuth0DCRProviderAuth0.Active)
			r.Active = r.Auth0.Active
			r.Auth0.CreatedAt = types.StringValue(resp.DCRProviderAuth0DCRProviderAuth0.CreatedAt.Format(time.RFC3339Nano))
			r.Auth0.DcrConfig.InitialClientAudience = types.StringPointerValue(resp.DCRProviderAuth0DCRProviderAuth0.DcrConfig.InitialClientAudience)
			r.Auth0.DcrConfig.InitialClientID = types.StringValue(resp.DCRProviderAuth0DCRProviderAuth0.DcrConfig.InitialClientID)
			r.Auth0.DcrConfig.UseDeveloperManagedScopes = types.BoolValue(resp.DCRProviderAuth0DCRProviderAuth0.DcrConfig.UseDeveloperManagedScopes)
			r.Auth0.DisplayName = types.StringPointerValue(resp.DCRProviderAuth0DCRProviderAuth0.DisplayName)
			r.DisplayName = r.Auth0.DisplayName
			r.Auth0.ID = types.StringValue(resp.DCRProviderAuth0DCRProviderAuth0.ID)
			r.ID = r.Auth0.ID
			r.Auth0.Issuer = types.StringValue(resp.DCRProviderAuth0DCRProviderAuth0.Issuer)
			r.Issuer = r.Auth0.Issuer
			if len(resp.DCRProviderAuth0DCRProviderAuth0.Labels) > 0 {
				r.Auth0.Labels = make(map[string]types.String)
				for key, value := range resp.DCRProviderAuth0DCRProviderAuth0.Labels {
					r.Auth0.Labels[key] = types.StringValue(value)
				}
			}
			r.Auth0.Name = types.StringValue(resp.DCRProviderAuth0DCRProviderAuth0.Name)
			r.Name = r.Auth0.Name
			r.Auth0.ProviderType = types.StringValue(string(resp.DCRProviderAuth0DCRProviderAuth0.ProviderType))
			r.Auth0.UpdatedAt = types.StringValue(resp.DCRProviderAuth0DCRProviderAuth0.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.DCRProviderAzureADDCRProviderAzureAD != nil {
			r.AzureAd = &tfTypes.DCRProviderAzureADDCRProviderAzureAD{}
			r.AzureAd.Active = types.BoolValue(resp.DCRProviderAzureADDCRProviderAzureAD.Active)
			r.Active = r.AzureAd.Active
			r.AzureAd.CreatedAt = types.StringValue(resp.DCRProviderAzureADDCRProviderAzureAD.CreatedAt.Format(time.RFC3339Nano))
			r.AzureAd.DcrConfig.InitialClientID = types.StringValue(resp.DCRProviderAzureADDCRProviderAzureAD.DcrConfig.InitialClientID)
			r.AzureAd.DisplayName = types.StringPointerValue(resp.DCRProviderAzureADDCRProviderAzureAD.DisplayName)
			r.DisplayName = r.AzureAd.DisplayName
			r.AzureAd.ID = types.StringValue(resp.DCRProviderAzureADDCRProviderAzureAD.ID)
			r.ID = r.AzureAd.ID
			r.AzureAd.Issuer = types.StringValue(resp.DCRProviderAzureADDCRProviderAzureAD.Issuer)
			r.Issuer = r.AzureAd.Issuer
			if len(resp.DCRProviderAzureADDCRProviderAzureAD.Labels) > 0 {
				r.AzureAd.Labels = make(map[string]types.String)
				for key1, value1 := range resp.DCRProviderAzureADDCRProviderAzureAD.Labels {
					r.AzureAd.Labels[key1] = types.StringValue(value1)
				}
			}
			r.AzureAd.Name = types.StringValue(resp.DCRProviderAzureADDCRProviderAzureAD.Name)
			r.Name = r.AzureAd.Name
			r.AzureAd.ProviderType = types.StringValue(string(resp.DCRProviderAzureADDCRProviderAzureAD.ProviderType))
			r.AzureAd.UpdatedAt = types.StringValue(resp.DCRProviderAzureADDCRProviderAzureAD.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.DCRProviderCurityDCRProviderCurity != nil {
			r.Curity = &tfTypes.DCRProviderAzureADDCRProviderAzureAD{}
			r.Curity.Active = types.BoolValue(resp.DCRProviderCurityDCRProviderCurity.Active)
			r.Active = r.Curity.Active
			r.Curity.CreatedAt = types.StringValue(resp.DCRProviderCurityDCRProviderCurity.CreatedAt.Format(time.RFC3339Nano))
			r.Curity.DcrConfig.InitialClientID = types.StringValue(resp.DCRProviderCurityDCRProviderCurity.DcrConfig.InitialClientID)
			r.Curity.DisplayName = types.StringPointerValue(resp.DCRProviderCurityDCRProviderCurity.DisplayName)
			r.DisplayName = r.Curity.DisplayName
			r.Curity.ID = types.StringValue(resp.DCRProviderCurityDCRProviderCurity.ID)
			r.ID = r.Curity.ID
			r.Curity.Issuer = types.StringValue(resp.DCRProviderCurityDCRProviderCurity.Issuer)
			r.Issuer = r.Curity.Issuer
			if len(resp.DCRProviderCurityDCRProviderCurity.Labels) > 0 {
				r.Curity.Labels = make(map[string]types.String)
				for key2, value2 := range resp.DCRProviderCurityDCRProviderCurity.Labels {
					r.Curity.Labels[key2] = types.StringValue(value2)
				}
			}
			r.Curity.Name = types.StringValue(resp.DCRProviderCurityDCRProviderCurity.Name)
			r.Name = r.Curity.Name
			r.Curity.ProviderType = types.StringValue(string(resp.DCRProviderCurityDCRProviderCurity.ProviderType))
			r.Curity.UpdatedAt = types.StringValue(resp.DCRProviderCurityDCRProviderCurity.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.DCRProviderHTTPDCRProviderHTTP != nil {
			r.HTTP = &tfTypes.DCRProviderHTTPDCRProviderHTTP{}
			r.HTTP.Active = types.BoolValue(resp.DCRProviderHTTPDCRProviderHTTP.Active)
			r.Active = r.HTTP.Active
			r.HTTP.CreatedAt = types.StringValue(resp.DCRProviderHTTPDCRProviderHTTP.CreatedAt.Format(time.RFC3339Nano))
			r.HTTP.DcrConfig.DcrBaseURL = types.StringValue(resp.DCRProviderHTTPDCRProviderHTTP.DcrConfig.DcrBaseURL)
			r.HTTP.DcrConfig.DisableEventHooks = types.BoolPointerValue(resp.DCRProviderHTTPDCRProviderHTTP.DcrConfig.DisableEventHooks)
			r.HTTP.DcrConfig.DisableRefreshSecret = types.BoolPointerValue(resp.DCRProviderHTTPDCRProviderHTTP.DcrConfig.DisableRefreshSecret)
			r.HTTP.DisplayName = types.StringPointerValue(resp.DCRProviderHTTPDCRProviderHTTP.DisplayName)
			r.DisplayName = r.HTTP.DisplayName
			r.HTTP.ID = types.StringValue(resp.DCRProviderHTTPDCRProviderHTTP.ID)
			r.ID = r.HTTP.ID
			r.HTTP.Issuer = types.StringValue(resp.DCRProviderHTTPDCRProviderHTTP.Issuer)
			r.Issuer = r.HTTP.Issuer
			if len(resp.DCRProviderHTTPDCRProviderHTTP.Labels) > 0 {
				r.HTTP.Labels = make(map[string]types.String)
				for key3, value3 := range resp.DCRProviderHTTPDCRProviderHTTP.Labels {
					r.HTTP.Labels[key3] = types.StringValue(value3)
				}
			}
			r.HTTP.Name = types.StringValue(resp.DCRProviderHTTPDCRProviderHTTP.Name)
			r.Name = r.HTTP.Name
			r.HTTP.ProviderType = types.StringValue(string(resp.DCRProviderHTTPDCRProviderHTTP.ProviderType))
			r.HTTP.UpdatedAt = types.StringValue(resp.DCRProviderHTTPDCRProviderHTTP.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.DCRProviderOKTADCRProviderOKTA != nil {
			r.Okta = &tfTypes.DCRProviderOKTADCRProviderOKTA{}
			r.Okta.Active = types.BoolValue(resp.DCRProviderOKTADCRProviderOKTA.Active)
			r.Active = r.Okta.Active
			r.Okta.CreatedAt = types.StringValue(resp.DCRProviderOKTADCRProviderOKTA.CreatedAt.Format(time.RFC3339Nano))
			r.Okta.DisplayName = types.StringPointerValue(resp.DCRProviderOKTADCRProviderOKTA.DisplayName)
			r.DisplayName = r.Okta.DisplayName
			r.Okta.ID = types.StringValue(resp.DCRProviderOKTADCRProviderOKTA.ID)
			r.ID = r.Okta.ID
			r.Okta.Issuer = types.StringValue(resp.DCRProviderOKTADCRProviderOKTA.Issuer)
			r.Issuer = r.Okta.Issuer
			if len(resp.DCRProviderOKTADCRProviderOKTA.Labels) > 0 {
				r.Okta.Labels = make(map[string]types.String)
				for key4, value4 := range resp.DCRProviderOKTADCRProviderOKTA.Labels {
					r.Okta.Labels[key4] = types.StringValue(value4)
				}
			}
			r.Okta.Name = types.StringValue(resp.DCRProviderOKTADCRProviderOKTA.Name)
			r.Name = r.Okta.Name
			r.Okta.ProviderType = types.StringValue(string(resp.DCRProviderOKTADCRProviderOKTA.ProviderType))
			r.Okta.UpdatedAt = types.StringValue(resp.DCRProviderOKTADCRProviderOKTA.UpdatedAt.Format(time.RFC3339Nano))
		}
	}
}
