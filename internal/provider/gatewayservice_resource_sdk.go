// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayServiceResourceModel) ToSharedServiceInput() *shared.ServiceInput {
	var caCertificates []string = []string{}
	for _, caCertificatesItem := range r.CaCertificates {
		caCertificates = append(caCertificates, caCertificatesItem.ValueString())
	}
	var clientCertificate *shared.ClientCertificate
	if r.ClientCertificate != nil {
		id := new(string)
		if !r.ClientCertificate.ID.IsUnknown() && !r.ClientCertificate.ID.IsNull() {
			*id = r.ClientCertificate.ID.ValueString()
		} else {
			id = nil
		}
		clientCertificate = &shared.ClientCertificate{
			ID: id,
		}
	}
	connectTimeout := new(int64)
	if !r.ConnectTimeout.IsUnknown() && !r.ConnectTimeout.IsNull() {
		*connectTimeout = r.ConnectTimeout.ValueInt64()
	} else {
		connectTimeout = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var host string
	host = r.Host.ValueString()

	id1 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueString()
	} else {
		id1 = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	path := new(string)
	if !r.Path.IsUnknown() && !r.Path.IsNull() {
		*path = r.Path.ValueString()
	} else {
		path = nil
	}
	var port int64
	port = r.Port.ValueInt64()

	protocol := shared.Protocol(r.Protocol.ValueString())
	readTimeout := new(int64)
	if !r.ReadTimeout.IsUnknown() && !r.ReadTimeout.IsNull() {
		*readTimeout = r.ReadTimeout.ValueInt64()
	} else {
		readTimeout = nil
	}
	retries := new(int64)
	if !r.Retries.IsUnknown() && !r.Retries.IsNull() {
		*retries = r.Retries.ValueInt64()
	} else {
		retries = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	tlsVerify := new(bool)
	if !r.TLSVerify.IsUnknown() && !r.TLSVerify.IsNull() {
		*tlsVerify = r.TLSVerify.ValueBool()
	} else {
		tlsVerify = nil
	}
	tlsVerifyDepth := new(int64)
	if !r.TLSVerifyDepth.IsUnknown() && !r.TLSVerifyDepth.IsNull() {
		*tlsVerifyDepth = r.TLSVerifyDepth.ValueInt64()
	} else {
		tlsVerifyDepth = nil
	}
	writeTimeout := new(int64)
	if !r.WriteTimeout.IsUnknown() && !r.WriteTimeout.IsNull() {
		*writeTimeout = r.WriteTimeout.ValueInt64()
	} else {
		writeTimeout = nil
	}
	out := shared.ServiceInput{
		CaCertificates:    caCertificates,
		ClientCertificate: clientCertificate,
		ConnectTimeout:    connectTimeout,
		Enabled:           enabled,
		Host:              host,
		ID:                id1,
		Name:              name,
		Path:              path,
		Port:              port,
		Protocol:          protocol,
		ReadTimeout:       readTimeout,
		Retries:           retries,
		Tags:              tags,
		TLSVerify:         tlsVerify,
		TLSVerifyDepth:    tlsVerifyDepth,
		WriteTimeout:      writeTimeout,
	}
	return &out
}

func (r *GatewayServiceResourceModel) RefreshFromSharedService(resp *shared.Service) {
	if resp != nil {
		r.CaCertificates = []types.String{}
		for _, v := range resp.CaCertificates {
			r.CaCertificates = append(r.CaCertificates, types.StringValue(v))
		}
		if resp.ClientCertificate == nil {
			r.ClientCertificate = nil
		} else {
			r.ClientCertificate = &tfTypes.ACLConsumer{}
			r.ClientCertificate.ID = types.StringPointerValue(resp.ClientCertificate.ID)
		}
		r.ConnectTimeout = types.Int64PointerValue(resp.ConnectTimeout)
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.Host = types.StringValue(resp.Host)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.Path = types.StringPointerValue(resp.Path)
		r.Port = types.Int64Value(resp.Port)
		r.Protocol = types.StringValue(string(resp.Protocol))
		r.ReadTimeout = types.Int64PointerValue(resp.ReadTimeout)
		r.Retries = types.Int64PointerValue(resp.Retries)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.TLSVerify = types.BoolPointerValue(resp.TLSVerify)
		r.TLSVerifyDepth = types.Int64PointerValue(resp.TLSVerifyDepth)
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
		r.WriteTimeout = types.Int64PointerValue(resp.WriteTimeout)
	}
}
