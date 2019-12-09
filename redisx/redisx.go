package redisx

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
	)

type commands interface {
	redis.Cmdable
	Do(args ...interface{}) *redis.Cmd
}

type Redisx interface {
	Do(command string, args ...interface{}) *Result
	Del(key ...string) (int64, error)
	Expire(key string, expiration time.Duration) (bool, error)
	Exists(key string) (bool, error)
	Get(key string) (string, error)
	Set(key string, val interface{}) (string, error)
	SetExpire(key string, val interface{}, timeout time.Duration) (string, error)
	SetNX(key string, val interface{}, timeout time.Duration) (bool, error)
	SetXX(key string, val interface{}, timeout time.Duration) (bool, error)
	HGet(key, field string) (string, error)
	HSet(key, field string, value interface{}) (bool, error)
	HDel(key string, fields []string) (int64, error)
	HExists(key string, field string) (bool, error)
	LPop(key string) (string, error)
	LPush(key string, v ...interface{}) (int64, error)
	RPop(key string) (string, error)
	RPush(key string, value ...interface{}) (int64, error)
	LLen(key string) (int64, error)
}

type Client struct {
	cmd       commands
	mode      string
	scriptDic sync.Map
}

func NewClient(config *RedisConfig) *Client {
	if config == nil {
		config = &RedisConfig{}
	}
	config.DefaultValue()

	var client *Client
	if config.IsCluster {
		client = &Client{
			cmd: redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:        config.Addrs,
				Password:     config.Password,
				PoolSize:     32,
				MinIdleConns: 4,
				MaxConnAge:   time.Duration(config.MaxConnAge) * time.Second,
				MaxRetries:   3,
				IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second,
			}),
			mode: "cluster"}
	} else {
		client = &Client{
			cmd: redis.NewClient(&redis.Options{
				Addr:         config.Addrs[0],
				Password:     config.Password,
				PoolSize:     32,
				MinIdleConns: 4,
				MaxConnAge:   time.Duration(config.MaxConnAge) * time.Second,
				MaxRetries:   3,
				IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second,
			}),
			mode: "single"}
	}
	return client
}

//Do run command customer
func (rc *Client) Do(command string, args ...interface{}) *Result {
	r := &Result{}

	if err := isInvalidCmd(command, rc.mode); err != nil {
		r.err = err
		return r
	}
	argsL := []interface{}{command}
	argsL = append(argsL, args...)
	r.val, r.err = rc.cmd.Do(argsL...).Result()
	return r
}

// 以下为key的操作
//
// Del 删除keys
// *注意，大key请用scan删除
func (rc *Client) Del(key ...string) (int64, error) {
	return rc.cmd.Del(key...).Result()
}

// Expire 为key添加过期时间
func (rc *Client) Expire(key string, expiration time.Duration) (bool, error) {
	return rc.cmd.Expire(key, expiration).Result()
}

// Exists 判断key是否存在
func (rc *Client) Exists(key string) (bool, error) {
	rs := rc.cmd.Exists(key)
	return rs.Val() == 1, rs.Err()
}

// 以下为String的操作 string默认decode&encode
//
// Get 查询key
func (rc *Client) Get(key string) (string, error) {
	return rc.cmd.Get(key).Result()
}

// Set 存入值
func (rc *Client) Set(key string, val interface{}) (string, error) {
	return rc.cmd.Set(key, val, 0).Result()
}

// SetWithExpire 存入值并添加过期时间
func (rc *Client) SetExpire(key string, val interface{}, timeout time.Duration) (string, error) {
	return rc.cmd.Set(key, val, timeout).Result()
}

// SetNX 仅当key不存在时set
func (rc *Client) SetNX(key string, val interface{}, timeout time.Duration) (bool, error) {
	return rc.cmd.SetNX(key, val, timeout).Result()
}

// SetNX 仅当key存在时set
func (rc *Client) SetXX(key string, val interface{}, timeout time.Duration) (bool, error) {
	return rc.cmd.SetXX(key, val, timeout).Result()
}

// 以下为hash的操作
//
// HGet 返回哈希表 key 中给定 field 的值。
func (rc *Client) HGet(key, field string) (string, error) {
	return rc.cmd.HGet(key, field).Result()
}

// HSet hash set
func (rc *Client) HSet(key, field string, value interface{}) (bool, error) {
	return rc.cmd.HSet(key, field, value).Result()
}

// HDel
func (rc *Client) HDel(key string, fields []string) (int64, error) {
	return rc.cmd.HDel(key, fields...).Result()
}

// HExists
func (rc *Client) HExists(key string, field string) (bool, error) {
	return rc.cmd.HExists(key, field).Result()
}

// 以下为 List 的操作
//
// LPop 移除并返回列表 key 的头元素。
func (rc *Client) LPop(key string) (string, error) {
	return rc.cmd.LPop(key).Result()
}

// LPush 将 value 插入到列表 key 的表头
func (rc *Client) LPush(key string, v ...interface{}) (int64, error) {
	return rc.cmd.LPush(key, v...).Result()
}

// RPop 移除并返回列表 key 的尾元素。
func (rc *Client) RPop(key string) (string, error) {
	return rc.cmd.RPop(key).Result()
}

// RPush 将 value 插入到列表 key 的表尾(最右边)。
func (rc *Client) RPush(key string, value ...interface{}) (int64, error) {
	return rc.cmd.RPush(key, value...).Result()
}

// LLen 返回列表 key 的长度。
func (rc *Client) LLen(key string) (int64, error) {
	return rc.cmd.LLen(key).Result()
}
