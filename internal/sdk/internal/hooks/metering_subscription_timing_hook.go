package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

// MeteringSubscriptionTimingHook removes the `timing` key from POST request
// bodies for metering subscription create and change endpoints.
type MeteringSubscriptionTimingHook struct{}

var konnectMatchMeteringSubscriptionPOST = func(req *http.Request) bool {
	if req == nil || req.URL == nil {
		return false
	}

	match, err := regexp.MatchString(`^/v3/openmeter/subscriptions$`, req.URL.Path)
	if err != nil {
		return false
	}

	return match && req.Method == http.MethodPost
}

// BeforeRequest removes the `timing` key from the request body.
func (h MeteringSubscriptionTimingHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if !konnectMatchMeteringSubscriptionPOST(req) {
		return req, nil
	}

	if req.Body == nil {
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
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	if _, exists := bodyMap["timing"]; !exists {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	delete(bodyMap, "timing")

	newBody, err := json.Marshal(bodyMap)
	if err != nil {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	req.Body = io.NopCloser(bytes.NewReader(newBody))
	req.ContentLength = int64(len(newBody))
	return req, nil
}
