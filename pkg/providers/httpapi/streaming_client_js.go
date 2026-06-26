//go:build js

package httpapi

import "net/http"

func streamingClient(base *http.Client) *http.Client {
	return base
}
