package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/domain"
	"github.com/davecgh/go-spew/spew"
)

type Dialogs struct {
	Client client.Client
}

func (d *Dialogs) Get(params domain.GetDialogsParams) (domain.GetDialogsResp, error) {
	var baseUrl string

	if params.Id != 0 {
		baseUrl = d.Client.Config.BaseUrl + fmt.Sprintf("dialogs/%v", params.Id)
	} else {
		baseUrl = d.Client.Config.BaseUrl + "dialogs"
	}

	urlValues := url.Values{}

	if params.Limit != 0 {
		urlValues.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.OperatorId != 0 {
		urlValues.Set("operator_id", strconv.Itoa(params.OperatorId))
	}
	if params.State != "" {
		urlValues.Set("state", params.State)
	}
	if params.Order != "" {
		urlValues.Set("order", params.Order)
	}

	resultUrl := baseUrl + "?" + urlValues.Encode()

	spew.Dump(resultUrl)

	req, err := http.NewRequest("GET", resultUrl, nil)
	if err != nil {
		return domain.GetDialogsResp{}, err
	}

	resp := new(domain.GetDialogsResp)

	err = d.Client.Exec(*req, resp)
	if err != nil {
		return domain.GetDialogsResp{}, err
	}

	return *resp, nil
}
