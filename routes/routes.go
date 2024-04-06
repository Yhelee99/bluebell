package routes

import (
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":     fmt.Sprintf("Welcome To %s!", mod.Conf.App.Name),
			"version": mod.Conf.App.Version,
		})
	})
	return r
}
