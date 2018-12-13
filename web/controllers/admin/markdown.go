package admin

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"z-blog/web/model"
	"z-blog/web/service"
)

func MarkdownIndex(ctx *macaron.Context, f *session.Flash) {
	markdowns, err := service.MarkdownIndex()
	if err != nil {
		f.Error("获取Markdown列表失败")
	}
	ctx.Data["Markdowns"] = markdowns

	categories, err := service.CategoryIndex()
	if err != nil {
		f.Error("获取分类列表失败")
	}
	ctx.Data["Categories"] = categories

	ctx.HTML(200, "admin/markdown/index")
}

func MarkdownCreate(ctx *macaron.Context) {
	ctx.HTML(200, "admin/markdown/create")
}

func MarkdownUpload(ctx *macaron.Context) {
	ctx.Req.ParseMultipartForm(32 << 20)
	file, handler, err := ctx.Req.FormFile("files[]")
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"result": false,
			"msg":    "上传文件[" + handler.Filename + "]失败",
		})
		return
	}
	defer file.Close()

	originName := handler.Filename
	fileName := time.Now().Format("2006-01-02-15-04-05") + "_" + originName

	f, err := os.OpenFile("storage/markdown/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"result": false,
			"msg":    "上传文件[" + originName + "]失败",
		})
		return
	}
	defer f.Close()
	io.Copy(f, file)

	title := strings.Trim(originName, ".md")

	markdown := model.Markdown{Title: title, Path: fileName}
	error := service.MarkdownStore(markdown)

	if error != nil {
		ctx.JSON(200, map[string]interface{}{
			"result": false,
			"msg":    "文件[" + originName + "]导入数据库失败",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"result": true,
		"msg":    "上传文件[" + originName + "]成功",
	})
}

func MarkdownDestroy(ctx *macaron.Context) {
	result := false
	cid := ctx.ParamsInt(":id")
	err := service.MarkdownDeleteById(cid)
	if err == nil {
		result = true
	}

	ctx.JSON(200, map[string]interface{}{
		"success": result,
	})
}

func MarkdownPublish(ctx *macaron.Context, f *session.Flash) {
	categoryId := ctx.QueryInt("category_id")
	markdownIds := ctx.QueryTrim("ids")

	errMsg := ""
	idSlices := strings.Split(markdownIds, ",")
	for _, id := range idSlices {
		int, err := strconv.Atoi(id)
		if err != nil {
			errMsg = errMsg + "文章[" + id + "]发布失败；"
			continue
		}

		markdown, _ := service.MarkdownById(int)
		bytes, err := ioutil.ReadFile("storage/markdown/" + markdown.Path)
		if err != nil {
			errMsg = errMsg + "文章[" + markdown.Title + "]发布失败；"
			continue
		}

		article := model.Article{CategoryId: categoryId, Title: markdown.Title, Content: string(bytes), Tag: ""}
		newArticle, err := service.ArticleStore(article)
		if err != nil {
			errMsg = errMsg + "文章[" + markdown.Title + "]发布失败；"
			continue
		}

		update := model.Markdown{ArticleId: newArticle.Id}
		err = service.MarkdownUpdate(markdown, update)
		if err != nil {
			errMsg = errMsg + "文章[" + markdown.Title + "]发布失败；"
			continue
		}
	}

	if errMsg != "" {
		f.Error(errMsg)
	} else {
		f.Success("批量发布文章成功")
	}

	ctx.Redirect("/admin/article")
}
