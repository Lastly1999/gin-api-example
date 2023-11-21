package models

type UserInfo struct {
	Base
	Unionid   string `gorm:"comment:用户unionid;not null;unique" json:"unionid"`
	AvatarUrl string `gorm:"comment:用户头像" json:"avatarUrl"`
	NickName  string `gorm:"comment:用户昵称" json:"nickName"`
	Phone     string `gorm:"comment:用户手机号" json:"phone"`
}

func (m *UserInfo) TableName() string {
	return "wx_user_info"
}
