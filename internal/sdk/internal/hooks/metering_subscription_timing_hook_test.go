package hooks

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMeteringSubscriptionTimingHook_BeforeRequest(t *testing.T) {
	hook := MeteringSubscriptionTimingHook{}
	ctx := BeforeRequestContext{}

	makeReq := func(method, path, body string) *http.Request {
		return &http.Request{
			Method: method,
			URL:    &url.URL{Path: path},
			Body:   io.NopCloser(bytes.NewBufferString(body)),
		}
	}

	readBody := func(t *testing.T, r *http.Request) string {
		t.Helper()
		b, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		return string(b)
	}

	t.Run("removes timing from create subscription POST", func(t *testing.T) {
		req := makeReq(http.MethodPost, "/v3/openmeter/subscriptions", `{
			"customer": {"id": "cust-123"},
			"plan": {"id": "plan-456"},
			"timing": "immediate"
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"customer": {"id": "cust-123"},
			"plan": {"id": "plan-456"}
		}`, readBody(t, result))
	})

	t.Run("does not remove timing from change subscription POST", func(t *testing.T) {
		req := makeReq(http.MethodPost, "/v3/openmeter/subscriptions/sub-789/change", `{
			"customer": {"id": "cust-123"},
			"plan": {"id": "plan-456"},
			"timing": "next_billing_cycle"
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"customer": {"id": "cust-123"},
			"plan": {"id": "plan-456"},
			"timing": "next_billing_cycle"
		}`, readBody(t, result))
	})

	t.Run("passthrough when no timing key", func(t *testing.T) {
		body := `{"customer": {"id": "cust-123"}, "plan": {"id": "plan-456"}}`
		req := makeReq(http.MethodPost, "/v3/openmeter/subscriptions", body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, body, readBody(t, result))
	})

	t.Run("does not modify GET requests", func(t *testing.T) {
		body := `{"customer": {"id": "cust-123"}, "timing": "immediate"}`
		req := makeReq(http.MethodGet, "/v3/openmeter/subscriptions", body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, body, readBody(t, result))
	})

	t.Run("does not modify PUT requests", func(t *testing.T) {
		body := `{"customer": {"id": "cust-123"}, "timing": "immediate"}`
		req := makeReq(http.MethodPut, "/v3/openmeter/subscriptions", body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, body, readBody(t, result))
	})

	t.Run("does not modify DELETE requests", func(t *testing.T) {
		body := `{"customer": {"id": "cust-123"}, "timing": "immediate"}`
		req := makeReq(http.MethodDelete, "/v3/openmeter/subscriptions", body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, body, readBody(t, result))
	})

	t.Run("skips non-matching paths", func(t *testing.T) {
		body := `{"customer": {"id": "cust-123"}, "timing": "immediate"}`
		req := makeReq(http.MethodPost, "/v1/other-endpoint", body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, body, readBody(t, result))
	})

	t.Run("handles empty body", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodPost,
			URL:    &url.URL{Path: "/v3/openmeter/subscriptions"},
			Body:   io.NopCloser(bytes.NewBufferString("")),
		}
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, "", readBody(t, result))
	})

	t.Run("handles nil body", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodPost,
			URL:    &url.URL{Path: "/v3/openmeter/subscriptions"},
			Body:   nil,
		}
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Nil(t, result.Body)
	})

	t.Run("handles invalid JSON gracefully", func(t *testing.T) {
		req := makeReq(http.MethodPost, "/v3/openmeter/subscriptions", `not valid json`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, `not valid json`, readBody(t, result))
	})

	t.Run("updates content length after modification", func(t *testing.T) {
		req := makeReq(http.MethodPost, "/v3/openmeter/subscriptions", `{
			"customer": {"id": "cust-123"},
			"plan": {"id": "plan-456"},
			"timing": "immediate"
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		body := readBody(t, result)
		assert.Equal(t, int64(len(body)), result.ContentLength)
	})

	t.Run("preserves other fields when removing timing", func(t *testing.T) {
		req := makeReq(http.MethodPost, "/v3/openmeter/subscriptions", `{
			"customer": {"id": "cust-123"},
			"plan": {"id": "plan-456"},
			"billing_anchor": "2025-01-01T00:00:00Z",
			"labels": {"env": "test"},
			"timing": "immediate"
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"customer": {"id": "cust-123"},
			"plan": {"id": "plan-456"},
			"billing_anchor": "2025-01-01T00:00:00Z",
			"labels": {"env": "test"}
		}`, readBody(t, result))
	})
}
