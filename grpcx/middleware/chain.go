package middleware

import (
	"google.golang.org/grpc"
	"context"
)

// ChainUnaryServer 将传入的拦截器们链接起来，统一创建为一个服务端拦截器
// 执行顺序按照从左到右顺序执行，包括传递的上下文
// 例如：ChainUnaryServer（one, two, three），执行one->执行two->执行three
// 前面两个拦截器的上下文发生了改变，第三个拦截器可以获取到
func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			var (
				chainHandler grpc.UnaryHandler
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				if curI == lastI {
					return handler(currentCtx, currentReq)
				}
				curI++
				resp, err := interceptors[curI](currentCtx, currentReq, info, chainHandler)
				curI--
				return resp, err
			}

			return interceptors[0](ctx, req, info, chainHandler)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; 虚拟拦截器保持向后兼容性以避免返回nil
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}

// ChainUnaryClient 将传入的拦截器们链接起来，统一创建为一个客户端拦截器
// 执行顺序按照从左到右顺序执行，包括传递的上下文
// 例如：ChainUnaryClient（one, two, three），执行one->执行two->执行three
// 前面两个拦截器的上下文发生了改变，第三个拦截器可以获取到
func ChainUnaryClient(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	n := len(interceptors)

	if n > 1 {
		lastI := n - 1
		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			var (
				chainHandler grpc.UnaryInvoker
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentMethod string, currentReq, currentRepl interface{}, currentConn *grpc.ClientConn, currentOpts ...grpc.CallOption) error {
				if curI == lastI {
					return invoker(currentCtx, currentMethod, currentReq, currentRepl, currentConn, currentOpts...)
				}
				curI++
				err := interceptors[curI](currentCtx, currentMethod, currentReq, currentRepl, currentConn, chainHandler, currentOpts...)
				curI--
				return err
			}

			return interceptors[0](ctx, method, req, reply, cc, chainHandler, opts...)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	// n == 0; 虚拟拦截器保持向后兼容性以避免返回nil
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
