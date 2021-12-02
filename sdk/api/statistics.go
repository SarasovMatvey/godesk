package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/SarasovMatvey/godesk/client"
	"github.com/SarasovMatvey/godesk/sdk/domain"
)

type Statistics struct {
	Client client.Client
}

func (s *Statistics) GetRequestStats(date string, meta *domain.MetaParams) (domain.GetRequestStatsResp, error) {
	urlValues := url.Values{}

	if meta != nil {
		if meta.Limit != 0 {
			urlValues.Set("limit", strconv.Itoa(meta.Limit))
		}
		if meta.Offset != 0 {
			urlValues.Set("offset", strconv.Itoa(meta.Offset))
		}
	}

	baseUrl := s.Client.Config.BaseUrl + fmt.Sprintf("statistics?report=request_stats&date=%v", date)
	resultUrl := baseUrl + urlValues.Encode()

	req, err := http.NewRequest("GET", resultUrl, nil)
	if err != nil {
		return domain.GetRequestStatsResp{}, err
	}

	resp := new(domain.GetRequestStatsResp)

	err = s.Client.Exec(*req, resp)
	if err != nil {
		return domain.GetRequestStatsResp{}, err
	}

	return *resp, nil
}
