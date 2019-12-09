package metadata

import (
	"google.golang.org/grpc"
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/liuchonglin/go-tools/commonContext"
	"github.com/liuchonglin/go-utils"
	"encoding/json"
)

var errDataLossMetadata = status.Error(codes.DataLoss, "无法获取元数据")

func ServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errDataLossMetadata
		}
		for k, v := range md {
			switch k {
			case TraceIdMedCtxKey:
				ctx = commonContext.SetTraceId(ctx, v[0])
			case TokenMedCtxKey:
				ctx = commonContext.SetToken(ctx, v[0])
			case TokenDataMedCtxKey:
				jsonData := v[0]
				if !utils.IsEmpty(jsonData) {
					tokenData := make(map[string]interface{})
					err := json.Unmarshal([]byte(jsonData), &tokenData)
					if err == nil {
						ctx = commonContext.SetTokenData(ctx, tokenData)
					}
				}
			}
		}
		
		return handler(ctx, req)
	}
}
