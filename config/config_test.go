package config_test

import (
	"os"
	"testing"

	"github.com/cao7113/golang/config"
	"github.com/stretchr/testify/assert"
)

func TestAppEnv(t *testing.T) {
	os.Setenv("APP_ENV", "")
	assert.Equal(t, true, config.IsDev())
}

func TestDbURL(t *testing.T) {
	url := "mysql://xxx"
	os.Setenv("DB_URL", url)
	assert.NotEqual(t, url, config.Settings.DbURL)
}
