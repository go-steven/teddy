package teddy

import (
	"errors"
	"github.com/bububa/ljson"
	"strings"
)

type MultiSendRequest struct {
	BaseRequest

	Phones   []string
	Content  string
	Sign     string
	CName    string
	SendTime string
	UKey     string
}

type MultiSendResponse struct {
	OrderId string `json:"orderId" codec:"orderId"`
	Result  bool   `json:"result" codec:"result"`
}

func MultiSend(apiReq *MultiSendRequest) (string, error) {
	if len(apiReq.Phones) == 0 {
		return "", errors.New("no phones.")
	}

	client := NewClient(apiReq.Account, apiReq.Password, apiReq.IsTest)
	req := NewRequest(URL_MULTI_SEND, apiReq.UKey)
	req.Params["phones"] = strings.Join(apiReq.Phones, ",")
	req.Params["content"] = apiReq.Content
	req.Params["sign"] = apiReq.Sign
	if apiReq.CName != "" {
		req.Params["cname"] = apiReq.CName
	}
	if apiReq.SendTime != "" {
		req.Params["sendtime"] = apiReq.SendTime
	}

	response, err := client.Execute(req)
	if err != nil {
		return "", err
	}
	//fmt.Printf("ret: %s\n", Json(response))
	j := MultiSendResponse{}
	err = ljson.Unmarshal([]byte(Json(response)), &j)
	if err != nil {
		return "", err
	}
	if !j.Result {
		return "", errors.New(Json(response))
	}

	return j.OrderId, nil
}
