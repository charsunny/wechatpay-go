package transactions

import (
	"time"
)

// Amount
type Amount struct {
	// 订单总金额，单位为分
	Total int `json:"total"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency string `json:"currency,omitempty"`
}

// CloseOrderRequest
type CloseOrderRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	// 直连商户号
	Mchid string `json:"mchid"`
}

// CloseRequest
type CloseRequest struct {
	// 直连商户号
	Mchid string `json:"mchid"`
}

// Detail 优惠功能
type Detail struct {
	// 1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。 2.当订单原价与支付金额不相等，则不享受优惠。 3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	CostPrice *int32 `json:"cost_price,omitempty"`
	// 商家小票ID。
	InvoiceId   string         `json:"invoice_id,omitempty"`
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"`
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。
	MerchantGoodsId string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）。
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称。
	GoodsName string `json:"goods_name,omitempty"`
	// 用户购买的数量。
	Quantity int32 `json:"quantity"`
	// 商品单价，单位为分。
	UnitPrice int32 `json:"unit_price"`
}

// H5Info
type H5Info struct {
	// 场景类型 iOS, Android, Wap
	Type string `json:"type"`
	// 应用名称
	AppName string `json:"app_name,omitempty"`
	// 网站URL
	AppUrl string `json:"app_url,omitempty"`
	// iOS平台BundleID
	BundleId string `json:"bundle_id,omitempty"`
	// Android平台PackageName
	PackageName string `json:"package_name,omitempty"`
}

// PrepayRequest
type PrepayRequest struct {
	// 服务商传递
	SpMchid  string `json:"sp_mchid,omitempty"`
	SubMchid string `json:"sub_mchid,omitempty"`
	SpAppid  string `json:"sp_appid,omitempty"`
	SubAppid string `json:"sub_appid,omitempty"`
	// 直营商户公众号ID
	Appid string `json:"appid,omitempty"`
	Mchid string `json:"mchid,omitempty"`
	// 商品描述
	Description *string `json:"description"`
	// 商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	// 订单生成时间，格式为rfc3339格式
	TimeExpire *time.Time `json:"time_expire,omitempty"`
	// 附加数据
	Attach string `json:"attach,omitempty"`
	// 有效性：1. HTTPS；2. 不允许携带查询串。
	NotifyUrl string `json:"notify_url"`
	// 商品标记，代金券或立减优惠功能的参数。
	GoodsTag string `json:"goods_tag,omitempty"`
	// 指定支付方式
	LimitPay []string `json:"limit_pay,omitempty"`
	// 传入true时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效。
	SupportFapiao *bool       `json:"support_fapiao,omitempty"`
	Payer         *Payer      `json:"payer,omitempty"`
	Amount        *Amount     `json:"amount"`
	Detail        *Detail     `json:"detail,omitempty"`
	SceneInfo     *SceneInfo  `json:"scene_info"`
	SettleInfo    *SettleInfo `json:"settle_info,omitempty"`
}

// PrepayResponse
type PrepayResponse struct {
	// 支付跳转链接
	H5Url *string `json:"h5_url"`
}

// QueryOrderByIdRequest
type QueryOrderByIdRequest struct {
	TransactionId *string `json:"transaction_id"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

// QueryOrderByOutTradeNoRequest
type QueryOrderByOutTradeNoRequest struct {
	OutTradeNo *string `json:"out_trade_no"`
	// 直连商户号
	Mchid *string `json:"mchid"`
}

// SceneInfo 支付场景描述
type SceneInfo struct {
	// 用户终端IP
	PayerClientIp string `json:"payer_client_ip"`
	// 商户端设备号
	DeviceId  string     `json:"device_id,omitempty"`
	StoreInfo *StoreInfo `json:"store_info,omitempty"`
	// h5下单为必须传字段
	H5Info *H5Info `json:"h5_info,omitempty"`
}

// SettleInfo
type SettleInfo struct {
	// 是否指定分账
	ProfitSharing *bool `json:"profit_sharing,omitempty"`
}

// StoreInfo 商户门店信息
type StoreInfo struct {
	// 商户侧门店编号
	Id string `json:"id"`
	// 商户侧门店名称
	Name string `json:"name,omitempty"`
	// 地区编码，详细请见微信支付提供的文档
	AreaCode string `json:"area_code,omitempty"`
	// 详细的商户门店地址
	Address string `json:"address,omitempty"`
}

// PromotionDetail
type PromotionDetail struct {
	// 券ID
	CouponId *string `json:"coupon_id,omitempty"`
	// 优惠名称
	Name *string `json:"name,omitempty"`
	// GLOBAL：全场代金券；SINGLE：单品优惠
	Scope *string `json:"scope,omitempty"`
	// CASH：充值；NOCASH：预充值。
	Type string `json:"type,omitempty"`
	// 优惠券面额
	Amount int32 `json:"amount,omitempty"`
	// 活动ID，批次ID
	StockId *string `json:"stock_id,omitempty"`
	// 单位为分
	WechatpayContribute *int32 `json:"wechatpay_contribute,omitempty"`
	// 单位为分
	MerchantContribute *int32 `json:"merchant_contribute,omitempty"`
	// 单位为分
	OtherContribute *int32 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency    *string                `json:"currency,omitempty"`
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

// PromotionGoodsDetail
type PromotionGoodsDetail struct {
	// 商品编码
	GoodsId *string `json:"goods_id"`
	// 商品数量
	Quantity *int32 `json:"quantity"`
	// 商品价格
	UnitPrice *int32 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount *int32 `json:"discount_amount"`
	// 商品备注
	GoodsRemark *string `json:"goods_remark,omitempty"`
}

// Transaction
type Transaction struct {
	// 服务商返回
	SpMchid  string `json:"sp_mchid,omitempty"`
	SubMchid string `json:"sub_mchid,omitempty"`
	SpAppid  string `json:"sp_appid,omitempty"`
	SubAppid string `json:"sub_appid,omitempty"`
	// 正常返回
	Amount          *TransactionAmount `json:"amount,omitempty"`
	Appid           *string            `json:"appid,omitempty"`
	Attach          *string            `json:"attach,omitempty"`
	BankType        *string            `json:"bank_type,omitempty"`
	Mchid           *string            `json:"mchid,omitempty"`
	OutTradeNo      *string            `json:"out_trade_no,omitempty"`
	Payer           *Payer             `json:"payer,omitempty"`
	PromotionDetail []PromotionDetail  `json:"promotion_detail,omitempty"`
	SuccessTime     *string            `json:"success_time,omitempty"`
	TradeState      *string            `json:"trade_state,omitempty"`
	TradeStateDesc  *string            `json:"trade_state_desc,omitempty"`
	TradeType       *string            `json:"trade_type,omitempty"`
	TransactionId   *string            `json:"transaction_id,omitempty"`
}

// TransactionAmount
type TransactionAmount struct {
	Currency      *string `json:"currency,omitempty"`
	PayerCurrency *string `json:"payer_currency,omitempty"`
	PayerTotal    *int32  `json:"payer_total,omitempty"`
	Total         *int32  `json:"total,omitempty"`
}

// TransactionPayer
type Payer struct {
	// 服务商使用
	SpOpenid  string `json:"sp_openid,omitempty"`
	SubOpenid string `json:"sub_openid,omitempty"`
	// 商户使用
	Openid string `json:"openid,omitempty"`
}
