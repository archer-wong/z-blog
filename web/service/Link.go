
package service

import (
	"z-blog/web/model"
)
func LinkIndex() ([]model.Link, error){
	var links []model.Link
	err := model.Db.Order("created_at desc").Find(&links).Error

	return links, err
}

func LinkStore(link model.Link) error{
	err := model.Db.Create(&link).Error

	return err
}

func LinkUpdate(link, update model.Link) error{
	err := model.Db.Model(&link).Update(update).Error

	return err
}

func LinkDeleteById(id int) (error){
	var link model.Link
	err := model.Db.Where("id = ?", id).Delete(&link).Error

	return err
}


