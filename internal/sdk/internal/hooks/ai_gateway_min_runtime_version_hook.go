package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"time"
)

// AIGatewayMinRuntimeVersionHook removes min_runtime_version from AI Gateway
// updates when Konnect would reject it.
//
// Konnect returns min_runtime_version for every AI Gateway, but rejects it in
// an update of a managed or serverless gateway ("min_runtime_version: must not
// be set for a managed or serverless deployment"). The attribute is computed,
// so the provider sends the value it read back from state on every update,
// and any change to a managed gateway fails.
//
// Before such an update, the hook reads the gateway. If it is managed or
// serverless and the request carries the value the server already has (an echo
// of state, not a change), the field is removed. A different value is still
// sent, so an explicit change gets Konnect's error instead of being dropped.
type AIGatewayMinRuntimeVersionHook struct {
	// Client is used to read the gateway; nil uses a client with a timeout.
	Client *http.Client
}

var aiGatewayPathRe = regexp.MustCompile(`^/v1/ai-gateways/[^/]+/?$`)

func matchAIGatewayUpdate(req *http.Request) bool {
	return req != nil && req.URL != nil && req.Body != nil &&
		(req.Method == http.MethodPut || req.Method == http.MethodPatch) &&
		aiGatewayPathRe.MatchString(req.URL.Path)
}

// BeforeRequest implements beforeRequestHook.
func (h AIGatewayMinRuntimeVersionHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if !matchAIGatewayUpdate(req) {
		return req, nil
	}

	bodyBytes, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		return req, err
	}
	restore := func() (*http.Request, error) {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	var body map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return restore()
	}
	sent, ok := body["min_runtime_version"].(string)
	if !ok {
		return restore()
	}

	current, err := h.readGateway(req)
	if err != nil || current == nil {
		// Could not read the gateway: send the request unchanged.
		return restore()
	}
	switch current.DeploymentType {
	case "managed", "serverless":
	default:
		return restore()
	}
	if current.MinRuntimeVersion != sent {
		return restore()
	}

	delete(body, "min_runtime_version")
	newBody, err := json.Marshal(body)
	if err != nil {
		return restore()
	}
	req.Body = io.NopCloser(bytes.NewReader(newBody))
	req.ContentLength = int64(len(newBody))
	return req, nil
}

type aiGatewayRuntimeFields struct {
	DeploymentType    string `json:"deployment_type"`
	MinRuntimeVersion string `json:"min_runtime_version"`
}

func (h AIGatewayMinRuntimeVersionHook) readGateway(req *http.Request) (*aiGatewayRuntimeFields, error) {
	client := h.Client
	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}

	get, err := http.NewRequestWithContext(req.Context(), http.MethodGet, req.URL.String(), nil)
	if err != nil {
		return nil, err
	}
	for _, name := range []string{"Authorization", "User-Agent"} {
		if v := req.Header.Get(name); v != "" {
			get.Header.Set(name, v)
		}
	}
	get.Header.Set("Accept", "application/json")

	res, err := client.Do(get)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, nil
	}

	var gw aiGatewayRuntimeFields
	if err := json.NewDecoder(res.Body).Decode(&gw); err != nil {
		return nil, err
	}
	return &gw, nil
}
