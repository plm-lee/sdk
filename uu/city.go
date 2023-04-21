package uu

import (
	"encoding/json"
	"errors"
	"github.com/plm-lee/sdk/libs/httplib"
	"github.com/plm-lee/sdk/libs/utils"
	"time"
)

type GetCityListRequestParams struct {
	OpenId string `json:"openid"`
	BasicRequestParams
}

type GetCityListResponse struct {
	CityList []struct {
		CityName   string `json:"city_name"`
		CountyName string `json:"county_name"`
	} `json:"citylist"`
	BasicResponseParams
}

// GetCityList 查询已开通配送服务城市列表
func GetCityList(openId string) (resp GetCityListResponse, err error) {
	req := GetCityListRequestParams{
		OpenId: openId,
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

	url := getUrl(api_getCityList)
	body, err := httplib.PostForm(url, params)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return
	}

	if resp.ReturnCode != "ok" {
		err = errors.New(resp.ReturnMsg)
	}

	return
}
