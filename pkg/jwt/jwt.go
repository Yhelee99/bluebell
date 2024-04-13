package jwt

import (
	"bluebell/mod"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var ErrorTokenInvalid = errors.New("Toekn不合法！")

// GenToken 生成Token
func GenToken(username string, userid int64) (string, error) {
	c := mod.MyClaims{
		UserName: username,
		UserId:   userid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "BlueBell",
			ExpiresAt: time.Now().Add(mod.AccessTokenExpireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mod.MySecret)
}

// ParseAccessToken 解析AccessToken，验证AccessToken是否合法
func ParseAccessToken(tokenString string) (*mod.MyClaims, error) {
	var m *mod.MyClaims
	token, err := jwt.ParseWithClaims(tokenString, m, getSecret)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return m, nil
	}
	return nil, ErrorTokenInvalid
}

// getSecret 获取生成token时加签的秘钥
func getSecret(token *jwt.Token) (interface{}, error) {
	return mod.MySecret, nil
}

func GenTokenWithRefreshToken(username string, userid int64) (accessToken, refreshToken string, err error) {
	c := mod.MyClaims{
		UserName: username,
		UserId:   userid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "BlueBell",
			ExpiresAt: time.Now().Add(mod.AccessTokenExpireDuration).Unix(),
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mod.MySecret)

	mc := jwt.StandardClaims{
		Issuer:    "BlueBell",
		ExpiresAt: time.Now().Add(mod.RefreshTokenExpireDuration).Unix(),
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, mc).SignedString(mod.MySecret)
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// 判断refresh是否有效,无效直接返回
	if _, err = jwt.Parse(refreshToken, getSecret); err != nil {
		return
	}
	// 取出claims数据
	var m mod.MyClaims
	_, err = jwt.ParseWithClaims(accessToken, &m, getSecret)
	v, _ := err.(*jwt.ValidationError)
	// 当AccessToken是过期错误时且RefreshToken未过期时,才生成新AccessToken
	if v.Errors == jwt.ValidationErrorExpired {
		return GenTokenWithRefreshToken(m.UserName, m.UserId)
	}
	return
}
