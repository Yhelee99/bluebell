package jwt

import (
	"bluebell/mod"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenToken 生成Token
func GenToken(username string, userid int64) (string, error) {
	c := mod.MyClaims{
		UserName: username,
		UserId:   userid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "BlueBell",
			ExpiresAt: time.Now().Add(mod.TokenExpireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mod.MySecret)
}

// ParseToken 解析Token，验证Token是否合法
func ParseToken(tokenString string) (*mod.MyClaims, error) {
	var mc = new(mod.MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, getSecret)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("token不合法！")
}

// getSecret 获取生成token时加签的秘钥
func getSecret(token *jwt.Token) (interface{}, error) {
	return mod.MySecret, nil
}
