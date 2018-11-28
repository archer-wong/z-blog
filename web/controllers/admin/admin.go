package admin

import (
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"net/http"
	"z-blog/web/service"
)

func Index(ctx *macaron.Context){
	ctx.HTML(200, "admin/index")
}


func Login(ctx *macaron.Context) {
	ctx.HTML(200, "admin/login")
}

func DoLogin(ctx *macaron.Context, cpt *captcha.Captcha, sess session.Store, flash *session.Flash) {
	username := ctx.QueryTrim("username")
	password := ctx.QueryTrim("password")
	captchaId := ctx.QueryTrim("captcha_id")

	if len(username) == 0 || len(password) == 0 {
		flash.Error("请填写用户名和密码")
		ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
		return
	}

	if len(captchaId) != 0 && !cpt.VerifyReq(ctx.Req){
		flash.Error("验证码错误")
		ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
		return
	}

	user := service.ValidateAccount(username, password)
	if user != nil {
		sess.Set("userId", user.Id)
		ctx.Redirect(ctx.URLFor("admin.index"), http.StatusSeeOther)
	} else {
		flash.Error("用户名或密码错误")
		ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
	}
}

func Logout(ctx *macaron.Context, sess session.Store) {
	sess.Flush()
	ctx.Redirect("/admin/login")
}

