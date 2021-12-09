package api

import (
	"fmt"
	"net/http"

	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/domain"
	"github.com/SarasovMatvey/godesk/sdk/helper"
)

type Clients struct {
	Client client.Client
}

func (c *Clients) Get(clientId int) (domain.GetClientResp, error) {
	url := fmt.Sprintf("clients/%v", clientId)
	req, err := http.NewRequest("GET", c.Client.Config.BaseUrl+url, nil)
	if err != nil {
		return domain.GetClientResp{}, err
	}

	resp := new(domain.GetClientResp)

	err = c.Client.Exec(*req, resp)

	return *resp, nil
}

func (c *Clients) GetAll() (domain.GetAllClientsResp, error) {
	req, err := http.NewRequest("GET", c.Client.Config.BaseUrl+"clients", nil)
	if err != nil {
		return domain.GetAllClientsResp{}, err
	}

	resp := new(domain.GetAllClientsResp)

	err = c.Client.Exec(*req, resp)

	return *resp, nil
}

func (c *Clients) Add(params domain.AddClientParams) (domain.GetClientResp, error) {
	body := helper.ObjectToJsonReader(params)

	req, err := http.NewRequest("POST", c.Client.Config.BaseUrl+"clients", body)
	if err != nil {
		return domain.GetClientResp{}, err
	}

	resp := new(domain.GetClientResp)

	err = c.Client.Exec(*req, resp)

	return *resp, nil
}

func (c *Clients) SetClientExtraInfo(params domain.SetClientExtraInfoParams) (domain.StatusResp, error) {
	body := helper.ObjectToJsonReader(params)

	req, err := http.NewRequest("PUT", c.Client.Config.BaseUrl+"clients", body)
	if err != nil {
		return domain.StatusResp{}, err
	}

	resp := new(domain.StatusResp)

	err = c.Client.Exec(*req, resp)

	return *resp, nil
}
