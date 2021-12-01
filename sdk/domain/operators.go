package domain

type Operator struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Online     int    `json:"online"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ExternalId string `json:"external_id"`
	StatusId   int    `json:"status_id"`
}

type RespMeta struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetOperatorsResp struct {
	Data   []Operator `json:"data"`
	Meta   RespMeta   `json:"meta"`
	Status string     `json:"status"`
}
