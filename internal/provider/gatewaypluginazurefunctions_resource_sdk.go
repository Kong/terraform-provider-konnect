// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/v2/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/v2/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginAzureFunctionsResourceModel) ToSharedAzureFunctionsPluginInput() *shared.AzureFunctionsPluginInput {
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
	var ordering *shared.AzureFunctionsPluginOrdering
	if r.Ordering != nil {
		var after *shared.AzureFunctionsPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.AzureFunctionsPluginAfter{
				Access: access,
			}
		}
		var before *shared.AzureFunctionsPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.AzureFunctionsPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.AzureFunctionsPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	apikey := new(string)
	if !r.Config.Apikey.IsUnknown() && !r.Config.Apikey.IsNull() {
		*apikey = r.Config.Apikey.ValueString()
	} else {
		apikey = nil
	}
	appname := new(string)
	if !r.Config.Appname.IsUnknown() && !r.Config.Appname.IsNull() {
		*appname = r.Config.Appname.ValueString()
	} else {
		appname = nil
	}
	clientid := new(string)
	if !r.Config.Clientid.IsUnknown() && !r.Config.Clientid.IsNull() {
		*clientid = r.Config.Clientid.ValueString()
	} else {
		clientid = nil
	}
	functionname := new(string)
	if !r.Config.Functionname.IsUnknown() && !r.Config.Functionname.IsNull() {
		*functionname = r.Config.Functionname.ValueString()
	} else {
		functionname = nil
	}
	hostdomain := new(string)
	if !r.Config.Hostdomain.IsUnknown() && !r.Config.Hostdomain.IsNull() {
		*hostdomain = r.Config.Hostdomain.ValueString()
	} else {
		hostdomain = nil
	}
	https := new(bool)
	if !r.Config.HTTPS.IsUnknown() && !r.Config.HTTPS.IsNull() {
		*https = r.Config.HTTPS.ValueBool()
	} else {
		https = nil
	}
	httpsVerify := new(bool)
	if !r.Config.HTTPSVerify.IsUnknown() && !r.Config.HTTPSVerify.IsNull() {
		*httpsVerify = r.Config.HTTPSVerify.ValueBool()
	} else {
		httpsVerify = nil
	}
	keepalive := new(float64)
	if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
		*keepalive, _ = r.Config.Keepalive.ValueBigFloat().Float64()
	} else {
		keepalive = nil
	}
	routeprefix := new(string)
	if !r.Config.Routeprefix.IsUnknown() && !r.Config.Routeprefix.IsNull() {
		*routeprefix = r.Config.Routeprefix.ValueString()
	} else {
		routeprefix = nil
	}
	timeout := new(float64)
	if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
		*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
	} else {
		timeout = nil
	}
	config := shared.AzureFunctionsPluginConfig{
		Apikey:       apikey,
		Appname:      appname,
		Clientid:     clientid,
		Functionname: functionname,
		Hostdomain:   hostdomain,
		HTTPS:        https,
		HTTPSVerify:  httpsVerify,
		Keepalive:    keepalive,
		Routeprefix:  routeprefix,
		Timeout:      timeout,
	}
	var consumer *shared.AzureFunctionsPluginConsumer
	if r.Consumer != nil {
		id1 := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id1 = r.Consumer.ID.ValueString()
		} else {
			id1 = nil
		}
		consumer = &shared.AzureFunctionsPluginConsumer{
			ID: id1,
		}
	}
	var protocols []shared.AzureFunctionsPluginProtocols = []shared.AzureFunctionsPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.AzureFunctionsPluginProtocols(protocolsItem.ValueString()))
	}
	var route *shared.AzureFunctionsPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.AzureFunctionsPluginRoute{
			ID: id2,
		}
	}
	var service *shared.AzureFunctionsPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.AzureFunctionsPluginService{
			ID: id3,
		}
	}
	out := shared.AzureFunctionsPluginInput{
		Enabled:      enabled,
		ID:           id,
		InstanceName: instanceName,
		Ordering:     ordering,
		Tags:         tags,
		Config:       config,
		Consumer:     consumer,
		Protocols:    protocols,
		Route:        route,
		Service:      service,
	}
	return &out
}

func (r *GatewayPluginAzureFunctionsResourceModel) RefreshFromSharedAzureFunctionsPlugin(resp *shared.AzureFunctionsPlugin) {
	if resp != nil {
		r.Config.Apikey = types.StringPointerValue(resp.Config.Apikey)
		r.Config.Appname = types.StringPointerValue(resp.Config.Appname)
		r.Config.Clientid = types.StringPointerValue(resp.Config.Clientid)
		r.Config.Functionname = types.StringPointerValue(resp.Config.Functionname)
		r.Config.Hostdomain = types.StringPointerValue(resp.Config.Hostdomain)
		r.Config.HTTPS = types.BoolPointerValue(resp.Config.HTTPS)
		r.Config.HTTPSVerify = types.BoolPointerValue(resp.Config.HTTPSVerify)
		if resp.Config.Keepalive != nil {
			r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
		} else {
			r.Config.Keepalive = types.NumberNull()
		}
		r.Config.Routeprefix = types.StringPointerValue(resp.Config.Routeprefix)
		if resp.Config.Timeout != nil {
			r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
		} else {
			r.Config.Timeout = types.NumberNull()
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLWithoutParentsConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.ACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.ACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.ACLPluginAfter{}
				r.Ordering.Before.Access = []types.String{}
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = []types.String{}
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
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
