package redis

import (
	"bluebell/mod"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math"
	"strconv"
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

// getPageSize 根据page,size,key,获取数据并排序返回date
func getPageSize(key string, page, size int64) ([]string, error) {
	//对redis进行查询，使用ZRevRange  ===> 可查询redis命令手册  https://redis.com.cn/commands.html
	start := (page - 1) * size
	end := start + size - 1
	return rds.ZRevRange(key, start, end).Result() //此处传的是索引
	/*Zrevrange 命令返回有序集中，指定区间内的成员*/
}

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

	return getPageSize(key, p.Page, p.Size)
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

// GetPostListByCommunity 根据分区，获取帖子列表并排序
func GetPostListByCommunity(p *mod.ParamsGetPostListByCommunity) ([]string, error) {
	//使用zinterstore把分区的帖子set与帖子分数的zset取交集，生成一个zset
	//针对新的zset按照之前的逻辑取数据

	//先算排序用的orderkey
	orderkey := getRedisKey(KeyPostTime)
	if p.Type == "score" {
		orderkey = getRedisKey(KeyPostScore)
	}
	//拼接要查询的社区key
	ckey := getRedisKey(KeyCommunity) + strconv.Itoa(int(p.CommunityId))
	zap.L().Debug("ckey", zap.Any("cky", ckey))
	//存储到数据的newkey
	newkey := orderkey + ":" + strconv.Itoa(int(p.CommunityId))
	zap.L().Debug("newkey", zap.Any("newkey", newkey))

	//利用缓存来减少ZInterStore的计算量
	if rds.Exists(newkey).Val() < 1 { //计算newkey类型是否存在，<1即为不存在，需要进行ZInterStore计算
		pipeline := rds.Pipeline()
		pipeline.ZInterStore(newkey, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderkey) //新集合的key、参与的集合数、参与交集的集合key
		pipeline.Expire(newkey, 60*time.Second) //设置newkey过期时间
		/*Redis 使用 expire 命令设置一个键的过期时间，到时间后 Redis 会自动删除它*/
		_, err := pipeline.Exec() //记得执行事务
		if err != nil {
			return nil, err
		}
	}
	//如果在60s内，直接返回数据
	return getPageSize(newkey, p.Page, p.Size)
}
