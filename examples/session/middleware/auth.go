package middleware

import (
	"net/http"

	"github.com/skyareas/skyjet"
)

func AuthMiddleware(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	if !req.Session.IsValid() {
		res.Redirect("/auth/login", http.StatusPermanentRedirect)
	}
	return nil
}
