package uu

import (
	"encoding/json"
	"errors"
	"sdk/libs/httplib"
	"sdk/libs/utils"
	"time"
)

type BasicRequestParams struct {
	Sign      string `json:"sign"`
	NonceStr  string `json:"nonce_str"`
	Timestamp string `json:"timestamp"`
	Appid     string `json:"appid"`
}

type BasicResponseParams struct {
	Sign       string `json:"sign"`
	AppId      string `json:"appid"`
	ReturnMsg  string `json:"return_msg"`
	ReturnCode string `json:"return_code"`
}

type SendUserPhoneCodeRequestParams struct {
	UserMobile string `json:"user_mobile"`
	UserIp     string `json:"user_ip"`
	ImageCode  string `json:"image_code,omitempty"`
	BasicRequestParams
}

type SendUserPhoneCodeResponse struct {
	NonceStr   string `json:"nonce_str"`
	Sign       string `json:"sign"`
	AppId      string `json:"appid"`
	ReturnMsg  string `json:"return_msg"`
	ReturnCode string `json:"return_code"`
}

type GetUserOpenIdRequestParams struct {
	UserMobile   string `json:"user_mobile"`
	ValidateCode string `json:"validate_code"`
	CityName     string `json:"city_name"`
	CountyName   string `json:"county_name,omitempty"`
	RegIp        string `json:"reg_ip,omitempty"`
	BasicRequestParams
}

type GetUserOpenIdResponse struct {
	OpenId      string `json:"openid"`
	OpenService string `json:"openservice,omitempty"` // 开通配送服务类型（UU专送/UU团送）
	BasicResponseParams
}

// SendUserPhoneCode 给uu用户发送授权验证码
func SendUserPhoneCode(phone, ip string) (err error) {
	req := SendUserPhoneCodeRequestParams{
		UserMobile: phone,
		UserIp:     ip,
		BasicRequestParams: BasicRequestParams{
			NonceStr:  utils.UUID(32),
			Timestamp: utils.ToString(time.Now().Unix()),
			Appid:     getInstance().appId,
		},
	}
	// 参数以form－data方式传递，不是json方式
	tmp, err := json.Marshal(req)
	if err != nil {
		return
	}

	params := make(map[string]interface{})
	err = json.Unmarshal(tmp, &params)
	if err != nil {
		return
	}

	params["sign"] = getSign(params)

	url := getUrl(api_bindUser_sendPhoneCode)
	body, err := httplib.PostForm(url, params)
	if err != nil {
		return
	}

	var resp SendUserPhoneCodeResponse
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return
	}

	if resp.ReturnCode != "ok" {
		err = errors.New(resp.ReturnMsg)
	}

	return
}

// GetUserOpenId 获取授权用户openId
func GetUserOpenId(phone, code, city string) (openId, service string, err error) {
	req := GetUserOpenIdRequestParams{
		UserMobile:   phone,
		ValidateCode: code,
		CityName:     city,
		BasicRequestParams: BasicRequestParams{
			NonceStr:  utils.UUID(32),
			Timestamp: utils.ToString(time.Now().Unix()),
			Appid:     getInstance().appId,
		},
	}
	// 参数以form－data方式传递，不是json方式
	tmp, err := json.Marshal(req)
	if err != nil {
		return
	}

	params := make(map[string]interface{})
	err = json.Unmarshal(tmp, &params)
	if err != nil {
		return
	}

	params["sign"] = getSign(params)

	url := getUrl(api_bindUser_getOpenId)
	body, err := httplib.PostForm(url, params)
	if err != nil {
		return
	}

	var resp GetUserOpenIdResponse
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return
	}

	if resp.ReturnCode != "ok" {
		err = errors.New(resp.ReturnMsg)
		return
	}

	openId = resp.OpenId
	service = resp.OpenService

	return
}
