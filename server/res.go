package server

import (
	"encoding/json"
	"fmt"
	"github.com/akaahmedkamal/go-server/config"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// HttpResponse struct represents an Http response.
type HttpResponse struct {
	r    *http.Request
	w    http.ResponseWriter
	sent bool
}

func NewHttpResponse(req *http.Request, w http.ResponseWriter) *HttpResponse {
	return &HttpResponse{r: req, w: w, sent: false}
}

// Writer returns a pointer to the underlying http.ResponseWriter
func (r *HttpResponse) Writer() http.ResponseWriter {
	return r.w
}

// Download prompts a file to be downloaded.
func (r *HttpResponse) Download() {}

// End ends the response process.
func (r *HttpResponse) End() {
	r.sent = true
}

// Json sends a JSON response.
func (r *HttpResponse) Json(v interface{}, statusCode ...int) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err = r.w.Write(b); err != nil {
		return err
	}
	if len(statusCode) > 0 {
		r.Status(statusCode[0])
	}
	return err
}

// Jsonp sends a JSON response with JSONP support.
func (r *HttpResponse) Jsonp(d interface{}) error {
	return nil
}

// Redirect redirects a request, optionally specifying the status code.
func (r *HttpResponse) Redirect(url string, statusCode ...int) {
	code := http.StatusMovedPermanently
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	http.Redirect(r.w, r.r, url, code)
	r.End()
}

// Render renders a view template with data.
func (r *HttpResponse) Render(temp string, data ...interface{}) error {
	vp := config.Shared().Http.ViewsPath
	f, err := os.ReadFile(filepath.Join(vp, temp))
	if err != nil {
		return err
	}

	t := template.New(temp)
	if t, err = t.Parse(string(f)); err != nil {
		fmt.Println(err)
	}

	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}

	r.w.Header().Set("Content-Type", "text/html")
	return t.Execute(r.w, d)
}

// Send sends a response of raw bytes value,
// optionally specifying the status code.
func (r *HttpResponse) Send(v []byte, statusCode ...int) error {
	_, err := r.w.Write(v)
	if len(statusCode) > 0 {
		r.Status(statusCode[0])
	}
	return err
}

// SendFile sends a file as an octet stream,
// optionally specifying the status code.
func (r *HttpResponse) SendFile(filename string, statusCode ...int) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	r.Status(code)
	r.w.Header().Set("Content-Type", "application/octet-stream")
	_, err = r.w.Write(b)
	return err
}

// Status set the response status code.
func (r *HttpResponse) Status(statusCode int) {
	r.w.WriteHeader(statusCode)
}

// SendStatus set the response status code and send its
// string representation as the response body.
func (r *HttpResponse) SendStatus(statusCode int) error {
	r.Status(statusCode)
	_, err := r.w.Write([]byte(http.StatusText(statusCode)))
	return err
}

// Sent returns a bool indicates whether the response
// has been sent to the client.
func (r *HttpResponse) Sent() bool {
	return r.sent
}
