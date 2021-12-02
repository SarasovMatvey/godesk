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

type GetOperatorsResp struct {
	Data   []Operator `json:"data"`
	Meta   MetaResp   `json:"meta"`
	Status string     `json:"status"`
}
