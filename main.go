package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/mod"
	snowflake "bluebell/pkg"
	"bluebell/routes"
	"bluebell/setting"
	"go.uber.org/zap"
)

func main() {

	//初始化配置
	setting.Init()

	//初始化日志
	logger.Init()

	//初始化数据库连接
	mysql.Init()

	//初始化redis
	redis.Init()

	//初始化翻译器
	if err := controller.InitTrans("zh"); err != nil {
		zap.L().Error("翻译器初始化失败！", zap.Error(err))
		return
	}

	//初始化雪花函数
	if err := snowflake.Init(mod.Conf.Snowflake.StartTime, mod.Conf.Snowflake.MachineId); err != nil {
		zap.L().Debug("雪花算法初始化失败！")
		zap.Error(err)
		return
	}
	//注册路由
	r := routes.Setup(mod.Conf.App.Mode)

	//优雅关机
	controller.ShutDown(r)
}
