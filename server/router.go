package server

import (
	"net/http"

	"github.com/akaahmedkamal/go-cli/v1"
)

type Router struct {
	app    *cli.App
	routes map[string]Route
}

func NewRouter(app *cli.App) *Router {
	return &Router{app: app, routes: make(map[string]Route)}
}

func (r *Router) Get(pattern string, route Route) {
	r.routes[pattern] = route
}

func (r *Router) Post(pattern string, route Route) {
	r.routes[pattern] = route
}

func (r *Router) Put(pattern string, route Route) {
	r.routes[pattern] = route
}

func (r *Router) Delete(pattern string, route Route) {
	r.routes[pattern] = route
}

func (r *Router) All(pattern string, route Route) {
	r.routes[pattern] = route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error

	route := r.routes[req.URL.Path]

	if route == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404 page not found"))
		return
	}

	_req := new(HttpRequest)
	_req.App = r.app
	_req.Request = req

	_res := new(HttpResponse)
	_res.w = w

	if err = route.HandleRequest(_req, _res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}
