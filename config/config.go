package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
type config struct {
	DbURL string
}

// AppEnv app running env
var AppEnv string

// Settings config from env
var Settings *config

func init() {
	AppEnv = os.Getenv("APP_ENV")
	if AppEnv == "" {
		AppEnv = "development"
	}

	// load config from .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Warning: no .env file found")
	} else {
		log.Println("==loaded .env")
	}

	Settings = &config{
		DbURL: getEnv("DB_URL", ""),
	}
}

// IsTest is in testing?
func IsTest() bool {
	return AppEnv == "testing"
}

// IsDev is in develop env?
func IsDev() bool {
	return AppEnv == "development"
}

// IsProd is in production?
func IsProd() bool {
	return AppEnv == "production"
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// // Simple helper function to read an environment variable into integer or return a default value
// func getEnvAsInt(name string, defaultVal int) int {
// 	valueStr := getEnv(name, "")
// 	if value, err := strconv.Atoi(valueStr); err == nil {
// 		return value
// 	}
//
// 	return defaultVal
// }
//
// // Helper to read an environment variable into a bool or return default value
// func getEnvAsBool(name string, defaultVal bool) bool {
// 	valStr := getEnv(name, "")
// 	if val, err := strconv.ParseBool(valStr); err == nil {
// 		return val
// 	}
//
// 	return defaultVal
// }
