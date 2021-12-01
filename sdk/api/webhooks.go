package api

import (
	"net/http"

	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/domain"
	"github.com/SarasovMatvey/godesk/sdk/helper"
)

type Webhooks struct {
	Client client.Client
}

func (w *Webhooks) Add(params domain.AddWebhooksParams) error {
	body := helper.ObjectToJsonReader(params)

	req, err := http.NewRequest("POST", w.Client.Config.BaseUrl+"webhooks", body)

	if err != nil {
		return err
	}

	err = w.Client.Exec(*req, nil)
	if err != nil {
		return err
	}

	return nil
}
