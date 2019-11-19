package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	viper.SetDefault("ContentDir", "content")

	viper.SetConfigName("settings") // name of config file (without extension)
	// viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			t.Log("not found")
		} else {
			// Config file was found but another error was produced
			t.Error(err)
		}
	}

	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })

}
