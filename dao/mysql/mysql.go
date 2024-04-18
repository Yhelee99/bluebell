package mysql

import (
	"bluebell/mod"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", //parseTime=true&loc=Local,设置时间自动转化，时区为本机时区
		mod.Conf.Mysql.Username,
		mod.Conf.Mysql.Password,
		mod.Conf.Mysql.Host,
		mod.Conf.Mysql.Port,
		mod.Conf.Mysql.Dbname,
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("连接数据库失败！")
		return
	}
	return
}
