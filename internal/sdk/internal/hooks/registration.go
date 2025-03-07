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
    h.registerBeforeRequestHook(&MeshDefaultsHook{})

    // Domain customization - enable usage with non-prod domains
    h.registerBeforeRequestHook(&CustomizeKongDomainHook{})

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
