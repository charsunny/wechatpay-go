package favor

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type FavorService services.Service

// 创建代金券批次API
// 适用对象： 服务商
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/coupon-stocks
// 请求方式：POST
// 接口频率：单个商户号调用频率60/min，所有商户号调用频率为1000/min
// 接口耗时：80ms
func (s *FavorService) CouponStocks(ctx context.Context, req *CouponStockRequest) (resp *CouponStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/coupon-stocks"

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CouponStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 激活代金券批次API
// 适用对象：服务商
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/start
// 请求方式：POST
// path 指该参数需在请求URL传参
// query 指该参数需在请求URL传参
// body 指该参数需在请求JSON传参
func (s *FavorService) StartStock(ctx context.Context, stock_creator_mchid string, stock_id string) (resp *CouponStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s/start", stock_id)

	// Setup Body Params
	localVarPostBody = map[string]string{
		"stock_creator_mchid": stock_creator_mchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CouponStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 暂停代金券批次API
// 适用对象：服务商
// 通过此接口可暂停指定代金券批次。暂停后，该代金券批次暂停发放，用户无法通过任何渠道再领取该批次的券。
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/start
// 请求方式：POST
// path 指该参数需在请求URL传参
// query 指该参数需在请求URL传参
// body 指该参数需在请求JSON传参
func (s *FavorService) PauseStock(ctx context.Context, stock_creator_mchid string, stock_id string) (resp *CouponStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s/pause", stock_id)

	// Setup Body Params
	localVarPostBody = map[string]string{
		"stock_creator_mchid": stock_creator_mchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CouponStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 重启代金券批次API
// 适用对象：服务商
// 通过此接口可重启指定代金券批次。重启后，该代金券批次可以再次发放。
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/start
// 请求方式：POST
// path 指该参数需在请求URL传参
// query 指该参数需在请求URL传参
// body 指该参数需在请求JSON传参
func (s *FavorService) RestartStock(ctx context.Context, stock_creator_mchid string, stock_id string) (resp *CouponStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s/restart", stock_id)

	// Setup Body Params
	localVarPostBody = map[string]string{
		"stock_creator_mchid": stock_creator_mchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CouponStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (s *FavorService) SendCoupons(ctx context.Context, openid string, req *CouponSendRequest) (resp *CouponSendResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/users/%s/coupons", openid)

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(CouponSendResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 设置消息通知地址API
// 适用对象：直连商户
// 接口频率：不区分来源 1000/min 单ip 5/min
// 幂等规则：接口支持幂等重入
// 接口耗时：100ms
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/callbacks
// 请求方式：POST
func (s *FavorService) SetNotifyURL(ctx context.Context, req *Callbacks) (resp *Callbacks, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/callbacks"

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(Callbacks)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 查询批次详情API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}
// 请求方式：GET
// 接口频率：不区分来源 1000/s 单ip 500/s
// 接口耗时：1S
// 幂等规则：接口支持幂等重入
func (s *FavorService) GetCouponStocks(ctx context.Context, stockID, mchID string) (resp *FavorStockInfo, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("stock_creator_mchid", mchID)
	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s", stockID)

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(FavorStockInfo)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 条件查询批次列表API
// 通过此接口可查询多个批次的信息，包括批次的配置信息以及批次概况数据。
// 适用对象：直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks
// 请求方式：GET
// 接口频率：不区分来源 1000/s 单ip 500/s
// 接口耗时：1S
// 幂等规则：接口支持幂等重入
func (s *FavorService) GetFavorStockList(ctx context.Context, openid string, req *FavorStockListRequest) (resp *FavorStockList, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks"

	// Setup Body Params
	// localVarPostBody = req
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("stock_creator_mchid", req.StockCreatorMchid)
	localVarQueryParams.Add("create_start_time", req.CreateStartTime)
	localVarQueryParams.Add("create_end_time", req.CreateEndTime)
	localVarQueryParams.Add("status", req.Status)
	localVarQueryParams.Add("offset", fmt.Sprintf("%d", req.Offset))
	localVarQueryParams.Add("limit", fmt.Sprintf("%d", req.Limit))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(FavorStockList)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 查询代金券详情API
// 通过此接口可查询代金券信息，包括代金券的基础信息、状态。如代金券已核销，会包括代金券核销的订单信息（订单号、单品信息等）。
// 适用对象：直连商户
// 接口频率：不区分来源 1000/s， 单ip 500/s
// 接口耗时：1S
// 幂等规则：接口支持幂等重入
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/users/{openid}/coupons/{coupon_id}
// 请求方式：GET
func (s *FavorService) GetUserFavor(ctx context.Context, couponID, appid, openid string) (resp *UserFavor, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("appid", appid)
	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/users/%s/coupons/%s", openid, couponID)

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(UserFavor)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 查询代金券可用商户API
// 通过调用此接口可查询批次的可用商户号，判断券是否在某商户号可用，来决定是否展示。
// 适用对象：直连商户
// 接口频率：不区分来源 1000/min，单ip 200/min
// 接口耗时：1S
// 幂等规则：接口支持幂等重入
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/merchants
// 请求方式：GET
func (s *FavorService) GetFavorMerchants(ctx context.Context, openid string, req *FavorMerchantsOrItemsRequest) (resp *FavorMerchantsOrItems, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s/merchants", req.StockID)

	// Setup Body Params
	// localVarPostBody = req
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("stock_creator_mchid", req.StockCreatorMchid)
	localVarQueryParams.Add("offset", fmt.Sprintf("%d", req.Offset))
	localVarQueryParams.Add("limit", fmt.Sprintf("%d", req.Limit))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(FavorMerchantsOrItems)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 查询代金券可用商户API
// 通过此接口可查询批次的可用商品编码，判断券是否可用于某些商品，来决定是否展示。
// 适用对象：直连商户
// 接口频率：不区分来源 1000/min，单ip 200/min
// 接口耗时：1S
// 幂等规则：接口支持幂等重入
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/items
// 请求方式：GET
func (s *FavorService) GetFavorItems(ctx context.Context, openid string, req *FavorMerchantsOrItemsRequest) (resp *FavorMerchantsOrItems, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s/items ", req.StockID)

	// Setup Body Params
	// localVarPostBody = req
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("stock_creator_mchid", req.StockCreatorMchid)
	localVarQueryParams.Add("offset", fmt.Sprintf("%d", req.Offset))
	localVarQueryParams.Add("limit", fmt.Sprintf("%d", req.Limit))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(FavorMerchantsOrItems)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 根据商户号查用户的券
// 可通过该接口查询用户在某商户号可用的全部券，可用于商户的小程序/H5中，用户"我的代金券"或"提交订单页"展示优惠信息。无法查询到微信支付立减金。本接口查不到用户的微信支付立减金（又称“全平台通用券”），即在所有商户都可以使用的券，例如：摇摇乐红包；当按可用商户号查询时，无法查询用户已经核销的券
// 适用对象：直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/users/{openid}/coupons
// 请求方式：GET
// 接口频率：不区分来源 2000/s 单ip 500/s
// 接口耗时：平均100ms以内
// 幂等规则：接口支持幂等重入
func (s *FavorService) GetUserFavorsByMchid(ctx context.Context, openid string, req *UserFavorsRequest) (resp *UserFavors, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/users/%s/coupons ", openid)

	// Setup Body Params
	// localVarPostBody = req
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("appid", req.Appid)
	localVarQueryParams.Add("stock_id", req.StockID)
	localVarQueryParams.Add("status", req.Status)
	localVarQueryParams.Add("sender_mchid", req.SenderMchid)
	localVarQueryParams.Add("available_mchid", req.AvailableMchid)
	localVarQueryParams.Add("stock_creator_mchid", req.StockCreatorMchid)
	localVarQueryParams.Add("offset", fmt.Sprintf("%d", req.Offset))
	localVarQueryParams.Add("limit", fmt.Sprintf("%d", req.Limit))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(UserFavors)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 下载批次核销明细API
// 可获取到某批次的核销明细数据，包括订单号、单品信息、银行流水号等，用于对账/数据分析。
// • 账单文件的下载地址的有效时间为30s。
// • 强烈建议商户将实际账单文件的哈希值和之前从接口获取到的哈希值进行比对，以确认数据的完整性。
// • 该接口响应的信息请求头中不包含微信接口响应的签名值，因此需要跳过验签的流程。
// • 该接口需要活动结束后次日10点才可以下载。
// • 该接口只能使用创建活动方进行调用。
// 适用对象：直连商户
// 接口频率：不区分来源 1000/min， 单ip 5/min
// 幂等规则：接口支持幂等重入
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/use-flow
// 请求方式：GET
func (s *FavorService) GetFavorUseFlow(ctx context.Context, stockID string) (resp *Flow, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s/use-flow ", stockID)

	// Setup Body Params
	// localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(Flow)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 下载批次退款明细API
// 可获取到某批次的退款明细数据，包括订单号、单品信息、银行流水号等，用于对账/数据分析
// 适用对象：直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/refund-flow
// 请求方式：GET
func (s *FavorService) GetFavorRefundFlow(ctx context.Context, stockID string) (resp *Flow, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/favor/stocks/%s/use-flow ", stockID)

	// Setup Body Params
	// localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(Flow)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
