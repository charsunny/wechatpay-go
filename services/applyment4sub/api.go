package applyment4sub

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type Applyment4SubService services.Service

// 提交申请单API
// 适用对象： 服务商
// 请求URL：https://api.mch.weixin.qq.com/v3/applyment4sub/applyment/
// 请求方式：POST
// 接口频率：单个商户号调用频率60/min，所有商户号调用频率为1000/min
// 接口耗时：80ms
func (s *Applyment4SubService) ApplySubMch(ctx context.Context, req *ApplymentRequest) (resp *ApplymentResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/applyment4sub/applyment"

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
	resp = new(ApplymentResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 通过业务申请编号查询申请状态
// 适用对象：服务商
// 通过此接口可重启指定代金券批次。重启后，该代金券批次可以再次发放。
// 请求URL：https://api.mch.weixin.qq.com/v3/applyment4sub/applyment/business_code/{business_code}
// GET
func (s *Applyment4SubService) GetAuditStatusByBusinessCode(ctx context.Context, business_code string) (resp *ApplyAuditResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/applyment4sub/applyment/business_code/%s", business_code)

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
	resp = new(ApplyAuditResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 通过申请单号查询申请状态
// 适用对象：服务商
// 通过此接口可重启指定代金券批次。重启后，该代金券批次可以再次发放。
// 请求URL：https://api.mch.weixin.qq.com/v3/applyment4sub/applyment/business_code/{business_code}
// GET
func (s *Applyment4SubService) GetAuditStatusByApplymentId(ctx context.Context, applyment_id string) (resp *ApplyAuditResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/applyment4sub/applyment/applyment_id/%s", applyment_id)

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
	resp = new(ApplyAuditResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// 修改结算账号API
// 适用对象： 服务商
// 请求URL：https://api.mch.weixin.qq.com/v3/apply4sub/sub_merchants/{sub_mchid}/modify-settlement
// 请求方式：POST
// 接口频率：单个商户号调用频率60/min，所有商户号调用频率为1000/min
// 接口耗时：80ms
func (s *Applyment4SubService) ModifySettlement(ctx context.Context, sub_mchid string, req *BankAccountInfo) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/applyment4sub/applyment/applyment_id/%s", sub_mchid)

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	return
}

// 通过申请单号查询申请状态
// 适用对象：服务商
// 通过此接口可重启指定代金券批次。重启后，该代金券批次可以再次发放。
// 请求URL：https://api.mch.weixin.qq.com/v3/apply4sub/sub_merchants/{sub_mchid}/settlement
// GET
func (s *Applyment4SubService) QuerySettlement(ctx context.Context, sub_mchid string) (resp *BankAccountInfoResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams nethttp.Header
	)

	localVarPath := consts.WechatPayAPIServer + fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/settlement", sub_mchid)

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
	resp = new(BankAccountInfoResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
