package grpcx

import (
	"google.golang.org/grpc"
	"sync"
	"github.com/liuchonglin/go-utils"
	"errors"
	"google.golang.org/grpc/connectivity"
	"github.com/liuchonglin/go-tools/grpcx/middleware"
	"github.com/liuchonglin/go-tools/grpcx/middleware/timeOut"
	"github.com/liuchonglin/go-tools/grpcx/middleware/metadata"
)

var clientPool sync.Map

func GetClient(address string) (*grpc.ClientConn, error) {
	if utils.IsEmpty(address) {
		return nil, errors.New("address can not be empty")
	}
	temp, ok := clientPool.Load(address)
	if ok {
		coon, coonOk := temp.(*grpc.ClientConn)
		if !coonOk {
			return nil, errors.New("grpc client pool type err")
		}
		state := coon.GetState()
		// 检查链接的状态
		switch state {
		case connectivity.TransientFailure:
			fallthrough
		case connectivity.Shutdown:
			coon.Close()
			clientPool.Delete(address)
		default:
			return coon, nil
		}
	}

	newConn, err := newClient(address)
	if err != nil {
		return nil, err
	}
	clientPool.Store(address, newConn)
	return newConn, nil
}

func newClient(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(middleware.ChainUnaryClient(
			timeOut.ClientInterceptor(),
			metadata.ClientInterceptor())))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
