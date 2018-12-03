package routers

import (
	"mime/multipart"
	"z-blog/web/controllers"
	"z-blog/web/controllers/admin"
	"z-blog/web/controllers/home"
	"z-blog/web/middlewares"
)

type UploadForm struct {
	Title      string                `form:"title"`
	TextUpload *multipart.FileHeader `form:"txtUpload"`
}
func initRouter() {
	//前台路由
	m.Get("/", home.Index).Name("index")
	m.Get("/article/:id:int", home.ArticleShow).Name("home.article_show")
	m.Get("/category/:id:int", home.ArticlesByCategory).Name("home.articles_by_category")

	m.Group("admin", func() {
		m.Get("/login", admin.Login).Name("admin.login")
		m.Post("/do-login", admin.DoLogin).Name("admin.do_login")

	}, middlewares.AuthLogin)

	m.Group("admin", func() {
		m.Get("/", admin.Index).Name("admin.index")
		m.Post("/password", admin.PasswordUpdate).Name("admin.password_update")
		m.Get("/logout", admin.Logout).Name("admin.logout")

		//分类路由
		m.Get("/category", admin.CategoryIndex).Name("admin.category_index")
		m.Get("/category/create", admin.CategoryCreate).Name("admin.category_create")
		m.Post("/category", admin.CategoryStore).Name("admin.category_store")
		m.Get("/category/:id:int/edit", admin.CategoryEdit).Name("admin.category_edit")
		m.Post("/category/:id:int", admin.CategoryUpdate).Name("admin.category_update")
		m.Delete("/category/:id:int", admin.CategoryDestroy).Name("admin.category_destroy")

		//文章路由
		m.Get("/article", admin.ArticleIndex).Name("admin.article_index")
		m.Get("/article/create", admin.ArticleCreate).Name("admin.article_create")
		m.Post("/article", admin.ArticleStore).Name("admin.article_store")
		m.Get("/article/:id:int/edit", admin.ArticleEdit).Name("admin.article_edit")
		m.Post("/article/:id:int", admin.ArticleUpdate).Name("admin.article_update")
		m.Delete("/article/:id:int", admin.ArticleDestroy).Name("admin.article_destroy")

		//友情链接
		m.Get("/link", admin.LinkIndex).Name("admin.link_index")
		m.Post("/link", admin.LinkStore).Name("admin.link_store")
		m.Post("/link/:id:int", admin.LinkUpdate).Name("admin.link_update")
		m.Delete("/link/:id:int", admin.LinkDestroy).Name("admin.link_destroy")

		//批量上传
		m.Get("/markdown", admin.MarkdownIndex).Name("admin.markdown_index")
		m.Get("/markdown/create", admin.MarkdownCreate).Name("admin.markdown_create")
		m.Post("/markdown-upload", admin.MarkdownUpload).Name("admin.markdown_upload")
		m.Delete("/markdown/:id:int", admin.MarkdownDestroy).Name("admin.markdown_destroy")
		m.Post("/markdown/publish", admin.MarkdownPublish).Name("admin.markdown_publish")

	}, middlewares.AuthPermission)


	m.NotFound(controllers.NotFound)
}
