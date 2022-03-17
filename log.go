package skyjet

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewJsonLogger Creates a new logger, pre-configured to use JSONFormatter.
func NewJsonLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.JSONFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: true,
	}
}
