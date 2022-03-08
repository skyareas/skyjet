package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akaahmedkamal/go-server/config"
)

// HttpResponse struct represents an Http response.
type HttpResponse struct {
	r      *http.Request
	w      http.ResponseWriter
	Header http.Header
	sent   bool
}

func NewHttpResponse(req *http.Request, w http.ResponseWriter) *HttpResponse {
	return &HttpResponse{r: req, w: w, Header: w.Header(), sent: false}
}

// Writer returns a pointer to the underlying http.ResponseWriter
func (ref *HttpResponse) Writer() http.ResponseWriter {
	return ref.w
}

// Download prompts a file to be downloaded.
func (ref *HttpResponse) Download() {}

// End ends the response process.
func (ref *HttpResponse) End() {
	ref.sent = true
}

// Json sends a JSON response.
func (ref *HttpResponse) Json(v interface{}, statusCode ...int) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	ref.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return ref.Send(b, statusCode...)
}

// Jsonp sends a JSON response with JSONP support.
func (ref *HttpResponse) Jsonp(d interface{}) error {
	return nil
}

// Redirect redirects a request, optionally specifying the status code.
func (ref *HttpResponse) Redirect(url string, statusCode ...int) {
	code := http.StatusMovedPermanently
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	http.Redirect(ref.w, ref.r, url, code)
	ref.End()
}

// Render renders a view template with data.
func (ref *HttpResponse) Render(temp string, data ...interface{}) error {
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

	ref.w.Header().Set("Content-Type", "text/html")
	err = t.Execute(ref.w, d)
	ref.End()

	return err
}

// Send sends a response of raw bytes value,
// optionally specifying the status code.
func (ref *HttpResponse) Send(v []byte, statusCode ...int) error {
	if len(statusCode) > 0 {
		ref.Status(statusCode[0])
	}
	_, err := ref.w.Write(v)
	ref.End()
	return err
}

// SendFile sends a file as an octet stream,
// optionally specifying the status code.
func (ref *HttpResponse) SendFile(filename string, statusCode ...int) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	m := mime.TypeByExtension(filepath.Ext(filename))
	if err != nil {
		return err
	}

	if m == "" {
		m = "application/octet-stream"
	}

	ref.w.Header().Set("Content-Type", m)

	return ref.Send(b, statusCode...)
}

// Status set the response status code.
func (ref *HttpResponse) Status(statusCode int) {
	ref.w.WriteHeader(statusCode)
}

// SendStatus set the response status code and send its
// string representation as the response body.
func (ref *HttpResponse) SendStatus(statusCode int) error {
	ref.Status(statusCode)
	_, err := ref.w.Write([]byte(http.StatusText(statusCode)))
	ref.End()
	return err
}

// Sent returns a bool indicates whether the response
// has been sent to the client.
func (ref *HttpResponse) Sent() bool {
	return ref.sent
}

// SetCookie a wrapper around the http.SetCookie() function.
func (ref *HttpResponse) SetCookie(cookie *http.Cookie) {
	http.SetCookie(ref.w, cookie)
}
