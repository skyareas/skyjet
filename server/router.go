package server

import (
	"github.com/akaahmedkamal/go-server/app"
	"net/http"
)

type Router struct {
	config *RouterConfig
	routes []*RouteEntry
}

type PathMatchingStrategy int

const (
	PathMatchingStrategyExact  PathMatchingStrategy = 0
	PathMatchingStrategyPrefix PathMatchingStrategy = 1
)

type RouterConfig struct {
	PathMatchingStrategy PathMatchingStrategy
}

func NewRouter(cfg *RouterConfig) *Router {
	return &Router{config: cfg, routes: make([]*RouteEntry, 0)}
}

func (r *Router) appendOrdered(pattern, method string, route Route) {
	if pattern == "" {
		app.Shared().Log().Fatalln("server: invalid pattern")
	}
	if route == nil {
		app.Shared().Log().Fatalln("server: nil route")
	}
	if r.hasEntry(pattern) {
		app.Shared().Log().Fatalf("server: multiple registrations for \"%s\"\n", pattern)
	}

	i := r.searchEntry(func(idx int) bool {
		return len(pattern) > len(r.routes[idx].pattern)
	})

	r.routes = append(r.routes, nil)
	copy(r.routes[i+1:], r.routes[i:])
	r.routes[i] = &RouteEntry{
		pattern: pattern,
		method:  method,
		route:   route,
	}
}

func (r *Router) searchEntry(f func(i int) bool) int {
	i, j := 0, len(r.routes)
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (r *Router) hasEntry(pattern string) bool {
	for _, entry := range r.routes {
		if entry.pattern == pattern {
			return true
		}
	}
	return false
}

func (r *Router) Get(pattern string, route Route) {
	r.appendOrdered(pattern, http.MethodGet, route)
}

func (r *Router) Post(pattern string, route Route) {
	r.appendOrdered(pattern, http.MethodPost, route)
}

func (r *Router) Put(pattern string, route Route) {
	r.appendOrdered(pattern, http.MethodPut, route)
}

func (r *Router) Delete(pattern string, route Route) {
	r.appendOrdered(pattern, http.MethodDelete, route)
}

func (r *Router) All(pattern string, route Route) {
	r.appendOrdered(pattern, "*", route)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			w.WriteHeader(http.StatusInternalServerError)
			r.writeError(w, rec.(error).Error())
		}
	}()

	var found bool
	for _, entry := range r.routes {
		match, params := entry.Match(req, r.config.PathMatchingStrategy)
		if match {
			found = true

			_req := NewHttpRequest(req, params)
			_res := NewHttpResponse(req, w)
			if err := entry.route.HandleRequest(_req, _res); err != nil {
				panic(err)
			}

			if _res.Sent() {
				break
			}
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		r.writeError(w, "404 page not found")
	}
}

func (r *Router) writeError(w http.ResponseWriter, errMsg string) {
	if _, err := w.Write([]byte(errMsg)); err != nil {
		panic(err)
	}
}
