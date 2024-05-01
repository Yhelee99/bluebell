package mod

import "time"

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

// 内存对齐
type Post struct {
	Post_id      int64     `json:"post_id" db:"post_id"`                              //帖子id
	Author_id    int64     `json:"author_id" db:"author_id"`                          //作者id
	Community_id int64     `json:"community_id" db:"community_id" binding:"required"` //社区id
	Status       int32     `json:"status" db:"status"`                                //帖子状态
	Title        string    `json:"title" db:"title" binding:"required"`               //帖子标题
	Content      string    `json:"content" db:"content" binding:"required"`           //帖子内容
	Creattime    time.Time `json:"creat_time" db:"create_time"`                       //发帖时间
}

type ApiPost struct {
	Author_name    string //作者名称
	PostApproveNum int64  `json:"post_approve_num"`       //帖子被点赞数
	CommunityId    int64  `json:"community_id,omitempty"` //社区id
	*Post
	*CommunityInfo
}

type PostVoted struct {
	PostId    string `json:"post_id" binding:"required"`       // 帖子id，必传
	Direction int8   `json:"direction" binding:"oneof=1 -1 0"` // 投票类型，1赞同，-1不赞同，0取消投票
	/*oneof 判断值在一个范围里	required如果传0，会判断为空*/
	/*Userid从c中获取*/
}
