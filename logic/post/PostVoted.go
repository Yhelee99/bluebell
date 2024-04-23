package logic

import (
	"bluebell/dao/redis"
	"bluebell/mod"
)

func PostVoted(p *mod.PostVoted, userid string) error {
	return redis.VoteForPost(userid, p.PostId, float64(p.Direction))
}
