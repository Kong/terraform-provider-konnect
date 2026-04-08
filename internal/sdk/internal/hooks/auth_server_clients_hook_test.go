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

func TestAuthServerClientCreateHook_BeforeRequest(t *testing.T) {
	testCases := []struct {
		name           string
		req            *http.Request
		expectedMethod string
		expectedPath   string
		expectedBody   string
		expectError    bool
	}{
		{
			name: "converts POST with id to PUT",
			req: &http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path: "/v1/auth-servers/server-123/clients",
				},
				Body: io.NopCloser(bytes.NewBufferString(`{"id":"client-456","name":"test-client"}`)),
			},
			expectedMethod: http.MethodPut,
			expectedPath:   "/v1/auth-servers/server-123/clients/client-456",
			expectedBody:   `{"id":"client-456","name":"test-client"}`,
			expectError:    false,
		},
		{
			name: "keeps POST when no id field",
			req: &http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path: "/v1/auth-servers/server-123/clients",
				},
				Body: io.NopCloser(bytes.NewBufferString(`{"name":"test-client"}`)),
			},
			expectedMethod: http.MethodPost,
			expectedPath:   "/v1/auth-servers/server-123/clients",
			expectedBody:   `{"name":"test-client"}`,
			expectError:    false,
		},
		{
			name: "keeps POST when id is empty string",
			req: &http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path: "/v1/auth-servers/server-123/clients",
				},
				Body: io.NopCloser(bytes.NewBufferString(`{"id":"","name":"test-client"}`)),
			},
			expectedMethod: http.MethodPost,
			expectedPath:   "/v1/auth-servers/server-123/clients",
			expectedBody:   `{"id":"","name":"test-client"}`,
			expectError:    false,
		},
		{
			name: "does not modify non-matching requests",
			req: &http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path: "/v1/other-endpoint",
				},
				Body: io.NopCloser(bytes.NewBufferString(`{"id":"test-123"}`)),
			},
			expectedMethod: http.MethodPost,
			expectedPath:   "/v1/other-endpoint",
			expectedBody:   `{"id":"test-123"}`,
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hook := AuthServerClientCreateHook{}
			result, err := hook.BeforeRequest(BeforeRequestContext{}, tc.req)

			if tc.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			assert.Equal(t, tc.expectedMethod, result.Method)
			assert.Equal(t, tc.expectedPath, result.URL.Path)

			bodyBytes, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expectedBody, string(bodyBytes))
		})
	}
}
