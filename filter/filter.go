package filter

import (
	"github.com/liuchonglin/go-utils"
	"context"
	"github.com/liuchonglin/go-tools/result"
	"time"
	"golang.org/x/time/rate"
	"github.com/gin-gonic/gin"
	"github.com/liuchonglin/go-tools/commonContext"
	"google.golang.org/grpc/codes"
	"github.com/liuchonglin/go-tools/grpcx/code"
)

const (
	logTag = "core.filter"
)

var limiter = rate.NewLimiter(40000, 20000)

// 限流过滤器
func RateFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// AllowN标识在时间now的时候，n个事件是否可以同时发生(也意思就是now的时候是否可以从令牌池中取n个令牌)。
		// 如果你需要在事件超出频率的时候丢弃或跳过事件
		if !limiter.AllowN(time.Now(), 1) {
			c.JSON(code.GRPCCodeToHTTPStatus(codes.ResourceExhausted), result.NewError(codes.ResourceExhausted, code.GetMessage(codes.ResourceExhausted)))
			// 终止
			c.Abort()
			return
		}
		// 执行下一个中间件
		c.Next()
	}
}

// 上下文处理
func ContextHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader(commonContext.CtxKeyTraceId)
		if utils.IsEmpty(traceId) || len(traceId) != 36 {
			c.JSON(code.GRPCCodeToHTTPStatus(codes.InvalidArgument), result.NewError(codes.InvalidArgument, code.GetMessage(codes.InvalidArgument)))
			c.Abort()
			return
		}
		ctx := context.WithValue(context.Background(), commonContext.CtxKeyTraceId, traceId)

		c.Set(commonContext.CtxKeyCommon, ctx)
		c.Next()
	}
}
