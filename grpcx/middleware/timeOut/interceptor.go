package timeOut

import (
	"google.golang.org/grpc"
	"context"
	"time"
)

func ClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		tCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		return invoker(tCtx, method, req, reply, cc, opts...)
	}
}
