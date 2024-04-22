package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
	"go.uber.org/zap"
)

func GetPostList(page, size int64) (date []*mod.ApiPost, err error) {

	p := new(mod.Post)
	u := new(mod.User)
	c := new(mod.CommunityInfo)

	post, err := mysql.GetPostList(page, size)
	if err != nil {
		return
	}
	for _, posts := range post {

		//根据pid查帖子详情
		p, err = mysql.GetPostDetail(posts.Post_id)
		if err != nil {
			zap.L().Error("GetPostDetail查库失败！", zap.Error(err))
			return
		}

		//根据用户id查询用户信息
		u, err = mysql.GetUserInfoById(p.Author_id)
		if err != nil {
			zap.L().Error("GetUserInfo查库失败！", zap.Error(err))
			return
		}

		//根据社区id查询社区详情
		c, err = mysql.CommunityGetInfo(p.Community_id)
		if err != nil {
			zap.L().Error("CommunityGetInfo查库失败！", zap.Error(err))
			return
		}
		postDetail := &mod.ApiPost{
			Author_name:   u.Username,
			Post:          p,
			CommunityInfo: c,
		}
		date = append(date, postDetail)
	}
	return
}
