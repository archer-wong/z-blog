package service

import (
	"z-blog/web/model"
)
func ArticleIndex() ([]model.Article, error){
	var articles []model.Article
	err := model.Db.Order("created_at desc").Preload("Category").Find(&articles).Error

	return articles, err
}

func ArticleStore(article model.Article) (model.Article ,error){
	err := model.Db.Create(&article).Error

	return article, err
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
	err := model.Db.Where("id = ?", id).Preload("Category").First(&article).Error

	return article, err
}

func ArticlePre(id int) (model.Article, error){
	var article model.Article
	err := model.Db.Where("id < ?", id).Order("id desc").First(&article).Error

	return article, err
}

func ArticleNext(id int) (model.Article, error){
	var article model.Article
	err := model.Db.Where("id > ?", id).Order("id asc").First(&article).Error

	return article, err
}


func TopArticles() ([]model.Article, error){
	var articles []model.Article
	err := model.Db.Order("view desc").Limit(5).Find(&articles).Error

	return articles, err
}

func ArticleByCategory(categoryId int) ([]model.Article, error){
	var articles []model.Article
	err := model.Db.Where("category_id = ?", categoryId).Find(&articles).Error

	return articles, err
}

