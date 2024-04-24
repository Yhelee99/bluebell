package mysql

import (
	"bluebell/mod"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

// CreatPost 创建帖子
func CreatPost(p *mod.Post) (err error) {
	sqlStr := `insert into post (post_id, title, content, author_id, community_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.Post_id, p.Title, p.Content, p.Author_id, p.Community_id)
	if err != nil {
		zap.L().Debug("入库失败！", zap.Error(err))
		return
	}
	return
}

// GetPostDetail 根据post_id查询帖子详情
func GetPostDetail(pid int64) (p *mod.Post, err error) {
	p = new(mod.Post)
	sqlStr := `select post_id,author_id,community_id,title,content,create_time from post where post_id = ?`
	err = db.Get(p, sqlStr, pid)
	return
}

// GetPostList 获取帖子详情
func GetPostList(page, size int64) (posts []*mod.Post, err error) {
	sqlStr := `select post_id,author_id,community_id,title,content,create_time from post order by create_time desc limit ?,?`
	posts = make([]*mod.Post, 0, 2)
	if err = db.Select(&posts, sqlStr, (page-1)*size, size); err != nil {
		zap.L().Error("GetPostList查库失败！", zap.Error(err))
		return
	}
	return
}

// GetPostListPlus 获取帖子详情，可选排序方式
func GetPostListPlus(ids []string) (postlist []*mod.Post, err error) {

	sqlStr := `select post_id,author_id,community_id,title,content,create_time from post 
    where post_id in (?) order by FIND_IN_SET(post_id,?)` //FIND_IN_SET函数:查询id在给定id集合的数据并维持给定id集合的顺序	FIND_IN_SET(search_str, str_list)

	//sqlx使用：https://www.liwenzhou.com/posts/Go/sqlx/
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ",")) //args将是一个切片，包含需要传递给数据库驱动的参数值，即[1, 2, 3, 4, 5]
	if err != nil {
		return nil, err
	}

	// Rebind 用于将查询语句中的占位符（?）重新绑定为适合特定数据库驱动的占位符
	query = db.Rebind(query)
	err = db.Select(&postlist, query, args...) // !!!select要传args...
	return
}
