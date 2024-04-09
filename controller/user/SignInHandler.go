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

func SignInHandler(c *gin.Context) {

	//1:获取参数，校验参数
	p := new(mod.ParamSignIn)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Debug("解析json数据到请求体失败！", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "登录失败，参数不正确！",
				"code": 100002,
				"err":  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "登录失败，参数不正确！",
				"code": 100002,
				"err":  controller.RemoveTopStruct(errs.Translate(controller.Trans)),
			})
			return
		}
	}

	//2:业务处理
	//3:返回响应
	if err := logic.SignIn(*p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功！",
	})

}
