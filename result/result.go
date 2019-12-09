package result

import (
	"google.golang.org/grpc/codes"
	"github.com/liuchonglin/go-tools/grpcx/code"
	"google.golang.org/grpc/status"
)

type Result struct {
	Code    codes.Code  `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total,omitempty"`
}

func NewSuccess() *Result {
	return &Result{Code: codes.OK, Message: code.GetMessage(codes.OK)}
}

func NewError(code codes.Code, message string) *Result {
	return &Result{Code: code, Message: message}
}

func NewSuccessData(data interface{}) *Result {
	return &Result{Code: codes.OK, Message: code.GetMessage(codes.OK), Data: data}
}

func NewSuccessPage(data interface{}, total int64) *Result {
	return &Result{Code: codes.OK, Message: code.GetMessage(codes.OK), Data: data, Total: total}
}

func ErrorConvert(err error) (int, *Result) {
	stu := status.Convert(err)
	var message string
	if stu.Code() == codes.FailedPrecondition {
		message = stu.Message()
	} else {
		message = code.GetMessage(stu.Code())
	}
	httpStatus := code.GRPCCodeToHTTPStatus(stu.Code())
	return httpStatus, NewError(stu.Code(), message)
}
