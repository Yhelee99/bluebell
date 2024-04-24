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

func VoteForPost(userid string, postid string, ticketType float64) error {

	//判断可否投票
	//获取发布时间
	postTime := rds.ZScore(getKey(KeyPostTime), postid).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorTimeOut
	}

	//更改分数
	//先查当前用户是否投过票
	oldticket := rds.ZScore(getKey(KeyPostVotedP+postid), userid).Val()
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
	pipeline.ZIncrBy(getKey(KeyPostScore), oc*diff*oneTicket, postid)

	//记录用户为改帖子投票的记录
	if ticketType == 0 { //取消投票
		pipeline.ZRem(getKey(KeyPostVotedP+postid), userid)
	} else {
		pipeline.ZAdd(getKey(KeyPostVotedP+postid), redis.Z{
			Score:  ticketType, //新的投票类型
			Member: userid,     //记录用户id
		})
	}
	_, err := pipeline.Exec()
	return err
}

func GetPostListPlus(p *mod.ParamsGetPostListPlus) ([]string, error) {
	//确定要查的key
	zap.L().Debug("p传入的", zap.Any("key", p.Type))

	key := getKey(KeyPostTime)
	if p.Type == "score" {
		key = getKey(KeyPostScore)
	}

	zap.L().Debug("确定要查的key", zap.Any("key", key))

	//对redis进行查询，使用ZRevRange  ===> 可查询redis命令手册  https://redis.com.cn/commands.html
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return rds.ZRevRange(key, start, end).Result() //此处传的是索引
}
