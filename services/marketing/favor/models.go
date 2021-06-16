package favor

type CouponStockRequest struct {
	StockName          string        `json:"stock_name"`
	Comment            string        `json:"comment,omitempty"`
	BelongMerchant     string        `json:"belong_merchant"`
	AvailableBeginTime string        `json:"available_begin_time"`
	AvailableEndTime   string        `json:"available_end_time"`
	StockUseRule       StockUseRule  `json:"stock_use_rule"`
	PatternInfo        PatternInfo   `json:"pattern_info"`
	CouponUseRule      CouponUseRule `json:"coupon_use_rule,omitempty"`
	NoCash             bool          `json:"no_cash"`
	StockType          string        `json:"stock_type,omitempty"`
	OutRequestNo       string        `json:"out_request_no"`
	ExtInfo            string        `json:"ext_info,omitempty"`
}

type StockUseRule struct {
	MaxCoupons         uint64  `json:"max_coupons"`
	MaxAmount          uint64  `json:"max_amount"`
	MaxAmountByDay     *uint64 `json:"max_amount_by_day,omitempty"`
	MaxCouponsPerUser  uint32  `json:"max_coupons_per_user"`
	NaturalPersonLimit bool    `json:"natural_person_limit"`
	PreventApiAbuse    bool    `json:"prevent_api_abuse"`
}

type PatternInfo struct {
	Description     string `json:"description"`
	MerchantLogo    string `json:"merchant_logo,omitempty"`
	MerchantName    string `json:"merchant_name,omitempty"`
	BackgroundColor string `json:"background_color,omitempty"`
	CouponImage     string `json:"coupon_image,omitempty"`
}

type CouponUseRule struct {
	AvailableItems     []string          `json:"available_items,omitempty"`
	AvailableMerchants []string          `json:"available_merchants,omitempty"`
	CombineUse         bool              `json:"combine_use,omitempty"`
	FixedNormalCoupon  FixedNormalCoupon `json:"fixed_normal_coupon,omitempty"`
	GoodsTag           []string          `json:"goods_tag,omitempty"`
	TradeType          []string          `json:"trade_type,omitempty"`
}

type FixedNormalCoupon struct {
	CouponAmount       int64 `json:"coupon_amount"`
	TransactionMinimum int64 `json:"transaction_minimum"`
}

type CouponStockResponse struct {
	StockId    string `json:"stock_id"`
	CreateTime string `json:"create_time"`
}

type CouponSendRequest struct {
	StockID           string  `json:"stock_id"`
	OutRequestNo      string  `json:"out_request_no"`
	Appid             string  `json:"appid"`
	StockCreatorMchid string  `json:"stock_creator_mchid"`
	CouponValue       *uint64 `json:"coupon_value,omitempty"`
	CouponMinimum     *uint64 `json:"coupon_minimum,omitempty"`
}

type CouponSendResponse struct {
	CouponId string `json:"coupon_id"`
}
