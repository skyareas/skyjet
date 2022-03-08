package server

import (
	"net/http"
)

// HttpRequest struct represents an Http request.
type HttpRequest struct {
	Request *http.Request
	Body    *HttpRequestBody
	User    HttpRequestUser
	Session *HttpRequestSession
	params  map[string][]string
}

func NewHttpRequest(req *http.Request, params map[string][]string) (*HttpRequest, error) {
	ses, err := NewSession(req)
	return &HttpRequest{
		Request: req,
		Body: &HttpRequestBody{},
		User: nil,
		Session: ses,
		params: params,
	}, err
}

// Param returns all values for a specific path parameter.
func (r *HttpRequest) Param(name string) ([]string, bool) {
	p, ok := r.params[name]
	return p, ok
}

// ParamString returns the first value FROM RIGHT
// for a specific path parameter as string.
func (r *HttpRequest) ParamString(name string) (string, bool) {
	p, ok := r.params[name]
	if len(p) == 0 {
		return "", false
	}
	return p[len(p)-1], ok
}

// ReadBody reads the raw request body into the Body struct property.
func (r *HttpRequest) ReadBody() error {
	return r.Body.Read(r.Request.Body)
}
