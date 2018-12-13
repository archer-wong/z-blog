package controllers

import (
	"z-blog/web/model"
	"z-blog/web/service"
)

func Categories() []model.Category {
	categories, err := service.CategoryIndex()
	if err != nil {
		return nil
	}
	return categories
}

func TopArticles() []model.Article {
	articles, err := service.TopArticles()
	if err != nil {
		return nil
	}
	return articles
}

func Links() []model.Link {
	links, err := service.LinkIndex()
	if err != nil {
		return nil
	}
	return links
}

func Admin() *model.Admin {
	admin, err := service.GetAdmin()
	if err != nil {
		return nil
	}
	return admin
}
