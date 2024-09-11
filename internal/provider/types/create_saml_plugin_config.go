// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateSamlPluginConfig struct {
	Anonymous                      types.String                             `tfsdk:"anonymous"`
	AssertionConsumerPath          types.String                             `tfsdk:"assertion_consumer_path"`
	IdpCertificate                 types.String                             `tfsdk:"idp_certificate"`
	IdpSsoURL                      types.String                             `tfsdk:"idp_sso_url"`
	Issuer                         types.String                             `tfsdk:"issuer"`
	NameidFormat                   types.String                             `tfsdk:"nameid_format"`
	Redis                          *CreateKonnectApplicationAuthPluginRedis `tfsdk:"redis"`
	RequestDigestAlgorithm         types.String                             `tfsdk:"request_digest_algorithm"`
	RequestSignatureAlgorithm      types.String                             `tfsdk:"request_signature_algorithm"`
	RequestSigningCertificate      types.String                             `tfsdk:"request_signing_certificate"`
	RequestSigningKey              types.String                             `tfsdk:"request_signing_key"`
	ResponseDigestAlgorithm        types.String                             `tfsdk:"response_digest_algorithm"`
	ResponseEncryptionKey          types.String                             `tfsdk:"response_encryption_key"`
	ResponseSignatureAlgorithm     types.String                             `tfsdk:"response_signature_algorithm"`
	SessionAbsoluteTimeout         types.Number                             `tfsdk:"session_absolute_timeout"`
	SessionAudience                types.String                             `tfsdk:"session_audience"`
	SessionCookieDomain            types.String                             `tfsdk:"session_cookie_domain"`
	SessionCookieHTTPOnly          types.Bool                               `tfsdk:"session_cookie_http_only"`
	SessionCookieName              types.String                             `tfsdk:"session_cookie_name"`
	SessionCookiePath              types.String                             `tfsdk:"session_cookie_path"`
	SessionCookieSameSite          types.String                             `tfsdk:"session_cookie_same_site"`
	SessionCookieSecure            types.Bool                               `tfsdk:"session_cookie_secure"`
	SessionEnforceSameSubject      types.Bool                               `tfsdk:"session_enforce_same_subject"`
	SessionHashStorageKey          types.Bool                               `tfsdk:"session_hash_storage_key"`
	SessionHashSubject             types.Bool                               `tfsdk:"session_hash_subject"`
	SessionIdlingTimeout           types.Number                             `tfsdk:"session_idling_timeout"`
	SessionMemcachedHost           types.String                             `tfsdk:"session_memcached_host"`
	SessionMemcachedPort           types.Int64                              `tfsdk:"session_memcached_port"`
	SessionMemcachedPrefix         types.String                             `tfsdk:"session_memcached_prefix"`
	SessionMemcachedSocket         types.String                             `tfsdk:"session_memcached_socket"`
	SessionRemember                types.Bool                               `tfsdk:"session_remember"`
	SessionRememberAbsoluteTimeout types.Number                             `tfsdk:"session_remember_absolute_timeout"`
	SessionRememberCookieName      types.String                             `tfsdk:"session_remember_cookie_name"`
	SessionRememberRollingTimeout  types.Number                             `tfsdk:"session_remember_rolling_timeout"`
	SessionRequestHeaders          []types.String                           `tfsdk:"session_request_headers"`
	SessionResponseHeaders         []types.String                           `tfsdk:"session_response_headers"`
	SessionRollingTimeout          types.Number                             `tfsdk:"session_rolling_timeout"`
	SessionSecret                  types.String                             `tfsdk:"session_secret"`
	SessionStorage                 types.String                             `tfsdk:"session_storage"`
	SessionStoreMetadata           types.Bool                               `tfsdk:"session_store_metadata"`
	ValidateAssertionSignature     types.Bool                               `tfsdk:"validate_assertion_signature"`
}
