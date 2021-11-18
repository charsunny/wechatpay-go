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

// FavorStockInfo 代金券批次详情
type FavorStockInfo struct {
	StockID           string `json:"stock_id"`
	StockCreatorMchid string `json:"stock_creator_mchid"`
	StockName         string `json:"stock_name"`
	Status            string `json:"status"`
	CreateTime        string `json:"create_time"`
	Description       string `json:"description"`
	StockUseRule      struct {
		MaxCoupons        int      `json:"max_coupons"`
		MaxAmount         int      `json:"max_amount"`
		MaxAmountByDay    int      `json:"max_amount_by_day"`
		MaxCouponsPerUser int      `json:"max_coupons_per_user"`
		TradeType         []string `json:"trade_type"`
	} `json:"stock_use_rule,omitempty"`
	AvailableBeginTime string `json:"available_begin_time"`
	AvailableEndTime   string `json:"available_end_time"`
	DistributedCoupons int    `json:"distributed_coupons"`
	NoCash             bool   `json:"no_cash"`
	CutToMessage       struct {
		SinglePriceMax int `json:"single_price_max"`
		CutToPrice     int `json:"cut_to_price"`
	} `json:"cut_to_message,omitempty"`
	Singleitem bool   `json:"singleitem"`
	StockType  string `json:"stock_type"`
}

// FavorStockListRequest 条件查询批次列表请求参数
type FavorStockListRequest struct {
	StockCreatorMchid string `json:"stock_creator_mchid"`         // 批次创建方商户号。
	CreateStartTime   string `json:"create_start_time,omitempty"` //起始创建时间
	CreateEndTime     string `json:"create_end_time,omitempty"`   // 终止创建时间
	Status            string `json:"status,omitempty"`            // 批次状态，枚举值：unactivated：未激活 audit：审核中 running：运行中 stoped：已停止 paused：暂停发放
	Offset            int    `json:"offset,omitempty"`            // 分页页码
	Limit             int    `json:"limit,omitempty"`             // 分页大小
}

// FavorStockList 代金券批次列表
type FavorStockList struct {
	List       []FavorStockInfo `json:"data,omitempty"`
	TotalCount int              `json:"total_count"`
	Limit      int              `json:"limit"`
	Offset     int              `json:"offset"`
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

// Callbacks 消息通知地址
type Callbacks struct {
	MchID      string `json:"mchid"`                 // 微信支付商户的商户号
	NotifyURL  string `json:"notify_url"`            //  商户提供的用于接收商家券事件通知的url地址，必须支持https。2、地址要是一个可访问的地址3、不能携带参数。
	UpdateTime string `json:"update_time,omitempty"` // 修改时间
	Switch     bool   `json:"switch,omitempty"`      //如果商户不需要再接收营销事件通知，可通过该开关关闭。枚举值：true：开启推送false：停止推送校验规则：1、该参数设置为false暂不支持使用 2、如果想关闭通知可以notify_url置空请求即可，或者直接在商户平台关闭券回调通知的
}

// UserFavor 用户领取的代金券
type UserFavor struct {
	StockCreatorMchid string `json:"stock_creator_mchid"`
	StockID           string `json:"stock_id"`
	CouponID          string `json:"coupon_id"`
	CutToMessage      struct {
		SinglePriceMax int `json:"single_price_max"`
		CutToPrice     int `json:"cut_to_price"`
	} `json:"cut_to_message"`
	CouponName              string `json:"coupon_name"`
	Status                  string `json:"status"`
	Description             string `json:"description"`
	CreateTime              string `json:"create_time"`
	CouponType              string `json:"coupon_type"`
	NoCash                  bool   `json:"no_cash"`
	AvailableBeginTime      string `json:"available_begin_time"`
	AvailableEndTime        string `json:"available_end_time"`
	Singleitem              bool   `json:"singleitem"`
	NormalCouponInformation struct {
		CouponAmount       int `json:"coupon_amount"`
		TransactionMinimum int `json:"transaction_minimum"`
	} `json:"normal_coupon_information"`
}

// FavorMerchantsOrItems 代金券可用商户/单品列表请求参数
type FavorMerchantsOrItemsRequest struct {
	StockID           string `json:"stock_id"`            //  批次号
	StockCreatorMchid string `json:"stock_creator_mchid"` // 批次创建方商户号。
	Offset            int    `json:"offset,omitempty"`    // 分页页码
	Limit             int    `json:"limit,omitempty"`     // 分页大小
}

// FavorMerchantsOrItems  代金券可用商户/单品列表
type FavorMerchantsOrItems struct {
	Data       []string `json:"data,omitempty"`
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	StockID    string   `json:"stock_id"`
	TotalCount int      `json:"total_count"`
}

// UserFavosrRequest 根据商户号查用户的券请求参数
type UserFavorsRequest struct {
	Appid             string `json:"appid"`                         //微信为发券方商户分配的公众账号ID，接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的）或APP的appid（在open.weixin.qq.com申请的）。
	StockID           string `json:"stock_id,omitempty"`            //  批次号
	StockCreatorMchid string `json:"stock_creator_mchid,omitempty"` // 批次创建方商户号。三选一
	SenderMchid       string `json:"sender_mchid,omitempty"`        // 批次发放商户号 （该字段暂未开放）三选一
	AvailableMchid    string `json:"available_mchid,omitempty"`     //可用商户号 三选一
	Status            string `json:"status,omitempty"`
	Offset            int    `json:"offset,omitempty"` // 分页页码
	Limit             int    `json:"limit,omitempty"`  // 分页大小
}

// UserFavors  根据商户号查用户的券列表
type UserFavors struct {
	Data       []string `json:"data,omitempty"`
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	StockID    string   `json:"stock_id"`
	TotalCount int      `json:"total_count"`
}

type Flow struct {
	HashType  string `json:"hash_type"`
	HashValue string `json:"hash_value"`
	URL       string `json:"url"`
}
