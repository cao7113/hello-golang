package syntax

import (
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFor(t *testing.T) {
	i := 0
	for {
		i += 1
		logrus.Infof("i=%d", i)
		if i > 3 {
			break
		}
	}
	assert.Equal(t, 4, i)
}