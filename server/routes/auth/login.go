package auth

import (
	"github.com/akaahmedkamal/go-server/server"
)

type Login struct {
}

func (l *Login) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Send(server.D{"route": "login"})
}
