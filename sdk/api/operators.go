package api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/domain"
)

type Operators struct {
	Client client.Client
}

func (c *Operators) Get(limit int, offset int) (domain.GetOperatorsResp, error) {
	urlValues := url.Values{}

	if limit != 0 {
		urlValues.Set("limit", strconv.Itoa(limit))
	}
	if offset != 0 {
		urlValues.Set("offset", strconv.Itoa(offset))
	}

	baseUrl := c.Client.Config.BaseUrl + "operators"
	resultUrl := baseUrl + "?" + urlValues.Encode()

	req, err := http.NewRequest("GET", resultUrl, nil)
	if err != nil {
		return domain.GetOperatorsResp{}, err
	}

	resp := new(domain.GetOperatorsResp)

	err = c.Client.Exec(*req, resp)
	if err != nil {
		return domain.GetOperatorsResp{}, err
	}

	return *resp, nil
}
