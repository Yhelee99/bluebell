package mod

import "time"

type CommunityList struct {
	Cid   int    `json:"community_id" db:"community_id"` //社区ID
	Cname string `json:"cname" db:"community_name"`      //社区名称
}

type CommunityInfo struct {
	Cid          int64     `json:"community_id" db:"community_id"`     //社区id
	Cname        string    `json:"community_name" db:"community_name"` //社区名称
	Introduction string    `json:"introduction" db:"introduction"`     //描述信息
	Creattime    time.Time `json:"creattime" db:"creat_time"`          //创建时间
}
