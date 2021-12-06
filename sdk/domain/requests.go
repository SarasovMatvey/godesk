package domain

type Request struct {
	RequestId         int         `json:"request_id"`
	FirstOperatorId   interface{} `json:"first_operator_id"`
	FirstOperatorName string      `json:"first_operator_name"`
	RequestTime       int         `json:"request_time"`
	ReplyStart        interface{} `json:"reply_start"`
	ClientMessenger   string      `json:"client_messenger"`
	ClientId          int         `json:"client_id"`
	ClientName        string      `json:"client_name"`
}
