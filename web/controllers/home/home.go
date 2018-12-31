package home

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"html/template"
	"strconv"
	"z-blog/helper"
	"z-blog/web/controllers"
	"z-blog/web/model"
	"z-blog/web/service"
)

type IndexData struct {
	Articles []model.Article
	Total    int
	Page     int
	Category model.Category
}

func Index(ctx *macaron.Context) {
	page := ctx.QueryTrim("page")
	pageInt, _ := strconv.Atoi(page)
	if pageInt == 0 {
		pageInt = 1
	}

	articles, count, _ := service.ArticlePage(pageInt, 10)
	indexData := IndexData{Articles: articles, Total: count, Page: pageInt}

	funcMap := FunctionMap()
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("views/layout/home/home.html", "views/home/index.html", "views/layout/home/pagination.html"))

	t.ExecuteTemplate(ctx.Resp, "home.home", indexData)
}

func ArticlesByCategory(ctx *macaron.Context) {
	cid := ctx.ParamsInt(":id")
	page := ctx.QueryTrim("page")
	pageInt, _ := strconv.Atoi(page)
	if pageInt == 0 {
		pageInt = 1
	}

	articles, count, _ := service.ArticleByCategoryPage(cid, pageInt, 10)
	category, _ := service.CategoryById(cid)

	indexData := IndexData{Articles: articles, Total: count, Page: pageInt, Category: category}

	funcMap := FunctionMap()
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("views/layout/home/home.html", "views/home/index.html", "views/layout/home/pagination.html"))

	t.ExecuteTemplate(ctx.Resp, "home.home", indexData)
}

type ArticleShowData struct {
	Article     model.Article
	PreArticle  model.Article
	NextArticle model.Article
}

func ArticleShow(ctx *macaron.Context, sess session.Store) {
	aid := ctx.ParamsInt(":id")
	article, _ := service.ArticleById(aid)
	preArticle, _ := service.ArticlePre(aid)
	nextArticle, _ := service.ArticleNext(aid)
	articleShowData := ArticleShowData{Article: article, PreArticle: preArticle, NextArticle: nextArticle}

	//阅读统计
	if userId := sess.Get("userId"); userId == nil {
		view := article.View + 1
		update := model.Article{View: view}
		service.ArticleUpdate(article, update)
	}

	funcMap := FunctionMap()
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("views/layout/home/home.html", "views/home/article.html"))

	t.ExecuteTemplate(ctx.Resp, "home.home", articleShowData)
}

func FunctionMap() template.FuncMap {
	funcMap := template.FuncMap{"Categories": controllers.Categories, "TopArticles": controllers.TopArticles, "Links": controllers.Links, "Blog": controllers.Blog, "DateTime": helper.DateTime}

	return funcMap
}
