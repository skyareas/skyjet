package auth

import (
	"github.com/akaahmedkamal/go-server/server"
)

type Logout struct {
}

func (l *Logout) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Send(server.D{"route": "logout"})
}
