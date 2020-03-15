package config

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEnvConfig(t *testing.T) {
	assert.True(t, IsTest())
	logrus.Infof("app config: %+v", Config)
}
