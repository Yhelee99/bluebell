package controller

import (
	"bluebell/controller"
	logic "bluebell/logic/user"
	"bluebell/mod"
	"encoding/json"
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
	user, err := logic.SignIn(*p)
	if err != nil {
		controller.ResponseError(c, controller.ErrorCodeInvalidPassword)
		return
	}

	//userid序列化
	userIdRes, err := json.Marshal(user.UserId)

	controller.ResponseSuccess(c, gin.H{
		"UserName": user.Username,
		"UserId":   string(userIdRes), //int64范围 -2^64+1 ~ 2^64-1   前端(javascript)只能显示-2^53+1 ~ 2^53-1   会出现id值失真的问题
		"Token":    user.Token,
	})
	return
}
