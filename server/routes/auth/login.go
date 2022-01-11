package auth

import (
	"net/http"

	"github.com/akaahmedkamal/go-server/server"
)

// Login implements the login route.
type Login struct {
}

// HandleRequest handles the incoming requests.
func (l *Login) HandleRequest(req *server.HttpRequest, res *server.HttpResponse) error {
	if req.Request.Method == http.MethodGet {
		return l.get(req, res)
	}
	if req.Request.Method == http.MethodPost {
		return l.post(req, res)
	}
	return res.SendStatus(http.StatusMethodNotAllowed)
}

// get handles incoming GET requests.
func (l *Login) get(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Render("auth/login.html")
}

// post handles incoming POST requests.
func (l *Login) post(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Json(server.D{"Route": "Login.Post"})
}
