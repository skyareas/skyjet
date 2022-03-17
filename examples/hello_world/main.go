package main

import (
	"github.com/skyareas/skyjet"
)

type HelloController struct{}

func (c *HelloController) HandleRequest(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	name, nameParamExists := req.LookupParam("name")
	if !nameParamExists {
		name = req.Query().Get("name")
	}
	if name != "" {
		return res.Send([]byte("Hello, " + name + "!"))
	}
	return res.Send([]byte("Please tell me your name!"))
}

func main() {
	app := skyjet.SharedApp()
	controller := &HelloController{}
	app.Get("/", controller)
	app.Get("/:name", controller)
	app.Run()
}
