package mod

import "time"

// 内存对齐
type post struct {
	Post_id      int64     `json:"id" db:"post_id"`
	Author_id    int64     `json:"author_id" db:"author_id"`
	Community_id int64     `json:"community_id" db:"community_id"`
	Status       int32     `json:"status" db:"status"`
	Title        string    `json:"title" db:"title"`
	Content      string    `json:"content" db:"content"`
	Creattime    time.Time `json:"creat_time" db:"creat_time"`
}
