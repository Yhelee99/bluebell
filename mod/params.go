package mod

//定义请求的结构体参数

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //eqfield必须相同
}

type ParamSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamsGetPostList struct {
	Page        int64  `json:"page" form:"page"` //form 处理ShouldBindQuery参数的tag
	Size        int64  `json:"size" form:"size"`
	Type        string `json:"type" form:"type"`
	CommunityId int64  `json:"community_id" form:"community_id"`
}
