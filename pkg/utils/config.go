package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func NewEnv(key string) string {
	// read env from os
	env, err := strconv.ParseBool(os.Getenv("ENV_EXTERNAL"))
	if err != nil {
		log.Println("Use External Environment")
	}
	if !env {
		// load .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
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
