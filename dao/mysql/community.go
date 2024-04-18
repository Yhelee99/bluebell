package mysql

import (
	"bluebell/mod"
	"database/sql"
	"go.uber.org/zap"
)

func GetComnunityList() (cl []*mod.CommunityList, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&cl, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("没有社区记录！")
			err = nil
		}
	}
	return
}

// 根据id查询社区详情
func CommunityGetInfo(id int64) (cinfo *mod.CommunityInfo, err error) {
	cinfo = new(mod.CommunityInfo) //new初始化对象，得到对象的指针
	sqlStr := "select community_id,community_name,introduction,creat_time from community where community_id = ?"
	err = db.Get(cinfo, sqlStr, id) //查询使用db.Get
	if err != nil {
		if err == sql.ErrNoRows {
			err = DbErrorInvalidId
		}
	}
	return cinfo, err
}
