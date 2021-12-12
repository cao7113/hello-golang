package enumer

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGender(t *testing.T) {
	logrus.Infoln(Male.String())
}
