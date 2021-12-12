package lab

import (
	"log"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOsPrint(t *testing.T) {
	println("test", "os println")
}

func TestLog(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(os.Stdout) // default is os.Stderr
	//println("test", "log")
	log.Println("test", "log")
}

func TestLogrus(t *testing.T) {
	logrus.Infoln("grpc", "logrus")
}
