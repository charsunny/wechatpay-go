package favor

import "github.com/wechatpay-apiv3/wechatpay-go/services"

type FavorService services.Service

// 创建代金券批次API
// 适用对象： 服务商
// 请求URL：https://api.mch.weixin.qq.com/v3/marketing/favor/coupon-stocks
// 请求方式：POST
// 接口频率：单个商户号调用频率60/min，所有商户号调用频率为1000/min
// 接口耗时：80ms
func (s *FavorService) CouponStocks(req *CouponStockRequest) {

}
