package hooks

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKongIdentityDefaultRegionHook_BeforeRequest(t *testing.T) {
	testCases := []struct {
		name         string
		host         string
		path         string
		expectedHost string
	}{
		{"global directories list goes to us", "global.api.konghq.com", "/v2/directories", "us.api.konghq.com"},
		{"global nested identity path goes to us", "global.api.konghq.com", "/v2/directories/abc/principals/def", "us.api.konghq.com"},
		{"eu is left alone", "eu.api.konghq.com", "/v2/directories/abc", "eu.api.konghq.com"},
		{"au is left alone", "au.api.konghq.com", "/v2/directories", "au.api.konghq.com"},
		{"us is left alone", "us.api.konghq.com", "/v2/directories", "us.api.konghq.com"},
		{"global non-identity path is left alone", "global.api.konghq.com", "/v3/identity-providers", "global.api.konghq.com"},
		{"global similar prefix is left alone", "global.api.konghq.com", "/v2/directories-other", "global.api.konghq.com"},
		{"custom host is left alone", "global.api.example.com", "/v2/directories", "global.api.example.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := &http.Request{
				Method: http.MethodGet,
				URL:    &url.URL{Scheme: "https", Host: tc.host, Path: tc.path},
				Host:   tc.host,
			}

			hook := &KongIdentityDefaultRegionHook{}
			got, err := hook.BeforeRequest(BeforeRequestContext{}, req)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedHost, got.URL.Host)
			assert.Equal(t, tc.expectedHost, got.Host)
			assert.Equal(t, tc.path, got.URL.Path)
		})
	}
}
