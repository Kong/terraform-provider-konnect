package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrgNotificationDestinationStateHook(t *testing.T) {
	makeRequest := func(method, url, body string) *http.Request {
		req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
		require.NoError(t, err)
		return req
	}
	readBody := func(req *http.Request) map[string]any {
		data, err := io.ReadAll(req.Body)
		require.NoError(t, err)
		var got map[string]any
		require.NoError(t, json.Unmarshal(data, &got))
		return got
	}
	hook := OrgNotificationDestinationStateHook{}

	t.Run("removes pending verification state from destination patch", func(t *testing.T) {
		req := makeRequest(http.MethodPatch, "https://global.api.konghq.com/v1/notifications/destinations/dest-123", `{"name":"updated","state":"pending_verification"}`)
		got, err := hook.BeforeRequest(BeforeRequestContext{}, req)
		require.NoError(t, err)
		assert.Equal(t, map[string]any{"name": "updated"}, readBody(got))
		assert.Equal(t, got.ContentLength, req.ContentLength)
	})

	t.Run("keeps other destination state values", func(t *testing.T) {
		req := makeRequest(http.MethodPatch, "https://global.api.konghq.com/v1/notifications/destinations/dest-123", `{"state":"enabled"}`)
		got, err := hook.BeforeRequest(BeforeRequestContext{}, req)
		require.NoError(t, err)
		assert.Equal(t, map[string]any{"state": "enabled"}, readBody(got))
	})

	t.Run("ignores other endpoints and methods", func(t *testing.T) {
		for _, tc := range []struct {
			method string
			url    string
		}{
			{http.MethodPatch, "https://global.api.konghq.com/v1/notifications/destinations"},
			{http.MethodPatch, "https://global.api.konghq.com/v1/notifications/destinations/dest-123/test"},
			{http.MethodPut, "https://global.api.konghq.com/v1/notifications/destinations/dest-123"},
		} {
			req := makeRequest(tc.method, tc.url, `{"state":"pending_verification"}`)
			got, err := hook.BeforeRequest(BeforeRequestContext{}, req)
			require.NoError(t, err)
			assert.Equal(t, map[string]any{"state": "pending_verification"}, readBody(got))
		}
	})
}
