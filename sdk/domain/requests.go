package domain

type Request struct {
	RequestId         int         `json:"request_id"`
	FirstOperatorId   interface{} `json:"first_operator_id"`
	FirstOperatorName string      `json:"first_operator_name"`
	RequestTime       int         `json:"request_time"`
	RequestStart      interface{} `json:"request_start"`
	RequestStartD     string      `json:"request_start_d"`
	RequestStartT     string      `json:"request_start_t"`
	ReplyStart        interface{} `json:"reply_start"`
	ReplyStartD       string      `json:"reply_start_d"`
	ReplyStartT       string      `json:"reply_start_t"`
	ClientMessenger   string      `json:"client_messenger"`
	ClientId          int         `json:"client_id"`
	ClientName        string      `json:"client_name"`
}
