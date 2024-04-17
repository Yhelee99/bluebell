package mod

type CommunityList struct {
	Cid   int    `json:"cid" db:"community_id"`
	Cname string `json:"cname" db:"community_name"`
}
