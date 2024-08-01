package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port        string
	PortProduct string
	Domain      string
	Debug       string
	Timeout     string
	Address     string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPass      string
	DbName      string
	Migration   string
}

func NewEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func NewConfig() Config {
	return Config{
		Port:        NewEnv("port"),
		PortProduct: NewEnv("port_product"),
		Domain:      NewEnv("domain"),
		Debug:       NewEnv("debug"),
		Address:     NewEnv("address"),
		DbHost:      NewEnv("db_host"),
		DbPort:      NewEnv("db_port"),
		DbUser:      NewEnv("db_user"),
		DbPass:      NewEnv("db_pass"),
		DbName:      NewEnv("db_name"),
		Migration:   NewEnv("migration"),
	}
}
