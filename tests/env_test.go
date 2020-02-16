package tests

import (
	"testing"

	"github.com/cao7113/hellogolang/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEnvConfig(t *testing.T) {
	assert.True(t, config.IsTest())
	logrus.Infof("app config: %+v", config.Config)
}
