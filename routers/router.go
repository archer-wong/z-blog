package routers

import (
	"z-blog/web/controllers"
	"z-blog/web/controllers/admin"
	"z-blog/web/controllers/home"
	"z-blog/web/middlewares"
)

func initRouter() {
	m.Get("/", home.Index).Name("index")

	m.Group("admin", func() {
		m.Get("/login", admin.Login).Name("admin.login")
		m.Post("/do-login", admin.DoLogin).Name("admin.do-login")
	}, middlewares.AuthLogin)

	m.Group("admin", func() {
		m.Get("/", admin.Index).Name("admin.index")
	}, middlewares.AuthPermission)


	m.NotFound(controllers.NotFound)
}
