package code

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

// codeMessageMap 是为api service或api Gateway提供的，
// 这些消息都是展示给最终用户的，您不应该把服务端的详细错误信息暴露给最终用户（可能涉及隐私数据），
// 因为他们也无法修复服务端的错误，如果是因为开发人员导致的错误，这时我们应该通过
// 这些消息引导最终用户去做正确的事情。
// 注意：下面缺少了 codes.FailedPrecondition 错误码， 一旦我们遇到它，必须透传错误信息，
// 这是下游服务想要展示给最终用户的信息。除了codes.FailedPrecondition 错误码，您都应该
// 隐藏下游服务给您的详细错误信息。
var codeMessageMap = map[codes.Code]string{
	codes.OK:                "成功",
	codes.InvalidArgument:   "参数错误",
	codes.OutOfRange:        "参数范围错误",
	codes.Unauthenticated:   "身份认证无效",
	codes.PermissionDenied:  "权限不足",
	codes.NotFound:          "资源未找到",
	codes.Canceled:          "网络不稳定",
	codes.Aborted:           "并发冲突",
	codes.AlreadyExists:     "资源已存在",
	codes.ResourceExhausted: "请求超出限制，请稍后重试",
	codes.DataLoss:          "数据丢失",
	codes.Unknown:           "未知错误",
	codes.Internal:          "服务器错误，请稍后重试",
	codes.Unimplemented:     "功能未实现",
	codes.Unavailable:       "服务暂时不可用，请稍后重试",
	codes.DeadlineExceeded:  "请求超时，请稍后重试",
}

// 获取错误信息
func GetMessage(code codes.Code) string {
	if v, ok := codeMessageMap[code]; ok {
		return v
	}
	return codeMessageMap[codes.Unknown]
}

// GRPCCodeToHTTPStatus 将gRPC错误代码转换为相应的HTTP响应状态。
// 参考: https://cloud.google.com/apis/design/errors
func GRPCCodeToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Aborted, codes.AlreadyExists:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	}
	// 其他code: codes.DataLoss codes.Unknown codes.Internal
	return http.StatusInternalServerError
}
