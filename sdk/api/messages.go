package api

import (
	"net/http"

	"gitlab.com/SarasovMatvey/godesk/client"
	"gitlab.com/SarasovMatvey/godesk/sdk/domain"
	"gitlab.com/SarasovMatvey/godesk/sdk/helper"
)

type Messages struct {
	Client client.Client
}

func (m Messages) GetAll(params domain.GetAllMessagesParams) (domain.GetAllMessagesResp, error) {
	body := helper.ObjectToJsonReader(params)

	req, err := http.NewRequest("GET", m.Client.Config.BaseUrl+"messages", body)
	if err != nil {
		return domain.GetAllMessagesResp{}, err
	}

	resp := new(domain.GetAllMessagesResp)

	err = m.Client.Exec(*req, resp)
	if err != nil {
		return domain.GetAllMessagesResp{}, err
	}

	return *resp, nil
}

func (m Messages) AddInbox(params domain.AddInboxMessageParams) (domain.StatusResp, error) {
	body := helper.ObjectToJsonReader(params)

	req, err := http.NewRequest("POST", m.Client.Config.BaseUrl+"messages/inbox", body)
	if err != nil {
		return domain.StatusResp{}, err
	}

	resp := new(domain.StatusResp)

	err = m.Client.Exec(*req, resp)
	if err != nil {
		return domain.StatusResp{}, err
	}

	return *resp, nil
}
