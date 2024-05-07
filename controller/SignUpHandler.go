package controller

import (
	logic "bluebell/logic/user"
	"bluebell/mod"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 用户注册接口
// @Summary 用户注册接口
// @Description 用于用户注册
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param object body mod.ParamSignUp true "用户信息"
// @Security ApiKeyAuth
// @Success 200
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	//1.获取参数，校验参数
	p := new(mod.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUpHandler校验参数失败", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, ErrorCodeInvalidParams)
			return
		} else {
			ResponseErrorWithMessage(c, ErrorCodeInvalidPassword, errs.Translate(Trans))
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
	if err := logic.SignUp(p); err != nil {
		ResponseError(c, ErrorCodeUserAlreadyExist)
		return
	}
	//3.返回响应
	ResponseSuccess(c, SuccessCode)
}
