package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func DemoRedis() {
	fmt.Println("------redis---------")
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	//订阅
	sub := Subscribe("channel")
	//新协程
	go func() {
		for {
			//发布
			pub := Publish("channel", "message")
			if pub.Err() != nil {
				fmt.Println("Publish err", "message")
			} else {
				fmt.Println("Publish msg", "message")
			}
		}
	}()

	//从订阅获取信息，获取一次则程序结束
	msg, err := sub.ReceiveMessage()
	if err != nil {
		fmt.Println("Subscribe err", err)
	} else {
		fmt.Println("Subscribe msg", msg)
	}

}
