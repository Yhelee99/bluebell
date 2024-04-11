package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"time"
)

var node *sf.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	//time.Pares 字符串解析为时间 layout 参数指定了输入字符串的时间格式，value 参数是待解析的时间字符串
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		zap.L().Fatal("时间解析失败！")
		return
	}

	//snowflake.Epoch表示的是算法中的起始时间
	//time.UnixNano() 返回的是当前时间的纳秒级时间戳
	//1毫秒=1000000纳秒，st.UnixNano()/1000000表示把算法起始时间转化为毫秒级
	sf.Epoch = st.UnixNano() / 1000000 //指定时间戳级别

	node, err = sf.NewNode(machineId)
	return
}

// 生成一个int64的GenId（可指定其他数据类型）
func GetSnowflakeId() int64 {
	return node.Generate().Int64()
}
