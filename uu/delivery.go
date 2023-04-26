package uu

import (
	"encoding/json"
	"errors"
	"github.com/plm-lee/sdk/libs/httplib"
	"github.com/plm-lee/sdk/libs/utils"
	"time"
)

// SendOrderDeliveryRequestParams 订单发配送请求结构
type SendOrderDeliveryRequestParams struct {
	PriceToken      string `json:"price_token"`             // 金额令牌，计算订单价格接口返回的price_token
	OrderPrice      string `json:"order_price"`             // 订单金额，计算订单价格接口返回的total_money
	BalancePayMoney string `json:"balance_paymoney"`        // 实际余额支付金额计算订单价格接口返回的need_paymoney
	Receiver        string `json:"receiver"`                // 收件人
	ReceivePhone    string `json:"receiver_phone"`          // 收件人电话 手机号码； 虚拟号码格式（手机号_分机号码）例如：13700000000_1111
	Note            string `json:"note,omitempty"`          // 订单备注 最长140个汉字
	CallbackUrl     string `json:"callback_url,omitempty"`  // 订单提交成功后及状态变化的回调地址
	PushType        string `json:"push_type"`               // 推送方式（0 开放订单，2测试订单）默认传0即可
	SpecialType     string `json:"special_type"`            // 特殊处理类型，是否需要保温箱 1需要 0不需要
	CallmeWithtake  string `json:"callme_withtake"`         // 取件是否给我打电话 1需要 0不需要
	PubuserMobile   string `json:"pubusermobile,omitempty"` // 发件人电话，（如果为空则是用户注册的手机号）
	PayType         string `json:"pay_type,omitempty"`      // 支付方式：1=企业支付 0账户余额支付（企业余额不足自动转账户余额支付）
	OrderSource     string `json:"ordersource,omitempty"`   // 订单来源标示,示例（1=美团 2=饿了么 3=其他）
	ShortOrderNum   string `json:"shortordernum,omitempty"` // 订单平台短单号（0-10000），该单号会展示给骑手突出展示：“美团 #1”
	Sign            string `json:"sign"`
	NonceStr        string `json:"nonce_str"`
	Timestamp       string `json:"timestamp"`
	OpenId          string `json:"openid"`
	AppId           string `json:"appid"`
}

// SendOrderDeliveryResponse 订单发配送返回
type SendOrderDeliveryResponse struct {
	OrderCode  string `json:"ordercode"`
	OriginId   string `json:"origin_id"`
	NonceStr   string `json:"nonce_str"`
	Sign       string `json:"sign"`
	AppId      string `json:"appid"`
	ReturnMsg  string `json:"return_msg"`
	ReturnCode string `json:"return_code"`
}

// DeliveryCancelRequestParams 取消配送
type DeliveryCancelRequestParams struct {
	OriginId  string `json:"origin_id"`  // 第三方对接平台订单id，order_code和origin_id必须二选其一，如果都传，则只根据order_code返回
	OrderCode string `json:"order_code"` // UU跑腿订单编号，order_code和origin_id必须二选其一，如果都传，则只根据order_code返回
	Reason    string `json:"reason"`     // 取消原因
	Sign      string `json:"sign"`
	NonceStr  string `json:"nonce_str"`
	Timestamp string `json:"timestamp"`
	OpenId    string `json:"openid"`
	AppId     string `json:"appid"`
}

// DeliveryCancelResponse 取消配送接口返回
type DeliveryCancelResponse struct {
	OrderCode  string `json:"order_code"` // 订单号
	OriginId   string `json:"origin_id"`  // 第三方对接平台订单id
	NonceStr   string `json:"nonce_str"`
	Timestamp  string `json:"timestamp"`
	ReturnMsg  string `json:"return_msg"`
	ReturnCode string `json:"return_code"`
}

// SendOrderDelivery 发uu配送
func SendOrderDelivery(r OrderDeliveryRequest) (orderId string, err error) {
	order := &SendOrderDeliveryRequestParams{
		PriceToken:      r.PriceToken,
		OrderPrice:      r.OriginPrice,
		BalancePayMoney: r.PayedPrice,
		Receiver:        r.ReceiverName,
		ReceivePhone:    r.ReceiverPhone,
		Note:            r.Remarks,
		CallbackUrl:     getInstance().orderStatusCallbackUrl, //
		PushType:        "0",
		SpecialType:     "0", // 保温箱
		CallmeWithtake:  "1", // 需要电话通知
		PayType:         "1", // 1=企业支付
		NonceStr:        utils.UUID(32),
		Timestamp:       utils.ToString(time.Now().Unix()),
		OpenId:          r.UserOpenId,
		AppId:           getInstance().appId,
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

	url := getUrl(api_addOrder)
	body, err := httplib.PostForm(url, params)
	if err != nil {
		return
	}

	var result SendOrderDeliveryResponse
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return
	}

	if result.ReturnCode != "ok" {
		err = errors.New(result.ReturnMsg)
		return
	}

	orderId = result.OrderCode
	return
}

// CancelOrderDelivery 取消配送
func CancelOrderDelivery(r CancelOrderRequest) (err error) {
	order := &DeliveryCancelRequestParams{
		OrderCode: r.DeliveryId,
		OriginId:  r.OrderId,
		Reason:    r.CancelReason,
		Timestamp: utils.ToString(time.Now().Unix()),
		NonceStr:  utils.UUID(32),
		OpenId:    r.UserOpenId,
		AppId:     getInstance().appId,
	}
	if order.Reason == "" {
		order.Reason = "用户取消订单"
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

	url := getUrl(api_cancelOrder)
	body, err := httplib.PostForm(url, params)
	if err != nil {
		return
	}

	var result SendOrderDeliveryResponse
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return
	}

	if result.ReturnCode != "ok" {
		err = errors.New(result.ReturnMsg)
	}

	return
}
