package testutil

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

// NewTestServer is a http test server stub instance.
// Used for stub request controlling match path, response body and status code.
func NewTestServer(path string, body string, statusCode int) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, path) {
				w.WriteHeader(statusCode)
				w.Write([]byte(body))
			}
		}),
	)
}
