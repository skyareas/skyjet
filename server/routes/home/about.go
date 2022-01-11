package home

import "github.com/akaahmedkamal/go-server/server"

// About implements about route.
type About struct {
}

// HandleRequest handles the incoming requests.
func (i *About) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Render("about.html")
}
