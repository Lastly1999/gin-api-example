package models

type Base struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	CreatedAt *LocalTime `gorm:"type:datetime(0)" json:"created_at"`
	UpdatedAt *LocalTime `gorm:"type:datetime(0)" json:"updated_at"`
	DeletedAt *LocalTime `gorm:"type:datetime(0)" json:"deleted_at" sql:"index"`
}
