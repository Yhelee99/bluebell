package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseData struct {
	Code    ResCode     `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"` //json操作小技巧，omitempty忽略空值
}

// ResponseSuccess 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    SuccessCode,
		Message: SuccessCode.getMessage(),
		Data:    data,
	})
}

// ResponseError 返回错误
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: code.getMessage(),
		Data:    nil,
	})
}

// ResponseErrorWithMessage 自定义响应错误
func ResponseErrorWithMessage(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

// GetPageSize 获取分页信息
func GetPageSize(c *gin.Context) (int64, int64) {

	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
