package commonContext

import (
	"context"
	"github.com/liuchonglin/go-utils"
)

func SetTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, CtxKeyTraceId, traceId)
}

func GetTraceId(ctx context.Context) string {
	val := ctx.Value(CtxKeyTraceId)
	if utils.IsEmpty(val) {
		return ""
	}
	traceId, ok := val.(string)
	if !ok {
		return ""
	}
	return traceId
}

func SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, CtxKeyToken, token)
}

func GetToken(ctx context.Context) string {
	val := ctx.Value(CtxKeyToken)
	if utils.IsEmpty(val) {
		return ""
	}
	token, ok := val.(string)
	if !ok {
		return ""
	}
	return token
}

func SetTokenData(ctx context.Context, tokenData map[string]interface{}) context.Context {
	return context.WithValue(ctx, CtxKeyTokenData, tokenData)
}

func GetTokenData(ctx context.Context) map[string]interface{} {
	val := ctx.Value(CtxKeyTokenData)
	if utils.IsEmpty(val) {
		return nil
	}
	tokenData, ok := val.(map[string]interface{})
	if !ok {
		return nil
	}
	return tokenData
}

func SetClientIp(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, CtxKeyClientIp, ip)
}

func GetClientIp(ctx context.Context) string {
	val := ctx.Value(CtxKeyClientIp)
	if utils.IsEmpty(val) {
		return ""
	}
	clientIp, ok := val.(string)
	if !ok {
		return ""
	}
	return clientIp
}
