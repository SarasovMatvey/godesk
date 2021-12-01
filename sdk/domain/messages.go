package domain

type Message struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	Transport string `json:"transport"`
	ClientId  int    `json:"client_id"`
}

type GetAllMessagesParams struct {
	ChannelId int    `json:"channel_id,omitempty"`
	Transport string `json:"transport,omitempty"`
	Type      string `json:"type,omitempty"`
}

type GetAllMessagesResp struct {
	Data   []Message `json:"data"`
	Status string    `json:"status"`
}

type AddInboxMessageParams struct {
	ChannelId  string `json:"channel_id"`
	Body       string `json:"body"`
	FromClient struct {
		Id    string `json:"id"`
		Phone int    `json:"phone"`
	} `json:"from_client"`
}
