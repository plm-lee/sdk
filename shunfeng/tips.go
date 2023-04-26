package shunfeng

import (
	"encoding/json"
	"time"
)

// OrderAddTipRequestParams 顺丰订单加小费接口参数
type OrderAddTipRequestParams struct {
	DevId        int    `json:"dev_id"`
	OrderId      string `json:"order_id"`
	OrderType    int    `json:"order_type,omitempty"` // 1、顺丰订单号 2、商家订单号
	ShopId       string `json:"shop_id,omitempty"`
	ShopType     int    `json:"shop_type,omitempty"`
	GratuityFee  int    `json:"gratuity_fee"` // 订单小费，单位分，加小费最低不能少于100分
	SerialNumber string `json:"serial_number,omitempty"`
	PushTime     int64  `json:"push_time"` //推送时间；秒级时间戳
}

// OrderAddTips 顺丰订单加小费接口
func OrderAddTips(req AddTipRequest) (err error) {
	order := &OrderAddTipRequestParams{
		DevId:       getInstance().devId,
		OrderId:     req.DeliveryId, // 默认使用顺丰订单号
		OrderType:   1,
		ShopId:      req.ShopId,
		ShopType:    1, // 1: 使用顺丰店铺ID 2: 使用系统内部店铺id
		GratuityFee: int(req.Tips * 100),
		PushTime:    time.Now().Unix(),
	}

	sign := getSign(order)
	url := getUrl(api_orderAddTip, sign)

	body, err := json.Marshal(order)
	if err != nil {
		return
	}

	_, err = doPost(url, body)

	return
}
