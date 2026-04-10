package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

// struct that implements the BeforeRequestHook interface.
type AuthServerClientCreateHook struct{}

var konnectMatchAuthServerClientCreate = func(req *http.Request) bool {
	if req == nil || req.URL == nil {
		return false
	}

	match, err := regexp.MatchString(`^/v1/auth-servers/[^/]+/clients$`, req.URL.Path)
	if err != nil {
		return false
	}

	return match && req.Method == http.MethodPost
}

// BeforeRequest modifies the request before sending it.
func (e AuthServerClientCreateHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if konnectMatchAuthServerClientCreate(req) {
		// Change to PUT if creating with ID
		var bodyMap map[string]interface{}

		if req.Body != nil {
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
			defer req.Body.Close()

			if len(bodyBytes) > 0 {
				if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
					return nil, err
				}

				// Check if "id" field exists, if yes, use PUT to preserve the ID
				id, exists := bodyMap["id"]
				if idStr, ok := id.(string); ok && exists && idStr != "" {
					req.URL.Path = req.URL.Path + "/" + idStr
					req.Method = http.MethodPut
				}

				newBody, err := json.Marshal(bodyMap)
				if err != nil {
					return nil, err
				}

				req.Body = io.NopCloser(bytes.NewReader(newBody))
				req.ContentLength = int64(len(newBody))
			}
		}
	}

	return req, nil
}
