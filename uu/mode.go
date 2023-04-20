package uu

// QueryDeliveryFeeRequest  询价接口传入参数
type QueryDeliveryFeeRequest struct {
	OrderId         string  `json:"order_id"`    //第三方订单ID
	UserOpenId      string  `json:"user_openId"` // 商家自有账号id
	CityName        string  `json:"city_name"`   //订单所在城市
	OrderPrice      float64 `json:"order_price"` //订单金额（单位：元）
	FromAddress     string  `json:"from_address"`
	FromLat         string  `json:"from_lat"`         // 发单地址纬度
	FromLng         string  `json:"from_lng"`         // 发单地址经度
	ReceiverName    string  `json:"receiver_name"`    //收货人姓名
	ReceiverAddress string  `json:"receiver_address"` //收货人地址
	ReceiverLat     string  `json:"receiver_lat"`     //收货人地址纬度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度
	ReceiverLng     string  `json:"receiver_lng"`     //收货人地址经度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度
	ReceiverPhone   string  `json:"receiver_phone"`   //收货人手机号（手机号和座机号必填一项）
	Tips            float64 `json:"tips"`             //小费（单位：元，精确小数点后一位）
	Remarks         string  `json:"remarks"`          //订单备注
	Quantity        int     `json:"quantity"`         //订单商品数量
}

type OrderDeliveryRequest struct {
	PriceToken    string `json:"price_token"`
	OriginPrice   string `json:"origin_price"`
	PayedPrice    string `json:"payed_price"`
	ReceiverName  string `json:"receiver_name"`  //收货人姓名
	ReceiverPhone string `json:"receiver_phone"` //收货人手机号虚拟号码格式（手机号_分机号码）例如：13700000000_1111
	Remarks       string `json:"remarks"`        // 备注
	UserOpenId    string `json:"user_openId"`    // 发单用户
}

// 订单取消请求结构
type CancelOrderRequest struct {
	OrderId      string `json:"order_id"`
	UserOpenId   string `json:"user_openId"`
	CancelCode   int    `json:"cancel_code"`   // 取消原因
	CancelReason string `json:"cancel_reason"` //
}
