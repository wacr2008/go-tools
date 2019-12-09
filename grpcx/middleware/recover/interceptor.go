package recover

import (
	"google.golang.org/grpc"
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

var errInternal = status.Error(codes.Internal, "服务器内部错误")

func ServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		fmt.Println("FullMethod=", info.FullMethod)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("服务器内部发生恐慌:", r)
				// 将恐慌转成grpc服务器内部错误
				err = errInternal
			}
		}()
		return handler(ctx, req)
	}
}
