package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
	"go.uber.org/zap"
)

func GetPostDetail(pid int64) (date *mod.ApiPost, err error) {

	p := new(mod.Post)
	p, err = mysql.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("GetPostDetail查库失败！", zap.Error(err))
		return
	}
	//根据用户id查询用户信息
	u, err := mysql.GetUserInfoById(p.Author_id)
	if err != nil {
		zap.L().Error("GetUserInfo查库失败！", zap.Error(err))
		return
	}

	//根据社区id查询社区详情
	c, err := mysql.CommunityGetInfo(p.Community_id)
	if err != nil {
		zap.L().Error("CommunityGetInfo查库失败！", zap.Error(err))
		return
	}
	date = &mod.ApiPost{
		Author_name:   u.Username,
		Post:          p,
		CommunityInfo: c,
	}
	return
}
