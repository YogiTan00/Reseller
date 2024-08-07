package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func NewEnvDefault(key, value string) string {
	env := NewEnv(key)
	if env == "" {
		return value
	}
	return env
}
