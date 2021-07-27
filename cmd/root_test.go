package cmd

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"git.sr.ht/~kota/metservice-go"
)

// setup sets up a test HTTP server along with a Client that is configured to
// talk to that test server. Tests should register handlers on mux which
// provide mock responses for the API method being tested.
func setup() (client *metservice.Client, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the Client being tested and is configured to use test server.
	localURL, _ := url.Parse(server.URL + "/")
	localClient := &metservice.Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    localURL.String(),
	}

	return localClient, mux, server.Close
}
