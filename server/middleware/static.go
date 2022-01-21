package middleware

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/akaahmedkamal/go-server/server"
)

// Static implements static files handler middleware.
type Static struct {
	PublicPath string
}

// HandleRequest handles the incoming requests.
func (ref *Static) HandleRequest(req *server.HttpRequest, res *server.HttpResponse) error {
	err := res.SendFile(filepath.Join(ref.PublicPath, req.Request.URL.Path))

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return res.Send([]byte(http.StatusText(http.StatusNotFound)), http.StatusNotFound)
	}

	return err
}
