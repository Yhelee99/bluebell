package mysql

import (
	"bluebell/mod"
	"go.uber.org/zap"
)

func CreatPost(p *mod.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.Post_id, p.Title, p.Content, p.Author_id, p.Community_id)
	if err != nil {
		zap.L().Debug("入库失败！", zap.Error(err))
		return
	}
	return
}
