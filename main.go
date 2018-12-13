package main

import (
	"z-blog/routers"
)

func main() {
	app := routers.New()
	app.ListenAndServe()
}
