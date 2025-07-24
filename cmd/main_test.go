package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pluralsh/pr-governance-webhook/api"
	"github.com/pluralsh/pr-governance-webhook/pkg/client"
	"github.com/pluralsh/pr-governance-webhook/pkg/handler"
)

func startTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/open", handler.OpenHandler)
	mux.HandleFunc("/v1/confirm", handler.ConfirmHandler)
	mux.HandleFunc("/v1/close", handler.CloseHandler)

	return httptest.NewServer(mux)
}

func TestEndpointsWithRunningServer(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	var pr = &api.PullRequest{
		Url:   "https://example.com",
		Title: "Test PR",
		Body:  "Just testing",
		Ref:   "test-branch",
	}

	c := client.New(server.URL)
	resp, err := c.Open(pr)

	assert.NoError(t, err)
	assert.Equal(t, pr.Url, resp["url"])
	assert.Equal(t, pr.Ref, resp["ref"])

	err = c.Close(pr, map[string]any{})
	assert.NoError(t, err)

	err = c.Confirm(pr, map[string]any{})
	assert.NoError(t, err)

}
