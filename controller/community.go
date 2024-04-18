package controller

import (
	logic "bluebell/logic/communtiy"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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

func CommunityGetInfo(c *gin.Context) {
	//处理请求数据
	temp := c.Param("id")                     //获取到str
	id, err := strconv.ParseInt(temp, 10, 64) //base 进制  bitsize 位数
	if err != nil {
		ResponseSuccess(c, ErrorCodeInvalidParams)
		return
	}
	date, err := logic.CommunityGetInfo(id)
	if err != nil {
		zap.L().Error("查询失败！", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}
