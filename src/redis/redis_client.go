package redis

import (
	"github.com/go-redis/redis"
	"go-study/src/document"
	"go-study/src/log"
	"go.uber.org/zap"
)

//redis单机客户端
var client *redis.Client

//redis集群客户端
var redisClusterClient *redis.ClusterClient

//redis配置
var cfg configRedis

func init() {

	//用了配置文件读取
	if err := document.Properties.Decode(&cfg); err != nil {
		log.LmdbLogger.Error("read document fail", zap.Error(err))
	}

	//判断是否为集群配置
	if cfg.Cluster {
		//ClusterClient是一个Redis集群客户机，表示一个由0个或多个底层连接组成的池。它对于多个goroutine的并发使用是安全的。
		redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Password: cfg.Password,
			Addrs:    cfg.Addrs,
		})
		//Ping
		ping, err := redisClusterClient.Ping().Result()
		log.LmdbLogger.Info("Redis Ping", zap.String("ping", ping), zap.Error(err))

	} else {
		//Redis客户端，由零个或多个基础连接组成的池。它对于多个goroutine的并发使用是安全的。
		//更多参数参考Options结构体
		client = redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password, // no password set
			DB:       cfg.DB,       // use default DB
		})
		//Ping
		pong, err := client.Ping().Result()
		log.LmdbLogger.Info("Redis Ping", zap.String("pong", pong), zap.Error(err))
	}

}

//redis配置结构体
type configRedis struct {
	Addr     string   `properties:"redis.addr"`
	Password string   `properties:"redis.password"`
	DB       int      `properties:"redis.db,default=0"`
	Cluster  bool     `properties:"redis.cluster,default=false"`
	Addrs    []string `properties:"redis.addrs,default=localhost:2181;localhost:2182"`
}
