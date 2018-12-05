package home

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"strconv"
	"z-blog/web/model"
	"z-blog/web/service"
)

func Index(ctx *macaron.Context){
	page := ctx.QueryTrim("page")
	pageInt, _ :=strconv.Atoi(page)
	if pageInt == 0 {
		pageInt =1
	}

	articles, count, _ := service.ArticlePage(pageInt, 10)

	ctx.Data["Articles"] = articles
	ctx.Data["Total"] = count
	ctx.Data["Page"] = pageInt

	ctx.HTML(200, "home/index")
}

func ArticleShow(ctx *macaron.Context, sess session.Store){
	aid :=ctx.ParamsInt(":id")
	article, _ := service.ArticleById(aid)
	preArticle, _ := service.ArticlePre(aid)
	nextArticle, _ := service.ArticleNext(aid)

	//阅读统计
	if userId := sess.Get("userId"); userId == nil {
		view := article.View + 1
		update := model.Article{View:view}
		service.ArticleUpdate(article, update)
	}

	ctx.Data["Article"] = article
	ctx.Data["PreArticle"] = preArticle
	ctx.Data["NextArticle"] = nextArticle

	ctx.HTML(200, "home/article")
}

func ArticlesByCategory(ctx *macaron.Context){
	cid :=ctx.ParamsInt(":id")
	page := ctx.QueryTrim("page")
	pageInt, _ :=strconv.Atoi(page)
	if pageInt == 0 {
		pageInt =1
	}

	articles, count, _ := service.ArticleByCategoryPage(cid, pageInt,10)
	category, _ := service.CategoryById(cid)

	ctx.Data["Articles"] = articles
	ctx.Data["Category"] = category
	ctx.Data["Total"] = count
	ctx.Data["Page"] = pageInt

	ctx.HTML(200, "home/index")
}

