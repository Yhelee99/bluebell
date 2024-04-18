package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
)

func GetComnunityList() (date []*mod.CommunityList, err error) {

	return mysql.GetComnunityList()

}

func CommunityGetInfo(id int64) (*mod.CommunityInfo, error) {
	return mysql.CommunityGetInfo(id)
}
