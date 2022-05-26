package main

import (
	"github.com/skyareas/skyjet"
)

func main() {
	app := skyjet.App()
	app.Use("/assets", skyjet.Static())
	app.Get("/", func(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
		return res.SendFile("views/index.html")
	})
	app.Run()
}
