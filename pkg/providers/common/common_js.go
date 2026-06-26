//go:build js

package common

import (
	"net/http"
)

// NewHTTPClient creates an *http.Client for js/wasm builds.
// Custom transports and proxy settings are not supported in the browser environment.
// Returns a new client each time so callers can safely modify Timeout without
// affecting the global http.DefaultClient.
func NewHTTPClient(proxy string) *http.Client {
	return &http.Client{
		Timeout: DefaultRequestTimeout,
	}
}
