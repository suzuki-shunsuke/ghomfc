package log

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func New(version string) *logrus.Entry {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		DisableQuote: true,
	}
	return logger.WithFields(logrus.Fields{
		"version": version,
		"program": "ghomfc",
		"env":     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	})
}
