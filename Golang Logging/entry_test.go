package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestXxx(t *testing.T) {
	logger:= logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Info("Hello Logging")
	logger.WithField("username","Jason").Info("Hello Logging")

	entry:=logrus.NewEntry(logger)
	entry.WithField("username","Keshinryan")
	entry.Info("Hello Entry")
}