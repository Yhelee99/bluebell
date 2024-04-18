package mod

import (
	"fmt"
	"testing"
	"unsafe"
)

type s1 struct {
	a int8
	b string
	c int8
}

type s2 struct {
	a int8
	c int8
	b string
}

func Test_struct(t *testing.T) {
	a := s1{
		a: 1,
		b: "李银河",
		c: 2,
	}

	b := s2{
		a: 1,
		b: "李银河",
		c: 2,
	}

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}

/*
	在与要测试的代码相同的包中创建一个新的文件，文件名以_test.go结尾
	导入testing包
	编写测试函数，函数名以Test开头，接受一个*testing.T类型的参数
*/
