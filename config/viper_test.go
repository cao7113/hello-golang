package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestViper1(t *testing.T) {
	viper.SetDefault("ContentDir", "content")
}
