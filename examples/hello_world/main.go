package main

import (
	"github.com/skyareas/skyjet"
)

func main() {
	app := skyjet.SharedApp()
	app.Get("/", requestHandler)
	app.Get("/:name", requestHandler)
	app.Run()
}
