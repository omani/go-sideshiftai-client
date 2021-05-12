package sideshiftai

import "net/http"

// Config holds the configuration of a sideshiftai rpc client.
type Config struct {
	APIBaseAddress string
	APIVersion     string
	CustomHeaders  map[string]string
	Transport      http.RoundTripper
}
