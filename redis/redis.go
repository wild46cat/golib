package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/wild46cat/golib/config"
	"log"
	"sync"
	"time"
)

var redisClient RedisClient

type RedisClient struct {
	*redis.ClusterClient
}

var onceInit sync.Once
var clusterClient *redis.ClusterClient

func NewRedisClient() *RedisClient {
	onceInit.Do(func() {
		log.Println("Init redis cluster ...")
		Redis := config.Configuration.Redis
		addrs := Redis.Addrs
		password := Redis.Password
		poolSize := Redis.PoolSize
		dialTimeout := Redis.DialTimeout
		readTimeout := Redis.ReadTimeout
		writeTimeout := Redis.WriteTimeout
		minIdleConns := Redis.MinIdleConns
		poolTimeout := Redis.PoolTimeOut
		idleTimeout := Redis.IdleTimeOut
		log.Printf("redis addr is :%s", addrs)
		clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:          addrs,
			ReadOnly:       false,
			RouteByLatency: false,
			RouteRandomly:  false,
			Password:       password,
			DialTimeout:    time.Duration(dialTimeout) * time.Millisecond,
			ReadTimeout:    time.Duration(readTimeout) * time.Millisecond,
			WriteTimeout:   time.Duration(writeTimeout) * time.Millisecond,
			PoolSize:       poolSize,
			MinIdleConns:   minIdleConns,
			PoolTimeout:    time.Duration(poolTimeout) * time.Millisecond,
			IdleTimeout:    time.Duration(idleTimeout) * time.Millisecond,
			TLSConfig:      nil,
		})
		pong, err := clusterClient.Ping().Result()
		if err != nil {
			fmt.Println(pong)
			log.Fatalf("connect to redis %v and ping receive %v with err %v", addrs, pong, err)
		}
	})
	return &RedisClient{clusterClient}
}

func (r *RedisClient) Lock(key string, requestId interface{}, expiration time.Duration) bool {
	return r.SetNX(key, requestId, expiration).Val()
}

func (r *RedisClient) Unlock(key string, requestId interface{}) bool {
	script := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	res := r.Eval(script, []string{key}, requestId).Val()
	ok := res.(int64) == 1
	return ok
}
