package shunfeng

// MerchantAuthCallbackMessage 商家授权回调消息结构
type MerchantAuthCallbackMessage struct {
	UrlIndex  string `json:"url_index"`
	ShopId    int    `json:"shop_id"`
	OutShopId string `json:"out_shop_id"`
	PushTime  int    `json:"push_time"`
}
