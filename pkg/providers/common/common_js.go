//go:build js

package common

import (
	"net/http"
)

// NewHTTPClient returns http.DefaultClient for js/wasm builds.
// In js/wasm, only http.DefaultClient is wired to the browser fetch API.
// Any other *http.Client will fall back to Go's DNS resolver which does not work.
// Proxy settings are ignored — the browser handles networking.
func NewHTTPClient(proxy string) *http.Client {
	return http.DefaultClient
}
