package server

import (
	"net/http"

	"github.com/akaahmedkamal/go-cli/v1"
)

type HttpRequest struct {
	Request *http.Request
	App     *cli.App
}
