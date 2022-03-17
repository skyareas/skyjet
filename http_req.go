package skyjet

import (
	"net/http"
	"net/url"
)

// HttpRequest struct represents an Http request.
type HttpRequest struct {
	Request *http.Request
	Body    *HttpRequestBody
	User    HttpRequestUser
	Session *HttpRequestSession
	params  map[string]string
}

func NewHttpRequest(req *http.Request, params map[string]string) *HttpRequest {
	return &HttpRequest{
		Request: req,
		Body:    &HttpRequestBody{},
		User:    nil,
		Session: NewSession(req),
		params:  params,
	}
}

// Param returns the named path parameter.
func (r *HttpRequest) Param(name string) string {
	return r.params[name]
}

// LookupParam returns the named path parameter, and a bool
// value indicates whether the named param exists.
func (r *HttpRequest) LookupParam(name string) (string, bool) {
	p, ok := r.params[name]
	return p, ok
}

// Query returns the URL query parameters.
func (r *HttpRequest) Query() url.Values {
	return r.Request.URL.Query()
}

// ReadBody reads the raw request body into the Body struct property.
func (r *HttpRequest) ReadBody() error {
	return r.Body.Read(r.Request.Body)
}
