package middle

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(c *gin.Context) {
	//获取请求头参数
	authHeader := c.Request.Header.Get("Authorization")

	//先判断是否为空
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": controller.ErrorCodeInvalidParams,
			"msg":  "请求头Token为空！",
		})
		c.Abort() //立即停止请求并返回响应
		return
	}

	//处理请求头数据 SplitN返回切片
	post := strings.SplitN(authHeader, " ", 2)
	if !(len(post) != 2 && post[0] != "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code": controller.ErrorCodeInvalidParams,
			"msg":  "Token不合法！",
		})
		c.Abort()
		return
	}

	//验证token
	mc, err := jwt.ParseToken(post[1])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": controller.ErrorCodeInvalidParams,
			"msg":  "无效的Token！",
		})
	}

	//将当前的用户id保存到上下文c中
	c.Set("UserId", mc.UserId)
	c.Next() //验证通过，进行下个响应，后续处理可以通过c.Get("UserId")来获取当前用户信息
}
