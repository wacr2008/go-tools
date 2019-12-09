// Copyright 2019 go-tools Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// redis配置
package redisx

import (
	"github.com/liuchonglin/go-utils"
)

var redisConfig *RedisConfig
// redis 配置
type RedisConfig struct {
	// 是否是集群
	IsCluster bool `json:"isCluster" yaml:"isCluster"`
	// 连接地址
	Addrs []string `json:"addrs" yaml:"addrs"`
	// 最大连接超时时间
	MaxConnAge int `json:"maxConnAge" yaml:"maxConnAge"`
	// 空闲超时时间（最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭）
	IdleTimeout int `json:"idleTimeout" yaml:"idleTimeout"`
	// 密码
	Password string `json:"password" yaml:"password"`
}

func (r *RedisConfig) DefaultValue() {
	if utils.IsEmpty(r.Addrs) {
		r.Addrs = append(r.Addrs, "127.0.0.1:6379")
	}
	if r.MaxConnAge == 0 {
		r.MaxConnAge = 240
	}
	if r.IdleTimeout == 0 {
		r.IdleTimeout = 300
	}
}
