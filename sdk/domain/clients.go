package domain

type Client struct {
	Id           int    `json:"id"`
	Comment      string `json:"comment"`
	AssignedName string `json:"assigned_name"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	RegionId     string `json:"region_id"`
	CountryId    int    `json:"country_id"`
}

type GetClientResp struct {
	Data Client `json:"data"`
}

type GetAllClientsResp struct {
	Data []Client `json:"data"`
}

type AddClientParams struct {
	Phone         string `json:"phone"`
	AssignedPhone string `json:"assigned_phone"`
	Transport     string `json:"transport"`
	ChannelId     int    `json:"channel_id,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
}

type SetClientExtraInfoParams struct {
	Nickname   string `json:"nickname"`
	ExternalId string `json:"external_id"`
}
