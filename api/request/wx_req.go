package request

type WxSignReq struct {
	Code string `json:"code" binding:"required" msg:"code不能为空"`
}
