package hooks

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
)

// FeaturesHook is a struct that implements the BeforeRequestHook interface.
type FeaturesHook struct{}

// BeforeRequest modifies the request before sending it.
func (e FeaturesHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
    if req.Method == http.MethodPost && req.URL.Path == "/v1/mesh/control-planes" {
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
            }
        }

        if bodyMap == nil {
            bodyMap = make(map[string]interface{})
        }

        // Define the default "features" value
        defaultFeatures := []map[string]interface{}{
            {
                "type":         "MeshCreation",
                "meshCreation": map[string]bool{"enabled": false},
            },
            {
                "type":                      "HostnameGeneratorCreation",
                "hostnameGeneratorCreation": map[string]bool{"enabled": false},
            },
        }

        if _, exists := bodyMap["features"]; !exists {
            bodyMap["features"] = defaultFeatures
        }

        newBody, err := json.Marshal(bodyMap)
        if err != nil {
            return nil, err
        }

        req.Body = io.NopCloser(bytes.NewReader(newBody))
        req.ContentLength = int64(len(newBody))
        req.Header.Set("Content-Type", "application/json")
    }

    return req, nil
}
