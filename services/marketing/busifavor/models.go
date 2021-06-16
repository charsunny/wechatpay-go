package busifavor

type BusiFavorStockRequest struct {
	StockName          string              `json:"stock_name"`
	BelongMerchant     string              `json:"belong_merchant"`
	Comment            string              `json:"comment,omitempty"`
	GoodsName          string              `json:"goods_name"`
	StockType          string              `json:"stock_type"`
	CouponUseRule      *CouponUseRule      `json:"coupon_use_rule"`
	StockSendRule      *StockSendRule      `json:"stock_send_rule"`
	OutRequestNo       string              `json:"out_request_no"`
	CustomEntrance     *CustomEntrance     `json:"custom_entrance,omitempty"`
	DisplayPatternInfo *DisplayPatternInfo `json:"display_pattern_info,omitempty"`
	CouponCodeMode     string              `json:"coupon_code_mode"`
	NotifyConfig       *NotifyConfig       `json:"notify_config"`
	Subsidy            *bool               `json:"subsidy,omitempty"`
}

type AvailableDayTime struct {
	BeginTime int `json:"begin_time"`
	EndTime   int `json:"end_time"`
}
type AvailableWeek struct {
	WeekDay          []int              `json:"week_day,omitempty"`
	AvailableDayTime []AvailableDayTime `json:"available_day_time,omitempty"`
}
type IrregularyAvaliableTime struct {
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
}
type CouponAvailableTime struct {
	AvailableBeginTime       string                     `json:"available_begin_time"`
	AvailableEndTime         string                     `json:"available_end_time"`
	AvailableDayAfterReceive int                        `json:"available_day_after_receive"`
	WaitDaysAfterReceive     int                        `json:"wait_days_after_receive"`
	AvailableWeek            *AvailableWeek             `json:"available_week,omitempty"`
	IrregularyAvaliableTime  []*IrregularyAvaliableTime `json:"irregulary_avaliable_time,omitempty"`
}
type FixedNormalCoupon struct {
	DiscountAmount     int `json:"discount_amount"`
	TransactionMinimum int `json:"transaction_minimum"`
}
type CouponUseRule struct {
	CouponAvailableTime CouponAvailableTime `json:"coupon_available_time"`
	FixedNormalCoupon   FixedNormalCoupon   `json:"fixed_normal_coupon"`
	UseMethod           string              `json:"use_method"`
	MiniProgramsAppid   string              `json:"mini_programs_appid"`
	MiniProgramsPath    string              `json:"mini_programs_path"`
}
type StockSendRule struct {
	MaxCoupons         int  `json:"max_coupons"`
	MaxCouponsPerUser  int  `json:"max_coupons_per_user"`
	MaxCouponsByDay    int  `json:"max_coupons_by_day,omitempty"`
	NaturalPersonLimit bool `json:"natural_person_limit"`
	PreventAPIAbuse    bool `json:"prevent_api_abuse"`
	Transferable       bool `json:"transferable"`
	Shareable          bool `json:"shareable"`
}
type MiniProgramsInfo struct {
	MiniProgramsAppid string `json:"mini_programs_appid"`
	MiniProgramsPath  string `json:"mini_programs_path"`
	EntranceWords     string `json:"entrance_words"`
	GuidingWords      string `json:"guiding_words,omitempty"`
}
type CustomEntrance struct {
	MiniProgramsInfo MiniProgramsInfo `json:"mini_programs_info,omitempty"`
	Appid            string           `json:"appid,omitempty"`
	HallID           string           `json:"hall_id,omitempty"`
	StoreID          string           `json:"store_id,omitempty"`
}
type DisplayPatternInfo struct {
	Description     string `json:"description,omitempty"`
	MerchantLogoURL string `json:"merchant_logo_url,omitempty"`
	MerchantName    string `json:"merchant_name,omitempty"`
	BackgroundColor string `json:"background_color,omitempty"`
	CouponImageURL  string `json:"coupon_image_url,omitempty"`
}

type CouponCodeCount struct {
	TotalCount     int `json:"total_count"`
	AvailableCount int `json:"available_count"`
}

type NotifyConfig struct {
	NotifyAppid string `json:"notify_appid,omitempty"`
}

type BusiFavorStockResponse struct {
	StockId    string `json:"stock_id"`
	CreateTime string `json:"create_time"`
}

type BusiFavorStockInfo struct {
	StockName          string             `json:"stock_name"`
	BelongMerchant     string             `json:"belong_merchant"`
	Comment            string             `json:"comment"`
	GoodsName          string             `json:"goods_name"`
	StockType          string             `json:"stock_type"`
	CouponUseRule      CouponUseRule      `json:"coupon_use_rule"`
	StockSendRule      StockSendRule      `json:"stock_send_rule"`
	CustomEntrance     CustomEntrance     `json:"custom_entrance"`
	DisplayPatternInfo DisplayPatternInfo `json:"display_pattern_info"`
	StockState         string             `json:"stock_state"`
	CouponCodeMode     string             `json:"coupon_code_mode"`
	StockID            string             `json:"stock_id"`
	NotifyConfig       *NotifyConfig      `json:"notify_config"`
	CouponCodeCount    CouponCodeCount    `json:"coupon_code_count"`
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

type CouponConsumeRequest struct {
	CouponCode   string `json:"coupon_code"`
	StockID      string `json:"stock_id"`
	Appid        string `json:"appid"`
	UseTime      string `json:"use_time"`
	UseRequestNo string `json:"use_request_no"`
	Openid       string `json:"openid"`
}

type CouponConsumeResponse struct {
	StockID          string `json:"stock_id"`
	Openid           string `json:"openid"`
	WechatpayUseTime string `json:"wechatpay_use_time"`
}
