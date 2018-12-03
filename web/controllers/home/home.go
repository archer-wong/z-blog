package home

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"z-blog/web/model"
	"z-blog/web/service"
)

func Index(ctx *macaron.Context){
	articles, _ := service.ArticleIndex()
	ctx.Data["Articles"] = articles

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
	articles, _ := service.ArticleByCategory(cid)
	category, _ := service.CategoryById(cid)

	ctx.Data["Articles"] = articles
	ctx.Data["Category"] = category

	ctx.HTML(200, "home/index")
}

