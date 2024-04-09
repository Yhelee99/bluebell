package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
)

func SignIn(p mod.ParamSignIn) error {
	user := &mod.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
