package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
	"bluebell/pkg/jwt"
)

func SignIn(p mod.ParamSignIn) (user *mod.User, err error) {
	user = &mod.User{
		Username: p.Username,
		Password: p.Password,
	}

	//因为Login传递的是指针，本层的user就能读到userid
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	var accessToken string
	accessToken, err = jwt.GenToken(user.Username, user.UserId)
	user.Token = accessToken
	if err != nil {
		return user, err
	}
	return
}
