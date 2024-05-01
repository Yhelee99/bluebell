package mod

//定义请求的结构体参数

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`                     //用户名
	Password   string `json:"password" binding:"required"`                     //密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //确认密码
	//eqfield必须相同
}

type ParamSignIn struct {
	Username string `json:"username" binding:"required"` //用户名
	Password string `json:"password" binding:"required"` //密码
}

type ParamsGetPostList struct {
	Page        int64  `json:"page" form:"page"`                 // 页码
	Size        int64  `json:"size" form:"size"`                 // 每页数据量
	Type        string `json:"type" form:"type" example:"score"` // 排序依据
	CommunityId int64  `json:"community_id" form:"community_id"` // 可以为空
	//form 处理ShouldBindQuery参数的tag
}
