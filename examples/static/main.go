package main

import (
	"github.com/skyareas/skyjet"
)

func main() {
	app := skyjet.SharedApp()
	app.Get("/", func(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
		return res.SendFile("index.html")
	})
	app.Use("/assets", skyjet.Static())
	app.Run()
}
