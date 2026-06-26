//go:build !js

package common

import (
	"log"
	"net/http"
	"net/url"
)

// NewHTTPClient creates an *http.Client with an optional proxy and the default timeout.
func NewHTTPClient(proxy string) *http.Client {
	client := &http.Client{
		Timeout: DefaultRequestTimeout,
	}
	if proxy != "" {
		parsed, err := url.Parse(proxy)
		if err == nil {
			// Preserve http.DefaultTransport settings (TLS, HTTP/2, timeouts, etc.)
			if base, ok := http.DefaultTransport.(*http.Transport); ok {
				tr := base.Clone()
				tr.Proxy = http.ProxyURL(parsed)
				client.Transport = tr
			} else {
				// Fallback: minimal transport if DefaultTransport is not *http.Transport.
				client.Transport = &http.Transport{
					Proxy: http.ProxyURL(parsed),
				}
			}
		} else {
			log.Printf("common: invalid proxy URL %q: %v", proxy, err)
		}
	}
	return client
}
