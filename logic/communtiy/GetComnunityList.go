package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
)

func GetComnunityList() (date []*mod.CommunityList, err error) {

	return mysql.GetComnunityList()

}
