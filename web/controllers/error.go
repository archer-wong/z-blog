package controllers

import (
	"gopkg.in/macaron.v1"
	"net/http"
)

func NotFound(ctx *macaron.Context) {
	ctx.HTML(http.StatusNotFound, "errors/404")
}
