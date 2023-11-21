package response

type WxloginRes struct {
	Token  string `json:"token"`
	OpenId string `json:"openId"`
}
