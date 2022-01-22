package server

import (
	"net/http"

	"github.com/akaahmedkamal/go-server/app"
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
}

func NewRouter(cfg *RouterConfig) *Router {
	return &Router{config: cfg, routes: make([]*RouteEntry, 0)}
}

func (ref *Router) append(pattern, method string, route Route, pathMatchingStrategy PathMatchingStrategy) {
	if pattern == "" {
		app.Shared().Log().Fatalln("server: invalid pattern")
	}
	if route == nil {
		app.Shared().Log().Fatalln("server: nil route")
	}
	ref.routes = append(ref.routes, &RouteEntry{
		pattern:              pattern,
		method:               method,
		route:                route,
		pathMatchingStrategy: pathMatchingStrategy,
	})
}

// func (ref *Router) appendOrdered(pattern, method string, route Route) {
// 	if pattern == "" {
// 		app.Shared().Log().Fatalln("server: invalid pattern")
// 	}
// 	if route == nil {
// 		app.Shared().Log().Fatalln("server: nil route")
// 	}

// 	i := ref.searchEntry(func(idx int) bool {
// 		return len(pattern) > len(ref.routes[idx].pattern)
// 	})

// 	ref.routes = append(ref.routes, nil)
// 	copy(ref.routes[i+1:], ref.routes[i:])
// 	ref.routes[i] = &RouteEntry{
// 		pattern: pattern,
// 		method:  method,
// 		route:   route,
// 	}
// }

// func (ref *Router) searchEntry(f func(i int) bool) int {
// 	i, j := 0, len(ref.routes)
// 	for i < j {
// 		h := int(uint(i+j) >> 1)
// 		if !f(h) {
// 			i = h + 1
// 		} else {
// 			j = h
// 		}
// 	}
// 	return i
// }

// func (ref *Router) hasEntry(pattern string) bool {
// 	for _, entry := range ref.routes {
// 		if entry.pattern == pattern {
// 			return true
// 		}
// 	}
// 	return false
// }

func (ref *Router) Get(pattern string, route Route) {
	ref.append(pattern, http.MethodGet, route, PathMatchingStrategyExact)
}

func (ref *Router) Post(pattern string, route Route) {
	ref.append(pattern, http.MethodPost, route, PathMatchingStrategyExact)
}

func (ref *Router) Put(pattern string, route Route) {
	ref.append(pattern, http.MethodPut, route, PathMatchingStrategyExact)
}

func (ref *Router) Delete(pattern string, route Route) {
	ref.append(pattern, http.MethodDelete, route, PathMatchingStrategyExact)
}

func (ref *Router) All(pattern string, route Route) {
	ref.append(pattern, "*", route, PathMatchingStrategyExact)
}

func (ref *Router) Use(pattern string, routes ...Route) {
	for _, route := range routes {
		ref.append(pattern, "*", route, PathMatchingStrategyPrefix)
	}
}

func (ref *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var recovered bool

	defer func() {
		if rec := recover(); rec != nil {
			recovered = true
			w.WriteHeader(http.StatusInternalServerError)
			ref.writeError(w, rec.(error).Error())
		}
	}()

	var found bool
	for _, entry := range ref.routes {
		match, params := entry.Match(req)
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

	if !found && !recovered {
		w.WriteHeader(http.StatusNotFound)
		ref.writeError(w, "404 page not found")
	}
}

func (ref *Router) writeError(w http.ResponseWriter, errMsg string) {
	if _, err := w.Write([]byte(errMsg)); err != nil {
		panic(err)
	}
}
