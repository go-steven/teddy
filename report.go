package teddy

import (
	"errors"
	"github.com/bububa/ljson"
)

type ReportRequest struct {
	BaseRequest

	OrderId string
	CName   string
}

type ReportResponse struct {
	SuccessNum     string      `json:"success_num" codec:"success_num"`
	FailNum        string      `json:"fail_num" codec:"fail_num"`
	TotalNum       string      `json:"total_num" codec:"total_num"`
	DownloadReport interface{} `json:"download_report" codec:"download_report"`
}

func Report(apiReq *ReportRequest) (*ReportResponse, error) {
	if apiReq.OrderId == "" {
		return nil, errors.New("no orderId.")
	}

	client := NewClient(apiReq.Account, apiReq.Password, apiReq.IsTest)
	req := NewRequest(URL_REPORT, "")
	req.Params["orderId"] = apiReq.OrderId
	if apiReq.CName != "" {
		req.Params["cname"] = apiReq.CName
	}

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
