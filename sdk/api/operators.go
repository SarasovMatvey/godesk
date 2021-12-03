package api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/domain"
	"github.com/davecgh/go-spew/spew"
)

type Operators struct {
	Client client.Client
}

func (c *Operators) Get(meta *domain.MetaParams) (domain.GetOperatorsResp, error) {
	urlValues := url.Values{}

	if meta != nil {
		if meta.Limit != 0 {
			urlValues.Set("limit", strconv.Itoa(meta.Limit))
		}
		if meta.Offset != 0 {
			urlValues.Set("offset", strconv.Itoa(meta.Offset))
		}
	}

	baseUrl := c.Client.Config.BaseUrl + "operators"
	resultUrl := baseUrl + "?" + urlValues.Encode()

	spew.Dump(resultUrl)

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
