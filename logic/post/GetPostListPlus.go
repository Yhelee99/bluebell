package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/mod"
	"go.uber.org/zap"
)

func GetPostListPlus(pl *mod.ParamsGetPostListPlus) (date []*mod.ApiPost, err error) {

	p := new(mod.Post)
	u := new(mod.User)
	c := new(mod.CommunityInfo)

	//2:从redis中获取post_id
	ids, err := redis.GetPostListPlus(pl)
	zap.L().Debug("ids:", zap.Any("ids:", ids))
	if err != nil {
		zap.L().Error("redis.GetPostListPlus查询失败！", zap.Error(err))
		return
	}
	//3:从数据库中查询对应的详细数据并按redis取出的顺序返回
	postlist, err := mysql.GetPostListPlus(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListPlus查询失败！", zap.Error(err))
		return
	}

	//根据postlist中的数据，查询帖子详情并返回
	for _, v := range postlist {
		//根据pid查帖子详情
		p, err = mysql.GetPostDetail(v.Post_id)
		if err != nil {
			zap.L().Error("GetPostDetail查库失败！", zap.Error(err))
			return
		}
		//根据用户id查询用户信息
		u, err = mysql.GetUserInfoById(v.Author_id)
		if err != nil {
			zap.L().Error("GetUserInfo查库失败！", zap.Error(err))
			return
		}

		//根据社区id查询社区详情
		c, err = mysql.CommunityGetInfo(v.Community_id)
		if err != nil {
			zap.L().Error("CommunityGetInfo查库失败！", zap.Error(err))
			return
		}

		postdetail := &mod.ApiPost{
			Author_name:   u.Username,
			Post:          p,
			CommunityInfo: c,
		}
		date = append(date, postdetail)
	}
	return
}
