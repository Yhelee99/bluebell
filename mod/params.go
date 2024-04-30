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
	//form 处理ShouldBindQuery参数的tag
	Page        int64  `json:"page" form:"page"`                 // 页码
	Size        int64  `json:"size" form:"size"`                 // 每页数据量
	Type        string `json:"type" form:"type" example:"score"` // 排序依据
	CommunityId int64  `json:"community_id" form:"community_id"` // 可以为空
}
