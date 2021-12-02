package domain

type GetRequestStatsResp struct {
	Data []Request `json:"data"`
	Meta MetaResp  `json:"meta"`
}
