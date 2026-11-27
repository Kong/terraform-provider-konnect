package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

var orgNotificationDestinationPath = regexp.MustCompile(`^/v1/notifications/destinations/[^/]+/?$`)

// OrgNotificationDestinationStateHook removes pending_verification from
// destination update requests. That state is returned by reads, but it is not
// accepted as an update value.
type OrgNotificationDestinationStateHook struct{}

func (OrgNotificationDestinationStateHook) BeforeRequest(_ BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if req == nil || req.URL == nil || req.Body == nil || req.Method != http.MethodPatch || !orgNotificationDestinationPath.MatchString(req.URL.Path) {
		return req, nil
	}

	bodyBytes, err := io.ReadAll(req.Body)
	_ = req.Body.Close()
	if err != nil {
		return req, err
	}
	restore := func() (*http.Request, error) {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		return req, nil
	}

	var body map[string]json.RawMessage
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return restore()
	}
	var state string
	if err := json.Unmarshal(body["state"], &state); err != nil || state != "pending_verification" {
		return restore()
	}

	delete(body, "state")
	updatedBody, err := json.Marshal(body)
	if err != nil {
		return restore()
	}
	req.Body = io.NopCloser(bytes.NewReader(updatedBody))
	req.ContentLength = int64(len(updatedBody))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(updatedBody)), nil
	}
	return req, nil
}
