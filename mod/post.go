package mod

import "time"

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

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
	Author_name    string
	PostApproveNum int64 `json:"post_approve_num"`
	*Post
	*CommunityInfo
}

type PostVoted struct {
	//Userid从c中获取
	PostId    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction" binding:"oneof=1 -1 0"` // oneof 判断值在一个范围里	required如果传0，会判断为空
}
