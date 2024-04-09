package mysql

import (
	"bluebell/mod"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mod.Conf.Mysql.Username,
		mod.Conf.Mysql.Password,
		mod.Conf.Mysql.Host,
		mod.Conf.Mysql.Port,
		mod.Conf.Mysql.Dbname,
	)

	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Debug("连接数据库失败！")
		panic(err)
	}
}
