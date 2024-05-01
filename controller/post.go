package controller

import (
	logic "bluebell/logic/post"
	"bluebell/mod"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

// CreatPostHandler 创建帖子
// @Summary 创建帖子
// @Description 创建帖子接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param object body mod.Post true "创建参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /createpost [post]
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
// @Summary 获取帖子详情
// @Description 根据postid获取帖子详情
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param postid path int true "帖子id" Format(int64)
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePost
// @Router /post/:id [get]
func GetPostDetail(c *gin.Context) {
	//解析数据
	pidstr := c.Param("postid")
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
// @Summary 获取帖子列表
// @Description 获取帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param page query int false "页数"
// @Param size query int false "每页数量"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePost
// @Router /posts [get]
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

// GetPostListDetermineCommunityId 可选排序方式的获取帖子列表的接口
// @Summary 获取帖子列表
// @Description 可选排序方式的获取帖子列表的接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param object body mod.ParamsGetPostList true "获取帖子的参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePost
// @Router /getpostslist [get]
func GetPostListDetermineCommunityId(c *gin.Context) {
	//1:处理参数
	p := &mod.ParamsGetPostList{
		Page: 0,
		Size: 10,
		Type: mod.OrderByScore, //尽量避免代码中出现magic string 即"time"这种写法
	} //通过定义结构体指定默认值
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("解析失败！", zap.Error(err))
		return
	}

	date, err := logic.GetPostListDetermineCommunityId(p)
	if err != nil {
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)

}

// PostVoted 帖子投票功能
// @Summary 帖子投票
// @Description 帖子投票接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param object body mod.PostVoted true "获取帖子参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /post/voted [post]
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
		var es validator.ValidationErrors
		ok := errors.As(err, &es)
		if !ok {
			zap.L().Error("eerrrroooo")
			ResponseError(c, ErrorCodeServerBusy)
			return
		}
		msg := RemoveTopStruct(es.Translate(Trans))
		ResponseErrorWithMessage(c, ErrorCodeNeedLogin, msg)
		return

	}
	ResponseSuccess(c, nil)

}

/*// GetPostListByCommunity 按社区获取帖子点赞列表
func GetPostListByCommunity(c *gin.Context) {
	//1:处理参数
	p := &mod.ParamsGetPostListByCommunity{
		ParamsGetPostListPlus: &mod.ParamsGetPostListPlus{
			Page: 0,
			Size: 10,
			Type: mod.OrderByTime,
		},
	} //通过定义结构体指定默认值
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("解析失败！", zap.Error(err))
		return
	}

	zap.L().Debug("Parms:", zap.Any("parms", p))
	date, err := logic.GetPostListByCommunity(p)
	if err != nil {
		zap.L().Error("logic.GetPostListByCommunity出错", zap.Error(err))
		ResponseError(c, ErrorCodeServerBusy)
		return
	}
	ResponseSuccess(c, date)
}*/
