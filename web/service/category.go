
package service

import (
	"z-blog/web/model"
)
func CategoryIndex() ([]model.Category, error){
	var categories []model.Category
	err := model.Db.Order("created_at desc").Find(&categories).Error

	return categories, err
}

func CategoryStore(category model.Category) error{
	err := model.Db.Create(&category).Error

	return err
}

func CategoryUpdate(category, update model.Category) error{
	err := model.Db.Model(&category).Update(update).Error

	return err
}

func CategoryDeleteById(id int) (error){
	var category model.Category
	err := model.Db.Where("id = ?", id).Delete(&category).Error

	return err
}

func CategoryById(id int) (model.Category, error){
	var category model.Category
	err := model.Db.Where("id = ?", id).First(&category).Error

	return category, err
}


