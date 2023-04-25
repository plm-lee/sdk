package shunfeng

import (
	"encoding/json"
	"time"
)

type OrderQueryFeeRequestParams struct {
	DevId            int         `json:"dev_id"`                       //api开发者ID
	ShopId           string      `json:"shop_id"`                      //店铺ID 城市大区店铺
	ShopType         int         `json:"shop_type,omitempty"`          //店铺ID类型 1:顺丰店铺ID 2:接入方店铺ID
	UserLng          string      `json:"user_lng,omitempty"`           //用户地址经度
	UserLat          string      `json:"user_lat,omitempty"`           //用户地址纬度
	UserAddress      string      `json:"user_address"`                 //用户详细地址
	CityName         string      `json:"city_name,omitempty"`          //发单城市
	Weight           int         `json:"weight"`                       //物品重量（单位：克）
	ProductType      int         `json:"product_type"`                 //物品类型
	TotalPrice       int         `json:"total_price,omitempty"`        //用户订单总金额（单位：分）
	IsAppoint        int         `json:"is_appoint"`                   //是否是预约单	0：非预约单；1：预约单
	AppointType      int         `json:"appoint_type,omitempty"`       //预约单类型	 预约单的时候传入，1：预约单送达单；2：预约单上门单
	ExpectTime       int         `json:"expect_time,omitempty"`        //用户期望送达时间 若传入自此段且时间大于配送时效，则按照预约送达单处理，时间小于配送时效按照立即单处理；appoint_type=1时需必传,秒级时间戳
	ExpectPickupTime int         `json:"expect_pickup_time,omitempty"` //用户期望上门时间 appoint_type=2时需必传,秒级时间戳
	LbsType          int         `json:"lbs_type,omitempty"`           //坐标类型，1：百度坐标，2：高德坐标
	PayType          int         `json:"pay_type"`                     //用户支付方式：1、已支付 0、货到付款
	ReceiveUserMoney int         `json:"receive_user_money,omitempty"` //代收金额
	IsInsured        int         `json:"is_insured"`                   //是否保价，0：非保价；1：保价
	IsPersonDirect   int         `json:"is_person_direct"`             //是否是专人直送订单，0：否；1：是
	Vehicle          int         `json:"vehicle,omitempty"`            //配送交通工具，0：否；1：电动车；2：小轿车
	DeclaredValue    int         `json:"declared_value,omitempty"`     //保价金额	单位：分
	GratuityFee      int         `json:"gratuity_fee,omitempty"`       //订单小费，不传或者传0为不加小费 单位分，加小费最低不能少于100分
	RiderPickMethod  int         `json:"rider_pick_method,omitempty"`  //物流流向 1：从门店取件送至用户； 2：从用户取件送至门店
	ReturnFlag       int         `json:"return_flag,omitempty"`        //返回字段控制标志位（二进制） 1:商品总价格，2:配送距离，4:物品重量，8:起送时间，16:期望送达时间，32:支付费用，64:实际支持金额，128:优惠卷总金额，256:结算方式 例如全部返回为填入511
	PushTime         int64       `json:"push_time"`                    //推单时间
	Shop             *ShopParams `json:"shop,omitempty"`               //发货店铺信息 平台级开发者需要传入
}

// ShopParams 店铺信息
type ShopParams struct {
	ShopName    string `json:"shop_name"`
	ShopPhone   string `json:"shop_phone"`
	ShopAddress string `json:"shop_address"`
	ShopLng     string `json:"shop_lng"`
	ShopLat     string `json:"shop_lat"`
}

// OrderQueryFeeResponse 查询配送费返回
type OrderQueryFeeResponse struct {
	DeliveryType          int         `json:"delivery_type"`           //0：预约送达单 1：立即单 3：预约上门单
	ExpectTime            int64       `json:"expect_time"`             ////预计送达（上门）时间
	StartTime             int64       `json:"start_time"`              ////预计开始配送的时间
	PromiseDeliveryTime   int         `json:"promise_delivery_time"`   ////预计配送时间（单位: 分）
	DeliveryDistanceMeter interface{} `json:"delivery_distance_meter"` ////配送距离
	ChargePriceList       struct {
		ShopPayPrice  int `json:"shop_pay_price"` //配送费总额（单位:分）
		ChargesDetail struct {
			BasicFee        int `json:"basic_fee"`         ////常规配送费=起步价+超距离费+超重量费
			Basic           int `json:"basic"`             //起步价
			OverDistance    int `json:"over_distance"`     //超距离费用
			OverWeight      int `json:"over_weight"`       //超重量费用
			CancelExcessFee int `json:"cancel_excess_fee"` //拒收扣费
			SpecialTimeFee  int `json:"special_time_fee"`  //特殊时段费
			VasFee          int `json:"vas_fee"`           //增值服务费
			VasFeeDetail    struct {
				PackingFee      int `json:"packing_fee"`        //包材费
				LowTempFee      int `json:"low_temp_fee"`       //低温服务费
				TakeGoodsSmsFee int `json:"take_goods_sms_fee"` //取货短信费
				Insured         struct {
					Fee           int `json:"fee"`
					DeclaredPrice int `json:"declared_price"`
				} `json:"insured"` //增值服务费详情
				ExtraFee int `json:"extra_fee"`
			}
		}
	}
	GratuityFee int   `json:"gratuity_fee"` //订单小费
	PushPime    int64 `json:"push_time"`
	//以下字段受请求参数中 return_flag 控制：return_flag中未包含的，此字段将不存在，请注意！
	TotalPrice     int `json:"total_price"`     //配送费总额，当return_flag中包含1时返回，单位分（值为计算出来此单总价）
	WeightGram     int `json:"weight_gram"`     //商品重量，当return_flag中包含4时返回，单位克（值为下单传入参数回传）
	TotalPayMoney  int `json:"total_pay_money"` //支付费用，当return_flag中包含32时返回，单位分
	RealPayMoney   int `json:"real_pay_money"`  //实际支付金额，当return_flag中包含64时返回，单位分（实际支付金额=总金额-优惠卷总金额）
	SettlementType int `json:"settlement_type"` //结算方式，当return_flag中包含256时返回
}

// GetOrderPrice 获取订单配送费用
func GetOrderPrice(req QueryDeliveryFeeRequest) (resp OrderQueryFeeResponse, err error) {
	order := &OrderQueryFeeRequestParams{
		DevId:       getInstance().devId,
		ShopId:      req.ShopId,
		ShopType:    1, // 1: 使用顺丰店铺ID 2: 使用系统内部店铺id
		UserLng:     req.ReceiverLng,
		UserLat:     req.ReceiverLat,
		UserAddress: req.ReceiverAddress,
		Weight:      1000,
		ProductType: 1, // 快餐
		LbsType:     2, // 默认使用高德坐标系
		GratuityFee: int(req.Tips * 100),
		PushTime:    time.Now().Unix(),
		ReturnFlag:  511,
	}

	// 平台发单，需要传入发单门店信息
	if req.IsJuhe {
		shop := ShopParams{
			ShopName:    req.FromShopName,
			ShopAddress: req.FromAddress,
			ShopPhone:   req.ShopId,
			ShopLng:     req.FromLng,
			ShopLat:     req.FromLat,
		}
		order.Shop = &shop
	}
	sign := getSign(order)
	url := getUrl(api_queryDeliverFee, sign)

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
