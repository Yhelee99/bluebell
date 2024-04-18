package mod

import "time"

type CommunityList struct {
	Cid   int    `json:"community_id" db:"community_id"`
	Cname string `json:"cname" db:"community_name"`
}

type CommunityInfo struct {
	Cid          int64     `json:"community_id" db:"community_id"`
	Cname        string    `json:"community_name" db:"community_name"`
	Introduction string    `json:"introduction" db:"introduction"`
	Creattime    time.Time `json:"creattime" db:"creat_time"`
}
