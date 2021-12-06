package domain

type Request struct {
	ClientPhone          string      `json:"client_phone"`
	RequestStart         int         `json:"request_start"`
	ClientRegion         string      `json:"client_region"`
	ClientExtra2         string      `json:"client_extra_2"`
	RequestStartMode     string      `json:"request_start_mode"`
	RequestStarter       string      `json:"request_starter"`
	RequestTime          int         `json:"request_time"`
	FirstOperatorRole    string      `json:"first_operator_role"`
	RequestTags          string      `json:"request_tags"`
	WorkingReactionTime  interface{} `json:"working_reaction_time"`
	ReplyStart           interface{} `json:"reply_start"`
	ChannelId            int         `json:"channel_id"`
	ClientCountry        string      `json:"client_country"`
	ClientInfo           string      `json:"client_info"`
	BanStatus            string      `json:"ban_status"`
	ClientAssignedName   string      `json:"client_assigned_name"`
	IncomingMenuMessages int         `json:"incoming_menu_messages"`
	IncomingMessages     int         `json:"incoming_messages"`
	OutgoingMenuMessages int         `json:"outgoing_menu_messages"`
	FirstOperatorGroup1  string      `json:"first_operator_group1"`
	FirstOperatorName    string      `json:"first_operator_name"`
	OutgoingMessages     int         `json:"outgoing_messages"`
	ClientTags           string      `json:"client_tags"`
	OperatorsInRequest   int         `json:"operators_in_request"`
	RequestStartT        string      `json:"request_start_t"`
	ClientExtra1         string      `json:"client_extra_1"`
	FirstOperatorGroup2  string      `json:"first_operator_group2"`
	ChannelName          string      `json:"channel_name"`
	RequestId            int         `json:"request_id"`
	ClientComment        string      `json:"client_comment"`
	ClientName           string      `json:"client_name"`
	ClientExtra3         string      `json:"client_extra_3"`
	ReactionTime         interface{} `json:"reaction_time"`
	RequestTerminator    string      `json:"request_terminator"`
	ReplyStartD          string      `json:"reply_start_d"`
	FirstOperatorId      interface{} `json:"first_operator_id"`
	RequestStartD        string      `json:"request_start_d"`
	ClientId             int         `json:"client_id"`
	ClientMessenger      string      `json:"client_messenger"`
	ClientIsNew          string      `json:"client_is_new"`
	ReplyStartT          string      `json:"reply_start_t"`
}
