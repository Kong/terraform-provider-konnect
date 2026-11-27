package hooks

import "os"

/*
 * This file is only ever generated once on the first generation and then is free to be modified.
 * Any hooks you wish to add should be registered in the initHooks function. Feel free to define
 * your hooks in this file or in separate files in the hooks package.
 *
 * Hooks are registered per SDK instance, and are valid for the lifetime of the SDK instance.
 */

func initHooks(h *Hooks) {
	// This supports importing auth server clients with ID and client secret in input
	h.registerBeforeRequestHook(&AuthServerClientCreateHook{})
	// Remove fields covered by partials so the server uses the partial's values
	h.registerBeforeRequestHook(&PluginWithPartialHook{})
	// The API returns pending_verification, but does not accept it on updates.
	h.registerBeforeRequestHook(&OrgNotificationDestinationStateHook{})

	// Kong Identity is not served by the global endpoint, so send the default to the US region.
	// Registered before CustomizeKongDomainHook, which rewrites the konghq.com host.
	h.registerBeforeRequestHook(&KongIdentityDefaultRegionHook{})

	// Domain customization - enable usage with non-prod domains
	h.registerBeforeRequestHook(&CustomizeKongDomainHook{
		Enabled:           os.Getenv("KONG_CUSTOM_DOMAIN") != "",
		Domain:            os.Getenv("KONG_CUSTOM_DOMAIN"),
		ReplaceFullDomain: os.Getenv("KONG_CUSTOM_DOMAIN_REPLACE_FULL_DOMAIN") == "true",
	})

	// Konnect rejects min_runtime_version in updates of managed and serverless AI
	// Gateways; drop the value echoed from state. Registered after the domain
	// customization so the gateway is read from the final host.
	h.registerBeforeRequestHook(&AIGatewayMinRuntimeVersionHook{})

	// Debug hooks - dump request/response
	h.registerBeforeRequestHook(&HTTPDumpRequestHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_REQUEST") == "true",
	})
	h.registerAfterSuccessHook(&HTTPDumpResponseHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_RESPONSE") == "true",
	})
	h.registerAfterErrorHook(&HTTPDumpResponseErrorHook{
		Enabled: os.Getenv("KONNECT_SDK_HTTP_DUMP_RESPONSE_ERROR") == "true",
	})
}
