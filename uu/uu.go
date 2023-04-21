package uu

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/plm-lee/sdk/libs/utils"
	"sort"
	"strings"
)

const (
	register_code              = "UU"
	server_host_test           = "http://openapi.test.uupt.com/v2_0/" // 测试环境接口地址
	server_host                = "https://openapi.uupt.com/v2_0/"     // 正式环境接口地址
	api_queryDeliverFee        = "getorderprice.ashx"
	api_bindUser_sendPhoneCode = "binduserapply.ashx"
	api_bindUser_getOpenId     = "bindusersubmit.ashx"
	api_addOrder               = "addorder.ashx"
	api_cancelOrder            = "cancelorder.ashx"
	api_gethomeservicefee      = "gethomeservicefee.ashx" // 查违约金
	api_addTips                = "/payonlinefee.ashx"     // 加小费
	api_getCityList            = "citylist.ashx"          // 开通城市列表
)

var (
	app *AppConfig
)

type AppConfig struct {
	isDevelopment          bool
	appId                  string
	appKey                 string
	openId                 string
	orderStatusCallbackUrl string
}

func InitAppConfig(appId, appKey, callbackUrl string, development bool) {
	app = &AppConfig{
		appId:                  appId,
		appKey:                 appKey,
		orderStatusCallbackUrl: callbackUrl,
		isDevelopment:          development,
	}
}

func getInstance() *AppConfig {
	return app
}

func getSign(params map[string]interface{}) string {
	// 排除空值参数，根据key字典序排序
	sortKey := make([]string, 0)
	for key, v := range params {
		if v == "" {
			continue
		}
		sortKey = append(sortKey, key)
	}
	sort.Strings(sortKey) // 排序

	character := "" //签名排序串
	for _, key := range sortKey {
		v := utils.ToString(params[key])
		character += fmt.Sprintf("&%s=%s", key, v)
	}

	character = strings.TrimPrefix(character, "&")
	character += "&key=" + getInstance().appKey // 加上应用密钥
	character = strings.ToUpper(character)      // 所有字母转为大写

	tmp := md5.Sum([]byte(character))
	sign := hex.EncodeToString(tmp[:])
	sign = strings.ToUpper(sign) // 转大写
	return sign
}

func getUrl(url string) string {
	if getInstance().isDevelopment {
		return server_host_test + url
	}

	return server_host + url
}
