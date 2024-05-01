package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostVoted(t *testing.T) {

	url := "/api/v1/createpost"

	//设置运行模式
	gin.SetMode(gin.TestMode)

	//创建路由,不直接调用是为了不造成互相引用问题
	r := gin.Default()
	r.POST(url, CreatPostHandler)

	body := `{
    "title":"你好",
    "content":"用户李卓的帖子测试完成了的1111"
}`

	//创建请求
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	//创建响应对象
	w := httptest.NewRecorder()
	//发送请求
	r.ServeHTTP(w, req)

	//校验响应
	//方法1：判断返回的内容中是否包含指定的字符串
	//assert.Contains(t, w.Body.String(), "请先登录！")

	//方法2：反序列化返回再做比较
	p := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), p); err != nil {
		t.Fatalf("反序列化失败!错误：%v\n", err)
	}
	assert.Equal(t, ErrorCodeInvalidParams, p.Code)
}
