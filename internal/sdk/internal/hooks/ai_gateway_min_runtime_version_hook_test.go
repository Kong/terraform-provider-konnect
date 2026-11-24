package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAIGatewayMinRuntimeVersionHook_BeforeRequest(t *testing.T) {
	// serve returns a server that answers GET /v1/ai-gateways/{id} with the
	// given gateway, and counts the reads.
	serve := func(t *testing.T, deploymentType, minRuntimeVersion string) (*httptest.Server, *int) {
		t.Helper()
		reads := 0
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodGet, r.Method)
			require.Equal(t, "Bearer token", r.Header.Get("Authorization"))
			reads++
			_ = json.NewEncoder(w).Encode(map[string]string{
				"id":                  "gw-1",
				"deployment_type":     deploymentType,
				"min_runtime_version": minRuntimeVersion,
			})
		}))
		t.Cleanup(srv.Close)
		return srv, &reads
	}

	makeReq := func(t *testing.T, method, url, body string) *http.Request {
		t.Helper()
		req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer token")
		return req
	}

	readBody := func(t *testing.T, r *http.Request) map[string]interface{} {
		t.Helper()
		b, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(b, &m))
		return m
	}

	body := `{"display_name":"x","name":"gw","description":"changed","min_runtime_version":"2.3","runtime_auto_upgrade":true}`

	for _, deploymentType := range []string{"managed", "serverless"} {
		t.Run("removes the echoed value for a "+deploymentType+" gateway", func(t *testing.T) {
			srv, reads := serve(t, deploymentType, "2.3")
			hook := AIGatewayMinRuntimeVersionHook{Client: srv.Client()}

			req, err := hook.BeforeRequest(BeforeRequestContext{}, makeReq(t, http.MethodPut, srv.URL+"/v1/ai-gateways/gw-1", body))
			require.NoError(t, err)

			got := readBody(t, req)
			assert.NotContains(t, got, "min_runtime_version")
			assert.Equal(t, "changed", got["description"])
			assert.Equal(t, true, got["runtime_auto_upgrade"])
			assert.Equal(t, 1, *reads)
		})
	}

	t.Run("keeps an explicit change on a managed gateway", func(t *testing.T) {
		srv, _ := serve(t, "managed", "2.2")
		hook := AIGatewayMinRuntimeVersionHook{Client: srv.Client()}

		req, err := hook.BeforeRequest(BeforeRequestContext{}, makeReq(t, http.MethodPut, srv.URL+"/v1/ai-gateways/gw-1", body))
		require.NoError(t, err)
		assert.Equal(t, "2.3", readBody(t, req)["min_runtime_version"])
	})

	t.Run("keeps the value for a hybrid gateway", func(t *testing.T) {
		srv, _ := serve(t, "hybrid", "2.3")
		hook := AIGatewayMinRuntimeVersionHook{Client: srv.Client()}

		req, err := hook.BeforeRequest(BeforeRequestContext{}, makeReq(t, http.MethodPut, srv.URL+"/v1/ai-gateways/gw-1", body))
		require.NoError(t, err)
		assert.Equal(t, "2.3", readBody(t, req)["min_runtime_version"])
	})

	t.Run("does not read the gateway when the field is absent", func(t *testing.T) {
		srv, reads := serve(t, "managed", "2.3")
		hook := AIGatewayMinRuntimeVersionHook{Client: srv.Client()}

		req, err := hook.BeforeRequest(BeforeRequestContext{}, makeReq(t, http.MethodPut, srv.URL+"/v1/ai-gateways/gw-1", `{"display_name":"x","name":"gw"}`))
		require.NoError(t, err)
		assert.NotContains(t, readBody(t, req), "min_runtime_version")
		assert.Equal(t, 0, *reads)
	})

	t.Run("ignores other paths and methods", func(t *testing.T) {
		srv, reads := serve(t, "managed", "2.3")
		hook := AIGatewayMinRuntimeVersionHook{Client: srv.Client()}

		for _, c := range []struct{ method, path string }{
			{http.MethodPost, "/v1/ai-gateways"},
			{http.MethodPut, "/v1/ai-gateways/gw-1/models/m-1"},
			{http.MethodPut, "/v1/event-gateways/gw-1"},
		} {
			req, err := hook.BeforeRequest(BeforeRequestContext{}, makeReq(t, c.method, srv.URL+c.path, body))
			require.NoError(t, err)
			assert.Equal(t, "2.3", readBody(t, req)["min_runtime_version"], c.method+" "+c.path)
		}
		assert.Equal(t, 0, *reads)
	})

	t.Run("sends the request unchanged when the gateway cannot be read", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
		}))
		t.Cleanup(srv.Close)
		hook := AIGatewayMinRuntimeVersionHook{Client: srv.Client()}

		req, err := hook.BeforeRequest(BeforeRequestContext{}, makeReq(t, http.MethodPut, srv.URL+"/v1/ai-gateways/gw-1", body))
		require.NoError(t, err)
		assert.Equal(t, "2.3", readBody(t, req)["min_runtime_version"])
	})
}
