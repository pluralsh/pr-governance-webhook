//go:build e2e

package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestOpenEndpoint(t *testing.T) {
	payload := map[string]any{
		"pr": map[string]any{
			"url":   "https://example.com",
			"title": "E2E Test PR",
			"body":  "Some body",
			"ref":   "feature/e2e",
		},
	}
	body, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:8080/v1/open", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to call /v1/open: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}
}
