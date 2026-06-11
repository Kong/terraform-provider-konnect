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

func TestPluginWithPartialHook_BeforeRequest(t *testing.T) {
	hook := PluginWithPartialHook{}
	ctx := BeforeRequestContext{}

	makeReq := func(method, body string) *http.Request {
		return &http.Request{
			Method: method,
			URL:    &url.URL{Path: "/v2/control-planes/cp-1/core-entities/plugins"},
			Body:   io.NopCloser(bytes.NewBufferString(body)),
		}
	}

	readBody := func(t *testing.T, r *http.Request) string {
		t.Helper()
		b, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		return string(b)
	}

	t.Run("removes top-level path covered by partial", func(t *testing.T) {
		req := makeReq(http.MethodPost, `{
			"name": "rate-limiting-advanced",
			"config": {
				"limit": [100],
				"redis": {"host": "localhost", "port": 6379}
			},
			"partials": [{"id": "abc", "path": "config.redis"}]
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"name": "rate-limiting-advanced",
			"config": {"limit": [100]},
			"partials": [{"id": "abc", "path": "config.redis"}]
		}`, readBody(t, result))
	})

	t.Run("removes multiple paths from multiple partials", func(t *testing.T) {
		req := makeReq(http.MethodPost, `{
			"config": {
				"vectordb": {"strategy": "pgvector"},
				"embeddings": {"model": {"provider": "openai"}},
				"cache_ttl": 300
			},
			"partials": [
				{"id": "p1", "path": "config.vectordb"},
				{"id": "p2", "path": "config.embeddings"}
			]
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"config": {"cache_ttl": 300},
			"partials": [
				{"id": "p1", "path": "config.vectordb"},
				{"id": "p2", "path": "config.embeddings"}
			]
		}`, readBody(t, result))
	})

	t.Run("works on PUT requests", func(t *testing.T) {
		req := makeReq(http.MethodPut, `{
			"config": {"redis": {"host": "redis.example.com"}, "limit": [10]},
			"partials": [{"id": "r1", "path": "config.redis"}]
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"config": {"limit": [10]},
			"partials": [{"id": "r1", "path": "config.redis"}]
		}`, readBody(t, result))
	})

	t.Run("does not modify GET requests", func(t *testing.T) {
		body := `{"config":{"redis":{"host":"localhost"}},"partials":[{"id":"x","path":"config.redis"}]}`
		req := makeReq(http.MethodGet, body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, body, readBody(t, result))
	})

	t.Run("does not modify DELETE requests", func(t *testing.T) {
		body := `{"config":{"redis":{"host":"localhost"}},"partials":[{"id":"x","path":"config.redis"}]}`
		req := makeReq(http.MethodDelete, body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, body, readBody(t, result))
	})

	t.Run("passthrough when no partials key", func(t *testing.T) {
		body := `{"config":{"redis":{"host":"localhost"}}}`
		req := makeReq(http.MethodPost, body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, body, readBody(t, result))
	})

	t.Run("passthrough when partials array is empty", func(t *testing.T) {
		body := `{"config":{"redis":{"host":"localhost"}},"partials":[]}`
		req := makeReq(http.MethodPost, body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, body, readBody(t, result))
	})

	t.Run("skips partial with no path", func(t *testing.T) {
		body := `{"config":{"redis":{"host":"localhost"}},"partials":[{"id":"x"}]}`
		req := makeReq(http.MethodPost, body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, body, readBody(t, result))
	})

	t.Run("passthrough when path does not exist in body", func(t *testing.T) {
		body := `{"config":{"limit":[10]},"partials":[{"id":"x","path":"config.redis"}]}`
		req := makeReq(http.MethodPost, body)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.JSONEq(t, body, readBody(t, result))
	})

	t.Run("passthrough when body is empty", func(t *testing.T) {
		req := &http.Request{
			Method: http.MethodPost,
			URL:    &url.URL{Path: "/v2/plugins"},
			Body:   io.NopCloser(bytes.NewBufferString("")),
		}
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, "", readBody(t, result))
	})

	t.Run("skips body read for POST to non-matching path", func(t *testing.T) {
		body := `{"config":{"redis":{"host":"localhost"}},"partials":[{"id":"x","path":"config.redis"}]}`
		req := &http.Request{
			Method: http.MethodPost,
			URL:    &url.URL{Path: "/v1/control-planes/cp-1/core-entities/plugins"},
			Body:   io.NopCloser(bytes.NewBufferString(body)),
		}
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		// Body must be untouched — config.redis should still be present.
		assert.JSONEq(t, body, readBody(t, result))
	})

	t.Run("content length is updated after body modification", func(t *testing.T) {
		req := makeReq(http.MethodPost, `{
			"config": {"redis": {"host": "localhost"}, "limit": [10]},
			"partials": [{"id": "r1", "path": "config.redis"}]
		}`)
		result, err := hook.BeforeRequest(ctx, req)
		require.NoError(t, err)
		body := readBody(t, result)
		assert.Equal(t, int64(len(body)), result.ContentLength)
	})
}
