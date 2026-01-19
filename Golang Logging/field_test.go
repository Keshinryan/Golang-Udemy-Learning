package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger:=logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	
	logger.WithField("username","keshinryan").Info("Hello World")
	logger.WithField("username","Jason").WithField("name","Jason Patrick").Info("Hello World")

}

func TestFields(t *testing.T){
	logger:=logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username":"keshinryan",
		"name":"Jason Patrick",
	}).Info("Hello World")

}