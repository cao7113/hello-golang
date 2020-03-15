package try

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	println("test log")
}

func TestLogrus(t *testing.T) {
	logrus.Infoln("hello", "logrus")
}
