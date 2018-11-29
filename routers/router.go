package routers

import (
	"z-blog/web/controllers"
	"z-blog/web/controllers/admin"
	"z-blog/web/controllers/home"
	"z-blog/web/middlewares"
)

func initRouter() {
	m.Get("/", home.Index).Name("index")
	m.Get("/admin", admin.Index).Name("admin.index")
	m.Get("/admin/logout", admin.Logout).Name("admin.logout")

	//分类路由
	m.Get("/admin/category", admin.CategoryIndex).Name("admin.category_index")
	m.Get("/admin/category/create", admin.CategoryCreate).Name("admin.category_create")
	m.Post("/admin/category", admin.CategoryStore).Name("admin.category_store")
	m.Get("/admin/category/:id:int/edit", admin.CategoryEdit).Name("admin.category_edit")
	m.Post("/admin/category/:id:int", admin.CategoryUpdate).Name("admin.category_update")
	m.Delete("/admin/category/:id:int", admin.CategoryDestroy).Name("admin.category_destroy")

	//文章路由
	m.Get("/admin/article", admin.ArticleIndex).Name("admin.article_index")
	m.Get("/admin/article/create", admin.ArticleCreate).Name("admin.article_create")
	m.Post("/admin/article", admin.ArticleStore).Name("admin.article_store")
	m.Get("/admin/article/:id:int/edit", admin.ArticleEdit).Name("admin.article_edit")
	m.Post("/admin/article/:id:int", admin.ArticleUpdate).Name("admin.article_update")
	m.Delete("/admin/article/:id:int", admin.ArticleDestroy).Name("admin.article_destroy")

	//友情链接
	m.Get("/admin/link", admin.LinkIndex).Name("admin.link_index")
	m.Post("/admin/link", admin.LinkStore).Name("admin.link_store")
	m.Post("/admin/link/:id:int", admin.LinkUpdate).Name("admin.link_update")
	m.Delete("/admin/link/:id:int", admin.LinkDestroy).Name("admin.link_destroy")


	//前台路由
	m.Get("/article/:id:int", home.ArticleShow).Name("home.article_show")
	m.Get("/category/:id:int", home.ArticlesByCategory).Name("home.articles_by_category")

	m.Group("admin", func() {
		m.Get("/login", admin.Login).Name("admin.login")
		m.Post("/do-login", admin.DoLogin).Name("admin.do-login")

	}, middlewares.AuthLogin)

	m.Group("admin", func() {
		//m.Get("/", admin.Index).Name("admin.index")
	}, middlewares.AuthPermission)


	m.NotFound(controllers.NotFound)
}
