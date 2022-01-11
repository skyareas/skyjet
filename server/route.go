package server

import (
	"net/http"
	"strings"
)

type Route interface {
	HandleRequest(*HttpRequest, *HttpResponse) error
}

type RouteEntry struct {
	pattern string
	method  string
	route   Route
}

func (r *RouteEntry) Match(req *http.Request, strategy PathMatchingStrategy) (bool, map[string][]string) {
	params := make(map[string][]string)

	// match exactly
	if r.pattern == req.URL.Path && r.method == req.Method {
		return true, params
	}

	// match wildcard
	if (r.pattern == "*" && r.method == "*") ||
		(r.pattern == req.URL.Path && r.method == "*") ||
		(r.pattern == "*" && r.method == req.Method) {
		return true, params
	}

	// match prefix
	if strategy == PathMatchingStrategyPrefix &&
		strings.HasPrefix(req.URL.Path, r.pattern) &&
		(r.method == "*" || r.method == req.Method) {
		return true, params
	}

	// match with params
	if strings.ContainsRune(r.pattern, ':') {
		return r.matchWithParams(req.URL.Path)
	}

	// doesn't match
	return false, params
}

func (r *RouteEntry) matchWithParams(path string) (bool, map[string][]string) {
	patternParts := strings.Split(r.pattern, "/")
	pathParts := strings.Split(path, "/")
	params := make(map[string][]string)

	if len(patternParts) != len(pathParts) {
		return false, params
	}

	for i, part := range patternParts {
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = append(params[part[1:]], pathParts[i])
			continue
		}
		if part != pathParts[i] {
			return false, params
		}
	}

	return true, params
}
