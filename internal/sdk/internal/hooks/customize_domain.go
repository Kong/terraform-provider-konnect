package hooks

import (
	"net/http"
	"strings"
)

type CustomizeKongDomainHook struct {
	Enabled           bool
	Domain            string
	ReplaceFullDomain bool
}

// Replace `.konghq.com` with the custom domain set in the `KONG_CUSTOM_DOMAIN` environment variable.
func (i *CustomizeKongDomainHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if !i.Enabled {
		return req, nil
	}

	customDomain := i.Domain
	var host string

	if i.ReplaceFullDomain {
		host = i.Domain
	} else {
		host = strings.Replace(req.URL.Host, ".konghq.com", "."+customDomain, 1)
	}

	req.URL.Host = host
	req.Host = host

	return req, nil
}
