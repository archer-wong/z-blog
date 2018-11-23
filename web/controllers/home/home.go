package home


import "gopkg.in/macaron.v1"

func Index(ctx *macaron.Context) string{
	ctx.Data["msg"] = "welcome to z-blog home page"
	return  "welcome to z-blog home page"
}
