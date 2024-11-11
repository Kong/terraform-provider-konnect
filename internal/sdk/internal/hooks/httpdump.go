package hooks

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// HTTPDumpRequestHook is a hook that dumps the request to stdout.
type HTTPDumpRequestHook struct {
	Enabled bool
}

var _ beforeRequestHook = (*HTTPDumpRequestHook)(nil)

// BeforeRequest dumps the request to stdout if enabled.
func (i *HTTPDumpRequestHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if !i.Enabled {
		return req, nil
	}
	if req == nil {
		return nil, nil
	}

	fmt.Println("=================")
	b, err := httputil.DumpRequestOut(req, true)
	fmt.Println("=================")
	if err != nil {
		fmt.Printf("Error dumping request: %v\n", err)
	} else {
		fmt.Printf("request:\n%s\n\n", b)
	}

	return req, nil
}

// HTTPDumpResponseHook is a hook that dumps the response to stdout.
type HTTPDumpResponseHook struct {
	Enabled bool
}

var _ afterSuccessHook = (*HTTPDumpResponseHook)(nil)

// AfterSuccess dumps the response to stdout if enabled.
func (i *HTTPDumpResponseHook) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if !i.Enabled {
		return res, nil
	}
	return dumpResponse(res), nil
}

// HTTPDumpResponseErrorHook is a hook that dumps the error response to stdout.
type HTTPDumpResponseErrorHook struct {
	Enabled bool
}

var _ afterErrorHook = (*HTTPDumpResponseErrorHook)(nil)

// AfterError dumps the error response to stdout if enabled.
func (i *HTTPDumpResponseErrorHook) AfterError(hookCtx AfterErrorContext, res *http.Response, err error) (*http.Response, error) {
	if !i.Enabled {
		return res, nil
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return dumpResponse(res), err
}

func dumpResponse(res *http.Response) *http.Response {
	if res == nil {
		return nil
	}

	fmt.Println("-----------------")
	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Printf("Error dumping respone: %v\n", err)
	} else {
		fmt.Printf("response:\n%s\n\n", b)
	}
	fmt.Println("-----------------")

	return res
}
