package cache

import (
	"common/configs"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var RedisContext = context.Background()

func InitRedis() {
	addr := configs.GetConfig().Redis.Addr
	password := configs.GetConfig().Redis.Password
	port := configs.GetConfig().Redis.Port
	db := configs.GetConfig().Redis.Db

	// 创建一个 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", addr, port), // Redis 服务器地址
		Password:     password,                         // 密码设置
		DB:           db,                               // 默认数据库
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	RedisClient = client
}
