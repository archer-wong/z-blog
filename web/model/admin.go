package model

type Admin struct {
	Id       int    `gorm:"primary_key"`
	Username string `gorm:not null;unique;size:15"`
	Password string `gorm:not null;size:32"`
}

func (Admin) TableName() string {
	return "admin"
}
