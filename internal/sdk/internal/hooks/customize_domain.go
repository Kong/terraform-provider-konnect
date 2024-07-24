package hooks

import (
	"net/http"
	"os"
	"strings"
)

type CustomizeKongDomainHook struct{}

// Replace `.konghq.com` with the custom domain set in the `KONG_CUSTOM_DOMAIN` environment variable.
func (i *CustomizeKongDomainHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	customDomain := os.Getenv("KONG_CUSTOM_DOMAIN")
	if customDomain != "" {
		req.URL.Host = strings.Replace(req.URL.Host, ".konghq.com", "."+customDomain, 1)
	}

	return req, nil
}
