package routes

import (
	"bluebell/controller/user"
	"bluebell/logger"
	"bluebell/mod"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Setup() http.Handler {
	r := gin.New()
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(zap.L(), true))

	//注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":     fmt.Sprintf("Welcome To %s!", mod.Conf.App.Name),
			"version": mod.Conf.App.Version,
		})
	})
	return r
}
