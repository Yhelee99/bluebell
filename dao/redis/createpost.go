package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

func CreatePost(postId int64) (err error) {

	//生成一个事务
	pipeline := rds.TxPipeline()

	//初始化帖子时间
	pipeline.ZAdd(getKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	//初始化帖子分数
	pipeline.ZAdd(getKey(KeyPostScore), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	if _, err = pipeline.Exec(); err != nil {
		zap.L().Debug("createpost事务执行失败！", zap.Error(err))
	}
	return
}