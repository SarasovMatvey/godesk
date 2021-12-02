package domain

type StatusResp struct {
	Status string `json:"status"`
}

type MetaParams struct {
	Limit  int
	Offset int
}

type MetaResp struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
