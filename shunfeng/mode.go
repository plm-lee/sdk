package shunfeng

// QueryDeliveryFeeRequest 询价请求结构
type QueryDeliveryFeeRequest struct {
	ShopId          string  `json:"sf_shopId"`               // 顺丰门店ID
	OrderId         string  `json:"order_id"`                // 第三方订单ID
	IsJuhe          bool    `json:"is_juhe"`                 // 是否聚合发单，使用平台账号发单
	Tips            float64 `json:"tips"`                    // 小费（单位：元，精确小数点后一位）
	FromShopName    string  `json:"from_shopName,omitempty"` // 发单门店名 聚合发单传入
	FromAddress     string  `json:"from_address,omitempty"`  // 发单门店地址 聚合发单传入
	FromLat         string  `json:"from_lat,omitempty"`      // 发单地址纬度 聚合发单传入
	FromLng         string  `json:"from_lng,omitempty"`      // 发单地址经度 聚合发单传入
	ReceiverAddress string  `json:"receiver_address"`        // 收货人地址
	ReceiverLat     string  `json:"receiver_lat"`            // 收货人地址纬度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度
	ReceiverLng     string  `json:"receiver_lng"`            // 收货人地址经度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度
}

// QueryDeliveryFeeResult 询价返回结构
type QueryDeliveryFeeResult struct {
	OrderId     string  `json:"order_id"`     //订单id
	Distance    float64 `json:"distance"`     //配送距离(单位：米)
	DeliveryNo  string  `json:"deliveryNo"`   //平台订单号
	PayedPrice  float64 `json:"payed_price"`  //实际运费(单位：元)，运费减去优惠券费用
	OriginPrice float64 `json:"origin_price"` //运费(单位：元)
	Tips        float64 `json:"tips"`         //小费(单位：元)
}

// OrderDeliveryRequest 订单发配送请求结构
type OrderDeliveryRequest struct {
	OrderId         string  `json:"order_id"`                //第三方订单ID
	OrderPrice      string  `json:"order_price"`             // 订单价格
	ReceiverName    string  `json:"receiver_name"`           //收货人姓名
	ReceiverPhone   string  `json:"receiver_phone"`          //收货人手机号虚拟号码格式（手机号_分机号码）例如：13700000000_1111
	Remarks         string  `json:"remarks"`                 // 备注
	ShopId          string  `json:"sf_shopId"`               // 顺丰门店ID
	IsJuhe          bool    `json:"is_juhe"`                 // 是否聚合发单，使用平台账号发单
	Tips            float64 `json:"tips"`                    // 小费（单位：元，精确小数点后一位）
	FromShopName    string  `json:"from_shopName,omitempty"` // 发单门店名 聚合发单传入
	FromAddress     string  `json:"from_address,omitempty"`  // 发单门店地址 聚合发单传入
	FromLat         string  `json:"from_lat,omitempty"`      // 发单地址纬度 聚合发单传入
	FromLng         string  `json:"from_lng,omitempty"`      // 发单地址经度 聚合发单传入
	FromPhone       string  `json:"from_phone"`              // 发单人联系号码
	ReceiverAddress string  `json:"receiver_address"`        // 收货人地址
	ReceiverLat     string  `json:"receiver_lat"`            // 收货人地址纬度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度
	ReceiverLng     string  `json:"receiver_lng"`            // 收货人地址经度（高德坐标系，若是其他地图经纬度需要转化成高德地图经纬度
	CityName        string  `json:"city_name"`               //订单所在城市
}

// CancelOrderRequest 订单取消请求结构
type CancelOrderRequest struct {
	SfOrderId    string `json:"sf_orderId"`
	OrderId      string `json:"order_id"`
	ShopId       string `json:"sf_shopId"`
	CancelReason string `json:"cancel_reason"` //
}

// AddTipRequest 订单加小费请求结构
type AddTipRequest struct {
	SfOrderId string  `json:"sf_orderId"`
	OrderId   string  `json:"order_id"` // 第三方订单编号
	Tips      float64 `json:"tips"`     // 消费金额（精确到小数点后一位，元）
	Remarks   string  `json:"remarks"`  // 备注
}
