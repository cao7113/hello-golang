package lab

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestTime(t *testing.T) {
	tm := time.Now().Add(-3 * time.Hour)
	logrus.Infof("app config: %s", tm)
}

func TestTimezone(t *testing.T) {
	// todo china location
}

type obj struct {
	t1 *time.Time
	t2 *time.Time
}

func TestTimevar(t *testing.T) {

	o := &obj{}
	tm1 := time.Now()
	tm2 := time.Now().Add(-3 * 24 * time.Hour)

	var tm *time.Time
	tm = &tm1
	o.t1 = tm

	tm = &tm2
	o.t2 = tm
	logrus.Infof("obj: %+v", o)
	logrus.Infof("t1 = %s t2 = %s", o.t1, o.t2)
}
