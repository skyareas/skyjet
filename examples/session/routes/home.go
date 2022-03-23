package routes

import (
	"time"

	"github.com/skyareas/skyjet"
)

func NewHomeRouter() *skyjet.Router {
	r := skyjet.NewRouter()
	r.Get("/", home)
	r.Get("/about", about)
	return r
}

func home(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	username := req.Session.Get("username")
	year := time.Now().Format("2006")
	return res.Render("index.html", skyjet.D{"Username": username, "CurrentYear": year})
}

func about(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	username := req.Session.Get("username")
	year := time.Now().Format("2006")
	return res.Render("about.html", skyjet.D{"Username": username, "CurrentYear": year})
}
