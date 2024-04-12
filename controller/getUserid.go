package controller

import (
	"bluebell/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// getUserId 获取用户id
func getUserId(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(middleware.CtxUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
