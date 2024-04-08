package mysql

import (
	"bluebell/mod"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

/*
	把每一步数据库操作封装成函数
	等待logic层调用
*/

const secret = "Yhelee99"

// InsertUser 将用户数据插入数据库
func InsertUser(u mod.User) error {
	//1.执行SQL入库
	sql := "insert into user(user_id,username,password) values (?,?,?)"
	u.Password = encryptPassword(u.Password)
	if _, err := db.Exec(sql, u.UserId, u.Username, u.Password); err != nil {
		return err
	}
	return nil
}

// CheckUserExist 查库，判断用户是否存在
func CheckUserExist(username string) error {
	sql := "select count(username) from user where username = ?"
	var count int
	if err := db.Get(&count, sql, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在！")
	}
	return nil
}

// encryptPassword 给密码加密
func encryptPassword(psw string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(psw)))
}
