package busifavor

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type BusiFavorService services.Service

// 创建商家券API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/stocks
// 请求方式：POST
func (s *BusiFavorService) CreateCouponStocks(ctx context.Context, req *BusiFavorStockRequest) (resp *BusiFavorStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks"

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
	resp = new(BusiFavorStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 修改批次预算API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/stocks/{stock_id}/budget
// 请求方式：Patch
// 前置条件： 已创建商家券批次，且修改时间位于有效期结束时间前
func (s *BusiFavorService) ModifyCouponStocksBudget(ctx context.Context, stock_id string, req *CouponStocksBudgetRequset) (resp *StockSendRule, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/busifavor/stocks/%s/budget", stock_id)

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
	resp = new(StockSendRule)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, err
}

// 修改商家券基本信息API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/stocks/{stock_id}
// 请求方式：Patch
// 前置条件： 已创建商家券批次，且修改时间位于有效期结束时间前
func (s *BusiFavorService) ModifyCouponStocks(ctx context.Context, stock_id string, req *BusiFavorStockRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/busifavor/stocks/%s", stock_id)

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	return result, err
}

// 查询商家券详情API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/stocks/{stock_id}
// 请求方式：GET
func (s *BusiFavorService) GetCouponStocks(ctx context.Context, stock_id string) (resp *BusiFavorStockInfo, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/busifavor/stocks/%s", stock_id)

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
	resp = new(BusiFavorStockInfo)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 核销用户券API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/coupons/use
// 请求方式：POST
// 接口频率：500QPS
func (s *BusiFavorService) ConsumeCoupons(ctx context.Context, openid string, req *CouponConsumeRequest) (resp *CouponConsumeResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/coupons/use"

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
	resp = new(CouponConsumeResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 根据过滤条件查询用户券API
// 商户自定义筛选条件（如创建商户号、归属商户号、发放商户号等），查询指定微信用户卡包中满足对应条件的所有商家券信息
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/users/{openid}/coupons
// 请求方式：GET
func (s *BusiFavorService) GetUserCoupons(ctx context.Context, openid string, req *UserCouponRequest) (resp *UserCouponList, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/busifavor/users/%s/coupons", openid)

	// Setup Body Params
	// localVarPostBody = req

	localVarQueryParams.Add("appid", req.Appid)
	localVarQueryParams.Add("stock_id", req.StockID)
	localVarQueryParams.Add("coupon_state", req.CouponState)
	localVarQueryParams.Add("creator_merchant", req.CreatorMerchant)
	localVarQueryParams.Add("belong_merchant", req.BelongMerchant)
	localVarQueryParams.Add("sender_merchant", req.SenderMerchant)
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
	resp = new(UserCouponList)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 查询用户单张券详情API
// 商户可通过该接口查询微信用户卡包中某一张商家券的详情信息。
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/users/{openid}/coupons/{coupon_code}/appids/{appid}
// 请求方式：GET
func (s *BusiFavorService) GetUserCoupon(ctx context.Context, openid, couponCode, appid string) (resp *UserCoupon, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/marketing/busifavor/users/%s/coupons/%s/appids/%s/coupons", openid, couponCode, appid)

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
	resp = new(UserCoupon)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 设置商家券事件通知地址API
// 用于设置接收商家券相关事件通知的URL，可接收商家券相关的事件通知、包括发放通知等。需要设置接收通知的URL，并在商户平台开通营销事件推送的能力，即可接收到相关通知。
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/callbacks
// 请求方式：POST
func (s *BusiFavorService) SetNotifyURL(ctx context.Context, req *Callbacks) (resp *Callbacks, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/callbacks"

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

// 查询商家券事件通知地址API
// 通过调用此接口可查询设置的通知URL。
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/callbacks
// 请求方式：GET
// mchID 微信支付商户的商户号，由微信支付生成并下发，不填默认查询调用方商户的通知URL。
func (s *BusiFavorService) GetNotifyURL(ctx context.Context, mchID string) (resp *Callbacks, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/callbacks"

	// Setup Body Params
	// localVarPostBody = req

	localVarQueryParams.Add("mchid", mchID)

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

// 关联订单信息API
// 适用对象： 直连商户
// 请求URL： https://api.mch.weixin.qq.com/v3/marketing/busifavor/coupons/associate
// 请求方式： POST
// 接口频率： 500QPS
// 前置条件： 已为用户发券，且调用查询接口查到用户的券code、批次ID等信息
func (s *BusiFavorService) AssociateCoupons(ctx context.Context, req *CouponAssociateRequest) (resp *CouponAssociateResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/coupons/associate"

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
	resp = new(CouponAssociateResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 取消关联订单信息API
// 适用对象： 直连商户
// 请求URL： https://api.mch.weixin.qq.com/v3/marketing/busifavor/coupons/disassociate
// 请求方式： POST
// 接口频率： 500QPS
// 前置条件： 已调用关联接口为券创建过关联关系
func (s *BusiFavorService) DisassociateCoupons(ctx context.Context, req *CouponAssociateRequest) (resp *CouponDisassociateResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/coupons/disassociate"

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
	resp = new(CouponDisassociateResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 申请退券API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/coupons/return
// 请求方式：POST
// 前置条件：券的状态为USED
func (s *BusiFavorService) ReturnCoupons(ctx context.Context, req *CouponReturnRequest) (resp *CouponReturnResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/coupons/return"

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
	resp = new(CouponReturnResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 使券失效API
// 适用对象： 直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/coupons/deactivate
// 请求方式：POST
// 前置条件：券的状态为SENDED
func (s *BusiFavorService) DeactivateCoupons(ctx context.Context, req *CouponDeactivateRequest) (resp *CouponDeactivateResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/coupons/deactivate"

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
	resp = new(CouponDeactivateResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 营销补差付款API
// 该API主要用于商户营销补贴场景，支持基于单张券进行不同商户账户间的资金补差，从而提升财务结算、资金利用效率。https://pay.weixin.qq.com/wiki/doc/apiv3/download/%E5%95%86%E5%AE%B6%E5%88%B8%E8%A1%A5%E5%B7%AE%E4%BA%A7%E5%93%81%E6%93%8D%E4%BD%9C%E6%96%87%E6%A1%A3.pdf
// 适用对象：直连商户
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/busifavor/subsidy/pay-receipts
// 请求方式：POST
// 前置条件：商家必须核销了商家券且发起了微信支付收款
// 是否支持幂等：是
func (s *BusiFavorService) SubsidyPayReceipts(ctx context.Context, req *CouponDeactivateRequest) (resp *CouponDeactivateResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/coupons/deactivate"

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
	resp = new(CouponDeactivateResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
