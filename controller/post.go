package controller

import (
	logic "bluebell/logic/post"
	"bluebell/mod"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// GetPostListPlus 可选排序方式的获取帖子列表的接口
func GetPostListPlus(c *gin.Context) {
	//1:处理参数
	p := &mod.ParamsGetPostListPlus{
		Page: 0,
		Size: 10,
		Type: mod.OrderByScore, //尽量避免代码中出现magic string 即"time"这种写法
	} //通过定义结构体指定默认值
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("解析失败！", zap.Error(err))
		return
	}

	date, err := logic.GetPostListPlus(p)
	if err != nil {
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)

}

// PostVoted 帖子投票功能
func PostVoted(c *gin.Context) {
	//处理参数
	p := new(mod.PostVoted)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, ErrorCodeInvalidParams)
			return
		}
		msg := RemoveTopStruct(errs.Translate(Trans)) //Trans定义的全局翻译器	RemoveTopStruct去除提示信息中的结构体名称
		ResponseErrorWithMessage(c, ErrorCodeInvalidParams, msg)
		return
	}
	userid, _ := getUserId(c)
	if err := logic.PostVoted(p, strconv.Itoa(int(userid))); err != nil {
		zap.L().Error("logic.PostVoted出错！", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)

	}
	ResponseSuccess(c, nil)

}
