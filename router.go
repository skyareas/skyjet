package skyjet

import (
	"net/http"
)

type Router struct {
	routes []*RouteEntry
}

type PathMatchingStrategy int

const (
	PathMatchingStrategyExact  PathMatchingStrategy = 0
	PathMatchingStrategyPrefix PathMatchingStrategy = 1
)

func NewRouter() *Router {
	return &Router{routes: make([]*RouteEntry, 0)}
}

func (r *Router) append(pattern, method string, route Route, pathMatchingStrategy PathMatchingStrategy) {
	if pattern == "" {
		app.log.Fatalln("server: invalid pattern")
	}
	if route == nil {
		app.log.Fatalln("server: nil route")
	}
	r.routes = append(r.routes, &RouteEntry{
		pattern:              pattern,
		method:               method,
		route:                route,
		pathMatchingStrategy: pathMatchingStrategy,
	})
}

func (r *Router) Get(pattern string, route Route) {
	r.append(pattern, http.MethodGet, route, PathMatchingStrategyExact)
}

func (r *Router) Post(pattern string, route Route) {
	r.append(pattern, http.MethodPost, route, PathMatchingStrategyExact)
}

func (r *Router) Put(pattern string, route Route) {
	r.append(pattern, http.MethodPut, route, PathMatchingStrategyExact)
}

func (r *Router) Delete(pattern string, route Route) {
	r.append(pattern, http.MethodDelete, route, PathMatchingStrategyExact)
}

func (r *Router) All(pattern string, route Route) {
	r.append(pattern, "*", route, PathMatchingStrategyExact)
}

func (r *Router) Use(pattern string, routes ...Route) {
	for _, route := range routes {
		r.append(pattern, "*", route, PathMatchingStrategyPrefix)
	}
}

func (r *Router) UseFilter(pattern string, routes ...Route) {
	for _, route := range routes {
		r.append(pattern, "*", route, PathMatchingStrategyPrefix)
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var recovered bool

	defer func() {
		if rec := recover(); rec != nil {
			recovered = true
			w.WriteHeader(http.StatusInternalServerError)
			r.writeError(w, rec.(error).Error())
		}
	}()

	var found bool

	_req := NewHttpRequest(req, map[string]string{})
	_res := NewHttpResponse(req, w)

	for _, entry := range r.routes {
		match, params := entry.Match(req)
		if match {
			found = true

			_req.params = params

			if err := entry.route.HandleRequest(_req, _res); err != nil {
				panic(err)
			}

			if _res.Sent() {
				break
			}
		}
	}

	if !found && !recovered {
		w.WriteHeader(http.StatusNotFound)
		r.writeError(w, "404 page not found")
	}
}

func (r *Router) writeError(w http.ResponseWriter, errMsg string) {
	if _, err := w.Write([]byte(errMsg)); err != nil {
		app.log.Fatalln(err)
	}
}
