package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cao7113/golang/config"
	"github.com/stretchr/testify/assert"
)

func TestDbURL(t *testing.T) {
	url := "mysql://xxx"
	os.Setenv("DB_URL", url)
	assert.NotEqual(t, url, config.Settings.DbURL)
}

func TestConf(t *testing.T) {
	config.ConfInfo()
}

func TestAppRoot(t *testing.T) {
	fmt.Println(config.GetAppRoot())
}
