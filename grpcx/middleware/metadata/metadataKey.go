package metadata

// grpc元数据上下文常量
const (
	// 用户唯一标识
	UserIDMedCtxKey = "med-ctx-user-id"
	// 请求链跟踪唯一标识
	TraceIdMedCtxKey = "med-ctx-trace-id"
	// 用户令牌
	TokenMedCtxKey = "med-ctx-token"
	// 令牌载荷
	TokenDataMedCtxKey = "med-ctx-token-date"
	// 请求服务名
	OriginNameMedCtxKey = "med-ctx-origin-name"
	// 请求服务IP地址
	IPMedCtxKey = "med-ctx-ip"
)
