package logic

import (
	"bluebell/mod"
	"go.uber.org/zap"
)

// GetPostListDetermineCommunityId 判断communityid是否为空，为空获取全量，不为空按社区取数据
func GetPostListDetermineCommunityId(p *mod.ParamsGetPostList) (date []*mod.ApiPost, err error) {
	if p.CommunityId == 0 {
		date, err = getPostListPlus(p)
	} else {
		date, err = getPostListByCommunity(p)
	}
	if err != nil {
		zap.L().Error("GetPostListDetermineCommunityId失败！", zap.Error(err))
		return nil, err
	}
	return
}
