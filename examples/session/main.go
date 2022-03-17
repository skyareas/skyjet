package main

import (
	"github.com/skyareas/skyjet"
	"net/http"
	"time"
)

type AuthController struct{}

func (c *AuthController) HandleRequest(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	if req.Request.Method == http.MethodGet && req.Request.URL.Path == "/auth/login" {
		year := time.Now().Format("2006")
		return res.Render("login.html", skyjet.D{"CurrentYear": year})
	}
	if req.Request.Method == http.MethodPost && req.Request.URL.Path == "/auth/login" {

	}
	if req.Request.Method == http.MethodPost && req.Request.URL.Path == "/auth/logout" {

	}
	return nil
}

type AuthFilter struct{}

func (c *AuthFilter) HandleRequest(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	if !req.Session.IsValid() {
		res.Redirect("/auth/login", http.StatusPermanentRedirect)
	}
	return nil
}

type HomeController struct{}

func (c *HomeController) HandleRequest(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	username := req.Session.Get("username")
	year := time.Now().Format("2006")
	return res.Render("index.html", skyjet.D{"Username": username, "CurrentYear": year})
}

func main() {
	app := skyjet.SharedApp()
	app.Use("/auth", &AuthController{})
	app.Use("/", &AuthFilter{}, &HomeController{})
	app.Run()
}
