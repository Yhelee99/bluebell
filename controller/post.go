package controller

import (
	logic "bluebell/logic/post"
	"bluebell/mod"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatPostHandler 创建帖子
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

// GetPostDetail 获取帖子详情
func GetPostDetail(c *gin.Context) {
	//解析数据
	pidstr := c.Param("id")
	pid, err := strconv.ParseInt(pidstr, 10, 64)
	if err != nil {
		zap.L().Error("解析参数为空！", zap.Error(err))
		ResponseError(c, ErrorCodeInvalidParams)
		return
	}
	//查询数据
	date, err := logic.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("查库失败！", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}

// GetPostList 获取帖子列表
func GetPostList(c *gin.Context) {
	//获取数据
	//获取分页信息
	page, size := GetPageSize(c)

	date, err := logic.GetPostList(page, size)
	//返回
	if err != nil {
		zap.L().Error("获取帖子列表失败！", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)

}
