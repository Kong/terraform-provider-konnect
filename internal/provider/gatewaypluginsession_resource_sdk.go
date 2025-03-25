// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginSessionResourceModel) ToSharedSessionPluginInput() *shared.SessionPluginInput {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	ordering := make(map[string]string)
	for orderingKey, orderingValue := range r.Ordering {
		var orderingInst string
		orderingInst = orderingValue.ValueString()

		ordering[orderingKey] = orderingInst
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	absoluteTimeout := new(float64)
	if !r.Config.AbsoluteTimeout.IsUnknown() && !r.Config.AbsoluteTimeout.IsNull() {
		*absoluteTimeout, _ = r.Config.AbsoluteTimeout.ValueBigFloat().Float64()
	} else {
		absoluteTimeout = nil
	}
	audience := new(string)
	if !r.Config.Audience.IsUnknown() && !r.Config.Audience.IsNull() {
		*audience = r.Config.Audience.ValueString()
	} else {
		audience = nil
	}
	cookieDomain := new(string)
	if !r.Config.CookieDomain.IsUnknown() && !r.Config.CookieDomain.IsNull() {
		*cookieDomain = r.Config.CookieDomain.ValueString()
	} else {
		cookieDomain = nil
	}
	cookieHTTPOnly := new(bool)
	if !r.Config.CookieHTTPOnly.IsUnknown() && !r.Config.CookieHTTPOnly.IsNull() {
		*cookieHTTPOnly = r.Config.CookieHTTPOnly.ValueBool()
	} else {
		cookieHTTPOnly = nil
	}
	cookieName := new(string)
	if !r.Config.CookieName.IsUnknown() && !r.Config.CookieName.IsNull() {
		*cookieName = r.Config.CookieName.ValueString()
	} else {
		cookieName = nil
	}
	cookiePath := new(string)
	if !r.Config.CookiePath.IsUnknown() && !r.Config.CookiePath.IsNull() {
		*cookiePath = r.Config.CookiePath.ValueString()
	} else {
		cookiePath = nil
	}
	cookieSameSite := new(shared.CookieSameSite)
	if !r.Config.CookieSameSite.IsUnknown() && !r.Config.CookieSameSite.IsNull() {
		*cookieSameSite = shared.CookieSameSite(r.Config.CookieSameSite.ValueString())
	} else {
		cookieSameSite = nil
	}
	cookieSecure := new(bool)
	if !r.Config.CookieSecure.IsUnknown() && !r.Config.CookieSecure.IsNull() {
		*cookieSecure = r.Config.CookieSecure.ValueBool()
	} else {
		cookieSecure = nil
	}
	idlingTimeout := new(float64)
	if !r.Config.IdlingTimeout.IsUnknown() && !r.Config.IdlingTimeout.IsNull() {
		*idlingTimeout, _ = r.Config.IdlingTimeout.ValueBigFloat().Float64()
	} else {
		idlingTimeout = nil
	}
	var logoutMethods []shared.SessionPluginLogoutMethods = []shared.SessionPluginLogoutMethods{}
	for _, logoutMethodsItem := range r.Config.LogoutMethods {
		logoutMethods = append(logoutMethods, shared.SessionPluginLogoutMethods(logoutMethodsItem.ValueString()))
	}
	logoutPostArg := new(string)
	if !r.Config.LogoutPostArg.IsUnknown() && !r.Config.LogoutPostArg.IsNull() {
		*logoutPostArg = r.Config.LogoutPostArg.ValueString()
	} else {
		logoutPostArg = nil
	}
	logoutQueryArg := new(string)
	if !r.Config.LogoutQueryArg.IsUnknown() && !r.Config.LogoutQueryArg.IsNull() {
		*logoutQueryArg = r.Config.LogoutQueryArg.ValueString()
	} else {
		logoutQueryArg = nil
	}
	readBodyForLogout := new(bool)
	if !r.Config.ReadBodyForLogout.IsUnknown() && !r.Config.ReadBodyForLogout.IsNull() {
		*readBodyForLogout = r.Config.ReadBodyForLogout.ValueBool()
	} else {
		readBodyForLogout = nil
	}
	remember := new(bool)
	if !r.Config.Remember.IsUnknown() && !r.Config.Remember.IsNull() {
		*remember = r.Config.Remember.ValueBool()
	} else {
		remember = nil
	}
	rememberAbsoluteTimeout := new(float64)
	if !r.Config.RememberAbsoluteTimeout.IsUnknown() && !r.Config.RememberAbsoluteTimeout.IsNull() {
		*rememberAbsoluteTimeout, _ = r.Config.RememberAbsoluteTimeout.ValueBigFloat().Float64()
	} else {
		rememberAbsoluteTimeout = nil
	}
	rememberCookieName := new(string)
	if !r.Config.RememberCookieName.IsUnknown() && !r.Config.RememberCookieName.IsNull() {
		*rememberCookieName = r.Config.RememberCookieName.ValueString()
	} else {
		rememberCookieName = nil
	}
	rememberRollingTimeout := new(float64)
	if !r.Config.RememberRollingTimeout.IsUnknown() && !r.Config.RememberRollingTimeout.IsNull() {
		*rememberRollingTimeout, _ = r.Config.RememberRollingTimeout.ValueBigFloat().Float64()
	} else {
		rememberRollingTimeout = nil
	}
	var requestHeaders []shared.RequestHeaders = []shared.RequestHeaders{}
	for _, requestHeadersItem := range r.Config.RequestHeaders {
		requestHeaders = append(requestHeaders, shared.RequestHeaders(requestHeadersItem.ValueString()))
	}
	var responseHeaders []shared.SessionPluginResponseHeaders = []shared.SessionPluginResponseHeaders{}
	for _, responseHeadersItem := range r.Config.ResponseHeaders {
		responseHeaders = append(responseHeaders, shared.SessionPluginResponseHeaders(responseHeadersItem.ValueString()))
	}
	rollingTimeout := new(float64)
	if !r.Config.RollingTimeout.IsUnknown() && !r.Config.RollingTimeout.IsNull() {
		*rollingTimeout, _ = r.Config.RollingTimeout.ValueBigFloat().Float64()
	} else {
		rollingTimeout = nil
	}
	secret := new(string)
	if !r.Config.Secret.IsUnknown() && !r.Config.Secret.IsNull() {
		*secret = r.Config.Secret.ValueString()
	} else {
		secret = nil
	}
	staleTTL := new(float64)
	if !r.Config.StaleTTL.IsUnknown() && !r.Config.StaleTTL.IsNull() {
		*staleTTL, _ = r.Config.StaleTTL.ValueBigFloat().Float64()
	} else {
		staleTTL = nil
	}
	storage := new(shared.SessionPluginStorage)
	if !r.Config.Storage.IsUnknown() && !r.Config.Storage.IsNull() {
		*storage = shared.SessionPluginStorage(r.Config.Storage.ValueString())
	} else {
		storage = nil
	}
	config := shared.SessionPluginConfig{
		AbsoluteTimeout:         absoluteTimeout,
		Audience:                audience,
		CookieDomain:            cookieDomain,
		CookieHTTPOnly:          cookieHTTPOnly,
		CookieName:              cookieName,
		CookiePath:              cookiePath,
		CookieSameSite:          cookieSameSite,
		CookieSecure:            cookieSecure,
		IdlingTimeout:           idlingTimeout,
		LogoutMethods:           logoutMethods,
		LogoutPostArg:           logoutPostArg,
		LogoutQueryArg:          logoutQueryArg,
		ReadBodyForLogout:       readBodyForLogout,
		Remember:                remember,
		RememberAbsoluteTimeout: rememberAbsoluteTimeout,
		RememberCookieName:      rememberCookieName,
		RememberRollingTimeout:  rememberRollingTimeout,
		RequestHeaders:          requestHeaders,
		ResponseHeaders:         responseHeaders,
		RollingTimeout:          rollingTimeout,
		Secret:                  secret,
		StaleTTL:                staleTTL,
		Storage:                 storage,
	}
	var protocols []shared.SessionPluginProtocols = []shared.SessionPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.SessionPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.SessionPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.SessionPluginRoute{
			ID: id1,
		}
	}
	var service *shared.SessionPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.SessionPluginService{
			ID: id2,
		}
	}
	out := shared.SessionPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginSessionResourceModel) RefreshFromSharedSessionPlugin(resp *shared.SessionPlugin) {
	if resp != nil {
		if resp.Config.AbsoluteTimeout != nil {
			r.Config.AbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.AbsoluteTimeout)))
		} else {
			r.Config.AbsoluteTimeout = types.NumberNull()
		}
		r.Config.Audience = types.StringPointerValue(resp.Config.Audience)
		r.Config.CookieDomain = types.StringPointerValue(resp.Config.CookieDomain)
		r.Config.CookieHTTPOnly = types.BoolPointerValue(resp.Config.CookieHTTPOnly)
		r.Config.CookieName = types.StringPointerValue(resp.Config.CookieName)
		r.Config.CookiePath = types.StringPointerValue(resp.Config.CookiePath)
		if resp.Config.CookieSameSite != nil {
			r.Config.CookieSameSite = types.StringValue(string(*resp.Config.CookieSameSite))
		} else {
			r.Config.CookieSameSite = types.StringNull()
		}
		r.Config.CookieSecure = types.BoolPointerValue(resp.Config.CookieSecure)
		if resp.Config.IdlingTimeout != nil {
			r.Config.IdlingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.IdlingTimeout)))
		} else {
			r.Config.IdlingTimeout = types.NumberNull()
		}
		r.Config.LogoutMethods = make([]types.String, 0, len(resp.Config.LogoutMethods))
		for _, v := range resp.Config.LogoutMethods {
			r.Config.LogoutMethods = append(r.Config.LogoutMethods, types.StringValue(string(v)))
		}
		r.Config.LogoutPostArg = types.StringPointerValue(resp.Config.LogoutPostArg)
		r.Config.LogoutQueryArg = types.StringPointerValue(resp.Config.LogoutQueryArg)
		r.Config.ReadBodyForLogout = types.BoolPointerValue(resp.Config.ReadBodyForLogout)
		r.Config.Remember = types.BoolPointerValue(resp.Config.Remember)
		if resp.Config.RememberAbsoluteTimeout != nil {
			r.Config.RememberAbsoluteTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.RememberAbsoluteTimeout)))
		} else {
			r.Config.RememberAbsoluteTimeout = types.NumberNull()
		}
		r.Config.RememberCookieName = types.StringPointerValue(resp.Config.RememberCookieName)
		if resp.Config.RememberRollingTimeout != nil {
			r.Config.RememberRollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.RememberRollingTimeout)))
		} else {
			r.Config.RememberRollingTimeout = types.NumberNull()
		}
		r.Config.RequestHeaders = make([]types.String, 0, len(resp.Config.RequestHeaders))
		for _, v := range resp.Config.RequestHeaders {
			r.Config.RequestHeaders = append(r.Config.RequestHeaders, types.StringValue(string(v)))
		}
		r.Config.ResponseHeaders = make([]types.String, 0, len(resp.Config.ResponseHeaders))
		for _, v := range resp.Config.ResponseHeaders {
			r.Config.ResponseHeaders = append(r.Config.ResponseHeaders, types.StringValue(string(v)))
		}
		if resp.Config.RollingTimeout != nil {
			r.Config.RollingTimeout = types.NumberValue(big.NewFloat(float64(*resp.Config.RollingTimeout)))
		} else {
			r.Config.RollingTimeout = types.NumberNull()
		}
		r.Config.Secret = types.StringPointerValue(resp.Config.Secret)
		if resp.Config.StaleTTL != nil {
			r.Config.StaleTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.StaleTTL)))
		} else {
			r.Config.StaleTTL = types.NumberNull()
		}
		if resp.Config.Storage != nil {
			r.Config.Storage = types.StringValue(string(*resp.Config.Storage))
		} else {
			r.Config.Storage = types.StringNull()
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering != nil {
			r.Ordering = make(map[string]types.String, len(resp.Ordering))
			for key, value := range resp.Ordering {
				r.Ordering[key] = types.StringValue(value)
			}
		}
		r.Protocols = make([]types.String, 0, len(resp.Protocols))
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLWithoutParentsConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLWithoutParentsConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = make([]types.String, 0, len(resp.Tags))
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
