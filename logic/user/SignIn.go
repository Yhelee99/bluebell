package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
	"bluebell/pkg/jwt"
)

func SignIn(p mod.ParamSignIn) (accessToken string, err error) {
	user := &mod.User{
		Username: p.Username,
		Password: p.Password,
	}

	//因为Login传递的是指针，本层的user就能读到userid
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	accessToken, err = jwt.GenToken(user.Username, user.UserId)
	if err != nil {
		return "", err
	}
	return
}
