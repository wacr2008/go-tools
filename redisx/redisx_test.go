package redisx

import (
	"reflect"
	"sync"
	"testing"
	"time"
	"fmt"
	"os"
)

var config = &RedisConfig{
	Addrs:       []string{"127.0.0.1:6379"},
	MaxConnAge:  16,
	IdleTimeout: 300,
	Password:    "liu5522112",
}
var client *Client

func TestNewClient(t *testing.T) {
	type args struct {
		config *RedisConfig
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "ok",
			args: args{
				config: config,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.config)
			fmt.Printf("NewClient = %v", got)
		})
	}
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		command string
		args    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			if got := rc.Do(tt.args.command, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Del(t *testing.T) {
	type args struct {
		key []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				key: []string{"name"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Del(tt.args.key...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_Expire(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key        string
		expiration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.Expire(tt.args.key, tt.args.expiration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Expire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Expire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Exists(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.Exists(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Get(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Set(t *testing.T) {
	type args struct {
		key string
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				key: "name",
				val: "Evan",
			},
			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Set(tt.args.key, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetExpire(t *testing.T) {
	type args struct {
		key     string
		val     interface{}
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				key:     "name",
				val:     "Evan",
				timeout: 1 * 60 * time.Second,
			},
			want:    "OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.SetExpire(tt.args.key, tt.args.val, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SetExpire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.SetExpire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetNX(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key     string
		val     interface{}
		timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.SetNX(tt.args.key, tt.args.val, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SetNX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.SetNX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetXX(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key     string
		val     interface{}
		timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.SetXX(tt.args.key, tt.args.val, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SetXX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.SetXX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_HGet(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.HGet(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.HGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.HGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_HSet(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key   string
		field string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.HSet(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.HSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.HSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_HDel(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key    string
		fields []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.HDel(tt.args.key, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.HDel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.HDel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_HExists(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.HExists(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.HExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.HExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LPop(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.LPop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LPop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.LPop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LPush(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key string
		v   []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.LPush(tt.args.key, tt.args.v...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LPush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.LPush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RPop(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.RPop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RPop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.RPop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RPush(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key   string
		value []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.RPush(tt.args.key, tt.args.value...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RPush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.RPush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LLen(t *testing.T) {
	type fields struct {
		cmd       commands
		mode      string
		scriptDic sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &Client{
				cmd:       tt.fields.cmd,
				mode:      tt.fields.mode,
				scriptDic: tt.fields.scriptDic,
			}
			got, err := rc.LLen(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.LLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.LLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	client = NewClient(config)
	os.Exit(m.Run())
}
