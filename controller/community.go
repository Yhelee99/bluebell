package controller

import (
	logic "bluebell/logic/communtiy"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
controller层，处理请求参数，完成路由转发
*/

// -------社区相关

func CommunityHandler(c *gin.Context) {
	date, err := logic.GetComnunityList()
	if err != nil {
		zap.L().Error("logic.GetComnunityList()函数执行出错！", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}
