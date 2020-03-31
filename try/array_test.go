package try

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestArray(t *testing.T) {
	var ints []int
	ints = append(ints, 2, 3)
	logrus.Infof("integers=%+v", ints)

	strs := []string{
		"hello",
		"world",
	}
	logrus.Infof("strings: %+v", strs)
}
