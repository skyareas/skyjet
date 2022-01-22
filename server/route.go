package server

import (
	"net/http"
	"strings"
)

type Route interface {
	HandleRequest(*HttpRequest, *HttpResponse) error
}

type RouteEntry struct {
	pattern              string
	method               string
	route                Route
	pathMatchingStrategy PathMatchingStrategy
}

func (ref *RouteEntry) Match(req *http.Request) (bool, map[string][]string) {
	params := make(map[string][]string)

	// match exactly
	if ref.pattern == req.URL.Path && ref.method == req.Method {
		return true, params
	}

	// match wildcard
	if (ref.pattern == "*" && ref.method == "*") ||
		(ref.pattern == req.URL.Path && ref.method == "*") ||
		(ref.pattern == "*" && ref.method == req.Method) {
		return true, params
	}

	// match prefix
	if ref.pathMatchingStrategy == PathMatchingStrategyPrefix &&
		strings.HasPrefix(req.URL.Path, ref.pattern) &&
		(ref.method == "*" || ref.method == req.Method) {
		return true, params
	}

	// match with params
	if strings.ContainsRune(ref.pattern, ':') {
		return ref.matchWithParams(req.URL.Path)
	}

	// doesn't match
	return false, params
}

func (ref *RouteEntry) matchWithParams(path string) (bool, map[string][]string) {
	patternParts := strings.Split(ref.pattern, "/")
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
