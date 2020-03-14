package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// AppEnv run mode e.g. test
var AppEnv string

// AppRoot root dir
var AppRoot string

// Config from env
var Config EnvConfig

// EnvConfig thanks to https://github.com/caarlos0/env
type EnvConfig struct {
	AppEnv  string
	Verbose bool `env:"VERBOSE"`

	// Settings
	DbURL string `env:"DATABASE_URL,required"`

	// DingTalk
	DingdingURL   string `env:"DINGDING_URL" envDefault:"https://oapi.dingtalk.com/robot/send"`
	DingdingToken string `env:"DINGDING_TOKEN,required"`
}

func init() {
	if AppEnv = strings.ToLower(os.Getenv("APP_ENV")); "" == AppEnv {
		AppEnv = "test"
	}
	logrus.Infof("run in env: %s ", AppEnv)

	if !IsProd() { // bug in docker container
		AppRoot = GetAppRoot("go.mod")
		err := loadEnvFiles(AppRoot)
		if err != nil {
			logrus.Fatalln(err)
		}
	}

	Config = EnvConfig{
		AppEnv: AppEnv, // just ref top config
	}
	if err := env.Parse(&Config); err != nil {
		logrus.Fatalf("invalid config: %+v", err)
	}
	logrus.Infoln("env config loaded!")
	if Verbose() {
		logrus.Infof("config: %+v", Config)
	}
}

// IsTest in testing mode
func IsTest() bool {
	return AppEnv == "test"
}

// IsProd in production mode
func IsProd() bool {
	return AppEnv == "production"
}

// Verbose more info
func Verbose() bool {
	return Config.Verbose
}

// GetAppRoot contains a file
func GetAppRoot(rootFile string) string {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Fatalln("can not get working dir")
	}

	for {
		file := filepath.Join(wd, rootFile)
		if fileExists(file) {
			return wd
		}

		wd = filepath.Dir(wd)

		if "/" == wd {
			logrus.Fatalln("not found app root")
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func loadEnvFiles(dir string) error {
	parts := []string{"", ".local"}
	for _, part := range parts {
		file := strings.Join([]string{dir, ".env." + AppEnv + part}, "/")
		if fileExists(file) {
			err := godotenv.Overload(file)
			if err != nil {
				return err
			}
			logrus.Infof(".env file: %s", file)
		}
	}

	return nil
}
