package hooks

import (
	"net/http"
	"strings"
)

const (
	globalKonnectHost  = "global.api.konghq.com"
	defaultRegionHost  = "us.api.konghq.com"
	kongIdentityPrefix = "/v2/directories"
)

// KongIdentityDefaultRegionHook sends Kong Identity requests that target the global
// endpoint to the US region instead. The global endpoint does not serve Kong Identity,
// and it is the provider default when server_url is not set. A request to any regional
// endpoint is left as it is, so server_url selects the region.
type KongIdentityDefaultRegionHook struct{}

func (h *KongIdentityDefaultRegionHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if req.URL == nil || req.URL.Host != globalKonnectHost {
		return req, nil
	}
	if req.URL.Path != kongIdentityPrefix && !strings.HasPrefix(req.URL.Path, kongIdentityPrefix+"/") {
		return req, nil
	}

	req.URL.Host = defaultRegionHost
	req.Host = defaultRegionHost

	return req, nil
}
