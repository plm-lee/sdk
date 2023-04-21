package uu

import (
	"encoding/json"
	"errors"
	"github.com/plm-lee/sdk/libs/httplib"
	"github.com/plm-lee/sdk/libs/utils"
	"time"
)

type OrderAddTipRequestParams struct {
	OrderCode string `json:"order_code"` // 订单编号 uu平台配送订单号
	OnlineFee int    `json:"onlinefee"`  // 支付小费金额，最低1元，只能整数（不支持小数）
	OriginId  string `json:"origin_id"`  // 第三方订单编号
	OpenId    string `json:"openid"`
	BasicRequestParams
}

type OrderAddTipResponse struct {
	BasicResponseParams
}

func OrderAddTips(r AddTipRequest) (err error) {
	order := &OrderAddTipRequestParams{
		OrderCode: r.DeliveryId,
		OnlineFee: int(r.Tips),
		OriginId:  r.OrderId,
		BasicRequestParams: BasicRequestParams{
			NonceStr:  utils.UUID(32),
			Timestamp: utils.ToString(time.Now().Unix()),
			Appid:     getInstance().appId,
		},
		OpenId: r.UserOpenId,
	}

	// 参数以form－data方式传递，不是json方式
	tmp, err := json.Marshal(order)
	if err != nil {
		return
	}

	params := make(map[string]interface{})
	err = json.Unmarshal(tmp, &params)
	if err != nil {
		return
	}

	params["sign"] = getSign(params)

	url := getUrl(api_addTips)
	body, err := httplib.PostForm(url, params)
	if err != nil {
		return
	}

	var resp OrderAddTipResponse
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return
	}

	if resp.ReturnCode != "ok" {
		err = errors.New(resp.ReturnMsg)
	}

	return
}
