package routes

import (
	"net/http"

	"github.com/akaahmedkamal/go-server/server"
)

// NotFound implements page not found (404) route.
type NotFound struct {
}

// HandleRequest handles the incoming requests.
func (ref *NotFound) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Send([]byte(http.StatusText(http.StatusNotFound)), http.StatusNotFound)
}
