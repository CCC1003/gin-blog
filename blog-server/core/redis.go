package core

import (
	"Blog/global"
	"context"
	"github.com/go-redis/redis"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}
func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       0,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Logger.Errorf("redis连接失败 %s", redisConf.Addr())
		return nil
	}
	return rdb
}
