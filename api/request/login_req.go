package request

type LoginReq struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空!"`
	Password string `json:"password" binding:"required" msg:"密码不能为空!"`
}
