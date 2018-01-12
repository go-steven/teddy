package teddy

import (
	"errors"
	"github.com/bububa/ljson"
)

type ReportByTagRequest struct {
	BaseRequest

	SubmitDate string
	CName      string
}

func ReportByTag(apiReq *ReportByTagRequest) (*ReportResponse, error) {
	if apiReq.CName == "" || apiReq.SubmitDate == "" {
		return nil, errors.New("invalid params.")
	}

	client := NewClient(apiReq.Account, apiReq.Password, apiReq.IsTest)
	req := NewRequest(URL_REPORT_BY_TAG, "")
	req.Params["submitdate"] = apiReq.SubmitDate
	req.Params["cname"] = apiReq.CName

	response, err := client.Execute(req)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("ret: %s\n", Json(response))
	j := ReportResponse{}
	err = ljson.Unmarshal([]byte(Json(response)), &j)
	if err != nil {
		return nil, err
	}

	return &j, nil
}
