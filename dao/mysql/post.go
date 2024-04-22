package mysql

import (
	"bluebell/mod"
	"fmt"
	"go.uber.org/zap"
)

func CreatPost(p *mod.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?,?,?,?,?)`
	fmt.Printf("CreatPost=======%v\n", p.Author_id)
	_, err = db.Exec(sqlStr, p.Post_id, p.Title, p.Content, p.Author_id, p.Community_id)
	if err != nil {
		zap.L().Debug("入库失败！", zap.Error(err))
		return
	}
	return
}

func GetPostDetail(pid int64) (p *mod.Post, err error) {
	p = new(mod.Post)
	sqlStr := `select post_id,author_id,community_id,title,content,create_time from post where post_id = ?`
	err = db.Get(p, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (posts []*mod.Post, err error) {
	sqlStr := `select post_id,author_id,community_id,title,content,create_time from post limit ?,?`
	posts = make([]*mod.Post, 0, 2)
	if err = db.Select(&posts, sqlStr, (page-1)*size, size); err != nil {
		zap.L().Error("GetPostList查库失败！", zap.Error(err))
		return
	}
	return
}
