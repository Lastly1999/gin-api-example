package models

type UserWx struct {
	Base
	OpenId    string `gorm:"type:varchar(255);not null;unique;comment:微信openid" json:"openId"`
	Unionid   string `gorm:"type:varchar(255);not null;unique;comment:微信unionid" json:"unionid"`
	NickName  string `gorm:"type:varchar(255);not null;unique;comment:用户昵称" json:"nickName"`
	AvatarUrl string `gorm:"type:varchar(255);not null;comment:用户头像" json:"avatarUrl"`
	City      string `gorm:"type:varchar(255);null;comment:城市" json:"city"`
	Province  string `gorm:"type:varchar(255);null;comment:省份" json:"province"`
	Country   string `gorm:"type:varchar(255);null;comment:国家" json:"country"`
	Type      int    `gorm:"type:int(1);not null;comment:类型 0-小程序 1-公众号 2-H5 3-APP" json:"type"`
}

func (m *UserWx) TableName() string {
	return "wx_user"
}
