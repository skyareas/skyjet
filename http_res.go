package skyjet

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// HttpResponse struct represents an Http response.
type HttpResponse struct {
	r       *http.Request
	w       http.ResponseWriter
	Header  http.Header
	session *HttpRequestSession
	sent    bool
}

func NewHttpResponse(req *http.Request, w http.ResponseWriter, session *HttpRequestSession) *HttpResponse {
	return &HttpResponse{r: req, w: w, Header: w.Header(), session: session, sent: false}
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
	r.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return r.Send(b, statusCode...)
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

func (r *HttpResponse) templateFiles() []string {
	files := make([]string, 0)
	vp := app.cfg.Http.ViewsPath
	_ = filepath.WalkDir(vp, func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

// Render renders a view template with data.
func (r *HttpResponse) Render(temp string, data ...interface{}) error {
	vp := app.cfg.Http.ViewsPath
	f, err := os.ReadFile(filepath.Join(vp, temp))
	if err != nil {
		return err
	}

	t, err := template.ParseFiles(r.templateFiles()...)
	if err != nil {
		return err
	}

	if t, err = t.Parse(string(f)); err != nil {
		return err
	}

	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}

	if err = r.setSession(); err != nil {
		return err
	}

	r.w.Header().Set("Content-Type", "text/html")
	if err = t.Execute(r.w, d); err != nil {
		return err
	}

	r.End()
	return nil
}

// Send sends a response of raw bytes value,
// optionally specifying the status code.
func (r *HttpResponse) Send(v []byte, statusCode ...int) error {
	if len(statusCode) > 0 {
		r.Status(statusCode[0])
	}

	if err := r.setSession(); err != nil {
		return err
	}

	if _, err := r.w.Write(v); err != nil {
		return err
	}

	r.End()
	return nil
}

// SendFile sends a file as an octet stream,
// optionally specifying the status code.
func (r *HttpResponse) SendFile(filename string, statusCode ...int) error {
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

	r.w.Header().Set("Content-Type", m)

	return r.Send(b, statusCode...)
}

// Status set the response status code.
func (r *HttpResponse) Status(statusCode int) {
	r.w.WriteHeader(statusCode)
}

// SendStatus set the response status code and send its
// string representation as the response body.
func (r *HttpResponse) SendStatus(statusCode int) error {
	return r.Send([]byte(http.StatusText(statusCode)), statusCode)
}

// Sent returns a bool indicates whether the response
// has been sent to the client.
func (r *HttpResponse) Sent() bool {
	return r.sent
}

// SetCookie a wrapper around the http.SetCookie() function.
func (r *HttpResponse) SetCookie(cookie *http.Cookie) {
	http.SetCookie(r.w, cookie)
}

func (r *HttpResponse) setSession() error {
	ses, err := r.session.Cookie()
	if err != nil {
		return err
	}
	r.SetCookie(ses)
	return nil
}
