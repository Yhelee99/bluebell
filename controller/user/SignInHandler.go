package controller

import (
	"bluebell/controller"
	logic "bluebell/logic/user"
	"bluebell/mod"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignInHandler(c *gin.Context) {

	//1:获取参数，校验参数
	p := new(mod.ParamSignIn)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Debug("解析json数据到请求体失败！", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			controller.ResponseError(c, controller.ErrorCodeInvalidParams)
			return
		} else {
			controller.ResponseErrorWithMessage(c, controller.ErrorCodeInvalidParams, errs.Translate(controller.Trans))
			return
		}
	}

	//2:业务处理
	//3:返回响应
	if err := logic.SignIn(*p); err != nil {
		controller.ResponseError(c, controller.ErrorCodeInvalidPassword)
		return
	}
	controller.ResponseSuccess(c, controller.SuccessCode)
	return
}
