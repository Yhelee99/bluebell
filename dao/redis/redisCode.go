package redis

import "errors"

var (
	ErrorPostRepeated = errors.New("请勿重复投票！")
	ErrorTimeOut      = errors.New("已超时，无法投票")
)
