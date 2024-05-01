package controller

import "bluebell/mod"

type _ResponsePost struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
	*mod.ApiPost
}

type _ResponseCommunity struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
	*mod.CommunityList
	*mod.CommunityInfo
}

type _ResponseUser struct {
	Code ResCode     `json:"code"` // 业务响应状态码
	Date interface{} `json:"date"` //返回数据
}
