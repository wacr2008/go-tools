package commonContext

const (
	CtxKeyTraceId   = "traceId"   // 全链路追踪唯一标识
	CtxKeyToken     = "token"     // 用户令牌
	CtxKeyTokenData = "tokenData" // 用户令牌中的数据
	CtxKeyClientIp  = "clientIp"  // 客户端IP地址
	CtxKeyCommon    = "commonContext"
)
