package service

import (
	"z-blog/web/model"
)
func ArticleIndex() ([]model.ArticleWithCategory, error){
	var articles []model.Article
	err := model.Db.Order("created_at desc").Find(&articles).Error

	var articlesWithCategory []model.ArticleWithCategory
	for _, article := range articles{
		var category model.Category
		model.Db.Where("id = ?", article.CategoryId).First(&category)
		articlesWithCategory = append(articlesWithCategory, model.ArticleWithCategory{article, category.Title})
	}

	return articlesWithCategory, err
}

func ArticleStore(article model.Article) error{
	err := model.Db.Create(&article).Error

	return err
}

func ArticleUpdate(article, update model.Article) error{
	err := model.Db.Model(&article).Update(update).Error

	return err
}

func ArticleDeleteById(id int) (error){
	var article model.Article
	err := model.Db.Where("id = ?", id).Delete(&article).Error

	return err
}

func ArticleById(id int) (model.Article, error){
	var article model.Article
	err := model.Db.Where("id = ?", id).First(&article).Error

	return article, err
}

