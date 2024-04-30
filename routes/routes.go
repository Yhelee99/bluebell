package routes

import (
	"bluebell/controller"
	user "bluebell/controller/user"
	"bluebell/logger"
	"bluebell/middleware"
	"bluebell/mod"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Setup(mode string) http.Handler {

	//判断是否是开发模式，如果是不显示gin框架的debug信息
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(zap.L(), true))

	//swagger页面渲染
	r.GET("/swagger/*any")

	//注册业务路由
	//登录
	r.POST("/signin", user.SignInHandler)

	//注册
	r.POST("/signup", user.SignUpHandler)

	// api/v1
	{
		//创建路由组
		v1 := r.Group("/api/v1")
		//使用中间件
		v1.Use(middleware.JwtAuthMiddleware())

		{
			//查看社区列表
			v1.GET("/community", controller.CommunityHandler)

			//根据社区id获取详情
			v1.GET("/community/:id", controller.CommunityGetInfo)

			//帖子发表
			v1.POST("/createpost", controller.CreatPostHandler)

			//根据帖子id查询帖子详情
			v1.GET("/post/:id", controller.GetPostDetail)

			//获取帖子列表
			v1.GET("/posts", controller.GetPostList)
			//获取帖子列表（可按点赞量或时间排序）
			v1.GET("/getpostslist", controller.GetPostListDetermineCommunityId)

			//帖子点赞功能
			v1.POST("/post/voted", controller.PostVoted)

			//根据社区id获取帖子列表并排序
			//v1.GET("/postlistByCommunity", controller.GetPostListByCommunity)

		}
	}

	//返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": 404,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":     fmt.Sprintf("Welcome To %s!", mod.Conf.App.Name),
			"version": mod.Conf.App.Version,
		})
	})
	return r
}
