package middleware

import (
	"strconv"
	"strings"

	"github.com/akaahmedkamal/go-server/server"
)

// Cors implements Http CORS middleware.
type Cors struct {
	Origin           []string
	Methods          []string
	Headers          []string
	AllowCredentials bool
}

// HandleRequest handles the incoming requests.
func (ref *Cors) HandleRequest(_ *server.HttpRequest, res *server.HttpResponse) error {
	res.Header.Set("Access-Control-Allow-Origin", strings.Join(ref.Origin, ","))
	res.Header.Set("Access-Control-Allow-Methods", strings.Join(ref.Methods, ","))
	res.Header.Set("Access-Control-Allow-Headers", strings.Join(ref.Headers, ","))
	res.Header.Set("Access-Control-Allow-Credentials", strconv.FormatBool(ref.AllowCredentials))
	return nil
}
