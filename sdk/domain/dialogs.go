package domain

type LastMessage struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	Coordinates string `json:"coordinates"`
	Transport   string `json:"transport"`
	Type        string `json:"type"`
	Read        int    `json:"read"`
	Created     string `json:"created"`
	OperatorId  int    `json:"operator_id"`
	ClientId    int    `json:"client_id"`
}

type Dialog struct {
	Id          int         `json:"id"`
	State       string      `json:"state"`
	Begin       string      `json:"begin"`
	End         string      `json:"end"`
	OperatorId  int         `json:"operator_id"`
	LastMessage LastMessage `json:"last_message"`
}

type GetDialogsParams struct {
	Id         int
	Limit      int
	OperatorId int
	State      string
	Order      string
}

type GetDialogsResp struct {
	Data   []Dialog `json:"data"`
	Status string   `json:"status"`
}
