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

// @title Bluebell项目
// @version 1.0
// @description 一个后端项目
// @termsOfService https://github.com/Yhelee99

// @contact.name Yhelee
// @contact.url https://github.com/Yhelee99
// @contact.email yhelee99@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {

	//初始化配置
	setting.Init()

	//初始化日志
	logger.Init()

	//初始化数据库连接
	dbConfig := &mod.Mysql{
		Port:     mod.Conf.Mysql.Port,
		Password: mod.Conf.Mysql.Password,
		Host:     mod.Conf.Mysql.Host,
		Dbname:   mod.Conf.Mysql.Dbname,
		Username: mod.Conf.Mysql.Username,
	}
	mysql.Init(dbConfig)

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
