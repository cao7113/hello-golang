package try

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

func TestFor(t *testing.T) {
	i := 0
	cnt := 3
	for {
		i += 1
		logrus.Infof("i=%d", i)
		if i > cnt {
			break
		}
	}
	assert.Equal(t, cnt+1, i)
}

func TestForSlice(t *testing.T) {
	m := []int{1, 2, 3}
	// idx 0-based
	for idx, e := range m {
		logrus.Infoln(idx, e)
	}
}

func TestForRange(t *testing.T) {
	m := map[string]string{
		"name": "cao",
	}

	for k, v := range m {
		logrus.Infoln(k, v)
	}
}
