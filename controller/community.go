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

// CommunityHandler 获取社区列表
// @Summary 获取社区列表
// @Description 获取社区列表接口
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunity
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	date, err := logic.GetComnunityList()
	if err != nil {
		zap.L().Error("logic.GetComnunityList()函数执行出错！", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}

// CommunityGetInfo 根据社区id查询社区详情
// @Summary 获取社区列表
// @Description 根据社区id查询社区详情
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param id path int false "社区ID,可不传,不传获取全量" Format(int64)
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunity
// @Router /community/:id [get]
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
