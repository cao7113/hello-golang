package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// AppEnv app running env
var AppEnv string

// Settings config from env
var Settings *config

// RawLogger a new instance of the logger. You can have any number of instances.
var RawLogger *logrus.Logger

// Logger logger
var Logger *logrus.Entry

// https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
type config struct {
	DbURL string
}

func init() {
	AppEnv = os.Getenv("APP_ENV")
	if AppEnv == "" {
		AppEnv = "development"
	}

	RawLogger = logrus.New()
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	RawLogger.Out = os.Stdout
	if IsProd() {
		RawLogger.Formatter = &logrus.JSONFormatter{}
	}
	if IsDev() {
		RawLogger.Level = logrus.DebugLevel
	} else {
		RawLogger.Level = logrus.InfoLevel
	}
	Logger = RawLogger.WithFields(logrus.Fields{"f1": "t1"})

	if !IsProd() || os.Getenv("USE_DOT_ENV") != "" {
		loadDotEnvFiles()
	}

	Settings = &config{
		DbURL: getEnv("DB_URL", ""),
	}
	Logger.Infof("in env: %s", AppEnv)
	if !IsProd() {
		ConfInfo()
	}
}

// ConfInfo conf
func ConfInfo() {
	tmpl := `
AppEnv: %s
Settings: %+v
`
	Logger.Infof(tmpl, AppEnv, Settings)
}

func loadDotEnvFiles() {
	root := GetAppRoot()
	// top first
	loadDotEnvFile(root + "/.env." + AppEnv + ".local")
	loadDotEnvFile(root + "/.env." + AppEnv)
	loadDotEnvFile(root + "/.env")
}

func loadDotEnvFile(filePath string) {
	if !fileExists(filePath) {
		// Logger.Infof("Not found: %s", filePath)
		return
	}

	if err := godotenv.Load(filePath); err != nil {
		Logger.Errorf("load %s error: %+v", filePath, err)
	} else {
		Logger.Debugf("loaded file: %s", filePath)
	}
}

// GetAppRoot get
func GetAppRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		panic("Can not get working dir")
	}

	for {
		file := filepath.Join(wd, "go.mod")
		if fileExists(file) {
			return wd
		}

		wd = filepath.Dir(wd)

		if "/" == wd {
			panic("Not found root")
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

// IsTest is in testing?
func IsTest() bool {
	return AppEnv == "test"
}

// IsDev is in develop env?
func IsDev() bool {
	return AppEnv == "development" || AppEnv == "dev"
}

// IsProd is in production?
func IsProd() bool {
	return AppEnv == "production" || AppEnv == "prod"
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
