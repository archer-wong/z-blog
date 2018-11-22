package main

import "gopkg.in/macaron.v1"

func main() {
    m := macaron.Classic()
    m.Get("/", func() string {
        return "z-blog init"
    })
    m.Run()
}
