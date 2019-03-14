package redis

import (
	"github.com/go-redis/redis"
	"time"
)

//set
func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if cfg.Cluster {
		return redisClusterClient.Set(key, value, expiration)
	} else {
		return client.Set(key, value, expiration)
	}
}

//get
func Get(key string) *redis.StringCmd {
	if cfg.Cluster {
		return redisClusterClient.Get(key)
	} else {
		return client.Get(key)
	}
}

//订阅
func Subscribe(channels ...string) *redis.PubSub {
	if cfg.Cluster {
		return redisClusterClient.Subscribe(channels...)
	} else {
		return client.Subscribe(channels...)
	}
}

//发布
func Publish(channel string, message interface{}) *redis.IntCmd {
	if cfg.Cluster {
		return redisClusterClient.Publish(channel, message)
	} else {
		return client.Publish(channel, message)
	}
}
