//go:build js

package openai_compat

import "net/http"

func streamingClient(base *http.Client) *http.Client {
	return base
}
