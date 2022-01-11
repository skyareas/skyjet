package auth

import (
	"net/http"

	"github.com/akaahmedkamal/go-server/server"
)

// Register implements the register route.
type Register struct {
}

// HandleRequest handles the incoming requests.
func (r *Register) HandleRequest(req *server.HttpRequest, res *server.HttpResponse) error {
	if req.Request.Method == http.MethodGet {
		return r.get(req, res)
	}
	if req.Request.Method == http.MethodPost {
		return r.post(req, res)
	}
	return res.SendStatus(http.StatusMethodNotAllowed)
}

// get handles incoming GET requests.
func (r *Register) get(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Render("auth/register.html")
}

// post handles incoming POST requests.
func (r *Register) post(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Json(server.D{"Route": "Register.Post"})
}
