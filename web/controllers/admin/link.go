package admin

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"z-blog/web/model"
	"z-blog/web/service"
)

func LinkIndex(ctx *macaron.Context, f *session.Flash){
	links, err := service.LinkIndex()
	if err != nil {
		f.Error("获取友情链接列表失败")
	}
	ctx.Data["Links"] = links
	ctx.HTML(200, "admin/link/index")
}
func LinkStore(ctx *macaron.Context, f *session.Flash){
	title := ctx.QueryTrim("title")
	url := ctx.QueryTrim("url")

	link := model.Link{Title:title, Url:url}
	err := service.LinkStore(link)

	if err != nil {
		f.Error("创建友情链接[ " + title + " ]失败")
		ctx.Redirect("/admin/link/create")
		return
	}

	f.Success("创建友情链接[ " + title + " ]成功")
	ctx.Redirect("/admin/link")
}

func LinkUpdate(ctx *macaron.Context, f *session.Flash){
	cid :=ctx.ParamsInt(":id")
	title := ctx.QueryTrim("title")

	link := model.Link{Id: cid}
	update := model.Link{Title:title}
	err := service.LinkUpdate(link, update)

	if err != nil {
		idString :=ctx.Params(":id")
		f.Error("修改友情链接[ " + title + " ]失败")
		ctx.Redirect("/admin/link/"+ idString +"/edit")
		return
	}

	f.Success("修改友情链接[ " + title + " ]成功")
	ctx.Redirect("/admin/link")
}

func LinkDestroy(ctx *macaron.Context){
	result := false
	cid :=ctx.ParamsInt(":id")
	err := service.LinkDeleteById(cid)
	if err == nil {
		result = true
	}

	ctx.JSON(200, map[string]interface{}{
		"success": result,
	})
}


