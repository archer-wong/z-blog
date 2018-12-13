package middlewares

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"net/http"
)

func AuthLogin(ctx *macaron.Context, sess session.Store) {
	if userId := sess.Get("userId"); userId == nil {
		ctx.Next()
	} else {
		ctx.Redirect(ctx.URLFor("admin.index"), http.StatusSeeOther)
	}
}

func AuthPermission(ctx *macaron.Context, sess session.Store) {
	if userId := sess.Get("userId"); userId == nil {
		ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
	} else {
		ctx.Next()
	}
}
