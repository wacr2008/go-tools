package redisx

import (
	"testing"

	"github.com/go-redis/redis"
	"fmt"
)

func TestSub(t *testing.T) {
	type args struct {
		callBack func(msg *redis.Message)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				callBack: func(msg *redis.Message) {
					fmt.Println("具体的消息 = ", msg.Payload)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sub(tt.args.callBack)
		})
	}
}

func TestPub(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pub()
		})
	}
}
