package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadConfig () {
	runMode := os.Getenv("RUN_MODE")
    envFile := ".env"

    switch runMode {
    case "staging":
        envFile = ".env.staging"
    case "production":
        envFile = ".env.production"
    default:
        envFile = ".env"
    }
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file", envFile)
	}
}

func GetEnv (key string) string {
	return os.Getenv(key)
}