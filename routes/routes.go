package routes

import (
	"bluebell/controller/user"
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

	//注册业务路由

	//注册
	r.POST("/signup", controller.SignUpHandler)
	//登录
	r.POST("/signin", middleware.JwtAuthMiddleware(), controller.SignInHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":     fmt.Sprintf("Welcome To %s!", mod.Conf.App.Name),
			"version": mod.Conf.App.Version,
		})
	})
	return r
}
