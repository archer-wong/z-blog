package home

import (
	"gopkg.in/macaron.v1"
	"z-blog/web/service"
)

func Index(ctx *macaron.Context){
	articles, _ := service.ArticleIndex()
	ctx.Data["Articles"] = articles

	ctx.HTML(200, "home/index")
}

func ArticleShow(ctx *macaron.Context){
	aid :=ctx.ParamsInt(":id")
	article, _ := service.ArticleById(aid)
	preArticle, _ := service.ArticlePre(aid)
	nextArticle, _ := service.ArticleNext(aid)

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

