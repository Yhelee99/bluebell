package logic

import (
	mysql "bluebell/dao/mysql"
	snowflake "bluebell/pkg"
)

func SignUp() {
	//1.判断用户存不存在
	mysql.QueryUserByUsername()
	//2.生成UID
	snowflake.GetSnowflakeId()
	//3.入库
	mysql.InsertUser()
	//4.
}
