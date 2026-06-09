package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// PluginWithPartialHook removes request body fields that are managed by partials.
// When a request body contains a `partials` array with `path` entries, the
// property at each path (dot-notation) is deleted from the request body so that
// those fields are not overwritten — they are owned by the referenced partial.
type PluginWithPartialHook struct{}

// BeforeRequest removes any fields from the request body that are covered by
// a partial's path.
func (h PluginWithPartialHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if req.Body == nil {
		return req, nil
	}

	if req.Method != http.MethodPost && req.Method != http.MethodPut {
		return req, nil
	}

	bodyBytes, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		return req, err
	}
	if len(bodyBytes) == 0 {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	var bodyMap map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
		// Not JSON — restore body unchanged.
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	partials, ok := bodyMap["partials"]
	if !ok {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	partialsList, ok := partials.([]interface{})
	if !ok {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	modified := false
	for _, p := range partialsList {
		partial, ok := p.(map[string]interface{})
		if !ok {
			continue
		}
		path, ok := partial["path"].(string)
		if !ok || path == "" {
			continue
		}
		if deleteNestedKey(bodyMap, strings.Split(path, ".")) {
			modified = true
		}
	}

	if !modified {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	newBody, err := json.Marshal(bodyMap)
	if err != nil {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	req.Body = io.NopCloser(bytes.NewReader(newBody))
	req.ContentLength = int64(len(newBody))
	return req, nil
}

// deleteNestedKey deletes the leaf key at the given dot-notation path segments
// from m. Returns true if a key was actually removed.
func deleteNestedKey(m map[string]interface{}, keys []string) bool {
	if len(keys) == 0 {
		return false
	}
	if len(keys) == 1 {
		if _, exists := m[keys[0]]; exists {
			delete(m, keys[0])
			return true
		}
		return false
	}
	next, ok := m[keys[0]].(map[string]interface{})
	if !ok {
		return false
	}
	return deleteNestedKey(next, keys[1:])
}
