package redis

import (
	"github.com/go-redis/redis"
	"go-study/src/log"
	"go.uber.org/zap"
)

func DemoRedis() {
	log.LmdbLogger.Info("------redis---------")
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		log.LmdbLogger.Error("client.Set fail", zap.Error(err))
	}

	val, err := Get("key").Result()
	if err != nil {
		log.LmdbLogger.Error("client.Get fail", zap.Error(err))
	}
	log.LmdbLogger.Info("client.Get", zap.String("val", val))

	val2, err := Get("key2").Result()
	if err == redis.Nil {
		log.LmdbLogger.Info("key2 does not exist")
	} else if err != nil {
		log.LmdbLogger.Error("client.Get fail", zap.Error(err))
	} else {
		log.LmdbLogger.Info("client.Get", zap.String("val2", val2))
	}
	//订阅
	sub := Subscribe("channel")
	//新协程
	go func() {
		for {
			//发布
			pub := Publish("channel", "message")
			if pub.Err() != nil {
				log.LmdbLogger.Error("Publish err", zap.String("message", "message"), zap.Error(pub.Err()))
			} else {
				log.LmdbLogger.Info("Publish msg", zap.String("message", "message"))
			}
		}
	}()

	//从订阅获取信息，获取一次则程序结束
	msg, err := sub.ReceiveMessage()
	if err != nil {
		log.LmdbLogger.Error("Subscribe err", zap.Error(err))
	} else {
		log.LmdbLogger.Info("Subscribe msg", zap.String("msg", msg.String()))
	}

}
