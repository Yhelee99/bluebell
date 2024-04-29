package redis

import (
	"bluebell/mod"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math"
	"time"
)

/*
投票功能实现：
direction=1
	1、原来投过反对票	-1 => 1   432*2
	2、未投过票		0 => 1	  432
direction=0
	1、原来投过反对票  -1 => 0   432
	2、原来投过赞成票  1 ->0     -432
direction=-1
	1、未投过票		0 => -1   -432
	2、原来投过赞成票 	1 => -1   -432*2

判断帖子是否过期，超过7天不可投票

一票=432分
一天86400s，投200票可以续一天榜首

*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	oneTicket        = 432
)

// VoteForPost 进行帖子投票并存入Redis中
func VoteForPost(userid string, postid string, ticketType float64) error {

	//判断可否投票
	//获取发布时间
	postTime := rds.ZScore(getRedisKey(KeyPostTime), postid).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorTimeOut
	}

	//更改分数
	//先查当前用户是否投过票
	oldticket := rds.ZScore(getRedisKey(KeyPostVotedP+postid), userid).Val()
	if oldticket == ticketType {
		return ErrorPostRepeated
	}
	//确定差值
	diff := math.Abs(oldticket - ticketType)
	//确定运算符
	var oc float64
	if oldticket > ticketType {
		oc = -1 //oc:operational character
	} else {
		oc = 1
	}
	//变更
	pipeline := rds.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScore), oc*diff*oneTicket, postid)

	//记录用户为改帖子投票的记录
	if ticketType == 0 { //取消投票
		pipeline.ZRem(getRedisKey(KeyPostVotedP+postid), userid)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedP+postid), redis.Z{
			Score:  ticketType, //新的投票类型
			Member: userid,     //记录用户id
		})
	}
	_, err := pipeline.Exec()
	return err
}

// GetPostID 根据请求参数中的page和size获取选中的PostID
func GetPostID(p *mod.ParamsGetPostListPlus) ([]string, error) {
	//确定要查的key
	//zap.L().Debug("p传入的", zap.Any("key", p.Type))

	key := getRedisKey(KeyPostTime)
	if p.Type == "score" {
		key = getRedisKey(KeyPostScore)
	}

	zap.L().Debug("确定要查的key", zap.Any("key", key))

	//对redis进行查询，使用ZRevRange  ===> 可查询redis命令手册  https://redis.com.cn/commands.html
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return rds.ZRevRange(key, start, end).Result() //此处传的是索引
}

func GetUserPostType(ids []string) (date []int64, err error) {

	/*	//拼接key			---可以实现功能但不好
		date = make([]int64, 0, len(ids))
		for _, postid := range ids {
			key := getRedisKey(KeyPostVotedP + postid)
			v := rds.ZCount(key, "1", "1").Val()
			date = append(date, v)
		}	*/

	//使用Pipeline降低
	/*
		Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。
		区别于一个接一个地执行100个命令，你可以将这些命令放入 pipeline 中，然后使用1次读写操作像执行单个命令一样执行它们。
		这样做的好处是节省了执行命令的网络往返时间（RTT）
	*/

	//拼接redis命令
	pipeline := rds.Pipeline()
	for _, v := range ids {
		key := getRedisKey(KeyPostVotedP + v)
		pipeline.ZCount(key, "1", "1")
	}

	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	zap.L().Debug("---cmders", zap.Any("cmders", cmders))

	date = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val() //cmder转化为redis.IntCmd，调用Val()函数获取int64值
		date = append(date, v)
	}

	return
}
