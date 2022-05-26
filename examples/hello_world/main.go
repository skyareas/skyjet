package main

import "github.com/skyareas/skyjet"

func main() {
	app := skyjet.App()
	app.Get("/", func(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
		return res.Send([]byte("Hello, Skyjet!"))
	})
	app.Run()
}
