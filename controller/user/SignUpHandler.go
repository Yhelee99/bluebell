package controller

import (
	"bluebell/controller"
	logic "bluebell/logic/user"
	"bluebell/mod"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数，校验参数
	p := new(mod.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUpHandler校验参数失败", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "请求参数有误！",
				"code": 100001,
				"err":  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "请求参数有误！",
				"code": 100001,
				"err":  controller.RemoveTopStruct(errs.Translate(controller.Trans)),
			})
		}

		return
	}
	//参数业务校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg":  "请求参数有误！",
	//		"code": 100001,
	//	})
	//	return
	//}
	//2.业务处理
	logic.SignUp()
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功!",
	})
}
