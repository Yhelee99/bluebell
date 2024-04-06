package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/logger"
	"bluebell/routes"
	"bluebell/setting"
)

func main() {

	//初始化配置
	setting.Init()

	//初始化日志
	logger.Init()

	//初始化数据库连接
	mysql.Init()

	//注册路由
	r := routes.Setup()

	//优雅关机
	controller.ShutDown(r)
}
