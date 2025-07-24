package api

type PullRequest struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Ref   string `json:"ref"`
}

type OpenInput struct {
	Pr *PullRequest `json:"pr"`
}

type ConfirmInput struct {
	Pr    *PullRequest   `json:"pr"`
	State map[string]any `json:"state"`
}

type CloseInput struct {
	Pr    *PullRequest   `json:"pr"`
	State map[string]any `json:"state"`
}
