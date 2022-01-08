package auth

import "github.com/akaahmedkamal/go-server/server"

type Register struct {
}

func (r *Register) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	return res.Send(server.D{"route": "register"})
}
