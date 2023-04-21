package uu

// DeliveryNotifyMessage 订单配送状态回调参数
type DeliveryNotifyMessage struct {
	OrderCode    string `json:"order_code"`    // 订单号
	DriverName   string `json:"driver_name"`   // 跑男姓名(跑男接单后)
	DriverJobnum string `json:"driver_jobnum"` // 跑男工号(跑男接单后)
	DriverMobile string `json:"driver_mobile"` // 跑男电话(跑男接单后)
	State        string `json:"state"`         // 当前状态1下单成功 3跑男抢单 4已到达 5已取件 6到达目的地 10收件人已收货 -1订单取消
	StateText    string `json:"state_text"`    // 当前状态说明
	OriginId     string `json:"origin_id"`     // 第三方订单号
	DriverPhoto  string `json:"driver_photo"`  // 跑男头像(跑男接单后)
	BasicResponseParams
}
