package admin

import (
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"net/http"
	"z-blog/helper"
	"z-blog/web/model"
	"z-blog/web/service"
)

func Index(ctx *macaron.Context) {
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

	if len(captchaId) != 0 && !cpt.VerifyReq(ctx.Req) {
		flash.Error("验证码错误")
		ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
		return
	}

	user := service.ValidateAccount(username, password)
	if user != nil {
		sess.Set("userId", user.Id)
		ctx.Redirect(ctx.URLFor("admin.index"), http.StatusSeeOther)
		return
	}

	flash.Error("用户名或密码错误")
	ctx.Redirect(ctx.URLFor("admin.login"), http.StatusFound)
}

func Logout(ctx *macaron.Context, sess session.Store) {
	sess.Flush()
	ctx.Redirect("/admin/login")
}

func PasswordUpdate(ctx *macaron.Context, sess session.Store, f *session.Flash) {
	username := ctx.QueryTrim("username")
	password := ctx.QueryTrim("password")
	newPassword := ctx.QueryTrim("new_password")
	confirmPassword := ctx.QueryTrim("confirm_password")

	if len(newPassword) < 6 {
		f.Error("新密码不能小于6位")
		ctx.Redirect("/admin")
		return
	}

	if newPassword != confirmPassword {
		f.Error("两次输入的密码不一致")
		ctx.Redirect("/admin")
		return
	}

	user := service.ValidateAccount(username, password)
	if user == nil {
		f.Error("原密码错误")
		ctx.Redirect("/admin")
		return
	}

	encryptPassword := helper.Encrypt(newPassword)
	admin := model.Admin{Username: username}
	update := model.Admin{Password: encryptPassword}
	err := service.AdminUpdate(admin, update)

	if err != nil {
		f.Error("修改密码失败，请稍后重试")
		ctx.Redirect("/admin")
		return
	}

	f.Success("修改密码成功，请退出后重新登录")
	ctx.Redirect("/admin")
}
