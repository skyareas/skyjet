package app

import (
	"os"

	"github.com/akaahmedkamal/go-cli/v1"
)

// shared instance of the code/cli app.
var shared *cli.App

// Shared returns the shared/global instance of the
// core/cli app if found, otherwise, it initializes
// a new shared instance and returns it.
func Shared() *cli.App {
	if shared == nil {
		shared = cli.NewApp(os.Args[1:])
	}
	return shared
}
