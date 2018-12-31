package controllers

import (
	"z-blog/modules/setting"
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

type BlogData struct {
	Name        string
	Description string
	Keywords    string
	Github      string
	Username    string
}

func Blog() BlogData {
	name := setting.Cfg.Section("blog").Key("BLOG_NAME").String()
	description := setting.Cfg.Section("blog").Key("BLOG_DESCRIPTION").String()
	keywords := setting.Cfg.Section("blog").Key("BLOG_KEYWORDS").String()
	github := setting.Cfg.Section("blog").Key("GITHUB").String()
	username := setting.Cfg.Section("blog").Key("USER_NAME").String()

	blog := BlogData{Name: name, Description: description, Keywords: keywords, Github: github, Username: username}

	return blog
}

func Admin() *model.Admin {
	admin, err := service.GetAdmin()
	if err != nil {
		return nil
	}
	return admin
}
