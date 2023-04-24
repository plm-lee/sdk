package shunfeng

import (
	"encoding/json"
	"github.com/plm-lee/sdk/libs/utils"
	"time"
)

type SendOrderDeliveryRequestParams struct {
	DevId            int    `json:"dev_id"`                       //api开发者ID
	ShopId           string `json:"shop_id"`                      //店铺ID 城市大区店铺
	ShopType         int    `json:"shop_type,omitempty"`          //店铺ID类型 1:顺丰店铺ID 2:接入方店铺ID
	Weight           int    `json:"weight,omitempty"`             //物品重量（单位：克）
	IsAppoint        int    `json:"is_appoint"`                   //是否是预约单	0：非预约单；1：预约单
	AppointType      int    `json:"appoint_type,omitempty"`       //预约单类型	 预约单的时候传入，1：预约单送达单；2：预约单上门单
	ExpectTime       int    `json:"expect_time,omitempty"`        //用户期望送达时间 若传入自此段且时间大于配送时效，则按照预约送达单处理，时间小于配送时效按照立即单处理；appoint_type=1时需必传,秒级时间戳
	ExpectPickupTime int    `json:"expect_pickup_time,omitempty"` //用户期望上门时间 appoint_type=2时需必传,秒级时间戳
	LbsType          int    `json:"lbs_type,omitempty"`           //坐标类型，1：百度坐标，2：高德坐标
	PayType          int    `json:"pay_type"`                     //用户支付方式：1、已支付 0、货到付款
	ReceiveUserMoney int    `json:"receive_user_money,omitempty"` //代收金额
	IsInsured        int    `json:"is_insured"`                   //是否保价，0：非保价；1：保价
	IsPersonDirect   int    `json:"is_person_direct"`             //是否是专人直送订单，0：否；1：是
	Vehicle          int    `json:"vehicle,omitempty"`            //配送交通工具，0：否；1：电动车；2：小轿车
	DeclaredValue    int    `json:"declared_value,omitempty"`     //保价金额	单位：分
	GratuityFee      int    `json:"gratuity_fee,omitempty"`       //订单小费，不传或者传0为不加小费 单位分，加小费最低不能少于100分
	Remark           string `json:"remark,omitempty"`
	RiderPickMethod  int    `json:"rider_pick_method,omitempty"` //物流流向 1：从门店取件送至用户； 2：从用户取件送至门店
	ReturnFlag       int    `json:"return_flag,omitempty"`       //返回字段控制标志位（二进制） 1:商品总价格，2:配送距离，4:物品重量，8:起送时间，16:期望送达时间，32:支付费用，64:实际支持金额，128:优惠卷总金额，256:结算方式 例如全部返回为填入511
	PushTime         int64  `json:"push_time"`                   //推单时间
	Shop             *Shop  `json:"shop,omitempty"`              //发货店铺信息 平台级开发者需要传入

	// 发单
	ShopOrderId   string  `json:"shop_order_id,omitempty"`  //商家订单号
	OrderSource   string  `json:"order_source,omitempty"`   //订单来源1：美团；2：饿了么；3：百度；4：口碑；
	OrderSequence string  `json:"order_sequence,omitempty"` // 取货序号 与order_source配合使用
	OrderTime     int64   `json:"order_time,omitempty"`     // 用户下单时间
	Version       int     `json:"version,omitempty"`        //文档版本号1.7,version=17
	Receive       Receive `json:"receive,omitempty"`        // 收货人信息
	OrderDetail   Detail  `json:"order_detail,omitempty"`   // 订单详情 必须
}

// 店铺信息
type Shop struct {
	ShopName    string `json:"shop_name"`
	ShopPhone   string `json:"shop_phone"`
	ShopAddress string `json:"shop_address"`
	ShopLng     string `json:"shop_lng"`
	ShopLat     string `json:"shop_lat"`
}

// 收货人信息
type Receive struct {
	UserName    string `json:"user_name"`
	UserPhone   string `json:"user_phone"`
	UserAddress string `json:"user_address"`
	UserLng     string `json:"user_lng"`
	UserLat     string `json:"user_lat"`
	CityName    string `json:"city_name"`
}

// 菜品
type Detail struct {
	TotalPrice     int          `json:"total_price"`          // 订单商品总金额 分
	ProductType    int          `json:"product_type"`         // 物品类型
	UserMoney      int          `json:"user_money,omitempty"` //用户实付商家金额（单位：分）
	ShopMoney      int          `json:"shop_money,omitempty"` //商家实收
	WeightGram     int          `json:"weight_gram"`          //物品重量（单位：克）
	ProductNum     int          `json:"product_num"`          // 物品个数
	ProductTypeNum int          `json:"product_type_num"`     // 物品种类个数
	ProductDetail  []DetailInfo `json:"product_detail"`
}

type DetailInfo struct {
	ProductName  string `json:"product_name"` // 物品名称
	ProductNum   int    `json:"product_num"`
	ProductPrice int    `json:"product_price"` //物品价格
	ProductUnit  string `json:"product_unit"`  // 物品单位
}

type SendOrderDeliveryResponse struct {
	SfOrderId             string      `json:"sf_order_id"`
	SfBillId              string      `json:"sf_bill_id"`
	ShopOrderId           string      `json:"shop_order_id"`
	PushTime              int64       `json:"push_time"`
	TotalPrice            int         `json:"total_price"`
	DeliveryDistanceMeter interface{} `json:"delivery_distance_meter"`
	WeightGram            int         `json:"weight_gram"`
	StartTime             int64       `json:"start_time"`
	ExpectTime            int64       `json:"expect_time"`
	TotalPayMoney         int         `json:"total_pay_money"` //支付费用
	RealPayMoney          int         `json:"real_pay_money"`
	CouponsTotalFee       int         `json:"coupons_total_fee"`
	SettlementType        int         `json:"settlement_type"`
}

/*
300	计划有变，暂时不需要寄件了
302	填错订单信息，取消后重新提交
303	骑士要求取消
304	暂时无法提供待配送物品
306	重复下单，取消此单
309	骑士上门时间太长
312	无人接单，换用其他平台寄件
313	其他，请注明原因
*/
type CancelOrderDeliveryRequestParams struct {
	DevId        int    `json:"dev_id"`
	OrderId      string `json:"order_id"`                //订单ID
	OrderType    int    `json:"order_type,omitempty"`    //查询订单ID类型	1、顺丰订单号 2、商家订单号
	ShopId       string `json:"shop_id,omitempty"`       //店铺ID 城市大区店铺
	ShopType     int    `json:"shop_type,omitempty"`     //店铺ID类型 1:顺丰店铺ID 2:接入方店铺ID
	CancelCode   int    `json:"cancel_code,omitempty"`   //不填时默认cancel_code=313,cancel_reason=商家发起取消
	CancelReason string `json:"cancel_reason,omitempty"` //其他取消原因
	PushTime     int64  `json:"push_time"`               //取消时间；秒级时间戳
}

type CancelOrderDeliveryResponse struct {
	SfOrderId       string `json:"sf_order_id"`
	ShopOrderId     string `json:"shop_order_id"`
	DeductionDetail struct {
		DeductionFee    int `json:"deduction_fee"`     //取消收费金额（单位：分）
		ShopCancelTimes int `json:"shop_cancel_times"` //店铺维度累计的取消次数
		FreeCancelTimes int `json:"free_cancel_times"` //配置的免费取消次数
	} `json:"deduction_detail"`
	PushTime int64 `json:"push_time"`
}

// SendOrderDelivery 订单发顺丰配送
func SendOrderDelivery(req OrderDeliveryRequest) (resp SendOrderDeliveryResponse, err error) {
	order := &SendOrderDeliveryRequestParams{
		DevId:         getInstance().devId,
		ShopId:        req.ShopId,
		ShopType:      1,
		LbsType:       2,
		GratuityFee:   int(req.Tips * 100),
		PushTime:      time.Now().Unix(),
		ShopOrderId:   req.OrderId,
		OrderSource:   "", //美团/饿了么
		OrderSequence: "", // seq 订单流水号
		OrderTime:     time.Now().Unix(),
		Version:       17,
		ReturnFlag:    511,
		Remark:        req.Remarks,
	}

	// 收货人信息
	order.Receive = Receive{
		UserName:    req.ReceiverName,
		UserPhone:   req.ReceiverPhone,
		UserLng:     req.ReceiverLng,
		UserLat:     req.ReceiverLat,
		UserAddress: req.ReceiverAddress,
		CityName:    req.CityName,
	}

	// 平台发单时候，需要传入门店信息
	if req.IsJuhe {
		shop := Shop{
			ShopName:    req.FromShopName,
			ShopAddress: req.FromAddress,
			ShopPhone:   req.FromPhone,
			ShopLng:     req.FromLng,
			ShopLat:     req.FromLat,
		}
		order.Shop = &shop
	}

	// 订单详情
	order.OrderDetail = Detail{
		TotalPrice:  int(utils.ToFloat64(req.OrderPrice) * 100),
		ProductType: 1,
		WeightGram:  1 * 1000,
	}

	sign := getSign(order)
	url := getUrl(api_createOrder, sign)

	body, err := json.Marshal(order)
	if err != nil {
		return
	}

	respBody, err := doPost(url, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(respBody, &resp)
	return
}

// CancelOrderDelivery 订单取消配送
func CancelOrderDelivery(req CancelOrderRequest) (resp CancelOrderDeliveryResponse, err error) {
	order := &CancelOrderDeliveryRequestParams{
		DevId:        getInstance().devId,
		OrderId:      req.SfOrderId, // 默认使用顺丰订单号
		OrderType:    1,
		ShopId:       req.ShopId,
		CancelCode:   313,
		CancelReason: req.CancelReason,
		PushTime:     time.Now().Unix(),
	}

	// order_type=2时必传shop_id与shop_type
	if order.OrderId == "" {
		order.OrderId = req.OrderId
		order.OrderType = 2 // 使用商家订单号
		order.ShopType = 1
	}

	sign := getSign(order)
	url := getUrl(api_cancelOrder, sign)

	body, err := json.Marshal(order)
	if err != nil {
		return
	}

	respBody, err := doPost(url, body)
	if err != nil {
		return
	}

	err = json.Unmarshal(respBody, &resp)
	return
}
