package model

type Admin struct {
	Id       uint   `gorm:"primary_key"`
	Username           string
	Password           string
}


func (Admin) TableName() string {
	return "admin"
}

