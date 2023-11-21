package models

type User struct {
	Base
	Username string `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"-"`
	Email    string `gorm:"type:varchar(20);null;unique" json:"email"`
}

func (m *User) TableName() string {
	return "sys_user"
}
