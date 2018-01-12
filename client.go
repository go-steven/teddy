package teddy

import (
	"errors"
	"github.com/bububa/ljson"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	MethodUrl string
	Params    map[string]interface{}
	UKey      string
}

func NewRequest(methodUrl string, ukey string) *Request {
	return &Request{
		MethodUrl: methodUrl,
		Params:    make(map[string]interface{}),
		UKey:      strings.TrimSpace(ukey),
	}
}

type Response struct {
	ErrorCode uint8       `json:"error_code" codec:"error_code"`
	ErrorMsg  string      `json:"error_msg" codec:"error_msg"`
	Data      interface{} `json:"data" codec:"data"`
}

type Client struct {
	Account  string
	Password string
	IsTest   bool
}

//create new client
func NewClient(account, password string, isTest bool) (c *Client) {
	c = &Client{
		Account:  account,
		Password: password,
		IsTest:   isTest,
	}
	return
}

func (c *Client) Execute(req *Request) (interface{}, error) {
	sysParams := make(map[string]interface{})
	sysParams["account"] = c.Account
	sysParams["password"] = c.Password
	if req.UKey != "" {
		sysParams["ukey"] = req.UKey
	}
	if req.MethodUrl == URL_MULTI_SEND {
		sysParams["data"] = req.Params
	} else {
		for k, v := range req.Params {
			sysParams[k] = v
		}
	}

	gatewayUrl := GATEWAY_URL
	if c.IsTest {
		gatewayUrl = TEST_GATEWAY_URL
	}
	//fmt.Println(Json(sysParams))
	reqUrl := gatewayUrl + req.MethodUrl
	response, err := http.DefaultClient.Post(reqUrl, "application/json; charset=UTF-8", strings.NewReader(Json(sysParams)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	j := Response{}
	err = ljson.Unmarshal(body, &j)
	if err == nil {
		if err == nil && j.ErrorCode != 0 {
			return nil, errors.New(string(j.ErrorMsg))
		}
	}
	//fmt.Println(Json(j))
	return j.Data, nil
}
