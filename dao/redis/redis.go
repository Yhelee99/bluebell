package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rds *redis.Client

func Init() {

	rds = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.poolsize"),
	})

	_, err := rds.Ping().Result()
	if err != nil {
		zap.L().Error("redis连接失败！")
		return
	}
	return
}
