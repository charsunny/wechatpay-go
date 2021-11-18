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
	NotifyConfig       *NotifyConfig       `json:"notify_config,omitempty"`
	Subsidy            *bool               `json:"subsidy,omitempty"`
}

type ModifyBusiFavorStockRequest struct {
	CustomEntrance struct {
		MiniProgramsInfo struct {
			MiniProgramsAppid string `json:"mini_programs_appid,omitempty"`
			MiniProgramsPath  string `json:"mini_programs_path,omitempty"`
			EntranceWords     string `json:"entrance_words,omitempty"`
			GuidingWords      string `json:"guiding_words,omitempty"`
		} `json:"mini_programs_info,omitempty"`
		Appid           string `json:"appid,omitempty"`
		HallID          string `json:"hall_id,omitempty"`
		CodeDisplayMode string `json:"code_display_mode,omitempty"`
	} `json:"custom_entrance,omitempty"`
	Comment            string `json:"comment,omitempty"`
	GoodsName          string `json:"goods_name,omitempty"`
	OutRequestNo       string `json:"out_request_no"`
	DisplayPatternInfo struct {
		Description     string `json:"description,omitempty"`
		BackgroundColor string `json:"background_color,omitempty"`
		CouponImageURL  string `json:"coupon_image_url,omitempty"`
		FinderInfo      struct {
			FinderID                 string `json:"finder_id"`
			FinderVideoCoverImageURL string `json:"finder_video_cover_image_url"`
			FinderVideoID            string `json:"finder_video_id"`
		} `json:"finder_info,omitempty"`
	} `json:"display_pattern_info,omitempty"`
	CouponUseRule struct {
		UseMethod         string `json:"use_method,omitempty"`
		MiniProgramsAppid string `json:"mini_programs_appid,omitempty"`
		MiniProgramsPath  string `json:"mini_programs_path,omitempty"`
	} `json:"coupon_use_rule,omitempty"`
	StockSendRule struct {
		PreventAPIAbuse bool `json:"prevent_api_abuse,omitempty"`
	} `json:"stock_send_rule,omitempty"`
	NotifyConfig struct {
		NotifyAppid string `json:"notify_appid,omitempty"`
	} `json:"notify_config,omitempty"`
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
	AvailableDayAfterReceive int                        `json:"available_day_after_receive,omitempty"`
	WaitDaysAfterReceive     int                        `json:"wait_days_after_receive,omitempty"`
	AvailableWeek            *AvailableWeek             `json:"available_week,omitempty"`
	IrregularyAvaliableTime  []*IrregularyAvaliableTime `json:"irregulary_avaliable_time,omitempty"`
}
type FixedNormalCoupon struct {
	DiscountAmount     int `json:"discount_amount"`
	TransactionMinimum int `json:"transaction_minimum"`
}
type DiscountCoupon struct {
	DiscountPercent    int `json:"discount_percent"`
	TransactionMinimum int `json:"transaction_minimum"`
}
type ExchangeCoupon struct {
	ExchangePrice      int `json:"exchange_price"`
	TransactionMinimum int `json:"transaction_minimum"`
}
type CouponUseRule struct {
	CouponAvailableTime CouponAvailableTime `json:"coupon_available_time"`
	FixedNormalCoupon   FixedNormalCoupon   `json:"fixed_normal_coupon,omitempty"` //三选一 stock_type为NORMAL时必填。
	DiscountCoupon      DiscountCoupon      `json:"discount_coupon,omitempty"`     //三选一 stock_type为DISCOUNT时必填。
	ExchangeCoupon      ExchangeCoupon      `json:"exchange_coupon,omitempty"`     //三选一 stock_type为EXCHANGE时必填。
	UseMethod           string              `json:"use_method"`
	MiniProgramsAppid   string              `json:"mini_programs_appid,omitempty"`
	MiniProgramsPath    string              `json:"mini_programs_path,omitempty"`
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
	StockID    string `json:"stock_id"`
	CreateTime string `json:"create_time"`
}

type BusiFavorStockInfo struct {
	StockName            string               `json:"stock_name"`
	BelongMerchant       string               `json:"belong_merchant"`
	Comment              string               `json:"comment,omitempty"`
	GoodsName            string               `json:"goods_name"`
	StockType            string               `json:"stock_type"`
	CouponUseRule        CouponUseRule        `json:"coupon_use_rule"`
	StockSendRule        StockSendRule        `json:"stock_send_rule"`
	CustomEntrance       CustomEntrance       `json:"custom_entrance,omitempty"`
	DisplayPatternInfo   DisplayPatternInfo   `json:"display_pattern_info"`
	StockState           string               `json:"stock_state"`
	CouponCodeMode       string               `json:"coupon_code_mode"`
	StockID              string               `json:"stock_id"`
	NotifyConfig         *NotifyConfig        `json:"notify_config,omitempty"`
	CouponCodeCount      CouponCodeCount      `json:"coupon_code_count,omitempty"`
	CendCountInformation CendCountInformation `json:"send_count_information,omitempty"`
}

type CendCountInformation struct {
	CouponCode uint64 `json:"total_send_num,omitempty"`    // 批次已发放的券数量，满减、折扣、换购类型会返回该字段
	StockID    uint64 `json:"total_send_amount,omitempty"` // 批次已发放的预算金额，满减券类型会返回该字段
	Appid      uint64 `json:"today_send_num,omitempty"`    // 批次当天已发放的券数量，设置了单天发放上限的满减、折扣、换购类型返回该字段
	UseTime    uint64 `json:"today_send_amount,omitempty"` // 批次当天已发放的预算金额，设置了当天发放上限的满减券类型返回该字段
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

// UserCouponRequest 根据过滤条件查询用户券请求参数
type UserCouponRequest struct {
	Appid           string `json:"appid"`                      // 支持传入与当前调用接口商户号有绑定关系的appid。支持小程序appid与公众号appid。
	StockID         string `json:"stock_id,omitempty"`         // 微信为每个商家券批次分配的唯一ID，是否指定批次号查询。
	CouponState     string `json:"coupon_state,omitempty"`     // 券状态 枚举值：SENDED：可用 USED：已核销 EXPIRED：已过期
	CreatorMerchant string `json:"creator_merchant,omitempty"` // 批次创建方商户号 校验规则：当调用方商户号（即请求头中的商户号）为创建批次方商户号时，该参数必传
	BelongMerchant  string `json:"belong_merchant,omitempty"`  // 批次归属商户号 校验规则：当调用方商户号（即请求头中的商户号）为批次归属商户号时，该参数必传
	SenderMerchant  string `json:"sender_merchant,omitempty"`  // 批次发放商户号 校验规则：当调用方商户号（即请求头中的商户号）为批次发放商户号时，该参数必传；委托营销关系下，请填写委托发券的商户号
	Offset          int    `json:"offset,omitempty"`           // 分页页码
	Limit           int    `json:"limit,omitempty"`            // 分页大小
}

// UserCouponList 用户领取商家券列表
type UserCouponList struct {
	List       []UserCoupon `json:"data,omitempty"`
	TotalCount int          `json:"total_count"`
	Limit      int          `json:"limit"`
	Offset     int          `json:"offset"`
}

// UserCoupon 用户单张券详情
type UserCoupon struct {
	BelongMerchant     string             `json:"belong_merchant"`                 // 批次归属于哪个商户。
	StockName          string             `json:"stock_name"`                      // 批次名称，字数上限为21个，一个中文汉字/英文字母/数字均占用一个字数。
	Comment            string             `json:"comment,omitempty"`               // 仅配置商户可见，用于自定义信息。字数上限为20个，一个中文汉字/英文字母/数字均占用一个字数。
	GoodsName          string             `json:"goods_name"`                      // 适用商品范围，字数上限为15个，一个中文汉字/英文字母/数字均占用一个字数。
	StockType          string             `json:"stock_type"`                      // 批次类型 NORMAL：固定面额满减券批次 DISCOUNT：折扣券批次 EXCHANGE：换购券批次
	Transferable       bool               `json:"transferable,omitempty"`          // 不填默认否，枚举值： true：是 false：否 该字段暂未开放
	Shareable          string             `json:"shareable,omitempty"`             // 不填默认否，枚举值： true：是 false：否
	CouponState        string             `json:"coupon_state,omitempty"`          // 商家券状态 枚举值： SENDED：可用 USED：已核销 EXPIRED：已过期 DEACTIVATED：已失效
	DisplayPatternInfo DisplayPatternInfo `json:"display_pattern_info,,omitempty"` // 商家券详细信息
	CouponUseRule      CouponUseRule      `json:"coupon_use_rule"`                 // 券核销相关规则
	CustomEntrance     CustomEntrance     `json:"custom_entrance,omitempty"`       // 卡详情页面，可选择多种入口引导用户。
	CouponCode         string             `json:"coupon_code,omitempty"`           // 券的唯一标识。
	StockID            string             `json:"stock_id,omitempty"`              // 微信为每个商家券批次分配的唯一ID，是否指定批次号查询。
	AvailableStartTime string             `json:"available_start_time"`            // 用户领取到该张券实际可使用的开始时间 若券批次设置为领取后可用，则开始时间即为券的领取时间；若券批次设置为领取后第X天可用，则开始时间为券领取时间后第X天00:00:00可用。
	ExpireTime         string             `json:"expire_time"`                     // 用户领取到该张券的过期时间
	ReceiveTime        string             `json:"receive_time"`                    // 用户领取到该张券的时间
	SendRequestNo      string             `json:"send_request_no"`                 // 发券时传入的唯一凭证
	UseRequestNo       string             `json:"use_request_no,omitempty"`        // 核销时传入的唯一凭证（如券已被核销，将返回此字段）
	UseTime            string             `json:"use_time,omitempty"`              // 券被核销的时间（如券已被核销，将返回此字段）
}

// Callbacks 商家券事件通知地址
type Callbacks struct {
	MchID      string `json:"mchid,omitempty"`       // 微信支付商户的商户号，由微信支付生成并下发，不填默认查询调用方商户的通知URL。
	NotifyURL  string `json:"notify_url"`            //  商户提供的用于接收商家券事件通知的url地址，必须支持https。
	UpdateTime string `json:"update_time,omitempty"` // 修改时间
}

// CouponAssociateRequest 关联订单信息请求参数
type CouponAssociateRequest struct {
	CouponCode   string `json:"coupon_code"`    // 券的唯一标识
	OutTradeNo   string `json:"out_trade_no"`   // 微信支付下单时的商户订单号，欲与该商家券关联的微信支付
	StockID      string `json:"stock_id"`       // 微信为每个商家券批次分配的唯一ID，对于商户自定义code的批次，关联请求必须填写批次号
	OutRequestNo string `json:"out_request_no"` // 商户创建批次凭据号（格式：商户id+日期+流水号），商户侧需保持唯一性，可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号。
}

// CouponAssociateResponse 关联订单信息返回参数
type CouponAssociateResponse struct {
	WechatpayAssociateTime string `json:"wechatpay_associate_time"` // 系统关联券成功的时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE 示例值：2015-05-20T13:29:35+08:00
}

// CouponDisassociateResponse 取消关联订单信息返回参数
type CouponDisassociateResponse struct {
	WechatpayDisassociateTime string `json:"wechatpay_disassociate_time"` // 系统成功取消商家券与订单信息关联关系的时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE 示例值：2015-05-20T13:29:35+08:00
}

// CouponStocksBudgetRequset 修改批次预算请求参数
type CouponStocksBudgetRequset struct {
	ModifyBudgetRequestNo  string `json:"modify_budget_request_no"`             // 修改预算请求单据号
	TargetMaxCoupons       int    `json:"target_max_coupons,omitempty"`         // 批次最大发放个数 与 target_max_coupons_by_day 二选一
	TargetMaxCouponsByDay  int    `json:"target_max_coupons_by_day,omitempty"`  // 单天发放上限个数
	CurrentMaxCoupons      int    `json:"current_max_coupons,omitempty"`        // 当前批次最大发放个数，当传入target_max_coupons大于0时，current_max_coupons必传
	CurrentMaxCouponsByDay int    `json:"current_max_coupons_by_day,omitempty"` // 当前单天发放上限个数 ，当传入target_max_coupons_by_day大于0时，current_max_coupons_by_day必填
}

// CouponReturnRequest 申请退券请求参数
type CouponReturnRequest struct {
	CouponCode      string `json:"coupon_code"`       // 券的唯一标识
	StockID         string `json:"stock_id"`          // 券的所属批次号
	ReturnRequestNo string `json:"return_request_no"` // 每次退券请求的唯一标识，商户需保证唯一
}

// CouponReturnResponse 申请退券返回参数
type CouponReturnResponse struct {
	WechatpayReturnTime string `json:"wechatpay_return_time"`
}

// CouponReturnRequest 使券失效请求参数
type CouponDeactivateRequest struct {
	CouponCode          string `json:"coupon_code"`                 // 券的唯一标识
	StockID             string `json:"stock_id"`                    // 券的所属批次号
	DeactivateRequestNo string `json:"deactivate_request_no"`       // 每次失效请求的唯一标识，商户需保证唯一
	DeactivateReason    string `json:"deactivate_reason,omitempty"` // 商户失效券的原因
}

// CouponDeactivateResponse 返回使券失效请求参数
type CouponDeactivateResponse struct {
	WechatpayReturnTime string `json:"wechatpay_return_time"`
}

// SubsidyPayReceiptsRequest  营销补差付款请求参数
type SubsidyPayReceiptsRequest struct {
	StockID       string `json:"stock_id"`       // 由微信支付生成，调用创建商家券API成功时返回的唯一批次ID 仅支持“满减券”，“换购券”批次不支持
	CouponCode    string `json:"coupon_code"`    // 券的唯一标识。在WECHATPAY_MODE的券Code模式下，商家券Code是由微信支付生成的唯一ID；在MERCHANT_UPLOAD、MERCHANT_API的券Code模式下，商家券Code是由商户上传或指定，在批次下保证唯一；
	TransactionID string `json:"transaction_id"` // 微信支付下单支付成功返回的订单号
	PayerMerchant string `json:"payer_merchant"` // 营销补差扣款商户号注：补差扣款商户号 = 制券商户号 或 补差扣款商户号 = 归属商户号
	PayeeMerchant string `json:"payee_merchant"` //营销补差入账商户号注：补差入帐商户号 = 券归属商户号 或者和 券归属商户号有连锁品牌关系
	Amount        int    `json:"amount"`         // 单位为分，单笔订单补差金额不得超过券的优惠金额，最高补差金额为5000元 > 券的优惠金额定义：满减券：满减金额即为优惠金额换购券：不支持
	Description   string `json:"description"`    // 付款备注描述，查询的时候原样带回
	OutSubsidyNo  string `json:"out_subsidy_no"` // 商户侧需保证唯一性。可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号
}

// SubsidyPayReceiptsResponse 营销补差付款返回参数
type SubsidyPayReceiptsResponse struct {
	SubsidyReceiptID string `json:"subsidy_receipt_id"`     // 补差付款唯一单号，由微信支付生成，仅在补差付款成功后有返回
	StockID          string `json:"stock_id"`               // 由微信支付生成，调用创建商家券API成功时返回的唯一批次ID
	CouponCode       string `json:"coupon_code"`            // 券的唯一标识
	TransactionID    string `json:"transaction_id"`         //  微信支付下单支付成功返回的订单号
	PayerMerchant    string `json:"payer_merchant"`         // 营销补差扣款商户号
	PayeeMerchant    string `json:"payee_merchant"`         // 营销补差入账商户号
	Amount           int    `json:"amount"`                 // 单位为分，单笔订单补差金额不得超过券的优惠金额，最高补差金额为5000元 > 券的优惠金额定义：满减券：满减金额即为优惠金额换购券：不支持
	Description      string `json:"description"`            // 付款备注描述，查询的时候原样带回
	Status           string `json:"status"`                 // 补差付款单据状态ACCEPTED：受理成功SUCCESS：补差补款成功FAIL：补差付款失败RETURNING：补差回退中PARTIAL_RETURN：补差部分回退FULL_RETURN：补差全额回退
	SuccessTime      string `json:"success_time,omitempty"` // 仅在补差付款成功时，返回完成时间
	OutSubsidyNo     string `json:"out_subsidy_no"`         // 商户侧需保证唯一性。可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号
	CreateTime       string `json:"create_time,omitempty"`  //补差付款单据创建时间。
}

// BusiFavorSend 领券成功通知参数
type BusiFavorSend struct {
	EventType    string `json:"event_type"`
	CouponCode   string `json:"coupon_code"`
	StockID      string `json:"stock_id"`
	SendTime     string `json:"send_time"`
	Openid       string `json:"openid,omitempty"`
	Unionid      string `json:"unionid,omitempty"`
	SendChannel  string `json:"send_channel"`
	SendMerchant string `json:"send_merchant"`
	AttachInfo   string `json:"attach_info,omitempty"`
	// AttachInfo   struct {
	// 	TransactionID   string `json:"transaction_id,omitempty"`
	// 	ActCode         string `json:"act_code,omitempty"`
	// 	HallCode        string `json:"hall_code,omitempty"`
	// 	HallBelongMchID string `json:"hall_belong_mch_id,omitempty"`
	// 	CardID          string `json:"card_id,omitempty"`
	// 	Code            string `json:"code,omitempty"`
	// 	ActivityID      string `json:"activity_id,omitempty"`
	// } `json:"attach_info,omitempty"`
}
