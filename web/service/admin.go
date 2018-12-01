
package service

import (
	"z-blog/helper"
	"z-blog/web/model"
)

func ValidateAccount(username, password string) (*model.Admin){
	user, _ := GetAdmin()
	if user.Username == username && user.Password == helper.Encrypt(password) {
		return user
	}
	return nil
}

func GetAdmin() (*model.Admin, error) {
	user := new(model.Admin)
	err := model.Db.First(user).Error
	return user, err
}

func AdminUpdate(admin, update model.Admin) error{
	err := model.Db.Model(&admin).Update(update).Error

	return err
}
