package auth

import "github.com/akaahmedkamal/go-server/server"

// Logout implements the logout route.
type Logout struct {
}

// HandleRequest handles the incoming requests.
func (l *Logout) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	// TODO: implement the logout logic!
	res.Redirect("/")
	return nil
}
