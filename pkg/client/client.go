package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pluralsh/pr-governance-webhook/api"
)

type client struct {
	url        string
	httpClient *http.Client
}

type Client interface {
	Open(pr *api.PullRequest) (map[string]any, error)
	Confirm(pr *api.PullRequest, state map[string]any) error
	Close(pr *api.PullRequest, state map[string]any) error
}

func New(url string) Client {
	return &client{
		url:        url,
		httpClient: &http.Client{Timeout: 2 * time.Second},
	}
}

func (c *client) doPost(path string, data any, result any) error {
	url := c.url + path
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("http post: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close body")
		}
	}(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("bad response: %s", resp.Status)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

func (c *client) Open(pr *api.PullRequest) (map[string]any, error) {
	input := api.OpenInput{Pr: pr}
	var state map[string]any
	err := c.doPost("/v1/open", input, &state)
	return state, err
}

func (c *client) Confirm(pr *api.PullRequest, state map[string]any) error {
	input := api.ConfirmInput{Pr: pr, State: state}
	return c.doPost("/v1/confirm", input, nil)
}

func (c *client) Close(pr *api.PullRequest, state map[string]any) error {
	input := api.CloseInput{Pr: pr, State: state}
	return c.doPost("/v1/close", input, nil)
}
