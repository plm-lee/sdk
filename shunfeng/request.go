package shunfeng

import (
	"encoding/json"
	"errors"
	"github.com/plm-lee/sdk/libs/httplib"
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
	Result    interface{} `json:"result"`
}

// 发起请求
func doPost(url string, body []byte) (respBody []byte, err error) {
	req := httplib.Post(url)

	req.Body(body)
	rsp, err := req.Bytes()
	if err != nil {
		return nil, err
	}

	var result Response
	err = json.Unmarshal(rsp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrorCode != 0 {
		return nil, errors.New(result.ErrorMsg)
	}

	respBody, err = json.Marshal(result.Result)

	return
}
