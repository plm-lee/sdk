package uu

import (
	"encoding/json"
	"errors"
	"github.com/plm-lee/sdk/libs/httplib"
	"github.com/plm-lee/sdk/libs/utils"
	"time"
)

type GetOrderPriceRequestParams struct {
	OriginId        string `json:"origin_id"`                   // 第三方对接平台订单id
	FromAddress     string `json:"from_address"`                // 起始地址
	FromUserNote    string `json:"from_usernote,omitempty"`     // 起始地址具体门牌号
	ToAddress       string `json:"to_address"`                  // 目的地址
	ToUserNote      string `json:"to_usernote,omitempty"`       // 目的地址具体门牌号
	CityName        string `json:"city_name"`                   // 订单所在城市名 称(如郑州市就填”郑州市“，必须带上“市”)
	SubscribeType   string `json:"subscribe_type,omitempty"`    // 预约类型 0实时订单 1预约取件时间
	CountyName      string `json:"county_name,omitempty"`       // 订单所在县级地名称(如金水区就填“金水区”)
	SubscribeTime   string `json:"subscribe_time,omitempty"`    // 预约时间（如：2015-06-18 12:00:00）没有可以传空字符串
	GoodsType       string `json:"goods_type,omitempty"`        // 物品类型：（美食、鲜花、蛋糕、手机、钥匙、文件、其他）仅支持指定类型不包含按其他发布
	CouponId        string `json:"coupon_id,omitempty"`         // 优惠券ID(如果传入-1就不用优惠券否则系统自动匹配)
	SendType        string `json:"send_type"`                   // 订单小类 0帮我送(默认) 1帮我买
	ToLat           string `json:"to_lat"`                      // 目的地坐标纬度(坐标系为百度地图坐标系)
	ToLng           string `json:"to_lng"`                      // 目的地坐标经度(坐标系为百度地图坐标系)
	FromLat         string `json:"from_lat"`                    // 起始地坐标纬度(坐标系为百度地图坐标系)
	FromLng         string `json:"from_lng"`                    // 起始地坐标经度(坐标系为百度地图坐标系)
	CouponType      string `json:"coupon_type,omitempty"`       // 优惠券类型 默认空 个人优惠券 1=企业优惠券（必须企业帐号才可以使用）
	GoodsWeightCode string `json:"goods_weight_code,omitempty"` // 物品重量配置信息code 默认空（code编码详情见 获取计价物品重量配置信息接口）
	ShopId          string `json:"shop_id,omitempty"`           // 门店编号（门店列表可查看门店编号）
	Sign            string `json:"sign"`
	NonceStr        string `json:"nonce_str"` // 随机字符串，不长于32位
	Timestamp       string `json:"timestamp"`
	Openid          string `json:"openid"` // 用户openid,详情见 获取openid接口
	Appid           string `json:"appid"`
}

// GetOrderPriceResponse 查询订单价格接口返回
type GetOrderPriceResponse struct {
	OriginId            string `json:"origin_id"`
	PriceToken          string `json:"price_token"`    // 金额令牌，提交订单前必须先计算价格
	TotalMoney          string `json:"total_money"`    // 订单总金额
	NeedPaymoney        string `json:"need_paymoney"`  // 实际支付金额
	TotalPriceoff       string `json:"total_priceoff"` // 总优惠金额
	Distance            string `json:"distance"`
	FreightMoney        string `json:"freight_money"`        // 跑腿费
	CouponId            string `json:"couponid"`             // 优惠券ID
	CouponAmount        string `json:"coupon_amount"`        // 优惠券金额
	Addfee              string `json:"addfee"`               // 加价金额
	GoodsInsuranceMoney string `json:"goods_insurancemoney"` // 商品保价金额
	ExpiresIn           string `json:"expires_in"`           // price_token的过期时间（单位：秒）
	NonceStr            string `json:"nonce_str"`
	Sign                string `json:"sign"`
	AppId               string `json:"appid"`
	ReturnMsg           string `json:"return_msg"`
	ReturnCode          string `json:"return_code"`
}

// GetOrderPrice 从uu接口询价
func GetOrderPrice(r QueryDeliveryFeeRequest) (resp GetOrderPriceResponse, err error) {
	order := &GetOrderPriceRequestParams{
		OriginId:      r.OrderId + "_" + utils.UUID(4), // 订单号加上4位随机数
		ToAddress:     r.ReceiverAddress,
		CityName:      r.CityName,
		SubscribeType: "0",
		SendType:      "0",
		ToLat:         r.ReceiverLat,
		ToLng:         r.ReceiverLng,
		FromLat:       r.FromLat,
		FromLng:       r.FromLng,
		FromAddress:   r.FromAddress,
		NonceStr:      utils.UUID(32),
		Timestamp:     utils.ToString(time.Now().Unix()),
		Openid:        r.UserOpenId,
		Appid:         getInstance().appId,
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

	url := getUrl(api_queryDeliverFee)
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
