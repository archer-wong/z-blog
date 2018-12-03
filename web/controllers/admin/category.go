package admin

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"z-blog/web/model"
	"z-blog/web/service"
)

func CategoryIndex(ctx *macaron.Context, f *session.Flash){
	categories, err := service.CategoryIndex()
	if err != nil {
		f.Error("获取分类列表失败")
	}
	ctx.Data["Categories"] = categories
	ctx.HTML(200, "admin/category/index")
}

func CategoryCreate(ctx *macaron.Context){
	ctx.HTML(200, "admin/category/create")
}

func CategoryStore(ctx *macaron.Context, f *session.Flash){
	title := ctx.QueryTrim("title")

	category := model.Category{Title:title}
	err := service.CategoryStore(category)

	if err != nil {
		f.Error("创建新分类[ " + title + " ]失败")
		ctx.Redirect("/admin/category/create")
		return
	}

	f.Success("创建新分类[ " + title + " ]成功")
	ctx.Redirect("/admin/category")
}

func CategoryEdit(ctx *macaron.Context, f *session.Flash){
	cid :=ctx.ParamsInt(":id")
	category, err := service.CategoryById(cid)
	ctx.Data["Category"] = category
	if err != nil {
		f.Error("获取分类失败")
	}

	ctx.HTML(200, "admin/category/edit")
}

func CategoryUpdate(ctx *macaron.Context, f *session.Flash){
	cid :=ctx.ParamsInt(":id")
	title := ctx.QueryTrim("title")

	category := model.Category{Id: cid}
	update := model.Category{Title:title}
	err := service.CategoryUpdate(category, update)

	if err != nil {
		idString :=ctx.Params(":id")
		f.Error("修改分类[ " + title + " ]失败")
		ctx.Redirect("/admin/category/"+ idString +"/edit")
		return
	}

	f.Success("修改分类[ " + title + " ]成功")
	ctx.Redirect("/admin/category")
}

func CategoryDestroy(ctx *macaron.Context){
	cid :=ctx.ParamsInt(":id")
	articles, err := service.ArticleByCategory(cid)

	if len(articles) > 0 {
		ctx.JSON(200, map[string]interface{}{
			"success": false,
			"message": "该分类下仍有文章，不能被删除",
		})
		return
	}

	err = service.CategoryDeleteById(cid)
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"success": false,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"success": true,
	})
	return

}


