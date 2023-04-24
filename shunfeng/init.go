package shunfeng

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	server_host   = "https://openic.sf-express.com"
	register_code = "SF"

	api_queryDeliverFee = "/open/api/external/precreateorder"
	api_createOrder     = "/open/api/external/createorder"
	api_cancelOrder     = "/open/api/external/cancelorder"
	api_orderAddTip     = "/open/api/external/addordergratuityfee"
)

var (
	app *AppConfig
)

type AppConfig struct {
	devId  int
	devKey string
	host   string

	cityShopCode map[string]string
}

func InitAppConfig(devId int, devKey, host string) {
	app = &AppConfig{
		devId:        devId,
		devKey:       devKey,
		cityShopCode: make(map[string]string),
	}
}

func getInstance() *AppConfig {
	return app
}

// InitCityCode 初始化城市-code对应关系
func InitCityCode(data map[string]string) {
	if app != nil {
		app.cityShopCode = data
	}
}

func getSign(params interface{}) string {
	data, err := json.Marshal(params)
	if err != nil {
		return ""
	}

	character := fmt.Sprintf("%s&%d&%s", data, getInstance().devId, getInstance().devKey)
	tmp := md5.Sum([]byte(character))
	sign := hex.EncodeToString(tmp[:])
	sign = base64.StdEncoding.EncodeToString([]byte(sign))
	return sign
}

func getUrl(api, sign string) string {
	return fmt.Sprintf("%s%s?sign=%s", server_host, api, sign)
}

// GetCodeByCity 获取城市对应大区店铺Id
func GetCodeByCity(cityName string) string {
	if cityName == "" {
		return ""
	}

	if cityCode, ok := getInstance().cityShopCode[cityName]; ok {
		return cityCode
	}

	for k, v := range getInstance().cityShopCode {
		if strings.Contains(k, cityName) {
			return v
		}

		if strings.Contains(cityName, k) {
			return v
		}
	}

	return ""
}
