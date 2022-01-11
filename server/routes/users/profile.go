package users

import "github.com/akaahmedkamal/go-server/server"

// Profile implements the user profile route.
type Profile struct {
}

// HandleRequest handles the incoming requests.
func (p *Profile) HandleRequest(req *server.HttpRequest, res *server.HttpResponse) error {
	id, _ := req.ParamString("id")
	return res.Render("users/profile.html", server.D{"ID": id})
}
