package redis

import "errors"

const (
	KeyPerfix     = "bluebell"    // key前缀，方便业务区分
	KeyPostTime   = "post:time"   // zset 存储用户及投票时间
	KeyPostScore  = "post:score"  // zset 存储帖子点赞数
	KeyPostVotedP = "post:voted:" // zset 存储用户及投票类型 (1)点赞 (-1)点踩 (0)取消投票
)

var (
	ErrorTimeOut        = errors.New("已超时，无法投票")
	ErrorInitPostFailed = errors.New("初始化帖子失败")
)

func getKey(key string) string {
	return KeyPerfix + key
}