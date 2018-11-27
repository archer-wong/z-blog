package admin

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"z-blog/web/model"
	"z-blog/web/service"
)

func ArticleIndex(ctx *macaron.Context, f *session.Flash){
	articles, err := service.ArticleIndex()
	if err != nil {
		f.Error("获取文章列表失败")
	}
	ctx.Data["Articles"] = articles
	ctx.HTML(200, "admin/article/index")
}

func ArticleCreate(ctx *macaron.Context, f *session.Flash){
	categories, err := service.CategoryIndex()
	if err != nil {
		f.Error("获取分类列表失败")
	}
	ctx.Data["Categories"] = categories
	ctx.HTML(200, "admin/article/create")
}

func ArticleStore(ctx *macaron.Context, f *session.Flash){
	title := ctx.QueryTrim("title")
	content := ctx.QueryTrim("content")
	tag := ctx.QueryTrim("tag")
	categoryId := ctx.QueryInt("category_id")

	article := model.Article{CategoryId:categoryId, Title:title, Content:content, Tag:tag}
	err := service.ArticleStore(article)

	if err != nil {
		f.Error("创建新文章[ " + title + " ]失败")
		ctx.Redirect("/admin/article/create")
		return
	}

	f.Success("创建新文章[ " + title + " ]成功")
	ctx.Redirect("/admin/article")
}

func ArticleEdit(ctx *macaron.Context, f *session.Flash){
	categories, err := service.CategoryIndex()
	if err != nil {
		f.Error("获取分类列表失败")
	}
	ctx.Data["Categories"] = categories

	aid :=ctx.ParamsInt(":id")
	article, err := service.ArticleById(aid)
	ctx.Data["Article"] = article
	if err != nil {
		f.Error("获取文章内容失败")
	}

	ctx.HTML(200, "admin/article/edit")
}

func ArticleUpdate(ctx *macaron.Context, f *session.Flash){
	aid :=ctx.ParamsInt(":id")
	title := ctx.QueryTrim("title")
	content := ctx.QueryTrim("content")
	tag := ctx.QueryTrim("tag")
	categoryId := ctx.QueryInt("category_id")

	article := model.Article{Id: aid}
	update := model.Article{CategoryId:categoryId, Title:title, Content:content, Tag:tag}

	err := service.ArticleUpdate(article, update)
	if err != nil {
		idString :=ctx.Params(":id")
		f.Error("修改文章[ " + title + " ]失败")
		ctx.Redirect("/admin/article/"+ idString +"/edit")
		return
	}

	f.Success("修改文章[ " + title + " ]成功")
	ctx.Redirect("/admin/article")
}

func ArticleDestroy(ctx *macaron.Context){
	result := false
	cid :=ctx.ParamsInt(":id")
	err := service.ArticleDeleteById(cid)
	if err == nil {
		result = true
	}

	ctx.JSON(200, map[string]interface{}{
		"success": result,
	})
}


