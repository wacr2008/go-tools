package redisx

import (
	"github.com/go-redis/redis"
	"time"
	"fmt"
)

var cli *redis.Client

const channelName = "myChannel"

func init() {
	cli = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "liu5522112",
		PoolSize:     32,
		MinIdleConns: 4,
		MaxConnAge:   10 * time.Second,
		MaxRetries:   3,
		IdleTimeout:  20 * time.Second,
	})
}

func Sub(callBack func(msg *redis.Message)) {
	pubSub := cli.Subscribe(channelName)
	go func() {
		defer pubSub.Close()
		for {
			// pubSub.ReceiveMessage 如果没有消息会阻塞，直到接收到消息，会返回消息，内部也是用for 不停的在监听
			// 外层的for 循环是必须的， 因为pubSub.ReceiveMessage 一旦接收到消息会返回，整个协程生命周期会结束。
			msg, err := pubSub.ReceiveMessage()
			// TODO 一般不会出现错误，遇到错误，应该预警
			if err != nil {
				fmt.Println(err)
			}
			// 回调方法， 或者用 chan ，在调用Sub方法的时候 传入 chan，然后 select chan
			callBack(msg)
		}
	}()

	time.Sleep(10 * time.Minute)
}

func Pub() {
	intCmd := cli.Publish(channelName, "hello1")
	if intCmd.Err() != nil {
		fmt.Println(intCmd.Err())
		return
	}
	if intCmd.Val() == 0 {
		// TODO 等于零表示没有订阅者，也就是没有任何订阅者接收到消息，这个地方可能需要重试，或者设置重试次数和时间
		// TODO 一般重试时间会随着次数的增加 而增加， 比如： 第一次立刻重试 第二次：1分钟后重试 第三次：5分钟后重试 第四次：10分钟...
		fmt.Println(intCmd.Val())
	}
	fmt.Println(intCmd.String())
}
