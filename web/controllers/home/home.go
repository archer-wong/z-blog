package home

import (
	"gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context){
	ctx.Data["msg"] = "welcome to z-blog home page"
	ctx.HTML(200, "home/index")
}

