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
