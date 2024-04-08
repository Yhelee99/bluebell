package logic

import (
	dao "bluebell/dao/mysql"
	"bluebell/mod"
	snowflake "bluebell/pkg"
)

func SignUp(p *mod.ParamSignUp) error {
	//1.判断用户存不存在
	if err := dao.CheckUserExist(p.Username); err != nil {
		return err
	}
	//2.生成UID
	userid := snowflake.GetSnowflakeId()

	//3.生成用户
	u := &mod.User{
		userid,
		p.Username,
		p.Username,
	}

	//4.入库
	if err := dao.InsertUser(*u); err != nil {
		return err
	}
	return nil
}
