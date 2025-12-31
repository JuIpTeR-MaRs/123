package redisdb

import (
	"context"
	"fmt"
	"shop_server/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisDb *redis.Client
var RedisCtx = context.Background()

func Init() {
	fmt.Println("redis init", config.CONFIG.Redis.Host, config.CONFIG.Redis.Port)
	// 从配置中读取对应的信息
	db, _ := strconv.Atoi(config.CONFIG.Redis.Database)
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.CONFIG.Redis.Host + ":" + config.CONFIG.Redis.Port,
		Password: config.CONFIG.Redis.Password,
		DB:       db,
	})

	_, err := RedisDb.Ping(RedisCtx).Result()
	// 如果 Redis 无法连接，记录错误但不要 panic，允许服务继续运行
	if err != nil {
		fmt.Println("redis ping failed:", err)
		// 把客户端置空以便调用方检测
		RedisDb = nil
		return
	}
}
