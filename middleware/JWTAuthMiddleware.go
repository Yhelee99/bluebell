package middleware

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		//获取请求头参数
		authHeader := c.Request.Header.Get("Authorization")

		//先判断是否为空
		if authHeader == "" {
			controller.ResponseError(c, controller.ErrorCodeNeedLogin)
			c.Abort() //立即停止请求并返回响应
			return
		}

		//处理请求头数据 SplitN返回切片
		post := strings.SplitN(authHeader, " ", 2)
		if !(len(post) != 2 && post[0] != "Bearer") {
			controller.ResponseError(c, controller.ErrorCodInvalidAuth)
			c.Abort()
			return
		}

		//验证token
		mc, err := jwt.ParseToken(post[1])
		if err != nil {
			controller.ResponseError(c, controller.ErrorCodInvalidAuth)
			return
		}

		//将当前的用户id保存到上下文c中
		//c.Set("UserId", mc.UserId)	`项目中很多地方要用的参数，定义为常量`
		c.Set(controller.CtxUserIdKey, mc.UserId)
		c.Next() //验证通过，进行下个响应，后续处理可以通过c.Get("UserId")
	}
}
