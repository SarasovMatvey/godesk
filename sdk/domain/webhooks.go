package domain

type AddWebhooksParams struct {
	Name   string   `json:"name"`
	Url    string   `json:"url"`
	Events []string `json:"events"`
}
