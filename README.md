# pr-governance-webhook

This project is a simple HTTP server in Go that implements three JSON-based endpoints:

- `POST /v1/open`
- `POST /v1/confirm`
- `POST /v1/close`

Each endpoint accepts and returns JSON payloads.


## Project Structure

- `cmd/main.go` — Server implementation and handlers.
- `cmd/main_test.go` — Tests that run the server and hit endpoints via HTTP.
- `pkg/client/client.go` — Simple Go client to call the endpoints.
- `Makefile` — Build, run, test commands.

## API Types

```go
type PullRequest struct {
  Url   string
  Title string
  Body  string
  Ref   string
}

type OpenInput struct {
  Pr *PullRequest
}

type ConfirmInput struct {
  Pr    *PullRequest
  State map[string]any
}

type CloseInput struct {
  Pr    *PullRequest
  State map[string]any
}
```
## Go client

```go
import "github.com/pluralsh/pr-governance-webhook/pkg/client"

var pr = &api.PullRequest{
    Url:   "https://example.com",
    Title: "Test PR",
    Body:  "Just testing",
    Ref:   "test-branch",
}
c := client.New(server.URL)
resp, err := c.Open(pr) // (map[string]any, error), echo response
err = c.Close(pr, map[string]any{}) // err
err = c.Confirm(pr, map[string]any{}) // err


```