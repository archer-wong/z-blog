package service

import (
	"z-blog/web/model"
)

func MarkdownIndex() ([]model.Markdown, error) {
	var markdowns []model.Markdown
	err := model.Db.Order("created_at desc").Find(&markdowns).Error

	return markdowns, err
}

func MarkdownStore(markdown model.Markdown) error {
	err := model.Db.Create(&markdown).Error

	return err
}

func MarkdownUpdate(markdown, update model.Markdown) error {
	err := model.Db.Model(&markdown).Update(update).Error

	return err
}

func MarkdownDeleteById(id int) error {
	var markdown model.Markdown
	err := model.Db.Where("id = ?", id).Delete(&markdown).Error

	return err
}

func MarkdownById(id int) (model.Markdown, error) {
	var markdown model.Markdown
	err := model.Db.Where("id = ?", id).First(&markdown).Error

	return markdown, err
}
