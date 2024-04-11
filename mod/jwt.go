package mod

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("Fake it,until you make it.")

type MyClaims struct {
	UserName string
	UserId   int64
	jwt.StandardClaims
}
