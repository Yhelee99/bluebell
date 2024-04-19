package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
)

func GetPostDetail(pid int64) (*mod.Post, error) {
	return mysql.GetPostDetail(pid)
}
