package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/domain"
)

type Dialogs struct {
	Client client.Client
}

func (d *Dialogs) Get(params domain.GetDialogsParams, meta *domain.MetaParams) (domain.GetDialogsResp, error) {
	var baseUrl string

	if params.Id != 0 {
		baseUrl = d.Client.Config.BaseUrl + fmt.Sprintf("dialogs/%v", params.Id)
	} else {
		baseUrl = d.Client.Config.BaseUrl + "dialogs"
	}

	urlValues := url.Values{}

	if meta != nil {
		if meta.Limit != 0 {
			urlValues.Set("limit", strconv.Itoa(meta.Limit))
		}
		if meta.Offset != 0 {
			urlValues.Set("offset", strconv.Itoa(meta.Limit))
		}
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

	req, err := http.NewRequest("GET", resultUrl, nil)
	if err != nil {
		return domain.GetDialogsResp{}, err
	}

	resp := new(domain.GetDialogsResp)

	_ = d.Client.Exec(*req, resp)

	return *resp, nil
}
