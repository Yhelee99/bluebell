package mod

//定义请求的结构体参数

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostVoted struct {
	//Userid从c中获取
	PostId    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction" binding:"oneof=1 -1 0"` // oneof 判断值在一个范围里	required如果传0，会判断为空
}
