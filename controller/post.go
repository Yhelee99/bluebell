package controller

import (
	logic "bluebell/logic/post"
	"bluebell/mod"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatPostHandler(c *gin.Context) {
	//1：处理参数
	p := new(mod.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("ShouldBindJSON解析错误！", zap.Error(err))
		ResponseError(c, ErrorCodeInvalidParams)
		return
	}
	//2:业务处理
	uid, err := getUserId(c) //从c获取当前用户id
	if err != nil {
		zap.L().Debug("用户未登录！", zap.Error(err))
		ResponseError(c, ErrorCodeNeedLogin)
		return
	}
	p.Author_id = uid
	if err = logic.CreatPost(p); err != nil { //创建帖子
		zap.L().Debug("创建帖子失败！", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)
		return
	}

	//3：返回响应
	ResponseSuccess(c, nil)
	return
}
