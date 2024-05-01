package mysql

import (
	"bluebell/mod"
	"testing"
)

// 单元测试需要初始化db,避免使用db时出现空指针异常
func init() {
	dbCfg := &mod.Mysql{
		Port:     668,
		Password: "123456",
		Host:     "localhost",
		Dbname:   "bluebell",
		Username: "root",
	}

	Init(dbCfg)
}

func TestCreatPost(t *testing.T) {
	post := &mod.Post{
		Community_id: 2,
		Title:        "test",
		Content:      "testing",
	}
	err := CreatPost(post)
	if err != nil {
		t.Fatalf("入库失败！错误：%v\n", err)
	}
	//打印成功的日志
	t.Logf("测试成功！")
}
