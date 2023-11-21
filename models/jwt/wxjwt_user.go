package models

type WxJwtUserInfo struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	OpenId     string `json:"openId"`
	SessionKey string `json:"sessionKey"`
	Unionid    string `json:"unionid"`
}
