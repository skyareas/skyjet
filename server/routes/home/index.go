package home

import "github.com/akaahmedkamal/go-server/server"

// Index implements the home/index route.
type Index struct {
}

// HandleRequest handles the incoming requests.
func (ref *Index) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Render("index.html")
}
