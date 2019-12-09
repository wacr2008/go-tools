package metadata

import (
	"google.golang.org/grpc"
	"context"
	"google.golang.org/grpc/metadata"
	"github.com/liuchonglin/go-tools/commonContext"
	"github.com/liuchonglin/go-utils"
	"encoding/json"
)

func ClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		tokenData := commonContext.GetTokenData(ctx)
		tokenDataJson := ""
		if !utils.IsEmpty(tokenData) {
			jsonData, err := json.Marshal(tokenData)
			if err != nil {
				panic(err)
			}
			tokenDataJson = string(jsonData)
		}

		md := metadata.New(map[string]string{
			TraceIdMedCtxKey:   commonContext.GetTraceId(ctx),
			TokenMedCtxKey:     commonContext.GetToken(ctx),
			TokenDataMedCtxKey: tokenDataJson,
		})
		ctx = metadata.NewOutgoingContext(ctx, md)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
