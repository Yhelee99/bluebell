package logic

import (
	"bluebell/dao/mysql"
	"bluebell/mod"
	snowflake "bluebell/pkg"
)

func CreatPost(p *mod.Post) error {
	//1:生成一个Post_id
	p.Post_id = snowflake.GetSnowflakeId()
	//2：入库
	//return CreatPost(p)//有点无语，自己返回自己造成栈溢出....
	return mysql.CreatPost(p)
}
