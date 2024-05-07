package redis

const (
	KeyPerfix     = "bluebell:"   // key前缀，方便业务区分
	KeyPostTime   = "post:time"   // zset 存储用户及投票时间
	KeyPostScore  = "post:score"  // zset 存储帖子点赞数
	KeyPostVotedP = "post:voted:" // zset 存储用户及投票类型 (1)点赞 (-1)点踩 (0)取消投票
	KeyCommunity  = "community:"
)

func getRedisKey(key string) string {
	return KeyPerfix + key
}
