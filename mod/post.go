package mod

import "time"

// 内存对齐
type Post struct {
	Post_id      int64     `json:"post_id" db:"post_id"`
	Author_id    int64     `json:"author_id" db:"author_id"`
	Community_id int64     `json:"community_id" db:"community_id" binding:"required"`
	Status       int32     `json:"status" db:"status"`
	Title        string    `json:"title" db:"title" binding:"required"`
	Content      string    `json:"content" db:"content" binding:"required"`
	Creattime    time.Time `json:"creat_time" db:"create_time"`
}

type ApiPost struct {
	Author_name string
	*Post
	*CommunityInfo
}
