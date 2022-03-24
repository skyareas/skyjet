package skyjet

import (
	"net/http"
	"strings"
)

const HttpMethodAll = "All"

type Router struct {
	routes []RouteEntry
}

func NewRouter() *Router {
	return &Router{make([]RouteEntry, 0)}
}

func (r *Router) Get(pattern string, handlers ...RouteHandler) {
	for _, handler := range handlers {
		r.append(pattern, http.MethodGet, handler)
	}
}

func (r *Router) Post(pattern string, handlers ...RouteHandler) {
	for _, handler := range handlers {
		r.append(pattern, http.MethodPost, handler)
	}
}

func (r *Router) Put(pattern string, handlers ...RouteHandler) {
	for _, handler := range handlers {
		r.append(pattern, http.MethodPut, handler)
	}
}

func (r *Router) Delete(pattern string, handlers ...RouteHandler) {
	for _, handler := range handlers {
		r.append(pattern, http.MethodDelete, handler)
	}
}

func (r *Router) All(pattern string, handlers ...RouteHandler) {
	for _, handler := range handlers {
		r.append(pattern, HttpMethodAll, handler)
	}
}

func (r *Router) UseMiddleware(pattern string, handlers ...RouteHandler) {
	for _, handler := range handlers {
		r.append(pattern, HttpMethodAll, handler, RouteMatchingStrategyPrefix)
	}
}

func (r *Router) Use(pattern string, router *Router) {
	for _, route := range router.routes {
		r.append(pattern+route.pattern, route.method, route.handler, route.matching)
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, httpRequest *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			w.WriteHeader(http.StatusInternalServerError)
			switch v := rec.(type) {
			case string:
				r.writeError(w, v)
			case error:
				r.writeError(w, v.Error())
			default:
				app.Log().Fatalf("unable to recover from error: %v", v)
			}
		}
	}()

	ses := NewSession(httpRequest)
	req := NewHttpRequest(httpRequest, ses)
	res := NewHttpResponse(httpRequest, w, ses)

	err := r.handleRequest(req, res)
	if err != nil {
		panic(err)
	}

	if !res.sent {
		w.WriteHeader(http.StatusNotFound)
		r.writeError(w, http.StatusText(http.StatusNotFound))
	}
}

func (r *Router) append(pattern, method string, handler RouteHandler, matching ...RouteMatchingStrategy) {
	if pattern == "" {
		app.log.Fatalln("server: invalid pattern")
	}
	if handler == nil {
		app.log.Fatalln("server: nil handler")
	}

	var m RouteMatchingStrategy
	if len(matching) > 0 {
		m = matching[0]
	} else {
		m = RouteMatchingStrategyExact
	}

	r.routes = append(r.routes, RouteEntry{
		pattern:  r.cleanPattern(pattern),
		method:   method,
		handler:  handler,
		matching: m,
	})
}

func (r *Router) handleRequest(req *HttpRequest, res *HttpResponse) error {
	for _, entry := range r.routes {
		match, params := entry.Match(req.Request)
		if match {
			req.params = params

			if err := entry.handler(req, res); err != nil {
				return err
			}

			if res.Sent() {
				break
			}
		}
	}
	return nil
}

func (r *Router) writeError(w http.ResponseWriter, errMsg string) {
	if _, err := w.Write([]byte(errMsg)); err != nil {
		app.log.Fatalln(err)
	}
}

func (r *Router) cleanPattern(pattern string) string {
	p := strings.ReplaceAll(pattern, "//", "/")
	if len(p) > 1 && p[len(p)-1] == '/' {
		return p[:len(p)-1]
	}
	return p
}
