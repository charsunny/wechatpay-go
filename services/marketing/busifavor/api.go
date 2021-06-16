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

// 创建代金券批次API
// 适用对象： 服务商
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/coupon-stocks
// 请求方式：POST
// 接口频率：单个商户号调用频率60/min，所有商户号调用频率为1000/min
// 接口耗时：80ms
func (s *BusiFavorService) CreateCouponStocks(ctx context.Context, req *BusiFavorStockRequest) (resp *BusiFavorStockResponse, result *core.APIResult, err error) {
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
	resp = new(BusiFavorStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 创建代金券批次API
// 适用对象： 服务商
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/coupon-stocks
// 请求方式：POST
// 接口频率：单个商户号调用频率60/min，所有商户号调用频率为1000/min
// 接口耗时：80ms
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

// 重启代金券批次API
// 适用对象：服务商
// 通过此接口可重启指定代金券批次。重启后，该代金券批次可以再次发放。
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/{stock_id}/start
// 请求方式：POST
// path 指该参数需在请求URL传参
// query 指该参数需在请求URL传参
// body 指该参数需在请求JSON传参
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
