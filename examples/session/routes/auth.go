package routes

import (
	"time"

	"github.com/skyareas/skyjet"
)

func NewAuthRouter() *skyjet.Router {
	r := skyjet.NewRouter()
	r.Get("/login", getLogin)
	r.Post("/login", postLogin)
	r.Post("/logout", logout)
	return r
}

func getLogin(_ *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	year := time.Now().Format("2006")
	return res.Render("login.html", skyjet.D{"CurrentYear": year})
}

func postLogin(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	err := req.Request.ParseForm()
	if err != nil {
		return err
	}
	skyjet.App().Log().Printf("%v\n", req.Request.Form)
	req.Session.Set("username", req.Request.Form.Get("username"))
	res.Redirect("/")
	return nil
}

func logout(_ *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	res.Redirect("/")
	return nil
}
