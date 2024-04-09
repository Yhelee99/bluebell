package mysql

import (
	"bluebell/mod"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
)

/*
	把每一步数据库操作封装成函数
	等待logic层调用
*/

const secret = "Yhelee99"

// InsertUser 将用户数据插入数据库
func InsertUser(u mod.User) error {
	//1.执行SQL入库
	sqlString := "insert into user(user_id,username,password) values (?,?,?)"
	u.Password = encryptPassword(u.Password)
	if _, err := db.Exec(sqlString, u.UserId, u.Username, u.Password); err != nil {
		return err
	}
	return nil
}

// CheckUserExist 查库，判断用户是否存在
func CheckUserExist(username string) (err error) {
	sqlString := "select count(username) from user where username = ?"
	var count int
	if err = db.Get(&count, sqlString, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在！")
	}
	return
}

// encryptPassword 给密码加密
func encryptPassword(psw string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(psw)))
}

// checkPassword 校验密码
func Login(u *mod.User) (err error) {
	sqlString := "select password from user where username = ?"
	var temp string
	err = db.Get(&temp, sqlString, u.Username)
	if err != nil {
		return err
	} else if err == sql.ErrNoRows {
		zap.L().Debug("用户不存在", zap.String("username", u.Username), zap.Error(err))
		return errors.New("用户不存在！")
	}
	if temp != encryptPassword(u.Password) {
		return errors.New("密码错误！")
	}
	return
}
