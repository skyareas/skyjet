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

func (ref *RouteEntry) Match(req *http.Request) (bool, map[string]string) {
	params := make(map[string]string)

	// match exactly
	if ref.pattern == req.URL.Path && ref.method == req.Method {
		return true, params
	}

	// match wildcard
	if (ref.pattern == "*" && ref.method == HttpMethodAll) ||
		(ref.pattern == req.URL.Path && ref.method == HttpMethodAll) ||
		(ref.pattern == "*" && ref.method == req.Method) {
		return true, params
	}

	// match prefix
	if ref.matching == RouteMatchingStrategyPrefix &&
		strings.HasPrefix(req.URL.Path, ref.pattern) &&
		(ref.method == HttpMethodAll || ref.method == req.Method) {
		return true, params
	}

	// match with params
	if strings.ContainsRune(ref.pattern, ':') {
		if ref.method == HttpMethodAll || ref.method == req.Method {
			return ref.matchWithParams(req.URL.Path)
		}
	}

	// doesn't match
	return false, params
}

func (ref *RouteEntry) matchWithParams(path string) (bool, map[string]string) {
	patternParts := strings.Split(ref.pattern, "/")
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
