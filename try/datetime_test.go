package try

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestTime(t *testing.T) {
	tm := time.Now().Add(-3 * time.Hour)
	logrus.Infof("app config: %s", tm)
}
