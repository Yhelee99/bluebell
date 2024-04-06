package controller

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ShutDown(r http.Handler) {

	sv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		err := sv.ListenAndServe()
		if err != nil {
			zap.L().Fatal("监听服务失败！", zap.Error(err))
			return
		}
	}()

	c := make(chan os.Signal, 1)
	<-c
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	zap.L().Info("开始准备关机！")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := sv.Shutdown(ctx); err != nil {
		zap.L().Fatal("服务关闭错误！", zap.Error(err))
		return
	}
	zap.L().Info("服务已关机！")
}
