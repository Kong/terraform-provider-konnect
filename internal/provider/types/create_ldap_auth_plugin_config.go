// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateLdapAuthPluginConfig struct {
	Anonymous       types.String `tfsdk:"anonymous"`
	Attribute       types.String `tfsdk:"attribute"`
	BaseDn          types.String `tfsdk:"base_dn"`
	CacheTTL        types.Number `tfsdk:"cache_ttl"`
	HeaderType      types.String `tfsdk:"header_type"`
	HideCredentials types.Bool   `tfsdk:"hide_credentials"`
	Keepalive       types.Number `tfsdk:"keepalive"`
	LdapHost        types.String `tfsdk:"ldap_host"`
	LdapPort        types.Int64  `tfsdk:"ldap_port"`
	Ldaps           types.Bool   `tfsdk:"ldaps"`
	Realm           types.String `tfsdk:"realm"`
	StartTLS        types.Bool   `tfsdk:"start_tls"`
	Timeout         types.Number `tfsdk:"timeout"`
	VerifyLdapHost  types.Bool   `tfsdk:"verify_ldap_host"`
}
