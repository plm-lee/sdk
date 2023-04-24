package shunfeng

// DeliveryNotifyMessage 订单配送状态回调消息
type DeliveryNotifyMessage struct {
	ShopId        interface{} `json:"shop_id"`        // 店铺ID
	SfOrderId     string      `json:"sf_order_id"`    // 顺丰订单ID
	ShopOrderId   string      `json:"shop_order_id"`  // 商家订单id
	UrlIndex      string      `json:"url_index"`      // rider_status 状态更改 order_complete 完成 sf_cancel 取消
	OperatorName  string      `json:"operator_name"`  // 配送员姓名
	OperatorPhone string      `json:"operator_phone"` // 配送员电话
	RiderLng      interface{} `json:"rider_lng"`      // 配送员经度
	RiderLat      interface{} `json:"rider_lat"`      // 纬度
	OrderStatus   int         `json:"order_status"`   // 订单描述 10-配送员确认;12:配送员到店;15:配送员配送中
	StatusDesc    string      `json:"status_desc"`    // 状态描述
	PushTime      int         `json:"push_time"`      // 状态变更时间
	PickupPic     []string    `json:"pickup_pic"`     // 妥投照片
	CancelReason  string      `json:"cancel_reason"`  // 取消原因

	ExId       int    `json:"ex_id"`       //异常ID
	ExContent  string `json:"ex_content"`  //异常详情
	ExpectTime string `json:"expect_time"` //新的期望送达时间
}
