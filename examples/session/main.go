package main

import (
	"github.com/skyareas/skyjet"
	"github.com/skyareas/skyjet/examples/session/middleware"
	"github.com/skyareas/skyjet/examples/session/routes"
)

func main() {
	app := skyjet.SharedApp()
	app.Use("/auth", routes.NewAuthRouter())
	app.Middleware("/", middleware.AuthMiddleware)
	app.Use("/", routes.NewHomeRouter())
	app.Run()
}
