package skyjet

import (
	"net/http"
	"strings"
)

type RouteMatchingStrategy int

const (
	RouteMatchingStrategyExact  RouteMatchingStrategy = 0
	RouteMatchingStrategyPrefix RouteMatchingStrategy = 1
)

type RouteEntry struct {
	pattern  string
	method   string
	handler  RouteHandler
	matching RouteMatchingStrategy
}

func (r *RouteEntry) Match(req *http.Request) (bool, map[string]string) {
	params := make(map[string]string)

	// match exactly
	if r.pattern == req.URL.Path && r.method == req.Method {
		return true, params
	}

	// match wildcard
	if (r.pattern == "*" && r.method == HttpMethodAll) ||
		(r.pattern == req.URL.Path && r.method == HttpMethodAll) ||
		(r.pattern == "*" && r.method == req.Method) {
		return true, params
	}

	// match prefix
	if ((r.matching == RouteMatchingStrategyPrefix && strings.HasPrefix(req.URL.Path, r.pattern)) ||
		(strings.HasSuffix(r.pattern, "*") && strings.HasPrefix(req.URL.Path, strings.TrimSuffix(r.pattern, "*")))) &&
		(r.method == HttpMethodAll || r.method == req.Method) {
		return true, params
	}

	// match with params
	if strings.ContainsRune(r.pattern, ':') {
		if r.method == HttpMethodAll || r.method == req.Method {
			return r.matchWithParams(req.URL.Path)
		}
	}

	// doesn't match
	return false, params
}

func (r *RouteEntry) matchWithParams(path string) (bool, map[string]string) {
	patternParts := strings.Split(r.pattern, "/")
	pathParts := strings.Split(path, "/")
	params := make(map[string]string)

	if len(patternParts) != len(pathParts) {
		return false, params
	}

	for i, part := range patternParts {
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = pathParts[i]
			continue
		}
		if part != pathParts[i] {
			return false, params
		}
	}

	return true, params
}
